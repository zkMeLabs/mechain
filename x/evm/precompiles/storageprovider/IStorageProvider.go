// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storageprovider

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

// Description is an auto generated low-level Go binding around an user-defined struct.
type Description struct {
	Moniker         string
	Identity        string
	Website         string
	SecurityContact string
	Details         string
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

// SpStoragePrice is an auto generated low-level Go binding around an user-defined struct.
type SpStoragePrice struct {
	SpId          uint32
	UpdateTimeSec *big.Int
	ReadPrice     *big.Int
	FreeReadQuota uint64
	StorePrice    *big.Int
}

// StorageProvider is an auto generated low-level Go binding around an user-defined struct.
type StorageProvider struct {
	Id                 uint32
	OperatorAddress    string
	FundingAddress     string
	SealAddress        string
	ApprovalAddress    string
	GcAddress          string
	MaintenanceAddress string
	TotalDeposit       *big.Int
	Status             uint8
	Endpoint           string
	Description        Description
	BlsKey             string
}

// IStorageProviderMetaData contains all meta data concerning the IStorageProvider contract.
var IStorageProviderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"storageProvider\",\"type\":\"address\"}],\"name\":\"UpdateSPPrice\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"}],\"name\":\"storageProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"operator_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funding_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"seal_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"approval_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gc_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"maintenance_address\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"total_deposit\",\"type\":\"uint256\"},{\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"security_contact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"bls_key\",\"type\":\"string\"}],\"internalType\":\"structStorageProvider\",\"name\":\"storageProvider\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"storageProviderByOperatorAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"operator_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funding_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"seal_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"approval_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gc_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"maintenance_address\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"total_deposit\",\"type\":\"uint256\"},{\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"security_contact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"bls_key\",\"type\":\"string\"}],\"internalType\":\"structStorageProvider\",\"name\":\"storageProvider\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddress\",\"type\":\"address\"}],\"name\":\"storageProviderPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"sp_id\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"update_time_sec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"read_price\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"free_read_quota\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"store_price\",\"type\":\"uint256\"}],\"internalType\":\"structSpStoragePrice\",\"name\":\"spStoragePrice\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pagination\",\"type\":\"tuple\"}],\"name\":\"storageProviders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"id\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"operator_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"funding_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"seal_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"approval_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"gc_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"maintenance_address\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"total_deposit\",\"type\":\"uint256\"},{\"internalType\":\"enumStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"security_contact\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"internalType\":\"structDescription\",\"name\":\"description\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"bls_key\",\"type\":\"string\"}],\"internalType\":\"structStorageProvider[]\",\"name\":\"storageProviders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"readPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"freeReadQuota\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"storePrice\",\"type\":\"uint256\"}],\"name\":\"updateSPPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IStorageProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use IStorageProviderMetaData.ABI instead.
var IStorageProviderABI = IStorageProviderMetaData.ABI

// IStorageProvider is an auto generated Go binding around an Ethereum contract.
type IStorageProvider struct {
	IStorageProviderCaller     // Read-only binding to the contract
	IStorageProviderTransactor // Write-only binding to the contract
	IStorageProviderFilterer   // Log filterer for contract events
}

// IStorageProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStorageProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStorageProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStorageProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStorageProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStorageProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStorageProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStorageProviderSession struct {
	Contract     *IStorageProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStorageProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStorageProviderCallerSession struct {
	Contract *IStorageProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IStorageProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStorageProviderTransactorSession struct {
	Contract     *IStorageProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IStorageProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStorageProviderRaw struct {
	Contract *IStorageProvider // Generic contract binding to access the raw methods on
}

// IStorageProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStorageProviderCallerRaw struct {
	Contract *IStorageProviderCaller // Generic read-only contract binding to access the raw methods on
}

// IStorageProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStorageProviderTransactorRaw struct {
	Contract *IStorageProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStorageProvider creates a new instance of IStorageProvider, bound to a specific deployed contract.
func NewIStorageProvider(address common.Address, backend bind.ContractBackend) (*IStorageProvider, error) {
	contract, err := bindIStorageProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStorageProvider{IStorageProviderCaller: IStorageProviderCaller{contract: contract}, IStorageProviderTransactor: IStorageProviderTransactor{contract: contract}, IStorageProviderFilterer: IStorageProviderFilterer{contract: contract}}, nil
}

// NewIStorageProviderCaller creates a new read-only instance of IStorageProvider, bound to a specific deployed contract.
func NewIStorageProviderCaller(address common.Address, caller bind.ContractCaller) (*IStorageProviderCaller, error) {
	contract, err := bindIStorageProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStorageProviderCaller{contract: contract}, nil
}

// NewIStorageProviderTransactor creates a new write-only instance of IStorageProvider, bound to a specific deployed contract.
func NewIStorageProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*IStorageProviderTransactor, error) {
	contract, err := bindIStorageProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStorageProviderTransactor{contract: contract}, nil
}

