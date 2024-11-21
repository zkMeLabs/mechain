// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package authz

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

// GrantAuthorization is an auto generated low-level Go binding around an user-defined struct.
type GrantAuthorization struct {
	Granter       common.Address
	Grantee       common.Address
	Authorization string
	Expiration    int64
}

// GrantData is an auto generated low-level Go binding around an user-defined struct.
type GrantData struct {
	Authorization string
	Expiration    int64
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

// IAuthzMetaData contains all meta data concerning the IAuthz contract.
var IAuthzMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"}],\"name\":\"Exec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"authzType\",\"type\":\"string\"}],\"name\":\"Grant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msgTypeUrl\",\"type\":\"string\"}],\"name\":\"Revoke\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"msgs\",\"type\":\"string[]\"}],\"name\":\"exec\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"authzType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"authorization\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"limit\",\"type\":\"tuple[]\"},{\"internalType\":\"int64\",\"name\":\"expiration\",\"type\":\"int64\"}],\"name\":\"grant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"granteeGrants\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"authorization\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"expiration\",\"type\":\"int64\"}],\"internalType\":\"structGrantAuthorization[]\",\"name\":\"grants\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"granterGrants\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"authorization\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"expiration\",\"type\":\"int64\"}],\"internalType\":\"structGrantAuthorization[]\",\"name\":\"grants\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"granter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"msgTypeUrl\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"grants\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"authorization\",\"type\":\"string\"},{\"internalType\":\"int64\",\"name\":\"expiration\",\"type\":\"int64\"}],\"internalType\":\"structGrantData[]\",\"name\":\"grants\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"grantee\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"msgTypeUrl\",\"type\":\"string\"}],\"name\":\"revoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAuthzABI is the input ABI used to generate the binding from.
// Deprecated: Use IAuthzMetaData.ABI instead.
var IAuthzABI = IAuthzMetaData.ABI

// IAuthz is an auto generated Go binding around an Ethereum contract.
type IAuthz struct {
	IAuthzCaller     // Read-only binding to the contract
	IAuthzTransactor // Write-only binding to the contract
	IAuthzFilterer   // Log filterer for contract events
}

// IAuthzCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuthzCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthzTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuthzTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthzFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuthzFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuthzSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuthzSession struct {
	Contract     *IAuthz           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthzCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuthzCallerSession struct {
	Contract *IAuthzCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IAuthzTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuthzTransactorSession struct {
	Contract     *IAuthzTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuthzRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuthzRaw struct {
	Contract *IAuthz // Generic contract binding to access the raw methods on
}

// IAuthzCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuthzCallerRaw struct {
	Contract *IAuthzCaller // Generic read-only contract binding to access the raw methods on
}

// IAuthzTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuthzTransactorRaw struct {
	Contract *IAuthzTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuthz creates a new instance of IAuthz, bound to a specific deployed contract.
func NewIAuthz(address common.Address, backend bind.ContractBackend) (*IAuthz, error) {
	contract, err := bindIAuthz(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuthz{IAuthzCaller: IAuthzCaller{contract: contract}, IAuthzTransactor: IAuthzTransactor{contract: contract}, IAuthzFilterer: IAuthzFilterer{contract: contract}}, nil
}

// NewIAuthzCaller creates a new read-only instance of IAuthz, bound to a specific deployed contract.
func NewIAuthzCaller(address common.Address, caller bind.ContractCaller) (*IAuthzCaller, error) {
	contract, err := bindIAuthz(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthzCaller{contract: contract}, nil
}

// NewIAuthzTransactor creates a new write-only instance of IAuthz, bound to a specific deployed contract.
func NewIAuthzTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuthzTransactor, error) {
	contract, err := bindIAuthz(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuthzTransactor{contract: contract}, nil
}

// NewIAuthzFilterer creates a new log filterer instance of IAuthz, bound to a specific deployed contract.
func NewIAuthzFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuthzFilterer, error) {
	contract, err := bindIAuthz(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuthzFilterer{contract: contract}, nil
}

// bindIAuthz binds a generic wrapper to an already deployed contract.
func bindIAuthz(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAuthzMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuthz *IAuthzRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuthz.Contract.IAuthzCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuthz *IAuthzRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuthz.Contract.IAuthzTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuthz *IAuthzRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuthz.Contract.IAuthzTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuthz *IAuthzCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuthz.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuthz *IAuthzTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuthz.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuthz *IAuthzTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuthz.Contract.contract.Transact(opts, method, params...)
}

// GranteeGrants is a free data retrieval call binding the contract method 0xa531881a.
//
// Solidity: function granteeGrants(address grantee, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzCaller) GranteeGrants(opts *bind.CallOpts, grantee common.Address, pagination PageRequest) (struct {
	Grants       []GrantAuthorization
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IAuthz.contract.Call(opts, &out, "granteeGrants", grantee, pagination)

	outstruct := new(struct {
		Grants       []GrantAuthorization
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Grants = *abi.ConvertType(out[0], new([]GrantAuthorization)).(*[]GrantAuthorization)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// GranteeGrants is a free data retrieval call binding the contract method 0xa531881a.
//
// Solidity: function granteeGrants(address grantee, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzSession) GranteeGrants(grantee common.Address, pagination PageRequest) (struct {
	Grants       []GrantAuthorization
	PageResponse PageResponse
}, error) {
	return _IAuthz.Contract.GranteeGrants(&_IAuthz.CallOpts, grantee, pagination)
}

// GranteeGrants is a free data retrieval call binding the contract method 0xa531881a.
//
// Solidity: function granteeGrants(address grantee, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzCallerSession) GranteeGrants(grantee common.Address, pagination PageRequest) (struct {
	Grants       []GrantAuthorization
	PageResponse PageResponse
}, error) {
	return _IAuthz.Contract.GranteeGrants(&_IAuthz.CallOpts, grantee, pagination)
}

// GranterGrants is a free data retrieval call binding the contract method 0x9672a536.
//
// Solidity: function granterGrants(address granter, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzCaller) GranterGrants(opts *bind.CallOpts, granter common.Address, pagination PageRequest) (struct {
	Grants       []GrantAuthorization
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IAuthz.contract.Call(opts, &out, "granterGrants", granter, pagination)

	outstruct := new(struct {
		Grants       []GrantAuthorization
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Grants = *abi.ConvertType(out[0], new([]GrantAuthorization)).(*[]GrantAuthorization)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// GranterGrants is a free data retrieval call binding the contract method 0x9672a536.
//
// Solidity: function granterGrants(address granter, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzSession) GranterGrants(granter common.Address, pagination PageRequest) (struct {
	Grants       []GrantAuthorization
	PageResponse PageResponse
}, error) {
	return _IAuthz.Contract.GranterGrants(&_IAuthz.CallOpts, granter, pagination)
}

// GranterGrants is a free data retrieval call binding the contract method 0x9672a536.
//
// Solidity: function granterGrants(address granter, (bytes,uint64,uint64,bool,bool) pagination) view returns((address,address,string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzCallerSession) GranterGrants(granter common.Address, pagination PageRequest) (struct {
	Grants       []GrantAuthorization
	PageResponse PageResponse
}, error) {
	return _IAuthz.Contract.GranterGrants(&_IAuthz.CallOpts, granter, pagination)
}

