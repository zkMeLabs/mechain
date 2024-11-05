package staking

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type Contract struct {
	ctx           sdk.Context
	stakingKeeper *stakingkeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, stakingKeeper *stakingkeeper.Keeper) *Contract {
	return &Contract{
		ctx:           ctx,
		stakingKeeper: stakingKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return stakingAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case EditValidatorMethodName:
		return EditValidatorGas
	case DelegateMethodName:
		return DelegateGas
	case UndelegateMethodName:
		return UndelegateGas
	case RedelegateMethodName:
		return RedelegateGas
	case CancelUnbondingDelegationMethodName:
		return CancelUnbondingDelegationGas
	case DelegationMethodName:
		return DelegationGas
	case UnbondingDelegationMethodName:
		return UnbondingDelegationGas
	case ValidatorMethodName:
		return ValidatorGas
	case ValidatorsMethodName:
		return ValidatorsGas
	case ValidatorDelegationsMethodName:
		return ValidatorDelegationsGas
	case ValidatorUnbondingDelegationsMethodName:
		return ValidatorUnbondingDelegationsGas
	case DelegatorDelegationsMethodName:
		return DelegatorDelegationsGas
	case DelegatorUnbondingDelegationsMethodName:
		return DelegatorUnbondingDelegationsGas
	case RedelegationsMethodName:
		return RedelegationsGas
	case DelegatorValidatorsMethodName:
		return DelegatorValidatorsGas
	case DelegatorValidatorMethodName:
		return DelegatorValidatorGas
	case HistoricalInfoMethodName:
		return HistoricalInfoGas
	case PoolMethodName:
		return PoolGas
	case ParamsMethodName:
		return ParamsGas
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
		case EditValidatorMethodName:
			ret, err = c.EditValidator(ctx, evm, contract, readonly)
		case DelegateMethodName:
			ret, err = c.Delegate(ctx, evm, contract, readonly)
		case UndelegateMethodName:
			ret, err = c.Undelegate(ctx, evm, contract, readonly)
		case RedelegateMethodName:
			ret, err = c.Redelegatge(ctx, evm, contract, readonly)
		case CancelUnbondingDelegationMethodName:
			ret, err = c.CancelUnbondingDelegation(ctx, evm, contract, readonly)
		case DelegationMethodName:
			ret, err = c.Delegation(ctx, evm, contract, readonly)
		case UnbondingDelegationMethodName:
			ret, err = c.UnbondingDelegation(ctx, evm, contract, readonly)
		case ValidatorMethodName:
			ret, err = c.Validator(ctx, evm, contract, readonly)
		case ValidatorsMethodName:
			ret, err = c.Validators(ctx, evm, contract, readonly)
		case ValidatorDelegationsMethodName:
			ret, err = c.ValidatorDelegations(ctx, evm, contract, readonly)
		case ValidatorUnbondingDelegationsMethodName:
			ret, err = c.ValidatorUnbondingDelegations(ctx, evm, contract, readonly)
		case DelegatorDelegationsMethodName:
			ret, err = c.DelegatorDelegations(ctx, evm, contract, readonly)
		case DelegatorUnbondingDelegationsMethodName:
			ret, err = c.DelegatorUnbondingDelegations(ctx, evm, contract, readonly)
		case RedelegationsMethodName:
			ret, err = c.Redelegations(ctx, evm, contract, readonly)
		case DelegatorValidatorsMethodName:
			ret, err = c.DelegatorValidators(ctx, evm, contract, readonly)
		case DelegatorValidatorMethodName:
			ret, err = c.DelegatorValidator(ctx, evm, contract, readonly)
		case HistoricalInfoMethodName:
			ret, err = c.HistoricalInfo(ctx, evm, contract, readonly)
		case PoolMethodName:
			ret, err = c.Pool(ctx, evm, contract, readonly)
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
