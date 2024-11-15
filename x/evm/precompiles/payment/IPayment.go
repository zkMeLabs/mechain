// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payment

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

// AutoSettleRecord is an auto generated low-level Go binding around an user-defined struct.
type AutoSettleRecord struct {
	Timestamp int64
	Addr      string
}

// DelayedWithdrawalRecord is an auto generated low-level Go binding around an user-defined struct.
type DelayedWithdrawalRecord struct {
	Addr            string
	Amount          *big.Int
	From            string
	UnlockTimestamp int64
}

// DynamicBalance is an auto generated low-level Go binding around an user-defined struct.
type DynamicBalance struct {
	DynamicBalance   *big.Int
	StreamRecord     StreamRecord
	CurrentTimestamp int64
	BankBalance      *big.Int
	AvailableBalance *big.Int
	LockedFee        *big.Int
	ChangeRate       *big.Int
}

// OutFlow is an auto generated low-level Go binding around an user-defined struct.
type OutFlow struct {
	ToAddress string
	Rate      *big.Int
	Status    int32
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
	VersionedParams           VersionedParams
	PaymentAccountCountLimit  uint64
	ForcedSettleTime          uint64
	MaxAutoSettleFlowCount    uint64
	MaxAutoResumeFlowCount    uint64
	FeeDenom                  string
	WithdrawTimeLockThreshold *big.Int
	WithdrawTimeLockDuration  uint64
}

// PaymentAccount is an auto generated low-level Go binding around an user-defined struct.
type PaymentAccount struct {
	Addr       string
	Owner      string
	Refundable bool
}

// PaymentAccountCount is an auto generated low-level Go binding around an user-defined struct.
type PaymentAccountCount struct {
	Owner string
	Count uint64
}

// StreamRecord is an auto generated low-level Go binding around an user-defined struct.
type StreamRecord struct {
	Account           string
	CrudTimestamp     int64
	NetflowRate       *big.Int
	StaticBalance     *big.Int
	BufferBalance     *big.Int
	LockBalance       *big.Int
	Status            int32
	SettleTimestamp   int64
	OutFlowCount      uint64
	FrozenNetflowRate *big.Int
}

// VersionedParams is an auto generated low-level Go binding around an user-defined struct.
type VersionedParams struct {
	ReserveTime      uint64
	ValidatorTaxRate *big.Int
}

// IPaymentMetaData contains all meta data concerning the IPayment contract.
var IPaymentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreatePaymentAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DisableRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"UpdateParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"autoSettleRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"timestamp\",\"type\":\"int64\"},{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"internalType\":\"structAutoSettleRecord[]\",\"name\":\"autoSettleRecords\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createPaymentAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"}],\"name\":\"delayedWithdrawal\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"unlockTimestamp\",\"type\":\"int64\"}],\"internalType\":\"structDelayedWithdrawalRecord\",\"name\":\"delayedWithdrawal\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"disableRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"}],\"name\":\"dynamicBalance\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dynamicBalance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"crudTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"netflowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockBalance\",\"type\":\"uint256\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"internalType\":\"int64\",\"name\":\"settleTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"outFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"frozenNetflowRate\",\"type\":\"uint256\"}],\"internalType\":\"structStreamRecord\",\"name\":\"streamRecord\",\"type\":\"tuple\"},{\"internalType\":\"int64\",\"name\":\"currentTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"bankBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"changeRate\",\"type\":\"uint256\"}],\"internalType\":\"structDynamicBalance\",\"name\":\"dynamicBalance\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"}],\"name\":\"outFlows\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"toAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"}],\"internalType\":\"structOutFlow[]\",\"name\":\"outFlows\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"reserveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"validatorTaxRate\",\"type\":\"uint256\"}],\"internalType\":\"structVersionedParams\",\"name\":\"versionedParams\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"paymentAccountCountLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"forcedSettleTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxAutoSettleFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxAutoResumeFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"feeDenom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"withdrawTimeLockThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"withdrawTimeLockDuration\",\"type\":\"uint64\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"timestamp\",\"type\":\"int64\"}],\"name\":\"paramsByTimestamp\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"reserveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"validatorTaxRate\",\"type\":\"uint256\"}],\"internalType\":\"structVersionedParams\",\"name\":\"versionedParams\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"paymentAccountCountLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"forcedSettleTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxAutoSettleFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxAutoResumeFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"feeDenom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"withdrawTimeLockThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"withdrawTimeLockDuration\",\"type\":\"uint64\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"paymentAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"refundable\",\"type\":\"bool\"}],\"internalType\":\"structPaymentAccount\",\"name\":\"paymentAccount\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"}],\"name\":\"paymentAccountCount\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"internalType\":\"structPaymentAccountCount\",\"name\":\"paymentAccountCount\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"paymentAccountCounts\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"internalType\":\"structPaymentAccountCount[]\",\"name\":\"paymentAccountCounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"paymentAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"refundable\",\"type\":\"bool\"}],\"internalType\":\"structPaymentAccount[]\",\"name\":\"paymentAccounts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"}],\"name\":\"paymentAccountsByOwner\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"accounts\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"}],\"name\":\"streamRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"crudTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"netflowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockBalance\",\"type\":\"uint256\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"internalType\":\"int64\",\"name\":\"settleTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"outFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"frozenNetflowRate\",\"type\":\"uint256\"}],\"internalType\":\"structStreamRecord\",\"name\":\"streamRecord\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"streamRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"account\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"crudTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint256\",\"name\":\"netflowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"staticBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockBalance\",\"type\":\"uint256\"},{\"internalType\":\"int32\",\"name\":\"status\",\"type\":\"int32\"},{\"internalType\":\"int64\",\"name\":\"settleTimestamp\",\"type\":\"int64\"},{\"internalType\":\"uint64\",\"name\":\"outFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"frozenNetflowRate\",\"type\":\"uint256\"}],\"internalType\":\"structStreamRecord[]\",\"name\":\"streamRecords\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authority\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"reserveTime\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"validatorTaxRate\",\"type\":\"uint256\"}],\"internalType\":\"structVersionedParams\",\"name\":\"versionedParams\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"paymentAccountCountLimit\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"forcedSettleTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxAutoSettleFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxAutoResumeFlowCount\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"feeDenom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"withdrawTimeLockThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"withdrawTimeLockDuration\",\"type\":\"uint64\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"updateParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"from\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use IPaymentMetaData.ABI instead.
var IPaymentABI = IPaymentMetaData.ABI

