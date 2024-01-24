// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jsonutil

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

// JsonUtilMetaData contains all meta data concerning the JsonUtil contract.
var JsonUtilMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"}],\"name\":\"dataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"getBool\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"getInt\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"getUint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"values\",\"type\":\"bool[]\"}],\"name\":\"setBool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setBool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"int256[]\",\"name\":\"values\",\"type\":\"int256[]\"}],\"name\":\"setInt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"setInt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"rawBlob\",\"type\":\"string\"}],\"name\":\"setRaw\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"rawBlobs\",\"type\":\"string[]\"}],\"name\":\"setRaw\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"paths\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"setUint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"path\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setUint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"values\",\"type\":\"string[]\"}],\"name\":\"subReplace\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"subReplace\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"bool[]\",\"name\":\"values\",\"type\":\"bool[]\"}],\"name\":\"subReplaceBool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"subReplaceBool\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"subReplaceInt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"int256[]\",\"name\":\"values\",\"type\":\"int256[]\"}],\"name\":\"subReplaceInt\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replacePath\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"subReplaceUint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"jsonBlob\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"searchPath\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"replacePaths\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"subReplaceUint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// JsonUtilABI is the input ABI used to generate the binding from.
// Deprecated: Use JsonUtilMetaData.ABI instead.
var JsonUtilABI = JsonUtilMetaData.ABI

// JsonUtil is an auto generated Go binding around an Ethereum contract.
type JsonUtil struct {
	JsonUtilCaller     // Read-only binding to the contract
	JsonUtilTransactor // Write-only binding to the contract
	JsonUtilFilterer   // Log filterer for contract events
}

// JsonUtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type JsonUtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonUtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JsonUtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonUtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JsonUtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JsonUtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JsonUtilSession struct {
	Contract     *JsonUtil         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JsonUtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JsonUtilCallerSession struct {
	Contract *JsonUtilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// JsonUtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JsonUtilTransactorSession struct {
	Contract     *JsonUtilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// JsonUtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type JsonUtilRaw struct {
	Contract *JsonUtil // Generic contract binding to access the raw methods on
}

// JsonUtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JsonUtilCallerRaw struct {
	Contract *JsonUtilCaller // Generic read-only contract binding to access the raw methods on
}

// JsonUtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JsonUtilTransactorRaw struct {
	Contract *JsonUtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJsonUtil creates a new instance of JsonUtil, bound to a specific deployed contract.
func NewJsonUtil(address common.Address, backend bind.ContractBackend) (*JsonUtil, error) {
	contract, err := bindJsonUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JsonUtil{JsonUtilCaller: JsonUtilCaller{contract: contract}, JsonUtilTransactor: JsonUtilTransactor{contract: contract}, JsonUtilFilterer: JsonUtilFilterer{contract: contract}}, nil
}

// NewJsonUtilCaller creates a new read-only instance of JsonUtil, bound to a specific deployed contract.
func NewJsonUtilCaller(address common.Address, caller bind.ContractCaller) (*JsonUtilCaller, error) {
	contract, err := bindJsonUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JsonUtilCaller{contract: contract}, nil
}

// NewJsonUtilTransactor creates a new write-only instance of JsonUtil, bound to a specific deployed contract.
func NewJsonUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*JsonUtilTransactor, error) {
	contract, err := bindJsonUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JsonUtilTransactor{contract: contract}, nil
}

// NewJsonUtilFilterer creates a new log filterer instance of JsonUtil, bound to a specific deployed contract.
func NewJsonUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*JsonUtilFilterer, error) {
	contract, err := bindJsonUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JsonUtilFilterer{contract: contract}, nil
}

// bindJsonUtil binds a generic wrapper to an already deployed contract.
func bindJsonUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JsonUtilMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JsonUtil *JsonUtilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JsonUtil.Contract.JsonUtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JsonUtil *JsonUtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JsonUtil.Contract.JsonUtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JsonUtil *JsonUtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JsonUtil.Contract.JsonUtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JsonUtil *JsonUtilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JsonUtil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JsonUtil *JsonUtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JsonUtil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JsonUtil *JsonUtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JsonUtil.Contract.contract.Transact(opts, method, params...)
}

