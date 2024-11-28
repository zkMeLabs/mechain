package slashing

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	SigningInfoGas  = 30_000
	SigningInfosGas = 50_000
	paramsGas       = 30_000

	SigningInfoMethodName  = "signingInfo"
	SigningInfosMethodName = "signingInfos"
	ParamsMethodName       = "params"
)

func (c *Contract) SigningInfo(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(SigningInfoMethodName)

	var args SigningInfoArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &slashingtypes.QuerySigningInfoRequest{
		ConsAddress: args.GetConsAddress().String(),
	}

	res, err := c.slashingkeeper.SigningInfo(ctx, msg)
	if err != nil {
		return nil, err
	}

	valSigningInfo := OutPutValidatorSigningInfo(res.ValSigningInfo)

	return method.Outputs.Pack(valSigningInfo)
}

func (c *Contract) SigningInfos(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(SigningInfosMethodName)

	var args SigningInfosArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &slashingtypes.QuerySigningInfosRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.slashingkeeper.SigningInfos(ctx, msg)
	if err != nil {
		return nil, err
	}

	var valSigningInfos []ValidatorSigningInfo
	for _, valSigningInfo := range res.Info {
		valSigningInfos = append(valSigningInfos, OutPutValidatorSigningInfo(valSigningInfo))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(valSigningInfos, pageResponse)
}

func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ParamsMethodName)

	msg := &slashingtypes.QueryParamsRequest{}

	res, err := c.slashingkeeper.Params(ctx, msg)
	if err != nil {
		return nil, err
	}

	params := Params{
		SignedBlocksWindow:      res.Params.SignedBlocksWindow,
		MinSignedPerWindow:      res.Params.MinSignedPerWindow.String(),
		DowntimeJailDuration:    int64(res.Params.DowntimeJailDuration.Seconds()),
		SlashFractionDoubleSign: res.Params.SlashFractionDoubleSign.String(),
		SlashFractionDowntime:   res.Params.SlashFractionDowntime.String(),
	}

	return method.Outputs.Pack(params)
}

func OutPutValidatorSigningInfo(validatorSigningInfo slashingtypes.ValidatorSigningInfo) ValidatorSigningInfo {
	return ValidatorSigningInfo{
		ConsAddress:         common.HexToAddress(validatorSigningInfo.Address),
		StartHeight:         validatorSigningInfo.StartHeight,
		IndexOffset:         validatorSigningInfo.IndexOffset,
		JailedUntil:         validatorSigningInfo.JailedUntil.Unix(),
		Tombstoned:          validatorSigningInfo.Tombstoned,
		MissedBlocksCounter: validatorSigningInfo.MissedBlocksCounter,
	}
}
