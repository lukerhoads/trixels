// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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
)

// AbigenMetaData contains all meta data concerning the Abigen contract.
var AbigenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AbigenABI is the input ABI used to generate the binding from.
// Deprecated: Use AbigenMetaData.ABI instead.
var AbigenABI = AbigenMetaData.ABI

// Abigen is an auto generated Go binding around an Ethereum contract.
type Abigen struct {
	AbigenCaller     // Read-only binding to the contract
	AbigenTransactor // Write-only binding to the contract
	AbigenFilterer   // Log filterer for contract events
}

// AbigenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbigenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbigenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbigenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbigenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbigenSession struct {
	Contract     *Abigen           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbigenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbigenCallerSession struct {
	Contract *AbigenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbigenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbigenTransactorSession struct {
	Contract     *AbigenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbigenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbigenRaw struct {
	Contract *Abigen // Generic contract binding to access the raw methods on
}

// AbigenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbigenCallerRaw struct {
	Contract *AbigenCaller // Generic read-only contract binding to access the raw methods on
}

// AbigenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbigenTransactorRaw struct {
	Contract *AbigenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbigen creates a new instance of Abigen, bound to a specific deployed contract.
func NewAbigen(address common.Address, backend bind.ContractBackend) (*Abigen, error) {
	contract, err := bindAbigen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abigen{AbigenCaller: AbigenCaller{contract: contract}, AbigenTransactor: AbigenTransactor{contract: contract}, AbigenFilterer: AbigenFilterer{contract: contract}}, nil
}

// NewAbigenCaller creates a new read-only instance of Abigen, bound to a specific deployed contract.
func NewAbigenCaller(address common.Address, caller bind.ContractCaller) (*AbigenCaller, error) {
	contract, err := bindAbigen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbigenCaller{contract: contract}, nil
}

// NewAbigenTransactor creates a new write-only instance of Abigen, bound to a specific deployed contract.
func NewAbigenTransactor(address common.Address, transactor bind.ContractTransactor) (*AbigenTransactor, error) {
	contract, err := bindAbigen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbigenTransactor{contract: contract}, nil
}

// NewAbigenFilterer creates a new log filterer instance of Abigen, bound to a specific deployed contract.
func NewAbigenFilterer(address common.Address, filterer bind.ContractFilterer) (*AbigenFilterer, error) {
	contract, err := bindAbigen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbigenFilterer{contract: contract}, nil
}

// bindAbigen binds a generic wrapper to an already deployed contract.
func bindAbigen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbigenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abigen *AbigenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abigen.Contract.AbigenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abigen *AbigenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigen.Contract.AbigenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abigen *AbigenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abigen.Contract.AbigenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abigen *AbigenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abigen.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abigen *AbigenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigen.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abigen *AbigenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abigen.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Abigen *AbigenTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Abigen *AbigenSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Abigen.Contract.OnERC721Received(&_Abigen.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_Abigen *AbigenTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Abigen.Contract.OnERC721Received(&_Abigen.TransactOpts, operator, from, tokenId, data)
}