// DataURI is a free data retrieval call binding the contract method 0x08d1cf58.
//
// Solidity: function dataURI(string jsonBlob) view returns(string)
func (_JsonUtil *JsonUtilCaller) DataURI(opts *bind.CallOpts, jsonBlob string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "dataURI", jsonBlob)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataURI is a free data retrieval call binding the contract method 0x08d1cf58.
//
// Solidity: function dataURI(string jsonBlob) view returns(string)
func (_JsonUtil *JsonUtilSession) DataURI(jsonBlob string) (string, error) {
	return _JsonUtil.Contract.DataURI(&_JsonUtil.CallOpts, jsonBlob)
}

// DataURI is a free data retrieval call binding the contract method 0x08d1cf58.
//
// Solidity: function dataURI(string jsonBlob) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) DataURI(jsonBlob string) (string, error) {
	return _JsonUtil.Contract.DataURI(&_JsonUtil.CallOpts, jsonBlob)
}

// Exists is a free data retrieval call binding the contract method 0x2656554c.
//
// Solidity: function exists(string jsonBlob, string path) view returns(bool)
func (_JsonUtil *JsonUtilCaller) Exists(opts *bind.CallOpts, jsonBlob string, path string) (bool, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "exists", jsonBlob, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x2656554c.
//
// Solidity: function exists(string jsonBlob, string path) view returns(bool)
func (_JsonUtil *JsonUtilSession) Exists(jsonBlob string, path string) (bool, error) {
	return _JsonUtil.Contract.Exists(&_JsonUtil.CallOpts, jsonBlob, path)
}

// Exists is a free data retrieval call binding the contract method 0x2656554c.
//
// Solidity: function exists(string jsonBlob, string path) view returns(bool)
func (_JsonUtil *JsonUtilCallerSession) Exists(jsonBlob string, path string) (bool, error) {
	return _JsonUtil.Contract.Exists(&_JsonUtil.CallOpts, jsonBlob, path)
}

// Get is a free data retrieval call binding the contract method 0x3e10510b.
//
// Solidity: function get(string jsonBlob, string path) view returns(string)
func (_JsonUtil *JsonUtilCaller) Get(opts *bind.CallOpts, jsonBlob string, path string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "get", jsonBlob, path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x3e10510b.
//
// Solidity: function get(string jsonBlob, string path) view returns(string)
func (_JsonUtil *JsonUtilSession) Get(jsonBlob string, path string) (string, error) {
	return _JsonUtil.Contract.Get(&_JsonUtil.CallOpts, jsonBlob, path)
}

// Get is a free data retrieval call binding the contract method 0x3e10510b.
//
// Solidity: function get(string jsonBlob, string path) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) Get(jsonBlob string, path string) (string, error) {
	return _JsonUtil.Contract.Get(&_JsonUtil.CallOpts, jsonBlob, path)
}

// GetBool is a free data retrieval call binding the contract method 0x0586a6a5.
//
// Solidity: function getBool(string jsonBlob, string path) view returns(bool)
func (_JsonUtil *JsonUtilCaller) GetBool(opts *bind.CallOpts, jsonBlob string, path string) (bool, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "getBool", jsonBlob, path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x0586a6a5.
//
// Solidity: function getBool(string jsonBlob, string path) view returns(bool)
func (_JsonUtil *JsonUtilSession) GetBool(jsonBlob string, path string) (bool, error) {
	return _JsonUtil.Contract.GetBool(&_JsonUtil.CallOpts, jsonBlob, path)
}

// GetBool is a free data retrieval call binding the contract method 0x0586a6a5.
//
// Solidity: function getBool(string jsonBlob, string path) view returns(bool)
func (_JsonUtil *JsonUtilCallerSession) GetBool(jsonBlob string, path string) (bool, error) {
	return _JsonUtil.Contract.GetBool(&_JsonUtil.CallOpts, jsonBlob, path)
}

// GetInt is a free data retrieval call binding the contract method 0x4ab13a7d.
//
// Solidity: function getInt(string jsonBlob, string path) view returns(int256)
func (_JsonUtil *JsonUtilCaller) GetInt(opts *bind.CallOpts, jsonBlob string, path string) (*big.Int, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "getInt", jsonBlob, path)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInt is a free data retrieval call binding the contract method 0x4ab13a7d.
//
// Solidity: function getInt(string jsonBlob, string path) view returns(int256)
func (_JsonUtil *JsonUtilSession) GetInt(jsonBlob string, path string) (*big.Int, error) {
	return _JsonUtil.Contract.GetInt(&_JsonUtil.CallOpts, jsonBlob, path)
}

// GetInt is a free data retrieval call binding the contract method 0x4ab13a7d.
//
// Solidity: function getInt(string jsonBlob, string path) view returns(int256)
func (_JsonUtil *JsonUtilCallerSession) GetInt(jsonBlob string, path string) (*big.Int, error) {
	return _JsonUtil.Contract.GetInt(&_JsonUtil.CallOpts, jsonBlob, path)
}

// GetUint is a free data retrieval call binding the contract method 0x5df6d3cc.
//
// Solidity: function getUint(string jsonBlob, string path) view returns(uint256)
func (_JsonUtil *JsonUtilCaller) GetUint(opts *bind.CallOpts, jsonBlob string, path string) (*big.Int, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "getUint", jsonBlob, path)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0x5df6d3cc.
//
// Solidity: function getUint(string jsonBlob, string path) view returns(uint256)
func (_JsonUtil *JsonUtilSession) GetUint(jsonBlob string, path string) (*big.Int, error) {
	return _JsonUtil.Contract.GetUint(&_JsonUtil.CallOpts, jsonBlob, path)
}

// GetUint is a free data retrieval call binding the contract method 0x5df6d3cc.
//
// Solidity: function getUint(string jsonBlob, string path) view returns(uint256)
func (_JsonUtil *JsonUtilCallerSession) GetUint(jsonBlob string, path string) (*big.Int, error) {
	return _JsonUtil.Contract.GetUint(&_JsonUtil.CallOpts, jsonBlob, path)
}

// Remove is a free data retrieval call binding the contract method 0x44590a7e.
//
// Solidity: function remove(string jsonBlob, string path) view returns(string)
func (_JsonUtil *JsonUtilCaller) Remove(opts *bind.CallOpts, jsonBlob string, path string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "remove", jsonBlob, path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Remove is a free data retrieval call binding the contract method 0x44590a7e.
//
// Solidity: function remove(string jsonBlob, string path) view returns(string)
func (_JsonUtil *JsonUtilSession) Remove(jsonBlob string, path string) (string, error) {
	return _JsonUtil.Contract.Remove(&_JsonUtil.CallOpts, jsonBlob, path)
}

// Remove is a free data retrieval call binding the contract method 0x44590a7e.
//
// Solidity: function remove(string jsonBlob, string path) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) Remove(jsonBlob string, path string) (string, error) {
	return _JsonUtil.Contract.Remove(&_JsonUtil.CallOpts, jsonBlob, path)
}

// Set is a free data retrieval call binding the contract method 0xa35a52ec.
//
// Solidity: function set(string jsonBlob, string[] paths, string[] values) view returns(string)
func (_JsonUtil *JsonUtilCaller) Set(opts *bind.CallOpts, jsonBlob string, paths []string, values []string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "set", jsonBlob, paths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Set is a free data retrieval call binding the contract method 0xa35a52ec.
//
// Solidity: function set(string jsonBlob, string[] paths, string[] values) view returns(string)
func (_JsonUtil *JsonUtilSession) Set(jsonBlob string, paths []string, values []string) (string, error) {
	return _JsonUtil.Contract.Set(&_JsonUtil.CallOpts, jsonBlob, paths, values)
}

// Set is a free data retrieval call binding the contract method 0xa35a52ec.
//
// Solidity: function set(string jsonBlob, string[] paths, string[] values) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) Set(jsonBlob string, paths []string, values []string) (string, error) {
	return _JsonUtil.Contract.Set(&_JsonUtil.CallOpts, jsonBlob, paths, values)
}

// Set0 is a free data retrieval call binding the contract method 0xda465d74.
//
// Solidity: function set(string jsonBlob, string path, string value) view returns(string)
func (_JsonUtil *JsonUtilCaller) Set0(opts *bind.CallOpts, jsonBlob string, path string, value string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "set0", jsonBlob, path, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Set0 is a free data retrieval call binding the contract method 0xda465d74.
//
// Solidity: function set(string jsonBlob, string path, string value) view returns(string)
func (_JsonUtil *JsonUtilSession) Set0(jsonBlob string, path string, value string) (string, error) {
	return _JsonUtil.Contract.Set0(&_JsonUtil.CallOpts, jsonBlob, path, value)
}

// Set0 is a free data retrieval call binding the contract method 0xda465d74.
//
// Solidity: function set(string jsonBlob, string path, string value) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) Set0(jsonBlob string, path string, value string) (string, error) {
	return _JsonUtil.Contract.Set0(&_JsonUtil.CallOpts, jsonBlob, path, value)
}

// SetBool is a free data retrieval call binding the contract method 0xdf7218d1.
//
// Solidity: function setBool(string jsonBlob, string[] paths, bool[] values) view returns(string)
func (_JsonUtil *JsonUtilCaller) SetBool(opts *bind.CallOpts, jsonBlob string, paths []string, values []bool) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "setBool", jsonBlob, paths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SetBool is a free data retrieval call binding the contract method 0xdf7218d1.
//
// Solidity: function setBool(string jsonBlob, string[] paths, bool[] values) view returns(string)
func (_JsonUtil *JsonUtilSession) SetBool(jsonBlob string, paths []string, values []bool) (string, error) {
	return _JsonUtil.Contract.SetBool(&_JsonUtil.CallOpts, jsonBlob, paths, values)
}

// SetBool is a free data retrieval call binding the contract method 0xdf7218d1.
//
// Solidity: function setBool(string jsonBlob, string[] paths, bool[] values) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SetBool(jsonBlob string, paths []string, values []bool) (string, error) {
	return _JsonUtil.Contract.SetBool(&_JsonUtil.CallOpts, jsonBlob, paths, values)
}

// SetBool0 is a free data retrieval call binding the contract method 0xf057bd17.
//
// Solidity: function setBool(string jsonBlob, string path, bool value) view returns(string)
func (_JsonUtil *JsonUtilCaller) SetBool0(opts *bind.CallOpts, jsonBlob string, path string, value bool) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "setBool0", jsonBlob, path, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SetBool0 is a free data retrieval call binding the contract method 0xf057bd17.
//
// Solidity: function setBool(string jsonBlob, string path, bool value) view returns(string)
func (_JsonUtil *JsonUtilSession) SetBool0(jsonBlob string, path string, value bool) (string, error) {
	return _JsonUtil.Contract.SetBool0(&_JsonUtil.CallOpts, jsonBlob, path, value)
}

// SetBool0 is a free data retrieval call binding the contract method 0xf057bd17.
//
// Solidity: function setBool(string jsonBlob, string path, bool value) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SetBool0(jsonBlob string, path string, value bool) (string, error) {
	return _JsonUtil.Contract.SetBool0(&_JsonUtil.CallOpts, jsonBlob, path, value)
}