// NewIStorageProviderFilterer creates a new log filterer instance of IStorageProvider, bound to a specific deployed contract.
func NewIStorageProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*IStorageProviderFilterer, error) {
	contract, err := bindIStorageProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStorageProviderFilterer{contract: contract}, nil
}

// bindIStorageProvider binds a generic wrapper to an already deployed contract.
func bindIStorageProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IStorageProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStorageProvider *IStorageProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStorageProvider.Contract.IStorageProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStorageProvider *IStorageProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStorageProvider.Contract.IStorageProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStorageProvider *IStorageProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStorageProvider.Contract.IStorageProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStorageProvider *IStorageProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStorageProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStorageProvider *IStorageProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStorageProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStorageProvider *IStorageProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStorageProvider.Contract.contract.Transact(opts, method, params...)
}

// StorageProvider is a free data retrieval call binding the contract method 0x163c507d.
//
// Solidity: function storageProvider(uint32 id) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string) storageProvider)
func (_IStorageProvider *IStorageProviderCaller) StorageProvider(opts *bind.CallOpts, id uint32) (StorageProvider, error) {
	var out []interface{}
	err := _IStorageProvider.contract.Call(opts, &out, "storageProvider", id)

	if err != nil {
		return *new(StorageProvider), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageProvider)).(*StorageProvider)

	return out0, err

}

// StorageProvider is a free data retrieval call binding the contract method 0x163c507d.
//
// Solidity: function storageProvider(uint32 id) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string) storageProvider)
func (_IStorageProvider *IStorageProviderSession) StorageProvider(id uint32) (StorageProvider, error) {
	return _IStorageProvider.Contract.StorageProvider(&_IStorageProvider.CallOpts, id)
}

// StorageProvider is a free data retrieval call binding the contract method 0x163c507d.
//
// Solidity: function storageProvider(uint32 id) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string) storageProvider)
func (_IStorageProvider *IStorageProviderCallerSession) StorageProvider(id uint32) (StorageProvider, error) {
	return _IStorageProvider.Contract.StorageProvider(&_IStorageProvider.CallOpts, id)
}

// StorageProviderByOperatorAddress is a free data retrieval call binding the contract method 0xf26effb0.
//
// Solidity: function storageProviderByOperatorAddress(address operatorAddress) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string) storageProvider)
func (_IStorageProvider *IStorageProviderCaller) StorageProviderByOperatorAddress(opts *bind.CallOpts, operatorAddress common.Address) (StorageProvider, error) {
	var out []interface{}
	err := _IStorageProvider.contract.Call(opts, &out, "storageProviderByOperatorAddress", operatorAddress)

	if err != nil {
		return *new(StorageProvider), err
	}

	out0 := *abi.ConvertType(out[0], new(StorageProvider)).(*StorageProvider)

	return out0, err

}

// StorageProviderByOperatorAddress is a free data retrieval call binding the contract method 0xf26effb0.
//
// Solidity: function storageProviderByOperatorAddress(address operatorAddress) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string) storageProvider)
func (_IStorageProvider *IStorageProviderSession) StorageProviderByOperatorAddress(operatorAddress common.Address) (StorageProvider, error) {
	return _IStorageProvider.Contract.StorageProviderByOperatorAddress(&_IStorageProvider.CallOpts, operatorAddress)
}

// StorageProviderByOperatorAddress is a free data retrieval call binding the contract method 0xf26effb0.
//
// Solidity: function storageProviderByOperatorAddress(address operatorAddress) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string) storageProvider)
func (_IStorageProvider *IStorageProviderCallerSession) StorageProviderByOperatorAddress(operatorAddress common.Address) (StorageProvider, error) {
	return _IStorageProvider.Contract.StorageProviderByOperatorAddress(&_IStorageProvider.CallOpts, operatorAddress)
}

