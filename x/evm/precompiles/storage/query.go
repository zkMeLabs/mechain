package storage

import (
	"bytes"
	"encoding/hex"

	cmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
	permissiontypes "github.com/evmos/evmos/v12/x/permission/types"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
	vgtypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
)

const (
	ListBucketsMethodName                            = "listBuckets"
	ListObjectsMethodName                            = "listObjects"
	ListGroupsMethodName                             = "listGroups"
	ListObjectsByBucketIdMethodName                  = "listObjectsByBucketId"
	HeadBucketMethodName                             = "headBucket"
	HeadGroupMethodName                              = "headGroup"
	HeadGroupMemberMethodName                        = "headGroupMember"
	HeadObjectMethodName                             = "headObject"
	HeadObjectByIdMethodName                         = "headObjectById"
	HeadBucketByIdMethodName                         = "headBucketById"
	HeadBucketNFTMethodName                          = "headBucketNFT"
	HeadShadowObjectMethodName                       = "headShadowObject"
	HeadObjectNFTMethodName                          = "headObjectNFT"
	HeadGroupNFTMethodName                           = "headGroupNFT"
	HeadBucketExtraMethodName                        = "headBucketExtra"
	QueryPolicyForGroupMethodName                    = "queryPolicyForGroup"
	QueryPolicyForAccountMethodName                  = "queryPolicyForAccount"
	QueryParamsByTimestampMethodName                 = "queryParamsByTimestamp"
	QueryPolicyByIdMethodName                        = "queryPolicyById"
	QueryLockFeeMethodName                           = "queryLockFee"
	QueryIsPriceChangedMethodName                    = "queryIsPriceChanged"
	QueryQuotaUpdateTimeMethodName                   = "queryQuotaUpdateTime"
	QueryGroupMembersExistMethodName                 = "queryGroupMembersExist"
	QueryGroupsExistMethodName                       = "queryGroupsExist"
	QueryGroupsExistByIdMethodName                   = "queryGroupsExistById"
	QueryPaymentAccountBucketFlowRateLimitMethodName = "queryPaymentAccountBucketFlowRateLimit"
	ParamsMethodName                                 = "params"
	VerifyPermissionMethodName                       = "verifyPermission"
)

func (c *Contract) registerQuery() {
	c.registerMethod(ListBucketsMethodName, 50_000, c.ListBuckets, "")
	c.registerMethod(ListObjectsMethodName, 50_000, c.ListObjects, "")
	c.registerMethod(ListGroupsMethodName, 50_000, c.ListGroups, "")
	c.registerMethod(ListObjectsByBucketIdMethodName, 50_000, c.ListObjectsByBucketId, "")
	c.registerMethod(HeadBucketMethodName, 50_000, c.HeadBucket, "")
	c.registerMethod(HeadGroupMethodName, 50_000, c.HeadGroup, "")
	c.registerMethod(HeadGroupMemberMethodName, 50_000, c.HeadGroupMember, "")
	c.registerMethod(HeadObjectMethodName, 50_000, c.HeadObject, "")
	c.registerMethod(HeadObjectByIdMethodName, 50_000, c.HeadObjectById, "")
	c.registerMethod(HeadBucketByIdMethodName, 50_000, c.HeadBucketById, "")
	c.registerMethod(HeadBucketNFTMethodName, 50_000, c.HeadBucketNFT, "")
	c.registerMethod(HeadShadowObjectMethodName, 50_000, c.HeadShadowObject, "")
	c.registerMethod(HeadObjectNFTMethodName, 50_000, c.HeadObjectNFT, "")
	c.registerMethod(HeadGroupNFTMethodName, 50_000, c.HeadGroupNFT, "")
	c.registerMethod(HeadBucketExtraMethodName, 50_000, c.HeadBucketExtra, "")
	c.registerMethod(QueryPolicyForGroupMethodName, 50_000, c.QueryPolicyForGroup, "")
	c.registerMethod(QueryPolicyForAccountMethodName, 50_000, c.QueryPolicyForAccount, "")
	c.registerMethod(QueryParamsByTimestampMethodName, 50_000, c.QueryParamsByTimestamp, "")
	c.registerMethod(QueryPolicyByIdMethodName, 50_000, c.QueryPolicyById, "")
	c.registerMethod(QueryLockFeeMethodName, 50_000, c.QueryLockFee, "")
	c.registerMethod(QueryIsPriceChangedMethodName, 50_000, c.QueryIsPriceChanged, "")
	c.registerMethod(QueryQuotaUpdateTimeMethodName, 50_000, c.QueryQuotaUpdateTime, "")
	c.registerMethod(QueryGroupMembersExistMethodName, 50_000, c.QueryGroupMembersExist, "")
	c.registerMethod(QueryGroupsExistMethodName, 50_000, c.QueryGroupsExist, "")
	c.registerMethod(QueryGroupsExistByIdMethodName, 50_000, c.QueryGroupsExistById, "")
	c.registerMethod(QueryPaymentAccountBucketFlowRateLimitMethodName, 50_000, c.QueryPaymentAccountBucketFlowRateLimit, "")
	c.registerMethod(ParamsMethodName, 50_000, c.Params, "")
	c.registerMethod(VerifyPermissionMethodName, 50_000, c.VerifyPermission, "")
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
	bucketInfos := make([]BucketInfo, 0, len(res.BucketInfos))
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
	objectInfos := make([]ObjectInfo, 0, len(res.ObjectInfos))
	for _, objectInfo := range res.ObjectInfos {
		objectInfos = append(objectInfos, *outputObjectInfo(objectInfo))
	}
	return method.Outputs.Pack(objectInfos, outputPageResponse(res.Pagination))
}