// SetInt is a free data retrieval call binding the contract method 0xbabe6ab7.
//
// Solidity: function setInt(string jsonBlob, string[] paths, int256[] values) view returns(string)
func (_JsonUtil *JsonUtilCaller) SetInt(opts *bind.CallOpts, jsonBlob string, paths []string, values []*big.Int) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "setInt", jsonBlob, paths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SetInt is a free data retrieval call binding the contract method 0xbabe6ab7.
//
// Solidity: function setInt(string jsonBlob, string[] paths, int256[] values) view returns(string)
func (_JsonUtil *JsonUtilSession) SetInt(jsonBlob string, paths []string, values []*big.Int) (string, error) {
	return _JsonUtil.Contract.SetInt(&_JsonUtil.CallOpts, jsonBlob, paths, values)
}

// SetInt is a free data retrieval call binding the contract method 0xbabe6ab7.
//
// Solidity: function setInt(string jsonBlob, string[] paths, int256[] values) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SetInt(jsonBlob string, paths []string, values []*big.Int) (string, error) {
	return _JsonUtil.Contract.SetInt(&_JsonUtil.CallOpts, jsonBlob, paths, values)
}

// SetInt0 is a free data retrieval call binding the contract method 0xf38487a6.
//
// Solidity: function setInt(string jsonBlob, string path, int256 value) view returns(string)
func (_JsonUtil *JsonUtilCaller) SetInt0(opts *bind.CallOpts, jsonBlob string, path string, value *big.Int) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "setInt0", jsonBlob, path, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SetInt0 is a free data retrieval call binding the contract method 0xf38487a6.
//
// Solidity: function setInt(string jsonBlob, string path, int256 value) view returns(string)
func (_JsonUtil *JsonUtilSession) SetInt0(jsonBlob string, path string, value *big.Int) (string, error) {
	return _JsonUtil.Contract.SetInt0(&_JsonUtil.CallOpts, jsonBlob, path, value)
}

