// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Distributor

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

// DistributorMetaData contains all meta data concerning the Distributor contract.
var DistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sales\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"}],\"name\":\"recordSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numContribs\",\"type\":\"uint256\"}],\"internalType\":\"structIDistributor.Contributor[]\",\"name\":\"_contributors\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200122938038062001229833981810160405281019062000037919062000184565b80620000586200004c620000a160201b60201c565b620000a960201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000209565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000815190506200017e81620001ef565b92915050565b6000602082840312156200019d576200019c620001ea565b5b6000620001ad848285016200016d565b91505092915050565b6000620001c382620001ca565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b620001fa81620001b6565b81146200020657600080fd5b50565b61101080620002196000396000f3fe60806040526004361061007b5760003560e01c80638da5cb5b1161004e5780638da5cb5b146101075780638daf859614610132578063b5f522f71461015b578063f2fde38b146101985761007b565b806327e235e3146100805780633ccfd60b146100bd5780636f27f690146100d4578063715018a6146100f0575b600080fd5b34801561008c57600080fd5b506100a760048036038101906100a2919061099a565b6101c1565b6040516100b49190610be0565b60405180910390f35b3480156100c957600080fd5b506100d26101d9565b005b6100ee60048036038101906100e99190610a50565b6102a6565b005b3480156100fc57600080fd5b506101056102c1565b005b34801561011357600080fd5b5061011c610349565b6040516101299190610b3c565b60405180910390f35b34801561013e57600080fd5b50610159600480360381019061015491906109c7565b610372565b005b34801561016757600080fd5b50610182600480360381019061017d9190610a50565b6104cf565b60405161018f9190610be0565b60405180910390f35b3480156101a457600080fd5b506101bf60048036038101906101ba919061099a565b6104e7565b005b60036020528060005260406000206000915090505481565b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541161025b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025290610bc0565b60405180910390fd5b6102a433600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105df565b565b34600260008381526020019081526020016000208190555050565b6102c9610725565b73ffffffffffffffffffffffffffffffffffffffff166102e7610349565b73ffffffffffffffffffffffffffffffffffffffff161461033d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033490610ba0565b60405180910390fd5b610347600061072d565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61037a610725565b73ffffffffffffffffffffffffffffffffffffffff16610398610349565b73ffffffffffffffffffffffffffffffffffffffff16146103ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e590610ba0565b60405180910390fd5b60005b82518110156104ca57610384600260008481526020019081526020016000205484838151811061042457610423610e69565b5b60200260200101516020015161043a9190610cef565b6104449190610cbe565b6003600085848151811061045b5761045a610e69565b5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104b09190610c68565b9250508190555080806104c290610dc2565b9150506103f1565b505050565b60026020528060005260406000206000915090505481565b6104ef610725565b73ffffffffffffffffffffffffffffffffffffffff1661050d610349565b73ffffffffffffffffffffffffffffffffffffffff1614610563576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055a90610ba0565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156105d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ca90610b80565b60405180910390fd5b6105dc8161072d565b50565b6105e982826107f1565b61072157600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561065757600080fd5b505af115801561066b573d6000803e3d6000fd5b5050505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b81526004016106cd929190610b57565b602060405180830381600087803b1580156106e757600080fd5b505af11580156106fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071f9190610a23565b505b5050565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000808373ffffffffffffffffffffffffffffffffffffffff16836175309060405161081c90610b27565b600060405180830381858888f193505050503d806000811461085a576040519150601f19603f3d011682016040523d82523d6000602084013e61085f565b606091505b505090508091505092915050565b600061088061087b84610c20565b610bfb565b905080838252602082019050828560408602820111156108a3576108a2610ed1565b5b60005b858110156108d357816108b98882610935565b8452602084019350604083019250506001810190506108a6565b5050509392505050565b6000813590506108ec81610f95565b92915050565b600082601f83011261090757610906610ec7565b5b813561091784826020860161086d565b91505092915050565b60008151905061092f81610fac565b92915050565b60006040828403121561094b5761094a610ecc565b5b6109556040610bfb565b90506000610965848285016108dd565b600083015250602061097984828501610985565b60208301525092915050565b60008135905061099481610fc3565b92915050565b6000602082840312156109b0576109af610edb565b5b60006109be848285016108dd565b91505092915050565b600080604083850312156109de576109dd610edb565b5b600083013567ffffffffffffffff8111156109fc576109fb610ed6565b5b610a08858286016108f2565b9250506020610a1985828601610985565b9150509250929050565b600060208284031215610a3957610a38610edb565b5b6000610a4784828501610920565b91505092915050565b600060208284031215610a6657610a65610edb565b5b6000610a7484828501610985565b91505092915050565b610a8681610d49565b82525050565b6000610a99602683610c57565b9150610aa482610ef1565b604082019050919050565b6000610abc602083610c57565b9150610ac782610f40565b602082019050919050565b6000610adf601683610c57565b9150610aea82610f69565b602082019050919050565b6000610b02600083610c4c565b9150610b0d82610f92565b600082019050919050565b610b2181610d87565b82525050565b6000610b3282610af5565b9150819050919050565b6000602082019050610b516000830184610a7d565b92915050565b6000604082019050610b6c6000830185610a7d565b610b796020830184610b18565b9392505050565b60006020820190508181036000830152610b9981610a8c565b9050919050565b60006020820190508181036000830152610bb981610aaf565b9050919050565b60006020820190508181036000830152610bd981610ad2565b9050919050565b6000602082019050610bf56000830184610b18565b92915050565b6000610c05610c16565b9050610c118282610d91565b919050565b6000604051905090565b600067ffffffffffffffff821115610c3b57610c3a610e98565b5b602082029050602081019050919050565b600081905092915050565b600082825260208201905092915050565b6000610c7382610d87565b9150610c7e83610d87565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610cb357610cb2610e0b565b5b828201905092915050565b6000610cc982610d87565b9150610cd483610d87565b925082610ce457610ce3610e3a565b5b828204905092915050565b6000610cfa82610d87565b9150610d0583610d87565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615610d3e57610d3d610e0b565b5b828202905092915050565b6000610d5482610d67565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b610d9a82610ee0565b810181811067ffffffffffffffff82111715610db957610db8610e98565b5b80604052505050565b6000610dcd82610d87565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610e0057610dff610e0b565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f4e6f2062616c616e636520746f20776974686472617700000000000000000000600082015250565b50565b610f9e81610d49565b8114610fa957600080fd5b50565b610fb581610d5b565b8114610fc057600080fd5b50565b610fcc81610d87565b8114610fd757600080fd5b5056fea264697066735822122095a251cfbed21aa56176b5ef7d34f3d6c2bd54a0371111b3fc91d596c703b5b364736f6c63430008060033",
}

// DistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use DistributorMetaData.ABI instead.
var DistributorABI = DistributorMetaData.ABI

// DistributorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DistributorMetaData.Bin instead.
var DistributorBin = DistributorMetaData.Bin

// DeployDistributor deploys a new Ethereum contract, binding an instance of Distributor to it.
func DeployDistributor(auth *bind.TransactOpts, backend bind.ContractBackend, _weth common.Address) (common.Address, *types.Transaction, *Distributor, error) {
	parsed, err := DistributorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DistributorBin), backend, _weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Distributor{DistributorCaller: DistributorCaller{contract: contract}, DistributorTransactor: DistributorTransactor{contract: contract}, DistributorFilterer: DistributorFilterer{contract: contract}}, nil
}

// Distributor is an auto generated Go binding around an Ethereum contract.
type Distributor struct {
	DistributorCaller     // Read-only binding to the contract
	DistributorTransactor // Write-only binding to the contract
	DistributorFilterer   // Log filterer for contract events
}

// DistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistributorSession struct {
	Contract     *Distributor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistributorCallerSession struct {
	Contract *DistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistributorTransactorSession struct {
	Contract     *DistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistributorRaw struct {
	Contract *Distributor // Generic contract binding to access the raw methods on
}

// DistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistributorCallerRaw struct {
	Contract *DistributorCaller // Generic read-only contract binding to access the raw methods on
}

// DistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistributorTransactorRaw struct {
	Contract *DistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistributor creates a new instance of Distributor, bound to a specific deployed contract.
func NewDistributor(address common.Address, backend bind.ContractBackend) (*Distributor, error) {
	contract, err := bindDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Distributor{DistributorCaller: DistributorCaller{contract: contract}, DistributorTransactor: DistributorTransactor{contract: contract}, DistributorFilterer: DistributorFilterer{contract: contract}}, nil
}

// NewDistributorCaller creates a new read-only instance of Distributor, bound to a specific deployed contract.
func NewDistributorCaller(address common.Address, caller bind.ContractCaller) (*DistributorCaller, error) {
	contract, err := bindDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistributorCaller{contract: contract}, nil
}

// NewDistributorTransactor creates a new write-only instance of Distributor, bound to a specific deployed contract.
func NewDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*DistributorTransactor, error) {
	contract, err := bindDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistributorTransactor{contract: contract}, nil
}

// NewDistributorFilterer creates a new log filterer instance of Distributor, bound to a specific deployed contract.
func NewDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*DistributorFilterer, error) {
	contract, err := bindDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistributorFilterer{contract: contract}, nil
}

