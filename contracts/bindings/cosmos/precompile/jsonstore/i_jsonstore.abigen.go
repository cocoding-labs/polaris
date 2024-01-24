// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jsonstore

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

// JsonStoreMetaData contains all meta data concerning the JsonStore contract.
var JsonStoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"dataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"getBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"getInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"getUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"values\",\"type\":\"bool[]\"}],\"name\":\"setBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"int256[]\",\"name\":\"values\",\"type\":\"int256[]\"}],\"name\":\"setInt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setInt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"rawBlob\",\"type\":\"string\"}],\"name\":\"setRaw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"rawBlobs\",\"type\":\"string[]\"}],\"name\":\"setRaw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"setUint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setUint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"subReplace\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"subReplace\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"subReplaceBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"values\",\"type\":\"bool[]\"}],\"name\":\"subReplaceBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"subReplaceInt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"int256[]\",\"name\":\"values\",\"type\":\"int256[]\"}],\"name\":\"subReplaceInt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"subReplaceUint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"subReplaceUint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// JsonStoreABI is the input ABI used to generate the binding from.
// Deprecated: Use JsonStoreMetaData.ABI instead.
var JsonStoreABI = JsonStoreMetaData.ABI

// JsonStore is an auto generated Go binding around an Ethereum contract.
type JsonStore struct {
	JsonStoreCaller     // Read-only binding to the contract
	JsonStoreTransactor // Write-only binding to the contract
	JsonStoreFilterer   // Log filterer for contract events
}

// JsonStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type JsonStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JsonStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JsonStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JsonStoreSession struct {
	Contract     *JsonStore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JsonStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JsonStoreCallerSession struct {
	Contract *JsonStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// JsonStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JsonStoreTransactorSession struct {
	Contract     *JsonStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// JsonStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type JsonStoreRaw struct {
	Contract *JsonStore // Generic contract binding to access the raw methods on
}

// JsonStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JsonStoreCallerRaw struct {
	Contract *JsonStoreCaller // Generic read-only contract binding to access the raw methods on
}

// JsonStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JsonStoreTransactorRaw struct {
	Contract *JsonStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJsonStore creates a new instance of JsonStore, bound to a specific deployed contract.
func NewJsonStore(address common.Address, backend bind.ContractBackend) (*JsonStore, error) {
	contract, err := bindJsonStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JsonStore{JsonStoreCaller: JsonStoreCaller{contract: contract}, JsonStoreTransactor: JsonStoreTransactor{contract: contract}, JsonStoreFilterer: JsonStoreFilterer{contract: contract}}, nil
}

// NewJsonStoreCaller creates a new read-only instance of JsonStore, bound to a specific deployed contract.
func NewJsonStoreCaller(address common.Address, caller bind.ContractCaller) (*JsonStoreCaller, error) {
	contract, err := bindJsonStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JsonStoreCaller{contract: contract}, nil
}

// NewJsonStoreTransactor creates a new write-only instance of JsonStore, bound to a specific deployed contract.
func NewJsonStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*JsonStoreTransactor, error) {
	contract, err := bindJsonStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JsonStoreTransactor{contract: contract}, nil
}

// NewJsonStoreFilterer creates a new log filterer instance of JsonStore, bound to a specific deployed contract.
func NewJsonStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*JsonStoreFilterer, error) {
	contract, err := bindJsonStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JsonStoreFilterer{contract: contract}, nil
}

// bindJsonStore binds a generic wrapper to an already deployed contract.
func bindJsonStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JsonStoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JsonStore *JsonStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JsonStore.Contract.JsonStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JsonStore *JsonStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JsonStore.Contract.JsonStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JsonStore *JsonStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JsonStore.Contract.JsonStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JsonStore *JsonStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JsonStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JsonStore *JsonStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JsonStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JsonStore *JsonStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JsonStore.Contract.contract.Transact(opts, method, params...)
}

