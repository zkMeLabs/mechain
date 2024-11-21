// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bank

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

// DenomOwner is an auto generated low-level Go binding around an user-defined struct.
type DenomOwner struct {
	AccountAddress common.Address
	Balance        Coin
}

// DenomUnit is an auto generated low-level Go binding around an user-defined struct.
type DenomUnit struct {
	Denom    string
	Exponent uint32
	Aliases  []string
}

// Metadata is an auto generated low-level Go binding around an user-defined struct.
type Metadata struct {
	Description string
	DenomUnits  []DenomUnit
	Base        string
	Display     string
	Name        string
	Symbol      string
	Uri         string
	UriHash     string
}

// Output is an auto generated low-level Go binding around an user-defined struct.
type Output struct {
	ToAddress common.Address
	Amount    []Coin
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
	SendEnabled        []SendEnabled
	DefaultSendEnabled bool
}

// SendEnabled is an auto generated low-level Go binding around an user-defined struct.
type SendEnabled struct {
	Denom   string
	Enabled bool
}

// IBankMetaData contains all meta data concerning the IBank contract.
var IBankMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"MultiSend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"amount\",\"type\":\"string\"}],\"name\":\"Send\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"allBalances\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"balance\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"balance\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"denomMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"exponent\",\"type\":\"uint32\"},{\"internalType\":\"string[]\",\"name\":\"aliases\",\"type\":\"string[]\"}],\"internalType\":\"structDenomUnit[]\",\"name\":\"denomUnits\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"base\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"display\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uriHash\",\"type\":\"string\"}],\"internalType\":\"structMetadata\",\"name\":\"metadata\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"denomOwners\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"balance\",\"type\":\"tuple\"}],\"internalType\":\"structDenomOwner[]\",\"name\":\"denomOwners\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"denomsMetadata\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"exponent\",\"type\":\"uint32\"},{\"internalType\":\"string[]\",\"name\":\"aliases\",\"type\":\"string[]\"}],\"internalType\":\"structDenomUnit[]\",\"name\":\"denomUnits\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"base\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"display\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uriHash\",\"type\":\"string\"}],\"internalType\":\"structMetadata[]\",\"name\":\"metadatas\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"internalType\":\"structOutput[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"}],\"name\":\"multiSend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structSendEnabled[]\",\"name\":\"sendEnabled\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"defaultSendEnabled\",\"type\":\"bool\"}],\"internalType\":\"structParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"amount\",\"type\":\"tuple[]\"}],\"name\":\"send\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"denoms\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"sendEnabled\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"internalType\":\"structSendEnabled[]\",\"name\":\"sendEnableds\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"spendableBalanceByDenom\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"balance\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"spendableBalances\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"supplyOf\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin\",\"name\":\"amount\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"offset\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"limit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"countTotal\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reverse\",\"type\":\"bool\"}],\"internalType\":\"structPageRequest\",\"name\":\"pageRequest\",\"type\":\"tuple\"}],\"name\":\"totalSupply\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCoin[]\",\"name\":\"supply\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"nextKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"total\",\"type\":\"uint64\"}],\"internalType\":\"structPageResponse\",\"name\":\"pageResponse\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBankABI is the input ABI used to generate the binding from.
// Deprecated: Use IBankMetaData.ABI instead.
var IBankABI = IBankMetaData.ABI

// IBank is an auto generated Go binding around an Ethereum contract.
type IBank struct {
	IBankCaller     // Read-only binding to the contract
	IBankTransactor // Write-only binding to the contract
	IBankFilterer   // Log filterer for contract events
}

// IBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBankSession struct {
	Contract     *IBank            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBankCallerSession struct {
	Contract *IBankCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBankTransactorSession struct {
	Contract     *IBankTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBankRaw struct {
	Contract *IBank // Generic contract binding to access the raw methods on
}

// IBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBankCallerRaw struct {
	Contract *IBankCaller // Generic read-only contract binding to access the raw methods on
}

// IBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBankTransactorRaw struct {
	Contract *IBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBank creates a new instance of IBank, bound to a specific deployed contract.
func NewIBank(address common.Address, backend bind.ContractBackend) (*IBank, error) {
	contract, err := bindIBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBank{IBankCaller: IBankCaller{contract: contract}, IBankTransactor: IBankTransactor{contract: contract}, IBankFilterer: IBankFilterer{contract: contract}}, nil
}

// NewIBankCaller creates a new read-only instance of IBank, bound to a specific deployed contract.
func NewIBankCaller(address common.Address, caller bind.ContractCaller) (*IBankCaller, error) {
	contract, err := bindIBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBankCaller{contract: contract}, nil
}

// NewIBankTransactor creates a new write-only instance of IBank, bound to a specific deployed contract.
func NewIBankTransactor(address common.Address, transactor bind.ContractTransactor) (*IBankTransactor, error) {
	contract, err := bindIBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBankTransactor{contract: contract}, nil
}

// NewIBankFilterer creates a new log filterer instance of IBank, bound to a specific deployed contract.
func NewIBankFilterer(address common.Address, filterer bind.ContractFilterer) (*IBankFilterer, error) {
	contract, err := bindIBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBankFilterer{contract: contract}, nil
}

// bindIBank binds a generic wrapper to an already deployed contract.
func bindIBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBankMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBank *IBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBank.Contract.IBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBank *IBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBank.Contract.IBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBank *IBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBank.Contract.IBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBank *IBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBank *IBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBank *IBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBank.Contract.contract.Transact(opts, method, params...)
}

