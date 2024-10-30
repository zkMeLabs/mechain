// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// BlockID is an auto generated low-level Go binding around an user-defined struct.
type BlockID struct {
	Hash          string
	PartSetHeader PartSetHeader
}

// Coin is an auto generated low-level Go binding around an user-defined struct.
type Coin struct {
	Denom  string
	Amount *big.Int
}

// Commission is an auto generated low-level Go binding around an user-defined struct.
type Commission struct {
	CommissionRates CommissionRates
	UpdateTime      int64
}

// CommissionRates is an auto generated low-level Go binding around an user-defined struct.
type CommissionRates struct {
	Rate          *big.Int
	MaxRate       *big.Int
	MaxChangeRate *big.Int
}

// Consensus is an auto generated low-level Go binding around an user-defined struct.
type Consensus struct {
	Block uint64
	App   uint64
}

// Dec is an auto generated low-level Go binding around an user-defined struct.
type Dec struct {
	Amount    *big.Int
	Precision uint8
}

// Delegation is an auto generated low-level Go binding around an user-defined struct.
type Delegation struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Shares           Dec
}

// DelegationResponse is an auto generated low-level Go binding around an user-defined struct.
type DelegationResponse struct {
	Delegation Delegation
	Balance    Coin
}

// Description is an auto generated low-level Go binding around an user-defined struct.
type Description struct {
	Moniker         string
	Identity        string
	Website         string
	SecurityContact string
	Details         string
}

// Header is an auto generated low-level Go binding around an user-defined struct.
type Header struct {
	Version            Consensus
	ChainId            string
	Height             int64
	Time               int64
	LastBlockId        BlockID
	LastCommitHash     string
	DataHash           string
	ValidatorsHash     string
	NextValidatorsHash string
	ConsensusHash      string
	AppHash            string
	LastResultsHash    string
	EvidenceHash       string
	ProposerAddress    string
}

// HistoricalInfo is an auto generated low-level Go binding around an user-defined struct.
type HistoricalInfo struct {
	Header Header
	Valset []Validator
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
	UnbondingTime     int64
	MaxValidators     uint32
	MaxEntries        uint32
	HistoricalEntries uint32
	BondDenom         string
	MinCommissionRate *big.Int
}

// PartSetHeader is an auto generated low-level Go binding around an user-defined struct.
type PartSetHeader struct {
	Total uint32
	Hash  string
}

// Pool is an auto generated low-level Go binding around an user-defined struct.
type Pool struct {
	NotBondedTokens *big.Int
	BondedTokens    *big.Int
}

// Redelegation is an auto generated low-level Go binding around an user-defined struct.
type Redelegation struct {
	DelegatorAddress    common.Address
	ValidatorSrcAddress common.Address
	ValidatorDstAddress common.Address
	Entries             []RedelegationEntry
}

// RedelegationEntry is an auto generated low-level Go binding around an user-defined struct.
type RedelegationEntry struct {
	CreationHeight int64
	CompletionTime int64
	InitialBalance *big.Int
	ShareDst       *big.Int
}

// RedelegationEntryResponse is an auto generated low-level Go binding around an user-defined struct.
type RedelegationEntryResponse struct {
	RedelegationEntry RedelegationEntry
	Balance           *big.Int
}

// RedelegationResponse is an auto generated low-level Go binding around an user-defined struct.
type RedelegationResponse struct {
	Redelegation Redelegation
	Entries      []RedelegationEntryResponse
}

// UnbondingDelegation is an auto generated low-level Go binding around an user-defined struct.
type UnbondingDelegation struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Entries          []UnbondingDelegationEntry
}

// UnbondingDelegationEntry is an auto generated low-level Go binding around an user-defined struct.
type UnbondingDelegationEntry struct {
	CreationHeight int64
	CompletionTime int64
	InitialBalance *big.Int
	Balance        *big.Int
}

// Validator is an auto generated low-level Go binding around an user-defined struct.
type Validator struct {
	OperatorAddress         common.Address
	ConsensusPubkey         string
	Jailed                  bool
	Status                  uint8
	Tokens                  *big.Int
	DelegatorShares         *big.Int
	Description             Description
	UnbondingHeight         int64
	UnbondingTime           int64
	Commission              Commission
	MinSelfDelegation       *big.Int
	UnbondingOnHoldRefCount int64
	UnbondingIds            []uint64
	SelfDelAddress          string
	RelayerAddress          string
	ChallengerAddress       string
	BlsKey                  string
}

