// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ETHMover

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

// ETHMoverMetaData contains all meta data concerning the ETHMover contract.
var ETHMoverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516101553803806101558339818101604052810190610032919061008d565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610108565b600081519050610087816100f1565b92915050565b6000602082840312156100a3576100a26100ec565b5b60006100b184828501610078565b91505092915050565b60006100c5826100cc565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6100fa816100ba565b811461010557600080fd5b50565b603f806101166000396000f3fe6080604052600080fdfea2646970667358221220f7b4ee89a81dec41b18494de2219fd62bd0b2de59d5e0c9afc1559fe3d499f6564736f6c63430008060033",
}

// ETHMoverABI is the input ABI used to generate the binding from.
// Deprecated: Use ETHMoverMetaData.ABI instead.
var ETHMoverABI = ETHMoverMetaData.ABI

// ETHMoverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ETHMoverMetaData.Bin instead.
var ETHMoverBin = ETHMoverMetaData.Bin

// DeployETHMover deploys a new Ethereum contract, binding an instance of ETHMover to it.
func DeployETHMover(auth *bind.TransactOpts, backend bind.ContractBackend, _weth common.Address) (common.Address, *types.Transaction, *ETHMover, error) {
	parsed, err := ETHMoverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ETHMoverBin), backend, _weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ETHMover{ETHMoverCaller: ETHMoverCaller{contract: contract}, ETHMoverTransactor: ETHMoverTransactor{contract: contract}, ETHMoverFilterer: ETHMoverFilterer{contract: contract}}, nil
}

// ETHMover is an auto generated Go binding around an Ethereum contract.
type ETHMover struct {
	ETHMoverCaller     // Read-only binding to the contract
	ETHMoverTransactor // Write-only binding to the contract
	ETHMoverFilterer   // Log filterer for contract events
}

// ETHMoverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHMoverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHMoverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHMoverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHMoverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHMoverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHMoverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHMoverSession struct {
	Contract     *ETHMover         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHMoverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHMoverCallerSession struct {
	Contract *ETHMoverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ETHMoverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHMoverTransactorSession struct {
	Contract     *ETHMoverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ETHMoverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHMoverRaw struct {
	Contract *ETHMover // Generic contract binding to access the raw methods on
}

// ETHMoverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHMoverCallerRaw struct {
	Contract *ETHMoverCaller // Generic read-only contract binding to access the raw methods on
}

// ETHMoverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHMoverTransactorRaw struct {
	Contract *ETHMoverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETHMover creates a new instance of ETHMover, bound to a specific deployed contract.
func NewETHMover(address common.Address, backend bind.ContractBackend) (*ETHMover, error) {
	contract, err := bindETHMover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHMover{ETHMoverCaller: ETHMoverCaller{contract: contract}, ETHMoverTransactor: ETHMoverTransactor{contract: contract}, ETHMoverFilterer: ETHMoverFilterer{contract: contract}}, nil
}

// NewETHMoverCaller creates a new read-only instance of ETHMover, bound to a specific deployed contract.
func NewETHMoverCaller(address common.Address, caller bind.ContractCaller) (*ETHMoverCaller, error) {
	contract, err := bindETHMover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHMoverCaller{contract: contract}, nil
}

// NewETHMoverTransactor creates a new write-only instance of ETHMover, bound to a specific deployed contract.
func NewETHMoverTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHMoverTransactor, error) {
	contract, err := bindETHMover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHMoverTransactor{contract: contract}, nil
}

// NewETHMoverFilterer creates a new log filterer instance of ETHMover, bound to a specific deployed contract.
func NewETHMoverFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHMoverFilterer, error) {
	contract, err := bindETHMover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHMoverFilterer{contract: contract}, nil
}

// bindETHMover binds a generic wrapper to an already deployed contract.
func bindETHMover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ETHMoverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHMover *ETHMoverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHMover.Contract.ETHMoverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHMover *ETHMoverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHMover.Contract.ETHMoverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHMover *ETHMoverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHMover.Contract.ETHMoverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHMover *ETHMoverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHMover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHMover *ETHMoverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHMover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHMover *ETHMoverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHMover.Contract.contract.Transact(opts, method, params...)
}