// IPayment is an auto generated Go binding around an Ethereum contract.
type IPayment struct {
	IPaymentCaller     // Read-only binding to the contract
	IPaymentTransactor // Write-only binding to the contract
	IPaymentFilterer   // Log filterer for contract events
}

// IPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPaymentSession struct {
	Contract     *IPayment         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPaymentCallerSession struct {
	Contract *IPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPaymentTransactorSession struct {
	Contract     *IPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPaymentRaw struct {
	Contract *IPayment // Generic contract binding to access the raw methods on
}

// IPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPaymentCallerRaw struct {
	Contract *IPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// IPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPaymentTransactorRaw struct {
	Contract *IPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPayment creates a new instance of IPayment, bound to a specific deployed contract.
func NewIPayment(address common.Address, backend bind.ContractBackend) (*IPayment, error) {
	contract, err := bindIPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPayment{IPaymentCaller: IPaymentCaller{contract: contract}, IPaymentTransactor: IPaymentTransactor{contract: contract}, IPaymentFilterer: IPaymentFilterer{contract: contract}}, nil
}

// NewIPaymentCaller creates a new read-only instance of IPayment, bound to a specific deployed contract.
func NewIPaymentCaller(address common.Address, caller bind.ContractCaller) (*IPaymentCaller, error) {
	contract, err := bindIPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymentCaller{contract: contract}, nil
}

// NewIPaymentTransactor creates a new write-only instance of IPayment, bound to a specific deployed contract.
func NewIPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*IPaymentTransactor, error) {
	contract, err := bindIPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPaymentTransactor{contract: contract}, nil
}

// NewIPaymentFilterer creates a new log filterer instance of IPayment, bound to a specific deployed contract.
func NewIPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*IPaymentFilterer, error) {
	contract, err := bindIPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPaymentFilterer{contract: contract}, nil
}

// bindIPayment binds a generic wrapper to an already deployed contract.
func bindIPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPaymentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPayment *IPaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPayment.Contract.IPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPayment *IPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPayment.Contract.IPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPayment *IPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPayment.Contract.IPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPayment *IPaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPayment *IPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPayment *IPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPayment.Contract.contract.Transact(opts, method, params...)
}

// AutoSettleRecords is a free data retrieval call binding the contract method 0x2e43a1dc.
//
// Solidity: function autoSettleRecords((bytes,uint64,uint64,bool,bool) pagination) view returns((int64,string)[] autoSettleRecords, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentCaller) AutoSettleRecords(opts *bind.CallOpts, pagination PageRequest) (struct {
	AutoSettleRecords []AutoSettleRecord
	PageResponse      PageResponse
}, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "autoSettleRecords", pagination)

	outstruct := new(struct {
		AutoSettleRecords []AutoSettleRecord
		PageResponse      PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AutoSettleRecords = *abi.ConvertType(out[0], new([]AutoSettleRecord)).(*[]AutoSettleRecord)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// AutoSettleRecords is a free data retrieval call binding the contract method 0x2e43a1dc.
//
// Solidity: function autoSettleRecords((bytes,uint64,uint64,bool,bool) pagination) view returns((int64,string)[] autoSettleRecords, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentSession) AutoSettleRecords(pagination PageRequest) (struct {
	AutoSettleRecords []AutoSettleRecord
	PageResponse      PageResponse
}, error) {
	return _IPayment.Contract.AutoSettleRecords(&_IPayment.CallOpts, pagination)
}

// AutoSettleRecords is a free data retrieval call binding the contract method 0x2e43a1dc.
//
// Solidity: function autoSettleRecords((bytes,uint64,uint64,bool,bool) pagination) view returns((int64,string)[] autoSettleRecords, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentCallerSession) AutoSettleRecords(pagination PageRequest) (struct {
	AutoSettleRecords []AutoSettleRecord
	PageResponse      PageResponse
}, error) {
	return _IPayment.Contract.AutoSettleRecords(&_IPayment.CallOpts, pagination)
}