// IStakingMetaData contains all meta data concerning the IStaking contract.
var IStakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"}],\"name\":\"CancelUnbondingDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"commissionRate\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"minSelfDelegation\",\"type\":\"int256\"}],\"name\":\"EditValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorSrcAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorDstAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"name\":\"Redelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"name\":\"Undelegate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationHeight\",\"type\":\"uint256\"}],\"name\":\"cancelUnbondingDelegation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"delegation\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDec\",\"name\":\"shares\",\"type\":\"tuple\"}],\"internalType\":\"structDelegation\",\"name\":\"delegation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"balance\",\"type\":\"tuple\"}],\"internalType\":\"structDelegationResponse\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"delegatorDelegations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDec\",\"name\":\"shares\",\"type\":\"tuple\"}],\"internalType\":\"structDelegation\",\"name\":\"delegation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"balance\",\"type\":\"tuple\"}],\"internalType\":\"structDelegationResponse[]\",\"name\":\"response\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"delegatorUnbondingDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structUnbondingDelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structUnbondingDelegation[]\",\"name\":\"response\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"delegatorValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"updateTime\",\"type\":\"int64\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string\",\"name\":\"selfDelAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"challengerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsKey\",\"type\":\"string\"}],\"internalType\":\"structValidator\",\"name\":\"validator\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"delegatorValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"updateTime\",\"type\":\"int64\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string\",\"name\":\"selfDelAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"challengerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsKey\",\"type\":\"string\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int256\",\"name\":\"commissionRate\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"minSelfDelegation\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"challengerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"blsKey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsProof\",\"type\":\"string\"}],\"name\":\"editValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"height\",\"type\":\"int64\"}],\"name\":\"historicalInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"app\",\"type\":\"uint64\"}],\"internalType\":\"structConsensus\",\"name\":\"version\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"chainId\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"height\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"time\",\"type\":\"int64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"total\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"internalType\":\"structPartSetHeader\",\"name\":\"partSetHeader\",\"type\":\"tuple\"}],\"internalType\":\"structBlockID\",\"name\":\"lastBlockId\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"lastCommitHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"validatorsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"nextValidatorsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"consensusHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"appHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lastResultsHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evidenceHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proposerAddress\",\"type\":\"string\"}],\"internalType\":\"structHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"updateTime\",\"type\":\"int64\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string\",\"name\":\"selfDelAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"challengerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsKey\",\"type\":\"string\"}],\"internalType\":\"structValidator[]\",\"name\":\"valset\",\"type\":\"tuple[]\"}],\"internalType\":\"structHistoricalInfo\",\"name\":\"historicalInfo\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"internalType\":\"uint32\",\"name\":\"maxValidators\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxEntries\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"historicalEntries\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"bondDenom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minCommissionRate\",\"type\":\"uint256\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"notBondedTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bondedTokens\",\"type\":\"uint256\"}],\"internalType\":\"structPool\",\"name\":\"pool\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorSrcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorDstAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcValidatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstValidatorAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"redelegations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorSrcAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorDstAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shareDst\",\"type\":\"uint256\"}],\"internalType\":\"structRedelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structRedelegation\",\"name\":\"redelegation\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shareDst\",\"type\":\"uint256\"}],\"internalType\":\"structRedelegationEntry\",\"name\":\"redelegationEntry\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structRedelegationEntryResponse[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structRedelegationResponse[]\",\"name\":\"redelegationResponses\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"unbondingDelegation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structUnbondingDelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structUnbondingDelegation\",\"name\":\"response\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"completionTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"validator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"updateTime\",\"type\":\"int64\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string\",\"name\":\"selfDelAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"challengerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsKey\",\"type\":\"string\"}],\"internalType\":\"structValidator\",\"name\":\"validator\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"validatorDelegations\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structDec\",\"name\":\"shares\",\"type\":\"tuple\"}],\"internalType\":\"structDelegation\",\"name\":\"delegation\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"balance\",\"type\":\"tuple\"}],\"internalType\":\"structDelegationResponse[]\",\"name\":\"response\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"validatorUnbondingDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"delegatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"creationHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"completionTime\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structUnbondingDelegationEntry[]\",\"name\":\"entries\",\"type\":\"tuple[]\"}],\"internalType\":\"structUnbondingDelegation[]\",\"name\":\"response\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"validators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"consensusPubkey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"enumBondStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorShares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"securityContact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"unbondingHeight\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"unbondingTime\",\"type\":\"int64\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxChangeRate\",\"type\":\"uint256\"}],\"internalType\":\"structCommissionRates\",\"name\":\"commissionRates\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"updateTime\",\"type\":\"int64\"}],\"internalType\":\"structCommission\",\"name\":\"commission\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"int64\",\"name\":\"unbondingOnHoldRefCount\",\"type\":\"int64\"},{\"internalType\":\"uint64[]\",\"name\":\"unbondingIds\",\"type\":\"uint64[]\"},{\"internalType\":\"string\",\"name\":\"selfDelAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"relayerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"challengerAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"blsKey\",\"type\":\"string\"}],\"internalType\":\"structValidator[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingMetaData.ABI instead.
var IStakingABI = IStakingMetaData.ABI

// IStaking is an auto generated Go binding around an Ethereum contract.
type IStaking struct {
	IStakingCaller     // Read-only binding to the contract
	IStakingTransactor // Write-only binding to the contract
	IStakingFilterer   // Log filterer for contract events
}

// IStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingSession struct {
	Contract     *IStaking         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingCallerSession struct {
	Contract *IStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingTransactorSession struct {
	Contract     *IStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingRaw struct {
	Contract *IStaking // Generic contract binding to access the raw methods on
}

// IStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingCallerRaw struct {
	Contract *IStakingCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingTransactorRaw struct {
	Contract *IStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStaking creates a new instance of IStaking, bound to a specific deployed contract.
func NewIStaking(address common.Address, backend bind.ContractBackend) (*IStaking, error) {
	contract, err := bindIStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStaking{IStakingCaller: IStakingCaller{contract: contract}, IStakingTransactor: IStakingTransactor{contract: contract}, IStakingFilterer: IStakingFilterer{contract: contract}}, nil
}

// NewIStakingCaller creates a new read-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingCaller(address common.Address, caller bind.ContractCaller) (*IStakingCaller, error) {
	contract, err := bindIStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingCaller{contract: contract}, nil
}

// NewIStakingTransactor creates a new write-only instance of IStaking, bound to a specific deployed contract.
func NewIStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingTransactor, error) {
	contract, err := bindIStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingTransactor{contract: contract}, nil
}

// NewIStakingFilterer creates a new log filterer instance of IStaking, bound to a specific deployed contract.
func NewIStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingFilterer, error) {
	contract, err := bindIStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingFilterer{contract: contract}, nil
}

// bindIStaking binds a generic wrapper to an already deployed contract.
func bindIStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.IStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.IStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStaking *IStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStaking *IStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStaking *IStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStaking.Contract.contract.Transact(opts, method, params...)
}