// DataURI is a free data retrieval call binding the contract method 0xda9e7501.
//
// Solidity: function dataURI(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreCaller) DataURI(opts *bind.CallOpts, key common.Address, slot *big.Int) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "dataURI", key, slot)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataURI is a free data retrieval call binding the contract method 0xda9e7501.
//
// Solidity: function dataURI(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreSession) DataURI(key common.Address, slot *big.Int) (string, error) {
	return _JsonStore.Contract.DataURI(&_JsonStore.CallOpts, key, slot)
}

// DataURI is a free data retrieval call binding the contract method 0xda9e7501.
//
// Solidity: function dataURI(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreCallerSession) DataURI(key common.Address, slot *big.Int) (string, error) {
	return _JsonStore.Contract.DataURI(&_JsonStore.CallOpts, key, slot)
}

// Exists is a free data retrieval call binding the contract method 0xaa1681a7.
//
// Solidity: function exists(address key, uint256 slot, string path) view returns(bool)
func (_JsonStore *JsonStoreCaller) Exists(opts *bind.CallOpts, key common.Address, slot *big.Int, path string) (bool, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "exists", key, slot, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xaa1681a7.
//
// Solidity: function exists(address key, uint256 slot, string path) view returns(bool)
func (_JsonStore *JsonStoreSession) Exists(key common.Address, slot *big.Int, path string) (bool, error) {
	return _JsonStore.Contract.Exists(&_JsonStore.CallOpts, key, slot, path)
}

// Exists is a free data retrieval call binding the contract method 0xaa1681a7.
//
// Solidity: function exists(address key, uint256 slot, string path) view returns(bool)
func (_JsonStore *JsonStoreCallerSession) Exists(key common.Address, slot *big.Int, path string) (bool, error) {
	return _JsonStore.Contract.Exists(&_JsonStore.CallOpts, key, slot, path)
}

// Get is a free data retrieval call binding the contract method 0x0334404d.
//
// Solidity: function get(address key, uint256 slot, string path) view returns(string)
func (_JsonStore *JsonStoreCaller) Get(opts *bind.CallOpts, key common.Address, slot *big.Int, path string) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "get", key, slot, path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x0334404d.
//
// Solidity: function get(address key, uint256 slot, string path) view returns(string)
func (_JsonStore *JsonStoreSession) Get(key common.Address, slot *big.Int, path string) (string, error) {
	return _JsonStore.Contract.Get(&_JsonStore.CallOpts, key, slot, path)
}

// Get is a free data retrieval call binding the contract method 0x0334404d.
//
// Solidity: function get(address key, uint256 slot, string path) view returns(string)
func (_JsonStore *JsonStoreCallerSession) Get(key common.Address, slot *big.Int, path string) (string, error) {
	return _JsonStore.Contract.Get(&_JsonStore.CallOpts, key, slot, path)
}

// Get0 is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreCaller) Get0(opts *bind.CallOpts, key common.Address, slot *big.Int) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "get0", key, slot)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get0 is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreSession) Get0(key common.Address, slot *big.Int) (string, error) {
	return _JsonStore.Contract.Get0(&_JsonStore.CallOpts, key, slot)
}

// Get0 is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreCallerSession) Get0(key common.Address, slot *big.Int) (string, error) {
	return _JsonStore.Contract.Get0(&_JsonStore.CallOpts, key, slot)
}

