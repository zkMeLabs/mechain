// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package virtualgroup

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

// Approval is an auto generated low-level Go binding around an user-defined struct.
type Approval struct {
	ExpiredHeight              uint64
	GlobalVirtualGroupFamilyId uint32
	Sig                        []byte
}

// Coin is an auto generated low-level Go binding around an user-defined struct.
type Coin struct {
	Denom  string
	Amount *big.Int
}

// GlobalVirtualGroupFamily is an auto generated low-level Go binding around an user-defined struct.
type GlobalVirtualGroupFamily struct {
	Id                    uint32
	PrimarySpId           uint32
	GlobalVirtualGroupIds []uint32
	VirtualPaymentAddress common.Address
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

// IVirtualGroupMetaData contains all meta data concerning the IVirtualGroup contract.
var IVirtualGroupMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"}],\"name\":\"CancelSwapIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"CompleteSPExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"}],\"name\":\"CompleteSwapIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"}],\"name\":\"CreateGlobalVirtualGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"}],\"name\":\"DeleteGlobalVirtualGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"}],\"name\":\"ReserveSwapIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"}],\"name\":\"SPExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"}],\"name\":\"SwapOut\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gvgFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"}],\"name\":\"cancelSwapIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"storageProvider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"operator\",\"type\":\"string\"}],\"name\":\"completeSPExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gvgFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"}],\"name\":\"completeSwapIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gvgFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"gvgIds\",\"type\":\"uint32[]\"}],\"name\":\"completeSwapOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"familyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"secondarySpIds\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"createGlobalVirtualGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"}],\"name\":\"deleteGlobalVirtualGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"globalVirtualGroupFamilies\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"primarySpId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"globalVirtualGroupIds\",\"type\":\"uint32[]\"},{\"internalType\":\"address\",\"name\":\"virtualPaymentAddress\",\"type\":\"address\"}],\"internalType\":\"structGlobalVirtualGroupFamily[]\",\"name\":\"gvgFamilies\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"familyId\",\"type\":\"uint32\"}],\"name\":\"globalVirtualGroupFamily\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"primarySpId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"globalVirtualGroupIds\",\"type\":\"uint32[]\"},{\"internalType\":\"address\",\"name\":\"virtualPaymentAddress\",\"type\":\"address\"}],\"internalType\":\"structGlobalVirtualGroupFamily\",\"name\":\"gvgfamily\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"targetSpId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gvgFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupId\",\"type\":\"uint32\"}],\"name\":\"reserveSwapIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spExit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gvgFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"gvgIds\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32\",\"name\":\"successorSpId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"expiredHeight\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"globalVirtualGroupFamilyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structApproval\",\"name\":\"successorSpApproval\",\"type\":\"tuple\"}],\"name\":\"swapOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IVirtualGroupABI is the input ABI used to generate the binding from.
// Deprecated: Use IVirtualGroupMetaData.ABI instead.
var IVirtualGroupABI = IVirtualGroupMetaData.ABI

// IVirtualGroup is an auto generated Go binding around an Ethereum contract.
type IVirtualGroup struct {
	IVirtualGroupCaller     // Read-only binding to the contract
	IVirtualGroupTransactor // Write-only binding to the contract
	IVirtualGroupFilterer   // Log filterer for contract events
}

// IVirtualGroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVirtualGroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVirtualGroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVirtualGroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVirtualGroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVirtualGroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVirtualGroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVirtualGroupSession struct {
	Contract     *IVirtualGroup    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVirtualGroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVirtualGroupCallerSession struct {
	Contract *IVirtualGroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IVirtualGroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVirtualGroupTransactorSession struct {
	Contract     *IVirtualGroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IVirtualGroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVirtualGroupRaw struct {
	Contract *IVirtualGroup // Generic contract binding to access the raw methods on
}

// IVirtualGroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVirtualGroupCallerRaw struct {
	Contract *IVirtualGroupCaller // Generic read-only contract binding to access the raw methods on
}

// IVirtualGroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVirtualGroupTransactorRaw struct {
	Contract *IVirtualGroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVirtualGroup creates a new instance of IVirtualGroup, bound to a specific deployed contract.
func NewIVirtualGroup(address common.Address, backend bind.ContractBackend) (*IVirtualGroup, error) {
	contract, err := bindIVirtualGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroup{IVirtualGroupCaller: IVirtualGroupCaller{contract: contract}, IVirtualGroupTransactor: IVirtualGroupTransactor{contract: contract}, IVirtualGroupFilterer: IVirtualGroupFilterer{contract: contract}}, nil
}

// NewIVirtualGroupCaller creates a new read-only instance of IVirtualGroup, bound to a specific deployed contract.
func NewIVirtualGroupCaller(address common.Address, caller bind.ContractCaller) (*IVirtualGroupCaller, error) {
	contract, err := bindIVirtualGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupCaller{contract: contract}, nil
}

// NewIVirtualGroupTransactor creates a new write-only instance of IVirtualGroup, bound to a specific deployed contract.
func NewIVirtualGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*IVirtualGroupTransactor, error) {
	contract, err := bindIVirtualGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupTransactor{contract: contract}, nil
}

