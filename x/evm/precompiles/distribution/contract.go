package distribution

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type Contract struct {
	ctx                sdk.Context
	distributionKeeper distributionkeeper.Keeper
}

func NewPrecompiledContract(ctx sdk.Context, distributionKeeper distributionkeeper.Keeper) *Contract {
	return &Contract{
		ctx:                ctx,
		distributionKeeper: distributionKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return distributionAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case SetWithdrawAddressMethodName:
		return SetWithdrawAddressGas
	case WithdrawDelegatorRewardMethodName:
		return WithdrawDelegatorRewardGas
	case WithdrawValidatorCommissionMethodName:
		return WithdrawValidatorCommissionGas
	case FundCommunityPoolMethodName:
		return FundCommunityPoolGas
	case ValidatorDistributionInfoMethodName:
		return ValidatorDistributionInfoGas
	case ValidatorOutstandingRewardsMethodName:
		return ValidatorOutstandingRewardsGas
	case ValidatorCommissionMethodName:
		return ValidatorCommissionGas
	case DelegationRewardsMethodName:
		return DelegationRewardsGas
	case DelegationTotalRewardsMethodName:
		return DelegationTotalRewardsGas
	case CommunityPoolMethodName:
		return CommunityPoolGas
	case ParamsMethodName:
		return ParamsGas
	case ValidatorSlashesMethodName:
		return ValidatorSlashesGas
	case DelegatorValidatorsMethodName:
		return DelegatorValidatorsGas
	case delegatorWithdrawAddressMethodName:
		return delegatorWithdrawAddressGas
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
		case SetWithdrawAddressMethodName:
			ret, err = c.SetWithdrawAddress(ctx, evm, contract, readonly)
		case WithdrawDelegatorRewardMethodName:
			ret, err = c.WithdrawDelegatorReward(ctx, evm, contract, readonly)
		case WithdrawValidatorCommissionMethodName:
			ret, err = c.WithdrawValidatorCommission(ctx, evm, contract, readonly)
		case FundCommunityPoolMethodName:
			ret, err = c.FundCommunityPool(ctx, evm, contract, readonly)
		case ValidatorDistributionInfoMethodName:
			ret, err = c.ValidatorDistributionInfo(ctx, evm, contract, readonly)
		case ValidatorOutstandingRewardsMethodName:
			ret, err = c.ValidatorOutstandingRewards(ctx, evm, contract, readonly)
		case ValidatorCommissionMethodName:
			ret, err = c.ValidatorCommission(ctx, evm, contract, readonly)
		case DelegationRewardsMethodName:
			ret, err = c.DelegationRewards(ctx, evm, contract, readonly)
		case DelegationTotalRewardsMethodName:
			ret, err = c.DelegationTotalRewards(ctx, evm, contract, readonly)
		case CommunityPoolMethodName:
			ret, err = c.CommunityPool(ctx, evm, contract, readonly)
		case ParamsMethodName:
			ret, err = c.Params(ctx, evm, contract, readonly)
		case ValidatorSlashesMethodName:
			ret, err = c.ValidatorSlashes(ctx, evm, contract, readonly)
		case DelegatorValidatorsMethodName:
			ret, err = c.DelegatorValidators(ctx, evm, contract, readonly)
		case delegatorWithdrawAddressMethodName:
			ret, err = c.DelegatorWithdrawAddress(ctx, evm, contract, readonly)
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
