// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAuctionHouse

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

// IAuctionHouseMetaData contains all meta data concerning the IAuctionHouse contract.
var IAuctionHouseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IAuctionHouseABI is the input ABI used to generate the binding from.
// Deprecated: Use IAuctionHouseMetaData.ABI instead.
var IAuctionHouseABI = IAuctionHouseMetaData.ABI

// IAuctionHouse is an auto generated Go binding around an Ethereum contract.
type IAuctionHouse struct {
	IAuctionHouseCaller     // Read-only binding to the contract
	IAuctionHouseTransactor // Write-only binding to the contract
	IAuctionHouseFilterer   // Log filterer for contract events
}

// IAuctionHouseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuctionHouseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionHouseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuctionHouseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionHouseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuctionHouseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionHouseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuctionHouseSession struct {
	Contract     *IAuctionHouse    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuctionHouseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuctionHouseCallerSession struct {
	Contract *IAuctionHouseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAuctionHouseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuctionHouseTransactorSession struct {
	Contract     *IAuctionHouseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAuctionHouseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuctionHouseRaw struct {
	Contract *IAuctionHouse // Generic contract binding to access the raw methods on
}

// IAuctionHouseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuctionHouseCallerRaw struct {
	Contract *IAuctionHouseCaller // Generic read-only contract binding to access the raw methods on
}

// IAuctionHouseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuctionHouseTransactorRaw struct {
	Contract *IAuctionHouseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuctionHouse creates a new instance of IAuctionHouse, bound to a specific deployed contract.
func NewIAuctionHouse(address common.Address, backend bind.ContractBackend) (*IAuctionHouse, error) {
	contract, err := bindIAuctionHouse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuctionHouse{IAuctionHouseCaller: IAuctionHouseCaller{contract: contract}, IAuctionHouseTransactor: IAuctionHouseTransactor{contract: contract}, IAuctionHouseFilterer: IAuctionHouseFilterer{contract: contract}}, nil
}

// NewIAuctionHouseCaller creates a new read-only instance of IAuctionHouse, bound to a specific deployed contract.
func NewIAuctionHouseCaller(address common.Address, caller bind.ContractCaller) (*IAuctionHouseCaller, error) {
	contract, err := bindIAuctionHouse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuctionHouseCaller{contract: contract}, nil
}

// NewIAuctionHouseTransactor creates a new write-only instance of IAuctionHouse, bound to a specific deployed contract.
func NewIAuctionHouseTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuctionHouseTransactor, error) {
	contract, err := bindIAuctionHouse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuctionHouseTransactor{contract: contract}, nil
}

// NewIAuctionHouseFilterer creates a new log filterer instance of IAuctionHouse, bound to a specific deployed contract.
func NewIAuctionHouseFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuctionHouseFilterer, error) {
	contract, err := bindIAuctionHouse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuctionHouseFilterer{contract: contract}, nil
}

