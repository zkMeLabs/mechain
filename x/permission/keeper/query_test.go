package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/x/permission/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := makeKeeper(t)
	params := types.DefaultParams()
	err := keeper.SetParams(ctx, params)
	require.NoError(t, err)

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