// bindDistributor binds a generic wrapper to an already deployed contract.
func bindDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributor *DistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distributor.Contract.DistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributor *DistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.Contract.DistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributor *DistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributor.Contract.DistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Distributor *DistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Distributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Distributor *DistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Distributor *DistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Distributor.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Distributor *DistributorCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Distributor *DistributorSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Distributor.Contract.Balances(&_Distributor.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Distributor *DistributorCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Distributor.Contract.Balances(&_Distributor.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distributor *DistributorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distributor *DistributorSession) Owner() (common.Address, error) {
	return _Distributor.Contract.Owner(&_Distributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Distributor *DistributorCallerSession) Owner() (common.Address, error) {
	return _Distributor.Contract.Owner(&_Distributor.CallOpts)
}

// Sales is a free data retrieval call binding the contract method 0xb5f522f7.
//
// Solidity: function sales(uint256 ) view returns(uint256)
func (_Distributor *DistributorCaller) Sales(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Distributor.contract.Call(opts, &out, "sales", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sales is a free data retrieval call binding the contract method 0xb5f522f7.
//
// Solidity: function sales(uint256 ) view returns(uint256)
func (_Distributor *DistributorSession) Sales(arg0 *big.Int) (*big.Int, error) {
	return _Distributor.Contract.Sales(&_Distributor.CallOpts, arg0)
}

// Sales is a free data retrieval call binding the contract method 0xb5f522f7.
//
// Solidity: function sales(uint256 ) view returns(uint256)
func (_Distributor *DistributorCallerSession) Sales(arg0 *big.Int) (*big.Int, error) {
	return _Distributor.Contract.Sales(&_Distributor.CallOpts, arg0)
}

// Distribute is a paid mutator transaction binding the contract method 0x8daf8596.
//
// Solidity: function distribute((address,uint256)[] _contributors, uint256 _tokenID) returns()
func (_Distributor *DistributorTransactor) Distribute(opts *bind.TransactOpts, _contributors []IDistributorContributor, _tokenID *big.Int) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "distribute", _contributors, _tokenID)
}

// Distribute is a paid mutator transaction binding the contract method 0x8daf8596.
//
// Solidity: function distribute((address,uint256)[] _contributors, uint256 _tokenID) returns()
func (_Distributor *DistributorSession) Distribute(_contributors []IDistributorContributor, _tokenID *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.Distribute(&_Distributor.TransactOpts, _contributors, _tokenID)
}

// Distribute is a paid mutator transaction binding the contract method 0x8daf8596.
//
// Solidity: function distribute((address,uint256)[] _contributors, uint256 _tokenID) returns()
func (_Distributor *DistributorTransactorSession) Distribute(_contributors []IDistributorContributor, _tokenID *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.Distribute(&_Distributor.TransactOpts, _contributors, _tokenID)
}

// RecordSale is a paid mutator transaction binding the contract method 0x6f27f690.
//
// Solidity: function recordSale(uint256 _tokenID) payable returns()
func (_Distributor *DistributorTransactor) RecordSale(opts *bind.TransactOpts, _tokenID *big.Int) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "recordSale", _tokenID)
}

// RecordSale is a paid mutator transaction binding the contract method 0x6f27f690.
//
// Solidity: function recordSale(uint256 _tokenID) payable returns()
func (_Distributor *DistributorSession) RecordSale(_tokenID *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.RecordSale(&_Distributor.TransactOpts, _tokenID)
}

// RecordSale is a paid mutator transaction binding the contract method 0x6f27f690.
//
// Solidity: function recordSale(uint256 _tokenID) payable returns()
func (_Distributor *DistributorTransactorSession) RecordSale(_tokenID *big.Int) (*types.Transaction, error) {
	return _Distributor.Contract.RecordSale(&_Distributor.TransactOpts, _tokenID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distributor *DistributorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distributor *DistributorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Distributor.Contract.RenounceOwnership(&_Distributor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Distributor *DistributorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Distributor.Contract.RenounceOwnership(&_Distributor.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distributor *DistributorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distributor *DistributorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Distributor.Contract.TransferOwnership(&_Distributor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Distributor *DistributorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Distributor.Contract.TransferOwnership(&_Distributor.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Distributor *DistributorTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Distributor.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Distributor *DistributorSession) Withdraw() (*types.Transaction, error) {
	return _Distributor.Contract.Withdraw(&_Distributor.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Distributor *DistributorTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Distributor.Contract.Withdraw(&_Distributor.TransactOpts)
}

// DistributorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Distributor contract.
type DistributorOwnershipTransferredIterator struct {
	Event *DistributorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DistributorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistributorOwnershipTransferred)
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
		it.Event = new(DistributorOwnershipTransferred)
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
func (it *DistributorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistributorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistributorOwnershipTransferred represents a OwnershipTransferred event raised by the Distributor contract.
type DistributorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Distributor *DistributorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DistributorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Distributor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DistributorOwnershipTransferredIterator{contract: _Distributor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Distributor *DistributorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DistributorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Distributor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistributorOwnershipTransferred)
				if err := _Distributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Distributor *DistributorFilterer) ParseOwnershipTransferred(log types.Log) (*DistributorOwnershipTransferred, error) {
	event := new(DistributorOwnershipTransferred)
	if err := _Distributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
