package storage

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	cmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/contracts"
	gtypes "github.com/evmos/evmos/v12/types"
	mechaincommon "github.com/evmos/evmos/v12/types/common"
	permTypes "github.com/evmos/evmos/v12/x/permission/types"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	CreateBucketMethodName                = "createBucket"
	DeleteBucketMethodName                = "deleteBucket"
	DiscontinueBucketMethodName           = "discontinueBucket"
	MigrateBucketMethodName               = "migrateBucket"
	CompleteMigrateBucketMethodName       = "completeMigrateBucket"
	RejectMigrateBucketMethodName         = "rejectMigrateBucket"
	CancelMigrateBucketMethodName         = "cancelMigrateBucket"
	SetBucketFlowRateLimitMethodName      = "setBucketFlowRateLimit"
	MirrorBucketMethodName                = "mirrorBucket"
	CreateObjectMethodName                = "createObject"
	CopyObjectMethodName                  = "copyObject"
	DeleteObjectMethodName                = "deleteObject"
	CancelCreateObjectMethodName          = "cancelCreateObject"
	SealObjectMethodName                  = "sealObject"
	SealObjectV2MethodName                = "sealObjectV2"
	RejectSealObjectMethodName            = "rejectSealObject"
	DelegateCreateObjectMethodName        = "delegateCreateObject"
	DelegateUpdateObjectContentMethodName = "delegateUpdateObjectContent"
	UpdateObjectInfoMethodName            = "updateObjectInfo"
	UpdateObjectContentMethodName         = "updateObjectContent"
	DiscontinueObjectMethodName           = "discontinueObject"
	MirrorObjectMethodName                = "mirrorObject"
	CreateGroupMethodName                 = "createGroup"
	UpdateGroupMethodName                 = "updateGroup"
	UpdateGroupExtraMethodName            = "updateGroupExtra"
	DeleteGroupMethodName                 = "deleteGroup"
	LeaveGroupMethodName                  = "leaveGroup"
	RenewGroupMemberMethodName            = "renewGroupMember"
	MirrorGroupMethodName                 = "mirrorGroup"
	SetTagMethodName                      = "setTag"
	UpdateBucketInfoMethodName            = "updateBucketInfo"
	PutPolicyMethodName                   = "putPolicy"
	DeletePolicyMethodName                = "deletePolicy"
	ToggleSPAsDelegatedAgentMethodName    = "toggleSPAsDelegatedAgent"
	UpdateParamsMethodName                = "updateParams"
)

func (c *Contract) registerTx() {
	c.registerMethod(CreateBucketMethodName, 60_000, c.CreateBucket, "CreateBucket")
	c.registerMethod(DeleteBucketMethodName, 60_000, c.DeleteBucket, "DeleteBucket")
	c.registerMethod(DiscontinueBucketMethodName, 60_000, c.DiscontinueBucket, "DiscontinueBucket")
	c.registerMethod(MigrateBucketMethodName, 60_000, c.MigrateBucket, "MigrateBucket")
	c.registerMethod(CompleteMigrateBucketMethodName, 60_000, c.CompleteMigrateBucket, "CompleteMigrateBucket")
	c.registerMethod(RejectMigrateBucketMethodName, 60_000, c.RejectMigrateBucket, "RejectMigrateBucket")
	c.registerMethod(CancelMigrateBucketMethodName, 60_000, c.SetBucketFlowRateLimit, "SetBucketFlowRateLimit")
	c.registerMethod(SetBucketFlowRateLimitMethodName, 60_000, c.CancelMigrateBucket, "CancelMigrateBucket")
	c.registerMethod(MirrorBucketMethodName, 60_000, c.MirrorBucket, "MirrorBucket")
	c.registerMethod(CreateObjectMethodName, 60_000, c.CreateObject, "CreateObject")
	c.registerMethod(CopyObjectMethodName, 60_000, c.CopyObject, "CopyObject")
	c.registerMethod(DeleteObjectMethodName, 60_000, c.DeleteObject, "DeleteObject")
	c.registerMethod(CancelCreateObjectMethodName, 60_000, c.CancelCreateObject, "CancelCreateObject")
	c.registerMethod(SealObjectMethodName, 100_000, c.SealObject, "SealObject")
	c.registerMethod(SealObjectV2MethodName, 100_000, c.SealObjectV2, "SealObjectV2")
	c.registerMethod(RejectSealObjectMethodName, 100_000, c.RejectSealObject, "RejectSealObject")
	c.registerMethod(DelegateCreateObjectMethodName, 100_000, c.DelegateCreateObject, "DelegateCreateObject")
	c.registerMethod(DelegateUpdateObjectContentMethodName, 100_000, c.DelegateUpdateObjectContent, "DelegateUpdateObjectContent")
	c.registerMethod(UpdateObjectInfoMethodName, 60_000, c.UpdateObjectInfo, "UpdateObjectInfo")
	c.registerMethod(UpdateObjectContentMethodName, 60_000, c.UpdateObjectContent, "UpdateObjectContent")
	c.registerMethod(DiscontinueObjectMethodName, 60_000, c.DiscontinueObject, "DiscontinueObject")
	c.registerMethod(MirrorObjectMethodName, 60_000, c.MirrorObject, "MirrorObject")
	c.registerMethod(CreateGroupMethodName, 60_000, c.CreateGroup, "CreateGroup")
	c.registerMethod(UpdateGroupMethodName, 60_000, c.UpdateGroup, "UpdateGroup")
	c.registerMethod(UpdateGroupExtraMethodName, 60_000, c.UpdateGroupExtra, "UpdateGroupExtra")
	c.registerMethod(DeleteGroupMethodName, 60_000, c.DeleteGroup, "DeleteGroup")
	c.registerMethod(LeaveGroupMethodName, 60_000, c.LeaveGroup, "LeaveGroup")
	c.registerMethod(RenewGroupMemberMethodName, 60_000, c.RenewGroupMember, "RenewGroupMember")
	c.registerMethod(MirrorGroupMethodName, 60_000, c.MirrorGroup, "MirrorGroup")
	c.registerMethod(SetTagMethodName, 60_000, c.SetTag, "SetTag")
	c.registerMethod(UpdateBucketInfoMethodName, 60_000, c.UpdateBucketInfo, "UpdateBucketInfo")
	c.registerMethod(PutPolicyMethodName, 60_000, c.PutPolicy, "PutPolicy")
	c.registerMethod(DeletePolicyMethodName, 60_000, c.DeletePolicy, "DeletePolicy")
	c.registerMethod(ToggleSPAsDelegatedAgentMethodName, 60_000, c.ToggleSPAsDelegatedAgent, "ToggleSPAsDelegatedAgent")
	c.registerMethod(UpdateParamsMethodName, 60_000, c.UpdateParams, "UpdateParams")
}