// GetBool is a free data retrieval call binding the contract method 0x577cbfc7.
//
// Solidity: function getBool(address key, uint256 slot, string path) view returns(bool)
func (_JsonStore *JsonStoreCaller) GetBool(opts *bind.CallOpts, key common.Address, slot *big.Int, path string) (bool, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "getBool", key, slot, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x577cbfc7.
//
// Solidity: function getBool(address key, uint256 slot, string path) view returns(bool)
func (_JsonStore *JsonStoreSession) GetBool(key common.Address, slot *big.Int, path string) (bool, error) {
	return _JsonStore.Contract.GetBool(&_JsonStore.CallOpts, key, slot, path)
}

// GetBool is a free data retrieval call binding the contract method 0x577cbfc7.
//
// Solidity: function getBool(address key, uint256 slot, string path) view returns(bool)
func (_JsonStore *JsonStoreCallerSession) GetBool(key common.Address, slot *big.Int, path string) (bool, error) {
	return _JsonStore.Contract.GetBool(&_JsonStore.CallOpts, key, slot, path)
}

// GetInt is a free data retrieval call binding the contract method 0x47a4d6da.
//
// Solidity: function getInt(address key, uint256 slot, string path) view returns(int256)
func (_JsonStore *JsonStoreCaller) GetInt(opts *bind.CallOpts, key common.Address, slot *big.Int, path string) (*big.Int, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "getInt", key, slot, path)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInt is a free data retrieval call binding the contract method 0x47a4d6da.
//
// Solidity: function getInt(address key, uint256 slot, string path) view returns(int256)
func (_JsonStore *JsonStoreSession) GetInt(key common.Address, slot *big.Int, path string) (*big.Int, error) {
	return _JsonStore.Contract.GetInt(&_JsonStore.CallOpts, key, slot, path)
}

// GetInt is a free data retrieval call binding the contract method 0x47a4d6da.
//
// Solidity: function getInt(address key, uint256 slot, string path) view returns(int256)
func (_JsonStore *JsonStoreCallerSession) GetInt(key common.Address, slot *big.Int, path string) (*big.Int, error) {
	return _JsonStore.Contract.GetInt(&_JsonStore.CallOpts, key, slot, path)
}

// GetUint is a free data retrieval call binding the contract method 0xf853705c.
//
// Solidity: function getUint(address key, uint256 slot, string path) view returns(uint256)
func (_JsonStore *JsonStoreCaller) GetUint(opts *bind.CallOpts, key common.Address, slot *big.Int, path string) (*big.Int, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "getUint", key, slot, path)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0xf853705c.
//
// Solidity: function getUint(address key, uint256 slot, string path) view returns(uint256)
func (_JsonStore *JsonStoreSession) GetUint(key common.Address, slot *big.Int, path string) (*big.Int, error) {
	return _JsonStore.Contract.GetUint(&_JsonStore.CallOpts, key, slot, path)
}

// GetUint is a free data retrieval call binding the contract method 0xf853705c.
//
// Solidity: function getUint(address key, uint256 slot, string path) view returns(uint256)
func (_JsonStore *JsonStoreCallerSession) GetUint(key common.Address, slot *big.Int, path string) (*big.Int, error) {
	return _JsonStore.Contract.GetUint(&_JsonStore.CallOpts, key, slot, path)
}

// Remove is a paid mutator transaction binding the contract method 0x9939b3a2.
//
// Solidity: function remove(uint256 slot, string path) returns(bool)
func (_JsonStore *JsonStoreTransactor) Remove(opts *bind.TransactOpts, slot *big.Int, path string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "remove", slot, path)
}

// Remove is a paid mutator transaction binding the contract method 0x9939b3a2.
//
// Solidity: function remove(uint256 slot, string path) returns(bool)
func (_JsonStore *JsonStoreSession) Remove(slot *big.Int, path string) (*types.Transaction, error) {
	return _JsonStore.Contract.Remove(&_JsonStore.TransactOpts, slot, path)
}

// Remove is a paid mutator transaction binding the contract method 0x9939b3a2.
//
// Solidity: function remove(uint256 slot, string path) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) Remove(slot *big.Int, path string) (*types.Transaction, error) {
	return _JsonStore.Contract.Remove(&_JsonStore.TransactOpts, slot, path)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 slot, string jsonBlob) returns(bool)
func (_JsonStore *JsonStoreTransactor) Set(opts *bind.TransactOpts, slot *big.Int, jsonBlob string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "set", slot, jsonBlob)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 slot, string jsonBlob) returns(bool)
func (_JsonStore *JsonStoreSession) Set(slot *big.Int, jsonBlob string) (*types.Transaction, error) {
	return _JsonStore.Contract.Set(&_JsonStore.TransactOpts, slot, jsonBlob)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 slot, string jsonBlob) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) Set(slot *big.Int, jsonBlob string) (*types.Transaction, error) {
	return _JsonStore.Contract.Set(&_JsonStore.TransactOpts, slot, jsonBlob)
}