// ListGroups queries the user's total groups.
func (c *Contract) ListGroups(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ListGroupsMethodName)
	var args ListGroupsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &storagetypes.QueryListGroupsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
		GroupOwner: args.GroupOwner.String(),
	}
	res, err := c.storageKeeper.ListGroups(ctx, msg)
	if err != nil {
		return nil, err
	}
	groupInfos := make([]GroupInfo, 0, len(res.GroupInfos))
	for _, groupInfo := range res.GroupInfos {
		groupInfos = append(groupInfos, GroupInfo{
			Owner:      common.HexToAddress(groupInfo.Owner),
			GroupName:  groupInfo.GroupName,
			SourceType: uint8(groupInfo.SourceType),
			Id:         groupInfo.Id.BigInt(),
			Extra:      groupInfo.Extra,
			Tags:       outputTags(groupInfo.Tags),
		})
	}
	return method.Outputs.Pack(groupInfos, outputPageResponse(res.Pagination))
}

// ListObjects queries a list of object items under the bucket.
func (c *Contract) ListObjectsByBucketId(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ListObjectsByBucketIdMethodName)
	var args ListObjectsByBucketIdArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &storagetypes.QueryListObjectsByBucketIdRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
		BucketId: args.BucketId,
	}
	res, err := c.storageKeeper.ListObjectsByBucketId(ctx, msg)
	if err != nil {
		return nil, err
	}
	objectInfos := make([]ObjectInfo, 0, len(res.ObjectInfos))
	for _, objectInfo := range res.ObjectInfos {
		objectInfos = append(objectInfos, *outputObjectInfo(objectInfo))
	}
	return method.Outputs.Pack(objectInfos, outputPageResponse(res.Pagination))
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

// HeadBucket queries the bucket's info by id.
func (c *Contract) HeadBucketById(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadBucketByIdMethodName)
	var args HeadBucketByIdArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadBucketByIdRequest{
		BucketId: args.BucketId,
	}
	res, err := c.storageKeeper.HeadBucketById(ctx, msg)
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

