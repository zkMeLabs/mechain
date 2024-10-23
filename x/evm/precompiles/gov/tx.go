package gov

import (
	"encoding/json"
	"fmt"

	"github.com/evmos/evmos/v12/utils"
	bridgetypes "github.com/evmos/evmos/v12/x/bridge/types"
	challengetypes "github.com/evmos/evmos/v12/x/challenge/types"
	erc20types "github.com/evmos/evmos/v12/x/erc20/types"
	"github.com/evmos/evmos/v12/x/evm/types"
	feemarkettypes "github.com/evmos/evmos/v12/x/feemarket/types"
	gensptypes "github.com/evmos/evmos/v12/x/gensp/types"
	paymenttypes "github.com/evmos/evmos/v12/x/payment/types"
	permissiontypes "github.com/evmos/evmos/v12/x/permission/types"
	sptypes "github.com/evmos/evmos/v12/x/sp/types"
	storagetypes "github.com/evmos/evmos/v12/x/storage/types"
	virtualgrouptypes "github.com/evmos/evmos/v12/x/virtualgroup/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	crosschaintypes "github.com/cosmos/cosmos-sdk/x/crosschain/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gashubtypes "github.com/cosmos/cosmos-sdk/x/gashub/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	oracletypes "github.com/cosmos/cosmos-sdk/x/oracle/types"
	proposaltypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	LegacySubmitProposalGas = 60_000
	SubmitProposalGas       = 60_000
	VoteGas                 = 60_000
	VoteWeightedGas         = 60_000
	DepositGas              = 60_000

	LegacySubmitProposalMethodName = "legacySubmitProposal"
	SubmitProposalMethodName       = "submitProposal"
	VoteMethodName                 = "vote"
	VoteWeightedMethodName         = "voteWeighted"
	DepositMethodName              = "deposit0"

	LegacySubmitProposalEventName = "LegacySubmitProposal"
	SubmitProposalEventName       = "SubmitProposal"
	VoteEventName                 = "Vote"
	VoteWeightedEventName         = "VoteWeighted"
	DepositEventName              = "Deposit"
)

