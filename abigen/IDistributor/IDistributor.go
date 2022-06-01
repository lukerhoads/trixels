// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDistributor

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

// IDistributorContributor is an auto generated low-level Go binding around an user-defined struct.
type IDistributorContributor struct {
	Addr        common.Address
	NumContribs *big.Int
}

// IDistributorMetaData contains all meta data concerning the IDistributor contract.
var IDistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"recordSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numContribs\",\"type\":\"uint256\"}],\"internalType\":\"structIDistributor.Contributor[]\",\"name\":\"_contributors\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use IDistributorMetaData.ABI instead.
var IDistributorABI = IDistributorMetaData.ABI

// IDistributor is an auto generated Go binding around an Ethereum contract.
type IDistributor struct {
	IDistributorCaller     // Read-only binding to the contract
	IDistributorTransactor // Write-only binding to the contract
	IDistributorFilterer   // Log filterer for contract events
}

// IDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDistributorSession struct {
	Contract     *IDistributor     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDistributorCallerSession struct {
	Contract *IDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDistributorTransactorSession struct {
	Contract     *IDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDistributorRaw struct {
	Contract *IDistributor // Generic contract binding to access the raw methods on
}

// IDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDistributorCallerRaw struct {
	Contract *IDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// IDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDistributorTransactorRaw struct {
	Contract *IDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDistributor creates a new instance of IDistributor, bound to a specific deployed contract.
func NewIDistributor(address common.Address, backend bind.ContractBackend) (*IDistributor, error) {
	contract, err := bindIDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDistributor{IDistributorCaller: IDistributorCaller{contract: contract}, IDistributorTransactor: IDistributorTransactor{contract: contract}, IDistributorFilterer: IDistributorFilterer{contract: contract}}, nil
}

// NewIDistributorCaller creates a new read-only instance of IDistributor, bound to a specific deployed contract.
func NewIDistributorCaller(address common.Address, caller bind.ContractCaller) (*IDistributorCaller, error) {
	contract, err := bindIDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDistributorCaller{contract: contract}, nil
}

// NewIDistributorTransactor creates a new write-only instance of IDistributor, bound to a specific deployed contract.
func NewIDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*IDistributorTransactor, error) {
	contract, err := bindIDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDistributorTransactor{contract: contract}, nil
}

// NewIDistributorFilterer creates a new log filterer instance of IDistributor, bound to a specific deployed contract.
func NewIDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*IDistributorFilterer, error) {
	contract, err := bindIDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDistributorFilterer{contract: contract}, nil
}

// bindIDistributor binds a generic wrapper to an already deployed contract.
func bindIDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDistributor *IDistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDistributor.Contract.IDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDistributor *IDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistributor.Contract.IDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDistributor *IDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDistributor.Contract.IDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDistributor *IDistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDistributor *IDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDistributor *IDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDistributor.Contract.contract.Transact(opts, method, params...)
}

// Distribute is a paid mutator transaction binding the contract method 0x8daf8596.
//
// Solidity: function distribute((address,uint256)[] _contributors, uint256 _salePrice) returns()
func (_IDistributor *IDistributorTransactor) Distribute(opts *bind.TransactOpts, _contributors []IDistributorContributor, _salePrice *big.Int) (*types.Transaction, error) {
	return _IDistributor.contract.Transact(opts, "distribute", _contributors, _salePrice)
}

// Distribute is a paid mutator transaction binding the contract method 0x8daf8596.
//
// Solidity: function distribute((address,uint256)[] _contributors, uint256 _salePrice) returns()
func (_IDistributor *IDistributorSession) Distribute(_contributors []IDistributorContributor, _salePrice *big.Int) (*types.Transaction, error) {
	return _IDistributor.Contract.Distribute(&_IDistributor.TransactOpts, _contributors, _salePrice)
}

// Distribute is a paid mutator transaction binding the contract method 0x8daf8596.
//
// Solidity: function distribute((address,uint256)[] _contributors, uint256 _salePrice) returns()
func (_IDistributor *IDistributorTransactorSession) Distribute(_contributors []IDistributorContributor, _salePrice *big.Int) (*types.Transaction, error) {
	return _IDistributor.Contract.Distribute(&_IDistributor.TransactOpts, _contributors, _salePrice)
}

// RecordSale is a paid mutator transaction binding the contract method 0x6f27f690.
//
// Solidity: function recordSale(uint256 _tokenId) payable returns()
func (_IDistributor *IDistributorTransactor) RecordSale(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _IDistributor.contract.Transact(opts, "recordSale", _tokenId)
}

// RecordSale is a paid mutator transaction binding the contract method 0x6f27f690.
//
// Solidity: function recordSale(uint256 _tokenId) payable returns()
func (_IDistributor *IDistributorSession) RecordSale(_tokenId *big.Int) (*types.Transaction, error) {
	return _IDistributor.Contract.RecordSale(&_IDistributor.TransactOpts, _tokenId)
}

// RecordSale is a paid mutator transaction binding the contract method 0x6f27f690.
//
// Solidity: function recordSale(uint256 _tokenId) payable returns()
func (_IDistributor *IDistributorTransactorSession) RecordSale(_tokenId *big.Int) (*types.Transaction, error) {
	return _IDistributor.Contract.RecordSale(&_IDistributor.TransactOpts, _tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IDistributor *IDistributorTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDistributor.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IDistributor *IDistributorSession) Withdraw() (*types.Transaction, error) {
	return _IDistributor.Contract.Withdraw(&_IDistributor.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IDistributor *IDistributorTransactorSession) Withdraw() (*types.Transaction, error) {
	return _IDistributor.Contract.Withdraw(&_IDistributor.TransactOpts)
}
