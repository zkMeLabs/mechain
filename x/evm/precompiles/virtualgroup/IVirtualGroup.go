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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"familyId\",\"type\":\"uint256\"}],\"name\":\"CreateGlobalVirtualGroup\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"familyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"secondarySpIds\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"deposit\",\"type\":\"tuple\"}],\"name\":\"createGlobalVirtualGroup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"globalVirtualGroupFamilies\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"primarySpId\",\"type\":\"uint32\"},{\"internalType\":\"uint32[]\",\"name\":\"globalVirtualGroupIds\",\"type\":\"uint32[]\"},{\"internalType\":\"address\",\"name\":\"virtualPaymentAddress\",\"type\":\"address\"}],\"internalType\":\"structGlobalVirtualGroupFamily[]\",\"name\":\"gvgFamilies\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
