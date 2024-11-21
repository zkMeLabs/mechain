package bank

import (
	"bytes"

	"github.com/evmos/evmos/v12/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	BalanceGas                 = 30_000
	AllBalancesGas             = 50_000
	TotalSupplyGas             = 50_000
	SpendableBalancesGas       = 50_000
	SpendableBalanceByDenomGas = 50_000
	SupplyOfGas                = 30_000
	ParamsGas                  = 50_000
	DenomMetadataGas           = 30_000
	DenomsMetadataGas          = 50_000
	DenomOwnersGas             = 50_000
	SendEnabledGas             = 50_000

	BalanceMethodName                 = "balance"
	AllBalancesMethodName             = "allBalances"
	TotalSupplyMethodName             = "totalSupply"
	SpendableBalancesMethodName       = "spendableBalances"
	SpendableBalanceByDenomMethodName = "spendableBalanceByDenom"
	SupplyOfMethodName                = "supplyOf"
	ParamsMethodName                  = "params"
	DenomMetadataMethodName           = "denomMetadata"
	DenomsMetadataMethodName          = "denomsMetadata"
	DenomOwnersMethodName             = "denomOwners"
	SendEnabledMethodName             = "sendEnabled"
)

func (c *Contract) Balance(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(BalanceMethodName)

	var args BalanceArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &banktypes.QueryBalanceRequest{
		Address: args.AccountAddress.String(),
		Denom:   args.Denom,
	}

	res, err := c.bankKeeper.Balance(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(Coin{
		Denom:  res.Balance.Denom,
		Amount: res.Balance.Amount.BigInt(),
	})
}

func (c *Contract) AllBalances(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(AllBalancesMethodName)

	var args AllBalancesArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &banktypes.QueryAllBalancesRequest{
		Address: args.AccountAddress.String(),
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	res, err := c.bankKeeper.AllBalances(ctx, msg)
	if err != nil {
		return nil, err
	}

	var balances []Coin
	for _, balance := range res.Balances {
		balances = append(balances, Coin{
			Denom:  balance.Denom,
			Amount: balance.Amount.BigInt(),
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(balances, pageResponse)
}

// TotalSupply queries the total supply of all coins.
func (c *Contract) TotalSupply(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(TotalSupplyMethodName)

	var args TotalSupplyArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &banktypes.QueryTotalSupplyRequest{
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	res, err := c.bankKeeper.TotalSupply(ctx, msg)
	if err != nil {
		return nil, err
	}

	var balances []Coin
	for _, balance := range res.Supply {
		balances = append(balances, Coin{
			Denom:  balance.Denom,
			Amount: balance.Amount.BigInt(),
		})
	}
	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(balances, pageResponse)
}

// SpendableBalanceByDenom queries an account's spendable balance for a specific denom.
func (c *Contract) SpendableBalanceByDenom(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(SpendableBalanceByDenomMethodName)

	var args SpendableBalanceByDenomArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: args.AccountAddress.String(),
		Denom:   args.Denom,
	}

	res, err := c.bankKeeper.SpendableBalanceByDenom(ctx, msg)
	if err != nil {
		return nil, err
	}

	balance := Coin{
		Denom:  res.Balance.Denom,
		Amount: res.Balance.Amount.BigInt(),
	}

	return method.Outputs.Pack(balance)
}

// SpendableBalances queries the spenable balance of all coins for a single account.
func (c *Contract) SpendableBalances(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(SpendableBalancesMethodName)

	var args SpendableBalancesArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &banktypes.QuerySpendableBalancesRequest{
		Address: args.AccountAddress.String(),
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	res, err := c.bankKeeper.SpendableBalances(ctx, msg)
	if err != nil {
		return nil, err
	}

	var balances []Coin
	for _, balance := range res.Balances {
		balances = append(balances, Coin{
			Denom:  balance.Denom,
			Amount: balance.Amount.BigInt(),
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(balances, pageResponse)
}

// SupplyOf queries the supply of a single coin.
func (c *Contract) SupplyOf(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(SupplyOfMethodName)

	var args SupplyOfArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &banktypes.QuerySupplyOfRequest{
		Denom: args.Denom,
	}

	res, err := c.bankKeeper.SupplyOf(ctx, msg)
	if err != nil {
		return nil, err
	}

	amount := Coin{
		Denom:  res.Amount.Denom,
		Amount: res.Amount.Amount.BigInt(),
	}

	return method.Outputs.Pack(amount)
}

// Params queries the parameters of x/bank module.
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ParamsMethodName)

	msg := &banktypes.QueryParamsRequest{}

	res, err := c.bankKeeper.Params(ctx, msg)
	if err != nil {
		return nil, err
	}

	var sendEnableds []SendEnabled
	for _, sendEnabled := range res.Params.SendEnabled {
		sendEnableds = append(sendEnableds, SendEnabled{
			Denom:   sendEnabled.Denom,
			Enabled: sendEnabled.Enabled,
		})
	}

	params := Params{
		SendEnabled:        sendEnableds,
		DefaultSendEnabled: res.Params.DefaultSendEnabled,
	}

	return method.Outputs.Pack(params)
}

// DenomMetadata queries the client metadata of a given coin denomination.
func (c *Contract) DenomMetadata(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DenomMetadataMethodName)

	var args DenomMetadataArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &banktypes.QueryDenomMetadataRequest{
		Denom: args.Denom,
	}

	res, err := c.bankKeeper.DenomMetadata(ctx, msg)
	if err != nil {
		return nil, err
	}

	var DenomUnits []DenomUnit
	for _, denomUnit := range res.Metadata.DenomUnits {
		DenomUnits = append(DenomUnits, DenomUnit{
			Denom:    denomUnit.Denom,
			Exponent: denomUnit.Exponent,
			Aliases:  denomUnit.Aliases,
		})
	}

	metaData := Metadata{
		Description: res.Metadata.Description,
		DenomUnits:  DenomUnits,
		Base:        res.Metadata.Base,
		Display:     res.Metadata.Display,
		Name:        res.Metadata.Name,
		Symbol:      res.Metadata.Symbol,
		Uri:         res.Metadata.URI,
		UriHash:     res.Metadata.URIHash,
	}

	return method.Outputs.Pack(metaData)
}

// DenomsMetadata queries the client metadata for all registered coin denominations.
func (c *Contract) DenomsMetadata(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DenomsMetadataMethodName)

	var args DenomsMetadataArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &banktypes.QueryDenomsMetadataRequest{
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	res, err := c.bankKeeper.DenomsMetadata(ctx, msg)
	if err != nil {
		return nil, err
	}

	var metaDatas []Metadata
	for _, metaData := range res.Metadatas {
		var DenomUnits []DenomUnit
		for _, denomUnit := range metaData.DenomUnits {
			DenomUnits = append(DenomUnits, DenomUnit{
				Denom:    denomUnit.Denom,
				Exponent: denomUnit.Exponent,
				Aliases:  denomUnit.Aliases,
			})
		}

		metaDatas = append(metaDatas, Metadata{
			Description: metaData.Description,
			DenomUnits:  DenomUnits,
			Base:        metaData.Base,
			Display:     metaData.Display,
			Name:        metaData.Name,
			Symbol:      metaData.Symbol,
			Uri:         metaData.URI,
			UriHash:     metaData.URIHash,
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(metaDatas, pageResponse)
}

// DenomOwners queries for all account addresses that own a particular token denomination.
func (c *Contract) DenomOwners(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DenomOwnersMethodName)

	var args DenomOwnersArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &banktypes.QueryDenomOwnersRequest{
		Denom: args.Denom,
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	res, err := c.bankKeeper.DenomOwners(ctx, msg)
	if err != nil {
		return nil, err
	}

	var denomOwners []DenomOwner
	for _, denomOwner := range res.DenomOwners {
		denomOwners = append(denomOwners, DenomOwner{
			AccountAddress: utils.AccAddressMustToHexAddress(denomOwner.Address),
			Balance: Coin{
				Denom:  denomOwner.Balance.Denom,
				Amount: denomOwner.Balance.Amount.BigInt(),
			},
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(denomOwners, pageResponse)
}

// SendEnabled queries for SendEnabled entries.
func (c *Contract) SendEnabled(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(SendEnabledMethodName)

	var args SendEnabledArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.PageRequest.Key, []byte{0}) {
		args.PageRequest.Key = nil
	}

	msg := &banktypes.QuerySendEnabledRequest{
		Denoms: args.Denoms,
		Pagination: &query.PageRequest{
			Key:        args.PageRequest.Key,
			Offset:     args.PageRequest.Offset,
			Limit:      args.PageRequest.Limit,
			CountTotal: args.PageRequest.CountTotal,
			Reverse:    args.PageRequest.Reverse,
		},
	}

	res, err := c.bankKeeper.SendEnabled(ctx, msg)
	if err != nil {
		return nil, err
	}

	var sendEnableds []SendEnabled
	for _, sendEnabled := range res.SendEnabled {
		sendEnableds = append(sendEnableds, SendEnabled{
			Denom:   sendEnabled.Denom,
			Enabled: sendEnabled.Enabled,
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(sendEnableds, pageResponse)
}