// HeadBucketExtra queries a bucket extra info (with gvg bindings and price time) with specify name.
func (c *Contract) HeadBucketExtra(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadBucketExtraMethodName)
	var args HeadBucketExtraArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadBucketExtraRequest{
		BucketName: args.BucketName,
	}
	res, err := c.storageKeeper.HeadBucketExtra(ctx, msg)
	if err != nil {
		return nil, err
	}
	localVirtualGroups := make([]LocalVirtualGroup, 0)
	for _, localVirtualGroup := range res.ExtraInfo.LocalVirtualGroups {
		localVirtualGroups = append(localVirtualGroups, LocalVirtualGroup{
			Id:                   localVirtualGroup.Id,
			GlobalVirtualGroupId: localVirtualGroup.GlobalVirtualGroupId,
			StoredSize:           localVirtualGroup.StoredSize,
			TotalChargeSize:      localVirtualGroup.TotalChargeSize,
		})
	}
	extraInfo := InternalBucketInfo{
		PriceTime:               res.ExtraInfo.PriceTime,
		TotalChargeSize:         res.ExtraInfo.TotalChargeSize,
		LocalVirtualGroups:      localVirtualGroups,
		NextLocalVirtualGroupId: res.ExtraInfo.NextLocalVirtualGroupId,
	}

	return method.Outputs.Pack(extraInfo)
}

// HeadBucket queries a bucket with EIP712 standard metadata info.
func (c *Contract) HeadBucketNFT(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadBucketNFTMethodName)
	var args HeadBucketNFTArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryNFTRequest{
		TokenId: args.TokenId,
	}
	res, err := c.storageKeeper.HeadBucketNFT(ctx, msg)
	if err != nil {
		return nil, err
	}
	attributes := make([]Trait, 0)
	for _, attribute := range res.MetaData.Attributes {
		attributes = append(attributes, Trait{
			TraitType: attribute.TraitType,
			Value:     attribute.Value,
		})
	}
	bucketMetaData := BucketMetaData{
		Description: res.MetaData.Description,
		ExternalUrl: res.MetaData.ExternalUrl,
		BucketName:  res.MetaData.BucketName,
		Image:       res.MetaData.Image,
		Attributes:  attributes,
	}

	return method.Outputs.Pack(bucketMetaData)
}

// HeadBucket queries a object with EIP712 standard metadata info.
func (c *Contract) HeadObjectNFT(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadObjectNFTMethodName)
	var args HeadObjectNFTArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryNFTRequest{
		TokenId: args.TokenId,
	}
	res, err := c.storageKeeper.HeadObjectNFT(ctx, msg)
	if err != nil {
		return nil, err
	}
	attributes := make([]Trait, 0)
	for _, attribute := range res.MetaData.Attributes {
		attributes = append(attributes, Trait{
			TraitType: attribute.TraitType,
			Value:     attribute.Value,
		})
	}
	objectMetaData := ObjectMetaData{
		Description: res.MetaData.Description,
		ExternalUrl: res.MetaData.ExternalUrl,
		ObjectName:  res.MetaData.ObjectName,
		Image:       res.MetaData.Image,
		Attributes:  attributes,
	}

	return method.Outputs.Pack(objectMetaData)
}

// HeadBucket queries a group with EIP712 standard metadata info.
func (c *Contract) HeadGroupNFT(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadGroupNFTMethodName)
	var args HeadGroupNFTArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryNFTRequest{
		TokenId: args.TokenId,
	}
	res, err := c.storageKeeper.HeadGroupNFT(ctx, msg)
	if err != nil {
		return nil, err
	}
	attributes := make([]Trait, 0)
	for _, attribute := range res.MetaData.Attributes {
		attributes = append(attributes, Trait{
			TraitType: attribute.TraitType,
			Value:     attribute.Value,
		})
	}
	groupMetaData := GroupMetaData{
		Description: res.MetaData.Description,
		ExternalUrl: res.MetaData.ExternalUrl,
		GroupName:   res.MetaData.GroupName,
		Image:       res.MetaData.Image,
		Attributes:  attributes,
	}

	return method.Outputs.Pack(groupMetaData)
}

// HeadGroup queries the group's info.
func (c *Contract) HeadGroup(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadGroupMethodName)
	var args HeadGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadGroupRequest{
		GroupOwner: args.GroupOwner.String(),
		GroupName:  args.GroupName,
	}
	res, err := c.storageKeeper.HeadGroup(ctx, msg)
	if err != nil {
		return nil, err
	}
	groupInfo := GroupInfo{
		Owner:      common.HexToAddress(res.GroupInfo.Owner),
		GroupName:  res.GroupInfo.GroupName,
		SourceType: uint8(res.GroupInfo.SourceType),
		Id:         res.GroupInfo.Id.BigInt(),
		Extra:      res.GroupInfo.Extra,
		Tags:       outputTags(res.GroupInfo.Tags),
	}
	return method.Outputs.Pack(groupInfo)
}

