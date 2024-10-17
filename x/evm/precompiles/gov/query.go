package gov

import (
	"bytes"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	proposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/evmos/evmos/v12/utils"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
	feemarkettypes "github.com/evmos/evmos/v12/x/feemarket/types"

	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	ProposalGas     = 30_000
	ProposalsGas    = 30_000
	VoteQueryGas    = 30_000
	VotesGas        = 30_000
	DepositQueryGas = 30_000
	DepositsGas     = 30_000
	TallyResultGas  = 30_000
	ParamsGas       = 30_000

	ProposalMethodName     = "proposal"
	ProposalsMethodName    = "proposals"
	VoteQueryMethodName    = "vote0"
	VotesMethodName        = "votes"
	DepositQueryMethodName = "deposit"
	DepositsMethodName     = "deposits"
	TallyResultMethodName  = "tallyResult"
	ParamsMethodName       = "params"
)

// Proposal returns proposal details based on ProposalID
func (c *Contract) Proposal(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ProposalMethodName)

	var args ProposalArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &govv1.QueryProposalRequest{
		ProposalId: args.ProposalId,
	}

	res, err := c.govKeeper.Proposal(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsProposal(*res.Proposal))
}

// Proposals queries all proposals based on given status.
func (c *Contract) Proposals(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ProposalsMethodName)

	var args ProposalsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	voter := ""
	if args.Voter != (common.Address{}) {
		voter = sdk.AccAddress(args.Voter.Bytes()).String()
	}

	depositor := ""
	if args.Depositor != (common.Address{}) {
		depositor = sdk.AccAddress(args.Depositor.Bytes()).String()
	}

	msg := &govv1.QueryProposalsRequest{
		ProposalStatus: govv1.ProposalStatus(args.Status),
		Voter:          voter,
		Depositor:      depositor,
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.govKeeper.Proposals(ctx, msg)
	if err != nil {
		return nil, err
	}

	var proposals []Proposal
	for _, proposal := range res.Proposals {
		proposals = append(proposals, OutputsProposal(*proposal))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(proposals, pageResponse)
}

// VoteQuery returns Voted information based on proposalID, voterAddr
func (c *Contract) VoteQuery(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(VoteQueryMethodName)

	var args VoteQueryArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &govv1.QueryVoteRequest{
		ProposalId: args.ProposalId,
		Voter:      sdk.AccAddress(args.Voter.Bytes()).String(),
	}

	res, err := c.govKeeper.Vote(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsVote(*res.Vote))
}

// Votes returns single proposal's votes
func (c *Contract) Votes(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(VotesMethodName)

	var args VotesArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &govv1.QueryVotesRequest{
		ProposalId: args.ProposalId,
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.govKeeper.Votes(ctx, msg)
	if err != nil {
		return nil, err
	}

	var votes []VoteData
	for _, vote := range res.Votes {
		votes = append(votes, OutputsVote(*vote))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(votes, pageResponse)
}

// DepositQuery queries single deposit information based on proposalID, depositAddr.
func (c *Contract) DepositQuery(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DepositQueryMethodName)

	var args DepositQueryArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &govv1.QueryDepositRequest{
		ProposalId: args.ProposalId,
		Depositor:  sdk.AccAddress(args.Depositor.Bytes()).String(),
	}

	res, err := c.govKeeper.Deposit(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(OutputsDeposit(*res.Deposit))
}

// Deposits returns single proposal's all deposits
func (c *Contract) Deposits(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(DepositsMethodName)

	var args DepositsArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}

	if bytes.Equal(args.Pagination.Key, []byte{0}) {
		args.Pagination.Key = nil
	}

	msg := &govv1.QueryDepositsRequest{
		ProposalId: args.ProposalId,
		Pagination: &query.PageRequest{
			Key:        args.Pagination.Key,
			Offset:     args.Pagination.Offset,
			Limit:      args.Pagination.Limit,
			CountTotal: args.Pagination.CountTotal,
			Reverse:    args.Pagination.Reverse,
		},
	}

	res, err := c.govKeeper.Deposits(ctx, msg)
	if err != nil {
		return nil, err
	}

	var deposits []DepositData
	for _, vote := range res.Deposits {
		deposits = append(deposits, OutputsDeposit(*vote))
	}

	var pageResponse PageResponse
	pageResponse.NextKey = res.Pagination.NextKey
	pageResponse.Total = res.Pagination.Total

	return method.Outputs.Pack(deposits, pageResponse)
}

// TallyResult queries the tally of a proposal vote.
func (c *Contract) TallyResult(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(TallyResultMethodName)

	var args ProposalArgs
	if err := types.ParseMethodArgs(method, &args, contract.Input[4:]); err != nil {
		return nil, err
	}
	msg := &govv1.QueryTallyResultRequest{
		ProposalId: args.ProposalId,
	}

	res, err := c.govKeeper.TallyResult(ctx, msg)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(TallyResult(*res.Tally))
}

// Params queries the staking parameters
func (c *Contract) Params(ctx sdk.Context, _ *vm.EVM, contract *vm.Contract, _ bool) ([]byte, error) {
	method := MustMethod(ParamsMethodName)

	msg1 := &govv1.QueryParamsRequest{
		ParamsType: govv1.ParamDeposit,
	}

	res1, err := c.govKeeper.Params(ctx, msg1)
	if err != nil {
		return nil, err
	}

	msg2 := &govv1.QueryParamsRequest{
		ParamsType: govv1.ParamVoting,
	}

	res2, err := c.govKeeper.Params(ctx, msg2)
	if err != nil {
		return nil, err
	}

	msg3 := &govv1.QueryParamsRequest{
		ParamsType: govv1.ParamTallying,
	}

	res3, err := c.govKeeper.Params(ctx, msg3)
	if err != nil {
		return nil, err
	}

	params := Params{
		MinDeposit: []Coin{
			{
				Denom:  res1.DepositParams.MinDeposit[0].Denom,
				Amount: res1.DepositParams.MinDeposit[0].Amount.BigInt(),
			},
		},
		MaxDepositPeriod: int64(res1.DepositParams.MaxDepositPeriod.Seconds()),
		VotingPeriod:     int64(res2.VotingParams.VotingPeriod.Seconds()),
		Quorum:           res3.TallyParams.Quorum,
		Threshold:        res3.TallyParams.Threshold,
		VetoThreshold:    res3.TallyParams.VetoThreshold,
	}

	return method.Outputs.Pack(params)
}

func OutputsProposal(proposal govv1.Proposal) Proposal {
	var messages []string
	msgs, err := proposal.GetMsgs()

	emptyProposal := Proposal{
		Id:               0,
		Messages:         nil,
		Status:           0,
		FinalTallyResult: TallyResult{},
		SubmitTime:       0,
		DepositEndTime:   0,
		TotalDeposit:     nil,
		VotingStartTime:  0,
		VotingEndTime:    0,
		Metadata:         "",
		Title:            "",
		Summary:          "",
	}

	if err != nil {
		return emptyProposal
	}

	interfaceRegistry := codectypes.NewInterfaceRegistry()

	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	distrtypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	govv1beta1.RegisterInterfaces(interfaceRegistry)
	govv1.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)
	feemarkettypes.RegisterInterfaces(interfaceRegistry)
	erc20types.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	proposaltypes.RegisterInterfaces(interfaceRegistry)

	mechainCodec := codec.NewProtoCodec(interfaceRegistry)

	for _, msg := range msgs {
		bytesMsg, err := mechainCodec.MarshalInterfaceJSON(msg)
		if err != nil {
			messages = append(messages, msg.String())
		}

		messages = append(messages, string(bytesMsg))
	}

	var totalDeposit []Coin
	for _, coin := range proposal.TotalDeposit {
		totalDeposit = append(totalDeposit, Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	var votingStartTime int64 = 0
	if proposal.VotingStartTime != nil {
		votingStartTime = proposal.VotingStartTime.Unix()
	}

	var votingEndTime int64 = 0
	if proposal.VotingEndTime != nil {
		votingEndTime = proposal.VotingEndTime.Unix()
	}

	return Proposal{
		Id:               proposal.Id,
		Messages:         messages,
		Status:           uint8(proposal.Status),
		FinalTallyResult: TallyResult(*proposal.FinalTallyResult),
		SubmitTime:       proposal.SubmitTime.Unix(),
		DepositEndTime:   proposal.DepositEndTime.Unix(),
		TotalDeposit:     totalDeposit,
		VotingStartTime:  votingStartTime,
		VotingEndTime:    votingEndTime,
		Metadata:         proposal.Metadata,
	}
}

func OutputsVote(vote govv1.Vote) VoteData {
	var options []WeightedVoteOption
	for _, option := range vote.Options {
		options = append(options, WeightedVoteOption{
			Option: uint8(option.Option),
			Weight: option.Weight,
		})
	}

	return VoteData{
		ProposalId: vote.ProposalId,
		Voter:      utils.AccAddressMustToHexAddress(vote.Voter),
		Options:    options,
		Metadata:   vote.Metadata,
	}
}

func OutputsDeposit(deposit govv1.Deposit) DepositData {
	var amount []Coin
	for _, coin := range deposit.Amount {
		amount = append(amount, Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.BigInt(),
		})
	}

	return DepositData{
		ProposalId: deposit.ProposalId,
		Depositor:  utils.AccAddressMustToHexAddress(deposit.Depositor),
		Amount:     amount,
	}
}
