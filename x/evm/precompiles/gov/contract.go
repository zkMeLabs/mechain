package gov

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	accountkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/evmos/evmos/v12/x/evm/types"
)

type Contract struct {
	ctx           sdk.Context
	govKeeper     govkeeper.Keeper
	accountKeeper accountkeeper.AccountKeeper
}

func NewPrecompiledContract(ctx sdk.Context, govKeeper govkeeper.Keeper, accountKeeper accountkeeper.AccountKeeper) *Contract {
	return &Contract{
		ctx:           ctx,
		govKeeper:     govKeeper,
		accountKeeper: accountKeeper,
	}
}

func (c *Contract) Address() common.Address {
	return govAddress
}

func (c *Contract) RequiredGas(input []byte) uint64 {
	method, err := GetMethodByID(input)
	if err != nil {
		return 0
	}

	switch method.Name {
	case LegacySubmitProposalMethodName:
		return LegacySubmitProposalGas
	case SubmitProposalMethodName:
		return SubmitProposalGas
	case VoteMethodName:
		return VoteGas
	case VoteWeightedMethodName:
		return VoteWeightedGas
	case DepositMethodName:
		return DepositGas
	case ProposalMethodName:
		return ProposalGas
	case ProposalsMethodName:
		return ProposalsGas
	case VoteQueryMethodName:
		return VoteQueryGas
	case VotesMethodName:
		return VotesGas
	case DepositQueryMethodName:
		return DepositQueryGas
	case DepositsMethodName:
		return DepositsGas
	case TallyResultMethodName:
		return TallyResultGas
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
		case LegacySubmitProposalMethodName:
			ret, err = c.LegacySubmitProposal(ctx, evm, contract, readonly)
		case SubmitProposalMethodName:
			ret, err = c.SubmitProposal(ctx, evm, contract, readonly)
		case VoteMethodName:
			ret, err = c.Vote(ctx, evm, contract, readonly)
		case VoteWeightedMethodName:
			ret, err = c.VoteWeighted(ctx, evm, contract, readonly)
		case DepositMethodName:
			ret, err = c.Deposit(ctx, evm, contract, readonly)
		case ProposalMethodName:
			ret, err = c.Proposal(ctx, evm, contract, readonly)
		case ProposalsMethodName:
			ret, err = c.Proposals(ctx, evm, contract, readonly)
		case VoteQueryMethodName:
			ret, err = c.VoteQuery(ctx, evm, contract, readonly)
		case VotesMethodName:
			ret, err = c.Votes(ctx, evm, contract, readonly)
		case DepositQueryMethodName:
			ret, err = c.DepositQuery(ctx, evm, contract, readonly)
		case DepositsMethodName:
			ret, err = c.Deposits(ctx, evm, contract, readonly)
		case TallyResultMethodName:
			ret, err = c.TallyResult(ctx, evm, contract, readonly)
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
