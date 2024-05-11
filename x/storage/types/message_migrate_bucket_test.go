package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/testutil/sample"
	"github.com/evmos/evmos/v12/types/common"
)

func TestMsgMigrateBucket_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMigrateBucket
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgMigrateBucket{
				Operator:   "invalid_address",
				BucketName: "bucketname",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgMigrateBucket{
				Operator:             sample.RandAccAddressHex(),
				BucketName:           "bucketname",
				DstPrimarySpId:       1,
				DstPrimarySpApproval: &common.Approval{ExpiredHeight: 10, Sig: []byte("XXXTentacion")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