// Delegation is a free data retrieval call binding the contract method 0x046d3307.
//
// Solidity: function delegation(address delegatorAddr, address validatorAddr) view returns(((address,address,(uint256,uint8)),(string,uint256)) response)
func (_IStaking *IStakingCaller) Delegation(opts *bind.CallOpts, delegatorAddr common.Address, validatorAddr common.Address) (DelegationResponse, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "delegation", delegatorAddr, validatorAddr)

	if err != nil {
		return *new(DelegationResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(DelegationResponse)).(*DelegationResponse)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0x046d3307.
//
// Solidity: function delegation(address delegatorAddr, address validatorAddr) view returns(((address,address,(uint256,uint8)),(string,uint256)) response)
func (_IStaking *IStakingSession) Delegation(delegatorAddr common.Address, validatorAddr common.Address) (DelegationResponse, error) {
	return _IStaking.Contract.Delegation(&_IStaking.CallOpts, delegatorAddr, validatorAddr)
}

// Delegation is a free data retrieval call binding the contract method 0x046d3307.
//
// Solidity: function delegation(address delegatorAddr, address validatorAddr) view returns(((address,address,(uint256,uint8)),(string,uint256)) response)
func (_IStaking *IStakingCallerSession) Delegation(delegatorAddr common.Address, validatorAddr common.Address) (DelegationResponse, error) {
	return _IStaking.Contract.Delegation(&_IStaking.CallOpts, delegatorAddr, validatorAddr)
}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x256bd907.
//
// Solidity: function delegatorDelegations(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,(uint256,uint8)),(string,uint256))[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCaller) DelegatorDelegations(opts *bind.CallOpts, delegatorAddr common.Address, pagination PageRequest) (struct {
	Response     []DelegationResponse
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "delegatorDelegations", delegatorAddr, pagination)

	outstruct := new(struct {
		Response     []DelegationResponse
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Response = *abi.ConvertType(out[0], new([]DelegationResponse)).(*[]DelegationResponse)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x256bd907.
//
// Solidity: function delegatorDelegations(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,(uint256,uint8)),(string,uint256))[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingSession) DelegatorDelegations(delegatorAddr common.Address, pagination PageRequest) (struct {
	Response     []DelegationResponse
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.DelegatorDelegations(&_IStaking.CallOpts, delegatorAddr, pagination)
}

// DelegatorDelegations is a free data retrieval call binding the contract method 0x256bd907.
//
// Solidity: function delegatorDelegations(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,(uint256,uint8)),(string,uint256))[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCallerSession) DelegatorDelegations(delegatorAddr common.Address, pagination PageRequest) (struct {
	Response     []DelegationResponse
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.DelegatorDelegations(&_IStaking.CallOpts, delegatorAddr, pagination)
}

// DelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0x155fd3ff.
//
// Solidity: function delegatorUnbondingDelegations(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,int64,uint256,uint256)[])[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCaller) DelegatorUnbondingDelegations(opts *bind.CallOpts, delegatorAddr common.Address, pagination PageRequest) (struct {
	Response     []UnbondingDelegation
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "delegatorUnbondingDelegations", delegatorAddr, pagination)

	outstruct := new(struct {
		Response     []UnbondingDelegation
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Response = *abi.ConvertType(out[0], new([]UnbondingDelegation)).(*[]UnbondingDelegation)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// DelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0x155fd3ff.
//
// Solidity: function delegatorUnbondingDelegations(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,int64,uint256,uint256)[])[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingSession) DelegatorUnbondingDelegations(delegatorAddr common.Address, pagination PageRequest) (struct {
	Response     []UnbondingDelegation
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.DelegatorUnbondingDelegations(&_IStaking.CallOpts, delegatorAddr, pagination)
}

// DelegatorUnbondingDelegations is a free data retrieval call binding the contract method 0x155fd3ff.
//
// Solidity: function delegatorUnbondingDelegations(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,int64,uint256,uint256)[])[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCallerSession) DelegatorUnbondingDelegations(delegatorAddr common.Address, pagination PageRequest) (struct {
	Response     []UnbondingDelegation
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.DelegatorUnbondingDelegations(&_IStaking.CallOpts, delegatorAddr, pagination)
}

// DelegatorValidator is a free data retrieval call binding the contract method 0x77007d17.
//
// Solidity: function delegatorValidator(address delegatorAddr, address validatorAddr) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string) validator)
func (_IStaking *IStakingCaller) DelegatorValidator(opts *bind.CallOpts, delegatorAddr common.Address, validatorAddr common.Address) (Validator, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "delegatorValidator", delegatorAddr, validatorAddr)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// DelegatorValidator is a free data retrieval call binding the contract method 0x77007d17.
//
// Solidity: function delegatorValidator(address delegatorAddr, address validatorAddr) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string) validator)
func (_IStaking *IStakingSession) DelegatorValidator(delegatorAddr common.Address, validatorAddr common.Address) (Validator, error) {
	return _IStaking.Contract.DelegatorValidator(&_IStaking.CallOpts, delegatorAddr, validatorAddr)
}

// DelegatorValidator is a free data retrieval call binding the contract method 0x77007d17.
//
// Solidity: function delegatorValidator(address delegatorAddr, address validatorAddr) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string) validator)
func (_IStaking *IStakingCallerSession) DelegatorValidator(delegatorAddr common.Address, validatorAddr common.Address) (Validator, error) {
	return _IStaking.Contract.DelegatorValidator(&_IStaking.CallOpts, delegatorAddr, validatorAddr)
}