// NewIVirtualGroupFilterer creates a new log filterer instance of IVirtualGroup, bound to a specific deployed contract.
func NewIVirtualGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*IVirtualGroupFilterer, error) {
	contract, err := bindIVirtualGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupFilterer{contract: contract}, nil
}

// bindIVirtualGroup binds a generic wrapper to an already deployed contract.
func bindIVirtualGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVirtualGroupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVirtualGroup *IVirtualGroupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVirtualGroup.Contract.IVirtualGroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVirtualGroup *IVirtualGroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.IVirtualGroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVirtualGroup *IVirtualGroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.IVirtualGroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVirtualGroup *IVirtualGroupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVirtualGroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVirtualGroup *IVirtualGroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVirtualGroup *IVirtualGroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.contract.Transact(opts, method, params...)
}

// GlobalVirtualGroupFamilies is a free data retrieval call binding the contract method 0xb975208b.
//
// Solidity: function globalVirtualGroupFamilies((bytes,uint64,uint64,bool,bool) pagination) view returns((uint32,uint32,uint32[],address)[] gvgFamilies, (bytes,uint64) pageResponse)
func (_IVirtualGroup *IVirtualGroupCaller) GlobalVirtualGroupFamilies(opts *bind.CallOpts, pagination PageRequest) (struct {
	GvgFamilies  []GlobalVirtualGroupFamily
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IVirtualGroup.contract.Call(opts, &out, "globalVirtualGroupFamilies", pagination)

	outstruct := new(struct {
		GvgFamilies  []GlobalVirtualGroupFamily
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GvgFamilies = *abi.ConvertType(out[0], new([]GlobalVirtualGroupFamily)).(*[]GlobalVirtualGroupFamily)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// GlobalVirtualGroupFamilies is a free data retrieval call binding the contract method 0xb975208b.
//
// Solidity: function globalVirtualGroupFamilies((bytes,uint64,uint64,bool,bool) pagination) view returns((uint32,uint32,uint32[],address)[] gvgFamilies, (bytes,uint64) pageResponse)
func (_IVirtualGroup *IVirtualGroupSession) GlobalVirtualGroupFamilies(pagination PageRequest) (struct {
	GvgFamilies  []GlobalVirtualGroupFamily
	PageResponse PageResponse
}, error) {
	return _IVirtualGroup.Contract.GlobalVirtualGroupFamilies(&_IVirtualGroup.CallOpts, pagination)
}

// GlobalVirtualGroupFamilies is a free data retrieval call binding the contract method 0xb975208b.
//
// Solidity: function globalVirtualGroupFamilies((bytes,uint64,uint64,bool,bool) pagination) view returns((uint32,uint32,uint32[],address)[] gvgFamilies, (bytes,uint64) pageResponse)
func (_IVirtualGroup *IVirtualGroupCallerSession) GlobalVirtualGroupFamilies(pagination PageRequest) (struct {
	GvgFamilies  []GlobalVirtualGroupFamily
	PageResponse PageResponse
}, error) {
	return _IVirtualGroup.Contract.GlobalVirtualGroupFamilies(&_IVirtualGroup.CallOpts, pagination)
}

// GlobalVirtualGroupFamily is a free data retrieval call binding the contract method 0x33fcaf0a.
//
// Solidity: function globalVirtualGroupFamily(uint32 familyId) view returns((uint32,uint32,uint32[],address) gvgfamily)
func (_IVirtualGroup *IVirtualGroupCaller) GlobalVirtualGroupFamily(opts *bind.CallOpts, familyId uint32) (GlobalVirtualGroupFamily, error) {
	var out []interface{}
	err := _IVirtualGroup.contract.Call(opts, &out, "globalVirtualGroupFamily", familyId)

	if err != nil {
		return *new(GlobalVirtualGroupFamily), err
	}

	out0 := *abi.ConvertType(out[0], new(GlobalVirtualGroupFamily)).(*GlobalVirtualGroupFamily)

	return out0, err

}

// GlobalVirtualGroupFamily is a free data retrieval call binding the contract method 0x33fcaf0a.
//
// Solidity: function globalVirtualGroupFamily(uint32 familyId) view returns((uint32,uint32,uint32[],address) gvgfamily)
func (_IVirtualGroup *IVirtualGroupSession) GlobalVirtualGroupFamily(familyId uint32) (GlobalVirtualGroupFamily, error) {
	return _IVirtualGroup.Contract.GlobalVirtualGroupFamily(&_IVirtualGroup.CallOpts, familyId)
}

// GlobalVirtualGroupFamily is a free data retrieval call binding the contract method 0x33fcaf0a.
//
// Solidity: function globalVirtualGroupFamily(uint32 familyId) view returns((uint32,uint32,uint32[],address) gvgfamily)
func (_IVirtualGroup *IVirtualGroupCallerSession) GlobalVirtualGroupFamily(familyId uint32) (GlobalVirtualGroupFamily, error) {
	return _IVirtualGroup.Contract.GlobalVirtualGroupFamily(&_IVirtualGroup.CallOpts, familyId)
}

// CancelSwapIn is a paid mutator transaction binding the contract method 0x460b4ce3.
//
// Solidity: function cancelSwapIn(uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) CancelSwapIn(opts *bind.TransactOpts, gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "cancelSwapIn", gvgFamilyId, globalVirtualGroupId)
}

// CancelSwapIn is a paid mutator transaction binding the contract method 0x460b4ce3.
//
// Solidity: function cancelSwapIn(uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) CancelSwapIn(gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CancelSwapIn(&_IVirtualGroup.TransactOpts, gvgFamilyId, globalVirtualGroupId)
}

// CancelSwapIn is a paid mutator transaction binding the contract method 0x460b4ce3.
//
// Solidity: function cancelSwapIn(uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) CancelSwapIn(gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CancelSwapIn(&_IVirtualGroup.TransactOpts, gvgFamilyId, globalVirtualGroupId)
}

// CompleteSPExit is a paid mutator transaction binding the contract method 0x44853205.
//
// Solidity: function completeSPExit(string storageProvider, string operator) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) CompleteSPExit(opts *bind.TransactOpts, storageProvider string, operator string) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "completeSPExit", storageProvider, operator)
}