// StorageProviderPrice is a free data retrieval call binding the contract method 0x9803fb33.
//
// Solidity: function storageProviderPrice(address operatorAddress) view returns((uint32,uint256,uint256,uint64,uint256) spStoragePrice)
func (_IStorageProvider *IStorageProviderCaller) StorageProviderPrice(opts *bind.CallOpts, operatorAddress common.Address) (SpStoragePrice, error) {
	var out []interface{}
	err := _IStorageProvider.contract.Call(opts, &out, "storageProviderPrice", operatorAddress)

	if err != nil {
		return *new(SpStoragePrice), err
	}

	out0 := *abi.ConvertType(out[0], new(SpStoragePrice)).(*SpStoragePrice)

	return out0, err

}

// StorageProviderPrice is a free data retrieval call binding the contract method 0x9803fb33.
//
// Solidity: function storageProviderPrice(address operatorAddress) view returns((uint32,uint256,uint256,uint64,uint256) spStoragePrice)
func (_IStorageProvider *IStorageProviderSession) StorageProviderPrice(operatorAddress common.Address) (SpStoragePrice, error) {
	return _IStorageProvider.Contract.StorageProviderPrice(&_IStorageProvider.CallOpts, operatorAddress)
}

// StorageProviderPrice is a free data retrieval call binding the contract method 0x9803fb33.
//
// Solidity: function storageProviderPrice(address operatorAddress) view returns((uint32,uint256,uint256,uint64,uint256) spStoragePrice)
func (_IStorageProvider *IStorageProviderCallerSession) StorageProviderPrice(operatorAddress common.Address) (SpStoragePrice, error) {
	return _IStorageProvider.Contract.StorageProviderPrice(&_IStorageProvider.CallOpts, operatorAddress)
}

