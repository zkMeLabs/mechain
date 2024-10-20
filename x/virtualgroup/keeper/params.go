package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/evmos/evmos/v12/x/virtualgroup/types"
)

func (k Keeper) DepositDenomForGVG(ctx sdk.Context) (res string) {
	params := k.GetParams(ctx)
	return params.DepositDenom
}

func (k Keeper) GVGStakingPerBytes(ctx sdk.Context) (res math.Int) {
	params := k.GetParams(ctx)
	return params.GvgStakingPerBytes
}

func (k Keeper) MaxGlobalVirtualGroupNumPerFamily(ctx sdk.Context) (res uint32) {
	params := k.GetParams(ctx)
	return params.MaxGlobalVirtualGroupNumPerFamily
}

func (k Keeper) MaxStoreSizePerFamily(ctx sdk.Context) (res uint64) {
	params := k.GetParams(ctx)
	return params.MaxStoreSizePerFamily
}

func (k Keeper) SwapInValidityPeriod(ctx sdk.Context) (res uint64) {
	params := k.GetParams(ctx)
	return params.SwapInValidityPeriod.Uint64()
}

func (k Keeper) SpConcurrentExitNum(ctx sdk.Context) (res uint32) {
	params := k.GetParams(ctx)
	return uint32(params.SpConcurrentExitNum.Uint64())
}

// GetParams returns the current sp module parameters.
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return p
	}

	k.cdc.MustUnmarshal(bz, &p)
	return p
}

// SetParams sets the params of sp module
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	if err := params.Validate(); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)

	return nil
}
