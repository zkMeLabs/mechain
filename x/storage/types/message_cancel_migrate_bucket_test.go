package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/testutil/sample"
)

func TestMsgCancelMigrateBucket_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelMigrateBucket
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCancelMigrateBucket{
				Operator:   "invalid_address",
				BucketName: testBucketName,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCancelMigrateBucket{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
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
