package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/testutil/sample"
	"github.com/evmos/evmos/v12/x/challenge/types"
)

func (s *TestSuite) TestUpdateParams() {
	params := types.DefaultParams()
	params.HeartbeatInterval = 10

	tests := []struct {
		name string
		msg  types.MsgUpdateParams
		err  bool
	}{
		{
			name: "invalid authority",
			msg: types.MsgUpdateParams{
				Authority: sample.RandAccAddressHex(),
			},
			err: true,
		}, {
			name: "success",
			msg: types.MsgUpdateParams{
				Authority: s.challengeKeeper.GetAuthority(),
				Params:    params,
			},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			msg := tt.msg
			_, err := s.msgServer.UpdateParams(s.ctx, &msg)
			if tt.err {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}

	// verify storage
	s.Require().Equal(params, s.challengeKeeper.GetParams(s.ctx))
}
