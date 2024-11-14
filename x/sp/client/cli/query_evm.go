package cli

import (
	"context"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	mtypes "github.com/evmos/evmos/v12/types"
	spp "github.com/evmos/evmos/v12/x/evm/precompiles/storageprovider"
	"github.com/evmos/evmos/v12/x/sp/types"
	"google.golang.org/grpc"
)

type QueryClientEVM struct {
	cc *ethclient.Client
}

func NewQueryClientEVM(cc *ethclient.Client) types.QueryClient {
	return &QueryClientEVM{cc}
}

func (c *QueryClientEVM) Params(ctx context.Context, in *types.QueryParamsRequest, opts ...grpc.CallOption) (*types.QueryParamsResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) StorageProviders(ctx context.Context, in *types.QueryStorageProvidersRequest, opts ...grpc.CallOption) (*types.QueryStorageProvidersResponse, error) {
	contract, err := spp.NewIStorageProvider(common.HexToAddress(mtypes.SpAddress), c.cc)
	if err != nil {
		return nil, err
	}
	if in.Pagination == nil {
		in.Pagination = &query.PageRequest{
			Limit:      100,
			CountTotal: true,
		}
	}
	r, err := contract.StorageProviders(&bind.CallOpts{}, *toSppPageReq(in.Pagination))
	if err != nil {
		return nil, err
	}
	res := &types.QueryStorageProvidersResponse{}
	for _, v := range r.StorageProviders {
		p := v
		res.Sps = append(res.Sps, toPbSP(&p))
	}
	res.Pagination = toPbPageResp(&r.PageResponse)
	return res, nil
}

func (c *QueryClientEVM) QuerySpStoragePrice(ctx context.Context, in *types.QuerySpStoragePriceRequest, opts ...grpc.CallOption) (*types.QuerySpStoragePriceResponse, error) {
	contract, err := spp.NewIStorageProvider(common.HexToAddress(mtypes.SpAddress), c.cc)
	if err != nil {
		return nil, err
	}
	r, err := contract.StorageProviderPrice(&bind.CallOpts{}, common.HexToAddress(in.SpAddr))
	if err != nil {
		return nil, err
	}
	res := &types.QuerySpStoragePriceResponse{
		SpStoragePrice: *toPbPrice(&r),
	}
	return res, nil
}

func (c *QueryClientEVM) QueryGlobalSpStorePriceByTime(ctx context.Context, in *types.QueryGlobalSpStorePriceByTimeRequest, opts ...grpc.CallOption) (*types.QueryGlobalSpStorePriceByTimeResponse, error) {
	return nil, nil
}

func (c *QueryClientEVM) StorageProvider(ctx context.Context, in *types.QueryStorageProviderRequest, opts ...grpc.CallOption) (*types.QueryStorageProviderResponse, error) {
	contract, err := spp.NewIStorageProvider(common.HexToAddress(mtypes.SpAddress), c.cc)
	if err != nil {
		return nil, err
	}
	r, err := contract.StorageProvider(&bind.CallOpts{}, in.Id)
	if err != nil {
		return nil, err
	}
	res := &types.QueryStorageProviderResponse{
		StorageProvider: toPbSP(&r),
	}
	return res, nil
}

func (c *QueryClientEVM) StorageProviderByOperatorAddress(ctx context.Context, in *types.QueryStorageProviderByOperatorAddressRequest, opts ...grpc.CallOption) (*types.QueryStorageProviderByOperatorAddressResponse, error) {
	contract, err := spp.NewIStorageProvider(common.HexToAddress(mtypes.SpAddress), c.cc)
	if err != nil {
		return nil, err
	}
	r, err := contract.StorageProviderByOperatorAddress(&bind.CallOpts{}, common.HexToAddress(in.OperatorAddress))
	if err != nil {
		return nil, err
	}
	res := &types.QueryStorageProviderByOperatorAddressResponse{
		StorageProvider: toPbSP(&r),
	}
	return res, nil
}

func (c *QueryClientEVM) StorageProviderMaintenanceRecordsByOperatorAddress(ctx context.Context, in *types.QueryStorageProviderMaintenanceRecordsRequest, opts ...grpc.CallOption) (*types.QueryStorageProviderMaintenanceRecordsResponse, error) {
	return nil, nil
}

func toPbSP(p *spp.StorageProvider) *types.StorageProvider {
	if p == nil {
		return nil
	}
	s := &types.StorageProvider{
		Id:                 p.Id,
		OperatorAddress:    p.OperatorAddress,
		FundingAddress:     p.FundingAddress,
		SealAddress:        p.SealAddress,
		ApprovalAddress:    p.ApprovalAddress,
		GcAddress:          p.GcAddress,
		MaintenanceAddress: p.MaintenanceAddress,
		TotalDeposit:       math.NewIntFromBigInt(p.TotalDeposit),
		Status:             types.Status(p.Status),
		Endpoint:           p.Endpoint,
		Description:        *toPbDescription(&p.Description),
		BlsKey:             []byte(p.BlsKey),
	}
	return s
}

func toPbDescription(d *spp.Description) *types.Description {
	if d == nil {
		return nil
	}
	return &types.Description{
		Moniker:         d.Moniker,
		Identity:        d.Identity,
		Website:         d.Website,
		SecurityContact: d.SecurityContact,
		Details:         d.Details,
	}
}

func toSppPageReq(in *query.PageRequest) *spp.PageRequest {
	if in == nil {
		return nil
	}
	return &spp.PageRequest{
		Key:        in.Key,
		Offset:     in.Offset,
		Limit:      in.Limit,
		CountTotal: in.CountTotal,
		Reverse:    in.Reverse,
	}
}

func toPbPageResp(p *spp.PageResponse) *query.PageResponse {
	if p == nil {
		return nil
	}
	return &query.PageResponse{
		NextKey: p.NextKey,
		Total:   p.Total,
	}
}

func toPbPrice(p *spp.SpStoragePrice) *types.SpStoragePrice {
	if p == nil {
		return nil
	}
	return &types.SpStoragePrice{
		SpId:          p.SpId,
		UpdateTimeSec: p.UpdateTimeSec.Int64(),
		ReadPrice:     math.LegacyNewDecFromInt(math.NewIntFromBigInt(p.ReadPrice)),
		FreeReadQuota: p.FreeReadQuota,
		StorePrice:    math.LegacyNewDecFromInt(math.NewIntFromBigInt(p.StorePrice)),
	}
}
