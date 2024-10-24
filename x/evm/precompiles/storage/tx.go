package storage

import (
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/contracts"
	gtypes "github.com/evmos/evmos/v12/types"
	mechaincommon "github.com/evmos/evmos/v12/types/common"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	CreateBucketMethodName                = "createBucket"
	DeleteBucketMethodName                = "deleteBucket"
	DiscontinueBucketMethodName           = "discontinueBucket"
	CompleteMigrateBucketMethodName       = "completeMigrateBucket"
	RejectMigrateBucketMethodName         = "rejectMigrateBucket"
	CreateObjectMethodName                = "createObject"
	SealObjectMethodName                  = "sealObject"
	SealObjectV2MethodName                = "sealObjectV2"
	RejectSealObjectMethodName            = "rejectSealObject"
	DelegateCreateObjectMethodName        = "delegateCreateObject"
	DelegateUpdateObjectContentMethodName = "delegateUpdateObjectContent"
	UpdateObjectInfoMethodName            = "updateObjectInfo"
	CreateGroupMethodName                 = "createGroup"
	UpdateGroupMethodName                 = "updateGroup"
	DeleteGroupMethodName                 = "deleteGroup"
	RenewGroupMemberMethodName            = "renewGroupMember"
	SetTagForGroupMethodName              = "setTagForGroup"
	UpdateBucketInfoMethodName            = "updateBucketInfo"
)

func (c *Contract) registerTx() {
	c.registerMethod(CreateBucketMethodName, 60_000, c.CreateBucket, "CreateBucket")
	c.registerMethod(DeleteBucketMethodName, 60_000, c.DeleteBucket, "DeleteBucket")
	c.registerMethod(DiscontinueBucketMethodName, 60_000, c.DiscontinueBucket, "DiscontinueBucket")
	c.registerMethod(CompleteMigrateBucketMethodName, 60_000, c.CompleteMigrateBucket, "CompleteMigrateBucket")
	c.registerMethod(RejectMigrateBucketMethodName, 60_000, c.RejectMigrateBucket, "RejectMigrateBucket")
	c.registerMethod(CreateObjectMethodName, 60_000, c.CreateObject, "CreateObject")
	c.registerMethod(SealObjectMethodName, 100_000, c.SealObject, "SealObject")
	c.registerMethod(SealObjectV2MethodName, 100_000, c.SealObjectV2, "SealObjectV2")
	c.registerMethod(RejectSealObjectMethodName, 100_000, c.RejectSealObject, "RejectSealObject")
	c.registerMethod(DelegateCreateObjectMethodName, 100_000, c.DelegateCreateObject, "DelegateCreateObject")
	c.registerMethod(DelegateUpdateObjectContentMethodName, 100_000, c.DelegateUpdateObjectContent, "DelegateUpdateObjectContent")
	c.registerMethod(UpdateObjectInfoMethodName, 60_000, c.UpdateObjectInfo, "UpdateObjectInfo")
	c.registerMethod(CreateGroupMethodName, 60_000, c.CreateGroup, "CreateGroup")
	c.registerMethod(UpdateGroupMethodName, 60_000, c.UpdateGroup, "UpdateGroup")
	c.registerMethod(DeleteGroupMethodName, 60_000, c.DeleteGroup, "DeleteGroup")
	c.registerMethod(RenewGroupMemberMethodName, 60_000, c.RenewGroupMember, "RenewGroupMember")
	c.registerMethod(SetTagForGroupMethodName, 60_000, c.SetTagForGroup, "SetTagForGroup")
	c.registerMethod(UpdateBucketInfoMethodName, 60_000, c.UpdateBucketInfo, "UpdateBucketInfo")
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
		GlobalVirtualGroupFamilyId: args.GlobalVirtualGroupFamilyId,
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

func (c *Contract) SetTagForGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("set tag for group method readonly")
	}
	method := GetAbiMethod(SetTagForGroupMethodName)
	var args SetTagForGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	addr, err := sdk.AccAddressFromHexUnsafe(contract.Caller().String())
	if err != nil {
		return nil, err
	}
	grn := gtypes.NewGroupGRN(addr, args.GroupName)
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
		Resource: grn.String(),
		Tags:     &tags,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _, err = server.SetTag(ctx, msg); err != nil {
		return nil, err
	}
	if err := c.AddLog(
		evm,
		GetAbiEvent(c.events[SetTagForGroupMethodName]),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
	); err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true)
}
