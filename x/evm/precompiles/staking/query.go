package staking

import (
	"bytes"
	"encoding/base64"
	"github.com/0xPolygon/polygon-edge/helper/hex"

	cometbfttypes "github.com/cometbft/cometbft/proto/tendermint/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	DelegationGas                    = 30_000
	UnbondingDelegationGas           = 30_000
	ValidatorGas                     = 30_000
	ValidatorsGas                    = 60_000
	ValidatorDelegationsGas          = 90_000
	ValidatorUnbondingDelegationsGas = 90_000
	DelegatorDelegationsGas          = 90_000
	DelegatorUnbondingDelegationsGas = 90_000
	RedelegationsGas                 = 90_000
	DelegatorValidatorsGas           = 60_000
	DelegatorValidatorGas            = 60_000
	HistoricalInfoGas                = 60_000
	PoolGas                          = 30_000
	ParamsGas                        = 30_000

	DelegationMethodName                    = "delegation"
	UnbondingDelegationMethodName           = "unbondingDelegation"
	ValidatorMethodName                     = "validator"
	ValidatorsMethodName                    = "validators"
	ValidatorDelegationsMethodName          = "validatorDelegations"
	ValidatorUnbondingDelegationsMethodName = "validatorUnbondingDelegations"
	DelegatorDelegationsMethodName          = "delegatorDelegations"
	DelegatorUnbondingDelegationsMethodName = "delegatorUnbondingDelegations"
	RedelegationsMethodName                 = "redelegations"
	DelegatorValidatorsMethodName           = "delegatorValidators"
	DelegatorValidatorMethodName            = "delegatorValidator"
	HistoricalInfoMethodName                = "historicalInfo"
	PoolMethodName                          = "pool"
	ParamsMethodName                        = "params"
)

func (c *Contract) Delegation(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegationMethodName)

	// parse args
	var args DelegationArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &stakingtypes.QueryDelegationRequest{
		DelegatorAddr: args.GetDelegator().String(),
		ValidatorAddr: args.GetValidator().String(),
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.Delegation(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsDelegation(*res.DelegationResponse))
}