// DelayedWithdrawal is a free data retrieval call binding the contract method 0xc23e3f0c.
//
// Solidity: function delayedWithdrawal(string account) view returns((string,uint256,string,int64) delayedWithdrawal)
func (_IPayment *IPaymentCaller) DelayedWithdrawal(opts *bind.CallOpts, account string) (DelayedWithdrawalRecord, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "delayedWithdrawal", account)

	if err != nil {
		return *new(DelayedWithdrawalRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(DelayedWithdrawalRecord)).(*DelayedWithdrawalRecord)

	return out0, err

}

// DelayedWithdrawal is a free data retrieval call binding the contract method 0xc23e3f0c.
//
// Solidity: function delayedWithdrawal(string account) view returns((string,uint256,string,int64) delayedWithdrawal)
func (_IPayment *IPaymentSession) DelayedWithdrawal(account string) (DelayedWithdrawalRecord, error) {
	return _IPayment.Contract.DelayedWithdrawal(&_IPayment.CallOpts, account)
}

// DelayedWithdrawal is a free data retrieval call binding the contract method 0xc23e3f0c.
//
// Solidity: function delayedWithdrawal(string account) view returns((string,uint256,string,int64) delayedWithdrawal)
func (_IPayment *IPaymentCallerSession) DelayedWithdrawal(account string) (DelayedWithdrawalRecord, error) {
	return _IPayment.Contract.DelayedWithdrawal(&_IPayment.CallOpts, account)
}

// DynamicBalance is a free data retrieval call binding the contract method 0xab840b84.
//
// Solidity: function dynamicBalance(string account) view returns((uint256,(string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256),int64,uint256,uint256,uint256,uint256) dynamicBalance)
func (_IPayment *IPaymentCaller) DynamicBalance(opts *bind.CallOpts, account string) (DynamicBalance, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "dynamicBalance", account)

	if err != nil {
		return *new(DynamicBalance), err
	}

	out0 := *abi.ConvertType(out[0], new(DynamicBalance)).(*DynamicBalance)

	return out0, err

}

// DynamicBalance is a free data retrieval call binding the contract method 0xab840b84.
//
// Solidity: function dynamicBalance(string account) view returns((uint256,(string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256),int64,uint256,uint256,uint256,uint256) dynamicBalance)
func (_IPayment *IPaymentSession) DynamicBalance(account string) (DynamicBalance, error) {
	return _IPayment.Contract.DynamicBalance(&_IPayment.CallOpts, account)
}

// DynamicBalance is a free data retrieval call binding the contract method 0xab840b84.
//
// Solidity: function dynamicBalance(string account) view returns((uint256,(string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256),int64,uint256,uint256,uint256,uint256) dynamicBalance)
func (_IPayment *IPaymentCallerSession) DynamicBalance(account string) (DynamicBalance, error) {
	return _IPayment.Contract.DynamicBalance(&_IPayment.CallOpts, account)
}

// OutFlows is a free data retrieval call binding the contract method 0x4222d636.
//
// Solidity: function outFlows(string account) view returns((string,uint256,int32)[] outFlows)
func (_IPayment *IPaymentCaller) OutFlows(opts *bind.CallOpts, account string) ([]OutFlow, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "outFlows", account)

	if err != nil {
		return *new([]OutFlow), err
	}

	out0 := *abi.ConvertType(out[0], new([]OutFlow)).(*[]OutFlow)

	return out0, err

}

// OutFlows is a free data retrieval call binding the contract method 0x4222d636.
//
// Solidity: function outFlows(string account) view returns((string,uint256,int32)[] outFlows)
func (_IPayment *IPaymentSession) OutFlows(account string) ([]OutFlow, error) {
	return _IPayment.Contract.OutFlows(&_IPayment.CallOpts, account)
}

// OutFlows is a free data retrieval call binding the contract method 0x4222d636.
//
// Solidity: function outFlows(string account) view returns((string,uint256,int32)[] outFlows)
func (_IPayment *IPaymentCallerSession) OutFlows(account string) ([]OutFlow, error) {
	return _IPayment.Contract.OutFlows(&_IPayment.CallOpts, account)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params)
