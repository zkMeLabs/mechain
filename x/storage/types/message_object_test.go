package types

import (
	"strings"
	"testing"

	"cosmossdk.io/math"
	"github.com/0xPolygon/polygon-edge/bls"
	"github.com/cometbft/cometbft/votepool"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/evmos/evmos/v12/testutil/sample"
	"github.com/evmos/evmos/v12/types/common"
	gnfderrors "github.com/evmos/evmos/v12/types/errors"
)

func TestMsgCreateObject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateObject
		err  error
	}{
		{
			name: "normal",
			msg: MsgCreateObject{
				Creator:           sample.RandAccAddressHex(),
				BucketName:        testBucketName,
				ObjectName:        testObjectName,
				PayloadSize:       1024,
				Visibility:        VISIBILITY_TYPE_PRIVATE,
				ContentType:       "content-type",
				PrimarySpApproval: &common.Approval{},
				ExpectChecksums:   [][]byte{sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum()},
			},
		}, {
			name: "invalid object name",
			msg: MsgCreateObject{
				Creator:           sample.RandAccAddressHex(),
				BucketName:        testBucketName,
				ObjectName:        "",
				PayloadSize:       1024,
				Visibility:        VISIBILITY_TYPE_PRIVATE,
				ContentType:       "content-type",
				PrimarySpApproval: &common.Approval{},
				ExpectChecksums:   [][]byte{sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum()},
			},
			err: gnfderrors.ErrInvalidObjectName,
		}, {
			name: "invalid object name",
			msg: MsgCreateObject{
				Creator:           sample.RandAccAddressHex(),
				BucketName:        testBucketName,
				ObjectName:        "../object",
				PayloadSize:       1024,
				Visibility:        VISIBILITY_TYPE_PRIVATE,
				ContentType:       "content-type",
				PrimarySpApproval: &common.Approval{},
				ExpectChecksums:   [][]byte{sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum()},
			},
			err: gnfderrors.ErrInvalidObjectName,
		}, {
			name: "invalid object name",
			msg: MsgCreateObject{
				Creator:           sample.RandAccAddressHex(),
				BucketName:        testBucketName,
				ObjectName:        "//object",
				PayloadSize:       1024,
				Visibility:        VISIBILITY_TYPE_PRIVATE,
				ContentType:       "content-type",
				PrimarySpApproval: &common.Approval{},
				ExpectChecksums:   [][]byte{sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum()},
			},
			err: gnfderrors.ErrInvalidObjectName,
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

func TestMsgCancelCreateObject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCancelCreateObject
		err  error
	}{
		{
			name: "basic",
			msg: MsgCancelCreateObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectName: testObjectName,
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

func TestMsgDeleteObject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteObject
		err  error
	}{
		{
			name: "normal",
			msg: MsgDeleteObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectName: testObjectName,
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

func TestMsgCopyObject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCopyObject
		err  error
	}{
		{
			name: "valid address",
			msg: MsgCopyObject{
				Operator:      sample.RandAccAddressHex(),
				SrcBucketName: testBucketName,
				SrcObjectName: testObjectName,
				DstBucketName: "dst" + testBucketName,
				DstObjectName: "dst" + testObjectName,
				DstPrimarySpApproval: &common.Approval{
					ExpiredHeight: 100,
					Sig:           []byte("xxx"),
				},
			},
		},
		{
			name: "invalid address",
			msg: MsgCopyObject{
				Operator:      "invalid address",
				SrcBucketName: testBucketName,
				SrcObjectName: testObjectName,
				DstBucketName: "dst" + testBucketName,
				DstObjectName: "dst" + testObjectName,
				DstPrimarySpApproval: &common.Approval{
					ExpiredHeight: 100,
					Sig:           []byte("xxx"),
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "empty approval",
			msg: MsgCopyObject{
				Operator:             sample.RandAccAddressHex(),
				SrcBucketName:        testBucketName,
				SrcObjectName:        testObjectName,
				DstBucketName:        "dst" + testBucketName,
				DstObjectName:        "dst" + testObjectName,
				DstPrimarySpApproval: nil,
			},
			err: ErrInvalidApproval,
		},
		{
			name: "invalid src bucket name",
			msg: MsgCopyObject{
				Operator:      sample.RandAccAddressHex(),
				SrcBucketName: "1.1.1.1",
				SrcObjectName: testObjectName,
				DstBucketName: "dst" + testBucketName,
				DstObjectName: "dst" + testObjectName,
				DstPrimarySpApproval: &common.Approval{
					ExpiredHeight: 100,
					Sig:           []byte("xxx"),
				},
			},
			err: gnfderrors.ErrInvalidBucketName,
		},
		{
			name: "invalid src object name",
			msg: MsgCopyObject{
				Operator:      sample.RandAccAddressHex(),
				SrcBucketName: testBucketName,
				SrcObjectName: "",
				DstBucketName: "dst" + testBucketName,
				DstObjectName: "dst" + testObjectName,
				DstPrimarySpApproval: &common.Approval{
					ExpiredHeight: 100,
					Sig:           []byte("xxx"),
				},
			},
			err: gnfderrors.ErrInvalidObjectName,
		},
		{
			name: "invalid dest bucket name",
			msg: MsgCopyObject{
				Operator:      sample.RandAccAddressHex(),
				SrcBucketName: testBucketName,
				SrcObjectName: testObjectName,
				DstBucketName: "1.1.1.1",
				DstObjectName: "dst" + testObjectName,
				DstPrimarySpApproval: &common.Approval{
					ExpiredHeight: 100,
					Sig:           []byte("xxx"),
				},
			},
			err: gnfderrors.ErrInvalidBucketName,
		},
		{
			name: "invalid dest object name",
			msg: MsgCopyObject{
				Operator:      sample.RandAccAddressHex(),
				SrcBucketName: testBucketName,
				SrcObjectName: testObjectName,
				DstBucketName: "dst" + testBucketName,
				DstObjectName: "",
				DstPrimarySpApproval: &common.Approval{
					ExpiredHeight: 100,
					Sig:           []byte("xxx"),
				},
			},
			err: gnfderrors.ErrInvalidObjectName,
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

func TestMsgSealObject_ValidateBasic(t *testing.T) {
	checksums := [][]byte{sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum(), sample.Checksum()}
	blsSignDoc := NewSecondarySpSealObjectSignDoc("mechain_5151-1", 1, math.NewUint(1), GenerateHash(checksums)).GetSignBytes()
	blsPrivKey, _ := bls.GenerateBlsKey()
	aggSig, _ := blsPrivKey.Sign(blsSignDoc, votepool.DST)
	aggSigBts, _ := aggSig.Marshal()
	tests := []struct {
		name string
		msg  MsgSealObject
		err  error
	}{
		{
			name: "normal",
			msg: MsgSealObject{
				Operator:                    sample.RandAccAddressHex(),
				BucketName:                  testBucketName,
				ObjectName:                  testObjectName,
				SecondarySpBlsAggSignatures: aggSigBts,
			},
		},
		{
			name: "invalid address",
			msg: MsgSealObject{
				Operator:                    "invalid address",
				BucketName:                  testBucketName,
				ObjectName:                  testObjectName,
				SecondarySpBlsAggSignatures: aggSigBts,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid bucket name",
			msg: MsgSealObject{
				Operator:                    sample.RandAccAddressHex(),
				BucketName:                  "1.1.1.1",
				ObjectName:                  testObjectName,
				SecondarySpBlsAggSignatures: aggSigBts,
			},
			err: gnfderrors.ErrInvalidBucketName,
		},
		{
			name: "invalid object name",
			msg: MsgSealObject{
				Operator:                    sample.RandAccAddressHex(),
				BucketName:                  testBucketName,
				ObjectName:                  "",
				SecondarySpBlsAggSignatures: aggSigBts,
			},
			err: gnfderrors.ErrInvalidObjectName,
		},
		{
			name: "invalid signature",
			msg: MsgSealObject{
				Operator:                    sample.RandAccAddressHex(),
				BucketName:                  testBucketName,
				ObjectName:                  testObjectName,
				SecondarySpBlsAggSignatures: []byte("invalid signature"),
			},
			err: gnfderrors.ErrInvalidBlsSignature,
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

func TestMsgRejectSealObject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRejectSealObject
		err  error
	}{
		{
			name: "normal",
			msg: MsgRejectSealObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectName: testObjectName,
			},
		},
		{
			name: "invalid address",
			msg: MsgRejectSealObject{
				Operator:   "invalid address",
				BucketName: "1.1.1.1",
				ObjectName: testObjectName,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid bucket name",
			msg: MsgRejectSealObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: "1.1.1.1",
				ObjectName: testObjectName,
			},
			err: gnfderrors.ErrInvalidBucketName,
		},
		{
			name: "invalid object name",
			msg: MsgRejectSealObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectName: "",
			},
			err: gnfderrors.ErrInvalidObjectName,
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

func TestMsgUpdateObjectInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateObjectInfo
		err  error
	}{
		{
			name: "normal",
			msg: MsgUpdateObjectInfo{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectName: testObjectName,
				Visibility: VISIBILITY_TYPE_INHERIT,
			},
		},
		{
			name: "invalid address",
			msg: MsgUpdateObjectInfo{
				Operator:   "invalid address",
				BucketName: testBucketName,
				ObjectName: testObjectName,
				Visibility: VISIBILITY_TYPE_INHERIT,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid bucket name",
			msg: MsgUpdateObjectInfo{
				Operator:   sample.RandAccAddressHex(),
				BucketName: "1.1.1.1",
				ObjectName: testObjectName,
				Visibility: VISIBILITY_TYPE_INHERIT,
			},
			err: gnfderrors.ErrInvalidBucketName,
		},
		{
			name: "invalid bucket name",
			msg: MsgUpdateObjectInfo{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectName: "",
				Visibility: VISIBILITY_TYPE_INHERIT,
			},
			err: gnfderrors.ErrInvalidObjectName,
		},
		{
			name: "invalid visibility",
			msg: MsgUpdateObjectInfo{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectName: testObjectName,
				Visibility: VISIBILITY_TYPE_UNSPECIFIED,
			},
			err: ErrInvalidVisibility,
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

func TestMsgMirrorObject_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgMirrorObject
		err  error
	}{
		{
			name: "normal",
			msg: MsgMirrorObject{
				Operator: sample.RandAccAddressHex(),
				Id:       math.NewUint(1),
			},
		},
		{
			name: "invalid address",
			msg: MsgMirrorObject{
				Operator: "wrong address",
				Id:       math.NewUint(1),
			},
			err: sdkerrors.ErrInvalidAddress,
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

func TestMsgDiscontinueObject_ValidateBasic(t *testing.T) {
	invalidObjectIDs := [MaxDiscontinueObjects + 1]Uint{}
	tests := []struct {
		name string
		msg  MsgDiscontinueObject
		err  error
	}{
		{
			name: "normal",
			msg: MsgDiscontinueObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectIds:  []Uint{math.NewUint(1)},
				Reason:     "valid reason",
			},
		},
		{
			name: "invalid address",
			msg: MsgDiscontinueObject{
				Operator:   "invalid address",
				BucketName: testBucketName,
				ObjectIds:  []Uint{math.NewUint(1)},
				Reason:     "valid reason",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid bucket name",
			msg: MsgDiscontinueObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: "1.11.1.1",
				ObjectIds:  []Uint{math.NewUint(1)},
				Reason:     "valid reason",
			},
			err: gnfderrors.ErrInvalidBucketName,
		},
		{
			name: "invalid object ids",
			msg: MsgDiscontinueObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectIds:  nil,
				Reason:     "valid reason",
			},
			err: ErrInvalidObjectIDs,
		},
		{
			name: "invalid object ids",
			msg: MsgDiscontinueObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectIds:  invalidObjectIDs[:],
				Reason:     "valid reason",
			},
			err: ErrInvalidObjectIDs,
		},
		{
			name: "invalid reason",
			msg: MsgDiscontinueObject{
				Operator:   sample.RandAccAddressHex(),
				BucketName: testBucketName,
				ObjectIds:  []Uint{math.NewUint(1)},
				Reason:     strings.Repeat("s", MaxDiscontinueReasonLen+1),
			},
			err: ErrInvalidReason,
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