// DelegatorValidators is a free data retrieval call binding the contract method 0xe8e4fe99.
//
// Solidity: function delegatorValidators(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[] validators, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCaller) DelegatorValidators(opts *bind.CallOpts, delegatorAddr common.Address, pagination PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "delegatorValidators", delegatorAddr, pagination)

	outstruct := new(struct {
		Validators   []Validator
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// DelegatorValidators is a free data retrieval call binding the contract method 0xe8e4fe99.
//
// Solidity: function delegatorValidators(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[] validators, (bytes,uint64) pageResponse)
func (_IStaking *IStakingSession) DelegatorValidators(delegatorAddr common.Address, pagination PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.DelegatorValidators(&_IStaking.CallOpts, delegatorAddr, pagination)
}

// DelegatorValidators is a free data retrieval call binding the contract method 0xe8e4fe99.
//
// Solidity: function delegatorValidators(address delegatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[] validators, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCallerSession) DelegatorValidators(delegatorAddr common.Address, pagination PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.DelegatorValidators(&_IStaking.CallOpts, delegatorAddr, pagination)
}

// HistoricalInfo is a free data retrieval call binding the contract method 0xc5e6e7d4.
//
// Solidity: function historicalInfo(int64 height) view returns((((uint64,uint64),string,int64,int64,(string,(uint32,string)),string,string,string,string,string,string,string,string,string),(address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[]) historicalInfo)
func (_IStaking *IStakingCaller) HistoricalInfo(opts *bind.CallOpts, height int64) (HistoricalInfo, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "historicalInfo", height)

	if err != nil {
		return *new(HistoricalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(HistoricalInfo)).(*HistoricalInfo)

	return out0, err

}

// HistoricalInfo is a free data retrieval call binding the contract method 0xc5e6e7d4.
//
// Solidity: function historicalInfo(int64 height) view returns((((uint64,uint64),string,int64,int64,(string,(uint32,string)),string,string,string,string,string,string,string,string,string),(address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[]) historicalInfo)
func (_IStaking *IStakingSession) HistoricalInfo(height int64) (HistoricalInfo, error) {
	return _IStaking.Contract.HistoricalInfo(&_IStaking.CallOpts, height)
}

// HistoricalInfo is a free data retrieval call binding the contract method 0xc5e6e7d4.
//
// Solidity: function historicalInfo(int64 height) view returns((((uint64,uint64),string,int64,int64,(string,(uint32,string)),string,string,string,string,string,string,string,string,string),(address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[]) historicalInfo)
func (_IStaking *IStakingCallerSession) HistoricalInfo(height int64) (HistoricalInfo, error) {
	return _IStaking.Contract.HistoricalInfo(&_IStaking.CallOpts, height)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((int64,uint32,uint32,uint32,string,uint256) params)
func (_IStaking *IStakingCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((int64,uint32,uint32,uint32,string,uint256) params)
func (_IStaking *IStakingSession) Params() (Params, error) {
	return _IStaking.Contract.Params(&_IStaking.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((int64,uint32,uint32,uint32,string,uint256) params)
func (_IStaking *IStakingCallerSession) Params() (Params, error) {
	return _IStaking.Contract.Params(&_IStaking.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns((uint256,uint256) pool)
func (_IStaking *IStakingCaller) Pool(opts *bind.CallOpts) (Pool, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(Pool), err
	}

	out0 := *abi.ConvertType(out[0], new(Pool)).(*Pool)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns((uint256,uint256) pool)
func (_IStaking *IStakingSession) Pool() (Pool, error) {
	return _IStaking.Contract.Pool(&_IStaking.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns((uint256,uint256) pool)
func (_IStaking *IStakingCallerSession) Pool() (Pool, error) {
	return _IStaking.Contract.Pool(&_IStaking.CallOpts)
}

// Redelegations is a free data retrieval call binding the contract method 0xeb5643a9.
//
// Solidity: function redelegations(address delegatorAddr, address srcValidatorAddr, address dstValidatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,address,(int64,int64,uint256,uint256)[]),((int64,int64,uint256,uint256),uint256)[])[] redelegationResponses, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCaller) Redelegations(opts *bind.CallOpts, delegatorAddr common.Address, srcValidatorAddr common.Address, dstValidatorAddr common.Address, pagination PageRequest) (struct {
	RedelegationResponses []RedelegationResponse
	PageResponse          PageResponse
}, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "redelegations", delegatorAddr, srcValidatorAddr, dstValidatorAddr, pagination)

	outstruct := new(struct {
		RedelegationResponses []RedelegationResponse
		PageResponse          PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RedelegationResponses = *abi.ConvertType(out[0], new([]RedelegationResponse)).(*[]RedelegationResponse)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Redelegations is a free data retrieval call binding the contract method 0xeb5643a9.
//
// Solidity: function redelegations(address delegatorAddr, address srcValidatorAddr, address dstValidatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,address,(int64,int64,uint256,uint256)[]),((int64,int64,uint256,uint256),uint256)[])[] redelegationResponses, (bytes,uint64) pageResponse)
func (_IStaking *IStakingSession) Redelegations(delegatorAddr common.Address, srcValidatorAddr common.Address, dstValidatorAddr common.Address, pagination PageRequest) (struct {
	RedelegationResponses []RedelegationResponse
	PageResponse          PageResponse
}, error) {
	return _IStaking.Contract.Redelegations(&_IStaking.CallOpts, delegatorAddr, srcValidatorAddr, dstValidatorAddr, pagination)
}

// Redelegations is a free data retrieval call binding the contract method 0xeb5643a9.
//
// Solidity: function redelegations(address delegatorAddr, address srcValidatorAddr, address dstValidatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,address,(int64,int64,uint256,uint256)[]),((int64,int64,uint256,uint256),uint256)[])[] redelegationResponses, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCallerSession) Redelegations(delegatorAddr common.Address, srcValidatorAddr common.Address, dstValidatorAddr common.Address, pagination PageRequest) (struct {
	RedelegationResponses []RedelegationResponse
	PageResponse          PageResponse
}, error) {
	return _IStaking.Contract.Redelegations(&_IStaking.CallOpts, delegatorAddr, srcValidatorAddr, dstValidatorAddr, pagination)
}

// UnbondingDelegation is a free data retrieval call binding the contract method 0x97e41907.
//
// Solidity: function unbondingDelegation(address delegatorAddr, address validatorAddr) view returns((address,address,(int64,int64,uint256,uint256)[]) response)
func (_IStaking *IStakingCaller) UnbondingDelegation(opts *bind.CallOpts, delegatorAddr common.Address, validatorAddr common.Address) (UnbondingDelegation, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "unbondingDelegation", delegatorAddr, validatorAddr)

	if err != nil {
		return *new(UnbondingDelegation), err
	}

	out0 := *abi.ConvertType(out[0], new(UnbondingDelegation)).(*UnbondingDelegation)

	return out0, err

}

// UnbondingDelegation is a free data retrieval call binding the contract method 0x97e41907.
//
// Solidity: function unbondingDelegation(address delegatorAddr, address validatorAddr) view returns((address,address,(int64,int64,uint256,uint256)[]) response)
func (_IStaking *IStakingSession) UnbondingDelegation(delegatorAddr common.Address, validatorAddr common.Address) (UnbondingDelegation, error) {
	return _IStaking.Contract.UnbondingDelegation(&_IStaking.CallOpts, delegatorAddr, validatorAddr)
}

// UnbondingDelegation is a free data retrieval call binding the contract method 0x97e41907.
//
// Solidity: function unbondingDelegation(address delegatorAddr, address validatorAddr) view returns((address,address,(int64,int64,uint256,uint256)[]) response)
func (_IStaking *IStakingCallerSession) UnbondingDelegation(delegatorAddr common.Address, validatorAddr common.Address) (UnbondingDelegation, error) {
	return _IStaking.Contract.UnbondingDelegation(&_IStaking.CallOpts, delegatorAddr, validatorAddr)
}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address validatorAddr) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string) validator)
func (_IStaking *IStakingCaller) Validator(opts *bind.CallOpts, validatorAddr common.Address) (Validator, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "validator", validatorAddr)

	if err != nil {
		return *new(Validator), err
	}

	out0 := *abi.ConvertType(out[0], new(Validator)).(*Validator)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address validatorAddr) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string) validator)