func (_IPayment *IPaymentCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params)
func (_IPayment *IPaymentSession) Params() (Params, error) {
	return _IPayment.Contract.Params(&_IPayment.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params)
func (_IPayment *IPaymentCallerSession) Params() (Params, error) {
	return _IPayment.Contract.Params(&_IPayment.CallOpts)
}

// ParamsByTimestamp is a free data retrieval call binding the contract method 0x5593d8b8.
//
// Solidity: function paramsByTimestamp(int64 timestamp) view returns(((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params)
func (_IPayment *IPaymentCaller) ParamsByTimestamp(opts *bind.CallOpts, timestamp int64) (Params, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "paramsByTimestamp", timestamp)

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// ParamsByTimestamp is a free data retrieval call binding the contract method 0x5593d8b8.
//
// Solidity: function paramsByTimestamp(int64 timestamp) view returns(((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params)
func (_IPayment *IPaymentSession) ParamsByTimestamp(timestamp int64) (Params, error) {
	return _IPayment.Contract.ParamsByTimestamp(&_IPayment.CallOpts, timestamp)
}

// ParamsByTimestamp is a free data retrieval call binding the contract method 0x5593d8b8.
//
// Solidity: function paramsByTimestamp(int64 timestamp) view returns(((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params)
func (_IPayment *IPaymentCallerSession) ParamsByTimestamp(timestamp int64) (Params, error) {
	return _IPayment.Contract.ParamsByTimestamp(&_IPayment.CallOpts, timestamp)
}

// PaymentAccount is a free data retrieval call binding the contract method 0xe64d2baf.
//
// Solidity: function paymentAccount(string addr) view returns((string,string,bool) paymentAccount)
func (_IPayment *IPaymentCaller) PaymentAccount(opts *bind.CallOpts, addr string) (PaymentAccount, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "paymentAccount", addr)

	if err != nil {
		return *new(PaymentAccount), err
	}

	out0 := *abi.ConvertType(out[0], new(PaymentAccount)).(*PaymentAccount)

	return out0, err

}

// PaymentAccount is a free data retrieval call binding the contract method 0xe64d2baf.
//
// Solidity: function paymentAccount(string addr) view returns((string,string,bool) paymentAccount)
func (_IPayment *IPaymentSession) PaymentAccount(addr string) (PaymentAccount, error) {
	return _IPayment.Contract.PaymentAccount(&_IPayment.CallOpts, addr)
}

// PaymentAccount is a free data retrieval call binding the contract method 0xe64d2baf.
//
// Solidity: function paymentAccount(string addr) view returns((string,string,bool) paymentAccount)
func (_IPayment *IPaymentCallerSession) PaymentAccount(addr string) (PaymentAccount, error) {
	return _IPayment.Contract.PaymentAccount(&_IPayment.CallOpts, addr)
}

// PaymentAccountCount is a free data retrieval call binding the contract method 0xd7e65eb5.
//
// Solidity: function paymentAccountCount(string owner) view returns((string,uint64) paymentAccountCount)
func (_IPayment *IPaymentCaller) PaymentAccountCount(opts *bind.CallOpts, owner string) (PaymentAccountCount, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "paymentAccountCount", owner)

	if err != nil {
		return *new(PaymentAccountCount), err
	}

	out0 := *abi.ConvertType(out[0], new(PaymentAccountCount)).(*PaymentAccountCount)

	return out0, err

}

// PaymentAccountCount is a free data retrieval call binding the contract method 0xd7e65eb5.
//
// Solidity: function paymentAccountCount(string owner) view returns((string,uint64) paymentAccountCount)
func (_IPayment *IPaymentSession) PaymentAccountCount(owner string) (PaymentAccountCount, error) {
	return _IPayment.Contract.PaymentAccountCount(&_IPayment.CallOpts, owner)
}

// PaymentAccountCount is a free data retrieval call binding the contract method 0xd7e65eb5.
//
// Solidity: function paymentAccountCount(string owner) view returns((string,uint64) paymentAccountCount)
func (_IPayment *IPaymentCallerSession) PaymentAccountCount(owner string) (PaymentAccountCount, error) {
	return _IPayment.Contract.PaymentAccountCount(&_IPayment.CallOpts, owner)
}

// PaymentAccountCounts is a free data retrieval call binding the contract method 0x1a101f68.
//
// Solidity: function paymentAccountCounts((bytes,uint64,uint64,bool,bool) pagination) view returns((string,uint64)[] paymentAccountCounts, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentCaller) PaymentAccountCounts(opts *bind.CallOpts, pagination PageRequest) (struct {
	PaymentAccountCounts []PaymentAccountCount
	PageResponse         PageResponse
}, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "paymentAccountCounts", pagination)

	outstruct := new(struct {
		PaymentAccountCounts []PaymentAccountCount
		PageResponse         PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PaymentAccountCounts = *abi.ConvertType(out[0], new([]PaymentAccountCount)).(*[]PaymentAccountCount)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// PaymentAccountCounts is a free data retrieval call binding the contract method 0x1a101f68.
//
// Solidity: function paymentAccountCounts((bytes,uint64,uint64,bool,bool) pagination) view returns((string,uint64)[] paymentAccountCounts, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentSession) PaymentAccountCounts(pagination PageRequest) (struct {
	PaymentAccountCounts []PaymentAccountCount
	PageResponse         PageResponse
}, error) {
	return _IPayment.Contract.PaymentAccountCounts(&_IPayment.CallOpts, pagination)
}

// PaymentAccountCounts is a free data retrieval call binding the contract method 0x1a101f68.
//
// Solidity: function paymentAccountCounts((bytes,uint64,uint64,bool,bool) pagination) view returns((string,uint64)[] paymentAccountCounts, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentCallerSession) PaymentAccountCounts(pagination PageRequest) (struct {
	PaymentAccountCounts []PaymentAccountCount
	PageResponse         PageResponse
}, error) {
	return _IPayment.Contract.PaymentAccountCounts(&_IPayment.CallOpts, pagination)
}

