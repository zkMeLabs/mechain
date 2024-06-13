package storage

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
)

const (
	ListBucketsGas = 50_000

	ListBucketsMethodName = "listBuckets"
)

// ListBuckets queries the total buckets.
func (c *Contract) ListBuckets(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ListBucketsMethodName)

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
		var tags []Tag
		if bucketInfo.Tags != nil {
			for _, tag := range bucketInfo.Tags.Tags {
				tags = append(tags, Tag{
					Key:   tag.Key,
					Value: tag.Value,
				})
			}
		}
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
			Tags:                       tags,
			SpAsDelegatedAgentDisabled: bucketInfo.SpAsDelegatedAgentDisabled,
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(bucketInfos, pageResponse)
}