func (_IStaking *IStakingSession) Validator(validatorAddr common.Address) (Validator, error) {
	return _IStaking.Contract.Validator(&_IStaking.CallOpts, validatorAddr)
}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address validatorAddr) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string) validator)
func (_IStaking *IStakingCallerSession) Validator(validatorAddr common.Address) (Validator, error) {
	return _IStaking.Contract.Validator(&_IStaking.CallOpts, validatorAddr)
}

// ValidatorDelegations is a free data retrieval call binding the contract method 0xc89017cd.
//
// Solidity: function validatorDelegations(address validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,(uint256,uint8)),(string,uint256))[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCaller) ValidatorDelegations(opts *bind.CallOpts, validatorAddr common.Address, pagination PageRequest) (struct {
	Response     []DelegationResponse
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "validatorDelegations", validatorAddr, pagination)

	outstruct := new(struct {
		Response     []DelegationResponse
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Response = *abi.ConvertType(out[0], new([]DelegationResponse)).(*[]DelegationResponse)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// ValidatorDelegations is a free data retrieval call binding the contract method 0xc89017cd.
//
// Solidity: function validatorDelegations(address validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,(uint256,uint8)),(string,uint256))[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingSession) ValidatorDelegations(validatorAddr common.Address, pagination PageRequest) (struct {
	Response     []DelegationResponse
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.ValidatorDelegations(&_IStaking.CallOpts, validatorAddr, pagination)
}

// ValidatorDelegations is a free data retrieval call binding the contract method 0xc89017cd.
//
// Solidity: function validatorDelegations(address validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns(((address,address,(uint256,uint8)),(string,uint256))[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCallerSession) ValidatorDelegations(validatorAddr common.Address, pagination PageRequest) (struct {
	Response     []DelegationResponse
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.ValidatorDelegations(&_IStaking.CallOpts, validatorAddr, pagination)
}

// ValidatorUnbondingDelegations is a free data retrieval call binding the contract method 0x51b55195.
//
// Solidity: function validatorUnbondingDelegations(address validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,int64,uint256,uint256)[])[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCaller) ValidatorUnbondingDelegations(opts *bind.CallOpts, validatorAddr common.Address, pagination PageRequest) (struct {
	Response     []UnbondingDelegation
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "validatorUnbondingDelegations", validatorAddr, pagination)

	outstruct := new(struct {
		Response     []UnbondingDelegation
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Response = *abi.ConvertType(out[0], new([]UnbondingDelegation)).(*[]UnbondingDelegation)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// ValidatorUnbondingDelegations is a free data retrieval call binding the contract method 0x51b55195.
//
// Solidity: function validatorUnbondingDelegations(address validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,int64,uint256,uint256)[])[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingSession) ValidatorUnbondingDelegations(validatorAddr common.Address, pagination PageRequest) (struct {
	Response     []UnbondingDelegation
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.ValidatorUnbondingDelegations(&_IStaking.CallOpts, validatorAddr, pagination)
}

// ValidatorUnbondingDelegations is a free data retrieval call binding the contract method 0x51b55195.
//
// Solidity: function validatorUnbondingDelegations(address validatorAddr, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,(int64,int64,uint256,uint256)[])[] response, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCallerSession) ValidatorUnbondingDelegations(validatorAddr common.Address, pagination PageRequest) (struct {
	Response     []UnbondingDelegation
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.ValidatorUnbondingDelegations(&_IStaking.CallOpts, validatorAddr, pagination)
}

