package slashing

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type Contract struct {
	ctx            sdk.Context
	slashingkeeper slashingkeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, slashingkeeper slashingkeeper.Keeper) *Contract {
	return &Contract{
		ctx:            ctx,
		slashingkeeper: slashingkeeper,
	}
}

func (c *Contract) Address() common.Address {
	return slashingAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case UnjailMethodName:
		return UnjailGas
	case SigningInfoMethodName:
		return SigningInfoGas
	case SigningInfosMethodName:
		return SigningInfosGas
	case ParamsMethodName:
		return paramsGas
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
		case UnjailMethodName:
			ret, err = c.Unjail(ctx, evm, contract, readonly)
		case SigningInfoMethodName:
			ret, err = c.SigningInfo(ctx, evm, contract, readonly)
		case SigningInfosMethodName:
			ret, err = c.SigningInfos(ctx, evm, contract, readonly)
		case ParamsMethodName:
			ret, err = c.Params(ctx, evm, contract, readonly)
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
