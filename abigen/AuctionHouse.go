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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_minBidIncrementPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_daoCut\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"contractIDAO\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoCut\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"contractIDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidIncrementPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"}],\"name\":\"setReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minBidIncrementPercentage\",\"type\":\"uint8\"}],\"name\":\"setMinIncrementPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 startDate, uint256 endDate, uint256 highestBid, address highestBidder, uint256 tokenId, bool settled)
func (_Abigen *AbigenCaller) Auction(opts *bind.CallOpts) (struct {
	StartDate     *big.Int
	EndDate       *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenId       *big.Int
	Settled       bool
}, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "auction")

	outstruct := new(struct {
		StartDate     *big.Int
		EndDate       *big.Int
		HighestBid    *big.Int
		HighestBidder common.Address
		TokenId       *big.Int
		Settled       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartDate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HighestBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Settled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 startDate, uint256 endDate, uint256 highestBid, address highestBidder, uint256 tokenId, bool settled)
func (_Abigen *AbigenSession) Auction() (struct {
	StartDate     *big.Int
	EndDate       *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenId       *big.Int
	Settled       bool
}, error) {
	return _Abigen.Contract.Auction(&_Abigen.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 startDate, uint256 endDate, uint256 highestBid, address highestBidder, uint256 tokenId, bool settled)
func (_Abigen *AbigenCallerSession) Auction() (struct {
	StartDate     *big.Int
	EndDate       *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenId       *big.Int
	Settled       bool
}, error) {
	return _Abigen.Contract.Auction(&_Abigen.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Abigen *AbigenCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Abigen *AbigenSession) Dao() (common.Address, error) {
	return _Abigen.Contract.Dao(&_Abigen.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Abigen *AbigenCallerSession) Dao() (common.Address, error) {
	return _Abigen.Contract.Dao(&_Abigen.CallOpts)
}

// DaoCut is a free data retrieval call binding the contract method 0x1ab3989a.
//
// Solidity: function daoCut() view returns(uint8)
func (_Abigen *AbigenCaller) DaoCut(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "daoCut")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DaoCut is a free data retrieval call binding the contract method 0x1ab3989a.
//
// Solidity: function daoCut() view returns(uint8)
func (_Abigen *AbigenSession) DaoCut() (uint8, error) {
	return _Abigen.Contract.DaoCut(&_Abigen.CallOpts)
}

// DaoCut is a free data retrieval call binding the contract method 0x1ab3989a.
//
// Solidity: function daoCut() view returns(uint8)
func (_Abigen *AbigenCallerSession) DaoCut() (uint8, error) {
	return _Abigen.Contract.DaoCut(&_Abigen.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Abigen *AbigenCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Abigen *AbigenSession) Distributor() (common.Address, error) {
	return _Abigen.Contract.Distributor(&_Abigen.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_Abigen *AbigenCallerSession) Distributor() (common.Address, error) {
	return _Abigen.Contract.Distributor(&_Abigen.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Abigen *AbigenCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Abigen *AbigenSession) Duration() (*big.Int, error) {
	return _Abigen.Contract.Duration(&_Abigen.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_Abigen *AbigenCallerSession) Duration() (*big.Int, error) {
	return _Abigen.Contract.Duration(&_Abigen.CallOpts)
}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_Abigen *AbigenCaller) MinBidIncrementPercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "minBidIncrementPercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_Abigen *AbigenSession) MinBidIncrementPercentage() (uint8, error) {
	return _Abigen.Contract.MinBidIncrementPercentage(&_Abigen.CallOpts)
}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_Abigen *AbigenCallerSession) MinBidIncrementPercentage() (uint8, error) {
	return _Abigen.Contract.MinBidIncrementPercentage(&_Abigen.CallOpts)
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

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Abigen *AbigenCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abigen.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Abigen *AbigenSession) ReservePrice() (*big.Int, error) {
	return _Abigen.Contract.ReservePrice(&_Abigen.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_Abigen *AbigenCallerSession) ReservePrice() (*big.Int, error) {
	return _Abigen.Contract.ReservePrice(&_Abigen.CallOpts)
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

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenID) payable returns()
func (_Abigen *AbigenTransactor) PlaceBid(opts *bind.TransactOpts, _tokenID *big.Int) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "placeBid", _tokenID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenID) payable returns()
func (_Abigen *AbigenSession) PlaceBid(_tokenID *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.PlaceBid(&_Abigen.TransactOpts, _tokenID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenID) payable returns()
func (_Abigen *AbigenTransactorSession) PlaceBid(_tokenID *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.PlaceBid(&_Abigen.TransactOpts, _tokenID)
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

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_Abigen *AbigenTransactor) SetDuration(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "setDuration", _duration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_Abigen *AbigenSession) SetDuration(_duration *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.SetDuration(&_Abigen.TransactOpts, _duration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_Abigen *AbigenTransactorSession) SetDuration(_duration *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.SetDuration(&_Abigen.TransactOpts, _duration)
}

// SetMinIncrementPercentage is a paid mutator transaction binding the contract method 0xee33c374.
//
// Solidity: function setMinIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_Abigen *AbigenTransactor) SetMinIncrementPercentage(opts *bind.TransactOpts, _minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "setMinIncrementPercentage", _minBidIncrementPercentage)
}

// SetMinIncrementPercentage is a paid mutator transaction binding the contract method 0xee33c374.
//
// Solidity: function setMinIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_Abigen *AbigenSession) SetMinIncrementPercentage(_minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _Abigen.Contract.SetMinIncrementPercentage(&_Abigen.TransactOpts, _minBidIncrementPercentage)
}

// SetMinIncrementPercentage is a paid mutator transaction binding the contract method 0xee33c374.
//
// Solidity: function setMinIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_Abigen *AbigenTransactorSession) SetMinIncrementPercentage(_minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _Abigen.Contract.SetMinIncrementPercentage(&_Abigen.TransactOpts, _minBidIncrementPercentage)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_Abigen *AbigenTransactor) SetReservePrice(opts *bind.TransactOpts, _reservePrice *big.Int) (*types.Transaction, error) {
	return _Abigen.contract.Transact(opts, "setReservePrice", _reservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_Abigen *AbigenSession) SetReservePrice(_reservePrice *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.SetReservePrice(&_Abigen.TransactOpts, _reservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_Abigen *AbigenTransactorSession) SetReservePrice(_reservePrice *big.Int) (*types.Transaction, error) {
	return _Abigen.Contract.SetReservePrice(&_Abigen.TransactOpts, _reservePrice)
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