// StorageProviders is a free data retrieval call binding the contract method 0x12196b83.
//
// Solidity: function storageProviders((bytes,uint64,uint64,bool,bool) pagination) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string)[] storageProviders, (bytes,uint64) pageResponse)
func (_IStorageProvider *IStorageProviderCaller) StorageProviders(opts *bind.CallOpts, pagination PageRequest) (struct {
	StorageProviders []StorageProvider
	PageResponse     PageResponse
}, error) {
	var out []interface{}
	err := _IStorageProvider.contract.Call(opts, &out, "storageProviders", pagination)

	outstruct := new(struct {
		StorageProviders []StorageProvider
		PageResponse     PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StorageProviders = *abi.ConvertType(out[0], new([]StorageProvider)).(*[]StorageProvider)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// StorageProviders is a free data retrieval call binding the contract method 0x12196b83.
//
// Solidity: function storageProviders((bytes,uint64,uint64,bool,bool) pagination) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string)[] storageProviders, (bytes,uint64) pageResponse)
func (_IStorageProvider *IStorageProviderSession) StorageProviders(pagination PageRequest) (struct {
	StorageProviders []StorageProvider
	PageResponse     PageResponse
}, error) {
	return _IStorageProvider.Contract.StorageProviders(&_IStorageProvider.CallOpts, pagination)
}

// StorageProviders is a free data retrieval call binding the contract method 0x12196b83.
//
// Solidity: function storageProviders((bytes,uint64,uint64,bool,bool) pagination) view returns((uint32,string,string,string,string,string,string,uint256,uint8,string,(string,string,string,string,string),string)[] storageProviders, (bytes,uint64) pageResponse)
func (_IStorageProvider *IStorageProviderCallerSession) StorageProviders(pagination PageRequest) (struct {
	StorageProviders []StorageProvider
	PageResponse     PageResponse
}, error) {
	return _IStorageProvider.Contract.StorageProviders(&_IStorageProvider.CallOpts, pagination)
}

// UpdateSPPrice is a paid mutator transaction binding the contract method 0x630d6e45.
//
// Solidity: function updateSPPrice(uint256 readPrice, uint64 freeReadQuota, uint256 storePrice) returns(bool success)
func (_IStorageProvider *IStorageProviderTransactor) UpdateSPPrice(opts *bind.TransactOpts, readPrice *big.Int, freeReadQuota uint64, storePrice *big.Int) (*types.Transaction, error) {
	return _IStorageProvider.contract.Transact(opts, "updateSPPrice", readPrice, freeReadQuota, storePrice)
}

// UpdateSPPrice is a paid mutator transaction binding the contract method 0x630d6e45.
//
// Solidity: function updateSPPrice(uint256 readPrice, uint64 freeReadQuota, uint256 storePrice) returns(bool success)
func (_IStorageProvider *IStorageProviderSession) UpdateSPPrice(readPrice *big.Int, freeReadQuota uint64, storePrice *big.Int) (*types.Transaction, error) {
	return _IStorageProvider.Contract.UpdateSPPrice(&_IStorageProvider.TransactOpts, readPrice, freeReadQuota, storePrice)
}

// UpdateSPPrice is a paid mutator transaction binding the contract method 0x630d6e45.
//
// Solidity: function updateSPPrice(uint256 readPrice, uint64 freeReadQuota, uint256 storePrice) returns(bool success)
func (_IStorageProvider *IStorageProviderTransactorSession) UpdateSPPrice(readPrice *big.Int, freeReadQuota uint64, storePrice *big.Int) (*types.Transaction, error) {
	return _IStorageProvider.Contract.UpdateSPPrice(&_IStorageProvider.TransactOpts, readPrice, freeReadQuota, storePrice)
}

// IStorageProviderUpdateSPPriceIterator is returned from FilterUpdateSPPrice and is used to iterate over the raw logs and unpacked data for UpdateSPPrice events raised by the IStorageProvider contract.
type IStorageProviderUpdateSPPriceIterator struct {
	Event *IStorageProviderUpdateSPPrice // Event containing the contract specifics and raw log

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
func (it *IStorageProviderUpdateSPPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStorageProviderUpdateSPPrice)
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
		it.Event = new(IStorageProviderUpdateSPPrice)
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
func (it *IStorageProviderUpdateSPPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStorageProviderUpdateSPPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStorageProviderUpdateSPPrice represents a UpdateSPPrice event raised by the IStorageProvider contract.
type IStorageProviderUpdateSPPrice struct {
	StorageProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateSPPrice is a free log retrieval operation binding the contract event 0x6d0de22fee7a8aa285edd35229fda9ba7ecdf8346300e3245e9baf7340969937.
//
// Solidity: event UpdateSPPrice(address indexed storageProvider)
func (_IStorageProvider *IStorageProviderFilterer) FilterUpdateSPPrice(opts *bind.FilterOpts, storageProvider []common.Address) (*IStorageProviderUpdateSPPriceIterator, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IStorageProvider.contract.FilterLogs(opts, "UpdateSPPrice", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return &IStorageProviderUpdateSPPriceIterator{contract: _IStorageProvider.contract, event: "UpdateSPPrice", logs: logs, sub: sub}, nil
}

// WatchUpdateSPPrice is a free log subscription operation binding the contract event 0x6d0de22fee7a8aa285edd35229fda9ba7ecdf8346300e3245e9baf7340969937.
//
// Solidity: event UpdateSPPrice(address indexed storageProvider)
func (_IStorageProvider *IStorageProviderFilterer) WatchUpdateSPPrice(opts *bind.WatchOpts, sink chan<- *IStorageProviderUpdateSPPrice, storageProvider []common.Address) (event.Subscription, error) {

	var storageProviderRule []interface{}
	for _, storageProviderItem := range storageProvider {
		storageProviderRule = append(storageProviderRule, storageProviderItem)
	}

	logs, sub, err := _IStorageProvider.contract.WatchLogs(opts, "UpdateSPPrice", storageProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStorageProviderUpdateSPPrice)
				if err := _IStorageProvider.contract.UnpackLog(event, "UpdateSPPrice", log); err != nil {
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

// ParseUpdateSPPrice is a log parse operation binding the contract event 0x6d0de22fee7a8aa285edd35229fda9ba7ecdf8346300e3245e9baf7340969937.
//
// Solidity: event UpdateSPPrice(address indexed storageProvider)
func (_IStorageProvider *IStorageProviderFilterer) ParseUpdateSPPrice(log types.Log) (*IStorageProviderUpdateSPPrice, error) {
	event := new(IStorageProviderUpdateSPPrice)
	if err := _IStorageProvider.contract.UnpackLog(event, "UpdateSPPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