// Grants is a free data retrieval call binding the contract method 0x60120384.
//
// Solidity: function grants(address granter, address grantee, string msgTypeUrl, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzCaller) Grants(opts *bind.CallOpts, granter common.Address, grantee common.Address, msgTypeUrl string, pagination PageRequest) (struct {
	Grants       []GrantData
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IAuthz.contract.Call(opts, &out, "grants", granter, grantee, msgTypeUrl, pagination)

	outstruct := new(struct {
		Grants       []GrantData
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Grants = *abi.ConvertType(out[0], new([]GrantData)).(*[]GrantData)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// Grants is a free data retrieval call binding the contract method 0x60120384.
//
// Solidity: function grants(address granter, address grantee, string msgTypeUrl, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzSession) Grants(granter common.Address, grantee common.Address, msgTypeUrl string, pagination PageRequest) (struct {
	Grants       []GrantData
	PageResponse PageResponse
}, error) {
	return _IAuthz.Contract.Grants(&_IAuthz.CallOpts, granter, grantee, msgTypeUrl, pagination)
}

// Grants is a free data retrieval call binding the contract method 0x60120384.
//
// Solidity: function grants(address granter, address grantee, string msgTypeUrl, (bytes,uint64,uint64,bool,bool) pagination) view returns((string,int64)[] grants, (bytes,uint64) pageResponse)
func (_IAuthz *IAuthzCallerSession) Grants(granter common.Address, grantee common.Address, msgTypeUrl string, pagination PageRequest) (struct {
	Grants       []GrantData
	PageResponse PageResponse
}, error) {
	return _IAuthz.Contract.Grants(&_IAuthz.CallOpts, granter, grantee, msgTypeUrl, pagination)
}