// Set0 is a paid mutator transaction binding the contract method 0x65560950.
//
// Solidity: function set(uint256 slot, string path, string value) returns(bool)
func (_JsonStore *JsonStoreTransactor) Set0(opts *bind.TransactOpts, slot *big.Int, path string, value string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "set0", slot, path, value)
}

// Set0 is a paid mutator transaction binding the contract method 0x65560950.
//
// Solidity: function set(uint256 slot, string path, string value) returns(bool)
func (_JsonStore *JsonStoreSession) Set0(slot *big.Int, path string, value string) (*types.Transaction, error) {
	return _JsonStore.Contract.Set0(&_JsonStore.TransactOpts, slot, path, value)
}

// Set0 is a paid mutator transaction binding the contract method 0x65560950.
//
// Solidity: function set(uint256 slot, string path, string value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) Set0(slot *big.Int, path string, value string) (*types.Transaction, error) {
	return _JsonStore.Contract.Set0(&_JsonStore.TransactOpts, slot, path, value)
}

// Set1 is a paid mutator transaction binding the contract method 0xb1e0da6d.
//
// Solidity: function set(uint256 slot, string[] paths, string[] values) returns(bool)
func (_JsonStore *JsonStoreTransactor) Set1(opts *bind.TransactOpts, slot *big.Int, paths []string, values []string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "set1", slot, paths, values)
}

// Set1 is a paid mutator transaction binding the contract method 0xb1e0da6d.
//
// Solidity: function set(uint256 slot, string[] paths, string[] values) returns(bool)
func (_JsonStore *JsonStoreSession) Set1(slot *big.Int, paths []string, values []string) (*types.Transaction, error) {
	return _JsonStore.Contract.Set1(&_JsonStore.TransactOpts, slot, paths, values)
}

// Set1 is a paid mutator transaction binding the contract method 0xb1e0da6d.
//
// Solidity: function set(uint256 slot, string[] paths, string[] values) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) Set1(slot *big.Int, paths []string, values []string) (*types.Transaction, error) {
	return _JsonStore.Contract.Set1(&_JsonStore.TransactOpts, slot, paths, values)
}

// SetBool is a paid mutator transaction binding the contract method 0x8519e0e5.
//
// Solidity: function setBool(uint256 slot, string path, bool value) returns(bool)
func (_JsonStore *JsonStoreTransactor) SetBool(opts *bind.TransactOpts, slot *big.Int, path string, value bool) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "setBool", slot, path, value)
}

// SetBool is a paid mutator transaction binding the contract method 0x8519e0e5.
//
// Solidity: function setBool(uint256 slot, string path, bool value) returns(bool)
func (_JsonStore *JsonStoreSession) SetBool(slot *big.Int, path string, value bool) (*types.Transaction, error) {
	return _JsonStore.Contract.SetBool(&_JsonStore.TransactOpts, slot, path, value)
}

// SetBool is a paid mutator transaction binding the contract method 0x8519e0e5.
//
// Solidity: function setBool(uint256 slot, string path, bool value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SetBool(slot *big.Int, path string, value bool) (*types.Transaction, error) {
	return _JsonStore.Contract.SetBool(&_JsonStore.TransactOpts, slot, path, value)
}

// SetBool0 is a paid mutator transaction binding the contract method 0xf30dd855.
//
// Solidity: function setBool(uint256 slot, string[] paths, bool[] values) returns(bool)
func (_JsonStore *JsonStoreTransactor) SetBool0(opts *bind.TransactOpts, slot *big.Int, paths []string, values []bool) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "setBool0", slot, paths, values)
}

// SetBool0 is a paid mutator transaction binding the contract method 0xf30dd855.
//
// Solidity: function setBool(uint256 slot, string[] paths, bool[] values) returns(bool)
func (_JsonStore *JsonStoreSession) SetBool0(slot *big.Int, paths []string, values []bool) (*types.Transaction, error) {
	return _JsonStore.Contract.SetBool0(&_JsonStore.TransactOpts, slot, paths, values)
}

// SetBool0 is a paid mutator transaction binding the contract method 0xf30dd855.
//
// Solidity: function setBool(uint256 slot, string[] paths, bool[] values) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SetBool0(slot *big.Int, paths []string, values []bool) (*types.Transaction, error) {
	return _JsonStore.Contract.SetBool0(&_JsonStore.TransactOpts, slot, paths, values)
}

