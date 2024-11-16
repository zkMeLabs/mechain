package distribution

import (
	"bytes"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/utils"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	ValidatorDistributionInfoGas   = 30_000
	ValidatorOutstandingRewardsGas = 30_000
	ValidatorCommissionGas         = 30_000
	DelegationRewardsGas           = 30_000
	DelegationTotalRewardsGas      = 30_000
	CommunityPoolGas               = 30_000
	ParamsGas                      = 30_000
	ValidatorSlashesGas            = 30_000
	DelegatorValidatorsGas         = 30_000
	delegatorWithdrawAddressGas    = 30_000

	ValidatorDistributionInfoMethodName   = "validatorDistributionInfo"
	ValidatorOutstandingRewardsMethodName = "validatorOutstandingRewards"
	ValidatorCommissionMethodName         = "validatorCommission"
	DelegationRewardsMethodName           = "delegationRewards"
	DelegationTotalRewardsMethodName      = "delegationTotalRewards"
	CommunityPoolMethodName               = "communityPool"
	ParamsMethodName                      = "params"
	ValidatorSlashesMethodName            = "validatorSlashes"
	DelegatorValidatorsMethodName         = "delegatorValidators"
	delegatorWithdrawAddressMethodName    = "delegatorWithdrawAddress"
)

// ValidatorDistributionInfo queries validator commision and self-delegation rewards for validator
func (c *Contract) ValidatorDistributionInfo(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorDistributionInfoMethodName)

	var args ValidatorAddressArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &distributiontypes.QueryValidatorDistributionInfoRequest{
		ValidatorAddress: args.ValidatorAddress.String(),
	}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.ValidatorDistributionInfo(ctx, msg)
	if err != nil {
		return nil, err
	}

	operatorAddress := utils.AccAddressMustToHexAddress(res.OperatorAddress)

	var selfBondRewards []DecCoin
	for _, reward := range res.SelfBondRewards {
		selfBondRewards = append(selfBondRewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	var commission []DecCoin
	for _, reward := range res.Commission {
		commission = append(commission, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(operatorAddress, selfBondRewards, commission)
}

// ValidatorOutstandingRewards queries rewards of a validator address.
func (c *Contract) ValidatorOutstandingRewards(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorOutstandingRewardsMethodName)

	var args ValidatorAddressArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &distributiontypes.QueryValidatorOutstandingRewardsRequest{
		ValidatorAddress: args.ValidatorAddress.String(),
	}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.ValidatorOutstandingRewards(ctx, msg)
	if err != nil {
		return nil, err
	}

	var rewards []DecCoin
	for _, reward := range res.Rewards.Rewards {
		rewards = append(rewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(rewards)
}

// ValidatorCommission queries accumulated commission for a validator.
func (c *Contract) ValidatorCommission(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorCommissionMethodName)

	var args ValidatorAddressArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &distributiontypes.QueryValidatorCommissionRequest{
		ValidatorAddress: args.ValidatorAddress.String(),
	}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.ValidatorCommission(ctx, msg)
	if err != nil {
		return nil, err
	}

	var rewards []DecCoin
	for _, reward := range res.Commission.Commission {
		rewards = append(rewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(rewards)
}

// DelegationRewards queries the total rewards accrued by a delegation.
func (c *Contract) DelegationRewards(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegationRewardsMethodName)

	var args DelegationRewardsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &distributiontypes.QueryDelegationRewardsRequest{
		DelegatorAddress: args.DelegatorAddress.String(),
		ValidatorAddress: args.ValidatorAddress.String(),
	}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.DelegationRewards(ctx, msg)
	if err != nil {
		return nil, err
	}

	var rewards []DecCoin
	for _, reward := range res.Rewards {
		rewards = append(rewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(rewards)
}

// DelegationTotalRewards queries the total rewards accrued by a each validator.
func (c *Contract) DelegationTotalRewards(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegationTotalRewardsMethodName)

	var args DelegatorAddressArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &distributiontypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: args.DelegatorAddress.String(),
	}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.DelegationTotalRewards(ctx, msg)
	if err != nil {
		return nil, err
	}

	var delegationDelegatorReward []DelegationDelegatorReward
	for _, reward := range res.Rewards {
		var rewards []DecCoin
		for _, r := range reward.Reward {
			rewards = append(rewards, DecCoin{
				Denom:     r.Denom,
				Amount:    r.Amount.BigInt(),
				Precision: uint8(sdk.Precision),
			})
		}
		delegationDelegatorReward = append(delegationDelegatorReward, DelegationDelegatorReward{
			ValidatorAddress: common.HexToAddress(reward.ValidatorAddress),
			Rewards:          rewards,
		})
	}

	var total []DecCoin
	for _, reward := range res.Total {
		total = append(total, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(delegationDelegatorReward, total)
}

// CommunityPool queries the community pool coins.
func (c *Contract) CommunityPool(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(CommunityPoolMethodName)

	msg := &distributiontypes.QueryCommunityPoolRequest{}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.CommunityPool(ctx, msg)
	if err != nil {
		return nil, err
	}

	var rewards []DecCoin
	for _, reward := range res.Pool {
		rewards = append(rewards, DecCoin{
			Denom:     reward.Denom,
			Amount:    reward.Amount.BigInt(),
			Precision: uint8(sdk.Precision),
		})
	}

	return method.Outputs.Pack(rewards)
}

// Params queries params of distribution module
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ParamsMethodName)

	msg := &distributiontypes.QueryParamsRequest{}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.Params(ctx, msg)
	if err != nil {
		return nil, err
	}

	params := Params{
		CommunityTax:        res.Params.CommunityTax.BigInt(),
		BaseProposerReward:  res.Params.BaseProposerReward.BigInt(),
		BonusProposerReward: res.Params.BonusProposerReward.BigInt(),
		WithdrawAddrEnabled: res.Params.WithdrawAddrEnabled,
	}

	return method.Outputs.Pack(params)
}

// ValidatorSlashes queries slash events of a validator
func (c *Contract) ValidatorSlashes(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorSlashesMethodName)

	var args ValidatorSlashesArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &distributiontypes.QueryValidatorSlashesRequest{
		ValidatorAddress: args.ValidatorAddress.String(),
		StartingHeight:   args.StartingHeight,
		EndingHeight:     args.EndingHeight,
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.ValidatorSlashes(ctx, msg)
	if err != nil {
		return nil, err
	}

	var slashEvents []ValidatorSlashEvent
	for _, slash := range res.Slashes {
		slashEvents = append(slashEvents, ValidatorSlashEvent{
			ValidatorPeriod: slash.ValidatorPeriod,
			Fraction:        slash.Fraction.BigInt(),
		})
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(slashEvents, pageResponse)
}

// DelegatorValidators queries the validators list of a delegator
func (c *Contract) DelegatorValidators(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegatorValidatorsMethodName)

	var args DelegatorAddressArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &distributiontypes.QueryDelegatorValidatorsRequest{
		DelegatorAddress: args.DelegatorAddress.String(),
	}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.DelegatorValidators(ctx, msg)
	if err != nil {
		return nil, err
	}

	var validators []common.Address
	for _, validator := range res.Validators {
		validators = append(validators, common.HexToAddress(validator))
	}

	return method.Outputs.Pack(validators)
}

// DelegatorWithdrawAddress queries Query/delegatorWithdrawAddress
func (c *Contract) DelegatorWithdrawAddress(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(delegatorWithdrawAddressMethodName)

	var args DelegatorAddressArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &distributiontypes.QueryDelegatorWithdrawAddressRequest{
		DelegatorAddress: args.DelegatorAddress.String(),
	}

	querier := distributionkeeper.Querier{Keeper: c.distributionKeeper}
	res, err := querier.DelegatorWithdrawAddress(ctx, msg)
	if err != nil {
		return nil, err
	}

	withdrawAddress := common.HexToAddress(res.WithdrawAddress)

	return method.Outputs.Pack(withdrawAddress)
}