// bindIAuctionHouse binds a generic wrapper to an already deployed contract.
func bindIAuctionHouse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuctionHouseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuctionHouse *IAuctionHouseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuctionHouse.Contract.IAuctionHouseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuctionHouse *IAuctionHouseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuctionHouse.Contract.IAuctionHouseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuctionHouse *IAuctionHouseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuctionHouse.Contract.IAuctionHouseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuctionHouse *IAuctionHouseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuctionHouse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuctionHouse *IAuctionHouseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuctionHouse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuctionHouse *IAuctionHouseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuctionHouse.Contract.contract.Transact(opts, method, params...)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_IAuctionHouse *IAuctionHouseTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuctionHouse.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_IAuctionHouse *IAuctionHouseSession) EndAuction() (*types.Transaction, error) {
	return _IAuctionHouse.Contract.EndAuction(&_IAuctionHouse.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_IAuctionHouse *IAuctionHouseTransactorSession) EndAuction() (*types.Transaction, error) {
	return _IAuctionHouse.Contract.EndAuction(&_IAuctionHouse.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenId) payable returns()
func (_IAuctionHouse *IAuctionHouseTransactor) PlaceBid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _IAuctionHouse.contract.Transact(opts, "placeBid", _tokenId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenId) payable returns()
func (_IAuctionHouse *IAuctionHouseSession) PlaceBid(_tokenId *big.Int) (*types.Transaction, error) {
	return _IAuctionHouse.Contract.PlaceBid(&_IAuctionHouse.TransactOpts, _tokenId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenId) payable returns()
func (_IAuctionHouse *IAuctionHouseTransactorSession) PlaceBid(_tokenId *big.Int) (*types.Transaction, error) {
	return _IAuctionHouse.Contract.PlaceBid(&_IAuctionHouse.TransactOpts, _tokenId)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_IAuctionHouse *IAuctionHouseTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuctionHouse.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_IAuctionHouse *IAuctionHouseSession) StartAuction() (*types.Transaction, error) {
	return _IAuctionHouse.Contract.StartAuction(&_IAuctionHouse.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_IAuctionHouse *IAuctionHouseTransactorSession) StartAuction() (*types.Transaction, error) {
	return _IAuctionHouse.Contract.StartAuction(&_IAuctionHouse.TransactOpts)
}

// IAuctionHouseAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the IAuctionHouse contract.
type IAuctionHouseAuctionEndedIterator struct {
	Event *IAuctionHouseAuctionEnded // Event containing the contract specifics and raw log

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
func (it *IAuctionHouseAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionHouseAuctionEnded)
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
		it.Event = new(IAuctionHouseAuctionEnded)
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
func (it *IAuctionHouseAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionHouseAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionHouseAuctionEnded represents a AuctionEnded event raised by the IAuctionHouse contract.
type IAuctionHouseAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_IAuctionHouse *IAuctionHouseFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*IAuctionHouseAuctionEndedIterator, error) {

	logs, sub, err := _IAuctionHouse.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &IAuctionHouseAuctionEndedIterator{contract: _IAuctionHouse.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_IAuctionHouse *IAuctionHouseFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *IAuctionHouseAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _IAuctionHouse.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionHouseAuctionEnded)
				if err := _IAuctionHouse.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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
func (_IAuctionHouse *IAuctionHouseFilterer) ParseAuctionEnded(log types.Log) (*IAuctionHouseAuctionEnded, error) {
	event := new(IAuctionHouseAuctionEnded)
	if err := _IAuctionHouse.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuctionHouseAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the IAuctionHouse contract.
type IAuctionHouseAuctionStartedIterator struct {
	Event *IAuctionHouseAuctionStarted // Event containing the contract specifics and raw log

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
func (it *IAuctionHouseAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionHouseAuctionStarted)
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
		it.Event = new(IAuctionHouseAuctionStarted)
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
func (it *IAuctionHouseAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionHouseAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionHouseAuctionStarted represents a AuctionStarted event raised by the IAuctionHouse contract.
type IAuctionHouseAuctionStarted struct {
	TokenId   *big.Int
	StartDate *big.Int
	EndDate   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x44c53be110c6aa83aa83cd02e351ed172359268272ee1b5d31c0fe48db35c6c7.
//
// Solidity: event AuctionStarted(uint256 tokenId, uint256 startDate, uint256 endDate)
func (_IAuctionHouse *IAuctionHouseFilterer) FilterAuctionStarted(opts *bind.FilterOpts) (*IAuctionHouseAuctionStartedIterator, error) {

	logs, sub, err := _IAuctionHouse.contract.FilterLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return &IAuctionHouseAuctionStartedIterator{contract: _IAuctionHouse.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x44c53be110c6aa83aa83cd02e351ed172359268272ee1b5d31c0fe48db35c6c7.
//
// Solidity: event AuctionStarted(uint256 tokenId, uint256 startDate, uint256 endDate)
func (_IAuctionHouse *IAuctionHouseFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *IAuctionHouseAuctionStarted) (event.Subscription, error) {

	logs, sub, err := _IAuctionHouse.contract.WatchLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionHouseAuctionStarted)
				if err := _IAuctionHouse.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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
func (_IAuctionHouse *IAuctionHouseFilterer) ParseAuctionStarted(log types.Log) (*IAuctionHouseAuctionStarted, error) {
	event := new(IAuctionHouseAuctionStarted)
	if err := _IAuctionHouse.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuctionHouseBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the IAuctionHouse contract.
type IAuctionHouseBidPlacedIterator struct {
	Event *IAuctionHouseBidPlaced // Event containing the contract specifics and raw log

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
func (it *IAuctionHouseBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionHouseBidPlaced)
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
		it.Event = new(IAuctionHouseBidPlaced)
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
func (it *IAuctionHouseBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionHouseBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionHouseBidPlaced represents a BidPlaced event raised by the IAuctionHouse contract.
type IAuctionHouseBidPlaced struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address bidder, uint256 amount)
func (_IAuctionHouse *IAuctionHouseFilterer) FilterBidPlaced(opts *bind.FilterOpts) (*IAuctionHouseBidPlacedIterator, error) {

	logs, sub, err := _IAuctionHouse.contract.FilterLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return &IAuctionHouseBidPlacedIterator{contract: _IAuctionHouse.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address bidder, uint256 amount)
func (_IAuctionHouse *IAuctionHouseFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *IAuctionHouseBidPlaced) (event.Subscription, error) {

	logs, sub, err := _IAuctionHouse.contract.WatchLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionHouseBidPlaced)
				if err := _IAuctionHouse.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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
func (_IAuctionHouse *IAuctionHouseFilterer) ParseBidPlaced(log types.Log) (*IAuctionHouseBidPlaced, error) {
	event := new(IAuctionHouseBidPlaced)
	if err := _IAuctionHouse.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
