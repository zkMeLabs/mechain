// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gov

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Coin is an auto generated low-level Go binding around an user-defined struct.
type Coin struct {
	Denom  string
	Amount *big.Int
}

// DepositData is an auto generated low-level Go binding around an user-defined struct.
type DepositData struct {
	ProposalId uint64
	Depositor  common.Address
	Amount     []Coin
}

// PageRequest is an auto generated low-level Go binding around an user-defined struct.
type PageRequest struct {
	Key        []byte
	Offset     uint64
	Limit      uint64
	CountTotal bool
	Reverse    bool
}

// PageResponse is an auto generated low-level Go binding around an user-defined struct.
type PageResponse struct {
	NextKey []byte
	Total   uint64
}

// Params is an auto generated low-level Go binding around an user-defined struct.
type Params struct {
	MinDeposit                 []Coin
	MaxDepositPeriod           int64
	VotingPeriod               int64
	Quorum                     string
	Threshold                  string
	VetoThreshold              string
	MinInitialDepositRatio     string
	BurnVoteQuorum             bool
	BurnProposalDepositPrevote bool
	BurnVoteVeto               bool
}

// Proposal is an auto generated low-level Go binding around an user-defined struct.
type Proposal struct {
	Id               uint64
	Messages         []string
	Status           uint8
	FinalTallyResult TallyResult
	SubmitTime       int64
	DepositEndTime   int64
	TotalDeposit     []Coin
	VotingStartTime  int64
	VotingEndTime    int64
	Metadata         string
	Title            string
	Summary          string
	Proposer         common.Address
	FailedReason     string
}

// TallyResult is an auto generated low-level Go binding around an user-defined struct.
type TallyResult struct {
	YesCount        string
	AbstainCount    string
	NoCount         string
	NoWithVetoCount string
}

// VoteData is an auto generated low-level Go binding around an user-defined struct.
type VoteData struct {
	ProposalId uint64
	Voter      common.Address
	Options    []WeightedVoteOption
	Metadata   string
}

// WeightedVoteOption is an auto generated low-level Go binding around an user-defined struct.
type WeightedVoteOption struct {
	Option uint8
	Weight string
}

// IGovMetaData contains all meta data concerning the IGov contract.
var IGovMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"LegacySubmitProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"SubmitProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"option\",\"type\":\"uint8\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"VoteWeighted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"internalType\":\"structDepositData\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"deposits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"internalType\":\"structDepositData[]\",\"name\":\"deposits\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"initialDeposit\",\"type\":\"tuple[]\"}],\"name\":\"legacySubmitProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"minDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"int64\",\"name\":\"maxDepositPeriod\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"votingPeriod\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"quorum\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"threshold\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vetoThreshold\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"minInitialDepositRatio\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"burnVoteQuorum\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"burnProposalDepositPrevote\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"burnVoteVeto\",\"type\":\"bool\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"proposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"messages\",\"type\":\"string[]\"},{\"internalType\":\"enumProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structTallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"submitTime\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"depositEndTime\",\"type\":\"int64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"int64\",\"name\":\"votingStartTime\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"votingEndTime\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"failedReason\",\"type\":\"string\"}],\"internalType\":\"structProposal\",\"name\":\"proposal\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"proposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string[]\",\"name\":\"messages\",\"type\":\"string[]\"},{\"internalType\":\"enumProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structTallyResult\",\"name\":\"finalTallyResult\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"submitTime\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"depositEndTime\",\"type\":\"int64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"totalDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"int64\",\"name\":\"votingStartTime\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"votingEndTime\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"failedReason\",\"type\":\"string\"}],\"internalType\":\"structProposal[]\",\"name\":\"proposals\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"messages\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"initialDeposit\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"summary\",\"type\":\"string\"}],\"name\":\"submitProposal\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"}],\"name\":\"tallyResult\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"yesCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"abstainCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noCount\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"noWithVetoCount\",\"type\":\"string\"}],\"internalType\":\"structTallyResult\",\"name\":\"tallyResult\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"enumVoteOption\",\"name\":\"option\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumVoteOption\",\"name\":\"option\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"weight\",\"type\":\"string\"}],\"internalType\":\"structWeightedVoteOption[]\",\"name\":\"options\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structVoteData\",\"name\":\"vote\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"enumVoteOption\",\"name\":\"option\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"weight\",\"type\":\"string\"}],\"internalType\":\"structWeightedVoteOption[]\",\"name\":\"options\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"voteWeighted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"votes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"proposalId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumVoteOption\",\"name\":\"option\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"weight\",\"type\":\"string\"}],\"internalType\":\"structWeightedVoteOption[]\",\"name\":\"options\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structVoteData[]\",\"name\":\"votes\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IGovABI is the input ABI used to generate the binding from.
// Deprecated: Use IGovMetaData.ABI instead.
var IGovABI = IGovMetaData.ABI