// HeadGroupMember queries the group member's info.
func (c *Contract) HeadGroupMember(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadGroupMemberMethodName)
	var args HeadGroupMemberArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadGroupMemberRequest{
		Member:     args.Member.String(),
		GroupOwner: args.GroupOwner.String(),
		GroupName:  args.GroupName,
	}
	res, err := c.storageKeeper.HeadGroupMember(ctx, msg)
	if err != nil {
		return nil, err
	}
	var expirationTime int64
	if res.GroupMember.ExpirationTime != nil {
		expirationTime = res.GroupMember.ExpirationTime.Unix()
	} else {
		expirationTime = 0
	}
	groupMemberInfo := GroupMember{
		Id:             res.GroupMember.Id.BigInt(),
		GroupId:        res.GroupMember.GroupId.BigInt(),
		Member:         common.HexToAddress(res.GroupMember.Member),
		ExpirationTime: expirationTime,
	}
	return method.Outputs.Pack(groupMemberInfo)
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
func (c *Contract) HeadObjectById(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) { //nolint
	method := GetAbiMethod(HeadObjectByIdMethodName)
	var args HeadObjectByIDArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadObjectByIdRequest{
		ObjectId: args.ObjectID,
	}
	res, err := c.storageKeeper.HeadObjectById(ctx, msg)
	if err != nil {
		return nil, err
	}
	objectInfo := outputObjectInfo(res.ObjectInfo)
	gvg := outputsGlobalVirtualGroup(res.GlobalVirtualGroup)
	return method.Outputs.Pack(objectInfo, gvg)
}

// HeadShadowObject queries a shadow object with specify name.
func (c *Contract) HeadShadowObject(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(HeadShadowObjectMethodName)
	var args HeadShadowObjectArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryHeadShadowObjectRequest{
		BucketName: args.BucketName,
		ObjectName: args.ObjectName,
	}
	res, err := c.storageKeeper.HeadShadowObject(ctx, msg)
	if err != nil {
		return nil, err
	}
	checksums := []string{}
	for i := range res.ObjectInfo.Checksums {
		checksums = append(checksums, hex.EncodeToString(res.ObjectInfo.Checksums[i]))
	}
	objectInfo := ShadowObjectInfo{
		Operator:    res.ObjectInfo.Operator,
		Id:          res.ObjectInfo.Id.BigInt(),
		ContentType: res.ObjectInfo.ContentType,
		PayloadSize: res.ObjectInfo.PayloadSize,
		Checksums:   checksums,
		UpdatedAt:   res.ObjectInfo.UpdatedAt,
		Version:     res.ObjectInfo.Version,
	}
	return method.Outputs.Pack(objectInfo)
}

// QueryPolicyForGroup queries the group's policy.
func (c *Contract) QueryPolicyForGroup(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryPolicyForGroupMethodName)
	var args QueryPolicyForGroupArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryPolicyForGroupRequest{
		Resource:         args.Resource,
		PrincipalGroupId: cmath.NewUintFromBigInt(args.GroupId).String(),
	}
	res, err := c.storageKeeper.QueryPolicyForGroup(ctx, msg)
	if err != nil {
		return nil, err
	}
	var expirationTime int64
	statements := make([]Statement, 0)
	if res.Policy.Statements != nil {
		for _, statement := range res.Policy.Statements {
			actions := make([]int32, 0)
			for _, action := range statement.Actions {
				actions = append(actions, int32(action))
			}
			if statement.ExpirationTime != nil {
				expirationTime = statement.ExpirationTime.Unix()
			} else {
				expirationTime = 0
			}
			statements = append(statements, Statement{
				Effect:         int32(statement.Effect),
				Actions:        actions,
				Resources:      statement.Resources,
				ExpirationTime: expirationTime,
				LimitSize:      statement.LimitSize.Value,
			})
		}
	}
	if res.Policy.ExpirationTime != nil {
		expirationTime = res.Policy.ExpirationTime.Unix()
	} else {
		expirationTime = 0
	}
	policyInfo := Policy{
		Id:             res.Policy.Id.BigInt(),
		Principal:      Principal{PrincipalType: int32(res.Policy.Principal.Type), Value: res.Policy.Principal.Value},
		ResourceType:   int32(res.Policy.ResourceType),
		ResourceId:     res.Policy.ResourceId.BigInt(),
		Statements:     statements,
		ExpirationTime: expirationTime,
	}
	return method.Outputs.Pack(policyInfo)
}