// Validators is a free data retrieval call binding the contract method 0x9e8505bb.
//
// Solidity: function validators(uint8 status, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[] validators, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCaller) Validators(opts *bind.CallOpts, status uint8, pagination PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IStaking.contract.Call(opts, &out, "validators", status, pagination)

	outstruct := new(struct {
		Validators   []Validator
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]Validator)).(*[]Validator)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0x9e8505bb.
//
// Solidity: function validators(uint8 status, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[] validators, (bytes,uint64) pageResponse)
func (_IStaking *IStakingSession) Validators(status uint8, pagination PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.Validators(&_IStaking.CallOpts, status, pagination)
}

// Validators is a free data retrieval call binding the contract method 0x9e8505bb.
//
// Solidity: function validators(uint8 status, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,string,bool,uint8,uint256,uint256,(string,string,string,string,string),int64,int64,((uint256,uint256,uint256),int64),uint256,int64,uint64[],string,string,string,string)[] validators, (bytes,uint64) pageResponse)
func (_IStaking *IStakingCallerSession) Validators(status uint8, pagination PageRequest) (struct {
	Validators   []Validator
	PageResponse PageResponse
}, error) {
	return _IStaking.Contract.Validators(&_IStaking.CallOpts, status, pagination)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x50826aef.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, uint256 creationHeight) returns(bool success)
func (_IStaking *IStakingTransactor) CancelUnbondingDelegation(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "cancelUnbondingDelegation", validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x50826aef.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, uint256 creationHeight) returns(bool success)
func (_IStaking *IStakingSession) CancelUnbondingDelegation(validatorAddress common.Address, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.CancelUnbondingDelegation(&_IStaking.TransactOpts, validatorAddress, amount, creationHeight)
}

// CancelUnbondingDelegation is a paid mutator transaction binding the contract method 0x50826aef.
//
// Solidity: function cancelUnbondingDelegation(address validatorAddress, uint256 amount, uint256 creationHeight) returns(bool success)
func (_IStaking *IStakingTransactorSession) CancelUnbondingDelegation(validatorAddress common.Address, amount *big.Int, creationHeight *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.CancelUnbondingDelegation(&_IStaking.TransactOpts, validatorAddress, amount, creationHeight)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) returns(bool success)
func (_IStaking *IStakingTransactor) Delegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "delegate", validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) returns(bool success)
func (_IStaking *IStakingSession) Delegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.Delegate(&_IStaking.TransactOpts, validatorAddress, amount)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address validatorAddress, uint256 amount) returns(bool success)
func (_IStaking *IStakingTransactorSession) Delegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.Delegate(&_IStaking.TransactOpts, validatorAddress, amount)
}

// EditValidator is a paid mutator transaction binding the contract method 0xd54b30d1.
//
// Solidity: function editValidator((string,string,string,string,string) description, int256 commissionRate, int256 minSelfDelegation, address relayerAddress, address challengerAddress, string blsKey, string blsProof) returns(bool success)
func (_IStaking *IStakingTransactor) EditValidator(opts *bind.TransactOpts, description Description, commissionRate *big.Int, minSelfDelegation *big.Int, relayerAddress common.Address, challengerAddress common.Address, blsKey string, blsProof string) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "editValidator", description, commissionRate, minSelfDelegation, relayerAddress, challengerAddress, blsKey, blsProof)
}

// EditValidator is a paid mutator transaction binding the contract method 0xd54b30d1.
//
// Solidity: function editValidator((string,string,string,string,string) description, int256 commissionRate, int256 minSelfDelegation, address relayerAddress, address challengerAddress, string blsKey, string blsProof) returns(bool success)
func (_IStaking *IStakingSession) EditValidator(description Description, commissionRate *big.Int, minSelfDelegation *big.Int, relayerAddress common.Address, challengerAddress common.Address, blsKey string, blsProof string) (*types.Transaction, error) {
	return _IStaking.Contract.EditValidator(&_IStaking.TransactOpts, description, commissionRate, minSelfDelegation, relayerAddress, challengerAddress, blsKey, blsProof)
}

// EditValidator is a paid mutator transaction binding the contract method 0xd54b30d1.
//
// Solidity: function editValidator((string,string,string,string,string) description, int256 commissionRate, int256 minSelfDelegation, address relayerAddress, address challengerAddress, string blsKey, string blsProof) returns(bool success)
func (_IStaking *IStakingTransactorSession) EditValidator(description Description, commissionRate *big.Int, minSelfDelegation *big.Int, relayerAddress common.Address, challengerAddress common.Address, blsKey string, blsProof string) (*types.Transaction, error) {
	return _IStaking.Contract.EditValidator(&_IStaking.TransactOpts, description, commissionRate, minSelfDelegation, relayerAddress, challengerAddress, blsKey, blsProof)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrcAddress, address validatorDstAddress, uint256 amount) returns(uint256 completionTime)
