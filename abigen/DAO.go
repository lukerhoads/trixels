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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"inSupport\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"numProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"yay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nay\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_transactionData\",\"type\":\"bytes\"}],\"name\":\"makeProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_inSupport\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"unVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_transactionData\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_transactionData\",\"type\":\"bytes\"}],\"name\":\"checkProposalHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"validHash\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// CheckProposalHash is a free data retrieval call binding the contract method 0x2891d6f2.
//
// Solidity: function checkProposalHash(uint256 _proposalID, address _recipient, uint256 _amount, bytes _transactionData) view returns(bool validHash)
func (_Abigen *AbigenCaller) CheckProposalHash(opts *bind.CallOpts, _proposalID *big.Int, _recipient common.Address, _amount *big.Int, _transactionData []byte) (bool, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "checkProposalHash", _proposalID, _recipient, _amount, _transactionData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProposalHash is a free data retrieval call binding the contract method 0x2891d6f2.
//
// Solidity: function checkProposalHash(uint256 _proposalID, address _recipient, uint256 _amount, bytes _transactionData) view returns(bool validHash)
func (_Abigen *AbigenSession) CheckProposalHash(_proposalID *big.Int, _recipient common.Address, _amount *big.Int, _transactionData []byte) (bool, error) {
	return _Abigen.Contract.CheckProposalHash(&_Abigen.CallOpts, _proposalID, _recipient, _amount, _transactionData)
}

// CheckProposalHash is a free data retrieval call binding the contract method 0x2891d6f2.
//
// Solidity: function checkProposalHash(uint256 _proposalID, address _recipient, uint256 _amount, bytes _transactionData) view returns(bool validHash)
func (_Abigen *AbigenCallerSession) CheckProposalHash(_proposalID *big.Int, _recipient common.Address, _amount *big.Int, _transactionData []byte) (bool, error) {
	return _Abigen.Contract.CheckProposalHash(&_Abigen.CallOpts, _proposalID, _recipient, _amount, _transactionData)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address person) view returns(bool)
func (_Abigen *AbigenCaller) IsMember(opts *bind.CallOpts, person common.Address) (bool, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "isMember", person)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address person) view returns(bool)
func (_Abigen *AbigenSession) IsMember(person common.Address) (bool, error) {
	return _Abigen.Contract.IsMember(&_Abigen.CallOpts, person)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address person) view returns(bool)
func (_Abigen *AbigenCallerSession) IsMember(person common.Address) (bool, error) {
	return _Abigen.Contract.IsMember(&_Abigen.CallOpts, person)
}

// NumProposals is a free data retrieval call binding the contract method 0x400e3949.
//
// Solidity: function numProposals() view returns(uint256)
func (_Abigen *AbigenCaller) NumProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "numProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumProposals is a free data retrieval call binding the contract method 0x400e3949.
//
// Solidity: function numProposals() view returns(uint256)
func (_Abigen *AbigenSession) NumProposals() (*big.Int, error) {
	return _Abigen.Contract.NumProposals(&_Abigen.CallOpts)
}