// SetInt0 is a free data retrieval call binding the contract method 0xf38487a6.
//
// Solidity: function setInt(string jsonBlob, string path, int256 value) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SetInt0(jsonBlob string, path string, value *big.Int) (string, error) {
	return _JsonUtil.Contract.SetInt0(&_JsonUtil.CallOpts, jsonBlob, path, value)
}

// SetRaw is a free data retrieval call binding the contract method 0x9793f350.
//
// Solidity: function setRaw(string jsonBlob, string path, string rawBlob) view returns(string)
func (_JsonUtil *JsonUtilCaller) SetRaw(opts *bind.CallOpts, jsonBlob string, path string, rawBlob string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "setRaw", jsonBlob, path, rawBlob)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SetRaw is a free data retrieval call binding the contract method 0x9793f350.
//
// Solidity: function setRaw(string jsonBlob, string path, string rawBlob) view returns(string)
func (_JsonUtil *JsonUtilSession) SetRaw(jsonBlob string, path string, rawBlob string) (string, error) {
	return _JsonUtil.Contract.SetRaw(&_JsonUtil.CallOpts, jsonBlob, path, rawBlob)
}

// SetRaw is a free data retrieval call binding the contract method 0x9793f350.
//
// Solidity: function setRaw(string jsonBlob, string path, string rawBlob) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SetRaw(jsonBlob string, path string, rawBlob string) (string, error) {
	return _JsonUtil.Contract.SetRaw(&_JsonUtil.CallOpts, jsonBlob, path, rawBlob)
}