// CompleteSPExit is a paid mutator transaction binding the contract method 0x44853205.
//
// Solidity: function completeSPExit(string storageProvider, string operator) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) CompleteSPExit(storageProvider string, operator string) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CompleteSPExit(&_IVirtualGroup.TransactOpts, storageProvider, operator)
}

// CompleteSPExit is a paid mutator transaction binding the contract method 0x44853205.
//
// Solidity: function completeSPExit(string storageProvider, string operator) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) CompleteSPExit(storageProvider string, operator string) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CompleteSPExit(&_IVirtualGroup.TransactOpts, storageProvider, operator)
}

// CompleteSwapIn is a paid mutator transaction binding the contract method 0x82b669d0.
//
// Solidity: function completeSwapIn(uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) CompleteSwapIn(opts *bind.TransactOpts, gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "completeSwapIn", gvgFamilyId, globalVirtualGroupId)
}

// CompleteSwapIn is a paid mutator transaction binding the contract method 0x82b669d0.
//
// Solidity: function completeSwapIn(uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) CompleteSwapIn(gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CompleteSwapIn(&_IVirtualGroup.TransactOpts, gvgFamilyId, globalVirtualGroupId)
}

// CompleteSwapIn is a paid mutator transaction binding the contract method 0x82b669d0.
//
// Solidity: function completeSwapIn(uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) CompleteSwapIn(gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CompleteSwapIn(&_IVirtualGroup.TransactOpts, gvgFamilyId, globalVirtualGroupId)
}

// CompleteSwapOut is a paid mutator transaction binding the contract method 0xebd2beaa.
//
// Solidity: function completeSwapOut(uint32 gvgFamilyId, uint32[] gvgIds) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) CompleteSwapOut(opts *bind.TransactOpts, gvgFamilyId uint32, gvgIds []uint32) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "completeSwapOut", gvgFamilyId, gvgIds)
}

// CompleteSwapOut is a paid mutator transaction binding the contract method 0xebd2beaa.
//
// Solidity: function completeSwapOut(uint32 gvgFamilyId, uint32[] gvgIds) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) CompleteSwapOut(gvgFamilyId uint32, gvgIds []uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CompleteSwapOut(&_IVirtualGroup.TransactOpts, gvgFamilyId, gvgIds)
}

// CompleteSwapOut is a paid mutator transaction binding the contract method 0xebd2beaa.
//
// Solidity: function completeSwapOut(uint32 gvgFamilyId, uint32[] gvgIds) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) CompleteSwapOut(gvgFamilyId uint32, gvgIds []uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CompleteSwapOut(&_IVirtualGroup.TransactOpts, gvgFamilyId, gvgIds)
}

// CreateGlobalVirtualGroup is a paid mutator transaction binding the contract method 0x92f6284a.
//
// Solidity: function createGlobalVirtualGroup(uint32 familyId, uint32[] secondarySpIds, (string,uint256) deposit) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) CreateGlobalVirtualGroup(opts *bind.TransactOpts, familyId uint32, secondarySpIds []uint32, deposit Coin) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "createGlobalVirtualGroup", familyId, secondarySpIds, deposit)
}

// CreateGlobalVirtualGroup is a paid mutator transaction binding the contract method 0x92f6284a.
//
// Solidity: function createGlobalVirtualGroup(uint32 familyId, uint32[] secondarySpIds, (string,uint256) deposit) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) CreateGlobalVirtualGroup(familyId uint32, secondarySpIds []uint32, deposit Coin) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CreateGlobalVirtualGroup(&_IVirtualGroup.TransactOpts, familyId, secondarySpIds, deposit)
}

// CreateGlobalVirtualGroup is a paid mutator transaction binding the contract method 0x92f6284a.
//
// Solidity: function createGlobalVirtualGroup(uint32 familyId, uint32[] secondarySpIds, (string,uint256) deposit) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) CreateGlobalVirtualGroup(familyId uint32, secondarySpIds []uint32, deposit Coin) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.CreateGlobalVirtualGroup(&_IVirtualGroup.TransactOpts, familyId, secondarySpIds, deposit)
}

