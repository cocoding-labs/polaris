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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"}],\"name\":\"dataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"dataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"key\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"path\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"value\",\"type\":\"string[]\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"subReplace\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"subReplace\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"subReplace\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"subReplace\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// DataURI is a free data retrieval call binding the contract method 0x08d1cf58.
//
// Solidity: function dataURI(string jsonBlob) view returns(string)
func (_JsonStore *JsonStoreCaller) DataURI(opts *bind.CallOpts, jsonBlob string) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "dataURI", jsonBlob)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataURI is a free data retrieval call binding the contract method 0x08d1cf58.
//
// Solidity: function dataURI(string jsonBlob) view returns(string)
func (_JsonStore *JsonStoreSession) DataURI(jsonBlob string) (string, error) {
	return _JsonStore.Contract.DataURI(&_JsonStore.CallOpts, jsonBlob)
}

// DataURI is a free data retrieval call binding the contract method 0x08d1cf58.
//
// Solidity: function dataURI(string jsonBlob) view returns(string)
func (_JsonStore *JsonStoreCallerSession) DataURI(jsonBlob string) (string, error) {
	return _JsonStore.Contract.DataURI(&_JsonStore.CallOpts, jsonBlob)
}

// DataURI0 is a free data retrieval call binding the contract method 0xda9e7501.
//
// Solidity: function dataURI(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreCaller) DataURI0(opts *bind.CallOpts, key common.Address, slot *big.Int) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "dataURI0", key, slot)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataURI0 is a free data retrieval call binding the contract method 0xda9e7501.
//
// Solidity: function dataURI(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreSession) DataURI0(key common.Address, slot *big.Int) (string, error) {
	return _JsonStore.Contract.DataURI0(&_JsonStore.CallOpts, key, slot)
}

// DataURI0 is a free data retrieval call binding the contract method 0xda9e7501.
//
// Solidity: function dataURI(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreCallerSession) DataURI0(key common.Address, slot *big.Int) (string, error) {
	return _JsonStore.Contract.DataURI0(&_JsonStore.CallOpts, key, slot)
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

// Get0 is a free data retrieval call binding the contract method 0x3e10510b.
//
// Solidity: function get(string jsonBlob, string path) view returns(string)
func (_JsonStore *JsonStoreCaller) Get0(opts *bind.CallOpts, jsonBlob string, path string) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "get0", jsonBlob, path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get0 is a free data retrieval call binding the contract method 0x3e10510b.
//
// Solidity: function get(string jsonBlob, string path) view returns(string)
func (_JsonStore *JsonStoreSession) Get0(jsonBlob string, path string) (string, error) {
	return _JsonStore.Contract.Get0(&_JsonStore.CallOpts, jsonBlob, path)
}

// Get0 is a free data retrieval call binding the contract method 0x3e10510b.
//
// Solidity: function get(string jsonBlob, string path) view returns(string)
func (_JsonStore *JsonStoreCallerSession) Get0(jsonBlob string, path string) (string, error) {
	return _JsonStore.Contract.Get0(&_JsonStore.CallOpts, jsonBlob, path)
}

// Get1 is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreCaller) Get1(opts *bind.CallOpts, key common.Address, slot *big.Int) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "get1", key, slot)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get1 is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreSession) Get1(key common.Address, slot *big.Int) (string, error) {
	return _JsonStore.Contract.Get1(&_JsonStore.CallOpts, key, slot)
}

// Get1 is a free data retrieval call binding the contract method 0xb464631b.
//
// Solidity: function get(address key, uint256 slot) view returns(string)
func (_JsonStore *JsonStoreCallerSession) Get1(key common.Address, slot *big.Int) (string, error) {
	return _JsonStore.Contract.Get1(&_JsonStore.CallOpts, key, slot)
}

// Remove is a free data retrieval call binding the contract method 0x44590a7e.
//
// Solidity: function remove(string jsonBlob, string path) view returns(string)
func (_JsonStore *JsonStoreCaller) Remove(opts *bind.CallOpts, jsonBlob string, path string) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "remove", jsonBlob, path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Remove is a free data retrieval call binding the contract method 0x44590a7e.
//
// Solidity: function remove(string jsonBlob, string path) view returns(string)
func (_JsonStore *JsonStoreSession) Remove(jsonBlob string, path string) (string, error) {
	return _JsonStore.Contract.Remove(&_JsonStore.CallOpts, jsonBlob, path)
}

// Remove is a free data retrieval call binding the contract method 0x44590a7e.
//
// Solidity: function remove(string jsonBlob, string path) view returns(string)
func (_JsonStore *JsonStoreCallerSession) Remove(jsonBlob string, path string) (string, error) {
	return _JsonStore.Contract.Remove(&_JsonStore.CallOpts, jsonBlob, path)
}

// Set1 is a free data retrieval call binding the contract method 0xa35a52ec.
//
// Solidity: function set(string jsonBlob, string[] path, string[] value) view returns(string)
func (_JsonStore *JsonStoreCaller) Set1(opts *bind.CallOpts, jsonBlob string, path []string, value []string) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "set1", jsonBlob, path, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Set1 is a free data retrieval call binding the contract method 0xa35a52ec.
//
// Solidity: function set(string jsonBlob, string[] path, string[] value) view returns(string)
func (_JsonStore *JsonStoreSession) Set1(jsonBlob string, path []string, value []string) (string, error) {
	return _JsonStore.Contract.Set1(&_JsonStore.CallOpts, jsonBlob, path, value)
}