// QueryPolicyForAccount queries the account's policy.
func (c *Contract) QueryPolicyForAccount(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryPolicyForAccountMethodName)
	var args QueryPolicyForAccountArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryPolicyForAccountRequest{
		Resource:         args.Resource,
		PrincipalAddress: args.PrincipalAddr,
	}
	res, err := c.storageKeeper.QueryPolicyForAccount(ctx, msg)
	if err != nil {
		return nil, err
	}
	var expirationTime int64
	statements := make([]Statement, 0)
	if res.Policy.Statements != nil {
		for _, statement := range res.Policy.Statements {
			actions := make([]int32, 0)
			for _, action := range statement.Actions {
				actions = append(actions, int32(action))
			}
			if statement.ExpirationTime != nil {
				expirationTime = statement.ExpirationTime.Unix()
			} else {
				expirationTime = 0
			}
			statements = append(statements, Statement{
				Effect:         int32(statement.Effect),
				Actions:        actions,
				Resources:      statement.Resources,
				ExpirationTime: expirationTime,
				LimitSize:      statement.LimitSize.Value,
			})
		}
	}
	if res.Policy.ExpirationTime != nil {
		expirationTime = res.Policy.ExpirationTime.Unix()
	} else {
		expirationTime = 0
	}
	policyInfo := Policy{
		Id:             res.Policy.Id.BigInt(),
		Principal:      Principal{PrincipalType: int32(res.Policy.Principal.Type), Value: res.Policy.Principal.Value},
		ResourceType:   int32(res.Policy.ResourceType),
		ResourceId:     res.Policy.ResourceId.BigInt(),
		Statements:     statements,
		ExpirationTime: expirationTime,
	}
	return method.Outputs.Pack(policyInfo)
}

// QueryPolicyById queries a policy by policy id.
func (c *Contract) QueryPolicyById(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryPolicyByIdMethodName)
	var args QueryPolicyByIdArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryPolicyByIdRequest{
		PolicyId: args.PolicyId,
	}
	res, err := c.storageKeeper.QueryPolicyById(ctx, msg)
	if err != nil {
		return nil, err
	}
	var expirationTime int64
	statements := make([]Statement, 0)
	if res.Policy.Statements != nil {
		for _, statement := range res.Policy.Statements {
			actions := make([]int32, 0)
			for _, action := range statement.Actions {
				actions = append(actions, int32(action))
			}
			if statement.ExpirationTime != nil {
				expirationTime = statement.ExpirationTime.Unix()
			} else {
				expirationTime = 0
			}
			statements = append(statements, Statement{
				Effect:         int32(statement.Effect),
				Actions:        actions,
				Resources:      statement.Resources,
				ExpirationTime: expirationTime,
				LimitSize:      statement.LimitSize.Value,
			})
		}
	}
	if res.Policy.ExpirationTime != nil {
		expirationTime = res.Policy.ExpirationTime.Unix()
	} else {
		expirationTime = 0
	}
	policyInfo := Policy{
		Id:             res.Policy.Id.BigInt(),
		Principal:      Principal{PrincipalType: int32(res.Policy.Principal.Type), Value: res.Policy.Principal.Value},
		ResourceType:   int32(res.Policy.ResourceType),
		ResourceId:     res.Policy.ResourceId.BigInt(),
		Statements:     statements,
		ExpirationTime: expirationTime,
	}
	return method.Outputs.Pack(policyInfo)
}

