// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permission

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

// Params is an auto generated low-level Go binding around an user-defined struct.
type Params struct {
	MaximumStatementsNum                  uint64
	MaximumGroupNum                       uint64
	MaximumRemoveExpiredPoliciesIteration uint64
}

// IPermissionMetaData contains all meta data concerning the IPermission contract.
var IPermissionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"UpdateParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"maximumStatementsNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maximumGroupNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maximumRemoveExpiredPoliciesIteration\",\"type\":\"uint64\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"authority\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"maximumStatementsNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maximumGroupNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maximumRemoveExpiredPoliciesIteration\",\"type\":\"uint64\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"updateParams\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IPermissionABI is the input ABI used to generate the binding from.
// Deprecated: Use IPermissionMetaData.ABI instead.
var IPermissionABI = IPermissionMetaData.ABI

// IPermission is an auto generated Go binding around an Ethereum contract.
type IPermission struct {
	IPermissionCaller     // Read-only binding to the contract
	IPermissionTransactor // Write-only binding to the contract
	IPermissionFilterer   // Log filterer for contract events
}

// IPermissionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPermissionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPermissionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPermissionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPermissionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPermissionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPermissionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPermissionSession struct {
	Contract     *IPermission      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPermissionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPermissionCallerSession struct {
	Contract *IPermissionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IPermissionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPermissionTransactorSession struct {
	Contract     *IPermissionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IPermissionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPermissionRaw struct {
	Contract *IPermission // Generic contract binding to access the raw methods on
}

// IPermissionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPermissionCallerRaw struct {
	Contract *IPermissionCaller // Generic read-only contract binding to access the raw methods on
}

// IPermissionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPermissionTransactorRaw struct {
	Contract *IPermissionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPermission creates a new instance of IPermission, bound to a specific deployed contract.
func NewIPermission(address common.Address, backend bind.ContractBackend) (*IPermission, error) {
	contract, err := bindIPermission(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPermission{IPermissionCaller: IPermissionCaller{contract: contract}, IPermissionTransactor: IPermissionTransactor{contract: contract}, IPermissionFilterer: IPermissionFilterer{contract: contract}}, nil
}

// NewIPermissionCaller creates a new read-only instance of IPermission, bound to a specific deployed contract.
func NewIPermissionCaller(address common.Address, caller bind.ContractCaller) (*IPermissionCaller, error) {
	contract, err := bindIPermission(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPermissionCaller{contract: contract}, nil
}

// NewIPermissionTransactor creates a new write-only instance of IPermission, bound to a specific deployed contract.
func NewIPermissionTransactor(address common.Address, transactor bind.ContractTransactor) (*IPermissionTransactor, error) {
	contract, err := bindIPermission(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPermissionTransactor{contract: contract}, nil
}

// NewIPermissionFilterer creates a new log filterer instance of IPermission, bound to a specific deployed contract.
func NewIPermissionFilterer(address common.Address, filterer bind.ContractFilterer) (*IPermissionFilterer, error) {
	contract, err := bindIPermission(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPermissionFilterer{contract: contract}, nil
}

// bindIPermission binds a generic wrapper to an already deployed contract.
func bindIPermission(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPermissionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPermission *IPermissionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPermission.Contract.IPermissionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPermission *IPermissionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPermission.Contract.IPermissionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPermission *IPermissionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPermission.Contract.IPermissionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPermission *IPermissionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPermission.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPermission *IPermissionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPermission.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPermission *IPermissionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPermission.Contract.contract.Transact(opts, method, params...)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint64,uint64,uint64) params)
func (_IPermission *IPermissionCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _IPermission.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint64,uint64,uint64) params)
func (_IPermission *IPermissionSession) Params() (Params, error) {
	return _IPermission.Contract.Params(&_IPermission.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns((uint64,uint64,uint64) params)
func (_IPermission *IPermissionCallerSession) Params() (Params, error) {
	return _IPermission.Contract.Params(&_IPermission.CallOpts)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xcc63f332.
//
// Solidity: function updateParams(string authority, (uint64,uint64,uint64) params) returns(bool success)
func (_IPermission *IPermissionTransactor) UpdateParams(opts *bind.TransactOpts, authority string, params Params) (*types.Transaction, error) {
	return _IPermission.contract.Transact(opts, "updateParams", authority, params)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xcc63f332.
//
// Solidity: function updateParams(string authority, (uint64,uint64,uint64) params) returns(bool success)
func (_IPermission *IPermissionSession) UpdateParams(authority string, params Params) (*types.Transaction, error) {
	return _IPermission.Contract.UpdateParams(&_IPermission.TransactOpts, authority, params)
}

// UpdateParams is a paid mutator transaction binding the contract method 0xcc63f332.
//
// Solidity: function updateParams(string authority, (uint64,uint64,uint64) params) returns(bool success)
func (_IPermission *IPermissionTransactorSession) UpdateParams(authority string, params Params) (*types.Transaction, error) {
	return _IPermission.Contract.UpdateParams(&_IPermission.TransactOpts, authority, params)
}

// IPermissionUpdateParamsIterator is returned from FilterUpdateParams and is used to iterate over the raw logs and unpacked data for UpdateParams events raised by the IPermission contract.
type IPermissionUpdateParamsIterator struct {
	Event *IPermissionUpdateParams // Event containing the contract specifics and raw log

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
func (it *IPermissionUpdateParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPermissionUpdateParams)
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
		it.Event = new(IPermissionUpdateParams)
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
func (it *IPermissionUpdateParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPermissionUpdateParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPermissionUpdateParams represents a UpdateParams event raised by the IPermission contract.
type IPermissionUpdateParams struct {
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateParams is a free log retrieval operation binding the contract event 0xdb2e743561971fc97db45300fcd12c1e9e20b20d485aa6f0259c3374f4a4dafd.
//
// Solidity: event UpdateParams(address indexed creator)
func (_IPermission *IPermissionFilterer) FilterUpdateParams(opts *bind.FilterOpts, creator []common.Address) (*IPermissionUpdateParamsIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IPermission.contract.FilterLogs(opts, "UpdateParams", creatorRule)
	if err != nil {
		return nil, err
	}
	return &IPermissionUpdateParamsIterator{contract: _IPermission.contract, event: "UpdateParams", logs: logs, sub: sub}, nil
}

// WatchUpdateParams is a free log subscription operation binding the contract event 0xdb2e743561971fc97db45300fcd12c1e9e20b20d485aa6f0259c3374f4a4dafd.
//
// Solidity: event UpdateParams(address indexed creator)
func (_IPermission *IPermissionFilterer) WatchUpdateParams(opts *bind.WatchOpts, sink chan<- *IPermissionUpdateParams, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _IPermission.contract.WatchLogs(opts, "UpdateParams", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPermissionUpdateParams)
				if err := _IPermission.contract.UnpackLog(event, "UpdateParams", log); err != nil {
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
func (_IPermission *IPermissionFilterer) ParseUpdateParams(log types.Log) (*IPermissionUpdateParams, error) {
	event := new(IPermissionUpdateParams)
	if err := _IPermission.contract.UnpackLog(event, "UpdateParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
