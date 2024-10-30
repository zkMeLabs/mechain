package virtualgroup

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	virtualgroupkeeper "github.com/evmos/evmos/v12/x/virtualgroup/keeper"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type Contract struct {
	ctx                sdk.Context
	virtualGroupKeeper virtualgroupkeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, virtualGroupKeeper virtualgroupkeeper.Keeper) *Contract {
	return &Contract{
		ctx:                ctx,
		virtualGroupKeeper: virtualGroupKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return virtualGroupAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case CreateGlobalVirtualGroupMethodName:
		return CreateGlobalVirtualGroupGas
	case DeleteGlobalVirtualGroupMethodName:
		return DeleteGlobalVirtualGroupGas
	case SwapOutMethodName:
		return SwapOutGas
	case CompleteSwapOutMethodName:
		return CompleteSwapOutGas
	case SPExitMethodName:
		return SPExitGas
	case CompleteSPExitMethodName:
		return CompleteSPExitGas
	case DepositMethodName:
		return DepositGas
	case ReserveSwapInMethodName:
		return ReserveSwapInGas
	case CompleteSwapInMethodName:
		return CompleteSwapInGas
	case CancelSwapInMethodName:
		return CancelSwapInGas
	case GlobalVirtualGroupFamiliesMethodName:
		return GlobalVirtualGroupFamiliesGas
	case GlobalVirtualGroupFamilyMethodName:
		return GlobalVirtualGroupFamilyGas
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
		switch method.Name {
		case CreateGlobalVirtualGroupMethodName:
			ret, err = c.CreateGlobalVirtualGroup(ctx, evm, contract, readonly)
		case DeleteGlobalVirtualGroupMethodName:
			ret, err = c.DeleteGlobalVirtualGroup(ctx, evm, contract, readonly)
		case SwapOutMethodName:
			ret, err = c.SwapOut(ctx, evm, contract, readonly)
		case CompleteSwapOutMethodName:
			ret, err = c.CompleteSwapOut(ctx, evm, contract, readonly)
		case SPExitMethodName:
			ret, err = c.SPExit(ctx, evm, contract, readonly)
		case CompleteSPExitMethodName:
			ret, err = c.CompleteSPExit(ctx, evm, contract, readonly)
		case DepositMethodName:
			ret, err = c.Deposit(ctx, evm, contract, readonly)
		case ReserveSwapInMethodName:
			ret, err = c.ReserveSwapIn(ctx, evm, contract, readonly)
		case CompleteSwapInMethodName:
			ret, err = c.CompleteSwapIn(ctx, evm, contract, readonly)
		case CancelSwapInMethodName:
			ret, err = c.CancelSwapIn(ctx, evm, contract, readonly)
		case GlobalVirtualGroupFamiliesMethodName:
			ret, err = c.GlobalVirtualGroupFamilies(ctx, evm, contract, readonly)
		case GlobalVirtualGroupFamilyMethodName:
			ret, err = c.GlobalVirtualGroupFamily(ctx, evm, contract, readonly)
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