// DeleteGlobalVirtualGroup is a paid mutator transaction binding the contract method 0x3712b375.
//
// Solidity: function deleteGlobalVirtualGroup(uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) DeleteGlobalVirtualGroup(opts *bind.TransactOpts, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "deleteGlobalVirtualGroup", globalVirtualGroupId)
}

// DeleteGlobalVirtualGroup is a paid mutator transaction binding the contract method 0x3712b375.
//
// Solidity: function deleteGlobalVirtualGroup(uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) DeleteGlobalVirtualGroup(globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.DeleteGlobalVirtualGroup(&_IVirtualGroup.TransactOpts, globalVirtualGroupId)
}

// DeleteGlobalVirtualGroup is a paid mutator transaction binding the contract method 0x3712b375.
//
// Solidity: function deleteGlobalVirtualGroup(uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) DeleteGlobalVirtualGroup(globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.DeleteGlobalVirtualGroup(&_IVirtualGroup.TransactOpts, globalVirtualGroupId)
}

// Deposit is a paid mutator transaction binding the contract method 0xd4755928.
//
// Solidity: function deposit(uint32 globalVirtualGroupId, (string,uint256) deposit) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) Deposit(opts *bind.TransactOpts, globalVirtualGroupId uint32, deposit Coin) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "deposit", globalVirtualGroupId, deposit)
}

// Deposit is a paid mutator transaction binding the contract method 0xd4755928.
//
// Solidity: function deposit(uint32 globalVirtualGroupId, (string,uint256) deposit) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) Deposit(globalVirtualGroupId uint32, deposit Coin) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.Deposit(&_IVirtualGroup.TransactOpts, globalVirtualGroupId, deposit)
}

// Deposit is a paid mutator transaction binding the contract method 0xd4755928.
//
// Solidity: function deposit(uint32 globalVirtualGroupId, (string,uint256) deposit) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) Deposit(globalVirtualGroupId uint32, deposit Coin) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.Deposit(&_IVirtualGroup.TransactOpts, globalVirtualGroupId, deposit)
}

// ReserveSwapIn is a paid mutator transaction binding the contract method 0xdc2d3f12.
//
// Solidity: function reserveSwapIn(uint32 targetSpId, uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) ReserveSwapIn(opts *bind.TransactOpts, targetSpId uint32, gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "reserveSwapIn", targetSpId, gvgFamilyId, globalVirtualGroupId)
}

// ReserveSwapIn is a paid mutator transaction binding the contract method 0xdc2d3f12.
//
// Solidity: function reserveSwapIn(uint32 targetSpId, uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) ReserveSwapIn(targetSpId uint32, gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.ReserveSwapIn(&_IVirtualGroup.TransactOpts, targetSpId, gvgFamilyId, globalVirtualGroupId)
}

// ReserveSwapIn is a paid mutator transaction binding the contract method 0xdc2d3f12.
//
// Solidity: function reserveSwapIn(uint32 targetSpId, uint32 gvgFamilyId, uint32 globalVirtualGroupId) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) ReserveSwapIn(targetSpId uint32, gvgFamilyId uint32, globalVirtualGroupId uint32) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.ReserveSwapIn(&_IVirtualGroup.TransactOpts, targetSpId, gvgFamilyId, globalVirtualGroupId)
}

// SpExit is a paid mutator transaction binding the contract method 0xecdec1ab.
//
// Solidity: function spExit() returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) SpExit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "spExit")
}

// SpExit is a paid mutator transaction binding the contract method 0xecdec1ab.
//
// Solidity: function spExit() returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) SpExit() (*types.Transaction, error) {
	return _IVirtualGroup.Contract.SpExit(&_IVirtualGroup.TransactOpts)
}

// SpExit is a paid mutator transaction binding the contract method 0xecdec1ab.
//
// Solidity: function spExit() returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) SpExit() (*types.Transaction, error) {
	return _IVirtualGroup.Contract.SpExit(&_IVirtualGroup.TransactOpts)
}

// SwapOut is a paid mutator transaction binding the contract method 0x2cb3f1d6.
//
// Solidity: function swapOut(uint32 gvgFamilyId, uint32[] gvgIds, uint32 successorSpId, (uint64,uint32,bytes) successorSpApproval) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactor) SwapOut(opts *bind.TransactOpts, gvgFamilyId uint32, gvgIds []uint32, successorSpId uint32, successorSpApproval Approval) (*types.Transaction, error) {
	return _IVirtualGroup.contract.Transact(opts, "swapOut", gvgFamilyId, gvgIds, successorSpId, successorSpApproval)
}

// SwapOut is a paid mutator transaction binding the contract method 0x2cb3f1d6.
//
// Solidity: function swapOut(uint32 gvgFamilyId, uint32[] gvgIds, uint32 successorSpId, (uint64,uint32,bytes) successorSpApproval) returns(bool success)
func (_IVirtualGroup *IVirtualGroupSession) SwapOut(gvgFamilyId uint32, gvgIds []uint32, successorSpId uint32, successorSpApproval Approval) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.SwapOut(&_IVirtualGroup.TransactOpts, gvgFamilyId, gvgIds, successorSpId, successorSpApproval)
}

