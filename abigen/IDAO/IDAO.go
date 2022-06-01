// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IDAO

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

// IDAOMetaData contains all meta data concerning the IDAO contract.
var IDAOMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"inSupport\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_transactionData\",\"type\":\"bytes\"}],\"name\":\"makeProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_inSupport\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"unVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transactionData\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IDAOABI is the input ABI used to generate the binding from.
// Deprecated: Use IDAOMetaData.ABI instead.
var IDAOABI = IDAOMetaData.ABI

// IDAO is an auto generated Go binding around an Ethereum contract.
type IDAO struct {
	IDAOCaller     // Read-only binding to the contract
	IDAOTransactor // Write-only binding to the contract
	IDAOFilterer   // Log filterer for contract events
}

// IDAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDAOSession struct {
	Contract     *IDAO             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDAOCallerSession struct {
	Contract *IDAOCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IDAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDAOTransactorSession struct {
	Contract     *IDAOTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDAORaw is an auto generated low-level Go binding around an Ethereum contract.
type IDAORaw struct {
	Contract *IDAO // Generic contract binding to access the raw methods on
}

// IDAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDAOCallerRaw struct {
	Contract *IDAOCaller // Generic read-only contract binding to access the raw methods on
}

// IDAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDAOTransactorRaw struct {
	Contract *IDAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDAO creates a new instance of IDAO, bound to a specific deployed contract.
func NewIDAO(address common.Address, backend bind.ContractBackend) (*IDAO, error) {
	contract, err := bindIDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDAO{IDAOCaller: IDAOCaller{contract: contract}, IDAOTransactor: IDAOTransactor{contract: contract}, IDAOFilterer: IDAOFilterer{contract: contract}}, nil
}

// NewIDAOCaller creates a new read-only instance of IDAO, bound to a specific deployed contract.
func NewIDAOCaller(address common.Address, caller bind.ContractCaller) (*IDAOCaller, error) {
	contract, err := bindIDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOCaller{contract: contract}, nil
}

// NewIDAOTransactor creates a new write-only instance of IDAO, bound to a specific deployed contract.
func NewIDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*IDAOTransactor, error) {
	contract, err := bindIDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDAOTransactor{contract: contract}, nil
}

// NewIDAOFilterer creates a new log filterer instance of IDAO, bound to a specific deployed contract.
func NewIDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*IDAOFilterer, error) {
	contract, err := bindIDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDAOFilterer{contract: contract}, nil
}

// bindIDAO binds a generic wrapper to an already deployed contract.
func bindIDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDAOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAO *IDAORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAO.Contract.IDAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAO *IDAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAO.Contract.IDAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAO *IDAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAO.Contract.IDAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDAO *IDAOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDAO *IDAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDAO *IDAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDAO.Contract.contract.Transact(opts, method, params...)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 proposalID, bytes transactionData) returns(bool success)
func (_IDAO *IDAOTransactor) ExecuteProposal(opts *bind.TransactOpts, proposalID *big.Int, transactionData []byte) (*types.Transaction, error) {
	return _IDAO.contract.Transact(opts, "executeProposal", proposalID, transactionData)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 proposalID, bytes transactionData) returns(bool success)
func (_IDAO *IDAOSession) ExecuteProposal(proposalID *big.Int, transactionData []byte) (*types.Transaction, error) {
	return _IDAO.Contract.ExecuteProposal(&_IDAO.TransactOpts, proposalID, transactionData)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 proposalID, bytes transactionData) returns(bool success)
