package authz

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type Contract struct {
	ctx         sdk.Context
	authzKeeper authzkeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, authzKeeper authzkeeper.Keeper) *Contract {
	return &Contract{
		ctx:         ctx,
		authzKeeper: authzKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return authzAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case GrantMethodName:
		return GrantGas
	case GranterGrantsMethodName:
		return GranterGrantsGas
	case GranteeGrantsMethodName:
		return GranteeGrantsGas
	default:
		return 0
	}
}

func (c *Contract) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) (ret []byte, err error) {
	if len(contract.Input) < 4 {
		return types.PackRetError("invalid input")
	}

	cacheCtx, commit := c.ctx.CacheContext()
	snapshot := evm.StateDB.Snapshot()

	method, err := GetMethodByID(contract.Input)
	if err == nil {
		// parse input
		switch method.Name {
		case GrantMethodName:
			ret, err = c.Grant(cacheCtx, evm, contract, readonly)
		case GrantsMethodName:
			ret, err = c.Grants(cacheCtx, evm, contract, readonly)
		case GranterGrantsMethodName:
			ret, err = c.GranterGrants(cacheCtx, evm, contract, readonly)
		case GranteeGrantsMethodName:
			ret, err = c.GranteeGrants(cacheCtx, evm, contract, readonly)
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