// QueryLockFee queries lock fee for storing an object.
func (c *Contract) QueryLockFee(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryLockFeeMethodName)
	var args QueryLockFeeArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryLockFeeRequest{
		PrimarySpAddress: args.PrimarySpAddress,
		CreateAt:         args.CreateAt,
		PayloadSize:      args.PayloadSize,
	}
	res, err := c.storageKeeper.QueryLockFee(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.Amount.BigInt())
}

// QueryIsPriceChanged queries whether read and storage prices changed for the bucket.
func (c *Contract) QueryIsPriceChanged(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryIsPriceChangedMethodName)
	var args QueryIsPriceChangedArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryIsPriceChangedRequest{
		BucketName: args.BucketName,
	}
	res, err := c.storageKeeper.QueryIsPriceChanged(ctx, msg)
	if err != nil {
		return nil, err
	}
	isPriceChanged := IsPriceChanged{
		Changed:                    res.Changed,
		CurrentReadPrice:           res.CurrentReadPrice.BigInt(),
		CurrentPrimaryStorePrice:   res.CurrentPrimaryStorePrice.BigInt(),
		CurrentSecondaryStorePrice: res.CurrentSecondaryStorePrice.BigInt(),
		CurrentValidatorTaxRate:    res.CurrentValidatorTaxRate.BigInt(),
		NewReadPrice:               res.NewReadPrice.BigInt(),
		NewPrimaryStorePrice:       res.NewPrimaryStorePrice.BigInt(),
		NewSecondaryStorePrice:     res.NewSecondaryStorePrice.BigInt(),
		NewValidatorTaxRate:        res.NewValidatorTaxRate.BigInt(),
	}

	return method.Outputs.Pack(isPriceChanged)
}

// QueryQuotaUpdateTime queries quota update time for the bucket.
func (c *Contract) QueryQuotaUpdateTime(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryQuotaUpdateTimeMethodName)
	var args QueryQuotaUpdateTimeArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryQuoteUpdateTimeRequest{
		BucketName: args.BucketName,
	}
	res, err := c.storageKeeper.QueryQuotaUpdateTime(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.UpdateAt)
}

// QueryGroupMembersExist queries whether some members are in the group.
func (c *Contract) QueryGroupMembersExist(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryGroupMembersExistMethodName)
	var args QueryGroupMembersExistArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryGroupMembersExistRequest{
		GroupId: args.GroupId,
		Members: args.Members,
	}
	res, err := c.storageKeeper.QueryGroupMembersExist(ctx, msg)
	if err != nil {
		return nil, err
	}
	exists := make([]bool, 0)
	for _, member := range args.Members {
		exists = append(exists, res.Exists[member])
	}

	return method.Outputs.Pack(args.Members, exists)
}

// QueryGroupsExist queries whether some groups are exist.
func (c *Contract) QueryGroupsExist(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryGroupsExistMethodName)
	var args QueryGroupsExistArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryGroupsExistRequest{
		GroupOwner: args.GroupOwner,
		GroupNames: args.GroupNames,
	}
	res, err := c.storageKeeper.QueryGroupsExist(ctx, msg)
	if err != nil {
		return nil, err
	}
	exists := make([]bool, 0)
	for _, groupName := range args.GroupNames {
		exists = append(exists, res.Exists[groupName])
	}

	return method.Outputs.Pack(args.GroupNames, exists)
}

// QueryGroupsExistById queries whether some groups are exist by id.
func (c *Contract) QueryGroupsExistById(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryGroupsExistByIdMethodName)
	var args QueryGroupsExistByIdArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryGroupsExistByIdRequest{
		GroupIds: args.GroupIds,
	}
	res, err := c.storageKeeper.QueryGroupsExistById(ctx, msg)
	if err != nil {
		return nil, err
	}
	exists := make([]bool, 0)
	for _, groupId := range args.GroupIds {
		exists = append(exists, res.Exists[groupId])
	}

	return method.Outputs.Pack(args.GroupIds, exists)
}