// Exec is a paid mutator transaction binding the contract method 0x47d4efd6.
//
// Solidity: function exec(string[] msgs) returns(bool success)
func (_IAuthz *IAuthzTransactor) Exec(opts *bind.TransactOpts, msgs []string) (*types.Transaction, error) {
	return _IAuthz.contract.Transact(opts, "exec", msgs)
}

// Exec is a paid mutator transaction binding the contract method 0x47d4efd6.
//
// Solidity: function exec(string[] msgs) returns(bool success)
func (_IAuthz *IAuthzSession) Exec(msgs []string) (*types.Transaction, error) {
	return _IAuthz.Contract.Exec(&_IAuthz.TransactOpts, msgs)
}

// Exec is a paid mutator transaction binding the contract method 0x47d4efd6.
//
// Solidity: function exec(string[] msgs) returns(bool success)
func (_IAuthz *IAuthzTransactorSession) Exec(msgs []string) (*types.Transaction, error) {
	return _IAuthz.Contract.Exec(&_IAuthz.TransactOpts, msgs)
}

// Grant is a paid mutator transaction binding the contract method 0xdf508fbb.
//
// Solidity: function grant(address grantee, string authzType, string authorization, (string,uint256)[] limit, int64 expiration) returns(bool success)
func (_IAuthz *IAuthzTransactor) Grant(opts *bind.TransactOpts, grantee common.Address, authzType string, authorization string, limit []Coin, expiration int64) (*types.Transaction, error) {
	return _IAuthz.contract.Transact(opts, "grant", grantee, authzType, authorization, limit, expiration)
}

// Grant is a paid mutator transaction binding the contract method 0xdf508fbb.
//
// Solidity: function grant(address grantee, string authzType, string authorization, (string,uint256)[] limit, int64 expiration) returns(bool success)
func (_IAuthz *IAuthzSession) Grant(grantee common.Address, authzType string, authorization string, limit []Coin, expiration int64) (*types.Transaction, error) {
	return _IAuthz.Contract.Grant(&_IAuthz.TransactOpts, grantee, authzType, authorization, limit, expiration)
}

// Grant is a paid mutator transaction binding the contract method 0xdf508fbb.
//
// Solidity: function grant(address grantee, string authzType, string authorization, (string,uint256)[] limit, int64 expiration) returns(bool success)
func (_IAuthz *IAuthzTransactorSession) Grant(grantee common.Address, authzType string, authorization string, limit []Coin, expiration int64) (*types.Transaction, error) {
	return _IAuthz.Contract.Grant(&_IAuthz.TransactOpts, grantee, authzType, authorization, limit, expiration)
}

// Revoke is a paid mutator transaction binding the contract method 0xafd0224b.
//
// Solidity: function revoke(address grantee, string msgTypeUrl) returns(bool success)
func (_IAuthz *IAuthzTransactor) Revoke(opts *bind.TransactOpts, grantee common.Address, msgTypeUrl string) (*types.Transaction, error) {
	return _IAuthz.contract.Transact(opts, "revoke", grantee, msgTypeUrl)
}

