package storageprovider

import (
	"bytes"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"
)

const (
	StorageProviderMethodName  = "storageProvider"
	StorageProvidersMethodName = "storageProviders"
)

func (c *Contract) registerQuery() {
	c.registerMethod(StorageProviderMethodName, 50_000, c.StorageProvider, "")
	c.registerMethod(StorageProvidersMethodName, 80_000, c.StorageProviders, "")
}

// StorageProvider queries a storage provider with specify id.
func (c *Contract) StorageProvider(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(StorageProviderMethodName)
	// parse args
	var args StorageProviderArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &sptypes.QueryStorageProviderRequest{
		Id: args.Id,
	}
	res, err := c.spKeeper.StorageProvider(ctx, msg)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(outputStorageProviderInfo(res.StorageProvider))
}

// StorageProviders queries a list of GetStorageProviders items.
func (c *Contract) StorageProviders(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := GetAbiMethod(StorageProvidersMethodName)
	var args StorageProvidersArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &sptypes.QueryStorageProvidersRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}
	res, err := c.spKeeper.StorageProviders(ctx, msg)
	if err != nil {
		return nil, err
	}

	sps := make([]StorageProvider, 0, len(res.Sps))
	for _, objectInfo := range res.Sps {
		sps = append(sps, *outputStorageProviderInfo(objectInfo))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(sps, pageResponse)
}

func outputStorageProviderInfo(sp *sptypes.StorageProvider) *StorageProvider {
	n := &StorageProvider{
		Id:                 sp.Id,
		OperatorAddress:    sp.OperatorAddress,
		FundingAddress:     sp.FundingAddress,
		SealAddress:        sp.SealAddress,
		ApprovalAddress:    sp.ApprovalAddress,
		GcAddress:          sp.GcAddress,
		MaintenanceAddress: sp.MaintenanceAddress,
		TotalDeposit:       sp.TotalDeposit.BigInt(),
		Status:             uint8(sp.Status),
		Endpoint:           sp.Endpoint,
		Description: Description{
			Moniker:         sp.Description.Moniker,
			Identity:        sp.Description.Identity,
			Website:         sp.Description.Website,
			SecurityContact: sp.Description.SecurityContact,
			Details:         sp.Description.Details,
		},
		BlsKey: hex.EncodeToString(sp.BlsKey),
	}

	return n
}