// PaymentAccounts is a free data retrieval call binding the contract method 0x1757a1ee.
//
// Solidity: function paymentAccounts((bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool)[] paymentAccounts, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentCaller) PaymentAccounts(opts *bind.CallOpts, pagination PageRequest) (struct {
	PaymentAccounts []PaymentAccount
	PageResponse    PageResponse
}, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "paymentAccounts", pagination)

	outstruct := new(struct {
		PaymentAccounts []PaymentAccount
		PageResponse    PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PaymentAccounts = *abi.ConvertType(out[0], new([]PaymentAccount)).(*[]PaymentAccount)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// PaymentAccounts is a free data retrieval call binding the contract method 0x1757a1ee.
//
// Solidity: function paymentAccounts((bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool)[] paymentAccounts, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentSession) PaymentAccounts(pagination PageRequest) (struct {
	PaymentAccounts []PaymentAccount
	PageResponse    PageResponse
}, error) {
	return _IPayment.Contract.PaymentAccounts(&_IPayment.CallOpts, pagination)
}

// PaymentAccounts is a free data retrieval call binding the contract method 0x1757a1ee.
//
// Solidity: function paymentAccounts((bytes,uint64,uint64,bool,bool) pagination) view returns((string,string,bool)[] paymentAccounts, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentCallerSession) PaymentAccounts(pagination PageRequest) (struct {
	PaymentAccounts []PaymentAccount
	PageResponse    PageResponse
}, error) {
	return _IPayment.Contract.PaymentAccounts(&_IPayment.CallOpts, pagination)
}

// PaymentAccountsByOwner is a free data retrieval call binding the contract method 0x8909f228.
//
// Solidity: function paymentAccountsByOwner(string owner) view returns(string[] accounts)
func (_IPayment *IPaymentCaller) PaymentAccountsByOwner(opts *bind.CallOpts, owner string) ([]string, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "paymentAccountsByOwner", owner)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// PaymentAccountsByOwner is a free data retrieval call binding the contract method 0x8909f228.
//
// Solidity: function paymentAccountsByOwner(string owner) view returns(string[] accounts)
func (_IPayment *IPaymentSession) PaymentAccountsByOwner(owner string) ([]string, error) {
	return _IPayment.Contract.PaymentAccountsByOwner(&_IPayment.CallOpts, owner)
}

// PaymentAccountsByOwner is a free data retrieval call binding the contract method 0x8909f228.
//
// Solidity: function paymentAccountsByOwner(string owner) view returns(string[] accounts)
func (_IPayment *IPaymentCallerSession) PaymentAccountsByOwner(owner string) ([]string, error) {
	return _IPayment.Contract.PaymentAccountsByOwner(&_IPayment.CallOpts, owner)
}

// StreamRecord is a free data retrieval call binding the contract method 0x56965b8f.
//
// Solidity: function streamRecord(string account) view returns((string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256) streamRecord)
func (_IPayment *IPaymentCaller) StreamRecord(opts *bind.CallOpts, account string) (StreamRecord, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "streamRecord", account)

	if err != nil {
		return *new(StreamRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(StreamRecord)).(*StreamRecord)

	return out0, err

}

// StreamRecord is a free data retrieval call binding the contract method 0x56965b8f.
//
// Solidity: function streamRecord(string account) view returns((string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256) streamRecord)
func (_IPayment *IPaymentSession) StreamRecord(account string) (StreamRecord, error) {
	return _IPayment.Contract.StreamRecord(&_IPayment.CallOpts, account)
}

// StreamRecord is a free data retrieval call binding the contract method 0x56965b8f.
//
// Solidity: function streamRecord(string account) view returns((string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256) streamRecord)
func (_IPayment *IPaymentCallerSession) StreamRecord(account string) (StreamRecord, error) {
	return _IPayment.Contract.StreamRecord(&_IPayment.CallOpts, account)
}

// StreamRecords is a free data retrieval call binding the contract method 0x33e16e34.
//
// Solidity: function streamRecords((bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256)[] streamRecords, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentCaller) StreamRecords(opts *bind.CallOpts, pagination PageRequest) (struct {
	StreamRecords []StreamRecord
	PageResponse  PageResponse
}, error) {
	var out []interface{}
	err := _IPayment.contract.Call(opts, &out, "streamRecords", pagination)

	outstruct := new(struct {
		StreamRecords []StreamRecord
		PageResponse  PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StreamRecords = *abi.ConvertType(out[0], new([]StreamRecord)).(*[]StreamRecord)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// StreamRecords is a free data retrieval call binding the contract method 0x33e16e34.
//
// Solidity: function streamRecords((bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256)[] streamRecords, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentSession) StreamRecords(pagination PageRequest) (struct {
	StreamRecords []StreamRecord
	PageResponse  PageResponse
}, error) {
	return _IPayment.Contract.StreamRecords(&_IPayment.CallOpts, pagination)
}

// StreamRecords is a free data retrieval call binding the contract method 0x33e16e34.
//
// Solidity: function streamRecords((bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64,uint256,uint256,uint256,uint256,int32,int64,uint64,uint256)[] streamRecords, (bytes,uint64) pageResponse)
func (_IPayment *IPaymentCallerSession) StreamRecords(pagination PageRequest) (struct {
	StreamRecords []StreamRecord
	PageResponse  PageResponse
}, error) {
	return _IPayment.Contract.StreamRecords(&_IPayment.CallOpts, pagination)
}

// CreatePaymentAccount is a paid mutator transaction binding the contract method 0xe3229ba5.
//
// Solidity: function createPaymentAccount() returns(bool success)
func (_IPayment *IPaymentTransactor) CreatePaymentAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPayment.contract.Transact(opts, "createPaymentAccount")
}

