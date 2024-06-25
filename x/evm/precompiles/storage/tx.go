package storage

import (
	"encoding/base64"
	"errors"

	"github.com/ethereum/go-ethereum/common"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
	mechaincommon "github.com/evmos/evmos/v12/types/common"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	CreateBucketGas = 60_000
	UpdateBucketInfoGas = 60_000
	CreateObjectGas = 60_000
	SealObjectGas   = 100_000
	SealObjectV2Gas = 100_000
	CreateGroupGas  = 60_000

	CreateBucketMethodName = "createBucket"
	UpdateBucketInfoMethodName = "updateBucketInfo"
	CreateObjectMethodName = "createObject"
	SealObjectMethodName   = "sealObject"
	SealObjectV2MethodName = "sealObjectV2"
	CreateGroupMethodName  = "createGroup"

	CreateBucketEventName = "CreateBucket"
	UpdateBucketInfoEventName = "UpdateBucket"
	CreateObjectEventName = "CreateObject"
	SealObjectEventName   = "SealObject"
	SealObjectV2EventName = "SealObjectV2"
	CreateGroupEventName  = "CreateGroup"
)

func (c *Contract) CreateBucket(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("create bucket method readonly")
	}

	method := MustMethod(CreateBucketMethodName)

	var args CreateBucketArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
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

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(CreateBucketEventName),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash(args.PaymentAddress.Bytes()),
			common.BytesToHash(args.PrimarySpAddress.Bytes()),
		},
		res.BucketId.BigInt(),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) UpdateBucketInfo(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("update bucket method readonly")
	}
	method := MustMethod(UpdateBucketInfoMethodName)

	var args UpdateBucketArgs
	if err:= types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil,err
	}
	msg:=&storagetypes.MsgUpdateBucketInfo{
		Operator: contract.CallerAddress.String(),
		BucketName: args.BucketName,
		Visibility: storagetypes.VisibilityType(args.Visibility),
		PaymentAddress: args.PaymentAddress.String(),
		ChargedReadQuota:&mechaincommon.UInt64Value{Value: args.ChargedReadQuota},
	}
	if err:=msg.ValidateBasic();err!=nil {
		return nil,err
	}
	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	if _,err:=server.UpdateBucketInfo(ctx,msg); err!=nil {
		return nil,err
	}
	if err:=c.AddLog(evm, MustEvent(UpdateBucketInfoEventName),[]common.Hash{
		common.BytesToHash(contract.Caller().Bytes()),
		common.BytesToHash([]byte(args.BucketName)),
		common.BytesToHash(args.PaymentAddress.Bytes()),
	},args.Visibility)	;err!=nil {
		return nil,err
	}
	return method.Outputs.Pack(true)
}

func (c *Contract) CreateObject(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("create object method readonly")
	}

	method := MustMethod(CreateObjectMethodName)

	var args CreateObjectArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var expectChecksums [][]byte
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

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(CreateObjectEventName),
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

	method := MustMethod(SealObjectMethodName)

	var args SealObjectArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
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
		GlobalVirtualGroupId:        args.GlobalVirtualGroupId,
		SecondarySpBlsAggSignatures: secondarySpBlsAggSignatures,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.SealObject(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(SealObjectEventName),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash(args.SealAddress.Bytes()),
		},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) SealObjectV2(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("seal object V2 method readonly")
	}

	method := MustMethod(SealObjectV2MethodName)

	var args SealObjectV2Args
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	secondarySpBlsAggSignatures, err := base64.StdEncoding.DecodeString(args.SecondarySpBlsAggSignatures)
	if err != nil {
		return nil, err
	}

	var expectChecksums [][]byte
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
		GlobalVirtualGroupId:        args.GlobalVirtualGroupId,
		SecondarySpBlsAggSignatures: secondarySpBlsAggSignatures,
		ExpectChecksums:             expectChecksums,
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := storagekeeper.NewMsgServerImpl(c.storageKeeper)
	_, err = server.SealObjectV2(ctx, msg)
	if err != nil {
		return nil, err
	}

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(SealObjectV2EventName),
		[]common.Hash{
			common.BytesToHash(contract.Caller().Bytes()),
			common.BytesToHash(args.SealAddress.Bytes()),
		},
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) CreateGroup(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, errors.New("create group method readonly")
	}

	method := MustMethod(CreateGroupMethodName)

	var args CreateGroupArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
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

	// add log
	if err := c.AddLog(
		evm,
		MustEvent(CreateGroupEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.GroupId.BigInt(),
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