func (_IStaking *IStakingTransactor) Redelegate(opts *bind.TransactOpts, validatorSrcAddress common.Address, validatorDstAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "redelegate", validatorSrcAddress, validatorDstAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrcAddress, address validatorDstAddress, uint256 amount) returns(uint256 completionTime)
func (_IStaking *IStakingSession) Redelegate(validatorSrcAddress common.Address, validatorDstAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.Redelegate(&_IStaking.TransactOpts, validatorSrcAddress, validatorDstAddress, amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address validatorSrcAddress, address validatorDstAddress, uint256 amount) returns(uint256 completionTime)
func (_IStaking *IStakingTransactorSession) Redelegate(validatorSrcAddress common.Address, validatorDstAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.Redelegate(&_IStaking.TransactOpts, validatorSrcAddress, validatorDstAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns(uint256 completionTime)
func (_IStaking *IStakingTransactor) Undelegate(opts *bind.TransactOpts, validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.contract.Transact(opts, "undelegate", validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns(uint256 completionTime)
func (_IStaking *IStakingSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.Undelegate(&_IStaking.TransactOpts, validatorAddress, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validatorAddress, uint256 amount) returns(uint256 completionTime)
func (_IStaking *IStakingTransactorSession) Undelegate(validatorAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStaking.Contract.Undelegate(&_IStaking.TransactOpts, validatorAddress, amount)
}

// IStakingCancelUnbondingDelegationIterator is returned from FilterCancelUnbondingDelegation and is used to iterate over the raw logs and unpacked data for CancelUnbondingDelegation events raised by the IStaking contract.
type IStakingCancelUnbondingDelegationIterator struct {
	Event *IStakingCancelUnbondingDelegation // Event containing the contract specifics and raw log

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
func (it *IStakingCancelUnbondingDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingCancelUnbondingDelegation)
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
		it.Event = new(IStakingCancelUnbondingDelegation)
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
func (it *IStakingCancelUnbondingDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingCancelUnbondingDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingCancelUnbondingDelegation represents a CancelUnbondingDelegation event raised by the IStaking contract.
type IStakingCancelUnbondingDelegation struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Amount           *big.Int
	CreationHeight   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCancelUnbondingDelegation is a free log retrieval operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_IStaking *IStakingFilterer) FilterCancelUnbondingDelegation(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorAddress []common.Address) (*IStakingCancelUnbondingDelegationIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IStaking.contract.FilterLogs(opts, "CancelUnbondingDelegation", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingCancelUnbondingDelegationIterator{contract: _IStaking.contract, event: "CancelUnbondingDelegation", logs: logs, sub: sub}, nil
}

// WatchCancelUnbondingDelegation is a free log subscription operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_IStaking *IStakingFilterer) WatchCancelUnbondingDelegation(opts *bind.WatchOpts, sink chan<- *IStakingCancelUnbondingDelegation, delegatorAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IStaking.contract.WatchLogs(opts, "CancelUnbondingDelegation", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingCancelUnbondingDelegation)
				if err := _IStaking.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
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

// ParseCancelUnbondingDelegation is a log parse operation binding the contract event 0x6dbe2fb6b2613bdd8e3d284a6111592e06c3ab0af846ff89b6688d48f408dbb5.
//
// Solidity: event CancelUnbondingDelegation(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 creationHeight)
func (_IStaking *IStakingFilterer) ParseCancelUnbondingDelegation(log types.Log) (*IStakingCancelUnbondingDelegation, error) {
	event := new(IStakingCancelUnbondingDelegation)
	if err := _IStaking.contract.UnpackLog(event, "CancelUnbondingDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the IStaking contract.
type IStakingDelegateIterator struct {
	Event *IStakingDelegate // Event containing the contract specifics and raw log

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
func (it *IStakingDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingDelegate)
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
		it.Event = new(IStakingDelegate)
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
func (it *IStakingDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingDelegate represents a Delegate event raised by the IStaking contract.
type IStakingDelegate struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_IStaking *IStakingFilterer) FilterDelegate(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*IStakingDelegateIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IStaking.contract.FilterLogs(opts, "Delegate", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &IStakingDelegateIterator{contract: _IStaking.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_IStaking *IStakingFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *IStakingDelegate, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IStaking.contract.WatchLogs(opts, "Delegate", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingDelegate)
				if err := _IStaking.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_IStaking *IStakingFilterer) ParseDelegate(log types.Log) (*IStakingDelegate, error) {
	event := new(IStakingDelegate)
	if err := _IStaking.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingEditValidatorIterator is returned from FilterEditValidator and is used to iterate over the raw logs and unpacked data for EditValidator events raised by the IStaking contract.
type IStakingEditValidatorIterator struct {
	Event *IStakingEditValidator // Event containing the contract specifics and raw log

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
func (it *IStakingEditValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingEditValidator)
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
		it.Event = new(IStakingEditValidator)
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
func (it *IStakingEditValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingEditValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingEditValidator represents a EditValidator event raised by the IStaking contract.
type IStakingEditValidator struct {
	Validator         common.Address
	CommissionRate    *big.Int
	MinSelfDelegation *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEditValidator is a free log retrieval operation binding the contract event 0xdce27cf2792bd8d8f28df5d2cdf379cd593414f21332370ca808c1e703eb4e1f.
//
// Solidity: event EditValidator(address indexed validator, int256 commissionRate, int256 minSelfDelegation)
func (_IStaking *IStakingFilterer) FilterEditValidator(opts *bind.FilterOpts, validator []common.Address) (*IStakingEditValidatorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IStaking.contract.FilterLogs(opts, "EditValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return &IStakingEditValidatorIterator{contract: _IStaking.contract, event: "EditValidator", logs: logs, sub: sub}, nil
}

// WatchEditValidator is a free log subscription operation binding the contract event 0xdce27cf2792bd8d8f28df5d2cdf379cd593414f21332370ca808c1e703eb4e1f.
//
// Solidity: event EditValidator(address indexed validator, int256 commissionRate, int256 minSelfDelegation)
func (_IStaking *IStakingFilterer) WatchEditValidator(opts *bind.WatchOpts, sink chan<- *IStakingEditValidator, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _IStaking.contract.WatchLogs(opts, "EditValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingEditValidator)
				if err := _IStaking.contract.UnpackLog(event, "EditValidator", log); err != nil {
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

// ParseEditValidator is a log parse operation binding the contract event 0xdce27cf2792bd8d8f28df5d2cdf379cd593414f21332370ca808c1e703eb4e1f.
//
// Solidity: event EditValidator(address indexed validator, int256 commissionRate, int256 minSelfDelegation)
func (_IStaking *IStakingFilterer) ParseEditValidator(log types.Log) (*IStakingEditValidator, error) {
	event := new(IStakingEditValidator)
	if err := _IStaking.contract.UnpackLog(event, "EditValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingRedelegateIterator is returned from FilterRedelegate and is used to iterate over the raw logs and unpacked data for Redelegate events raised by the IStaking contract.
type IStakingRedelegateIterator struct {
	Event *IStakingRedelegate // Event containing the contract specifics and raw log

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
func (it *IStakingRedelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingRedelegate)
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
		it.Event = new(IStakingRedelegate)
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
func (it *IStakingRedelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingRedelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingRedelegate represents a Redelegate event raised by the IStaking contract.
type IStakingRedelegate struct {
	DelegatorAddress    common.Address
	ValidatorSrcAddress common.Address
	ValidatorDstAddress common.Address
	Amount              *big.Int
	CompletionTime      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedelegate is a free log retrieval operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_IStaking *IStakingFilterer) FilterRedelegate(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorSrcAddress []common.Address, validatorDstAddress []common.Address) (*IStakingRedelegateIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorSrcAddressRule []interface{}
	for _, validatorSrcAddressItem := range validatorSrcAddress {
		validatorSrcAddressRule = append(validatorSrcAddressRule, validatorSrcAddressItem)
	}
	var validatorDstAddressRule []interface{}
	for _, validatorDstAddressItem := range validatorDstAddress {
		validatorDstAddressRule = append(validatorDstAddressRule, validatorDstAddressItem)
	}

	logs, sub, err := _IStaking.contract.FilterLogs(opts, "Redelegate", delegatorAddressRule, validatorSrcAddressRule, validatorDstAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingRedelegateIterator{contract: _IStaking.contract, event: "Redelegate", logs: logs, sub: sub}, nil
}

// WatchRedelegate is a free log subscription operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_IStaking *IStakingFilterer) WatchRedelegate(opts *bind.WatchOpts, sink chan<- *IStakingRedelegate, delegatorAddress []common.Address, validatorSrcAddress []common.Address, validatorDstAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorSrcAddressRule []interface{}
	for _, validatorSrcAddressItem := range validatorSrcAddress {
		validatorSrcAddressRule = append(validatorSrcAddressRule, validatorSrcAddressItem)
	}
	var validatorDstAddressRule []interface{}
	for _, validatorDstAddressItem := range validatorDstAddress {
		validatorDstAddressRule = append(validatorDstAddressRule, validatorDstAddressItem)
	}

	logs, sub, err := _IStaking.contract.WatchLogs(opts, "Redelegate", delegatorAddressRule, validatorSrcAddressRule, validatorDstAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingRedelegate)
				if err := _IStaking.contract.UnpackLog(event, "Redelegate", log); err != nil {
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

// ParseRedelegate is a log parse operation binding the contract event 0x82b07f2421474f1e3f1e0b34738cb5ffb925273f408e7591d9c803dcae8da657.
//
// Solidity: event Redelegate(address indexed delegatorAddress, address indexed validatorSrcAddress, address indexed validatorDstAddress, uint256 amount, uint256 completionTime)
func (_IStaking *IStakingFilterer) ParseRedelegate(log types.Log) (*IStakingRedelegate, error) {
	event := new(IStakingRedelegate)
	if err := _IStaking.contract.UnpackLog(event, "Redelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStakingUndelegateIterator is returned from FilterUndelegate and is used to iterate over the raw logs and unpacked data for Undelegate events raised by the IStaking contract.
type IStakingUndelegateIterator struct {
	Event *IStakingUndelegate // Event containing the contract specifics and raw log

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
func (it *IStakingUndelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStakingUndelegate)
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
		it.Event = new(IStakingUndelegate)
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
func (it *IStakingUndelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStakingUndelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStakingUndelegate represents a Undelegate event raised by the IStaking contract.
type IStakingUndelegate struct {
	DelegatorAddress common.Address
	ValidatorAddress common.Address
	Amount           *big.Int
	CompletionTime   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUndelegate is a free log retrieval operation binding the contract event 0x54feacf7d18f42300ff65ba8829f17724df4915a8b1984c5fee2cc42d4027f07.
//
// Solidity: event Undelegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_IStaking *IStakingFilterer) FilterUndelegate(opts *bind.FilterOpts, delegatorAddress []common.Address, validatorAddress []common.Address) (*IStakingUndelegateIterator, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IStaking.contract.FilterLogs(opts, "Undelegate", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &IStakingUndelegateIterator{contract: _IStaking.contract, event: "Undelegate", logs: logs, sub: sub}, nil
}

// WatchUndelegate is a free log subscription operation binding the contract event 0x54feacf7d18f42300ff65ba8829f17724df4915a8b1984c5fee2cc42d4027f07.
//
// Solidity: event Undelegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_IStaking *IStakingFilterer) WatchUndelegate(opts *bind.WatchOpts, sink chan<- *IStakingUndelegate, delegatorAddress []common.Address, validatorAddress []common.Address) (event.Subscription, error) {

	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}
	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IStaking.contract.WatchLogs(opts, "Undelegate", delegatorAddressRule, validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStakingUndelegate)
				if err := _IStaking.contract.UnpackLog(event, "Undelegate", log); err != nil {
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

// ParseUndelegate is a log parse operation binding the contract event 0x54feacf7d18f42300ff65ba8829f17724df4915a8b1984c5fee2cc42d4027f07.
//
// Solidity: event Undelegate(address indexed delegatorAddress, address indexed validatorAddress, uint256 amount, uint256 completionTime)
func (_IStaking *IStakingFilterer) ParseUndelegate(log types.Log) (*IStakingUndelegate, error) {
	event := new(IStakingUndelegate)
	if err := _IStaking.contract.UnpackLog(event, "Undelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
