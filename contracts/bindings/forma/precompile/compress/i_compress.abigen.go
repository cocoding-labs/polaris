// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compress

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

// CompressMetaData contains all meta data concerning the Compress contract.
var CompressMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deflate\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inflate\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"}]",
}

// CompressABI is the input ABI used to generate the binding from.
// Deprecated: Use CompressMetaData.ABI instead.
var CompressABI = CompressMetaData.ABI

// Compress is an auto generated Go binding around an Ethereum contract.
type Compress struct {
	CompressCaller     // Read-only binding to the contract
	CompressTransactor // Write-only binding to the contract
	CompressFilterer   // Log filterer for contract events
}

// CompressCaller is an auto generated read-only Go binding around an Ethereum contract.
type CompressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CompressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CompressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CompressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CompressSession struct {
	Contract     *Compress         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CompressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CompressCallerSession struct {
	Contract *CompressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CompressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CompressTransactorSession struct {
	Contract     *CompressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CompressRaw is an auto generated low-level Go binding around an Ethereum contract.
type CompressRaw struct {
	Contract *Compress // Generic contract binding to access the raw methods on
}

// CompressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CompressCallerRaw struct {
	Contract *CompressCaller // Generic read-only contract binding to access the raw methods on
}

// CompressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CompressTransactorRaw struct {
	Contract *CompressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCompress creates a new instance of Compress, bound to a specific deployed contract.
func NewCompress(address common.Address, backend bind.ContractBackend) (*Compress, error) {
	contract, err := bindCompress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Compress{CompressCaller: CompressCaller{contract: contract}, CompressTransactor: CompressTransactor{contract: contract}, CompressFilterer: CompressFilterer{contract: contract}}, nil
}

// NewCompressCaller creates a new read-only instance of Compress, bound to a specific deployed contract.
func NewCompressCaller(address common.Address, caller bind.ContractCaller) (*CompressCaller, error) {
	contract, err := bindCompress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CompressCaller{contract: contract}, nil
}

// NewCompressTransactor creates a new write-only instance of Compress, bound to a specific deployed contract.
func NewCompressTransactor(address common.Address, transactor bind.ContractTransactor) (*CompressTransactor, error) {
	contract, err := bindCompress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CompressTransactor{contract: contract}, nil
}

// NewCompressFilterer creates a new log filterer instance of Compress, bound to a specific deployed contract.
func NewCompressFilterer(address common.Address, filterer bind.ContractFilterer) (*CompressFilterer, error) {
	contract, err := bindCompress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CompressFilterer{contract: contract}, nil
}

// bindCompress binds a generic wrapper to an already deployed contract.
func bindCompress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CompressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compress *CompressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compress.Contract.CompressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compress *CompressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compress.Contract.CompressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compress *CompressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compress.Contract.CompressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Compress *CompressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Compress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Compress *CompressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Compress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Compress *CompressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Compress.Contract.contract.Transact(opts, method, params...)
}

// Deflate is a free data retrieval call binding the contract method 0x24623720.
//
// Solidity: function deflate(bytes[] data) view returns(bytes[])
func (_Compress *CompressCaller) Deflate(opts *bind.CallOpts, data [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _Compress.contract.Call(opts, &out, "deflate", data)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// Deflate is a free data retrieval call binding the contract method 0x24623720.
//
// Solidity: function deflate(bytes[] data) view returns(bytes[])
func (_Compress *CompressSession) Deflate(data [][]byte) ([][]byte, error) {
	return _Compress.Contract.Deflate(&_Compress.CallOpts, data)
}

// Deflate is a free data retrieval call binding the contract method 0x24623720.
//
// Solidity: function deflate(bytes[] data) view returns(bytes[])
func (_Compress *CompressCallerSession) Deflate(data [][]byte) ([][]byte, error) {
	return _Compress.Contract.Deflate(&_Compress.CallOpts, data)
}

// Inflate is a free data retrieval call binding the contract method 0x595100df.
//
// Solidity: function inflate(bytes[] data) view returns(bytes[])
func (_Compress *CompressCaller) Inflate(opts *bind.CallOpts, data [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _Compress.contract.Call(opts, &out, "inflate", data)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// Inflate is a free data retrieval call binding the contract method 0x595100df.
//
// Solidity: function inflate(bytes[] data) view returns(bytes[])
func (_Compress *CompressSession) Inflate(data [][]byte) ([][]byte, error) {
	return _Compress.Contract.Inflate(&_Compress.CallOpts, data)
}

// Inflate is a free data retrieval call binding the contract method 0x595100df.
//
// Solidity: function inflate(bytes[] data) view returns(bytes[])
func (_Compress *CompressCallerSession) Inflate(data [][]byte) ([][]byte, error) {
	return _Compress.Contract.Inflate(&_Compress.CallOpts, data)
}
