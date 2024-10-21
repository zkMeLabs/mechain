package gov

import (
	"bytes"
	"fmt"
	"math/big"

	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v12/types"
)

var (
	govAddress = common.HexToAddress(types.GovAddress)
	govABI     = types.MustABIJson(IGovMetaData.ABI)
)

func GetAddress() common.Address {
	return govAddress
}

func GetMethod(name string) (abi.Method, error) {
	method := govABI.Methods[name]
	if method.ID == nil {
		return abi.Method{}, fmt.Errorf("method %s is not exist", name)
	}
	return method, nil
}

func GetMethodByID(input []byte) (abi.Method, error) {
	if len(input) < 4 {
		return abi.Method{}, fmt.Errorf("input length %d is too short", len(input))
	}
	for _, method := range govABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method, nil
		}
	}
	return abi.Method{}, fmt.Errorf("method id %s is not exist", string(input[:4]))
}

func MustMethod(name string) abi.Method {
	method, err := GetMethod(name)
	if err != nil {
		panic(err)
	}
	return method
}

func MustMethodByID(input []byte) abi.Method {
	if len(input) < 4 {
		panic(fmt.Errorf("input length %d is too short", len(input)))
	}
	for _, method := range govABI.Methods {
		if bytes.Equal(input[:4], method.ID) {
			return method
		}
	}
	panic(fmt.Errorf("method id %s is not exist", string(input[:4])))
}

func GetEvent(name string) (abi.Event, error) {
	event := govABI.Events[name]
	if event.ID == (common.Hash{}) {
		return abi.Event{}, fmt.Errorf("event %s is not exist", name)
	}
	return event, nil
}

func MustEvent(name string) abi.Event {
	event, err := GetEvent(name)
	if err != nil {
		panic(err)
	}
	return event
}

type (
	CoinJson        = Coin
	PageRequestJson = PageRequest
)

type LegacySubmitProposalArgs struct {
	Title          string     `abi:"title"`
	Description    string     `abi:"description"`
	InitialDeposit []CoinJson `abi:"initialDeposit"`
}

// Validate LegacySubmitProposal args
func (args *LegacySubmitProposalArgs) Validate() error {
	for _, deposit := range args.InitialDeposit {
		if deposit.Amount.Sign() < 0 {
			return fmt.Errorf("deposit %s amount is negative %s", deposit.Denom, deposit.Amount.String())
		}
	}

	return nil
}

type SubmitProposalArgs struct {
	Messages       string     `abi:"messages"`
	InitialDeposit []CoinJson `abi:"initialDeposit"`
	Metadata       string     `abi:"metadata"`
	Title          string     `abi:"title"`
	Summary        string     `abi:"summary"`
}

// Validate SubmitProposal args
func (args *SubmitProposalArgs) Validate() error {
	for _, deposit := range args.InitialDeposit {
		if deposit.Amount.Sign() < 0 {
			return fmt.Errorf("deposit %s amount is negative %s", deposit.Denom, deposit.Amount.String())
		}
	}

	return nil
}

type VoteArgs struct {
	ProposalId uint64 `abi:"proposalId"`
	Option     uint8  `abi:"option"`
	Metadata   string `abi:"metadata"`
}

// Validate Vote args
func (args *VoteArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}

	return nil
}

type (
	WeightedVoteOptionJson = WeightedVoteOption
	VoteWeightedArgs       struct {
		ProposalId uint64                   `abi:"proposalId"`
		Options    []WeightedVoteOptionJson `abi:"options"`
		Metadata   string                   `abi:"metadata"`
	}
)

// Validate VoteWeighted args
func (args *VoteWeightedArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}

	return nil
}

type DepositArgs struct {
	ProposalId uint64   `abi:"proposalId"`
	Amount     *big.Int `abi:"amount"`
}

// Validate VoteWeighted args
func (args *DepositArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}

	if args.Amount.Sign() <= 0 {
		return fmt.Errorf("deposit amount is %s, need to greater than 0", args.Amount.String())
	}

	return nil
}

type ProposalArgs struct {
	ProposalId uint64 `abi:"proposalId"`
}

// Validate VoteWeighted args
func (args *ProposalArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}
	return nil
}

type ProposalsArgs struct {
	Status     uint8           `abi:"status"`
	Voter      common.Address  `abi:"voter"`
	Depositor  common.Address  `abi:"depositor"`
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *ProposalsArgs) Validate() error {
	if args.Status > uint8(govv1.ProposalStatus_PROPOSAL_STATUS_FAILED) {
		return fmt.Errorf("invalid status: %d", args.Status)
	}

	return nil
}

// GetStatus returns the proposal status string
func (args *ProposalsArgs) GetStatus() string {
	switch args.Status {
	case 0:
		return ""
	case 1:
		return govv1.ProposalStatus_PROPOSAL_STATUS_UNSPECIFIED.String()
	case 2:
		return govv1.ProposalStatus_PROPOSAL_STATUS_DEPOSIT_PERIOD.String()
	case 3:
		return govv1.ProposalStatus_PROPOSAL_STATUS_VOTING_PERIOD.String()
	case 4:
		return govv1.ProposalStatus_PROPOSAL_STATUS_PASSED.String()
	case 5:
		return govv1.ProposalStatus_PROPOSAL_STATUS_FAILED.String()
	default:
		return ""
	}
	return ""
}

type VoteQueryArgs struct {
	ProposalId uint64         `abi:"proposalId"`
	Voter      common.Address `abi:"voter"`
}

// Validate Vote args
func (args *VoteQueryArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}

	if args.Voter == (common.Address{}) {
		return fmt.Errorf("voter is zero address")
	}

	return nil
}

type VotesArgs struct {
	ProposalId uint64          `abi:"proposalId"`
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *VotesArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}

	return nil
}

type DepositQueryArgs struct {
	ProposalId uint64         `abi:"proposalId"`
	Depositor  common.Address `abi:"depositor"`
}

// Validate validates the args
func (args *DepositQueryArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}

	if args.Depositor == (common.Address{}) {
		return fmt.Errorf("depositor is zero address")
	}

	return nil
}

type DepositsArgs struct {
	ProposalId uint64          `abi:"proposalId"`
	Pagination PageRequestJson `abi:"pagination"`
}

// Validate validates the args
func (args *DepositsArgs) Validate() error {
	if args.ProposalId == 0 {
		return fmt.Errorf("proposal id must greater than 0")
	}

	return nil
}
