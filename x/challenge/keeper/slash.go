package keeper

import (
	"encoding/binary"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/evmos/evmos/v12/x/challenge/types"
)

// SaveSlash set a specific slash in the store
func (k Keeper) SaveSlash(ctx sdk.Context, slash types.Slash) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SlashKeyPrefix)

	heightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBytes, slash.Height)

	store.Set(getSlashKeyBytes(slash.SpId, slash.ObjectId), heightBytes)
}

// RemoveSlashUntil removes slashes which are created earlier
func (k Keeper) RemoveSlashUntil(ctx sdk.Context, height uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SlashKeyPrefix)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		slashHeight := binary.BigEndian.Uint64(iterator.Value())
		if slashHeight <= height {
			store.Delete(iterator.Key())
		}
	}
}

// ExistsSlash check whether there exists recent slash for a pair of sp and object info or not
func (k Keeper) ExistsSlash(ctx sdk.Context, spID uint32, objectID sdkmath.Uint) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SlashKeyPrefix)

	return store.Has(getSlashKeyBytes(spID, objectID))
}

// getSlashKeyBytes returns the byte representation of Slash key
func getSlashKeyBytes(spID uint32, objectId sdkmath.Uint) []byte {
	idBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(idBytes, spID)
	allBytes := make([]byte, 0, len(idBytes)+len(objectId.Bytes()))
	copy(allBytes, idBytes)
	allBytes = append(allBytes, objectId.Bytes()...)
	return crypto.Keccak256(allBytes)
}

func (k Keeper) SetSpSlashAmount(ctx sdk.Context, spID uint32, amount sdkmath.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SlashAmountKeyPrefix)
	idBz := make([]byte, 4)
	binary.BigEndian.PutUint32(idBz, spID)
	amountBz, err := amount.Marshal()
	if err != nil {
		panic("cannot marshal amount")
	}
	store.Set(idBz, amountBz)
}

func (k Keeper) GetSpSlashAmount(ctx sdk.Context, spID uint32) sdkmath.Int {
	amount := sdkmath.ZeroInt()

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SlashAmountKeyPrefix)
	idBz := make([]byte, 4)
	binary.BigEndian.PutUint32(idBz, spID)
	amountBz := store.Get(idBz)
	if amountBz == nil {
		return amount
	}
	err := amount.Unmarshal(amountBz)
	if err != nil {
		panic("cannot unmarshal amount")
	}
	return amount
}

func (k Keeper) ClearSpSlashAmount(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SlashAmountKeyPrefix)

	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
}