// SwapOut is a paid mutator transaction binding the contract method 0x2cb3f1d6.
//
// Solidity: function swapOut(uint32 gvgFamilyId, uint32[] gvgIds, uint32 successorSpId, (uint64,uint32,bytes) successorSpApproval) returns(bool success)
func (_IVirtualGroup *IVirtualGroupTransactorSession) SwapOut(gvgFamilyId uint32, gvgIds []uint32, successorSpId uint32, successorSpApproval Approval) (*types.Transaction, error) {
	return _IVirtualGroup.Contract.SwapOut(&_IVirtualGroup.TransactOpts, gvgFamilyId, gvgIds, successorSpId, successorSpApproval)
}

// IVirtualGroupCancelSwapInIterator is returned from FilterCancelSwapIn and is used to iterate over the raw logs and unpacked data for CancelSwapIn events raised by the IVirtualGroup contract.
type IVirtualGroupCancelSwapInIterator struct {
	Event *IVirtualGroupCancelSwapIn // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupCancelSwapInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupCancelSwapIn)
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
		it.Event = new(IVirtualGroupCancelSwapIn)
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
func (it *IVirtualGroupCancelSwapInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupCancelSwapInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupCancelSwapIn represents a CancelSwapIn event raised by the IVirtualGroup contract.
type IVirtualGroupCancelSwapIn struct {
	StorageProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelSwapIn is a free log retrieval operation binding the contract event 0x47a1a3473759e2488a21c8e8211dd7d2c1beccd245308917baa38b92e88236a7.
//
// Solidity: event CancelSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterCancelSwapIn(opts *bind.FilterOpts, storageProvider []common.Address) (*IVirtualGroupCancelSwapInIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "CancelSwapIn", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupCancelSwapInIterator{contract: _IVirtualGroup.contract, event: "CancelSwapIn", logs: logs, sub: sub}, nil
}

// WatchCancelSwapIn is a free log subscription operation binding the contract event 0x47a1a3473759e2488a21c8e8211dd7d2c1beccd245308917baa38b92e88236a7.
//
// Solidity: event CancelSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchCancelSwapIn(opts *bind.WatchOpts, sink chan<- *IVirtualGroupCancelSwapIn, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "CancelSwapIn", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupCancelSwapIn)
				if err := _IVirtualGroup.contract.UnpackLog(event, "CancelSwapIn", log); err != nil {
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

// ParseCancelSwapIn is a log parse operation binding the contract event 0x47a1a3473759e2488a21c8e8211dd7d2c1beccd245308917baa38b92e88236a7.
//
// Solidity: event CancelSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseCancelSwapIn(log types.Log) (*IVirtualGroupCancelSwapIn, error) {
	event := new(IVirtualGroupCancelSwapIn)
	if err := _IVirtualGroup.contract.UnpackLog(event, "CancelSwapIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVirtualGroupCompleteSPExitIterator is returned from FilterCompleteSPExit and is used to iterate over the raw logs and unpacked data for CompleteSPExit events raised by the IVirtualGroup contract.
type IVirtualGroupCompleteSPExitIterator struct {
	Event *IVirtualGroupCompleteSPExit // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupCompleteSPExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupCompleteSPExit)
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
		it.Event = new(IVirtualGroupCompleteSPExit)
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
func (it *IVirtualGroupCompleteSPExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupCompleteSPExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupCompleteSPExit represents a CompleteSPExit event raised by the IVirtualGroup contract.
type IVirtualGroupCompleteSPExit struct {
	StorageProvider common.Address
	Operator        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompleteSPExit is a free log retrieval operation binding the contract event 0x1c77a3b4cf4d86d84d39911226dfeab9086419874ea60bf9a07314111138123b.
//
// Solidity: event CompleteSPExit(address indexed storageProvider, address indexed operator)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterCompleteSPExit(opts *bind.FilterOpts, storageProvider []common.Address, operator []common.Address) (*IVirtualGroupCompleteSPExitIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "CompleteSPExit", storageProviderRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupCompleteSPExitIterator{contract: _IVirtualGroup.contract, event: "CompleteSPExit", logs: logs, sub: sub}, nil
}

// WatchCompleteSPExit is a free log subscription operation binding the contract event 0x1c77a3b4cf4d86d84d39911226dfeab9086419874ea60bf9a07314111138123b.
//
// Solidity: event CompleteSPExit(address indexed storageProvider, address indexed operator)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchCompleteSPExit(opts *bind.WatchOpts, sink chan<- *IVirtualGroupCompleteSPExit, storageProvider []common.Address, operator []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "CompleteSPExit", storageProviderRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupCompleteSPExit)
				if err := _IVirtualGroup.contract.UnpackLog(event, "CompleteSPExit", log); err != nil {
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

// ParseCompleteSPExit is a log parse operation binding the contract event 0x1c77a3b4cf4d86d84d39911226dfeab9086419874ea60bf9a07314111138123b.
//
// Solidity: event CompleteSPExit(address indexed storageProvider, address indexed operator)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseCompleteSPExit(log types.Log) (*IVirtualGroupCompleteSPExit, error) {
	event := new(IVirtualGroupCompleteSPExit)
	if err := _IVirtualGroup.contract.UnpackLog(event, "CompleteSPExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVirtualGroupCompleteSwapInIterator is returned from FilterCompleteSwapIn and is used to iterate over the raw logs and unpacked data for CompleteSwapIn events raised by the IVirtualGroup contract.
type IVirtualGroupCompleteSwapInIterator struct {
	Event *IVirtualGroupCompleteSwapIn // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupCompleteSwapInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupCompleteSwapIn)
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
		it.Event = new(IVirtualGroupCompleteSwapIn)
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
func (it *IVirtualGroupCompleteSwapInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupCompleteSwapInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupCompleteSwapIn represents a CompleteSwapIn event raised by the IVirtualGroup contract.
type IVirtualGroupCompleteSwapIn struct {
	StorageProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCompleteSwapIn is a free log retrieval operation binding the contract event 0xa0b278230c83c530bcfd43f8569599a27da67b035f23b082e3bf172037e5f762.
//
// Solidity: event CompleteSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterCompleteSwapIn(opts *bind.FilterOpts, storageProvider []common.Address) (*IVirtualGroupCompleteSwapInIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "CompleteSwapIn", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupCompleteSwapInIterator{contract: _IVirtualGroup.contract, event: "CompleteSwapIn", logs: logs, sub: sub}, nil
}

// WatchCompleteSwapIn is a free log subscription operation binding the contract event 0xa0b278230c83c530bcfd43f8569599a27da67b035f23b082e3bf172037e5f762.
//
// Solidity: event CompleteSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchCompleteSwapIn(opts *bind.WatchOpts, sink chan<- *IVirtualGroupCompleteSwapIn, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "CompleteSwapIn", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupCompleteSwapIn)
				if err := _IVirtualGroup.contract.UnpackLog(event, "CompleteSwapIn", log); err != nil {
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

// ParseCompleteSwapIn is a log parse operation binding the contract event 0xa0b278230c83c530bcfd43f8569599a27da67b035f23b082e3bf172037e5f762.
//
// Solidity: event CompleteSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseCompleteSwapIn(log types.Log) (*IVirtualGroupCompleteSwapIn, error) {
	event := new(IVirtualGroupCompleteSwapIn)
	if err := _IVirtualGroup.contract.UnpackLog(event, "CompleteSwapIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVirtualGroupCreateGlobalVirtualGroupIterator is returned from FilterCreateGlobalVirtualGroup and is used to iterate over the raw logs and unpacked data for CreateGlobalVirtualGroup events raised by the IVirtualGroup contract.
type IVirtualGroupCreateGlobalVirtualGroupIterator struct {
	Event *IVirtualGroupCreateGlobalVirtualGroup // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupCreateGlobalVirtualGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupCreateGlobalVirtualGroup)
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
		it.Event = new(IVirtualGroupCreateGlobalVirtualGroup)
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
func (it *IVirtualGroupCreateGlobalVirtualGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupCreateGlobalVirtualGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupCreateGlobalVirtualGroup represents a CreateGlobalVirtualGroup event raised by the IVirtualGroup contract.
type IVirtualGroupCreateGlobalVirtualGroup struct {
	StorageProvider common.Address
	FamilyId        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateGlobalVirtualGroup is a free log retrieval operation binding the contract event 0x092288be07c9b3fdca11bd705993617a74c4d005de41b0ea7cf95e13a19d1947.
//
// Solidity: event CreateGlobalVirtualGroup(address indexed storageProvider, uint256 familyId)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterCreateGlobalVirtualGroup(opts *bind.FilterOpts, storageProvider []common.Address) (*IVirtualGroupCreateGlobalVirtualGroupIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "CreateGlobalVirtualGroup", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupCreateGlobalVirtualGroupIterator{contract: _IVirtualGroup.contract, event: "CreateGlobalVirtualGroup", logs: logs, sub: sub}, nil
}

// WatchCreateGlobalVirtualGroup is a free log subscription operation binding the contract event 0x092288be07c9b3fdca11bd705993617a74c4d005de41b0ea7cf95e13a19d1947.
//
// Solidity: event CreateGlobalVirtualGroup(address indexed storageProvider, uint256 familyId)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchCreateGlobalVirtualGroup(opts *bind.WatchOpts, sink chan<- *IVirtualGroupCreateGlobalVirtualGroup, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "CreateGlobalVirtualGroup", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupCreateGlobalVirtualGroup)
				if err := _IVirtualGroup.contract.UnpackLog(event, "CreateGlobalVirtualGroup", log); err != nil {
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

// ParseCreateGlobalVirtualGroup is a log parse operation binding the contract event 0x092288be07c9b3fdca11bd705993617a74c4d005de41b0ea7cf95e13a19d1947.
//
// Solidity: event CreateGlobalVirtualGroup(address indexed storageProvider, uint256 familyId)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseCreateGlobalVirtualGroup(log types.Log) (*IVirtualGroupCreateGlobalVirtualGroup, error) {
	event := new(IVirtualGroupCreateGlobalVirtualGroup)
	if err := _IVirtualGroup.contract.UnpackLog(event, "CreateGlobalVirtualGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVirtualGroupDeleteGlobalVirtualGroupIterator is returned from FilterDeleteGlobalVirtualGroup and is used to iterate over the raw logs and unpacked data for DeleteGlobalVirtualGroup events raised by the IVirtualGroup contract.
type IVirtualGroupDeleteGlobalVirtualGroupIterator struct {
	Event *IVirtualGroupDeleteGlobalVirtualGroup // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupDeleteGlobalVirtualGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupDeleteGlobalVirtualGroup)
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
		it.Event = new(IVirtualGroupDeleteGlobalVirtualGroup)
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
func (it *IVirtualGroupDeleteGlobalVirtualGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupDeleteGlobalVirtualGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupDeleteGlobalVirtualGroup represents a DeleteGlobalVirtualGroup event raised by the IVirtualGroup contract.
type IVirtualGroupDeleteGlobalVirtualGroup struct {
	StorageProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeleteGlobalVirtualGroup is a free log retrieval operation binding the contract event 0xf905269410b7831c9948f9fea7e5a62d9da6382d5bb44954cf67e71a90fd1527.
//
// Solidity: event DeleteGlobalVirtualGroup(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterDeleteGlobalVirtualGroup(opts *bind.FilterOpts, storageProvider []common.Address) (*IVirtualGroupDeleteGlobalVirtualGroupIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "DeleteGlobalVirtualGroup", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupDeleteGlobalVirtualGroupIterator{contract: _IVirtualGroup.contract, event: "DeleteGlobalVirtualGroup", logs: logs, sub: sub}, nil
}

// WatchDeleteGlobalVirtualGroup is a free log subscription operation binding the contract event 0xf905269410b7831c9948f9fea7e5a62d9da6382d5bb44954cf67e71a90fd1527.
//
// Solidity: event DeleteGlobalVirtualGroup(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchDeleteGlobalVirtualGroup(opts *bind.WatchOpts, sink chan<- *IVirtualGroupDeleteGlobalVirtualGroup, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "DeleteGlobalVirtualGroup", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupDeleteGlobalVirtualGroup)
				if err := _IVirtualGroup.contract.UnpackLog(event, "DeleteGlobalVirtualGroup", log); err != nil {
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

// ParseDeleteGlobalVirtualGroup is a log parse operation binding the contract event 0xf905269410b7831c9948f9fea7e5a62d9da6382d5bb44954cf67e71a90fd1527.
//
// Solidity: event DeleteGlobalVirtualGroup(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseDeleteGlobalVirtualGroup(log types.Log) (*IVirtualGroupDeleteGlobalVirtualGroup, error) {
	event := new(IVirtualGroupDeleteGlobalVirtualGroup)
	if err := _IVirtualGroup.contract.UnpackLog(event, "DeleteGlobalVirtualGroup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVirtualGroupDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IVirtualGroup contract.
type IVirtualGroupDepositIterator struct {
	Event *IVirtualGroupDeposit // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupDeposit)
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
		it.Event = new(IVirtualGroupDeposit)
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
func (it *IVirtualGroupDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupDeposit represents a Deposit event raised by the IVirtualGroup contract.
type IVirtualGroupDeposit struct {
	StorageProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x8ce0bd46ec50cf39f0d0ea8686a686eb226af5796dcda4231b26fb84b5ef1234.
//
// Solidity: event Deposit(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterDeposit(opts *bind.FilterOpts, storageProvider []common.Address) (*IVirtualGroupDepositIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "Deposit", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupDepositIterator{contract: _IVirtualGroup.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x8ce0bd46ec50cf39f0d0ea8686a686eb226af5796dcda4231b26fb84b5ef1234.
//
// Solidity: event Deposit(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IVirtualGroupDeposit, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "Deposit", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupDeposit)
				if err := _IVirtualGroup.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseDeposit(log types.Log) (*IVirtualGroupDeposit, error) {
	event := new(IVirtualGroupDeposit)
	if err := _IVirtualGroup.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVirtualGroupReserveSwapInIterator is returned from FilterReserveSwapIn and is used to iterate over the raw logs and unpacked data for ReserveSwapIn events raised by the IVirtualGroup contract.
type IVirtualGroupReserveSwapInIterator struct {
	Event *IVirtualGroupReserveSwapIn // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupReserveSwapInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupReserveSwapIn)
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
		it.Event = new(IVirtualGroupReserveSwapIn)
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
func (it *IVirtualGroupReserveSwapInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupReserveSwapInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupReserveSwapIn represents a ReserveSwapIn event raised by the IVirtualGroup contract.
type IVirtualGroupReserveSwapIn struct {
	StorageProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReserveSwapIn is a free log retrieval operation binding the contract event 0x4dde188629a6fe412c5d5ad81c5fcda63e7ec0aa07df7fa41a1b8406e85f7317.
//
// Solidity: event ReserveSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterReserveSwapIn(opts *bind.FilterOpts, storageProvider []common.Address) (*IVirtualGroupReserveSwapInIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "ReserveSwapIn", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupReserveSwapInIterator{contract: _IVirtualGroup.contract, event: "ReserveSwapIn", logs: logs, sub: sub}, nil
}

// WatchReserveSwapIn is a free log subscription operation binding the contract event 0x4dde188629a6fe412c5d5ad81c5fcda63e7ec0aa07df7fa41a1b8406e85f7317.
//
// Solidity: event ReserveSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchReserveSwapIn(opts *bind.WatchOpts, sink chan<- *IVirtualGroupReserveSwapIn, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "ReserveSwapIn", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupReserveSwapIn)
				if err := _IVirtualGroup.contract.UnpackLog(event, "ReserveSwapIn", log); err != nil {
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

// ParseReserveSwapIn is a log parse operation binding the contract event 0x4dde188629a6fe412c5d5ad81c5fcda63e7ec0aa07df7fa41a1b8406e85f7317.
//
// Solidity: event ReserveSwapIn(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseReserveSwapIn(log types.Log) (*IVirtualGroupReserveSwapIn, error) {
	event := new(IVirtualGroupReserveSwapIn)
	if err := _IVirtualGroup.contract.UnpackLog(event, "ReserveSwapIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVirtualGroupSPExitIterator is returned from FilterSPExit and is used to iterate over the raw logs and unpacked data for SPExit events raised by the IVirtualGroup contract.
type IVirtualGroupSPExitIterator struct {
	Event *IVirtualGroupSPExit // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupSPExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupSPExit)
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
		it.Event = new(IVirtualGroupSPExit)
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
func (it *IVirtualGroupSPExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupSPExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupSPExit represents a SPExit event raised by the IVirtualGroup contract.
type IVirtualGroupSPExit struct {
	StorageProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSPExit is a free log retrieval operation binding the contract event 0x413fcbc6e0e0f8308a3f54b8522f3301c66bc3c86922d32a7f169fffdbb85724.
//
// Solidity: event SPExit(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterSPExit(opts *bind.FilterOpts, storageProvider []common.Address) (*IVirtualGroupSPExitIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "SPExit", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupSPExitIterator{contract: _IVirtualGroup.contract, event: "SPExit", logs: logs, sub: sub}, nil
}

// WatchSPExit is a free log subscription operation binding the contract event 0x413fcbc6e0e0f8308a3f54b8522f3301c66bc3c86922d32a7f169fffdbb85724.
//
// Solidity: event SPExit(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchSPExit(opts *bind.WatchOpts, sink chan<- *IVirtualGroupSPExit, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "SPExit", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupSPExit)
				if err := _IVirtualGroup.contract.UnpackLog(event, "SPExit", log); err != nil {
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

// ParseSPExit is a log parse operation binding the contract event 0x413fcbc6e0e0f8308a3f54b8522f3301c66bc3c86922d32a7f169fffdbb85724.
//
// Solidity: event SPExit(address indexed storageProvider)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseSPExit(log types.Log) (*IVirtualGroupSPExit, error) {
	event := new(IVirtualGroupSPExit)
	if err := _IVirtualGroup.contract.UnpackLog(event, "SPExit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVirtualGroupSwapOutIterator is returned from FilterSwapOut and is used to iterate over the raw logs and unpacked data for SwapOut events raised by the IVirtualGroup contract.
type IVirtualGroupSwapOutIterator struct {
	Event *IVirtualGroupSwapOut // Event containing the contract specifics and raw log

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
func (it *IVirtualGroupSwapOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVirtualGroupSwapOut)
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
		it.Event = new(IVirtualGroupSwapOut)
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
func (it *IVirtualGroupSwapOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVirtualGroupSwapOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVirtualGroupSwapOut represents a SwapOut event raised by the IVirtualGroup contract.
type IVirtualGroupSwapOut struct {
	StorageProvider common.Address
	FamilyId        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSwapOut is a free log retrieval operation binding the contract event 0x096ca2225a0de1da75baf17d9f77a22d2d99a5e11f397174df2f654acf1e60fc.
//
// Solidity: event SwapOut(address indexed storageProvider, uint256 familyId)
func (_IVirtualGroup *IVirtualGroupFilterer) FilterSwapOut(opts *bind.FilterOpts, storageProvider []common.Address) (*IVirtualGroupSwapOutIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.FilterLogs(opts, "SwapOut", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IVirtualGroupSwapOutIterator{contract: _IVirtualGroup.contract, event: "SwapOut", logs: logs, sub: sub}, nil
}

// WatchSwapOut is a free log subscription operation binding the contract event 0x096ca2225a0de1da75baf17d9f77a22d2d99a5e11f397174df2f654acf1e60fc.
//
// Solidity: event SwapOut(address indexed storageProvider, uint256 familyId)
func (_IVirtualGroup *IVirtualGroupFilterer) WatchSwapOut(opts *bind.WatchOpts, sink chan<- *IVirtualGroupSwapOut, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IVirtualGroup.contract.WatchLogs(opts, "SwapOut", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVirtualGroupSwapOut)
				if err := _IVirtualGroup.contract.UnpackLog(event, "SwapOut", log); err != nil {
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

// ParseSwapOut is a log parse operation binding the contract event 0x096ca2225a0de1da75baf17d9f77a22d2d99a5e11f397174df2f654acf1e60fc.
//
// Solidity: event SwapOut(address indexed storageProvider, uint256 familyId)
func (_IVirtualGroup *IVirtualGroupFilterer) ParseSwapOut(log types.Log) (*IVirtualGroupSwapOut, error) {
	event := new(IVirtualGroupSwapOut)
	if err := _IVirtualGroup.contract.UnpackLog(event, "SwapOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
