package keeper_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/x/storage/keeper"
	"github.com/evmos/evmos/v12/x/storage/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := makeKeeper(t)
	params := types.DefaultParams()

	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	require.EqualValues(t, params, k.GetParams(ctx))
}

func GetVersionedParamsWithTimestamp(k *keeper.Keeper, ctx sdk.Context, ts int64) (val types.VersionedParams) {
	params, err := k.GetVersionedParamsWithTs(ctx, ts)
	if err != nil {
		fmt.Printf("GetParamsWithTimestamp err %s\n", err)
	}
	return params
}

func TestMultiVersionParams(t *testing.T) {
	k, ctx := makeKeeper(t)
	params := types.DefaultParams()

	blockTimeT1 := ctx.BlockTime().Unix()
	params.VersionedParams.MaxSegmentSize = 1
	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(1 * time.Hour))
	blockTimeT2 := ctx.BlockTime().Unix()
	params.VersionedParams.MaxSegmentSize = 2
	err = k.SetParams(ctx, params)
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(1 * time.Hour))
	blockTimeT3 := ctx.BlockTime().Unix()
	params.VersionedParams.MaxSegmentSize = 3
	err = k.SetParams(ctx, params)
	require.NoError(t, err)

	require.EqualValues(t, params, k.GetParams(ctx))
	// default params
	require.EqualValues(t, GetVersionedParamsWithTimestamp(k, ctx, blockTimeT1+1).MaxSegmentSize, 1)
	require.EqualValues(t, GetVersionedParamsWithTimestamp(k, ctx, blockTimeT2+1).MaxSegmentSize, 2)
	require.EqualValues(t, GetVersionedParamsWithTimestamp(k, ctx, blockTimeT3+1).MaxSegmentSize, 3)
}