// QueryPaymentAccountBucketFlowRateLimit queries the flow rate limit of a bucket for a payment account.
func (c *Contract) QueryPaymentAccountBucketFlowRateLimit(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryPaymentAccountBucketFlowRateLimitMethodName)
	var args QueryPaymentAccountBucketFlowRateLimitArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryPaymentAccountBucketFlowRateLimitRequest{
		PaymentAccount: args.PaymentAccount,
		BucketOwner:    args.BucketOwner,
		BucketName:     args.BucketName,
	}
	res, err := c.storageKeeper.QueryPaymentAccountBucketFlowRateLimit(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.IsSet, res.FlowRateLimit.BigInt())
}

// QueryParamsByTimestamp queries the parameters of the module by timestamp
func (c *Contract) QueryParamsByTimestamp(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(QueryParamsByTimestampMethodName)
	var args QueryParamsByTimestampArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryParamsByTimestampRequest{
		Timestamp: args.Timestamp,
	}

	res, err := c.storageKeeper.QueryParamsByTimestamp(ctx, msg)
	if err != nil {
		return nil, err
	}

	params := Params{
		VersionedParams: VersionedParams{
			MaxSegmentSize:          res.Params.VersionedParams.MaxSegmentSize,
			RedundantDataChunkNum:   res.Params.VersionedParams.RedundantDataChunkNum,
			RedundantParityChunkNum: res.Params.VersionedParams.RedundantParityChunkNum,
			MinChargeSize:           res.Params.VersionedParams.MinChargeSize,
		},
		MaxPayloadSize:                   res.Params.MaxPayloadSize,
		BscMirrorBucketRelayerFee:        res.Params.BscMirrorBucketRelayerFee,
		BscMirrorBucketAckRelayerFee:     res.Params.BscMirrorBucketAckRelayerFee,
		BscMirrorObjectRelayerFee:        res.Params.BscMirrorObjectRelayerFee,
		BscMirrorObjectAckRelayerFee:     res.Params.BscMirrorObjectAckRelayerFee,
		BscMirrorGroupRelayerFee:         res.Params.BscMirrorGroupRelayerFee,
		BscMirrorGroupAckRelayerFee:      res.Params.BscMirrorGroupAckRelayerFee,
		MaxBucketsPerAccount:             res.Params.MaxBucketsPerAccount,
		DiscontinueCountingWindow:        res.Params.DiscontinueCountingWindow,
		DiscontinueObjectMax:             res.Params.DiscontinueObjectMax,
		DiscontinueBucketMax:             res.Params.DiscontinueBucketMax,
		DiscontinueConfirmPeriod:         res.Params.DiscontinueConfirmPeriod,
		DiscontinueDeletionMax:           res.Params.DiscontinueDeletionMax,
		StalePolicyCleanupMax:            res.Params.StalePolicyCleanupMax,
		MinQuotaUpdateInterval:           res.Params.MinQuotaUpdateInterval,
		MaxLocalVirtualGroupNumPerBucket: res.Params.MaxLocalVirtualGroupNumPerBucket,
		OpMirrorBucketRelayerFee:         res.Params.OpMirrorBucketRelayerFee,
		OpMirrorBucketAckRelayerFee:      res.Params.OpMirrorBucketAckRelayerFee,
		OpMirrorObjectRelayerFee:         res.Params.OpMirrorObjectRelayerFee,
		OpMirrorObjectAckRelayerFee:      res.Params.OpMirrorObjectAckRelayerFee,
		OpMirrorGroupRelayerFee:          res.Params.OpMirrorGroupRelayerFee,
		OpMirrorGroupAckRelayerFee:       res.Params.OpMirrorGroupAckRelayerFee,
		PolygonMirrorBucketRelayerFee:    res.Params.PolygonMirrorBucketRelayerFee,
		PolygonMirrorBucketAckRelayerFee: res.Params.PolygonMirrorBucketAckRelayerFee,
	}

	return method.Outputs.Pack(params)
}