// NumProposals is a free data retrieval call binding the contract method 0x400e3949.
//
// Solidity: function numProposals() view returns(uint256)
func (_Abigen *AbigenCallerSession) NumProposals() (*big.Int, error) {
	return _Abigen.Contract.NumProposals(&_Abigen.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 proposalHash, address recipient, uint256 amount, string description, bool passed, uint256 yay, uint256 nay, address creator, uint256 createdAt)
func (_Abigen *AbigenCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProposalHash [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Description  string
	Passed       bool
	Yay          *big.Int
	Nay          *big.Int
	Creator      common.Address
	CreatedAt    *big.Int
}, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		ProposalHash [32]byte
		Recipient    common.Address
		Amount       *big.Int
		Description  string
		Passed       bool
		Yay          *big.Int
		Nay          *big.Int
		Creator      common.Address
		CreatedAt    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProposalHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Recipient = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Description = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Passed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Yay = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Nay = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Creator = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 proposalHash, address recipient, uint256 amount, string description, bool passed, uint256 yay, uint256 nay, address creator, uint256 createdAt)
func (_Abigen *AbigenSession) Proposals(arg0 *big.Int) (struct {
	ProposalHash [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Description  string
	Passed       bool
	Yay          *big.Int
	Nay          *big.Int
	Creator      common.Address
	CreatedAt    *big.Int
}, error) {
	return _Abigen.Contract.Proposals(&_Abigen.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 proposalHash, address recipient, uint256 amount, string description, bool passed, uint256 yay, uint256 nay, address creator, uint256 createdAt)
func (_Abigen *AbigenCallerSession) Proposals(arg0 *big.Int) (struct {
	ProposalHash [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Description  string
	Passed       bool
	Yay          *big.Int
	Nay          *big.Int
	Creator      common.Address
	CreatedAt    *big.Int
}, error) {
	return _Abigen.Contract.Proposals(&_Abigen.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Abigen *AbigenCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Abigen *AbigenSession) Token() (common.Address, error) {
	return _Abigen.Contract.Token(&_Abigen.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Abigen *AbigenCallerSession) Token() (common.Address, error) {
	return _Abigen.Contract.Token(&_Abigen.CallOpts)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 _proposalID, bytes _transactionData) returns(bool success)
func (_Abigen *AbigenTransactor) ExecuteProposal(opts *bind.TransactOpts, _proposalID *big.Int, _transactionData []byte) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "executeProposal", _proposalID, _transactionData)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 _proposalID, bytes _transactionData) returns(bool success)
func (_Abigen *AbigenSession) ExecuteProposal(_proposalID *big.Int, _transactionData []byte) (*types.Transaction, error) {
	return _Abigen.Contract.ExecuteProposal(&_Abigen.TransactOpts, _proposalID, _transactionData)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 _proposalID, bytes _transactionData) returns(bool success)
func (_Abigen *AbigenTransactorSession) ExecuteProposal(_proposalID *big.Int, _transactionData []byte) (*types.Transaction, error) {
	return _Abigen.Contract.ExecuteProposal(&_Abigen.TransactOpts, _proposalID, _transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_Abigen *AbigenTransactor) MakeProposal(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "makeProposal", _recipient, _amount, _description, _transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_Abigen *AbigenSession) MakeProposal(_recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _Abigen.Contract.MakeProposal(&_Abigen.TransactOpts, _recipient, _amount, _description, _transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_Abigen *AbigenTransactorSession) MakeProposal(_recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _Abigen.Contract.MakeProposal(&_Abigen.TransactOpts, _recipient, _amount, _description, _transactionData)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_Abigen *AbigenTransactor) UnVote(opts *bind.TransactOpts, _proposalID *big.Int) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "unVote", _proposalID)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_Abigen *AbigenSession) UnVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.UnVote(&_Abigen.TransactOpts, _proposalID)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_Abigen *AbigenTransactorSession) UnVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.UnVote(&_Abigen.TransactOpts, _proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_Abigen *AbigenTransactor) Vote(opts *bind.TransactOpts, _proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "vote", _proposalID, _inSupport)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_Abigen *AbigenSession) Vote(_proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _Abigen.Contract.Vote(&_Abigen.TransactOpts, _proposalID, _inSupport)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_Abigen *AbigenTransactorSession) Vote(_proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _Abigen.Contract.Vote(&_Abigen.TransactOpts, _proposalID, _inSupport)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abigen *AbigenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abigen.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abigen *AbigenSession) Receive() (*types.Transaction, error) {
	return _Abigen.Contract.Receive(&_Abigen.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Abigen *AbigenTransactorSession) Receive() (*types.Transaction, error) {
	return _Abigen.Contract.Receive(&_Abigen.TransactOpts)
}

// AbigenProposalAddedIterator is returned from FilterProposalAdded and is used to iterate over the raw logs and unpacked data for ProposalAdded events raised by the Abigen contract.
type AbigenProposalAddedIterator struct {
	Event *AbigenProposalAdded // Event containing the contract specifics and raw log

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
func (it *AbigenProposalAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbigenProposalAdded)
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
		it.Event = new(AbigenProposalAdded)
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
func (it *AbigenProposalAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbigenProposalAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbigenProposalAdded represents a ProposalAdded event raised by the Abigen contract.
type AbigenProposalAdded struct {
	ProposalID  *big.Int
	Recipient   common.Address
	Amount      *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalAdded is a free log retrieval operation binding the contract event 0x646fec02522b41e7125cfc859a64fd4f4cefd5dc3b6237ca0abe251ded1fa881.
//
// Solidity: event ProposalAdded(uint256 indexed proposalID, address recipient, uint256 amount, string description)
func (_Abigen *AbigenFilterer) FilterProposalAdded(opts *bind.FilterOpts, proposalID []*big.Int) (*AbigenProposalAddedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _Abigen.contract.FilterLogs(opts, "ProposalAdded", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &AbigenProposalAddedIterator{contract: _Abigen.contract, event: "ProposalAdded", logs: logs, sub: sub}, nil
}

// WatchProposalAdded is a free log subscription operation binding the contract event 0x646fec02522b41e7125cfc859a64fd4f4cefd5dc3b6237ca0abe251ded1fa881.
//
// Solidity: event ProposalAdded(uint256 indexed proposalID, address recipient, uint256 amount, string description)
func (_Abigen *AbigenFilterer) WatchProposalAdded(opts *bind.WatchOpts, sink chan<- *AbigenProposalAdded, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _Abigen.contract.WatchLogs(opts, "ProposalAdded", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbigenProposalAdded)
				if err := _Abigen.contract.UnpackLog(event, "ProposalAdded", log); err != nil {
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
func (_Abigen *AbigenFilterer) ParseProposalAdded(log types.Log) (*AbigenProposalAdded, error) {
	event := new(AbigenProposalAdded)
	if err := _Abigen.contract.UnpackLog(event, "ProposalAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbigenProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Abigen contract.
type AbigenProposalExecutedIterator struct {
	Event *AbigenProposalExecuted // Event containing the contract specifics and raw log

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
func (it *AbigenProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbigenProposalExecuted)
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
		it.Event = new(AbigenProposalExecuted)
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
func (it *AbigenProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbigenProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbigenProposalExecuted represents a ProposalExecuted event raised by the Abigen contract.
type AbigenProposalExecuted struct {
	ProposalID *big.Int
	Result     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x948f4a9cd986f1118c3fbd459f7a22b23c0693e1ca3ef06a6a8be5aa7d39cc03.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, bool result)
func (_Abigen *AbigenFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalID []*big.Int) (*AbigenProposalExecutedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _Abigen.contract.FilterLogs(opts, "ProposalExecuted", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &AbigenProposalExecutedIterator{contract: _Abigen.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x948f4a9cd986f1118c3fbd459f7a22b23c0693e1ca3ef06a6a8be5aa7d39cc03.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, bool result)
func (_Abigen *AbigenFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *AbigenProposalExecuted, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _Abigen.contract.WatchLogs(opts, "ProposalExecuted", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbigenProposalExecuted)
				if err := _Abigen.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
func (_Abigen *AbigenFilterer) ParseProposalExecuted(log types.Log) (*AbigenProposalExecuted, error) {
	event := new(AbigenProposalExecuted)
	if err := _Abigen.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbigenVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Abigen contract.
type AbigenVotedIterator struct {
	Event *AbigenVoted // Event containing the contract specifics and raw log

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
func (it *AbigenVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbigenVoted)
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
		it.Event = new(AbigenVoted)
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
func (it *AbigenVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbigenVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbigenVoted represents a Voted event raised by the Abigen contract.
type AbigenVoted struct {
	ProposalID *big.Int
	InSupport  bool
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed proposalID, bool inSupport, address indexed voter)
func (_Abigen *AbigenFilterer) FilterVoted(opts *bind.FilterOpts, proposalID []*big.Int, voter []common.Address) (*AbigenVotedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Abigen.contract.FilterLogs(opts, "Voted", proposalIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &AbigenVotedIterator{contract: _Abigen.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed proposalID, bool inSupport, address indexed voter)
func (_Abigen *AbigenFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *AbigenVoted, proposalID []*big.Int, voter []common.Address) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Abigen.contract.WatchLogs(opts, "Voted", proposalIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbigenVoted)
				if err := _Abigen.contract.UnpackLog(event, "Voted", log); err != nil {
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
func (_Abigen *AbigenFilterer) ParseVoted(log types.Log) (*AbigenVoted, error) {
	event := new(AbigenVoted)
	if err := _Abigen.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