// SetRaw0 is a free data retrieval call binding the contract method 0xf0cc5e13.
//
// Solidity: function setRaw(string jsonBlob, string[] paths, string[] rawBlobs) view returns(string)
func (_JsonUtil *JsonUtilCaller) SetRaw0(opts *bind.CallOpts, jsonBlob string, paths []string, rawBlobs []string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "setRaw0", jsonBlob, paths, rawBlobs)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SetRaw0 is a free data retrieval call binding the contract method 0xf0cc5e13.
//
// Solidity: function setRaw(string jsonBlob, string[] paths, string[] rawBlobs) view returns(string)
func (_JsonUtil *JsonUtilSession) SetRaw0(jsonBlob string, paths []string, rawBlobs []string) (string, error) {
	return _JsonUtil.Contract.SetRaw0(&_JsonUtil.CallOpts, jsonBlob, paths, rawBlobs)
}

// SetRaw0 is a free data retrieval call binding the contract method 0xf0cc5e13.
//
// Solidity: function setRaw(string jsonBlob, string[] paths, string[] rawBlobs) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SetRaw0(jsonBlob string, paths []string, rawBlobs []string) (string, error) {
	return _JsonUtil.Contract.SetRaw0(&_JsonUtil.CallOpts, jsonBlob, paths, rawBlobs)
}