func (_IDAO *IDAOTransactorSession) ExecuteProposal(proposalID *big.Int, transactionData []byte) (*types.Transaction, error) {
	return _IDAO.Contract.ExecuteProposal(&_IDAO.TransactOpts, proposalID, transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_IDAO *IDAOTransactor) MakeProposal(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _IDAO.contract.Transact(opts, "makeProposal", _recipient, _amount, _description, _transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_IDAO *IDAOSession) MakeProposal(_recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _IDAO.Contract.MakeProposal(&_IDAO.TransactOpts, _recipient, _amount, _description, _transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_IDAO *IDAOTransactorSession) MakeProposal(_recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _IDAO.Contract.MakeProposal(&_IDAO.TransactOpts, _recipient, _amount, _description, _transactionData)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_IDAO *IDAOTransactor) UnVote(opts *bind.TransactOpts, _proposalID *big.Int) (*types.Transaction, error) {
	return _IDAO.contract.Transact(opts, "unVote", _proposalID)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_IDAO *IDAOSession) UnVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _IDAO.Contract.UnVote(&_IDAO.TransactOpts, _proposalID)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_IDAO *IDAOTransactorSession) UnVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _IDAO.Contract.UnVote(&_IDAO.TransactOpts, _proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_IDAO *IDAOTransactor) Vote(opts *bind.TransactOpts, _proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _IDAO.contract.Transact(opts, "vote", _proposalID, _inSupport)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_IDAO *IDAOSession) Vote(_proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _IDAO.Contract.Vote(&_IDAO.TransactOpts, _proposalID, _inSupport)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_IDAO *IDAOTransactorSession) Vote(_proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _IDAO.Contract.Vote(&_IDAO.TransactOpts, _proposalID, _inSupport)
}

// IDAOProposalAddedIterator is returned from FilterProposalAdded and is used to iterate over the raw logs and unpacked data for ProposalAdded events raised by the IDAO contract.
type IDAOProposalAddedIterator struct {
	Event *IDAOProposalAdded // Event containing the contract specifics and raw log

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
func (it *IDAOProposalAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDAOProposalAdded)
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
		it.Event = new(IDAOProposalAdded)
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
func (it *IDAOProposalAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDAOProposalAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDAOProposalAdded represents a ProposalAdded event raised by the IDAO contract.
type IDAOProposalAdded struct {
	ProposalID  *big.Int
	Recipient   common.Address
	Amount      *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalAdded is a free log retrieval operation binding the contract event 0x646fec02522b41e7125cfc859a64fd4f4cefd5dc3b6237ca0abe251ded1fa881.
//
// Solidity: event ProposalAdded(uint256 indexed proposalID, address recipient, uint256 amount, string description)
func (_IDAO *IDAOFilterer) FilterProposalAdded(opts *bind.FilterOpts, proposalID []*big.Int) (*IDAOProposalAddedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _IDAO.contract.FilterLogs(opts, "ProposalAdded", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &IDAOProposalAddedIterator{contract: _IDAO.contract, event: "ProposalAdded", logs: logs, sub: sub}, nil
}

// WatchProposalAdded is a free log subscription operation binding the contract event 0x646fec02522b41e7125cfc859a64fd4f4cefd5dc3b6237ca0abe251ded1fa881.
//
// Solidity: event ProposalAdded(uint256 indexed proposalID, address recipient, uint256 amount, string description)
func (_IDAO *IDAOFilterer) WatchProposalAdded(opts *bind.WatchOpts, sink chan<- *IDAOProposalAdded, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _IDAO.contract.WatchLogs(opts, "ProposalAdded", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDAOProposalAdded)
				if err := _IDAO.contract.UnpackLog(event, "ProposalAdded", log); err != nil {
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

// ParseProposalAdded is a log parse operation binding the contract event 0x646fec02522b41e7125cfc859a64fd4f4cefd5dc3b6237ca0abe251ded1fa881.
//
// Solidity: event ProposalAdded(uint256 indexed proposalID, address recipient, uint256 amount, string description)
func (_IDAO *IDAOFilterer) ParseProposalAdded(log types.Log) (*IDAOProposalAdded, error) {
	event := new(IDAOProposalAdded)
	if err := _IDAO.contract.UnpackLog(event, "ProposalAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDAOProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the IDAO contract.
type IDAOProposalExecutedIterator struct {
	Event *IDAOProposalExecuted // Event containing the contract specifics and raw log

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
func (it *IDAOProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDAOProposalExecuted)
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
		it.Event = new(IDAOProposalExecuted)
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
func (it *IDAOProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDAOProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDAOProposalExecuted represents a ProposalExecuted event raised by the IDAO contract.
type IDAOProposalExecuted struct {
	ProposalID *big.Int
	Result     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x948f4a9cd986f1118c3fbd459f7a22b23c0693e1ca3ef06a6a8be5aa7d39cc03.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, bool result)
func (_IDAO *IDAOFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalID []*big.Int) (*IDAOProposalExecutedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _IDAO.contract.FilterLogs(opts, "ProposalExecuted", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &IDAOProposalExecutedIterator{contract: _IDAO.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x948f4a9cd986f1118c3fbd459f7a22b23c0693e1ca3ef06a6a8be5aa7d39cc03.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, bool result)
func (_IDAO *IDAOFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *IDAOProposalExecuted, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _IDAO.contract.WatchLogs(opts, "ProposalExecuted", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDAOProposalExecuted)
				if err := _IDAO.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x948f4a9cd986f1118c3fbd459f7a22b23c0693e1ca3ef06a6a8be5aa7d39cc03.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, bool result)
func (_IDAO *IDAOFilterer) ParseProposalExecuted(log types.Log) (*IDAOProposalExecuted, error) {
	event := new(IDAOProposalExecuted)
	if err := _IDAO.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDAOVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the IDAO contract.
type IDAOVotedIterator struct {
	Event *IDAOVoted // Event containing the contract specifics and raw log

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
func (it *IDAOVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDAOVoted)
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
		it.Event = new(IDAOVoted)
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
func (it *IDAOVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDAOVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDAOVoted represents a Voted event raised by the IDAO contract.
type IDAOVoted struct {
	ProposalID *big.Int
	InSupport  bool
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed proposalID, bool inSupport, address indexed voter)
func (_IDAO *IDAOFilterer) FilterVoted(opts *bind.FilterOpts, proposalID []*big.Int, voter []common.Address) (*IDAOVotedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IDAO.contract.FilterLogs(opts, "Voted", proposalIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &IDAOVotedIterator{contract: _IDAO.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed proposalID, bool inSupport, address indexed voter)
func (_IDAO *IDAOFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *IDAOVoted, proposalID []*big.Int, voter []common.Address) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IDAO.contract.WatchLogs(opts, "Voted", proposalIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDAOVoted)
				if err := _IDAO.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed proposalID, bool inSupport, address indexed voter)
func (_IDAO *IDAOFilterer) ParseVoted(log types.Log) (*IDAOVoted, error) {
	event := new(IDAOVoted)
	if err := _IDAO.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