func (c *Contract) UnbondingDelegation(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(UnbondingDelegationMethodName)

	// parse args
	var args UnbondingDelegationArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &stakingtypes.QueryUnbondingDelegationRequest{
		DelegatorAddr: args.GetDelegator().String(),
		ValidatorAddr: args.GetValidator().String(),
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.UnbondingDelegation(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsUnbondingDelegation(res.Unbond))
}

func (c *Contract) Validators(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorsMethodName)

	// parse args
	var args ValidatorsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &stakingtypes.QueryValidatorsRequest{
		Status: args.GetStatus(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.Validators(ctx, msg)
	if err != nil {
		return nil, err
	}

	var validators []Validator
	for _, validator := range res.Validators {
		validators = append(validators, OutputsValidator(validator))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(validators, pageResponse)
}

func (c *Contract) Validator(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorMethodName)

	// parse args
	var args ValidatorArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: args.GetValidator().String(),
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.Validator(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsValidator(res.Validator))
}

// ValidatorDelegations queries delegate info for given validator.
func (c *Contract) ValidatorDelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorDelegationsMethodName)

	// parse args
	var args ValidatorDelegationsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryValidatorDelegationsRequest{
		ValidatorAddr: args.GetValidator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.ValidatorDelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var delegations []DelegationResponse
	for _, delegation := range res.DelegationResponses {
		delegations = append(delegations, OutputsDelegation(delegation))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(delegations, pageResponse)
}

// ValidatorUnbondingDelegations queries unbonding delegations of a validator.
func (c *Contract) ValidatorUnbondingDelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ValidatorUnbondingDelegationsMethodName)

	// parse args
	var args ValidatorDelegationsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryValidatorUnbondingDelegationsRequest{
		ValidatorAddr: args.GetValidator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.ValidatorUnbondingDelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var unbondingDelegations []UnbondingDelegation
	for _, unbondingDelegation := range res.UnbondingResponses {
		unbondingDelegations = append(unbondingDelegations, OutputsUnbondingDelegation(unbondingDelegation))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(unbondingDelegations, pageResponse)
}

// DelegatorDelegations queries all delegations of a given delegator address.
func (c *Contract) DelegatorDelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegatorDelegationsMethodName)

	// parse args
	var args DelegatorDelegationsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryDelegatorDelegationsRequest{
		DelegatorAddr: args.GetDelegator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.DelegatorDelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var delegations []DelegationResponse
	for _, delegation := range res.DelegationResponses {
		delegations = append(delegations, OutputsDelegation(delegation))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(delegations, pageResponse)
}

// DelegatorUnbondingDelegations queries all unbonding delegations of a given
// delegator address.
func (c *Contract) DelegatorUnbondingDelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegatorUnbondingDelegationsMethodName)

	// parse args
	var args DelegatorDelegationsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryDelegatorUnbondingDelegationsRequest{
		DelegatorAddr: args.GetDelegator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.DelegatorUnbondingDelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var unbondingDelegations []UnbondingDelegation
	for _, unbondingDelegation := range res.UnbondingResponses {
		unbondingDelegations = append(unbondingDelegations, OutputsUnbondingDelegation(unbondingDelegation))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(unbondingDelegations, pageResponse)
}

// Redelegations queries redelegations of given address.
func (c *Contract) Redelegations(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(RedelegationsMethodName)

	// parse args
	var args Redelegations
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryRedelegationsRequest{
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	if args.DelegatorAddr == (common.Address{}) {
		msg.DelegatorAddr = ""
	} else {
		msg.DelegatorAddr = args.GetDelegator().String()
	}
	if args.SrcValidatorAddr == (common.Address{}) {
		msg.SrcValidatorAddr = ""
	} else {
		msg.SrcValidatorAddr = args.GetSrcValidator().String()
	}

	if args.DstValidatorAddr == (common.Address{}) {
		msg.DstValidatorAddr = ""
	} else {
		msg.DstValidatorAddr = args.GetDstValidator().String()
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.Redelegations(ctx, msg)
	if err != nil {
		return nil, err
	}

	var redelegationResponses []RedelegationResponse
	for _, redelegationResponse := range res.RedelegationResponses {
		redelegationResponses = append(redelegationResponses, OutputsRedelegation(redelegationResponse))
	}

	var pageResponse PageResponse
	if res.Pagination != nil {
		pageResponse.NextKey = res.Pagination.NextKey
		pageResponse.Total = res.Pagination.Total
	}

	return method.Outputs.Pack(redelegationResponses, pageResponse)
}

// DelegatorValidators queries all validators info for given delegator address
func (c *Contract) DelegatorValidators(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegatorValidatorsMethodName)

	// parse args
	var args DelegatorValidators
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}
	msg := &stakingtypes.QueryDelegatorValidatorsRequest{
		DelegatorAddr: args.GetDelegator().String(),
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.DelegatorValidators(ctx, msg)
	if err != nil {
		return nil, err
	}

	var validators []Validator
	for _, validator := range res.Validators {
		validators = append(validators, OutputsValidator(validator))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(validators, pageResponse)
}

// DelegatorValidator queries validator info for given delegator validator pair
func (c *Contract) DelegatorValidator(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DelegatorValidatorMethodName)

	// parse args
	var args DelegatorValidator
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &stakingtypes.QueryDelegatorValidatorRequest{
		DelegatorAddr: args.GetDelegator().String(),
		ValidatorAddr: args.GetValidator().String(),
	}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.DelegatorValidator(ctx, msg)
	if err != nil {
		return nil, err
	}

	validator := OutputsValidator(res.Validator)

	return method.Outputs.Pack(validator)
}

// HistoricalInfo queries the historical info for given height
func (c *Contract) HistoricalInfo(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(HistoricalInfoMethodName)

	// parse args
	var args HistoricalInfoRequest
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	msg := &stakingtypes.QueryHistoricalInfoRequest{Height: args.GetHeight()}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.HistoricalInfo(ctx, msg)
	if err != nil {
		return nil, err
	}

	var valsets []Validator
	for _, validator := range res.Hist.Valset {
		valsets = append(valsets, OutputsValidator(validator))
	}
	header := OutputsHeader(res.Hist.Header)
	historicalInfo := HistoricalInfo{
		Header: header,
		Valset: valsets,
	}

	return method.Outputs.Pack(historicalInfo)
}

// Pool queries the pool info
func (c *Contract) Pool(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(PoolMethodName)

	msg := &stakingtypes.QueryPoolRequest{}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.Pool(ctx, msg)
	if err != nil {
		return nil, err
	}

	pool := Pool{
		NotBondedTokens: res.Pool.NotBondedTokens.BigInt(),
		BondedTokens:    res.Pool.BondedTokens.BigInt(),
	}

	return method.Outputs.Pack(pool)
}

// Params queries the staking parameters
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ParamsMethodName)

	msg := &stakingtypes.QueryParamsRequest{}

	querier := stakingkeeper.Querier{Keeper: c.stakingKeeper}

	res, err := querier.Params(ctx, msg)
	if err != nil {
		return nil, err
	}

	params := Params{
		UnbondingTime:     int64(res.Params.UnbondingTime),
		MaxValidators:     res.Params.MaxValidators,
		MaxEntries:        res.Params.MaxEntries,
		HistoricalEntries: res.Params.HistoricalEntries,
		BondDenom:         res.Params.BondDenom,
		MinCommissionRate: res.Params.MinCommissionRate.BigInt(),
	}

	return method.Outputs.Pack(params)
}

func OutputsValidator(validator stakingtypes.Validator) Validator {
	return Validator{
		OperatorAddress: common.HexToAddress(validator.OperatorAddress),
		ConsensusPubkey: FormatConsensusPubkey(validator.ConsensusPubkey),
		Jailed:          validator.Jailed,
		Status:          uint8(validator.Status),
		Tokens:          validator.Tokens.BigInt(),
		DelegatorShares: validator.DelegatorShares.BigInt(),
		Description:     Description(validator.Description),
		UnbondingHeight: validator.UnbondingHeight,
		UnbondingTime:   validator.UnbondingTime.Unix(),
		Commission: Commission{
			CommissionRates: CommissionRates{
				Rate:          validator.Commission.Rate.BigInt(),
				MaxRate:       validator.Commission.MaxRate.BigInt(),
				MaxChangeRate: validator.Commission.MaxChangeRate.BigInt(),
			},
			UpdateTime: validator.Commission.UpdateTime.Unix(),
		},
		MinSelfDelegation:       validator.MinSelfDelegation.BigInt(),
		UnbondingOnHoldRefCount: validator.UnbondingOnHoldRefCount,
		UnbondingIds:            validator.UnbondingIds,
		SelfDelAddress:          validator.SelfDelAddress,
		RelayerAddress:          validator.RelayerAddress,
		ChallengerAddress:       validator.ChallengerAddress,
		BlsKey:                  hex.EncodeToHex(validator.BlsKey)[2:],
	}
}

func OutputsDelegation(delegationResponse stakingtypes.DelegationResponse) DelegationResponse {
	deletation := delegationResponse.Delegation
	balance := delegationResponse.Balance

	return DelegationResponse{
		Delegation: Delegation{
			DelegatorAddress: common.HexToAddress(deletation.DelegatorAddress),
			ValidatorAddress: common.HexToAddress(deletation.ValidatorAddress),
			Shares: Dec{
				Amount:    deletation.Shares.BigInt(),
				Precision: sdk.Precision,
			},
		},
		Balance: Coin{
			Denom:  balance.Denom,
			Amount: balance.Amount.BigInt(),
		},
	}
}

func OutputsUnbondingDelegation(unbondingDelegation stakingtypes.UnbondingDelegation) UnbondingDelegation {
	var entries []UnbondingDelegationEntry
	for _, entry := range unbondingDelegation.Entries {
		entries = append(entries, UnbondingDelegationEntry{
			CreationHeight: entry.CreationHeight,
			CompletionTime: entry.CompletionTime.Unix(),
			InitialBalance: entry.InitialBalance.BigInt(),
			Balance:        entry.Balance.BigInt(),
		})
	}

	return UnbondingDelegation{
		DelegatorAddress: common.HexToAddress(unbondingDelegation.DelegatorAddress),
		ValidatorAddress: common.HexToAddress(unbondingDelegation.ValidatorAddress),
		Entries:          entries,
	}
}

func OutputsRedelegation(redelegationResponse stakingtypes.RedelegationResponse) RedelegationResponse {
	var entries []RedelegationEntryResponse
	for _, entry := range redelegationResponse.Entries {
		entries = append(entries, RedelegationEntryResponse{
			RedelegationEntry: RedelegationEntry{
				CreationHeight: entry.RedelegationEntry.CreationHeight,
				CompletionTime: entry.RedelegationEntry.CompletionTime.Unix(),
				InitialBalance: entry.RedelegationEntry.InitialBalance.BigInt(),
				ShareDst:       entry.RedelegationEntry.SharesDst.BigInt(),
			},
			Balance: entry.Balance.BigInt(),
		})
	}

	var redelegationEntries []RedelegationEntry
	for _, entry := range redelegationResponse.Redelegation.Entries {
		redelegationEntries = append(redelegationEntries, RedelegationEntry{
			CreationHeight: entry.CreationHeight,
			CompletionTime: entry.CompletionTime.Unix(),
			InitialBalance: entry.InitialBalance.BigInt(),
			ShareDst:       entry.SharesDst.BigInt(),
		})
	}

	redelegation := Redelegation{
		DelegatorAddress:    common.HexToAddress(redelegationResponse.Redelegation.DelegatorAddress),
		ValidatorSrcAddress: common.HexToAddress(redelegationResponse.Redelegation.ValidatorSrcAddress),
		ValidatorDstAddress: common.HexToAddress(redelegationResponse.Redelegation.ValidatorDstAddress),
		Entries:             redelegationEntries,
	}

	return RedelegationResponse{
		Redelegation: redelegation,
		Entries:      entries,
	}
}

func OutputsHeader(header cometbfttypes.Header) Header {
	return Header{
		Version: Consensus{Block: header.Version.Block, App: header.Version.App},
		ChainId: header.ChainID,
		Height:  header.Height,
		Time:    header.Time.Unix(),
		LastBlockId: BlockID{
			Hash: hexutil.Encode(header.LastBlockId.Hash),
			PartSetHeader: PartSetHeader{
				Total: header.LastBlockId.PartSetHeader.Total,
				Hash:  hexutil.Encode(header.LastBlockId.PartSetHeader.Hash),
			},
		},
		LastCommitHash:     hexutil.Encode(header.LastCommitHash),
		DataHash:           hexutil.Encode(header.DataHash),
		ValidatorsHash:     hexutil.Encode(header.ValidatorsHash),
		NextValidatorsHash: hexutil.Encode(header.NextValidatorsHash),
		ConsensusHash:      hexutil.Encode(header.ConsensusHash),
		AppHash:            hexutil.Encode(header.AppHash),
		LastResultsHash:    hexutil.Encode(header.LastResultsHash),
		EvidenceHash:       hexutil.Encode(header.EvidenceHash),
		ProposerAddress:    hexutil.Encode(header.ProposerAddress),
	}
}

// FormatConsensusPubkey format ConsensusPubkey into a base64 string
func FormatConsensusPubkey(consensusPubkey *codectypes.Any) string {
	ed25519pk, ok := consensusPubkey.GetCachedValue().(cryptotypes.PubKey)
	if ok {
		return base64.StdEncoding.EncodeToString(ed25519pk.Bytes())
	}
	return consensusPubkey.String()
}