// SetUint is a free data retrieval call binding the contract method 0x17ceb2be.
//
// Solidity: function setUint(string jsonBlob, string[] paths, uint256[] values) view returns(string)
func (_JsonUtil *JsonUtilCaller) SetUint(opts *bind.CallOpts, jsonBlob string, paths []string, values []*big.Int) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "setUint", jsonBlob, paths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SetUint is a free data retrieval call binding the contract method 0x17ceb2be.
//
// Solidity: function setUint(string jsonBlob, string[] paths, uint256[] values) view returns(string)
func (_JsonUtil *JsonUtilSession) SetUint(jsonBlob string, paths []string, values []*big.Int) (string, error) {
	return _JsonUtil.Contract.SetUint(&_JsonUtil.CallOpts, jsonBlob, paths, values)
}

// SetUint is a free data retrieval call binding the contract method 0x17ceb2be.
//
// Solidity: function setUint(string jsonBlob, string[] paths, uint256[] values) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SetUint(jsonBlob string, paths []string, values []*big.Int) (string, error) {
	return _JsonUtil.Contract.SetUint(&_JsonUtil.CallOpts, jsonBlob, paths, values)
}

// SetUint0 is a free data retrieval call binding the contract method 0x56b44b2c.
//
// Solidity: function setUint(string jsonBlob, string path, uint256 value) view returns(string)
func (_JsonUtil *JsonUtilCaller) SetUint0(opts *bind.CallOpts, jsonBlob string, path string, value *big.Int) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "setUint0", jsonBlob, path, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SetUint0 is a free data retrieval call binding the contract method 0x56b44b2c.
//
// Solidity: function setUint(string jsonBlob, string path, uint256 value) view returns(string)
func (_JsonUtil *JsonUtilSession) SetUint0(jsonBlob string, path string, value *big.Int) (string, error) {
	return _JsonUtil.Contract.SetUint0(&_JsonUtil.CallOpts, jsonBlob, path, value)
}

// SetUint0 is a free data retrieval call binding the contract method 0x56b44b2c.
//
// Solidity: function setUint(string jsonBlob, string path, uint256 value) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SetUint0(jsonBlob string, path string, value *big.Int) (string, error) {
	return _JsonUtil.Contract.SetUint0(&_JsonUtil.CallOpts, jsonBlob, path, value)
}

