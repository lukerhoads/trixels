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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Abigen *AbigenTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Abigen *AbigenSession) EndAuction() (*types.Transaction, error) {
	return _Abigen.Contract.EndAuction(&_Abigen.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Abigen *AbigenTransactorSession) EndAuction() (*types.Transaction, error) {
	return _Abigen.Contract.EndAuction(&_Abigen.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenId) payable returns()
func (_Abigen *AbigenTransactor) PlaceBid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "placeBid", _tokenId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenId) payable returns()
func (_Abigen *AbigenSession) PlaceBid(_tokenId *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.PlaceBid(&_Abigen.TransactOpts, _tokenId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenId) payable returns()
func (_Abigen *AbigenTransactorSession) PlaceBid(_tokenId *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.PlaceBid(&_Abigen.TransactOpts, _tokenId)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_Abigen *AbigenTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_Abigen *AbigenSession) StartAuction() (*types.Transaction, error) {
	return _Abigen.Contract.StartAuction(&_Abigen.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_Abigen *AbigenTransactorSession) StartAuction() (*types.Transaction, error) {
	return _Abigen.Contract.StartAuction(&_Abigen.TransactOpts)
}

// AbigenAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the Abigen contract.
type AbigenAuctionEndedIterator struct {
	Event *AbigenAuctionEnded // Event containing the contract specifics and raw log

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
func (it *AbigenAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbigenAuctionEnded)
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
		it.Event = new(AbigenAuctionEnded)
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
func (it *AbigenAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbigenAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbigenAuctionEnded represents a AuctionEnded event raised by the Abigen contract.
type AbigenAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Abigen *AbigenFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*AbigenAuctionEndedIterator, error) {

	logs, sub, err := _Abigen.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &AbigenAuctionEndedIterator{contract: _Abigen.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Abigen *AbigenFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *AbigenAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _Abigen.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbigenAuctionEnded)
				if err := _Abigen.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Abigen *AbigenFilterer) ParseAuctionEnded(log types.Log) (*AbigenAuctionEnded, error) {
	event := new(AbigenAuctionEnded)
	if err := _Abigen.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbigenAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the Abigen contract.
type AbigenAuctionStartedIterator struct {
	Event *AbigenAuctionStarted // Event containing the contract specifics and raw log

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
func (it *AbigenAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbigenAuctionStarted)
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
		it.Event = new(AbigenAuctionStarted)
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
func (it *AbigenAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbigenAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbigenAuctionStarted represents a AuctionStarted event raised by the Abigen contract.
type AbigenAuctionStarted struct {
	TokenId   *big.Int
	StartDate *big.Int
	EndDate   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x44c53be110c6aa83aa83cd02e351ed172359268272ee1b5d31c0fe48db35c6c7.
//
// Solidity: event AuctionStarted(uint256 tokenId, uint256 startDate, uint256 endDate)
func (_Abigen *AbigenFilterer) FilterAuctionStarted(opts *bind.FilterOpts) (*AbigenAuctionStartedIterator, error) {

	logs, sub, err := _Abigen.contract.FilterLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return &AbigenAuctionStartedIterator{contract: _Abigen.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x44c53be110c6aa83aa83cd02e351ed172359268272ee1b5d31c0fe48db35c6c7.
//
// Solidity: event AuctionStarted(uint256 tokenId, uint256 startDate, uint256 endDate)
func (_Abigen *AbigenFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *AbigenAuctionStarted) (event.Subscription, error) {

	logs, sub, err := _Abigen.contract.WatchLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbigenAuctionStarted)
				if err := _Abigen.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// ParseAuctionStarted is a log parse operation binding the contract event 0x44c53be110c6aa83aa83cd02e351ed172359268272ee1b5d31c0fe48db35c6c7.
//
// Solidity: event AuctionStarted(uint256 tokenId, uint256 startDate, uint256 endDate)
func (_Abigen *AbigenFilterer) ParseAuctionStarted(log types.Log) (*AbigenAuctionStarted, error) {
	event := new(AbigenAuctionStarted)
	if err := _Abigen.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbigenBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the Abigen contract.
type AbigenBidPlacedIterator struct {
	Event *AbigenBidPlaced // Event containing the contract specifics and raw log

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
func (it *AbigenBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbigenBidPlaced)
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
		it.Event = new(AbigenBidPlaced)
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
func (it *AbigenBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbigenBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbigenBidPlaced represents a BidPlaced event raised by the Abigen contract.
type AbigenBidPlaced struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address bidder, uint256 amount)
func (_Abigen *AbigenFilterer) FilterBidPlaced(opts *bind.FilterOpts) (*AbigenBidPlacedIterator, error) {

	logs, sub, err := _Abigen.contract.FilterLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return &AbigenBidPlacedIterator{contract: _Abigen.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address bidder, uint256 amount)
func (_Abigen *AbigenFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *AbigenBidPlaced) (event.Subscription, error) {

	logs, sub, err := _Abigen.contract.WatchLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbigenBidPlaced)
				if err := _Abigen.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address bidder, uint256 amount)
func (_Abigen *AbigenFilterer) ParseBidPlaced(log types.Log) (*AbigenBidPlaced, error) {
	event := new(AbigenBidPlaced)
	if err := _Abigen.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
