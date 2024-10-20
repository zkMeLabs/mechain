package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/testutil/sample"
	gnfderrors "github.com/evmos/evmos/v12/types/errors"
)

func TestMsgSubmit_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmit
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSubmit{
				Challenger: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid bucket name",
			msg: MsgSubmit{
				Challenger:        sample.RandAccAddressHex(),
				SpOperatorAddress: sample.RandAccAddressHex(),
				BucketName:        "1",
			},
			err: gnfderrors.ErrInvalidBucketName,
		}, {
			name: "invalid object name",
			msg: MsgSubmit{
				Challenger:        sample.RandAccAddressHex(),
				SpOperatorAddress: sample.RandAccAddressHex(),
				BucketName:        "bucket",
				ObjectName:        "",
			},
			err: gnfderrors.ErrInvalidObjectName,
		}, {
			name: "valid message with random index",
			msg: MsgSubmit{
				Challenger:        sample.RandAccAddressHex(),
				SpOperatorAddress: sample.RandAccAddressHex(),
				BucketName:        "bucket",
				ObjectName:        "object",
				RandomIndex:       true,
				SegmentIndex:      10,
			},
		}, {
			name: "valid message with specific index",
			msg: MsgSubmit{
				Challenger:        sample.RandAccAddressHex(),
				SpOperatorAddress: sample.RandAccAddressHex(),
				BucketName:        "bucket",
				ObjectName:        "object",
				RandomIndex:       false,
				SegmentIndex:      2,
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
