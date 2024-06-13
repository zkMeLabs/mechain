package storage

import (
	"bytes"
	"encoding/base64"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
)

const (
	ListBucketsGas = 50_000
	ListObjectsGas = 50_000

	ListBucketsMethodName = "listBuckets"
	ListObjectsMethodName = "listObjects"
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

// ListObjects queries the total objects.
func (c *Contract) ListObjects(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ListObjectsMethodName)

	// parse args
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
		var tags []Tag

		if objectInfo.Tags != nil {
			for _, tag := range objectInfo.Tags.Tags {
				tags = append(tags, Tag{
					Key:   tag.Key,
					Value: tag.Value,
				})
			}
		}

		var checksums []string
		if objectInfo.Checksums != nil {
			for _, checksum := range objectInfo.Checksums {
				checksums = append(checksums, base64.StdEncoding.EncodeToString(checksum))
			}
		}

		objectInfos = append(objectInfos, ObjectInfo{
			Owner:               common.HexToAddress(objectInfo.Owner),
			Creator:             common.HexToAddress(objectInfo.Creator),
			BucketName:          objectInfo.BucketName,
			ObjectName:          objectInfo.ObjectName,
			Id:                  objectInfo.Id.BigInt(),
			LocalVirtualGroupId: objectInfo.LocalVirtualGroupId,
			PayloadSize:         objectInfo.PayloadSize,
			Visibility:          uint8(objectInfo.PayloadSize),
			ContentType:         objectInfo.ContentType,
			CreateAt:            objectInfo.CreateAt,
			ObjectStatus:        uint8(objectInfo.ObjectStatus),
			RedundancyType:      uint8(objectInfo.RedundancyType),
			SourceType:          uint8(objectInfo.SourceType),
			Checksums:           checksums,
			Tags:                tags,
			IsUpdating:          objectInfo.IsUpdating,
			UpdatedAt:           objectInfo.UpdatedAt,
			UpdatedBy:           common.HexToAddress(objectInfo.UpdatedBy),
			Version:             objectInfo.Version,
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(objectInfos, pageResponse)
}