// SetInt is a paid mutator transaction binding the contract method 0xc297d59a.
//
// Solidity: function setInt(uint256 slot, string[] paths, int256[] values) returns(bool)
func (_JsonStore *JsonStoreTransactor) SetInt(opts *bind.TransactOpts, slot *big.Int, paths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "setInt", slot, paths, values)
}

// SetInt is a paid mutator transaction binding the contract method 0xc297d59a.
//
// Solidity: function setInt(uint256 slot, string[] paths, int256[] values) returns(bool)
func (_JsonStore *JsonStoreSession) SetInt(slot *big.Int, paths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SetInt(&_JsonStore.TransactOpts, slot, paths, values)
}

// SetInt is a paid mutator transaction binding the contract method 0xc297d59a.
//
// Solidity: function setInt(uint256 slot, string[] paths, int256[] values) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SetInt(slot *big.Int, paths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SetInt(&_JsonStore.TransactOpts, slot, paths, values)
}

// SetInt0 is a paid mutator transaction binding the contract method 0xf10ff226.
//
// Solidity: function setInt(uint256 slot, string path, int256 value) returns(bool)
func (_JsonStore *JsonStoreTransactor) SetInt0(opts *bind.TransactOpts, slot *big.Int, path string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "setInt0", slot, path, value)
}

// SetInt0 is a paid mutator transaction binding the contract method 0xf10ff226.
//
// Solidity: function setInt(uint256 slot, string path, int256 value) returns(bool)
func (_JsonStore *JsonStoreSession) SetInt0(slot *big.Int, path string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SetInt0(&_JsonStore.TransactOpts, slot, path, value)
}

// SetInt0 is a paid mutator transaction binding the contract method 0xf10ff226.
//
// Solidity: function setInt(uint256 slot, string path, int256 value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SetInt0(slot *big.Int, path string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SetInt0(&_JsonStore.TransactOpts, slot, path, value)
}

// SetRaw is a paid mutator transaction binding the contract method 0x1a20af47.
//
// Solidity: function setRaw(uint256 slot, string path, string rawBlob) returns(bool)
func (_JsonStore *JsonStoreTransactor) SetRaw(opts *bind.TransactOpts, slot *big.Int, path string, rawBlob string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "setRaw", slot, path, rawBlob)
}

// SetRaw is a paid mutator transaction binding the contract method 0x1a20af47.
//
// Solidity: function setRaw(uint256 slot, string path, string rawBlob) returns(bool)
func (_JsonStore *JsonStoreSession) SetRaw(slot *big.Int, path string, rawBlob string) (*types.Transaction, error) {
	return _JsonStore.Contract.SetRaw(&_JsonStore.TransactOpts, slot, path, rawBlob)
}

// SetRaw is a paid mutator transaction binding the contract method 0x1a20af47.
//
// Solidity: function setRaw(uint256 slot, string path, string rawBlob) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SetRaw(slot *big.Int, path string, rawBlob string) (*types.Transaction, error) {
	return _JsonStore.Contract.SetRaw(&_JsonStore.TransactOpts, slot, path, rawBlob)
}

// SetRaw0 is a paid mutator transaction binding the contract method 0x64aa7eb1.
//
// Solidity: function setRaw(uint256 slot, string[] paths, string[] rawBlobs) returns(bool)
func (_JsonStore *JsonStoreTransactor) SetRaw0(opts *bind.TransactOpts, slot *big.Int, paths []string, rawBlobs []string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "setRaw0", slot, paths, rawBlobs)
}

// SetRaw0 is a paid mutator transaction binding the contract method 0x64aa7eb1.
//
// Solidity: function setRaw(uint256 slot, string[] paths, string[] rawBlobs) returns(bool)
func (_JsonStore *JsonStoreSession) SetRaw0(slot *big.Int, paths []string, rawBlobs []string) (*types.Transaction, error) {
	return _JsonStore.Contract.SetRaw0(&_JsonStore.TransactOpts, slot, paths, rawBlobs)
}

