package storage

import (
	"bytes"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	mechaincommon "github.com/evmos/evmos/v12/types/common"
	"github.com/evmos/evmos/v12/x/evm/types"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
)

const (
	CreateBucketMethodName     = "createBucket"
	ListBucketsMethodName      = "listBuckets"
	HeadBucketMethodName       = "headBucket"
	UpdateBucketInfoMethodName = "updateBucketInfo"
)

func (c *Contract) registerBucketMethod() {
	c.registerMethod(ListBucketsMethodName, 50_000, c.ListBuckets, "")
	c.registerMethod(HeadBucketMethodName, 50_000, c.HeadBucket, "")
	c.registerMethod(CreateBucketMethodName, 60_000, c.CreateBucket, "CreateBucket")
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
		msg.ChargedReadQuota = &mechaincommon.UInt64Value{Value: uint64(args.ChargedReadQuota.Uint64())}
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

// ListBuckets queries the total buckets.
func (c *Contract) ListBuckets(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ListBucketsMethodName)
	// parse args
	var args ListBucketsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &storagetypes.QueryListBucketsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}
	res, err := c.storageKeeper.ListBuckets(ctx, msg)
	if err != nil {
		return nil, err
	}
	var bucketInfos []BucketInfo
	for _, bucketInfo := range res.BucketInfos {
		bucketInfos = append(bucketInfos, BucketInfo{
			Owner:                      common.HexToAddress(bucketInfo.Owner),
			BucketName:                 bucketInfo.BucketName,
			Visibility:                 uint8(bucketInfo.Visibility),
			Id:                         bucketInfo.Id.BigInt(),
			SourceType:                 uint8(bucketInfo.SourceType),
			CreateAt:                   bucketInfo.CreateAt,
			PaymentAddress:             common.HexToAddress(bucketInfo.PaymentAddress),
			GlobalVirtualGroupFamilyId: bucketInfo.GlobalVirtualGroupFamilyId,
			ChargedReadQuota:           bucketInfo.ChargedReadQuota,
			BucketStatus:               uint8(bucketInfo.BucketStatus),
			Tags:                       outputTags(bucketInfo.Tags),
			SpAsDelegatedAgentDisabled: bucketInfo.SpAsDelegatedAgentDisabled,
		})
	}
	return method.Outputs.Pack(bucketInfos, outputPageResponse(res.Pagination))
}

// HeadBucket queries the bucket's info.
func (c *Contract) HeadBucket(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadBucketMethodName)
	var args HeadBucketArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadBucketRequest{
		BucketName: args.BucketName,
	}
	res, err := c.storageKeeper.HeadBucket(ctx, msg)
	if err != nil {
		return nil, err
	}
	bucketInfo := BucketInfo{
		Owner:                      common.HexToAddress(res.BucketInfo.Owner),
		BucketName:                 res.BucketInfo.BucketName,
		Visibility:                 uint8(res.BucketInfo.Visibility),
		Id:                         res.BucketInfo.Id.BigInt(),
		SourceType:                 uint8(res.BucketInfo.SourceType),
		CreateAt:                   res.BucketInfo.CreateAt,
		PaymentAddress:             common.HexToAddress(res.BucketInfo.PaymentAddress),
		GlobalVirtualGroupFamilyId: res.BucketInfo.GlobalVirtualGroupFamilyId,
		ChargedReadQuota:           res.BucketInfo.ChargedReadQuota,
		BucketStatus:               uint8(res.BucketInfo.BucketStatus),
		Tags:                       outputTags(res.BucketInfo.Tags),
		SpAsDelegatedAgentDisabled: res.BucketInfo.SpAsDelegatedAgentDisabled,
	}
	extraInfo := BucketExtraInfo{
		IsRateLimited:   res.ExtraInfo.IsRateLimited,
		FlowRateLimit:   res.ExtraInfo.FlowRateLimit.BigInt(),
		CurrentFlowRate: res.ExtraInfo.CurrentFlowRate.BigInt(),
	}

	return method.Outputs.Pack(bucketInfo, extraInfo)
}