// Revoke is a paid mutator transaction binding the contract method 0xafd0224b.
//
// Solidity: function revoke(address grantee, string msgTypeUrl) returns(bool success)
func (_IAuthz *IAuthzSession) Revoke(grantee common.Address, msgTypeUrl string) (*types.Transaction, error) {
	return _IAuthz.Contract.Revoke(&_IAuthz.TransactOpts, grantee, msgTypeUrl)
}

// Revoke is a paid mutator transaction binding the contract method 0xafd0224b.
//
// Solidity: function revoke(address grantee, string msgTypeUrl) returns(bool success)
func (_IAuthz *IAuthzTransactorSession) Revoke(grantee common.Address, msgTypeUrl string) (*types.Transaction, error) {
	return _IAuthz.Contract.Revoke(&_IAuthz.TransactOpts, grantee, msgTypeUrl)
}

// IAuthzExecIterator is returned from FilterExec and is used to iterate over the raw logs and unpacked data for Exec events raised by the IAuthz contract.
type IAuthzExecIterator struct {
	Event *IAuthzExec // Event containing the contract specifics and raw log

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
func (it *IAuthzExecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuthzExec)
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
		it.Event = new(IAuthzExec)
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
func (it *IAuthzExecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuthzExecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuthzExec represents a Exec event raised by the IAuthz contract.
type IAuthzExec struct {
	Grantee common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExec is a free log retrieval operation binding the contract event 0xd074dc2801f2a4136a42e9fae80ca2c91228872a7a4996e6b47e10879559c0bf.
//
// Solidity: event Exec(address indexed grantee)
func (_IAuthz *IAuthzFilterer) FilterExec(opts *bind.FilterOpts, grantee []common.Address) (*IAuthzExecIterator, error) {

	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _IAuthz.contract.FilterLogs(opts, "Exec", granteeRule)
	if err != nil {
		return nil, err
	}
	return &IAuthzExecIterator{contract: _IAuthz.contract, event: "Exec", logs: logs, sub: sub}, nil
}

// WatchExec is a free log subscription operation binding the contract event 0xd074dc2801f2a4136a42e9fae80ca2c91228872a7a4996e6b47e10879559c0bf.
//
// Solidity: event Exec(address indexed grantee)
func (_IAuthz *IAuthzFilterer) WatchExec(opts *bind.WatchOpts, sink chan<- *IAuthzExec, grantee []common.Address) (event.Subscription, error) {

	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _IAuthz.contract.WatchLogs(opts, "Exec", granteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuthzExec)
				if err := _IAuthz.contract.UnpackLog(event, "Exec", log); err != nil {
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

// ParseExec is a log parse operation binding the contract event 0xd074dc2801f2a4136a42e9fae80ca2c91228872a7a4996e6b47e10879559c0bf.
//
// Solidity: event Exec(address indexed grantee)
func (_IAuthz *IAuthzFilterer) ParseExec(log types.Log) (*IAuthzExec, error) {
	event := new(IAuthzExec)
	if err := _IAuthz.contract.UnpackLog(event, "Exec", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuthzGrantIterator is returned from FilterGrant and is used to iterate over the raw logs and unpacked data for Grant events raised by the IAuthz contract.
type IAuthzGrantIterator struct {
	Event *IAuthzGrant // Event containing the contract specifics and raw log

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
func (it *IAuthzGrantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuthzGrant)
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
		it.Event = new(IAuthzGrant)
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
func (it *IAuthzGrantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuthzGrantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuthzGrant represents a Grant event raised by the IAuthz contract.
type IAuthzGrant struct {
	Granter   common.Address
	Grantee   common.Address
	AuthzType string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGrant is a free log retrieval operation binding the contract event 0xbfab9413dcbf9e8d938e2cc64562caeeb065bece6869742e10dc389036faa79d.
//
// Solidity: event Grant(address indexed granter, address indexed grantee, string authzType)
func (_IAuthz *IAuthzFilterer) FilterGrant(opts *bind.FilterOpts, granter []common.Address, grantee []common.Address) (*IAuthzGrantIterator, error) {

	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _IAuthz.contract.FilterLogs(opts, "Grant", granterRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return &IAuthzGrantIterator{contract: _IAuthz.contract, event: "Grant", logs: logs, sub: sub}, nil
}

// WatchGrant is a free log subscription operation binding the contract event 0xbfab9413dcbf9e8d938e2cc64562caeeb065bece6869742e10dc389036faa79d.
//
// Solidity: event Grant(address indexed granter, address indexed grantee, string authzType)
func (_IAuthz *IAuthzFilterer) WatchGrant(opts *bind.WatchOpts, sink chan<- *IAuthzGrant, granter []common.Address, grantee []common.Address) (event.Subscription, error) {

	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _IAuthz.contract.WatchLogs(opts, "Grant", granterRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuthzGrant)
				if err := _IAuthz.contract.UnpackLog(event, "Grant", log); err != nil {
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

// ParseGrant is a log parse operation binding the contract event 0xbfab9413dcbf9e8d938e2cc64562caeeb065bece6869742e10dc389036faa79d.
//
// Solidity: event Grant(address indexed granter, address indexed grantee, string authzType)
func (_IAuthz *IAuthzFilterer) ParseGrant(log types.Log) (*IAuthzGrant, error) {
	event := new(IAuthzGrant)
	if err := _IAuthz.contract.UnpackLog(event, "Grant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuthzRevokeIterator is returned from FilterRevoke and is used to iterate over the raw logs and unpacked data for Revoke events raised by the IAuthz contract.
type IAuthzRevokeIterator struct {
	Event *IAuthzRevoke // Event containing the contract specifics and raw log

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
func (it *IAuthzRevokeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuthzRevoke)
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
		it.Event = new(IAuthzRevoke)
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
func (it *IAuthzRevokeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuthzRevokeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuthzRevoke represents a Revoke event raised by the IAuthz contract.
type IAuthzRevoke struct {
	Granter    common.Address
	Grantee    common.Address
	MsgTypeUrl string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevoke is a free log retrieval operation binding the contract event 0x89edca5e39ec72c8be42f61c849867ad405ab6f86a51818b624504b0c3f5f5b2.
//
// Solidity: event Revoke(address indexed granter, address indexed grantee, string msgTypeUrl)
func (_IAuthz *IAuthzFilterer) FilterRevoke(opts *bind.FilterOpts, granter []common.Address, grantee []common.Address) (*IAuthzRevokeIterator, error) {

	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _IAuthz.contract.FilterLogs(opts, "Revoke", granterRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return &IAuthzRevokeIterator{contract: _IAuthz.contract, event: "Revoke", logs: logs, sub: sub}, nil
}

// WatchRevoke is a free log subscription operation binding the contract event 0x89edca5e39ec72c8be42f61c849867ad405ab6f86a51818b624504b0c3f5f5b2.
//
// Solidity: event Revoke(address indexed granter, address indexed grantee, string msgTypeUrl)
func (_IAuthz *IAuthzFilterer) WatchRevoke(opts *bind.WatchOpts, sink chan<- *IAuthzRevoke, granter []common.Address, grantee []common.Address) (event.Subscription, error) {

	var granterRule []interface{}
	for _, granterItem := range granter {
		granterRule = append(granterRule, granterItem)
	}
	var granteeRule []interface{}
	for _, granteeItem := range grantee {
		granteeRule = append(granteeRule, granteeItem)
	}

	logs, sub, err := _IAuthz.contract.WatchLogs(opts, "Revoke", granterRule, granteeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuthzRevoke)
				if err := _IAuthz.contract.UnpackLog(event, "Revoke", log); err != nil {
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

// ParseRevoke is a log parse operation binding the contract event 0x89edca5e39ec72c8be42f61c849867ad405ab6f86a51818b624504b0c3f5f5b2.
//
// Solidity: event Revoke(address indexed granter, address indexed grantee, string msgTypeUrl)
func (_IAuthz *IAuthzFilterer) ParseRevoke(log types.Log) (*IAuthzRevoke, error) {
	event := new(IAuthzRevoke)
	if err := _IAuthz.contract.UnpackLog(event, "Revoke", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
