package storage

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	mechaincommon "github.com/evmos/evmos/v12/types/common"
	"github.com/evmos/evmos/v12/x/evm/types"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
)

const (
	CreateObjectMethodName     = "createObject"
	SealObjectMethodName       = "sealObject"
	SealObjectV2MethodName     = "sealObjectV2"
	ListObjectsMethodName      = "listObjects"
	HeadObjectMethodName       = "headObject"
	HeadObjectByIdMethodName   = "headObjectById"
	UpdateObjectInfoMethodName = "updateObjectInfo"
)

func (c *Contract) registerObjectMethod() {
	c.registerMethod(ListObjectsMethodName, 50_000, c.ListObjects, "")
	c.registerMethod(HeadObjectMethodName, 50_000, c.HeadObject, "")
	c.registerMethod(HeadObjectByIdMethodName, 50_000, c.HeadObjectById, "")
	c.registerMethod(CreateObjectMethodName, 60_000, c.CreateObject, "CreateObject")
	c.registerMethod(SealObjectMethodName, 100_000, c.SealObject, "SealObject")
	c.registerMethod(SealObjectV2MethodName, 100_000, c.SealObjectV2, "SealObjectV2")
	c.registerMethod(UpdateObjectInfoMethodName, 60_000, c.UpdateObjectInfo, "UpdateObjectInfo")
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
		GlobalVirtualGroupId:        args.GlobalVirtualGroupId,
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

// ListObjects queries the total objects.
func (c *Contract) ListObjects(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ListObjectsMethodName)
	var args ListObjectsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &storagetypes.QueryListObjectsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
		BucketName: args.BucketName,
	}
	res, err := c.storageKeeper.ListObjects(ctx, msg)
	if err != nil {
		return nil, err
	}
	var objectInfos []ObjectInfo
	for _, objectInfo := range res.ObjectInfos {
		objectInfos = append(objectInfos, *outputObjectInfo(objectInfo))
	}
	return method.Outputs.Pack(objectInfos, outputPageResponse(res.Pagination))
}

// HeadObject queries the object's info.
func (c *Contract) HeadObject(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadObjectMethodName)
	var args HeadObjectArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadObjectRequest{
		BucketName: args.BucketName,
		ObjectName: args.ObjectName,
	}
	res, err := c.storageKeeper.HeadObject(ctx, msg)
	if err != nil {
		return nil, err
	}
	objectInfo := outputObjectInfo(res.ObjectInfo)
	gvg := outputsGlobalVirtualGroup(res.GlobalVirtualGroup)
	return method.Outputs.Pack(objectInfo, gvg)
}

// HeadObjectById queries the object's info.
func (c *Contract) HeadObjectById(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadObjectByIdMethodName)
	var args HeadObjectByIdArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadObjectByIdRequest{
		ObjectId: args.ObjectId,
	}
	res, err := c.storageKeeper.HeadObjectById(ctx, msg)
	if err != nil {
		return nil, err
	}
	objectInfo := outputObjectInfo(res.ObjectInfo)
	gvg := outputsGlobalVirtualGroup(res.GlobalVirtualGroup)
	return method.Outputs.Pack(objectInfo, gvg)
}

func outputObjectInfo(o *storagetypes.ObjectInfo) *ObjectInfo {
	n := &ObjectInfo{
		Owner:               common.HexToAddress(o.Owner),
		Creator:             common.HexToAddress(o.Creator),
		BucketName:          o.BucketName,
		ObjectName:          o.ObjectName,
		Id:                  o.Id.BigInt(),
		LocalVirtualGroupId: o.LocalVirtualGroupId,
		PayloadSize:         o.PayloadSize,
		Visibility:          uint8(o.Visibility),
		ContentType:         o.ContentType,
		CreateAt:            o.CreateAt,
		ObjectStatus:        uint8(o.ObjectStatus),
		RedundancyType:      uint8(o.RedundancyType),
		SourceType:          uint8(o.SourceType),
		Checksums:           []string{},
		Tags:                outputTags(o.Tags),
		IsUpdating:          o.IsUpdating,
		UpdatedAt:           o.UpdatedAt,
		UpdatedBy:           common.HexToAddress(o.UpdatedBy),
		Version:             o.Version,
	}
	for i := range o.Checksums {
		n.Checksums = append(n.Checksums, hex.EncodeToString(o.Checksums[i]))
	}
	return n
}
