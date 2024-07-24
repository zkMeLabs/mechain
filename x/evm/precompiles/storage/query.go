package storage

import (
	"bytes"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
	vgtypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
)

const (
	ListBucketsMethodName     = "listBuckets"
	ListObjectsMethodName     = "listObjects"
	ListGroupsMethodName      = "listGroups"
	HeadBucketMethodName      = "headBucket"
	HeadGroupMethodName       = "headGroup"
	HeadGroupMemberMethodName = "headGroupMember"
	HeadObjectMethodName      = "headObject"
	HeadObjectByIDMethodName  = "headObjectById"
)

func (c *Contract) registerQuery() {
	c.registerMethod(ListBucketsMethodName, 50_000, c.ListBuckets, "")
	c.registerMethod(ListObjectsMethodName, 50_000, c.ListObjects, "")
	c.registerMethod(ListGroupsMethodName, 50_000, c.ListGroups, "")
	c.registerMethod(HeadBucketMethodName, 50_000, c.HeadBucket, "")
	c.registerMethod(HeadGroupMethodName, 50_000, c.HeadGroup, "")
	c.registerMethod(HeadGroupMemberMethodName, 50_000, c.HeadGroupMember, "")
	c.registerMethod(HeadObjectMethodName, 50_000, c.HeadObject, "")
	c.registerMethod(HeadObjectByIDMethodName, 50_000, c.HeadObjectById, "")
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
	expirationTime := res.GroupMember.ExpirationTime.Unix()
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
	method := GetAbiMethod(HeadObjectByIDMethodName)
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
