package storage

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	storagekeeper "github.com/evmos/evmos/v12/x/storage/keeper"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type (
	precompiledContractFunc func(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error)
	Contract                struct {
		ctx           sdk.Context
		storageKeeper storagekeeper.Keeper
		handlers      map[string]precompiledContractFunc
		gasMeters     map[string]uint64
		events        map[string]string
	}
)

func NewPrecompiledContract(ctx sdk.Context, storageKeeper storagekeeper.Keeper) *Contract {
	c := &Contract{
		ctx:           ctx,
		storageKeeper: storageKeeper,
		handlers:      make(map[string]precompiledContractFunc),
		gasMeters:     make(map[string]uint64),
		events:        make(map[string]string),
	}
	c.registerQuery()
	c.registerTx()
	return c
}

func (c *Contract) Address() common.Address {
	return storageAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}
	return c.gasMeters[method.Name]
}

func (c *Contract) Run(evm *vm.EVM, contract *vm.Contract, readonly bool) (ret []byte, err error) {
	if len(contract.Input) < 4 {
		return types.PackRetError("invalid input")
	}
	ctx, commit := c.ctx.CacheContext()
	snapshot := evm.StateDB.Snapshot()
	defer func() {
		if err != nil {
			evm.StateDB.RevertToSnapshot(snapshot)
		}
	}()
	method, err := GetMethodByID(contract.Input)
	if err != nil {
		return types.PackRetError(err.Error())
	}
	handler, ok := c.handlers[method.Name]
	if !ok {
		return types.PackRetError("method not handled")
	}
	ret, err = handler(ctx, evm, contract, readonly)
	if err != nil {
		return nil, err
	}
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

func (c *Contract) AddOtherLog(evm *vm.EVM, event abi.Event, address common.Address, topics []common.Hash, args ...interface{}) error {
	data, newTopic, err := types.PackTopicData(event, topics, args...)
	if err != nil {
		return err
	}
	evm.StateDB.AddLog(&ethtypes.Log{
		Address:     address,
		Topics:      newTopic,
		Data:        data,
		BlockNumber: evm.Context.BlockNumber.Uint64(),
	})
	return nil
}

func (c *Contract) registerMethod(methodName string, gas uint64, handler precompiledContractFunc, eventName string) {
	method, ok := storageABI.Methods[methodName]
	if !ok {
		panic(fmt.Errorf("method %s is not exist", methodName))
	}
	c.handlers[method.Name] = handler
	c.gasMeters[method.Name] = gas
	if eventName != "" {
		if _, ok := storageABI.Events[eventName]; !ok {
			panic(fmt.Errorf("event %s is not exist", eventName))
		}
		c.events[method.Name] = eventName
	}
}