// IGov is an auto generated Go binding around an Ethereum contract.
type IGov struct {
	IGovCaller     // Read-only binding to the contract
	IGovTransactor // Write-only binding to the contract
	IGovFilterer   // Log filterer for contract events
}

// IGovCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGovCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGovTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGovFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGovSession struct {
	Contract     *IGov             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGovCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGovCallerSession struct {
	Contract *IGovCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IGovTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGovTransactorSession struct {
	Contract     *IGovTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGovRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGovRaw struct {
	Contract *IGov // Generic contract binding to access the raw methods on
}

// IGovCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGovCallerRaw struct {
	Contract *IGovCaller // Generic read-only contract binding to access the raw methods on
}

// IGovTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGovTransactorRaw struct {
	Contract *IGovTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGov creates a new instance of IGov, bound to a specific deployed contract.
func NewIGov(address common.Address, backend bind.ContractBackend) (*IGov, error) {
	contract, err := bindIGov(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGov{IGovCaller: IGovCaller{contract: contract}, IGovTransactor: IGovTransactor{contract: contract}, IGovFilterer: IGovFilterer{contract: contract}}, nil
}

// NewIGovCaller creates a new read-only instance of IGov, bound to a specific deployed contract.
func NewIGovCaller(address common.Address, caller bind.ContractCaller) (*IGovCaller, error) {
	contract, err := bindIGov(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGovCaller{contract: contract}, nil
}

// NewIGovTransactor creates a new write-only instance of IGov, bound to a specific deployed contract.
func NewIGovTransactor(address common.Address, transactor bind.ContractTransactor) (*IGovTransactor, error) {
	contract, err := bindIGov(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGovTransactor{contract: contract}, nil
}

// NewIGovFilterer creates a new log filterer instance of IGov, bound to a specific deployed contract.
func NewIGovFilterer(address common.Address, filterer bind.ContractFilterer) (*IGovFilterer, error) {
	contract, err := bindIGov(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGovFilterer{contract: contract}, nil
}

// bindIGov binds a generic wrapper to an already deployed contract.
func bindIGov(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGovMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGov *IGovRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGov.Contract.IGovCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGov *IGovRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGov.Contract.IGovTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGov *IGovRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGov.Contract.IGovTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGov *IGovCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGov.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGov *IGovTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGov.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGov *IGovTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGov.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a free data retrieval call binding the contract method 0x576bcd2f.
//
// Solidity: function deposit(uint64 proposalId, address depositor) view returns((uint64,address,(string,uint256)[]) deposit)
func (_IGov *IGovCaller) Deposit(opts *bind.CallOpts, proposalId uint64, depositor common.Address) (DepositData, error) {
	var out []interface{}
	err := _IGov.contract.Call(opts, &out, "deposit", proposalId, depositor)

	if err != nil {
		return *new(DepositData), err
	}

	out0 := *abi.ConvertType(out[0], new(DepositData)).(*DepositData)

	return out0, err

}

// Deposit is a free data retrieval call binding the contract method 0x576bcd2f.
//
// Solidity: function deposit(uint64 proposalId, address depositor) view returns((uint64,address,(string,uint256)[]) deposit)
func (_IGov *IGovSession) Deposit(proposalId uint64, depositor common.Address) (DepositData, error) {
	return _IGov.Contract.Deposit(&_IGov.CallOpts, proposalId, depositor)
}

// Deposit is a free data retrieval call binding the contract method 0x576bcd2f.
//
// Solidity: function deposit(uint64 proposalId, address depositor) view returns((uint64,address,(string,uint256)[]) deposit)
func (_IGov *IGovCallerSession) Deposit(proposalId uint64, depositor common.Address) (DepositData, error) {
	return _IGov.Contract.Deposit(&_IGov.CallOpts, proposalId, depositor)
}

// Deposits is a free data retrieval call binding the contract method 0xeaf16768.
//
// Solidity: function deposits(uint64 proposalId, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(string,uint256)[])[] deposits, (bytes,uint64) pageResponse)
func (_IGov *IGovCaller) Deposits(opts *bind.CallOpts, proposalId uint64, pagination PageRequest) (struct {
	Deposits     []DepositData
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IGov.contract.Call(opts, &out, "deposits", proposalId, pagination)

	outstruct := new(struct {
		Deposits     []DepositData
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposits = *abi.ConvertType(out[0], new([]DepositData)).(*[]DepositData)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xeaf16768.
//
// Solidity: function deposits(uint64 proposalId, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(string,uint256)[])[] deposits, (bytes,uint64) pageResponse)
func (_IGov *IGovSession) Deposits(proposalId uint64, pagination PageRequest) (struct {
	Deposits     []DepositData
	PageResponse PageResponse
}, error) {
	return _IGov.Contract.Deposits(&_IGov.CallOpts, proposalId, pagination)
}

// Deposits is a free data retrieval call binding the contract method 0xeaf16768.
//
// Solidity: function deposits(uint64 proposalId, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(string,uint256)[])[] deposits, (bytes,uint64) pageResponse)
func (_IGov *IGovCallerSession) Deposits(proposalId uint64, pagination PageRequest) (struct {
	Deposits     []DepositData
	PageResponse PageResponse
}, error) {
	return _IGov.Contract.Deposits(&_IGov.CallOpts, proposalId, pagination)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((string,uint256)[],int64,int64,string,string,string,string,bool,bool,bool) params)
func (_IGov *IGovCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _IGov.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((string,uint256)[],int64,int64,string,string,string,string,bool,bool,bool) params)
func (_IGov *IGovSession) Params() (Params, error) {
	return _IGov.Contract.Params(&_IGov.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((string,uint256)[],int64,int64,string,string,string,string,bool,bool,bool) params)
func (_IGov *IGovCallerSession) Params() (Params, error) {
	return _IGov.Contract.Params(&_IGov.CallOpts)
}

// Proposal is a free data retrieval call binding the contract method 0x7afa0aa3.
//
// Solidity: function proposal(uint64 proposalId) view returns((uint64,string[],uint8,(string,string,string,string),int64,int64,(string,uint256)[],int64,int64,string,string,string,address,string) proposal)
func (_IGov *IGovCaller) Proposal(opts *bind.CallOpts, proposalId uint64) (Proposal, error) {
	var out []interface{}
	err := _IGov.contract.Call(opts, &out, "proposal", proposalId)

	if err != nil {
		return *new(Proposal), err
	}

	out0 := *abi.ConvertType(out[0], new(Proposal)).(*Proposal)

	return out0, err

}

// Proposal is a free data retrieval call binding the contract method 0x7afa0aa3.
//
// Solidity: function proposal(uint64 proposalId) view returns((uint64,string[],uint8,(string,string,string,string),int64,int64,(string,uint256)[],int64,int64,string,string,string,address,string) proposal)
func (_IGov *IGovSession) Proposal(proposalId uint64) (Proposal, error) {
	return _IGov.Contract.Proposal(&_IGov.CallOpts, proposalId)
}

// Proposal is a free data retrieval call binding the contract method 0x7afa0aa3.
//
// Solidity: function proposal(uint64 proposalId) view returns((uint64,string[],uint8,(string,string,string,string),int64,int64,(string,uint256)[],int64,int64,string,string,string,address,string) proposal)
func (_IGov *IGovCallerSession) Proposal(proposalId uint64) (Proposal, error) {
	return _IGov.Contract.Proposal(&_IGov.CallOpts, proposalId)
}

// Proposals is a free data retrieval call binding the contract method 0xefb5152c.
//
// Solidity: function proposals(uint8 status, address voter, address depositor, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,string[],uint8,(string,string,string,string),int64,int64,(string,uint256)[],int64,int64,string,string,string,address,string)[] proposals, (bytes,uint64) pageResponse)
func (_IGov *IGovCaller) Proposals(opts *bind.CallOpts, status uint8, voter common.Address, depositor common.Address, pagination PageRequest) (struct {
	Proposals    []Proposal
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IGov.contract.Call(opts, &out, "proposals", status, voter, depositor, pagination)

	outstruct := new(struct {
		Proposals    []Proposal
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposals = *abi.ConvertType(out[0], new([]Proposal)).(*[]Proposal)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0xefb5152c.
//
// Solidity: function proposals(uint8 status, address voter, address depositor, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,string[],uint8,(string,string,string,string),int64,int64,(string,uint256)[],int64,int64,string,string,string,address,string)[] proposals, (bytes,uint64) pageResponse)
func (_IGov *IGovSession) Proposals(status uint8, voter common.Address, depositor common.Address, pagination PageRequest) (struct {
	Proposals    []Proposal
	PageResponse PageResponse
}, error) {
	return _IGov.Contract.Proposals(&_IGov.CallOpts, status, voter, depositor, pagination)
}

// Proposals is a free data retrieval call binding the contract method 0xefb5152c.
//
// Solidity: function proposals(uint8 status, address voter, address depositor, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,string[],uint8,(string,string,string,string),int64,int64,(string,uint256)[],int64,int64,string,string,string,address,string)[] proposals, (bytes,uint64) pageResponse)
func (_IGov *IGovCallerSession) Proposals(status uint8, voter common.Address, depositor common.Address, pagination PageRequest) (struct {
	Proposals    []Proposal
	PageResponse PageResponse
}, error) {
	return _IGov.Contract.Proposals(&_IGov.CallOpts, status, voter, depositor, pagination)
}

// TallyResult is a free data retrieval call binding the contract method 0x101146be.
//
// Solidity: function tallyResult(uint64 proposalId) view returns((string,string,string,string) tallyResult)
func (_IGov *IGovCaller) TallyResult(opts *bind.CallOpts, proposalId uint64) (TallyResult, error) {
	var out []interface{}
	err := _IGov.contract.Call(opts, &out, "tallyResult", proposalId)

	if err != nil {
		return *new(TallyResult), err
	}

	out0 := *abi.ConvertType(out[0], new(TallyResult)).(*TallyResult)

	return out0, err

}

// TallyResult is a free data retrieval call binding the contract method 0x101146be.
//
// Solidity: function tallyResult(uint64 proposalId) view returns((string,string,string,string) tallyResult)
func (_IGov *IGovSession) TallyResult(proposalId uint64) (TallyResult, error) {
	return _IGov.Contract.TallyResult(&_IGov.CallOpts, proposalId)
}

// TallyResult is a free data retrieval call binding the contract method 0x101146be.
//
// Solidity: function tallyResult(uint64 proposalId) view returns((string,string,string,string) tallyResult)
func (_IGov *IGovCallerSession) TallyResult(proposalId uint64) (TallyResult, error) {
	return _IGov.Contract.TallyResult(&_IGov.CallOpts, proposalId)
}

// Vote0 is a free data retrieval call binding the contract method 0xca469089.
//
// Solidity: function vote(uint64 proposalId, address voter) view returns((uint64,address,(uint8,string)[],string) vote)
func (_IGov *IGovCaller) Vote0(opts *bind.CallOpts, proposalId uint64, voter common.Address) (VoteData, error) {
	var out []interface{}
	err := _IGov.contract.Call(opts, &out, "vote0", proposalId, voter)

	if err != nil {
		return *new(VoteData), err
	}

	out0 := *abi.ConvertType(out[0], new(VoteData)).(*VoteData)

	return out0, err

}

// Vote0 is a free data retrieval call binding the contract method 0xca469089.
//
// Solidity: function vote(uint64 proposalId, address voter) view returns((uint64,address,(uint8,string)[],string) vote)
func (_IGov *IGovSession) Vote0(proposalId uint64, voter common.Address) (VoteData, error) {
	return _IGov.Contract.Vote0(&_IGov.CallOpts, proposalId, voter)
}

// Vote0 is a free data retrieval call binding the contract method 0xca469089.
//
// Solidity: function vote(uint64 proposalId, address voter) view returns((uint64,address,(uint8,string)[],string) vote)
func (_IGov *IGovCallerSession) Vote0(proposalId uint64, voter common.Address) (VoteData, error) {
	return _IGov.Contract.Vote0(&_IGov.CallOpts, proposalId, voter)
}

// Votes is a free data retrieval call binding the contract method 0x8f077f11.
//
// Solidity: function votes(uint64 proposalId, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(uint8,string)[],string)[] votes, (bytes,uint64) pageResponse)
func (_IGov *IGovCaller) Votes(opts *bind.CallOpts, proposalId uint64, pagination PageRequest) (struct {
	Votes        []VoteData
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IGov.contract.Call(opts, &out, "votes", proposalId, pagination)

	outstruct := new(struct {
		Votes        []VoteData
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Votes = *abi.ConvertType(out[0], new([]VoteData)).(*[]VoteData)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Votes is a free data retrieval call binding the contract method 0x8f077f11.
//
// Solidity: function votes(uint64 proposalId, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(uint8,string)[],string)[] votes, (bytes,uint64) pageResponse)
func (_IGov *IGovSession) Votes(proposalId uint64, pagination PageRequest) (struct {
	Votes        []VoteData
	PageResponse PageResponse
}, error) {
	return _IGov.Contract.Votes(&_IGov.CallOpts, proposalId, pagination)
}

// Votes is a free data retrieval call binding the contract method 0x8f077f11.
//
// Solidity: function votes(uint64 proposalId, (bytes,uint64,uint64,bool,bool) pagination) view returns((uint64,address,(uint8,string)[],string)[] votes, (bytes,uint64) pageResponse)
func (_IGov *IGovCallerSession) Votes(proposalId uint64, pagination PageRequest) (struct {
	Votes        []VoteData
	PageResponse PageResponse
}, error) {
	return _IGov.Contract.Votes(&_IGov.CallOpts, proposalId, pagination)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6170c4b1.
//
// Solidity: function deposit(uint64 proposalId, uint256 amount) returns(bool success)
func (_IGov *IGovTransactor) Deposit0(opts *bind.TransactOpts, proposalId uint64, amount *big.Int) (*types.Transaction, error) {
	return _IGov.contract.Transact(opts, "deposit0", proposalId, amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6170c4b1.
//
// Solidity: function deposit(uint64 proposalId, uint256 amount) returns(bool success)
func (_IGov *IGovSession) Deposit0(proposalId uint64, amount *big.Int) (*types.Transaction, error) {
	return _IGov.Contract.Deposit0(&_IGov.TransactOpts, proposalId, amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6170c4b1.
//
// Solidity: function deposit(uint64 proposalId, uint256 amount) returns(bool success)
func (_IGov *IGovTransactorSession) Deposit0(proposalId uint64, amount *big.Int) (*types.Transaction, error) {
	return _IGov.Contract.Deposit0(&_IGov.TransactOpts, proposalId, amount)
}

// LegacySubmitProposal is a paid mutator transaction binding the contract method 0xad5aa33d.
//
// Solidity: function legacySubmitProposal(string title, string description, (string,uint256)[] initialDeposit) returns(uint64 proposalId)
func (_IGov *IGovTransactor) LegacySubmitProposal(opts *bind.TransactOpts, title string, description string, initialDeposit []Coin) (*types.Transaction, error) {
	return _IGov.contract.Transact(opts, "legacySubmitProposal", title, description, initialDeposit)
}

// LegacySubmitProposal is a paid mutator transaction binding the contract method 0xad5aa33d.
//
// Solidity: function legacySubmitProposal(string title, string description, (string,uint256)[] initialDeposit) returns(uint64 proposalId)
func (_IGov *IGovSession) LegacySubmitProposal(title string, description string, initialDeposit []Coin) (*types.Transaction, error) {
	return _IGov.Contract.LegacySubmitProposal(&_IGov.TransactOpts, title, description, initialDeposit)
}

// LegacySubmitProposal is a paid mutator transaction binding the contract method 0xad5aa33d.
//
// Solidity: function legacySubmitProposal(string title, string description, (string,uint256)[] initialDeposit) returns(uint64 proposalId)
func (_IGov *IGovTransactorSession) LegacySubmitProposal(title string, description string, initialDeposit []Coin) (*types.Transaction, error) {
	return _IGov.Contract.LegacySubmitProposal(&_IGov.TransactOpts, title, description, initialDeposit)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x077a94c9.
//
// Solidity: function submitProposal(string messages, (string,uint256)[] initialDeposit, string metadata, string title, string summary) returns(uint64 proposalId)
func (_IGov *IGovTransactor) SubmitProposal(opts *bind.TransactOpts, messages string, initialDeposit []Coin, metadata string, title string, summary string) (*types.Transaction, error) {
	return _IGov.contract.Transact(opts, "submitProposal", messages, initialDeposit, metadata, title, summary)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x077a94c9.
//
// Solidity: function submitProposal(string messages, (string,uint256)[] initialDeposit, string metadata, string title, string summary) returns(uint64 proposalId)
func (_IGov *IGovSession) SubmitProposal(messages string, initialDeposit []Coin, metadata string, title string, summary string) (*types.Transaction, error) {
	return _IGov.Contract.SubmitProposal(&_IGov.TransactOpts, messages, initialDeposit, metadata, title, summary)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x077a94c9.
//
// Solidity: function submitProposal(string messages, (string,uint256)[] initialDeposit, string metadata, string title, string summary) returns(uint64 proposalId)
func (_IGov *IGovTransactorSession) SubmitProposal(messages string, initialDeposit []Coin, metadata string, title string, summary string) (*types.Transaction, error) {
	return _IGov.Contract.SubmitProposal(&_IGov.TransactOpts, messages, initialDeposit, metadata, title, summary)
}

// Vote is a paid mutator transaction binding the contract method 0x528783d5.
//
// Solidity: function vote(uint64 proposalId, uint8 option, string metadata) returns(bool success)
func (_IGov *IGovTransactor) Vote(opts *bind.TransactOpts, proposalId uint64, option uint8, metadata string) (*types.Transaction, error) {
	return _IGov.contract.Transact(opts, "vote", proposalId, option, metadata)
}

// Vote is a paid mutator transaction binding the contract method 0x528783d5.
//
// Solidity: function vote(uint64 proposalId, uint8 option, string metadata) returns(bool success)
func (_IGov *IGovSession) Vote(proposalId uint64, option uint8, metadata string) (*types.Transaction, error) {
	return _IGov.Contract.Vote(&_IGov.TransactOpts, proposalId, option, metadata)
}

// Vote is a paid mutator transaction binding the contract method 0x528783d5.
//
// Solidity: function vote(uint64 proposalId, uint8 option, string metadata) returns(bool success)
func (_IGov *IGovTransactorSession) Vote(proposalId uint64, option uint8, metadata string) (*types.Transaction, error) {
	return _IGov.Contract.Vote(&_IGov.TransactOpts, proposalId, option, metadata)
}

// VoteWeighted is a paid mutator transaction binding the contract method 0xc1cffef3.
//
// Solidity: function voteWeighted(uint64 proposalId, (uint8,string)[] options, string metadata) returns(bool success)
func (_IGov *IGovTransactor) VoteWeighted(opts *bind.TransactOpts, proposalId uint64, options []WeightedVoteOption, metadata string) (*types.Transaction, error) {
	return _IGov.contract.Transact(opts, "voteWeighted", proposalId, options, metadata)
}

// VoteWeighted is a paid mutator transaction binding the contract method 0xc1cffef3.
//
// Solidity: function voteWeighted(uint64 proposalId, (uint8,string)[] options, string metadata) returns(bool success)
func (_IGov *IGovSession) VoteWeighted(proposalId uint64, options []WeightedVoteOption, metadata string) (*types.Transaction, error) {
	return _IGov.Contract.VoteWeighted(&_IGov.TransactOpts, proposalId, options, metadata)
}

// VoteWeighted is a paid mutator transaction binding the contract method 0xc1cffef3.
//
// Solidity: function voteWeighted(uint64 proposalId, (uint8,string)[] options, string metadata) returns(bool success)
func (_IGov *IGovTransactorSession) VoteWeighted(proposalId uint64, options []WeightedVoteOption, metadata string) (*types.Transaction, error) {
	return _IGov.Contract.VoteWeighted(&_IGov.TransactOpts, proposalId, options, metadata)
}

// IGovDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IGov contract.
type IGovDepositIterator struct {
	Event *IGovDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovDeposit represents a Deposit event raised by the IGov contract.
type IGovDeposit struct {
	Depositor  common.Address
	ProposalId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x0ee94a97c7c69ce2eb8cfb09bacc78d63a73b5e0fbed0d13a079190ff876ae3a.
//
// Solidity: event Deposit(address indexed depositor, uint64 proposalId)
func (_IGov *IGovFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address) (*IGovDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _IGov.contract.FilterLogs(opts, "Deposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return &IGovDepositIterator{contract: _IGov.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x0ee94a97c7c69ce2eb8cfb09bacc78d63a73b5e0fbed0d13a079190ff876ae3a.
//
// Solidity: event Deposit(address indexed depositor, uint64 proposalId)
func (_IGov *IGovFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IGovDeposit, depositor []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _IGov.contract.WatchLogs(opts, "Deposit", depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovDeposit)
				if err := _IGov.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x0ee94a97c7c69ce2eb8cfb09bacc78d63a73b5e0fbed0d13a079190ff876ae3a.
//
// Solidity: event Deposit(address indexed depositor, uint64 proposalId)
func (_IGov *IGovFilterer) ParseDeposit(log types.Log) (*IGovDeposit, error) {
	event := new(IGovDeposit)
	if err := _IGov.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGovLegacySubmitProposalIterator is returned from FilterLegacySubmitProposal and is used to iterate over the raw logs and unpacked data for LegacySubmitProposal events raised by the IGov contract.
type IGovLegacySubmitProposalIterator struct {
	Event *IGovLegacySubmitProposal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovLegacySubmitProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovLegacySubmitProposal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovLegacySubmitProposal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovLegacySubmitProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovLegacySubmitProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovLegacySubmitProposal represents a LegacySubmitProposal event raised by the IGov contract.
type IGovLegacySubmitProposal struct {
	Proposer   common.Address
	ProposalId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLegacySubmitProposal is a free log retrieval operation binding the contract event 0x749a006802f9b6dde9b6ee9f8f3a317d8a98cb9017684a7b59bfcf21bf93f477.
//
// Solidity: event LegacySubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) FilterLegacySubmitProposal(opts *bind.FilterOpts, proposer []common.Address) (*IGovLegacySubmitProposalIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _IGov.contract.FilterLogs(opts, "LegacySubmitProposal", proposerRule)
	if err != nil {
		return nil, err
	}
	return &IGovLegacySubmitProposalIterator{contract: _IGov.contract, event: "LegacySubmitProposal", logs: logs, sub: sub}, nil
}

// WatchLegacySubmitProposal is a free log subscription operation binding the contract event 0x749a006802f9b6dde9b6ee9f8f3a317d8a98cb9017684a7b59bfcf21bf93f477.
//
// Solidity: event LegacySubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) WatchLegacySubmitProposal(opts *bind.WatchOpts, sink chan<- *IGovLegacySubmitProposal, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _IGov.contract.WatchLogs(opts, "LegacySubmitProposal", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovLegacySubmitProposal)
				if err := _IGov.contract.UnpackLog(event, "LegacySubmitProposal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLegacySubmitProposal is a log parse operation binding the contract event 0x749a006802f9b6dde9b6ee9f8f3a317d8a98cb9017684a7b59bfcf21bf93f477.
//
// Solidity: event LegacySubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) ParseLegacySubmitProposal(log types.Log) (*IGovLegacySubmitProposal, error) {
	event := new(IGovLegacySubmitProposal)
	if err := _IGov.contract.UnpackLog(event, "LegacySubmitProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGovSubmitProposalIterator is returned from FilterSubmitProposal and is used to iterate over the raw logs and unpacked data for SubmitProposal events raised by the IGov contract.
type IGovSubmitProposalIterator struct {
	Event *IGovSubmitProposal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovSubmitProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovSubmitProposal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovSubmitProposal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovSubmitProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovSubmitProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovSubmitProposal represents a SubmitProposal event raised by the IGov contract.
type IGovSubmitProposal struct {
	Proposer   common.Address
	ProposalId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmitProposal is a free log retrieval operation binding the contract event 0xf49a3a8232aff8553333cfd734e3a7ef1ab4764cd0494eb145216773b64bf349.
//
// Solidity: event SubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) FilterSubmitProposal(opts *bind.FilterOpts, proposer []common.Address) (*IGovSubmitProposalIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _IGov.contract.FilterLogs(opts, "SubmitProposal", proposerRule)
	if err != nil {
		return nil, err
	}
	return &IGovSubmitProposalIterator{contract: _IGov.contract, event: "SubmitProposal", logs: logs, sub: sub}, nil
}

// WatchSubmitProposal is a free log subscription operation binding the contract event 0xf49a3a8232aff8553333cfd734e3a7ef1ab4764cd0494eb145216773b64bf349.
//
// Solidity: event SubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) WatchSubmitProposal(opts *bind.WatchOpts, sink chan<- *IGovSubmitProposal, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _IGov.contract.WatchLogs(opts, "SubmitProposal", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovSubmitProposal)
				if err := _IGov.contract.UnpackLog(event, "SubmitProposal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubmitProposal is a log parse operation binding the contract event 0xf49a3a8232aff8553333cfd734e3a7ef1ab4764cd0494eb145216773b64bf349.
//
// Solidity: event SubmitProposal(address indexed proposer, uint64 proposalId)
func (_IGov *IGovFilterer) ParseSubmitProposal(log types.Log) (*IGovSubmitProposal, error) {
	event := new(IGovSubmitProposal)
	if err := _IGov.contract.UnpackLog(event, "SubmitProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGovVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the IGov contract.
type IGovVoteIterator struct {
	Event *IGovVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovVote represents a Vote event raised by the IGov contract.
type IGovVote struct {
	Voter      common.Address
	ProposalId uint64
	Option     uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x71c096cfbbce3e73fe1d1e5943da8fcbdcd2ba95519bfa456d51c282c575c64a.
//
// Solidity: event Vote(address indexed voter, uint64 proposalId, uint8 option)
func (_IGov *IGovFilterer) FilterVote(opts *bind.FilterOpts, voter []common.Address) (*IGovVoteIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IGov.contract.FilterLogs(opts, "Vote", voterRule)
	if err != nil {
		return nil, err
	}
	return &IGovVoteIterator{contract: _IGov.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x71c096cfbbce3e73fe1d1e5943da8fcbdcd2ba95519bfa456d51c282c575c64a.
//
// Solidity: event Vote(address indexed voter, uint64 proposalId, uint8 option)
func (_IGov *IGovFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *IGovVote, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IGov.contract.WatchLogs(opts, "Vote", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovVote)
				if err := _IGov.contract.UnpackLog(event, "Vote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVote is a log parse operation binding the contract event 0x71c096cfbbce3e73fe1d1e5943da8fcbdcd2ba95519bfa456d51c282c575c64a.
//
// Solidity: event Vote(address indexed voter, uint64 proposalId, uint8 option)
func (_IGov *IGovFilterer) ParseVote(log types.Log) (*IGovVote, error) {
	event := new(IGovVote)
	if err := _IGov.contract.UnpackLog(event, "Vote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGovVoteWeightedIterator is returned from FilterVoteWeighted and is used to iterate over the raw logs and unpacked data for VoteWeighted events raised by the IGov contract.
type IGovVoteWeightedIterator struct {
	Event *IGovVoteWeighted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovVoteWeightedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovVoteWeighted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovVoteWeighted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovVoteWeightedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovVoteWeightedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovVoteWeighted represents a VoteWeighted event raised by the IGov contract.
type IGovVoteWeighted struct {
	Voter      common.Address
	ProposalId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteWeighted is a free log retrieval operation binding the contract event 0xba05d71c5068c3a8723b434cd80b62655c4cb23452233f495ad020e2a0bcabf3.
//
// Solidity: event VoteWeighted(address indexed voter, uint64 proposalId)
func (_IGov *IGovFilterer) FilterVoteWeighted(opts *bind.FilterOpts, voter []common.Address) (*IGovVoteWeightedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IGov.contract.FilterLogs(opts, "VoteWeighted", voterRule)
	if err != nil {
		return nil, err
	}
	return &IGovVoteWeightedIterator{contract: _IGov.contract, event: "VoteWeighted", logs: logs, sub: sub}, nil
}

// WatchVoteWeighted is a free log subscription operation binding the contract event 0xba05d71c5068c3a8723b434cd80b62655c4cb23452233f495ad020e2a0bcabf3.
//
// Solidity: event VoteWeighted(address indexed voter, uint64 proposalId)
func (_IGov *IGovFilterer) WatchVoteWeighted(opts *bind.WatchOpts, sink chan<- *IGovVoteWeighted, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IGov.contract.WatchLogs(opts, "VoteWeighted", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovVoteWeighted)
				if err := _IGov.contract.UnpackLog(event, "VoteWeighted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteWeighted is a log parse operation binding the contract event 0xba05d71c5068c3a8723b434cd80b62655c4cb23452233f495ad020e2a0bcabf3.
//
// Solidity: event VoteWeighted(address indexed voter, uint64 proposalId)
func (_IGov *IGovFilterer) ParseVoteWeighted(log types.Log) (*IGovVoteWeighted, error) {
	event := new(IGovVoteWeighted)
	if err := _IGov.contract.UnpackLog(event, "VoteWeighted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