// CreatePaymentAccount is a paid mutator transaction binding the contract method 0xe3229ba5.
//
// Solidity: function createPaymentAccount() returns(bool success)
func (_IPayment *IPaymentSession) CreatePaymentAccount() (*types.Transaction, error) {
	return _IPayment.Contract.CreatePaymentAccount(&_IPayment.TransactOpts)
}

// CreatePaymentAccount is a paid mutator transaction binding the contract method 0xe3229ba5.
//
// Solidity: function createPaymentAccount() returns(bool success)
func (_IPayment *IPaymentTransactorSession) CreatePaymentAccount() (*types.Transaction, error) {
	return _IPayment.Contract.CreatePaymentAccount(&_IPayment.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8e27d719.
//
// Solidity: function deposit(string to, uint256 amount) returns(bool success)
func (_IPayment *IPaymentTransactor) Deposit(opts *bind.TransactOpts, to string, amount *big.Int) (*types.Transaction, error) {
	return _IPayment.contract.Transact(opts, "deposit", to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8e27d719.
//
// Solidity: function deposit(string to, uint256 amount) returns(bool success)
func (_IPayment *IPaymentSession) Deposit(to string, amount *big.Int) (*types.Transaction, error) {
	return _IPayment.Contract.Deposit(&_IPayment.TransactOpts, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x8e27d719.
//
// Solidity: function deposit(string to, uint256 amount) returns(bool success)
func (_IPayment *IPaymentTransactorSession) Deposit(to string, amount *big.Int) (*types.Transaction, error) {
	return _IPayment.Contract.Deposit(&_IPayment.TransactOpts, to, amount)
}

// DisableRefund is a paid mutator transaction binding the contract method 0xac43b761.
//
// Solidity: function disableRefund(string addr) returns(bool success)
func (_IPayment *IPaymentTransactor) DisableRefund(opts *bind.TransactOpts, addr string) (*types.Transaction, error) {
	return _IPayment.contract.Transact(opts, "disableRefund", addr)
}

// DisableRefund is a paid mutator transaction binding the contract method 0xac43b761.
//
// Solidity: function disableRefund(string addr) returns(bool success)
func (_IPayment *IPaymentSession) DisableRefund(addr string) (*types.Transaction, error) {
	return _IPayment.Contract.DisableRefund(&_IPayment.TransactOpts, addr)
}

// DisableRefund is a paid mutator transaction binding the contract method 0xac43b761.
//
// Solidity: function disableRefund(string addr) returns(bool success)
func (_IPayment *IPaymentTransactorSession) DisableRefund(addr string) (*types.Transaction, error) {
	return _IPayment.Contract.DisableRefund(&_IPayment.TransactOpts, addr)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x6737b841.
//
// Solidity: function updateParams(string authority, ((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params) returns(bool success)
func (_IPayment *IPaymentTransactor) UpdateParams(opts *bind.TransactOpts, authority string, params Params) (*types.Transaction, error) {
	return _IPayment.contract.Transact(opts, "updateParams", authority, params)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x6737b841.
//
// Solidity: function updateParams(string authority, ((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params) returns(bool success)
func (_IPayment *IPaymentSession) UpdateParams(authority string, params Params) (*types.Transaction, error) {
	return _IPayment.Contract.UpdateParams(&_IPayment.TransactOpts, authority, params)
}

// UpdateParams is a paid mutator transaction binding the contract method 0x6737b841.
//
// Solidity: function updateParams(string authority, ((uint64,uint256),uint64,uint64,uint64,uint64,string,uint256,uint64) params) returns(bool success)
func (_IPayment *IPaymentTransactorSession) UpdateParams(authority string, params Params) (*types.Transaction, error) {
	return _IPayment.Contract.UpdateParams(&_IPayment.TransactOpts, authority, params)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string from, uint256 amount) returns(bool success)
func (_IPayment *IPaymentTransactor) Withdraw(opts *bind.TransactOpts, from string, amount *big.Int) (*types.Transaction, error) {
	return _IPayment.contract.Transact(opts, "withdraw", from, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string from, uint256 amount) returns(bool success)
func (_IPayment *IPaymentSession) Withdraw(from string, amount *big.Int) (*types.Transaction, error) {
	return _IPayment.Contract.Withdraw(&_IPayment.TransactOpts, from, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x30b39a62.
//
// Solidity: function withdraw(string from, uint256 amount) returns(bool success)
func (_IPayment *IPaymentTransactorSession) Withdraw(from string, amount *big.Int) (*types.Transaction, error) {
	return _IPayment.Contract.Withdraw(&_IPayment.TransactOpts, from, amount)
}

// IPaymentCreatePaymentAccountIterator is returned from FilterCreatePaymentAccount and is used to iterate over the raw logs and unpacked data for CreatePaymentAccount events raised by the IPayment contract.
type IPaymentCreatePaymentAccountIterator struct {
	Event *IPaymentCreatePaymentAccount // Event containing the contract specifics and raw log

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
func (it *IPaymentCreatePaymentAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentCreatePaymentAccount)
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
		it.Event = new(IPaymentCreatePaymentAccount)
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
func (it *IPaymentCreatePaymentAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentCreatePaymentAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentCreatePaymentAccount represents a CreatePaymentAccount event raised by the IPayment contract.
type IPaymentCreatePaymentAccount struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreatePaymentAccount is a free log retrieval operation binding the contract event 0xe51c69f2d31f4a4a9eb11e314f91e7edee9b6549fb95b8aba70df8b0ca3b502b.
//
// Solidity: event CreatePaymentAccount(address indexed creator)
func (_IPayment *IPaymentFilterer) FilterCreatePaymentAccount(opts *bind.FilterOpts, creator []common.Address) (*IPaymentCreatePaymentAccountIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IPayment.contract.FilterLogs(opts, "CreatePaymentAccount", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentCreatePaymentAccountIterator{contract: _IPayment.contract, event: "CreatePaymentAccount", logs: logs, sub: sub}, nil
}

// WatchCreatePaymentAccount is a free log subscription operation binding the contract event 0xe51c69f2d31f4a4a9eb11e314f91e7edee9b6549fb95b8aba70df8b0ca3b502b.
//
// Solidity: event CreatePaymentAccount(address indexed creator)
func (_IPayment *IPaymentFilterer) WatchCreatePaymentAccount(opts *bind.WatchOpts, sink chan<- *IPaymentCreatePaymentAccount, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IPayment.contract.WatchLogs(opts, "CreatePaymentAccount", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentCreatePaymentAccount)
				if err := _IPayment.contract.UnpackLog(event, "CreatePaymentAccount", log); err != nil {
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

// ParseCreatePaymentAccount is a log parse operation binding the contract event 0xe51c69f2d31f4a4a9eb11e314f91e7edee9b6549fb95b8aba70df8b0ca3b502b.
//
// Solidity: event CreatePaymentAccount(address indexed creator)
func (_IPayment *IPaymentFilterer) ParseCreatePaymentAccount(log types.Log) (*IPaymentCreatePaymentAccount, error) {
	event := new(IPaymentCreatePaymentAccount)
	if err := _IPayment.contract.UnpackLog(event, "CreatePaymentAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IPayment contract.
type IPaymentDepositIterator struct {
	Event *IPaymentDeposit // Event containing the contract specifics and raw log

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
func (it *IPaymentDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentDeposit)
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
		it.Event = new(IPaymentDeposit)
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
func (it *IPaymentDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentDeposit represents a Deposit event raised by the IPayment contract.
type IPaymentDeposit struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x8ce0bd46ec50cf39f0d0ea8686a686eb226af5796dcda4231b26fb84b5ef1234.
//
// Solidity: event Deposit(address indexed operator)
func (_IPayment *IPaymentFilterer) FilterDeposit(opts *bind.FilterOpts, operator []common.Address) (*IPaymentDepositIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IPayment.contract.FilterLogs(opts, "Deposit", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentDepositIterator{contract: _IPayment.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x8ce0bd46ec50cf39f0d0ea8686a686eb226af5796dcda4231b26fb84b5ef1234.
//
// Solidity: event Deposit(address indexed operator)
func (_IPayment *IPaymentFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IPaymentDeposit, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IPayment.contract.WatchLogs(opts, "Deposit", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentDeposit)
				if err := _IPayment.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x8ce0bd46ec50cf39f0d0ea8686a686eb226af5796dcda4231b26fb84b5ef1234.
//
// Solidity: event Deposit(address indexed operator)
func (_IPayment *IPaymentFilterer) ParseDeposit(log types.Log) (*IPaymentDeposit, error) {
	event := new(IPaymentDeposit)
	if err := _IPayment.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentDisableRefundIterator is returned from FilterDisableRefund and is used to iterate over the raw logs and unpacked data for DisableRefund events raised by the IPayment contract.
type IPaymentDisableRefundIterator struct {
	Event *IPaymentDisableRefund // Event containing the contract specifics and raw log

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
func (it *IPaymentDisableRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentDisableRefund)
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
		it.Event = new(IPaymentDisableRefund)
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
func (it *IPaymentDisableRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentDisableRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentDisableRefund represents a DisableRefund event raised by the IPayment contract.
type IPaymentDisableRefund struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDisableRefund is a free log retrieval operation binding the contract event 0x5dab65934c4c1b887963b93f7d8505184b47126e065b2b1a1743a50699b5c48b.
//
// Solidity: event DisableRefund(address indexed owner)
func (_IPayment *IPaymentFilterer) FilterDisableRefund(opts *bind.FilterOpts, owner []common.Address) (*IPaymentDisableRefundIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IPayment.contract.FilterLogs(opts, "DisableRefund", ownerRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentDisableRefundIterator{contract: _IPayment.contract, event: "DisableRefund", logs: logs, sub: sub}, nil
}

// WatchDisableRefund is a free log subscription operation binding the contract event 0x5dab65934c4c1b887963b93f7d8505184b47126e065b2b1a1743a50699b5c48b.
//
// Solidity: event DisableRefund(address indexed owner)
func (_IPayment *IPaymentFilterer) WatchDisableRefund(opts *bind.WatchOpts, sink chan<- *IPaymentDisableRefund, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _IPayment.contract.WatchLogs(opts, "DisableRefund", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentDisableRefund)
				if err := _IPayment.contract.UnpackLog(event, "DisableRefund", log); err != nil {
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

// ParseDisableRefund is a log parse operation binding the contract event 0x5dab65934c4c1b887963b93f7d8505184b47126e065b2b1a1743a50699b5c48b.
//
// Solidity: event DisableRefund(address indexed owner)
func (_IPayment *IPaymentFilterer) ParseDisableRefund(log types.Log) (*IPaymentDisableRefund, error) {
	event := new(IPaymentDisableRefund)
	if err := _IPayment.contract.UnpackLog(event, "DisableRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentUpdateParamsIterator is returned from FilterUpdateParams and is used to iterate over the raw logs and unpacked data for UpdateParams events raised by the IPayment contract.
type IPaymentUpdateParamsIterator struct {
	Event *IPaymentUpdateParams // Event containing the contract specifics and raw log

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
func (it *IPaymentUpdateParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentUpdateParams)
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
		it.Event = new(IPaymentUpdateParams)
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
func (it *IPaymentUpdateParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentUpdateParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentUpdateParams represents a UpdateParams event raised by the IPayment contract.
type IPaymentUpdateParams struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateParams is a free log retrieval operation binding the contract event 0xdb2e743561971fc97db45300fcd12c1e9e20b20d485aa6f0259c3374f4a4dafd.
//
// Solidity: event UpdateParams(address indexed creator)
func (_IPayment *IPaymentFilterer) FilterUpdateParams(opts *bind.FilterOpts, creator []common.Address) (*IPaymentUpdateParamsIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IPayment.contract.FilterLogs(opts, "UpdateParams", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentUpdateParamsIterator{contract: _IPayment.contract, event: "UpdateParams", logs: logs, sub: sub}, nil
}

// WatchUpdateParams is a free log subscription operation binding the contract event 0xdb2e743561971fc97db45300fcd12c1e9e20b20d485aa6f0259c3374f4a4dafd.
//
// Solidity: event UpdateParams(address indexed creator)
func (_IPayment *IPaymentFilterer) WatchUpdateParams(opts *bind.WatchOpts, sink chan<- *IPaymentUpdateParams, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IPayment.contract.WatchLogs(opts, "UpdateParams", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentUpdateParams)
				if err := _IPayment.contract.UnpackLog(event, "UpdateParams", log); err != nil {
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

// ParseUpdateParams is a log parse operation binding the contract event 0xdb2e743561971fc97db45300fcd12c1e9e20b20d485aa6f0259c3374f4a4dafd.
//
// Solidity: event UpdateParams(address indexed creator)
func (_IPayment *IPaymentFilterer) ParseUpdateParams(log types.Log) (*IPaymentUpdateParams, error) {
	event := new(IPaymentUpdateParams)
	if err := _IPayment.contract.UnpackLog(event, "UpdateParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPaymentWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IPayment contract.
type IPaymentWithdrawIterator struct {
	Event *IPaymentWithdraw // Event containing the contract specifics and raw log

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
func (it *IPaymentWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPaymentWithdraw)
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
		it.Event = new(IPaymentWithdraw)
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
func (it *IPaymentWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPaymentWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPaymentWithdraw represents a Withdraw event raised by the IPayment contract.
type IPaymentWithdraw struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf67611512e0a2d90c96fd3f08dca4971bc45fba9dc679eabe839a32abbe58a8e.
//
// Solidity: event Withdraw(address indexed creator)
func (_IPayment *IPaymentFilterer) FilterWithdraw(opts *bind.FilterOpts, creator []common.Address) (*IPaymentWithdrawIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IPayment.contract.FilterLogs(opts, "Withdraw", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IPaymentWithdrawIterator{contract: _IPayment.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf67611512e0a2d90c96fd3f08dca4971bc45fba9dc679eabe839a32abbe58a8e.
//
// Solidity: event Withdraw(address indexed creator)
func (_IPayment *IPaymentFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IPaymentWithdraw, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IPayment.contract.WatchLogs(opts, "Withdraw", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPaymentWithdraw)
				if err := _IPayment.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf67611512e0a2d90c96fd3f08dca4971bc45fba9dc679eabe839a32abbe58a8e.
//
// Solidity: event Withdraw(address indexed creator)
func (_IPayment *IPaymentFilterer) ParseWithdraw(log types.Log) (*IPaymentWithdraw, error) {
	event := new(IPaymentWithdraw)
	if err := _IPayment.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