// SetRaw0 is a paid mutator transaction binding the contract method 0x64aa7eb1.
//
// Solidity: function setRaw(uint256 slot, string[] paths, string[] rawBlobs) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SetRaw0(slot *big.Int, paths []string, rawBlobs []string) (*types.Transaction, error) {
	return _JsonStore.Contract.SetRaw0(&_JsonStore.TransactOpts, slot, paths, rawBlobs)
}

// SetUint is a paid mutator transaction binding the contract method 0x1a294636.
//
// Solidity: function setUint(uint256 slot, string[] paths, uint256[] values) returns(bool)
func (_JsonStore *JsonStoreTransactor) SetUint(opts *bind.TransactOpts, slot *big.Int, paths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "setUint", slot, paths, values)
}

// SetUint is a paid mutator transaction binding the contract method 0x1a294636.
//
// Solidity: function setUint(uint256 slot, string[] paths, uint256[] values) returns(bool)
func (_JsonStore *JsonStoreSession) SetUint(slot *big.Int, paths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SetUint(&_JsonStore.TransactOpts, slot, paths, values)
}

// SetUint is a paid mutator transaction binding the contract method 0x1a294636.
//
// Solidity: function setUint(uint256 slot, string[] paths, uint256[] values) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SetUint(slot *big.Int, paths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SetUint(&_JsonStore.TransactOpts, slot, paths, values)
}

// SetUint0 is a paid mutator transaction binding the contract method 0x711295fc.
//
// Solidity: function setUint(uint256 slot, string path, uint256 value) returns(bool)
func (_JsonStore *JsonStoreTransactor) SetUint0(opts *bind.TransactOpts, slot *big.Int, path string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "setUint0", slot, path, value)
}

// SetUint0 is a paid mutator transaction binding the contract method 0x711295fc.
//
// Solidity: function setUint(uint256 slot, string path, uint256 value) returns(bool)
func (_JsonStore *JsonStoreSession) SetUint0(slot *big.Int, path string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SetUint0(&_JsonStore.TransactOpts, slot, path, value)
}

// SetUint0 is a paid mutator transaction binding the contract method 0x711295fc.
//
// Solidity: function setUint(uint256 slot, string path, uint256 value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SetUint0(slot *big.Int, path string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SetUint0(&_JsonStore.TransactOpts, slot, path, value)
}

// SubReplace is a paid mutator transaction binding the contract method 0x03af1afa.
//
// Solidity: function subReplace(uint256 slot, string searchPath, string[] replacePaths, string[] values) returns(bool)
func (_JsonStore *JsonStoreTransactor) SubReplace(opts *bind.TransactOpts, slot *big.Int, searchPath string, replacePaths []string, values []string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "subReplace", slot, searchPath, replacePaths, values)
}

// SubReplace is a paid mutator transaction binding the contract method 0x03af1afa.
//
// Solidity: function subReplace(uint256 slot, string searchPath, string[] replacePaths, string[] values) returns(bool)
func (_JsonStore *JsonStoreSession) SubReplace(slot *big.Int, searchPath string, replacePaths []string, values []string) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplace(&_JsonStore.TransactOpts, slot, searchPath, replacePaths, values)
}

// SubReplace is a paid mutator transaction binding the contract method 0x03af1afa.
//
// Solidity: function subReplace(uint256 slot, string searchPath, string[] replacePaths, string[] values) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SubReplace(slot *big.Int, searchPath string, replacePaths []string, values []string) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplace(&_JsonStore.TransactOpts, slot, searchPath, replacePaths, values)
}

// SubReplace0 is a paid mutator transaction binding the contract method 0x40f95e6c.
//
// Solidity: function subReplace(uint256 slot, string searchPath, string replacePath, string value) returns(bool)
func (_JsonStore *JsonStoreTransactor) SubReplace0(opts *bind.TransactOpts, slot *big.Int, searchPath string, replacePath string, value string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "subReplace0", slot, searchPath, replacePath, value)
}

// SubReplace0 is a paid mutator transaction binding the contract method 0x40f95e6c.
//
// Solidity: function subReplace(uint256 slot, string searchPath, string replacePath, string value) returns(bool)
func (_JsonStore *JsonStoreSession) SubReplace0(slot *big.Int, searchPath string, replacePath string, value string) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplace0(&_JsonStore.TransactOpts, slot, searchPath, replacePath, value)
}

