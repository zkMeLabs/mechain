package virtualgroup

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
	virtualgrouptypes "github.com/evmos/evmos/v12/x/virtualgroup/types"
)

const (
	GlobalVirtualGroupFamiliesGas = 50_000
	GlobalVirtualGroupFamilyGas   = 20_000

	GlobalVirtualGroupFamiliesMethodName = "globalVirtualGroupFamilies"
	GlobalVirtualGroupFamilyMethodName   = "globalVirtualGroupFamily"
)

// GlobalVirtualGroupFamilies queries all the global virtual group family.
func (c *Contract) GlobalVirtualGroupFamilies(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(GlobalVirtualGroupFamiliesMethodName)

	// parse args
	var args GlobalVirtualGroupFamiliesArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &virtualgrouptypes.QueryGlobalVirtualGroupFamiliesRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.virtualGroupKeeper.GlobalVirtualGroupFamilies(ctx, msg)
	if err != nil {
		return nil, err
	}

	gvgFamilies := make([]GlobalVirtualGroupFamily, 0, len(res.GvgFamilies))
	for _, gvgFamily := range res.GvgFamilies {
		gvgFamilies = append(gvgFamilies, GlobalVirtualGroupFamily{
			Id:                    gvgFamily.Id,
			PrimarySpId:           gvgFamily.PrimarySpId,
			GlobalVirtualGroupIds: gvgFamily.GlobalVirtualGroupIds,
			VirtualPaymentAddress: common.HexToAddress(gvgFamily.VirtualPaymentAddress),
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(gvgFamilies, pageResponse)
}

// GlobalVirtualGroupFamily queries the global virtual group family by family id.
func (c *Contract) GlobalVirtualGroupFamily(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(GlobalVirtualGroupFamilyMethodName)

	// parse args
	var args GlobalVirtualGroupFamilyArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &virtualgrouptypes.QueryGlobalVirtualGroupFamilyRequest{
		FamilyId: args.FamilyId,
	}

	res, err := c.virtualGroupKeeper.GlobalVirtualGroupFamily(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(GlobalVirtualGroupFamily{
		Id:                    res.GlobalVirtualGroupFamily.Id,
		PrimarySpId:           res.GlobalVirtualGroupFamily.PrimarySpId,
		GlobalVirtualGroupIds: res.GlobalVirtualGroupFamily.GlobalVirtualGroupIds,
		VirtualPaymentAddress: common.HexToAddress(res.GlobalVirtualGroupFamily.VirtualPaymentAddress),
	})
}