// SubReplace is a free data retrieval call binding the contract method 0xb27341bc.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string[] replacePaths, string[] values) view returns(string)
func (_JsonUtil *JsonUtilCaller) SubReplace(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePaths []string, values []string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "subReplace", jsonBlob, searchPath, replacePaths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplace is a free data retrieval call binding the contract method 0xb27341bc.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string[] replacePaths, string[] values) view returns(string)
func (_JsonUtil *JsonUtilSession) SubReplace(jsonBlob string, searchPath string, replacePaths []string, values []string) (string, error) {
	return _JsonUtil.Contract.SubReplace(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplace is a free data retrieval call binding the contract method 0xb27341bc.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string[] replacePaths, string[] values) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SubReplace(jsonBlob string, searchPath string, replacePaths []string, values []string) (string, error) {
	return _JsonUtil.Contract.SubReplace(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplace0 is a free data retrieval call binding the contract method 0xdb28035a.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string replacePath, string value) view returns(string)
func (_JsonUtil *JsonUtilCaller) SubReplace0(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePath string, value string) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "subReplace0", jsonBlob, searchPath, replacePath, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplace0 is a free data retrieval call binding the contract method 0xdb28035a.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string replacePath, string value) view returns(string)
func (_JsonUtil *JsonUtilSession) SubReplace0(jsonBlob string, searchPath string, replacePath string, value string) (string, error) {
	return _JsonUtil.Contract.SubReplace0(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplace0 is a free data retrieval call binding the contract method 0xdb28035a.
//
// Solidity: function subReplace(string jsonBlob, string searchPath, string replacePath, string value) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SubReplace0(jsonBlob string, searchPath string, replacePath string, value string) (string, error) {
	return _JsonUtil.Contract.SubReplace0(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplaceBool is a free data retrieval call binding the contract method 0x525ca75e.
//
// Solidity: function subReplaceBool(string jsonBlob, string searchPath, string[] replacePaths, bool[] values) view returns(string)
func (_JsonUtil *JsonUtilCaller) SubReplaceBool(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePaths []string, values []bool) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "subReplaceBool", jsonBlob, searchPath, replacePaths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplaceBool is a free data retrieval call binding the contract method 0x525ca75e.
//
// Solidity: function subReplaceBool(string jsonBlob, string searchPath, string[] replacePaths, bool[] values) view returns(string)
func (_JsonUtil *JsonUtilSession) SubReplaceBool(jsonBlob string, searchPath string, replacePaths []string, values []bool) (string, error) {
	return _JsonUtil.Contract.SubReplaceBool(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplaceBool is a free data retrieval call binding the contract method 0x525ca75e.
//
// Solidity: function subReplaceBool(string jsonBlob, string searchPath, string[] replacePaths, bool[] values) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SubReplaceBool(jsonBlob string, searchPath string, replacePaths []string, values []bool) (string, error) {
	return _JsonUtil.Contract.SubReplaceBool(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplaceBool0 is a free data retrieval call binding the contract method 0xdc782e37.
//
// Solidity: function subReplaceBool(string jsonBlob, string searchPath, string replacePath, bool value) view returns(string)
func (_JsonUtil *JsonUtilCaller) SubReplaceBool0(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePath string, value bool) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "subReplaceBool0", jsonBlob, searchPath, replacePath, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplaceBool0 is a free data retrieval call binding the contract method 0xdc782e37.
//
// Solidity: function subReplaceBool(string jsonBlob, string searchPath, string replacePath, bool value) view returns(string)
func (_JsonUtil *JsonUtilSession) SubReplaceBool0(jsonBlob string, searchPath string, replacePath string, value bool) (string, error) {
	return _JsonUtil.Contract.SubReplaceBool0(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplaceBool0 is a free data retrieval call binding the contract method 0xdc782e37.
//
// Solidity: function subReplaceBool(string jsonBlob, string searchPath, string replacePath, bool value) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SubReplaceBool0(jsonBlob string, searchPath string, replacePath string, value bool) (string, error) {
	return _JsonUtil.Contract.SubReplaceBool0(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplaceInt is a free data retrieval call binding the contract method 0x111dac0d.
//
// Solidity: function subReplaceInt(string jsonBlob, string searchPath, string replacePath, int256 value) view returns(string)
func (_JsonUtil *JsonUtilCaller) SubReplaceInt(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePath string, value *big.Int) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "subReplaceInt", jsonBlob, searchPath, replacePath, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplaceInt is a free data retrieval call binding the contract method 0x111dac0d.
//
// Solidity: function subReplaceInt(string jsonBlob, string searchPath, string replacePath, int256 value) view returns(string)
func (_JsonUtil *JsonUtilSession) SubReplaceInt(jsonBlob string, searchPath string, replacePath string, value *big.Int) (string, error) {
	return _JsonUtil.Contract.SubReplaceInt(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplaceInt is a free data retrieval call binding the contract method 0x111dac0d.
//
// Solidity: function subReplaceInt(string jsonBlob, string searchPath, string replacePath, int256 value) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SubReplaceInt(jsonBlob string, searchPath string, replacePath string, value *big.Int) (string, error) {
	return _JsonUtil.Contract.SubReplaceInt(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplaceInt0 is a free data retrieval call binding the contract method 0x98807da2.
//
// Solidity: function subReplaceInt(string jsonBlob, string searchPath, string[] replacePaths, int256[] values) view returns(string)
func (_JsonUtil *JsonUtilCaller) SubReplaceInt0(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePaths []string, values []*big.Int) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "subReplaceInt0", jsonBlob, searchPath, replacePaths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplaceInt0 is a free data retrieval call binding the contract method 0x98807da2.
//
// Solidity: function subReplaceInt(string jsonBlob, string searchPath, string[] replacePaths, int256[] values) view returns(string)
func (_JsonUtil *JsonUtilSession) SubReplaceInt0(jsonBlob string, searchPath string, replacePaths []string, values []*big.Int) (string, error) {
	return _JsonUtil.Contract.SubReplaceInt0(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplaceInt0 is a free data retrieval call binding the contract method 0x98807da2.
//
// Solidity: function subReplaceInt(string jsonBlob, string searchPath, string[] replacePaths, int256[] values) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SubReplaceInt0(jsonBlob string, searchPath string, replacePaths []string, values []*big.Int) (string, error) {
	return _JsonUtil.Contract.SubReplaceInt0(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplaceUint is a free data retrieval call binding the contract method 0xd37ee641.
//
// Solidity: function subReplaceUint(string jsonBlob, string searchPath, string replacePath, uint256 value) view returns(string)
func (_JsonUtil *JsonUtilCaller) SubReplaceUint(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePath string, value *big.Int) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "subReplaceUint", jsonBlob, searchPath, replacePath, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplaceUint is a free data retrieval call binding the contract method 0xd37ee641.
//
// Solidity: function subReplaceUint(string jsonBlob, string searchPath, string replacePath, uint256 value) view returns(string)
func (_JsonUtil *JsonUtilSession) SubReplaceUint(jsonBlob string, searchPath string, replacePath string, value *big.Int) (string, error) {
	return _JsonUtil.Contract.SubReplaceUint(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplaceUint is a free data retrieval call binding the contract method 0xd37ee641.
//
// Solidity: function subReplaceUint(string jsonBlob, string searchPath, string replacePath, uint256 value) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SubReplaceUint(jsonBlob string, searchPath string, replacePath string, value *big.Int) (string, error) {
	return _JsonUtil.Contract.SubReplaceUint(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePath, value)
}

// SubReplaceUint0 is a free data retrieval call binding the contract method 0xfcc6cddf.
//
// Solidity: function subReplaceUint(string jsonBlob, string searchPath, string[] replacePaths, uint256[] values) view returns(string)
func (_JsonUtil *JsonUtilCaller) SubReplaceUint0(opts *bind.CallOpts, jsonBlob string, searchPath string, replacePaths []string, values []*big.Int) (string, error) {
	var out []interface{}
	err := _JsonUtil.contract.Call(opts, &out, "subReplaceUint0", jsonBlob, searchPath, replacePaths, values)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SubReplaceUint0 is a free data retrieval call binding the contract method 0xfcc6cddf.
//
// Solidity: function subReplaceUint(string jsonBlob, string searchPath, string[] replacePaths, uint256[] values) view returns(string)
func (_JsonUtil *JsonUtilSession) SubReplaceUint0(jsonBlob string, searchPath string, replacePaths []string, values []*big.Int) (string, error) {
	return _JsonUtil.Contract.SubReplaceUint0(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePaths, values)
}

// SubReplaceUint0 is a free data retrieval call binding the contract method 0xfcc6cddf.
//
// Solidity: function subReplaceUint(string jsonBlob, string searchPath, string[] replacePaths, uint256[] values) view returns(string)
func (_JsonUtil *JsonUtilCallerSession) SubReplaceUint0(jsonBlob string, searchPath string, replacePaths []string, values []*big.Int) (string, error) {
	return _JsonUtil.Contract.SubReplaceUint0(&_JsonUtil.CallOpts, jsonBlob, searchPath, replacePaths, values)
}