// SubReplace0 is a paid mutator transaction binding the contract method 0x40f95e6c.
//
// Solidity: function subReplace(uint256 slot, string searchPath, string replacePath, string value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SubReplace0(slot *big.Int, searchPath string, replacePath string, value string) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplace0(&_JsonStore.TransactOpts, slot, searchPath, replacePath, value)
}

// SubReplaceBool is a paid mutator transaction binding the contract method 0x378f3a6c.
//
// Solidity: function subReplaceBool(uint256 slot, string searchPath, string replacePath, bool value) returns(bool)
func (_JsonStore *JsonStoreTransactor) SubReplaceBool(opts *bind.TransactOpts, slot *big.Int, searchPath string, replacePath string, value bool) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "subReplaceBool", slot, searchPath, replacePath, value)
}

// SubReplaceBool is a paid mutator transaction binding the contract method 0x378f3a6c.
//
// Solidity: function subReplaceBool(uint256 slot, string searchPath, string replacePath, bool value) returns(bool)
func (_JsonStore *JsonStoreSession) SubReplaceBool(slot *big.Int, searchPath string, replacePath string, value bool) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceBool(&_JsonStore.TransactOpts, slot, searchPath, replacePath, value)
}

// SubReplaceBool is a paid mutator transaction binding the contract method 0x378f3a6c.
//
// Solidity: function subReplaceBool(uint256 slot, string searchPath, string replacePath, bool value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SubReplaceBool(slot *big.Int, searchPath string, replacePath string, value bool) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceBool(&_JsonStore.TransactOpts, slot, searchPath, replacePath, value)
}

// SubReplaceBool0 is a paid mutator transaction binding the contract method 0xfd0689f3.
//
// Solidity: function subReplaceBool(uint256 slot, string searchPath, string[] replacePaths, bool[] values) returns(bool)
func (_JsonStore *JsonStoreTransactor) SubReplaceBool0(opts *bind.TransactOpts, slot *big.Int, searchPath string, replacePaths []string, values []bool) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "subReplaceBool0", slot, searchPath, replacePaths, values)
}

// SubReplaceBool0 is a paid mutator transaction binding the contract method 0xfd0689f3.
//
// Solidity: function subReplaceBool(uint256 slot, string searchPath, string[] replacePaths, bool[] values) returns(bool)
func (_JsonStore *JsonStoreSession) SubReplaceBool0(slot *big.Int, searchPath string, replacePaths []string, values []bool) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceBool0(&_JsonStore.TransactOpts, slot, searchPath, replacePaths, values)
}

// SubReplaceBool0 is a paid mutator transaction binding the contract method 0xfd0689f3.
//
// Solidity: function subReplaceBool(uint256 slot, string searchPath, string[] replacePaths, bool[] values) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SubReplaceBool0(slot *big.Int, searchPath string, replacePaths []string, values []bool) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceBool0(&_JsonStore.TransactOpts, slot, searchPath, replacePaths, values)
}

// SubReplaceInt is a paid mutator transaction binding the contract method 0x5d97110c.
//
// Solidity: function subReplaceInt(uint256 slot, string searchPath, string replacePath, int256 value) returns(bool)
func (_JsonStore *JsonStoreTransactor) SubReplaceInt(opts *bind.TransactOpts, slot *big.Int, searchPath string, replacePath string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "subReplaceInt", slot, searchPath, replacePath, value)
}

// SubReplaceInt is a paid mutator transaction binding the contract method 0x5d97110c.
//
// Solidity: function subReplaceInt(uint256 slot, string searchPath, string replacePath, int256 value) returns(bool)
func (_JsonStore *JsonStoreSession) SubReplaceInt(slot *big.Int, searchPath string, replacePath string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceInt(&_JsonStore.TransactOpts, slot, searchPath, replacePath, value)
}

