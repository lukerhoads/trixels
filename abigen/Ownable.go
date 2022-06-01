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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abigen *AbigenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abigen *AbigenSession) Owner() (common.Address, error) {
	return _Abigen.Contract.Owner(&_Abigen.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abigen *AbigenCallerSession) Owner() (common.Address, error) {
	return _Abigen.Contract.Owner(&_Abigen.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abigen *AbigenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abigen *AbigenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abigen.Contract.RenounceOwnership(&_Abigen.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abigen *AbigenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abigen.Contract.RenounceOwnership(&_Abigen.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abigen *AbigenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abigen *AbigenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abigen.Contract.TransferOwnership(&_Abigen.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abigen *AbigenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abigen.Contract.TransferOwnership(&_Abigen.TransactOpts, newOwner)
}

// AbigenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abigen contract.
type AbigenOwnershipTransferredIterator struct {
	Event *AbigenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbigenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbigenOwnershipTransferred)
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
		it.Event = new(AbigenOwnershipTransferred)
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
func (it *AbigenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbigenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbigenOwnershipTransferred represents a OwnershipTransferred event raised by the Abigen contract.
type AbigenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abigen *AbigenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbigenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abigen.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbigenOwnershipTransferredIterator{contract: _Abigen.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abigen *AbigenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbigenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abigen.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbigenOwnershipTransferred)
				if err := _Abigen.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abigen *AbigenFilterer) ParseOwnershipTransferred(log types.Log) (*AbigenOwnershipTransferred, error) {
	event := new(AbigenOwnershipTransferred)
	if err := _Abigen.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