// Set1 is a free data retrieval call binding the contract method 0xa35a52ec.
//
// Solidity: function set(string jsonBlob, string[] path, string[] value) view returns(string)
func (_JsonStore *JsonStoreCallerSession) Set1(jsonBlob string, path []string, value []string) (string, error) {
	return _JsonStore.Contract.Set1(&_JsonStore.CallOpts, jsonBlob, path, value)
}

// Set3 is a free data retrieval call binding the contract method 0xda465d74.
//
// Solidity: function set(string jsonBlob, string path, string value) view returns(string)
func (_JsonStore *JsonStoreCaller) Set3(opts *bind.CallOpts, jsonBlob string, path string, value string) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "set3", jsonBlob, path, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Set3 is a free data retrieval call binding the contract method 0xda465d74.
//
// Solidity: function set(string jsonBlob, string path, string value) view returns(string)
func (_JsonStore *JsonStoreSession) Set3(jsonBlob string, path string, value string) (string, error) {
	return _JsonStore.Contract.Set3(&_JsonStore.CallOpts, jsonBlob, path, value)
}

// Set3 is a free data retrieval call binding the contract method 0xda465d74.
//
// Solidity: function set(string jsonBlob, string path, string value) view returns(string)
func (_JsonStore *JsonStoreCallerSession) Set3(jsonBlob string, path string, value string) (string, error) {
	return _JsonStore.Contract.Set3(&_JsonStore.CallOpts, jsonBlob, path, value)
}

// SubReplace1 is a free data retrieval call binding the contract method 0xb27341bc.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string[] replacePaths, string[] values) view returns(string)
func (_JsonStore *JsonStoreCaller) SubReplace1(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePaths []string, values []string) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "subReplace1", jsonBlob, searchPath, replacePaths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplace1 is a free data retrieval call binding the contract method 0xb27341bc.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string[] replacePaths, string[] values) view returns(string)
func (_JsonStore *JsonStoreSession) SubReplace1(jsonBlob string, searchPath string, replacePaths []string, values []string) (string, error) {
	return _JsonStore.Contract.SubReplace1(&_JsonStore.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplace1 is a free data retrieval call binding the contract method 0xb27341bc.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string[] replacePaths, string[] values) view returns(string)
func (_JsonStore *JsonStoreCallerSession) SubReplace1(jsonBlob string, searchPath string, replacePaths []string, values []string) (string, error) {
	return _JsonStore.Contract.SubReplace1(&_JsonStore.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplace2 is a free data retrieval call binding the contract method 0xdb28035a.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string replacePath, string value) view returns(string)
func (_JsonStore *JsonStoreCaller) SubReplace2(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePath string, value string) (string, error) {
	var out []interface{}
	err := _JsonStore.contract.Call(opts, &out, "subReplace2", jsonBlob, searchPath, replacePath, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplace2 is a free data retrieval call binding the contract method 0xdb28035a.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string replacePath, string value) view returns(string)
func (_JsonStore *JsonStoreSession) SubReplace2(jsonBlob string, searchPath string, replacePath string, value string) (string, error) {
	return _JsonStore.Contract.SubReplace2(&_JsonStore.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplace2 is a free data retrieval call binding the contract method 0xdb28035a.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string replacePath, string value) view returns(string)
func (_JsonStore *JsonStoreCallerSession) SubReplace2(jsonBlob string, searchPath string, replacePath string, value string) (string, error) {
	return _JsonStore.Contract.SubReplace2(&_JsonStore.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// Remove0 is a paid mutator transaction binding the contract method 0x9939b3a2.
//
// Solidity: function remove(uint256 slot, string path) returns(bool)
func (_JsonStore *JsonStoreTransactor) Remove0(opts *bind.TransactOpts, slot *big.Int, path string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "remove0", slot, path)
}

// Remove0 is a paid mutator transaction binding the contract method 0x9939b3a2.
//
// Solidity: function remove(uint256 slot, string path) returns(bool)
func (_JsonStore *JsonStoreSession) Remove0(slot *big.Int, path string) (*types.Transaction, error) {
	return _JsonStore.Contract.Remove0(&_JsonStore.TransactOpts, slot, path)
}

// Remove0 is a paid mutator transaction binding the contract method 0x9939b3a2.
//
// Solidity: function remove(uint256 slot, string path) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) Remove0(slot *big.Int, path string) (*types.Transaction, error) {
	return _JsonStore.Contract.Remove0(&_JsonStore.TransactOpts, slot, path)
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

// Set2 is a paid mutator transaction binding the contract method 0xb1e0da6d.
//
// Solidity: function set(uint256 slot, string[] path, string[] value) returns(bool)
func (_JsonStore *JsonStoreTransactor) Set2(opts *bind.TransactOpts, slot *big.Int, path []string, value []string) (*types.Transaction, error) {
	return _JsonStore.contract.Transact(opts, "set2", slot, path, value)
}

// Set2 is a paid mutator transaction binding the contract method 0xb1e0da6d.
//
// Solidity: function set(uint256 slot, string[] path, string[] value) returns(bool)
func (_JsonStore *JsonStoreSession) Set2(slot *big.Int, path []string, value []string) (*types.Transaction, error) {
	return _JsonStore.Contract.Set2(&_JsonStore.TransactOpts, slot, path, value)
}

// Set2 is a paid mutator transaction binding the contract method 0xb1e0da6d.
//
// Solidity: function set(uint256 slot, string[] path, string[] value) returns(bool)
func (_JsonStore *JsonStoreTransactorSession) Set2(slot *big.Int, path []string, value []string) (*types.Transaction, error) {
	return _JsonStore.Contract.Set2(&_JsonStore.TransactOpts, slot, path, value)
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