// AllBalances is a free data retrieval call binding the contract method 0x013fc38f.
//
// Solidity: function allBalances(address accountAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] balances, (bytes,uint64) pageResponse)
func (_IBank *IBankCaller) AllBalances(opts *bind.CallOpts, accountAddress common.Address, pageRequest PageRequest) (struct {
	Balances     []Coin
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "allBalances", accountAddress, pageRequest)

	outstruct := new(struct {
		Balances     []Coin
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balances = *abi.ConvertType(out[0], new([]Coin)).(*[]Coin)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// AllBalances is a free data retrieval call binding the contract method 0x013fc38f.
//
// Solidity: function allBalances(address accountAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] balances, (bytes,uint64) pageResponse)
func (_IBank *IBankSession) AllBalances(accountAddress common.Address, pageRequest PageRequest) (struct {
	Balances     []Coin
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.AllBalances(&_IBank.CallOpts, accountAddress, pageRequest)
}

// AllBalances is a free data retrieval call binding the contract method 0x013fc38f.
//
// Solidity: function allBalances(address accountAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] balances, (bytes,uint64) pageResponse)
func (_IBank *IBankCallerSession) AllBalances(accountAddress common.Address, pageRequest PageRequest) (struct {
	Balances     []Coin
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.AllBalances(&_IBank.CallOpts, accountAddress, pageRequest)
}

// Balance is a free data retrieval call binding the contract method 0x16cadeab.
//
// Solidity: function balance(address accountAddress, string denom) view returns((string,uint256) balance)
func (_IBank *IBankCaller) Balance(opts *bind.CallOpts, accountAddress common.Address, denom string) (Coin, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "balance", accountAddress, denom)

	if err != nil {
		return *new(Coin), err
	}

	out0 := *abi.ConvertType(out[0], new(Coin)).(*Coin)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x16cadeab.
//
// Solidity: function balance(address accountAddress, string denom) view returns((string,uint256) balance)
func (_IBank *IBankSession) Balance(accountAddress common.Address, denom string) (Coin, error) {
	return _IBank.Contract.Balance(&_IBank.CallOpts, accountAddress, denom)
}

// Balance is a free data retrieval call binding the contract method 0x16cadeab.
//
// Solidity: function balance(address accountAddress, string denom) view returns((string,uint256) balance)
func (_IBank *IBankCallerSession) Balance(accountAddress common.Address, denom string) (Coin, error) {
	return _IBank.Contract.Balance(&_IBank.CallOpts, accountAddress, denom)
}

// DenomMetadata is a free data retrieval call binding the contract method 0xbf167569.
//
// Solidity: function denomMetadata(string denom) view returns((string,(string,uint32,string[])[],string,string,string,string,string,string) metadata)
func (_IBank *IBankCaller) DenomMetadata(opts *bind.CallOpts, denom string) (Metadata, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "denomMetadata", denom)

	if err != nil {
		return *new(Metadata), err
	}

	out0 := *abi.ConvertType(out[0], new(Metadata)).(*Metadata)

	return out0, err

}

// DenomMetadata is a free data retrieval call binding the contract method 0xbf167569.
//
// Solidity: function denomMetadata(string denom) view returns((string,(string,uint32,string[])[],string,string,string,string,string,string) metadata)
func (_IBank *IBankSession) DenomMetadata(denom string) (Metadata, error) {
	return _IBank.Contract.DenomMetadata(&_IBank.CallOpts, denom)
}

// DenomMetadata is a free data retrieval call binding the contract method 0xbf167569.
//
// Solidity: function denomMetadata(string denom) view returns((string,(string,uint32,string[])[],string,string,string,string,string,string) metadata)
func (_IBank *IBankCallerSession) DenomMetadata(denom string) (Metadata, error) {
	return _IBank.Contract.DenomMetadata(&_IBank.CallOpts, denom)
}

// DenomOwners is a free data retrieval call binding the contract method 0x35e2bdea.
//
// Solidity: function denomOwners(string denom, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((address,(string,uint256))[] denomOwners, (bytes,uint64) pageResponse)
func (_IBank *IBankCaller) DenomOwners(opts *bind.CallOpts, denom string, pageRequest PageRequest) (struct {
	DenomOwners  []DenomOwner
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "denomOwners", denom, pageRequest)

	outstruct := new(struct {
		DenomOwners  []DenomOwner
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DenomOwners = *abi.ConvertType(out[0], new([]DenomOwner)).(*[]DenomOwner)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// DenomOwners is a free data retrieval call binding the contract method 0x35e2bdea.
//
// Solidity: function denomOwners(string denom, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((address,(string,uint256))[] denomOwners, (bytes,uint64) pageResponse)
func (_IBank *IBankSession) DenomOwners(denom string, pageRequest PageRequest) (struct {
	DenomOwners  []DenomOwner
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.DenomOwners(&_IBank.CallOpts, denom, pageRequest)
}

// DenomOwners is a free data retrieval call binding the contract method 0x35e2bdea.
//
// Solidity: function denomOwners(string denom, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((address,(string,uint256))[] denomOwners, (bytes,uint64) pageResponse)
func (_IBank *IBankCallerSession) DenomOwners(denom string, pageRequest PageRequest) (struct {
	DenomOwners  []DenomOwner
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.DenomOwners(&_IBank.CallOpts, denom, pageRequest)
}

// DenomsMetadata is a free data retrieval call binding the contract method 0xe41ddb11.
//
// Solidity: function denomsMetadata((bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,(string,uint32,string[])[],string,string,string,string,string,string)[] metadatas, (bytes,uint64) pageResponse)
func (_IBank *IBankCaller) DenomsMetadata(opts *bind.CallOpts, pageRequest PageRequest) (struct {
	Metadatas    []Metadata
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "denomsMetadata", pageRequest)

	outstruct := new(struct {
		Metadatas    []Metadata
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Metadatas = *abi.ConvertType(out[0], new([]Metadata)).(*[]Metadata)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// DenomsMetadata is a free data retrieval call binding the contract method 0xe41ddb11.
//
// Solidity: function denomsMetadata((bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,(string,uint32,string[])[],string,string,string,string,string,string)[] metadatas, (bytes,uint64) pageResponse)
func (_IBank *IBankSession) DenomsMetadata(pageRequest PageRequest) (struct {
	Metadatas    []Metadata
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.DenomsMetadata(&_IBank.CallOpts, pageRequest)
}

// DenomsMetadata is a free data retrieval call binding the contract method 0xe41ddb11.
//
// Solidity: function denomsMetadata((bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,(string,uint32,string[])[],string,string,string,string,string,string)[] metadatas, (bytes,uint64) pageResponse)
func (_IBank *IBankCallerSession) DenomsMetadata(pageRequest PageRequest) (struct {
	Metadatas    []Metadata
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.DenomsMetadata(&_IBank.CallOpts, pageRequest)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((string,bool)[],bool) params)
func (_IBank *IBankCaller) Params(opts *bind.CallOpts) (Params, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "params")

	if err != nil {
		return *new(Params), err
	}

	out0 := *abi.ConvertType(out[0], new(Params)).(*Params)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((string,bool)[],bool) params)
func (_IBank *IBankSession) Params() (Params, error) {
	return _IBank.Contract.Params(&_IBank.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(((string,bool)[],bool) params)
func (_IBank *IBankCallerSession) Params() (Params, error) {
	return _IBank.Contract.Params(&_IBank.CallOpts)
}

// SendEnabled is a free data retrieval call binding the contract method 0xc27940ea.
//
// Solidity: function sendEnabled(string[] denoms, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,bool)[] sendEnableds, (bytes,uint64) pageResponse)
func (_IBank *IBankCaller) SendEnabled(opts *bind.CallOpts, denoms []string, pageRequest PageRequest) (struct {
	SendEnableds []SendEnabled
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "sendEnabled", denoms, pageRequest)

	outstruct := new(struct {
		SendEnableds []SendEnabled
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SendEnableds = *abi.ConvertType(out[0], new([]SendEnabled)).(*[]SendEnabled)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// SendEnabled is a free data retrieval call binding the contract method 0xc27940ea.
//
// Solidity: function sendEnabled(string[] denoms, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,bool)[] sendEnableds, (bytes,uint64) pageResponse)
func (_IBank *IBankSession) SendEnabled(denoms []string, pageRequest PageRequest) (struct {
	SendEnableds []SendEnabled
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.SendEnabled(&_IBank.CallOpts, denoms, pageRequest)
}

// SendEnabled is a free data retrieval call binding the contract method 0xc27940ea.
//
// Solidity: function sendEnabled(string[] denoms, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,bool)[] sendEnableds, (bytes,uint64) pageResponse)
func (_IBank *IBankCallerSession) SendEnabled(denoms []string, pageRequest PageRequest) (struct {
	SendEnableds []SendEnabled
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.SendEnabled(&_IBank.CallOpts, denoms, pageRequest)
}

// SpendableBalanceByDenom is a free data retrieval call binding the contract method 0xfb6dd4fe.
//
// Solidity: function spendableBalanceByDenom(address accountAddress, string denom) view returns((string,uint256) balance)
func (_IBank *IBankCaller) SpendableBalanceByDenom(opts *bind.CallOpts, accountAddress common.Address, denom string) (Coin, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "spendableBalanceByDenom", accountAddress, denom)

	if err != nil {
		return *new(Coin), err
	}

	out0 := *abi.ConvertType(out[0], new(Coin)).(*Coin)

	return out0, err

}

// SpendableBalanceByDenom is a free data retrieval call binding the contract method 0xfb6dd4fe.
//
// Solidity: function spendableBalanceByDenom(address accountAddress, string denom) view returns((string,uint256) balance)
func (_IBank *IBankSession) SpendableBalanceByDenom(accountAddress common.Address, denom string) (Coin, error) {
	return _IBank.Contract.SpendableBalanceByDenom(&_IBank.CallOpts, accountAddress, denom)
}

// SpendableBalanceByDenom is a free data retrieval call binding the contract method 0xfb6dd4fe.
//
// Solidity: function spendableBalanceByDenom(address accountAddress, string denom) view returns((string,uint256) balance)
func (_IBank *IBankCallerSession) SpendableBalanceByDenom(accountAddress common.Address, denom string) (Coin, error) {
	return _IBank.Contract.SpendableBalanceByDenom(&_IBank.CallOpts, accountAddress, denom)
}

// SpendableBalances is a free data retrieval call binding the contract method 0xf8554168.
//
// Solidity: function spendableBalances(address accountAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] balances, (bytes,uint64) pageResponse)
func (_IBank *IBankCaller) SpendableBalances(opts *bind.CallOpts, accountAddress common.Address, pageRequest PageRequest) (struct {
	Balances     []Coin
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "spendableBalances", accountAddress, pageRequest)

	outstruct := new(struct {
		Balances     []Coin
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balances = *abi.ConvertType(out[0], new([]Coin)).(*[]Coin)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// SpendableBalances is a free data retrieval call binding the contract method 0xf8554168.
//
// Solidity: function spendableBalances(address accountAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] balances, (bytes,uint64) pageResponse)
func (_IBank *IBankSession) SpendableBalances(accountAddress common.Address, pageRequest PageRequest) (struct {
	Balances     []Coin
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.SpendableBalances(&_IBank.CallOpts, accountAddress, pageRequest)
}

// SpendableBalances is a free data retrieval call binding the contract method 0xf8554168.
//
// Solidity: function spendableBalances(address accountAddress, (bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] balances, (bytes,uint64) pageResponse)
func (_IBank *IBankCallerSession) SpendableBalances(accountAddress common.Address, pageRequest PageRequest) (struct {
	Balances     []Coin
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.SpendableBalances(&_IBank.CallOpts, accountAddress, pageRequest)
}

// SupplyOf is a free data retrieval call binding the contract method 0x3cda0103.
//
// Solidity: function supplyOf(string denom) view returns((string,uint256) amount)
func (_IBank *IBankCaller) SupplyOf(opts *bind.CallOpts, denom string) (Coin, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "supplyOf", denom)

	if err != nil {
		return *new(Coin), err
	}

	out0 := *abi.ConvertType(out[0], new(Coin)).(*Coin)

	return out0, err

}

// SupplyOf is a free data retrieval call binding the contract method 0x3cda0103.
//
// Solidity: function supplyOf(string denom) view returns((string,uint256) amount)
func (_IBank *IBankSession) SupplyOf(denom string) (Coin, error) {
	return _IBank.Contract.SupplyOf(&_IBank.CallOpts, denom)
}

// SupplyOf is a free data retrieval call binding the contract method 0x3cda0103.
//
// Solidity: function supplyOf(string denom) view returns((string,uint256) amount)
func (_IBank *IBankCallerSession) SupplyOf(denom string) (Coin, error) {
	return _IBank.Contract.SupplyOf(&_IBank.CallOpts, denom)
}

// TotalSupply is a free data retrieval call binding the contract method 0xaa52071f.
//
// Solidity: function totalSupply((bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] supply, (bytes,uint64) pageResponse)
func (_IBank *IBankCaller) TotalSupply(opts *bind.CallOpts, pageRequest PageRequest) (struct {
	Supply       []Coin
	PageResponse PageResponse
}, error) {
	var out []interface{}
	err := _IBank.contract.Call(opts, &out, "totalSupply", pageRequest)

	outstruct := new(struct {
		Supply       []Coin
		PageResponse PageResponse
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Supply = *abi.ConvertType(out[0], new([]Coin)).(*[]Coin)
	outstruct.PageResponse = *abi.ConvertType(out[1], new(PageResponse)).(*PageResponse)

	return *outstruct, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xaa52071f.
//
// Solidity: function totalSupply((bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] supply, (bytes,uint64) pageResponse)
func (_IBank *IBankSession) TotalSupply(pageRequest PageRequest) (struct {
	Supply       []Coin
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.TotalSupply(&_IBank.CallOpts, pageRequest)
}

// TotalSupply is a free data retrieval call binding the contract method 0xaa52071f.
//
// Solidity: function totalSupply((bytes,uint64,uint64,bool,bool) pageRequest) view returns((string,uint256)[] supply, (bytes,uint64) pageResponse)
func (_IBank *IBankCallerSession) TotalSupply(pageRequest PageRequest) (struct {
	Supply       []Coin
	PageResponse PageResponse
}, error) {
	return _IBank.Contract.TotalSupply(&_IBank.CallOpts, pageRequest)
}

// MultiSend is a paid mutator transaction binding the contract method 0xd6915a4b.
//
// Solidity: function multiSend((address,(string,uint256)[])[] outputs) returns(bool success)
func (_IBank *IBankTransactor) MultiSend(opts *bind.TransactOpts, outputs []Output) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "multiSend", outputs)
}

// MultiSend is a paid mutator transaction binding the contract method 0xd6915a4b.
//
// Solidity: function multiSend((address,(string,uint256)[])[] outputs) returns(bool success)
func (_IBank *IBankSession) MultiSend(outputs []Output) (*types.Transaction, error) {
	return _IBank.Contract.MultiSend(&_IBank.TransactOpts, outputs)
}

// MultiSend is a paid mutator transaction binding the contract method 0xd6915a4b.
//
// Solidity: function multiSend((address,(string,uint256)[])[] outputs) returns(bool success)
func (_IBank *IBankTransactorSession) MultiSend(outputs []Output) (*types.Transaction, error) {
	return _IBank.Contract.MultiSend(&_IBank.TransactOpts, outputs)
}

// Send is a paid mutator transaction binding the contract method 0x8f7f2b20.
//
// Solidity: function send(address toAddress, (string,uint256)[] amount) returns(bool success)
func (_IBank *IBankTransactor) Send(opts *bind.TransactOpts, toAddress common.Address, amount []Coin) (*types.Transaction, error) {
	return _IBank.contract.Transact(opts, "send", toAddress, amount)
}

// Send is a paid mutator transaction binding the contract method 0x8f7f2b20.
//
// Solidity: function send(address toAddress, (string,uint256)[] amount) returns(bool success)
func (_IBank *IBankSession) Send(toAddress common.Address, amount []Coin) (*types.Transaction, error) {
	return _IBank.Contract.Send(&_IBank.TransactOpts, toAddress, amount)
}

// Send is a paid mutator transaction binding the contract method 0x8f7f2b20.
//
// Solidity: function send(address toAddress, (string,uint256)[] amount) returns(bool success)
func (_IBank *IBankTransactorSession) Send(toAddress common.Address, amount []Coin) (*types.Transaction, error) {
	return _IBank.Contract.Send(&_IBank.TransactOpts, toAddress, amount)
}

// IBankMultiSendIterator is returned from FilterMultiSend and is used to iterate over the raw logs and unpacked data for MultiSend events raised by the IBank contract.
type IBankMultiSendIterator struct {
	Event *IBankMultiSend // Event containing the contract specifics and raw log

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
func (it *IBankMultiSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankMultiSend)
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
		it.Event = new(IBankMultiSend)
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
func (it *IBankMultiSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankMultiSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankMultiSend represents a MultiSend event raised by the IBank contract.
type IBankMultiSend struct {
	FromAddress common.Address
	Amount      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMultiSend is a free log retrieval operation binding the contract event 0xdb266aa59af276151406ed2fa19f5376eff4fc03c649e9474b2e69213dbc5e3f.
//
// Solidity: event MultiSend(address indexed fromAddress, string amount)
func (_IBank *IBankFilterer) FilterMultiSend(opts *bind.FilterOpts, fromAddress []common.Address) (*IBankMultiSendIterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _IBank.contract.FilterLogs(opts, "MultiSend", fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &IBankMultiSendIterator{contract: _IBank.contract, event: "MultiSend", logs: logs, sub: sub}, nil
}

// WatchMultiSend is a free log subscription operation binding the contract event 0xdb266aa59af276151406ed2fa19f5376eff4fc03c649e9474b2e69213dbc5e3f.
//
// Solidity: event MultiSend(address indexed fromAddress, string amount)
func (_IBank *IBankFilterer) WatchMultiSend(opts *bind.WatchOpts, sink chan<- *IBankMultiSend, fromAddress []common.Address) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _IBank.contract.WatchLogs(opts, "MultiSend", fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankMultiSend)
				if err := _IBank.contract.UnpackLog(event, "MultiSend", log); err != nil {
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

// ParseMultiSend is a log parse operation binding the contract event 0xdb266aa59af276151406ed2fa19f5376eff4fc03c649e9474b2e69213dbc5e3f.
//
// Solidity: event MultiSend(address indexed fromAddress, string amount)
func (_IBank *IBankFilterer) ParseMultiSend(log types.Log) (*IBankMultiSend, error) {
	event := new(IBankMultiSend)
	if err := _IBank.contract.UnpackLog(event, "MultiSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBankSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the IBank contract.
type IBankSendIterator struct {
	Event *IBankSend // Event containing the contract specifics and raw log

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
func (it *IBankSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBankSend)
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
		it.Event = new(IBankSend)
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
func (it *IBankSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBankSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBankSend represents a Send event raised by the IBank contract.
type IBankSend struct {
	FromAddress common.Address
	ToAddress   common.Address
	Amount      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0xb53a0a2251075d1bc8c4dcce157b80078af99958b6ff4f15bf48a414270d7779.
//
// Solidity: event Send(address indexed fromAddress, address indexed toAddress, string amount)
func (_IBank *IBankFilterer) FilterSend(opts *bind.FilterOpts, fromAddress []common.Address, toAddress []common.Address) (*IBankSendIterator, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _IBank.contract.FilterLogs(opts, "Send", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &IBankSendIterator{contract: _IBank.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0xb53a0a2251075d1bc8c4dcce157b80078af99958b6ff4f15bf48a414270d7779.
//
// Solidity: event Send(address indexed fromAddress, address indexed toAddress, string amount)
func (_IBank *IBankFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *IBankSend, fromAddress []common.Address, toAddress []common.Address) (event.Subscription, error) {

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _IBank.contract.WatchLogs(opts, "Send", fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBankSend)
				if err := _IBank.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0xb53a0a2251075d1bc8c4dcce157b80078af99958b6ff4f15bf48a414270d7779.
//
// Solidity: event Send(address indexed fromAddress, address indexed toAddress, string amount)
func (_IBank *IBankFilterer) ParseSend(log types.Log) (*IBankSend, error) {
	event := new(IBankSend)
	if err := _IBank.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
