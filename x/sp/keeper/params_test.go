package keeper_test

import (
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/x/sp/types"
)

func (s *KeeperTestSuite) TestGetParams() {
	k := s.spKeeper
	ctx := s.ctx
	params := types.DefaultParams()

	err := k.SetParams(ctx, params)
	s.Require().NoError(err)

	require.EqualValues(s.T(), params, k.GetParams(ctx))
}
