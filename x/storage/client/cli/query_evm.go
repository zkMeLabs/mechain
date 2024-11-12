package cli

import (
	"context"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	mtypes "github.com/evmos/evmos/v12/types"
	"github.com/evmos/evmos/v12/x/evm/precompiles/storage"
	"github.com/evmos/evmos/v12/x/storage/types"
	"google.golang.org/grpc"
)

type QueryClientEVM struct {
	cc *ethclient.Client
}

func NewQueryClientEVM(cc *ethclient.Client) *QueryClientEVM {
	return &QueryClientEVM{cc}
}

func (c *QueryClientEVM) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryParamsByTimestamp(ctx context.Context, in *types.QueryParamsByTimestampRequest, opts ...grpc.CallOption) (*types.QueryParamsByTimestampResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadBucket(ctx context.Context, in *types.QueryHeadBucketRequest, opts ...grpc.CallOption) (*types.QueryHeadBucketResponse, error) {
	contract, err := storage.NewIStorage(common.HexToAddress(mtypes.StorageAddress), c.cc)
	if err != nil {
		return nil, err
	}

	r, err := contract.HeadBucket(&bind.CallOpts{}, in.BucketName)
	if err != nil {
		return nil, err
	}
	res := &types.QueryHeadBucketResponse{
		BucketInfo: toPbBucketInfo(&r.BucketInfo),
		ExtraInfo:  toPbBucketExtraInfo(&r.BucketExtraInfo),
	}
	return res, nil
}

func toPbBucketInfo(r *storage.BucketInfo) *types.BucketInfo {
	return &types.BucketInfo{
		Owner:                      r.Owner.String(),
		BucketName:                 r.BucketName,
		Visibility:                 types.VisibilityType(r.Visibility),
		Id:                         math.NewUintFromBigInt(r.Id),
		SourceType:                 types.SourceType(r.SourceType),
		CreateAt:                   r.CreateAt,
		PaymentAddress:             r.PaymentAddress.String(),
		GlobalVirtualGroupFamilyId: r.GlobalVirtualGroupFamilyId,
		ChargedReadQuota:           r.ChargedReadQuota,
		BucketStatus:               types.BucketStatus(r.BucketStatus),
		Tags:                       toPbTags(r.Tags),
		SpAsDelegatedAgentDisabled: r.SpAsDelegatedAgentDisabled,
	}
}

func toPbTag(r *storage.Tag) *types.ResourceTags_Tag {
	return &types.ResourceTags_Tag{
		Key:   r.Key,
		Value: r.Value,
	}
}

func toPbTags(r []storage.Tag) *types.ResourceTags {
	var tags types.ResourceTags
	for i, tag := range r {
		t := tag
		tags.Tags[i] = *toPbTag(&t)
	}
	return &tags
}

func toPbBucketExtraInfo(r *storage.BucketExtraInfo) *types.BucketExtraInfo {
	return &types.BucketExtraInfo{
		IsRateLimited:   r.IsRateLimited,
		FlowRateLimit:   math.NewIntFromBigInt(r.FlowRateLimit),
		CurrentFlowRate: math.NewIntFromBigInt(r.CurrentFlowRate),
	}
}

func (c *QueryClientEVM) HeadBucketById(ctx context.Context, in *types.QueryHeadBucketByIdRequest, opts ...grpc.CallOption) (*types.QueryHeadBucketResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadBucketNFT(ctx context.Context, in *types.QueryNFTRequest, opts ...grpc.CallOption) (*types.QueryBucketNFTResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadObject(ctx context.Context, in *types.QueryHeadObjectRequest, opts ...grpc.CallOption) (*types.QueryHeadObjectResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadObjectById(ctx context.Context, in *types.QueryHeadObjectByIdRequest, opts ...grpc.CallOption) (*types.QueryHeadObjectResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadShadowObject(ctx context.Context, in *types.QueryHeadShadowObjectRequest, opts ...grpc.CallOption) (*types.QueryHeadShadowObjectResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadObjectNFT(ctx context.Context, in *types.QueryNFTRequest, opts ...grpc.CallOption) (*types.QueryObjectNFTResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) ListBuckets(ctx context.Context, in *types.QueryListBucketsRequest, opts ...grpc.CallOption) (*types.QueryListBucketsResponse, error) {
	contract, err := storage.NewIStorage(common.HexToAddress(mtypes.StorageAddress), c.cc)
	if err != nil {
		return nil, err
	}
	if in.Pagination == nil {
		in.Pagination = &query.PageRequest{
			Limit:      100,
			CountTotal: true,
		}
	}
	r, err := contract.ListBuckets(&bind.CallOpts{}, *toStoragePageReq(in.Pagination))
	if err != nil {
		return nil, err
	}
	res := &types.QueryListBucketsResponse{}
	for _, v := range r.BucketInfos {
		p := v
		res.BucketInfos = append(res.BucketInfos, toPbBucketInfo(&p))
	}
	res.Pagination = toPbPageResp(&r.PageResponse)
	return nil, nil
}

func toStoragePageReq(in *query.PageRequest) *storage.PageRequest {
	if in == nil {
		return nil
	}
	return &storage.PageRequest{
		Key:        in.Key,
		Offset:     in.Offset,
		Limit:      in.Limit,
		CountTotal: in.CountTotal,
		Reverse:    in.Reverse,
	}
}

func toPbPageResp(p *storage.PageResponse) *query.PageResponse {
	if p == nil {
		return nil
	}
	return &query.PageResponse{
		NextKey: p.NextKey,
		Total:   p.Total,
	}
}

func (c *QueryClientEVM) ListObjects(ctx context.Context, in *types.QueryListObjectsRequest, opts ...grpc.CallOption) (*types.QueryListObjectsResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) ListObjectsByBucketId(ctx context.Context, in *types.QueryListObjectsByBucketIdRequest, opts ...grpc.CallOption) (*types.QueryListObjectsResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadGroupNFT(ctx context.Context, in *types.QueryNFTRequest, opts ...grpc.CallOption) (*types.QueryGroupNFTResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryPolicyForAccount(ctx context.Context, in *types.QueryPolicyForAccountRequest, opts ...grpc.CallOption) (*types.QueryPolicyForAccountResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) VerifyPermission(ctx context.Context, in *types.QueryVerifyPermissionRequest, opts ...grpc.CallOption) (*types.QueryVerifyPermissionResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadGroup(ctx context.Context, in *types.QueryHeadGroupRequest, opts ...grpc.CallOption) (*types.QueryHeadGroupResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) ListGroups(ctx context.Context, in *types.QueryListGroupsRequest, opts ...grpc.CallOption) (*types.QueryListGroupsResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadGroupMember(ctx context.Context, in *types.QueryHeadGroupMemberRequest, opts ...grpc.CallOption) (*types.QueryHeadGroupMemberResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryPolicyForGroup(ctx context.Context, in *types.QueryPolicyForGroupRequest, opts ...grpc.CallOption) (*types.QueryPolicyForGroupResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryPolicyById(ctx context.Context, in *types.QueryPolicyByIdRequest, opts ...grpc.CallOption) (*types.QueryPolicyByIdResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryLockFee(ctx context.Context, in *types.QueryLockFeeRequest, opts ...grpc.CallOption) (*types.QueryLockFeeResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) HeadBucketExtra(ctx context.Context, in *types.QueryHeadBucketExtraRequest, opts ...grpc.CallOption) (*types.QueryHeadBucketExtraResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryIsPriceChanged(ctx context.Context, in *types.QueryIsPriceChangedRequest, opts ...grpc.CallOption) (*types.QueryIsPriceChangedResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryQuotaUpdateTime(ctx context.Context, in *types.QueryQuoteUpdateTimeRequest, opts ...grpc.CallOption) (*types.QueryQuoteUpdateTimeResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryGroupMembersExist(ctx context.Context, in *types.QueryGroupMembersExistRequest, opts ...grpc.CallOption) (*types.QueryGroupMembersExistResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryGroupsExist(ctx context.Context, in *types.QueryGroupsExistRequest, opts ...grpc.CallOption) (*types.QueryGroupsExistResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryGroupsExistById(ctx context.Context, in *types.QueryGroupsExistByIdRequest, opts ...grpc.CallOption) (*types.QueryGroupsExistResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) QueryPaymentAccountBucketFlowRateLimit(ctx context.Context, in *types.QueryPaymentAccountBucketFlowRateLimitRequest, opts ...grpc.CallOption) (*types.QueryPaymentAccountBucketFlowRateLimitResponse, error) {
	return nil, nil
}