func (c *Contract) CreateBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("create bucket method readonly")
	}
	method := GetAbiMethod(CreateBucketMethodName)
	var args CreateBucketArgs

	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgCreateBucket{
		Creator:          contract.Caller().String(),
		BucketName:       args.BucketName,
		Visibility:       storagetypes.VisibilityType(args.Visibility),
		PaymentAddress:   args.PaymentAddress.String(),
		PrimarySpAddress: args.PrimarySpAddress.String(),
		PrimarySpApproval: &mechaincommon.Approval{
			ExpiredHeight:              args.PrimarySpApproval.ExpiredHeight,
			GlobalVirtualGroupFamilyId: args.PrimarySpApproval.GlobalVirtualGroupFamilyId,
			Sig:                        args.PrimarySpApproval.Sig,
		},
		ChargedReadQuota: args.ChargedReadQuota,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	res, err := server.CreateBucket(ctx, msg)
	if err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CreateBucketMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash(args.PaymentAddress.Bytes()),
			common.BytesToHash(args.PrimarySpAddress.Bytes()),
		},
		res.BucketId.BigInt(),
	); err != nil {
		return nil, err
	}

	bucketInfo, found := c.storageKeeper.GetBucketInfo(ctx, args.BucketName)
	if found {
		if err := c.AddOtherLog(
			evm,
			GetAbiEvent("Transfer"),
			contracts.BucketERC721TokenAddress,
			[]common.Hash{
				common.BytesToHash(common.HexToAddress(gtypes.EmptyEvmAddress).Bytes()),
				common.BytesToHash(common.HexToAddress(bucketInfo.Owner).Bytes()),
				common.BytesToHash(bucketInfo.Id.Bytes()),
			},
		); err != nil {
			return nil, err
		}
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateBucketInfo(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update bucket method readonly")
	}
	method := GetAbiMethod(UpdateBucketInfoMethodName)
	var args UpdateBucketInfoArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgUpdateBucketInfo{
		Operator:       contract.CallerAddress.String(),
		BucketName:     args.BucketName,
		Visibility:     storagetypes.VisibilityType(args.Visibility),
		PaymentAddress: args.PaymentAddress.String(),
	}
	if args.PaymentAddress == (common.Address{}) {
		msg.PaymentAddress = ""
	}
	if args.ChargedReadQuota.Int64() == -1 {
		msg.ChargedReadQuota = nil
	} else {
		msg.ChargedReadQuota = &mechaincommon.UInt64Value{Value: args.ChargedReadQuota.Uint64()}
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.UpdateBucketInfo(ctx, msg); err != nil {
		return nil, err
	}
	bucketNameHash := crypto.Keccak256([]byte(args.BucketName))
	if err := c.AddLog(evm, GetAbiEvent(c.events[UpdateBucketInfoMethodName]), []common.Hash{
		common.BytesToHash(contract.Caller().Bytes()),
		common.BytesToHash(bucketNameHash),
		common.BytesToHash(args.PaymentAddress.Bytes()),
	}, args.Visibility); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) DeleteBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delete bucket method readonly")
	}

	method := GetAbiMethod(DeleteBucketMethodName)

	var args DeleteBucketArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgDeleteBucket{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.DeleteBucket(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DeleteBucketMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) DiscontinueBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("discontinue bucket method readonly")
	}

	method := GetAbiMethod(DiscontinueBucketMethodName)

	var args DiscontinueBucketArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgDiscontinueBucket{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
		Reason:     strings.TrimSpace(args.Reason),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.DiscontinueBucket(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DiscontinueBucketMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.BucketName)),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) MigrateBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("migrate bucket method readonly")
	}

	method := GetAbiMethod(MigrateBucketMethodName)

	var args MigrateBucketArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgMigrateBucket{
		Operator:       contract.Caller().String(),
		BucketName:     args.BucketName,
		DstPrimarySpId: args.DstPrimarySpId,
		DstPrimarySpApproval: &mechaincommon.Approval{
			ExpiredHeight: args.DstPrimarySpApproval.ExpiredHeight,
			// GlobalVirtualGroupFamilyId: args.DstPrimarySpApproval.GlobalVirtualGroupFamilyId,
			Sig: args.DstPrimarySpApproval.Sig,
		},
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.MigrateBucket(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[MigrateBucketMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.BucketName)),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) CompleteMigrateBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("complete migrate bucket method readonly")
	}

	method := GetAbiMethod(CompleteMigrateBucketMethodName)

	var args CompleteMigrateBucketArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}
	gvgMappings := make([]*storagetypes.GVGMapping, 0)
	if args.GvgMappings != nil {
		for _, gvgMapping := range args.GvgMappings {
			gvgMappings = append(gvgMappings, &storagetypes.GVGMapping{
				SrcGlobalVirtualGroupId: gvgMapping.SrcGlobalVirtualGroupId,
				DstGlobalVirtualGroupId: gvgMapping.DstGlobalVirtualGroupId,
				SecondarySpBlsSignature: gvgMapping.SecondarySpBlsSignature,
			})
		}
	}

	msg := &storagetypes.MsgCompleteMigrateBucket{
		Operator:                   contract.Caller().String(),
		BucketName:                 args.BucketName,
		GlobalVirtualGroupFamilyId: args.GvgFamilyId,
		GvgMappings:                gvgMappings,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.CompleteMigrateBucket(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CompleteMigrateBucketMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.BucketName)),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) RejectMigrateBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("reject migrate bucket method readonly")
	}

	method := GetAbiMethod(RejectMigrateBucketMethodName)

	var args RejectMigrateBucketArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgRejectMigrateBucket{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.RejectMigrateBucket(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[RejectMigrateBucketMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) CancelMigrateBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("cancel migrate bucket method readonly")
	}

	method := GetAbiMethod(CancelMigrateBucketMethodName)

	var args CancelMigrateBucketArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgCancelMigrateBucket{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.CancelMigrateBucket(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CancelMigrateBucketMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) SetBucketFlowRateLimit(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("set bucket flow rate limit method readonly")
	}

	method := GetAbiMethod(SetBucketFlowRateLimitMethodName)

	var args SetBucketFlowRateLimitArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgSetBucketFlowRateLimit{
		Operator:       contract.Caller().String(),
		BucketName:     args.BucketName,
		BucketOwner:    args.BucketOwner,
		PaymentAddress: args.PaymentAddress,
		FlowRateLimit:  cmath.NewIntFromBigInt(args.FlowRateLimit),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.SetBucketFlowRateLimit(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[SetBucketFlowRateLimitMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) MirrorBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("mirror bucket method readonly")
	}

	method := GetAbiMethod(MirrorBucketMethodName)

	var args MirrorBucketArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgMirrorBucket{
		Operator:    contract.Caller().String(),
		Id:          cmath.NewUintFromBigInt(args.BucketId),
		BucketName:  args.BucketName,
		DestChainId: args.DestChainId,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.MirrorBucket(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[MirrorBucketMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) CreateObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("create object method readonly")
	}
	method := GetAbiMethod(CreateObjectMethodName)
	var args CreateObjectArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	expectChecksums := make([][]byte, 0)
	for _, checksum := range args.ExpectChecksums {
		checksumBytes, err := base64.StdEncoding.DecodeString(checksum)
		if err != nil {
			return nil, err
		}
		expectChecksums = append(expectChecksums, checksumBytes)
	}
	msg := &storagetypes.MsgCreateObject{
		Creator:     contract.Caller().String(),
		BucketName:  args.BucketName,
		ObjectName:  args.ObjectName,
		PayloadSize: args.PayloadSize,
		Visibility:  storagetypes.VisibilityType(args.Visibility),
		ContentType: args.ContentType,
		PrimarySpApproval: &mechaincommon.Approval{
			ExpiredHeight:              args.PrimarySpApproval.ExpiredHeight,
			GlobalVirtualGroupFamilyId: args.PrimarySpApproval.GlobalVirtualGroupFamilyId,
			Sig:                        args.PrimarySpApproval.Sig,
		},
		ExpectChecksums: expectChecksums,
		RedundancyType:  storagetypes.RedundancyType(args.RedundancyType),
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	res, err := server.CreateObject(ctx, msg)
	if err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CreateObjectMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.ObjectId.BigInt(),
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) CopyObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("copy object method readonly")
	}
	method := GetAbiMethod(CopyObjectMethodName)
	var args CopyObjectArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgCopyObject{
		Operator:      contract.Caller().String(),
		SrcBucketName: args.SrcBucketName,
		DstBucketName: args.DstBucketName,
		SrcObjectName: args.SrcObjectName,
		DstObjectName: args.DstObjectName,
		DstPrimarySpApproval: &mechaincommon.Approval{
			ExpiredHeight: args.DstPrimarySpApproval.ExpiredHeight,
			// GlobalVirtualGroupFamilyId: args.DstPrimarySpApproval.GlobalVirtualGroupFamilyId,
			Sig: args.DstPrimarySpApproval.Sig,
		},
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err := server.CopyObject(ctx, msg)
	if err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CopyObjectMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) DeleteObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delete object method readonly")
	}

	method := GetAbiMethod(DeleteObjectMethodName)

	var args DeleteObjectArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgDeleteObject{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
		ObjectName: args.ObjectName,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.DeleteObject(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DeleteObjectMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) CancelCreateObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("cancel create object method readonly")
	}

	method := GetAbiMethod(CancelCreateObjectMethodName)

	var args CancelCreateObjectArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgCancelCreateObject{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
		ObjectName: args.ObjectName,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.CancelCreateObject(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CancelCreateObjectMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) SealObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("seal object method readonly")
	}
	method := GetAbiMethod(SealObjectMethodName)
	var args SealObjectArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	secondarySpBlsAggSignatures, err := base64.StdEncoding.DecodeString(args.SecondarySpBlsAggSignatures)
	if err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgSealObject{
		Operator:                    args.SealAddress.String(),
		BucketName:                  args.BucketName,
		ObjectName:                  args.ObjectName,
		GlobalVirtualGroupId:        args.GlobalVirtualGroupID,
		SecondarySpBlsAggSignatures: secondarySpBlsAggSignatures,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err = server.SealObject(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[SealObjectMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash(args.SealAddress.Bytes()),
		},
	); err != nil {
		return nil, err
	}

	objectInfo, found := c.storageKeeper.GetObjectInfo(ctx, args.BucketName, args.ObjectName)
	if found {
		if err := c.AddOtherLog(
			evm,
			GetAbiEvent("Transfer"),
			contracts.ObjectERC721TokenAddress,
			[]common.Hash{
				common.BytesToHash(common.HexToAddress(gtypes.EmptyEvmAddress).Bytes()),
				common.BytesToHash(common.HexToAddress(objectInfo.Owner).Bytes()),
				common.BytesToHash(objectInfo.Id.Bytes()),
			},
		); err != nil {
			return nil, err
		}
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) SealObjectV2(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("seal object V2 method readonly")
	}
	method := GetAbiMethod(SealObjectV2MethodName)
	var args SealObjectV2Args
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	secondarySpBlsAggSignatures, err := base64.StdEncoding.DecodeString(args.SecondarySpBlsAggSignatures)
	if err != nil {
		return nil, err
	}
	expectChecksums := make([][]byte, 0)
	for _, checksum := range args.ExpectChecksums {
		checksumBytes, err := base64.StdEncoding.DecodeString(checksum)
		if err != nil {
			return nil, err
		}
		expectChecksums = append(expectChecksums, checksumBytes)
	}
	msg := &storagetypes.MsgSealObjectV2{
		Operator:                    args.SealAddress.String(),
		BucketName:                  args.BucketName,
		ObjectName:                  args.ObjectName,
		GlobalVirtualGroupId:        args.GlobalVirtualGroupID,
		SecondarySpBlsAggSignatures: secondarySpBlsAggSignatures,
		ExpectChecksums:             expectChecksums,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err = server.SealObjectV2(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[SealObjectV2MethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash(args.SealAddress.Bytes()),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) RejectSealObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("reject seal object method readonly")
	}
	method := GetAbiMethod(RejectSealObjectMethodName)
	var args RejectSealObjectArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgRejectSealObject{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
		ObjectName: args.ObjectName,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.RejectSealObject(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[RejectSealObjectMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.ObjectName)),
		},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) DelegateCreateObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delegate create object method readonly")
	}
	method := GetAbiMethod(DelegateCreateObjectMethodName)
	var args DelegateCreateObjectArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	expectChecksums := make([][]byte, 0)
	for _, checksum := range args.ExpectChecksums {
		checksumBytes, err := base64.StdEncoding.DecodeString(checksum)
		if err != nil {
			return nil, err
		}
		expectChecksums = append(expectChecksums, checksumBytes)
	}

	msg := &storagetypes.MsgDelegateCreateObject{
		Operator:        contract.Caller().String(),
		Creator:         args.Creator,
		BucketName:      args.BucketName,
		ObjectName:      args.ObjectName,
		PayloadSize:     args.PayloadSize,
		ContentType:     args.ContentType,
		Visibility:      storagetypes.VisibilityType(args.Visibility),
		ExpectChecksums: expectChecksums,
		RedundancyType:  storagetypes.RedundancyType(args.RedundancyType),
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.DelegateCreateObject(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DelegateCreateObjectMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.ObjectName)),
		},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) DelegateUpdateObjectContent(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delegate update object content method readonly")
	}
	method := GetAbiMethod(DelegateUpdateObjectContentMethodName)
	var args DelegateUpdateObjectContentArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	expectChecksums := make([][]byte, 0)
	for _, checksum := range args.ExpectChecksums {
		checksumBytes, err := base64.StdEncoding.DecodeString(checksum)
		if err != nil {
			return nil, err
		}
		expectChecksums = append(expectChecksums, checksumBytes)
	}

	msg := &storagetypes.MsgDelegateUpdateObjectContent{
		Operator:        contract.Caller().String(),
		Updater:         args.Updater,
		BucketName:      args.BucketName,
		ObjectName:      args.ObjectName,
		PayloadSize:     args.PayloadSize,
		ContentType:     args.ContentType,
		ExpectChecksums: expectChecksums,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.DelegateUpdateObjectContent(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DelegateUpdateObjectContentMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.ObjectName)),
		},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateObjectInfo(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update object info method readonly")
	}
	method := GetAbiMethod(UpdateObjectInfoMethodName)
	var args UpdateObjectInfoArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgUpdateObjectInfo{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
		ObjectName: args.ObjectName,
		Visibility: storagetypes.VisibilityType(args.Visibility),
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.UpdateObjectInfo(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[UpdateObjectInfoMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateObjectContent(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update object content method readonly")
	}
	method := GetAbiMethod(UpdateObjectContentMethodName)
	var args UpdateObjectContentArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	expectChecksums := make([][]byte, 0)
	for _, checksum := range args.ExpectChecksums {
		checksumBytes, err := base64.StdEncoding.DecodeString(checksum)
		if err != nil {
			return nil, err
		}
		expectChecksums = append(expectChecksums, checksumBytes)
	}
	msg := &storagetypes.MsgUpdateObjectContent{
		Operator:        contract.Caller().String(),
		BucketName:      args.BucketName,
		ObjectName:      args.ObjectName,
		PayloadSize:     args.PayloadSize,
		ContentType:     args.ContentType,
		ExpectChecksums: expectChecksums,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.UpdateObjectContent(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[UpdateObjectContentMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.ObjectName)),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) DiscontinueObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("discontinue object method readonly")
	}

	method := GetAbiMethod(DiscontinueObjectMethodName)

	var args DiscontinueObjectArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	objectIDs := make([]storagetypes.Uint, 0)
	for _, id := range args.ObjectIds {
		if id.Cmp(big.NewInt(0)) < 0 {
			return nil, fmt.Errorf("object id should not be negative")
		}

		objectIDs = append(objectIDs, cmath.NewUintFromBigInt(id))
	}
	msg := &storagetypes.MsgDiscontinueObject{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
		ObjectIds:  objectIDs,
		Reason:     strings.TrimSpace(args.Reason),
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.DiscontinueObject(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DiscontinueObjectMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.BucketName)),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) MirrorObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("mirror object method readonly")
	}

	method := GetAbiMethod(MirrorObjectMethodName)

	var args MirrorObjectArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgMirrorObject{
		Operator:    contract.Caller().String(),
		Id:          cmath.NewUintFromBigInt(args.ObjectId),
		BucketName:  args.BucketName,
		ObjectName:  args.ObjectName,
		DestChainId: args.DestChainId,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.MirrorObject(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[MirrorObjectMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) CreateGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("create group method readonly")
	}
	method := GetAbiMethod(CreateGroupMethodName)
	var args CreateGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgCreateGroup{
		Creator:   contract.Caller().String(),
		GroupName: args.GroupName,
		Extra:     args.Extra,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	res, err := server.CreateGroup(ctx, msg)
	if err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[CreateGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.GroupId.BigInt(),
	); err != nil {
		return nil, err
	}

	address := sdk.MustAccAddressFromHex(contract.Caller().String())
	groupInfo, found := c.storageKeeper.GetGroupInfo(ctx, address, args.GroupName)
	if found {
		if err := c.AddOtherLog(
			evm,
			GetAbiEvent("Transfer"),
			contracts.GroupERC721TokenAddress,
			[]common.Hash{
				common.BytesToHash(common.HexToAddress(gtypes.EmptyEvmAddress).Bytes()),
				common.BytesToHash(common.HexToAddress(groupInfo.Owner).Bytes()),
				common.BytesToHash(groupInfo.Id.Bytes()),
			},
		); err != nil {
			return nil, err
		}
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update group method readonly")
	}
	method := GetAbiMethod(UpdateGroupMethodName)
	var args UpdateGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	membersToAdd := make([]*storagetypes.MsgGroupMember, 0)
	if args.MembersToAdd != nil {
		for i, members := range args.MembersToAdd {
			var exp time.Time
			if args.ExpirationTime[i] != 0 {
				exp = time.Unix(args.ExpirationTime[i], 0)
			} else {
				exp = storagetypes.MaxTimeStamp
			}
			membersToAdd = append(membersToAdd, &storagetypes.MsgGroupMember{
				Member:         members.String(),
				ExpirationTime: &exp,
			})
		}
	}
	var membersToDelete []string
	if args.MembersToDelete != nil {
		for _, members := range args.MembersToDelete {
			membersToDelete = append(membersToDelete, members.String())
		}
	}
	msg := &storagetypes.MsgUpdateGroupMember{
		Operator:        contract.Caller().String(),
		GroupOwner:      args.GroupOwner.String(),
		GroupName:       args.GroupName,
		MembersToAdd:    membersToAdd,
		MembersToDelete: membersToDelete,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)

	if _, err := server.UpdateGroupMember(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[UpdateGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateGroupExtra(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update group extra method readonly")
	}
	method := GetAbiMethod(UpdateGroupExtraMethodName)
	var args UpdateGroupExtraArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgUpdateGroupExtra{
		Operator:   contract.Caller().String(),
		GroupOwner: args.GroupOwner.String(),
		GroupName:  args.GroupName,
		Extra:      args.Extra,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)

	if _, err := server.UpdateGroupExtra(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[UpdateGroupExtraMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) DeleteGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delete group method readonly")
	}
	method := GetAbiMethod(DeleteGroupMethodName)
	var args DeleteGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgDeleteGroup{
		Operator:  contract.Caller().String(),
		GroupName: args.GroupName,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.DeleteGroup(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DeleteGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) LeaveGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("leave group method readonly")
	}
	method := GetAbiMethod(LeaveGroupMethodName)
	var args LeaveGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.MsgLeaveGroup{
		Member:     args.Member.String(),
		GroupOwner: args.GroupOwner.String(),
		GroupName:  args.GroupName,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.LeaveGroup(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[LeaveGroupMethodName]),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash([]byte(args.GroupName)),
		},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) RenewGroupMember(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("renew group member method readonly")
	}
	method := GetAbiMethod(RenewGroupMemberMethodName)
	var args RenewGroupMemberArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if args.GroupName == "" {
		return nil, errors.New("group name is empty")
	}
	if len(args.Members) == 0 {
		return nil, errors.New("no renew member")
	}
	if args.ExpirationTime != nil && len(args.Members) != len(args.ExpirationTime) {
		return nil, errors.New("please provide expirationTime for every renew member")
	}
	membersToRenew := make([]*storagetypes.MsgGroupMember, 0)
	if args.Members != nil {
		for i, members := range args.Members {
			var exp time.Time
			if args.ExpirationTime[i] != 0 {
				exp = time.Unix(args.ExpirationTime[i], 0)
			} else {
				exp = storagetypes.MaxTimeStamp
			}
			membersToRenew = append(membersToRenew, &storagetypes.MsgGroupMember{
				Member:         members.String(),
				ExpirationTime: &exp,
			})
		}
	}
	msg := &storagetypes.MsgRenewGroupMember{
		Operator:   contract.Caller().String(),
		GroupOwner: args.GroupOwner.String(),
		GroupName:  args.GroupName,
		Members:    membersToRenew,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.RenewGroupMember(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[RenewGroupMemberMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) MirrorGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("mirror group method readonly")
	}

	method := GetAbiMethod(MirrorGroupMethodName)

	var args MirrorGroupArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgMirrorGroup{
		Operator:    contract.Caller().String(),
		Id:          cmath.NewUintFromBigInt(args.GroupId),
		GroupName:   args.GroupName,
		DestChainId: args.DestChainId,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.MirrorGroup(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[MirrorGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) SetTag(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("set tag for group method readonly")
	}
	method := GetAbiMethod(SetTagMethodName)
	var args SetTagArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	var tags storagetypes.ResourceTags
	if args.Tags != nil {
		for _, tag := range args.Tags {
			tags.Tags = append(tags.Tags, storagetypes.ResourceTags_Tag{
				Key:   tag.Key,
				Value: tag.Value,
			})
		}
	}
	msg := &storagetypes.MsgSetTag{
		Operator: contract.Caller().String(),
		Resource: args.Resource,
		Tags:     &tags,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.SetTag(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[SetTagMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func parseResourceType(resource string) (ResourceType, error) {
	var resourceType ResourceType
	if strings.HasPrefix(resource, BucketResourcePrefix) {
		resourceType = BucketResourceType
	} else if strings.HasPrefix(resource, ObjectResourcePrefix) {
		resourceType = ObjectResourceType
	} else if strings.HasPrefix(resource, GroupResourcePrefix) {
		resourceType = GroupResourceType
	} else {
		return -1, toCmdErr(errors.New("invalid resource name"))
	}
	return resourceType, nil
}

func parseBucketResource(resourceName string) (string, error) {
	prefixLen := len(BucketResourcePrefix)
	if len(resourceName) <= prefixLen {
		return "", errors.New("invalid bucket resource name")
	}

	return resourceName[prefixLen:], nil
}

func toCmdErr(err error) error {
	if strings.Contains(err.Error(), noBalanceErr) {
		fmt.Println("The operator account have no balance, please transfer token to your account")
	} else {
		fmt.Printf("run command error: %s\n", err.Error())
	}
	return nil
}

func getBucketAction(action string) (permTypes.ActionType, bool, error) {
	switch action {
	case "update":
		return permTypes.ACTION_UPDATE_BUCKET_INFO, false, nil
	case "delete":
		return permTypes.ACTION_DELETE_BUCKET, false, nil
	case "list":
		return permTypes.ACTION_LIST_OBJECT, false, nil
	case "createObj":
		return permTypes.ACTION_CREATE_OBJECT, true, nil
	case "deleteObj":
		return permTypes.ACTION_DELETE_OBJECT, true, nil
	case "copyObj":
		return permTypes.ACTION_COPY_OBJECT, true, nil
	case "getObj":
		return permTypes.ACTION_GET_OBJECT, true, nil
	case "executeObj":
		return permTypes.ACTION_EXECUTE_OBJECT, true, nil
	case "all":
		return permTypes.ACTION_TYPE_ALL, true, nil
	default:
		return permTypes.ACTION_UNSPECIFIED, false, errors.New("invalid action :" + action)
	}
}

func getObjectAction(action string) (permTypes.ActionType, error) {
	switch action {
	case "create":
		return permTypes.ACTION_CREATE_OBJECT, nil
	case "delete":
		return permTypes.ACTION_DELETE_OBJECT, nil
	case "copy":
		return permTypes.ACTION_COPY_OBJECT, nil
	case "get":
		return permTypes.ACTION_GET_OBJECT, nil
	case "execute":
		return permTypes.ACTION_EXECUTE_OBJECT, nil
	case "list":
		return permTypes.ACTION_LIST_OBJECT, nil
	case "update":
		return permTypes.ACTION_UPDATE_OBJECT_INFO, nil
	case "all":
		return permTypes.ACTION_TYPE_ALL, nil
	default:
		return permTypes.ACTION_UNSPECIFIED, errors.New("invalid action:" + action)
	}
}

func getGroupAction(action string) (permTypes.ActionType, error) {
	switch action {
	case "update":
		return permTypes.ACTION_UPDATE_GROUP_MEMBER, nil
	case "delete":
		return permTypes.ACTION_DELETE_GROUP, nil
	case "all":
		return permTypes.ACTION_TYPE_ALL, nil
	default:
		return permTypes.ACTION_UNSPECIFIED, errors.New("invalid action:" + action)
	}
}

func parseActions(actionListStr string, resourceType ResourceType) ([]permTypes.ActionType, bool, error) {
	actions := make([]permTypes.ActionType, 0)
	if actionListStr == "" {
		return nil, false, errors.New("fail to parse actions")
	}

	actionList := strings.Split(actionListStr, ",")
	var isObjectActionInBucketPolicy bool
	for _, v := range actionList {
		var action permTypes.ActionType
		var err error
		if resourceType == ObjectResourceType {
			action, err = getObjectAction(v)
		} else if resourceType == BucketResourceType {
			action, isObjectActionInBucketPolicy, err = getBucketAction(v)
		} else if resourceType == GroupResourceType {
			action, err = getGroupAction(v)
		}

		if err != nil {
			return nil, isObjectActionInBucketPolicy, err
		}
		actions = append(actions, action)
	}

	return actions, isObjectActionInBucketPolicy, nil
}

func NewStatement(actions []permTypes.ActionType, effect permTypes.Effect,
	resource []string, opts NewStatementOptions,
) permTypes.Statement {
	statement := permTypes.Statement{
		Actions:        actions,
		Effect:         effect,
		Resources:      resource,
		ExpirationTime: opts.StatementExpireTime,
	}

	if opts.LimitSize != 0 {
		statement.LimitSize = &mechaincommon.UInt64Value{Value: opts.LimitSize}
	}

	return statement
}

func (c *Contract) PutPolicy(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("put policy to group or account method readonly")
	}
	method := GetAbiMethod(PutPolicyMethodName)
	var args PutPolicyArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	statements := make([]*permTypes.Statement, 0)
	if args.Statements != nil {
		for _, statement := range args.Statements {
			actions := make([]permTypes.ActionType, 0)
			for _, action := range statement.Actions {
				actions = append(actions, permTypes.ActionType(action))
			}
			s := &permTypes.Statement{
				Effect:    permTypes.Effect(statement.Effect),
				Actions:   actions,
				Resources: statement.Resources,
			}
			if statement.ExpirationTime != 0 {
				tm := time.Unix(statement.ExpirationTime, 0)
				s.ExpirationTime = &tm
			}
			if statement.LimitSize != 0 {
				s.LimitSize = &mechaincommon.UInt64Value{Value: statement.LimitSize}
			}
			statements = append(statements, s)
		}
	}

	var tmptr *time.Time
	if args.ExpirationTime != 0 {
		tm := time.Unix(args.ExpirationTime, 0)
		tmptr = &tm
	}

	msg := &storagetypes.MsgPutPolicy{
		Operator:       contract.Caller().String(),
		Principal:      &permTypes.Principal{Type: permTypes.PrincipalType(args.Principal.PrincipalType), Value: args.Principal.Value},
		Resource:       args.Resource,
		Statements:     statements,
		ExpirationTime: tmptr,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.PutPolicy(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[PutPolicyMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) DeletePolicy(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("delete policy of principal method readonly")
	}
	method := GetAbiMethod(DeletePolicyMethodName)
	var args DeletePolicyArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgDeletePolicy{
		Operator:  contract.Caller().String(),
		Principal: &permTypes.Principal{Type: permTypes.PrincipalType(args.Principal.PrincipalType), Value: args.Principal.Value},
		Resource:  args.Resource,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.DeletePolicy(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[DeletePolicyMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) ToggleSPAsDelegatedAgent(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("toggle SP as delegated agent method readonly")
	}
	method := GetAbiMethod(ToggleSPAsDelegatedAgentMethodName)
	var args ToggleSPAsDelegatedAgentArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgToggleSPAsDelegatedAgent{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.ToggleSPAsDelegatedAgent(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[ToggleSPAsDelegatedAgentMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateParams(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update params of modular storage method readonly")
	}
	method := GetAbiMethod(UpdateParamsMethodName)
	var args UpdateParamsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &storagetypes.MsgUpdateParams{
		Authority: args.Authority,
		Params: storagetypes.Params{
			VersionedParams: storagetypes.VersionedParams{
				MaxSegmentSize:          args.Params.VersionedParams.MaxSegmentSize,
				RedundantDataChunkNum:   args.Params.VersionedParams.RedundantDataChunkNum,
				RedundantParityChunkNum: args.Params.VersionedParams.RedundantParityChunkNum,
				MinChargeSize:           args.Params.VersionedParams.MinChargeSize,
			},
			MaxPayloadSize:                   args.Params.MaxPayloadSize,
			BscMirrorBucketRelayerFee:        args.Params.BscMirrorBucketRelayerFee,
			BscMirrorBucketAckRelayerFee:     args.Params.BscMirrorBucketAckRelayerFee,
			BscMirrorObjectRelayerFee:        args.Params.BscMirrorObjectRelayerFee,
			BscMirrorObjectAckRelayerFee:     args.Params.BscMirrorObjectAckRelayerFee,
			BscMirrorGroupRelayerFee:         args.Params.BscMirrorGroupRelayerFee,
			BscMirrorGroupAckRelayerFee:      args.Params.BscMirrorGroupAckRelayerFee,
			MaxBucketsPerAccount:             args.Params.MaxBucketsPerAccount,
			DiscontinueCountingWindow:        args.Params.DiscontinueCountingWindow,
			DiscontinueObjectMax:             args.Params.DiscontinueObjectMax,
			DiscontinueBucketMax:             args.Params.DiscontinueBucketMax,
			DiscontinueConfirmPeriod:         args.Params.DiscontinueConfirmPeriod,
			DiscontinueDeletionMax:           args.Params.DiscontinueDeletionMax,
			StalePolicyCleanupMax:            args.Params.StalePolicyCleanupMax,
			MinQuotaUpdateInterval:           args.Params.MinQuotaUpdateInterval,
			MaxLocalVirtualGroupNumPerBucket: args.Params.MaxLocalVirtualGroupNumPerBucket,
			OpMirrorBucketRelayerFee:         args.Params.OpMirrorBucketRelayerFee,
			OpMirrorBucketAckRelayerFee:      args.Params.OpMirrorBucketAckRelayerFee,
			OpMirrorObjectRelayerFee:         args.Params.OpMirrorObjectRelayerFee,
			OpMirrorObjectAckRelayerFee:      args.Params.OpMirrorObjectAckRelayerFee,
			OpMirrorGroupRelayerFee:          args.Params.OpMirrorGroupRelayerFee,
			OpMirrorGroupAckRelayerFee:       args.Params.OpMirrorGroupAckRelayerFee,
			PolygonMirrorBucketRelayerFee:    args.Params.PolygonMirrorBucketRelayerFee,
			PolygonMirrorBucketAckRelayerFee: args.Params.PolygonMirrorBucketAckRelayerFee,
		},
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err := server.UpdateParams(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[UpdateParamsMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}