// Params queries the storage parameters
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(ParamsMethodName)

	msg := &storagetypes.QueryParamsRequest{}

	res, err := c.storageKeeper.Params(ctx, msg)
	if err != nil {
		return nil, err
	}

	params := Params{
		VersionedParams: VersionedParams{
			MaxSegmentSize:          res.Params.VersionedParams.MaxSegmentSize,
			RedundantDataChunkNum:   res.Params.VersionedParams.RedundantDataChunkNum,
			RedundantParityChunkNum: res.Params.VersionedParams.RedundantParityChunkNum,
			MinChargeSize:           res.Params.VersionedParams.MinChargeSize,
		},
		MaxPayloadSize:                   res.Params.MaxPayloadSize,
		BscMirrorBucketRelayerFee:        res.Params.BscMirrorBucketRelayerFee,
		BscMirrorBucketAckRelayerFee:     res.Params.BscMirrorBucketAckRelayerFee,
		BscMirrorObjectRelayerFee:        res.Params.BscMirrorObjectRelayerFee,
		BscMirrorObjectAckRelayerFee:     res.Params.BscMirrorObjectAckRelayerFee,
		BscMirrorGroupRelayerFee:         res.Params.BscMirrorGroupRelayerFee,
		BscMirrorGroupAckRelayerFee:      res.Params.BscMirrorGroupAckRelayerFee,
		MaxBucketsPerAccount:             res.Params.MaxBucketsPerAccount,
		DiscontinueCountingWindow:        res.Params.DiscontinueCountingWindow,
		DiscontinueObjectMax:             res.Params.DiscontinueObjectMax,
		DiscontinueBucketMax:             res.Params.DiscontinueBucketMax,
		DiscontinueConfirmPeriod:         res.Params.DiscontinueConfirmPeriod,
		DiscontinueDeletionMax:           res.Params.DiscontinueDeletionMax,
		StalePolicyCleanupMax:            res.Params.StalePolicyCleanupMax,
		MinQuotaUpdateInterval:           res.Params.MinQuotaUpdateInterval,
		MaxLocalVirtualGroupNumPerBucket: res.Params.MaxLocalVirtualGroupNumPerBucket,
		OpMirrorBucketRelayerFee:         res.Params.OpMirrorBucketRelayerFee,
		OpMirrorBucketAckRelayerFee:      res.Params.OpMirrorBucketAckRelayerFee,
		OpMirrorObjectRelayerFee:         res.Params.OpMirrorObjectRelayerFee,
		OpMirrorObjectAckRelayerFee:      res.Params.OpMirrorObjectAckRelayerFee,
		OpMirrorGroupRelayerFee:          res.Params.OpMirrorGroupRelayerFee,
		OpMirrorGroupAckRelayerFee:       res.Params.OpMirrorGroupAckRelayerFee,
		PolygonMirrorBucketRelayerFee:    res.Params.PolygonMirrorBucketRelayerFee,
		PolygonMirrorBucketAckRelayerFee: res.Params.PolygonMirrorBucketAckRelayerFee,
	}

	return method.Outputs.Pack(params)
}

// VerifyPermission queries a list of VerifyPermission items.
func (c *Contract) VerifyPermission(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(VerifyPermissionMethodName)
	var args VerifyPermissionArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &storagetypes.QueryVerifyPermissionRequest{
		Operator:   contract.Caller().String(),
		BucketName: args.BucketName,
		ObjectName: args.ObjectName,
		ActionType: permissiontypes.ActionType(args.ActionType),
	}
	res, err := c.storageKeeper.VerifyPermission(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(int32(res.Effect))
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

func outputsGlobalVirtualGroup(g *vgtypes.GlobalVirtualGroup) *GlobalVirtualGroup {
	return &GlobalVirtualGroup{
		Id:                    g.Id,
		FamilyId:              g.FamilyId,
		PrimarySpId:           g.PrimarySpId,
		SecondarySpIds:        g.SecondarySpIds,
		StoredSize:            g.StoredSize,
		VirtualPaymentAddress: common.HexToAddress(g.VirtualPaymentAddress),
		TotalDeposit:          g.TotalDeposit.String(),
	}
}

func outputTags(tags *storagetypes.ResourceTags) []Tag {
	t := make([]Tag, 0)
	if tags == nil {
		return t
	}
	for _, tag := range tags.Tags {
		t = append(t, Tag{
			Key:   tag.Key,
			Value: tag.Value,
		})
	}
	return t
}

func outputPageResponse(p *query.PageResponse) *PageResponse {
	return &PageResponse{
		NextKey: p.NextKey,
		Total:   p.Total,
	}
}