// SubReplaceInt is a paid mutator transaction binding the contract method 0x5d97110c.
//
// Solidity: function subReplaceInt(uint256 slot, string searchPath, string replacePath, int256 value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SubReplaceInt(slot *big.Int, searchPath string, replacePath string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceInt(&_JsonStore.TransactOpts, slot, searchPath, replacePath, value)
}

// SubReplaceInt0 is a paid mutator transaction binding the contract method 0x7f3aac41.
//
// Solidity: function subReplaceInt(uint256 slot, string searchPath, string[] replacePaths, int256[] values) returns(bool)
func (_JsonStore *JsonStoreTransactor) SubReplaceInt0(opts *bind.TransactOpts, slot *big.Int, searchPath string, replacePaths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "subReplaceInt0", slot, searchPath, replacePaths, values)
}

// SubReplaceInt0 is a paid mutator transaction binding the contract method 0x7f3aac41.
//
// Solidity: function subReplaceInt(uint256 slot, string searchPath, string[] replacePaths, int256[] values) returns(bool)
func (_JsonStore *JsonStoreSession) SubReplaceInt0(slot *big.Int, searchPath string, replacePaths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceInt0(&_JsonStore.TransactOpts, slot, searchPath, replacePaths, values)
}

// SubReplaceInt0 is a paid mutator transaction binding the contract method 0x7f3aac41.
//
// Solidity: function subReplaceInt(uint256 slot, string searchPath, string[] replacePaths, int256[] values) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SubReplaceInt0(slot *big.Int, searchPath string, replacePaths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceInt0(&_JsonStore.TransactOpts, slot, searchPath, replacePaths, values)
}

// SubReplaceUint is a paid mutator transaction binding the contract method 0x095a7a8b.
//
// Solidity: function subReplaceUint(uint256 slot, string searchPath, string replacePath, uint256 value) returns(bool)
func (_JsonStore *JsonStoreTransactor) SubReplaceUint(opts *bind.TransactOpts, slot *big.Int, searchPath string, replacePath string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "subReplaceUint", slot, searchPath, replacePath, value)
}

// SubReplaceUint is a paid mutator transaction binding the contract method 0x095a7a8b.
//
// Solidity: function subReplaceUint(uint256 slot, string searchPath, string replacePath, uint256 value) returns(bool)
func (_JsonStore *JsonStoreSession) SubReplaceUint(slot *big.Int, searchPath string, replacePath string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceUint(&_JsonStore.TransactOpts, slot, searchPath, replacePath, value)
}

// SubReplaceUint is a paid mutator transaction binding the contract method 0x095a7a8b.
//
// Solidity: function subReplaceUint(uint256 slot, string searchPath, string replacePath, uint256 value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SubReplaceUint(slot *big.Int, searchPath string, replacePath string, value *big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceUint(&_JsonStore.TransactOpts, slot, searchPath, replacePath, value)
}

// SubReplaceUint0 is a paid mutator transaction binding the contract method 0x0cd72795.
//
// Solidity: function subReplaceUint(uint256 slot, string searchPath, string[] replacePaths, uint256[] values) returns(bool)
func (_JsonStore *JsonStoreTransactor) SubReplaceUint0(opts *bind.TransactOpts, slot *big.Int, searchPath string, replacePaths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "subReplaceUint0", slot, searchPath, replacePaths, values)
}

// SubReplaceUint0 is a paid mutator transaction binding the contract method 0x0cd72795.
//
// Solidity: function subReplaceUint(uint256 slot, string searchPath, string[] replacePaths, uint256[] values) returns(bool)
func (_JsonStore *JsonStoreSession) SubReplaceUint0(slot *big.Int, searchPath string, replacePaths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceUint0(&_JsonStore.TransactOpts, slot, searchPath, replacePaths, values)
}

// SubReplaceUint0 is a paid mutator transaction binding the contract method 0x0cd72795.
//
// Solidity: function subReplaceUint(uint256 slot, string searchPath, string[] replacePaths, uint256[] values) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) SubReplaceUint0(slot *big.Int, searchPath string, replacePaths []string, values []*big.Int) (*types.Transaction, error) {
	return _JsonStore.Contract.SubReplaceUint0(&_JsonStore.TransactOpts, slot, searchPath, replacePaths, values)
}
