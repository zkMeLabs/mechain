package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/x/challenge/keeper"
	"github.com/evmos/evmos/v12/x/challenge/types"
)

func createSlash(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Slash {
	items := make([]types.Slash, n)
	for i := range items {
		items[i].ObjectId = sdkmath.NewUint(uint64(i))
		items[i].Height = uint64(i)
		items[i].SpId = uint32(i + 1)
		keeper.SaveSlash(ctx, items[i])
	}
	return items
}

func TestRemoveRecentSlash(t *testing.T) {
	keeper, ctx := makeKeeper(t)
	items := createSlash(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSlashUntil(ctx, item.Height)
		found := keeper.ExistsSlash(ctx, item.SpId, item.ObjectId)
		require.False(t, found)
	}
}

func TestRemoveSpSlashAmount(t *testing.T) {
	keeper, ctx := makeKeeper(t)
	keeper.SetSpSlashAmount(ctx, 1, sdk.NewInt(100))
	keeper.SetSpSlashAmount(ctx, 2, sdk.NewInt(200))
	keeper.ClearSpSlashAmount(ctx)
	require.True(t, keeper.GetSpSlashAmount(ctx, 1).Int64() == 0)
	require.True(t, keeper.GetSpSlashAmount(ctx, 2).Int64() == 0)
}
