package bank

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/evmos/evmos/v12/x/evm/types"
	paymentkeeper "github.com/evmos/evmos/v12/x/payment/keeper"
)

type Contract struct {
	ctx           sdk.Context
	bankKeeper    bankkeeper.Keeper
	paymentKeeper paymentkeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, bankKeeper bankkeeper.Keeper, paymentKeeper paymentkeeper.Keeper) *Contract {
	return &Contract{
		ctx:           ctx,
		bankKeeper:    bankKeeper,
		paymentKeeper: paymentKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return bankAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case SendMethodName:
		return SendGas
	case MultiSendMethodName:
		return MultiSendGas
	case BalanceMethodName:
		return BalanceGas
	case AllBalancesMethodName:
		return AllBalancesGas
	case TotalSupplyMethodName:
		return TotalSupplyGas
	case SpendableBalancesMethodName:
		return SpendableBalancesGas
	case SpendableBalanceByDenomMethodName:
		return SpendableBalanceByDenomGas
	case SupplyOfMethodName:
		return SupplyOfGas
	case ParamsMethodName:
		return ParamsGas
	case DenomMetadataMethodName:
		return DenomMetadataGas
	case DenomsMetadataMethodName:
		return DenomsMetadataGas
	case DenomOwnersMethodName:
		return DenomOwnersGas
	case SendEnabledMethodName:
		return SendEnabledGas
	default:
		return 0
	}
}

func (c *Contract) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) (ret []byte, err error) {
	if len(contract.Input) < 4 {
		return types.PackRetError("invalid input")
	}

	ctx, commit := c.ctx.CacheContext()
	snapshot := evm.StateDB.Snapshot()

	method, err := GetMethodByID(contract.Input)
	if err == nil {
		// parse input
		switch method.Name {
		case SendMethodName:
			ret, err = c.Send(ctx, evm, contract, readonly)
		case MultiSendMethodName:
			ret, err = c.MultiSend(ctx, evm, contract, readonly)
		case BalanceMethodName:
			ret, err = c.Balance(ctx, evm, contract, readonly)
		case AllBalancesMethodName:
			ret, err = c.AllBalances(ctx, evm, contract, readonly)
		case TotalSupplyMethodName:
			ret, err = c.TotalSupply(ctx, evm, contract, readonly)
		case SpendableBalancesMethodName:
			ret, err = c.SpendableBalances(ctx, evm, contract, readonly)
		case SpendableBalanceByDenomMethodName:
			ret, err = c.SpendableBalanceByDenom(ctx, evm, contract, readonly)
		case SupplyOfMethodName:
			ret, err = c.SupplyOf(ctx, evm, contract, readonly)
		case ParamsMethodName:
			ret, err = c.Params(ctx, evm, contract, readonly)
		case DenomMetadataMethodName:
			ret, err = c.DenomMetadata(ctx, evm, contract, readonly)
		case DenomsMetadataMethodName:
			ret, err = c.DenomsMetadata(ctx, evm, contract, readonly)
		case DenomOwnersMethodName:
			ret, err = c.DenomOwners(ctx, evm, contract, readonly)
		case SendEnabledMethodName:
			ret, err = c.SendEnabled(ctx, evm, contract, readonly)
		default:
			err = fmt.Errorf("method %s is not handle", method.Name)
		}
	}

	if err != nil {
		// revert evm state
		evm.StateDB.RevertToSnapshot(snapshot)
		return types.PackRetError(err.Error())
	}

	// commit and append events
	commit()
	return ret, nil
}

func (c *Contract) AddLog(evm *vm.EVM, event abi.Event, topics []common.Hash, args ...interface{}) error {
	data, newTopic, err := types.PackTopicData(event, topics, args...)
	if err != nil {
		return err
	}
	evm.StateDB.AddLog(&ethtypes.Log{
		Address:     c.Address(),
		Topics:      newTopic,
		Data:        data,
		BlockNumber: evm.Context.BlockNumber.Uint64(),
	})
	return nil
}