func (c *Contract) LegacySubmitProposal(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(LegacySubmitProposalMethodName)

	var args LegacySubmitProposalArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var amount sdk.Coins
	for _, deposit := range args.InitialDeposit {
		if deposit.Amount.Sign() > 0 {
			amount = amount.Add(sdk.Coin{
				Denom:  deposit.Denom,
				Amount: sdk.NewIntFromBigInt(deposit.Amount),
			})
		}
	}

	content, _ := govv1beta1.ContentFromProposalType(args.Title, args.Description, govv1beta1.ProposalTypeText)
	msg, err := govv1beta1.NewMsgSubmitProposal(content, amount, contract.Caller().Bytes())
	if err != nil {
		return nil, fmt.Errorf("invalid message: %w", err)
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	msgServer := govkeeper.NewMsgServerImpl(&c.govKeeper)
	server := govkeeper.NewLegacyMsgServerImpl(c.accountKeeper.GetModuleAddress(govtypes.ModuleName).String(), msgServer)
	res, err := server.SubmitProposal(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(LegacySubmitProposalEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.ProposalId,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.ProposalId)
}

func (c *Contract) SubmitProposal(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(SubmitProposalMethodName)

	var args SubmitProposalArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var messages []json.RawMessage
	err = json.Unmarshal([]byte(args.Messages), &messages)
	if err != nil {
		return nil, err
	}

	interfaceRegistry := codectypes.NewInterfaceRegistry()

	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	consensustypes.RegisterInterfaces(interfaceRegistry)
	crosschaintypes.RegisterInterfaces(interfaceRegistry)
	gashubtypes.RegisterInterfaces(interfaceRegistry)
	oracletypes.RegisterInterfaces(interfaceRegistry)
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	distrtypes.RegisterInterfaces(interfaceRegistry)
	slashingtypes.RegisterInterfaces(interfaceRegistry)
	govv1beta1.RegisterInterfaces(interfaceRegistry)
	govv1.RegisterInterfaces(interfaceRegistry)
	crisistypes.RegisterInterfaces(interfaceRegistry)
	ibctransfertypes.RegisterInterfaces(interfaceRegistry)
	upgradetypes.RegisterInterfaces(interfaceRegistry)
	proposaltypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)

	bridgetypes.RegisterInterfaces(interfaceRegistry)
	challengetypes.RegisterInterfaces(interfaceRegistry)
	erc20types.RegisterInterfaces(interfaceRegistry)
	types.RegisterInterfaces(interfaceRegistry)
	feemarkettypes.RegisterInterfaces(interfaceRegistry)
	gensptypes.RegisterInterfaces(interfaceRegistry)
	paymenttypes.RegisterInterfaces(interfaceRegistry)
	permissiontypes.RegisterInterfaces(interfaceRegistry)
	sptypes.RegisterInterfaces(interfaceRegistry)
	storagetypes.RegisterInterfaces(interfaceRegistry)
	virtualgrouptypes.RegisterInterfaces(interfaceRegistry)

	protoCodec := codec.NewProtoCodec(interfaceRegistry)

	msgs := make([]sdk.Msg, len(messages))
	for i, message := range messages {
		var msg sdk.Msg
		err := protoCodec.UnmarshalInterfaceJSON(message, &msg)
		if err != nil {
			return nil, err
		}

		msgs[i] = msg
	}

	var amount sdk.Coins
	for _, deposit := range args.InitialDeposit {
		if deposit.Amount.Sign() > 0 {
			amount = amount.Add(sdk.Coin{
				Denom:  deposit.Denom,
				Amount: sdk.NewIntFromBigInt(deposit.Amount),
			})
		}
	}

	msg, err := govv1.NewMsgSubmitProposal(msgs, amount, sdk.AccAddress(contract.Caller().Bytes()).String(), args.Metadata, args.Title, args.Summary)
	if err != nil {
		return nil, fmt.Errorf("invalid message: %w", err)
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := govkeeper.NewMsgServerImpl(&c.govKeeper)
	res, err := server.SubmitProposal(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(SubmitProposalEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		res.ProposalId,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(res.ProposalId)
}

func (c *Contract) Vote(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(VoteMethodName)

	var args VoteArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	msg := &govv1.MsgVote{
		ProposalId: args.ProposalId,
		Voter:      sdk.AccAddress(contract.Caller().Bytes()).String(),
		Option:     govv1.VoteOption(args.Option),
		Metadata:   args.Metadata,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := govkeeper.NewMsgServerImpl(&c.govKeeper)
	_, err = server.Vote(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(VoteEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		args.ProposalId,
		args.Option,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) VoteWeighted(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(VoteWeightedMethodName)

	var args VoteWeightedArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var options []*govv1.WeightedVoteOption
	for _, option := range args.Options {
		options = append(options, &govv1.WeightedVoteOption{
			Option: govv1.VoteOption(option.Option),
			Weight: option.Weight,
		})
	}

	msg := &govv1.MsgVoteWeighted{
		ProposalId: args.ProposalId,
		Voter:      sdk.AccAddress(contract.Caller().Bytes()).String(),
		Options:    options,
		Metadata:   args.Metadata,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := govkeeper.NewMsgServerImpl(&c.govKeeper)
	_, err = server.VoteWeighted(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(VoteWeightedEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		args.ProposalId,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (c *Contract) Deposit(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(DepositMethodName)

	var args DepositArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	var amount sdk.Coins
	amount = amount.Add(sdk.Coin{
		Denom:  utils.BaseDenom,
		Amount: sdk.NewIntFromBigInt(args.Amount),
	})

	msg := &govv1.MsgDeposit{
		ProposalId: args.ProposalId,
		Depositor:  sdk.AccAddress(contract.Caller().Bytes()).String(),
		Amount:     amount,
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := govkeeper.NewMsgServerImpl(&c.govKeeper)
	_, err = server.Deposit(ctx, msg)
	if err != nil {
		return nil, err
	}

	if err := c.AddLog(
		evm,
		MustEvent(DepositEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		args.ProposalId,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
