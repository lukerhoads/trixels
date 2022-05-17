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

// ITrixelsPixel is an auto generated low-level Go binding around an user-defined struct.
type ITrixelsPixel struct {
	X          uint16
	Y          uint16
	Color      [3]byte
	LastEditor common.Address
}

// TrixelsMetaData contains all meta data concerning the Trixels contract.
var TrixelsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes3\",\"name\":\"oldColor\",\"type\":\"bytes3\"},{\"indexed\":true,\"internalType\":\"bytes8\",\"name\":\"newColor\",\"type\":\"bytes8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"PixelChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"colors\",\"outputs\":[{\"internalType\":\"bytes3\",\"name\":\"\",\"type\":\"bytes3\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastEditors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes3\",\"name\":\"newColor\",\"type\":\"bytes3\"}],\"name\":\"changePixel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"x\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"y\",\"type\":\"uint16\"},{\"internalType\":\"bytes3\",\"name\":\"color\",\"type\":\"bytes3\"},{\"internalType\":\"address\",\"name\":\"lastEditor\",\"type\":\"address\"}],\"internalType\":\"structITrixels.Pixel[]\",\"name\":\"newPixels\",\"type\":\"tuple[]\"}],\"name\":\"massChangePixels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610f218061010d6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063775878221161005b57806377587822146100c45780638da5cb5b146100f4578063f2fde38b14610112578063f317d7dd1461012e5761007d565b806304d366611461008257806320cc7af11461009e578063715018a6146100ba575b600080fd5b61009c600480360381019061009791906108a0565b61015e565b005b6100b860048036038101906100b39190610958565b610348565b005b6100c26104fb565b005b6100de60048036038101906100d991906109a5565b610583565b6040516100eb9190610a26565b60405180910390f35b6100fc6105cf565b6040516101099190610a26565b60405180910390f35b61012c60048036038101906101279190610a6d565b6105f8565b005b610148600480360381019061014391906109a5565b6106f0565b6040516101559190610aa9565b60405180910390f35b60c883106101a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019890610b21565b60405180910390fd5b60c882106101e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101db90610b8d565b60405180910390fd5b600060018460c881106101fa576101f9610bad565b5b601402018360c881106102105761020f610bad565b5b600a91828204019190066003029054906101000a900460e81b90508160018560c881106102405761023f610bad565b5b601402018460c8811061025657610255610bad565b5b600a91828204019190066003026101000a81548162ffffff021916908360e81c021790555033610fa18560c8811061029157610290610bad565b5b60c802018460c881106102a7576102a6610bad565b5b0160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550817cffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191683857f4ed8888d7862eb69fdce4cb367605f01ea41de731b9f2f0293044aa598ae1097843360405161033a929190610bdc565b60405180910390a450505050565b610350610732565b73ffffffffffffffffffffffffffffffffffffffff1661036e6105cf565b73ffffffffffffffffffffffffffffffffffffffff16146103c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bb90610c51565b60405180910390fd5b60005b828290508110156104f65760008383838181106103e7576103e6610bad565b5b9050608002018036038101906103fd9190610db4565b905080604001516001826000015161ffff1660c881106104205761041f610bad565b5b60140201826020015161ffff1660c8811061043e5761043d610bad565b5b600a91828204019190066003026101000a81548162ffffff021916908360e81c02179055508060600151610fa1826000015161ffff1660c8811061048557610484610bad565b5b60c80201826020015161ffff1660c881106104a3576104a2610bad565b5b0160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505080806104ee90610e10565b9150506103c7565b505050565b610503610732565b73ffffffffffffffffffffffffffffffffffffffff166105216105cf565b73ffffffffffffffffffffffffffffffffffffffff1614610577576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056e90610c51565b60405180910390fd5b610581600061073a565b565b610fa18260c8811061059457600080fd5b60c802018160c881106105a657600080fd5b016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610600610732565b73ffffffffffffffffffffffffffffffffffffffff1661061e6105cf565b73ffffffffffffffffffffffffffffffffffffffff1614610674576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066b90610c51565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156106e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106db90610ecb565b60405180910390fd5b6106ed8161073a565b50565b60018260c8811061070057600080fd5b601402018160c8811061071257600080fd5b600a9182820401919006600302915091509054906101000a900460e81b81565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61082581610812565b811461083057600080fd5b50565b6000813590506108428161081c565b92915050565b60007fffffff000000000000000000000000000000000000000000000000000000000082169050919050565b61087d81610848565b811461088857600080fd5b50565b60008135905061089a81610874565b92915050565b6000806000606084860312156108b9576108b8610808565b5b60006108c786828701610833565b93505060206108d886828701610833565b92505060406108e98682870161088b565b9150509250925092565b600080fd5b600080fd5b600080fd5b60008083601f840112610918576109176108f3565b5b8235905067ffffffffffffffff811115610935576109346108f8565b5b602083019150836080820283011115610951576109506108fd565b5b9250929050565b6000806020838503121561096f5761096e610808565b5b600083013567ffffffffffffffff81111561098d5761098c61080d565b5b61099985828601610902565b92509250509250929050565b600080604083850312156109bc576109bb610808565b5b60006109ca85828601610833565b92505060206109db85828601610833565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a10826109e5565b9050919050565b610a2081610a05565b82525050565b6000602082019050610a3b6000830184610a17565b92915050565b610a4a81610a05565b8114610a5557600080fd5b50565b600081359050610a6781610a41565b92915050565b600060208284031215610a8357610a82610808565b5b6000610a9184828501610a58565b91505092915050565b610aa381610848565b82525050565b6000602082019050610abe6000830184610a9a565b92915050565b600082825260208201905092915050565b7f7820636f6f726420696e76616c69640000000000000000000000000000000000600082015250565b6000610b0b600f83610ac4565b9150610b1682610ad5565b602082019050919050565b60006020820190508181036000830152610b3a81610afe565b9050919050565b7f7920636f6f726420696e76616c69640000000000000000000000000000000000600082015250565b6000610b77600f83610ac4565b9150610b8282610b41565b602082019050919050565b60006020820190508181036000830152610ba681610b6a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000604082019050610bf16000830185610a9a565b610bfe6020830184610a17565b9392505050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610c3b602083610ac4565b9150610c4682610c05565b602082019050919050565b60006020820190508181036000830152610c6a81610c2e565b9050919050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610cbf82610c76565b810181811067ffffffffffffffff82111715610cde57610cdd610c87565b5b80604052505050565b6000610cf16107fe565b9050610cfd8282610cb6565b919050565b600061ffff82169050919050565b610d1981610d02565b8114610d2457600080fd5b50565b600081359050610d3681610d10565b92915050565b600060808284031215610d5257610d51610c71565b5b610d5c6080610ce7565b90506000610d6c84828501610d27565b6000830152506020610d8084828501610d27565b6020830152506040610d948482850161088b565b6040830152506060610da884828501610a58565b60608301525092915050565b600060808284031215610dca57610dc9610808565b5b6000610dd884828501610d3c565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e1b82610812565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610e4e57610e4d610de1565b5b600182019050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610eb5602683610ac4565b9150610ec082610e59565b604082019050919050565b60006020820190508181036000830152610ee481610ea8565b905091905056fea26469706673582212207777e5d249874221a4045a5e5de204b316a7638a45e13358fb30d3962d7e131464736f6c634300080a0033",
}

// TrixelsABI is the input ABI used to generate the binding from.
// Deprecated: Use TrixelsMetaData.ABI instead.
var TrixelsABI = TrixelsMetaData.ABI

// TrixelsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TrixelsMetaData.Bin instead.
var TrixelsBin = TrixelsMetaData.Bin

// DeployTrixels deploys a new Ethereum contract, binding an instance of Trixels to it.
func DeployTrixels(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trixels, error) {
	parsed, err := TrixelsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TrixelsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trixels{TrixelsCaller: TrixelsCaller{contract: contract}, TrixelsTransactor: TrixelsTransactor{contract: contract}, TrixelsFilterer: TrixelsFilterer{contract: contract}}, nil
}

// Trixels is an auto generated Go binding around an Ethereum contract.
type Trixels struct {
	TrixelsCaller     // Read-only binding to the contract
	TrixelsTransactor // Write-only binding to the contract
	TrixelsFilterer   // Log filterer for contract events
}

// TrixelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrixelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrixelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrixelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrixelsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrixelsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrixelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrixelsSession struct {
	Contract     *Trixels          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrixelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrixelsCallerSession struct {
	Contract *TrixelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TrixelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrixelsTransactorSession struct {
	Contract     *TrixelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TrixelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrixelsRaw struct {
	Contract *Trixels // Generic contract binding to access the raw methods on
}

// TrixelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrixelsCallerRaw struct {
	Contract *TrixelsCaller // Generic read-only contract binding to access the raw methods on
}

// TrixelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrixelsTransactorRaw struct {
	Contract *TrixelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrixels creates a new instance of Trixels, bound to a specific deployed contract.
func NewTrixels(address common.Address, backend bind.ContractBackend) (*Trixels, error) {
	contract, err := bindTrixels(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trixels{TrixelsCaller: TrixelsCaller{contract: contract}, TrixelsTransactor: TrixelsTransactor{contract: contract}, TrixelsFilterer: TrixelsFilterer{contract: contract}}, nil
}

// NewTrixelsCaller creates a new read-only instance of Trixels, bound to a specific deployed contract.
func NewTrixelsCaller(address common.Address, caller bind.ContractCaller) (*TrixelsCaller, error) {
	contract, err := bindTrixels(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrixelsCaller{contract: contract}, nil
}

// NewTrixelsTransactor creates a new write-only instance of Trixels, bound to a specific deployed contract.
func NewTrixelsTransactor(address common.Address, transactor bind.ContractTransactor) (*TrixelsTransactor, error) {
	contract, err := bindTrixels(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrixelsTransactor{contract: contract}, nil
}

// NewTrixelsFilterer creates a new log filterer instance of Trixels, bound to a specific deployed contract.
func NewTrixelsFilterer(address common.Address, filterer bind.ContractFilterer) (*TrixelsFilterer, error) {
	contract, err := bindTrixels(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrixelsFilterer{contract: contract}, nil
}

// bindTrixels binds a generic wrapper to an already deployed contract.
func bindTrixels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrixelsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trixels *TrixelsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trixels.Contract.TrixelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trixels *TrixelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trixels.Contract.TrixelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trixels *TrixelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trixels.Contract.TrixelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trixels *TrixelsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trixels.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trixels *TrixelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trixels.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trixels *TrixelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trixels.Contract.contract.Transact(opts, method, params...)
}

// Colors is a free data retrieval call binding the contract method 0xf317d7dd.
//
// Solidity: function colors(uint256 , uint256 ) view returns(bytes3)
func (_Trixels *TrixelsCaller) Colors(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([3]byte, error) {
	var out []interface{}
	err := _Trixels.contract.Call(opts, &out, "colors", arg0, arg1)

	if err != nil {
		return *new([3]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([3]byte)).(*[3]byte)

	return out0, err

}

// Colors is a free data retrieval call binding the contract method 0xf317d7dd.
//
// Solidity: function colors(uint256 , uint256 ) view returns(bytes3)
func (_Trixels *TrixelsSession) Colors(arg0 *big.Int, arg1 *big.Int) ([3]byte, error) {
	return _Trixels.Contract.Colors(&_Trixels.CallOpts, arg0, arg1)
}

// Colors is a free data retrieval call binding the contract method 0xf317d7dd.
//
// Solidity: function colors(uint256 , uint256 ) view returns(bytes3)
func (_Trixels *TrixelsCallerSession) Colors(arg0 *big.Int, arg1 *big.Int) ([3]byte, error) {
	return _Trixels.Contract.Colors(&_Trixels.CallOpts, arg0, arg1)
}

// LastEditors is a free data retrieval call binding the contract method 0x77587822.
//
// Solidity: function lastEditors(uint256 , uint256 ) view returns(address)
func (_Trixels *TrixelsCaller) LastEditors(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Trixels.contract.Call(opts, &out, "lastEditors", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LastEditors is a free data retrieval call binding the contract method 0x77587822.
//
// Solidity: function lastEditors(uint256 , uint256 ) view returns(address)
func (_Trixels *TrixelsSession) LastEditors(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Trixels.Contract.LastEditors(&_Trixels.CallOpts, arg0, arg1)
}

// LastEditors is a free data retrieval call binding the contract method 0x77587822.
//
// Solidity: function lastEditors(uint256 , uint256 ) view returns(address)
func (_Trixels *TrixelsCallerSession) LastEditors(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Trixels.Contract.LastEditors(&_Trixels.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trixels *TrixelsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Trixels.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trixels *TrixelsSession) Owner() (common.Address, error) {
	return _Trixels.Contract.Owner(&_Trixels.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Trixels *TrixelsCallerSession) Owner() (common.Address, error) {
	return _Trixels.Contract.Owner(&_Trixels.CallOpts)
}

// ChangePixel is a paid mutator transaction binding the contract method 0x04d36661.
//
// Solidity: function changePixel(uint256 x, uint256 y, bytes3 newColor) returns()
func (_Trixels *TrixelsTransactor) ChangePixel(opts *bind.TransactOpts, x *big.Int, y *big.Int, newColor [3]byte) (*types.Transaction, error) {
	return _Trixels.contract.Transact(opts, "changePixel", x, y, newColor)
}

// ChangePixel is a paid mutator transaction binding the contract method 0x04d36661.
//
// Solidity: function changePixel(uint256 x, uint256 y, bytes3 newColor) returns()
func (_Trixels *TrixelsSession) ChangePixel(x *big.Int, y *big.Int, newColor [3]byte) (*types.Transaction, error) {
	return _Trixels.Contract.ChangePixel(&_Trixels.TransactOpts, x, y, newColor)
}

// ChangePixel is a paid mutator transaction binding the contract method 0x04d36661.
//
// Solidity: function changePixel(uint256 x, uint256 y, bytes3 newColor) returns()
func (_Trixels *TrixelsTransactorSession) ChangePixel(x *big.Int, y *big.Int, newColor [3]byte) (*types.Transaction, error) {
	return _Trixels.Contract.ChangePixel(&_Trixels.TransactOpts, x, y, newColor)
}

// MassChangePixels is a paid mutator transaction binding the contract method 0x20cc7af1.
//
// Solidity: function massChangePixels((uint16,uint16,bytes3,address)[] newPixels) returns()
func (_Trixels *TrixelsTransactor) MassChangePixels(opts *bind.TransactOpts, newPixels []ITrixelsPixel) (*types.Transaction, error) {
	return _Trixels.contract.Transact(opts, "massChangePixels", newPixels)
}

// MassChangePixels is a paid mutator transaction binding the contract method 0x20cc7af1.
//
// Solidity: function massChangePixels((uint16,uint16,bytes3,address)[] newPixels) returns()
func (_Trixels *TrixelsSession) MassChangePixels(newPixels []ITrixelsPixel) (*types.Transaction, error) {
	return _Trixels.Contract.MassChangePixels(&_Trixels.TransactOpts, newPixels)
}

// MassChangePixels is a paid mutator transaction binding the contract method 0x20cc7af1.
//
// Solidity: function massChangePixels((uint16,uint16,bytes3,address)[] newPixels) returns()
func (_Trixels *TrixelsTransactorSession) MassChangePixels(newPixels []ITrixelsPixel) (*types.Transaction, error) {
	return _Trixels.Contract.MassChangePixels(&_Trixels.TransactOpts, newPixels)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trixels *TrixelsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trixels.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trixels *TrixelsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trixels.Contract.RenounceOwnership(&_Trixels.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Trixels *TrixelsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Trixels.Contract.RenounceOwnership(&_Trixels.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trixels *TrixelsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trixels.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trixels *TrixelsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trixels.Contract.TransferOwnership(&_Trixels.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trixels *TrixelsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trixels.Contract.TransferOwnership(&_Trixels.TransactOpts, newOwner)
}

// TrixelsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Trixels contract.
type TrixelsOwnershipTransferredIterator struct {
	Event *TrixelsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TrixelsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsOwnershipTransferred)
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
		it.Event = new(TrixelsOwnershipTransferred)
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
func (it *TrixelsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsOwnershipTransferred represents a OwnershipTransferred event raised by the Trixels contract.
type TrixelsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trixels *TrixelsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TrixelsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trixels.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TrixelsOwnershipTransferredIterator{contract: _Trixels.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Trixels *TrixelsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TrixelsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Trixels.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsOwnershipTransferred)
				if err := _Trixels.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Trixels *TrixelsFilterer) ParseOwnershipTransferred(log types.Log) (*TrixelsOwnershipTransferred, error) {
	event := new(TrixelsOwnershipTransferred)
	if err := _Trixels.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsPixelChangedIterator is returned from FilterPixelChanged and is used to iterate over the raw logs and unpacked data for PixelChanged events raised by the Trixels contract.
type TrixelsPixelChangedIterator struct {
	Event *TrixelsPixelChanged // Event containing the contract specifics and raw log

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
func (it *TrixelsPixelChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsPixelChanged)
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
		it.Event = new(TrixelsPixelChanged)
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
func (it *TrixelsPixelChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsPixelChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsPixelChanged represents a PixelChanged event raised by the Trixels contract.
type TrixelsPixelChanged struct {
	X        *big.Int
	Y        *big.Int
	OldColor [3]byte
	NewColor [8]byte
	User     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPixelChanged is a free log retrieval operation binding the contract event 0x4ed8888d7862eb69fdce4cb367605f01ea41de731b9f2f0293044aa598ae1097.
//
// Solidity: event PixelChanged(uint256 indexed x, uint256 indexed y, bytes3 oldColor, bytes8 indexed newColor, address user)
func (_Trixels *TrixelsFilterer) FilterPixelChanged(opts *bind.FilterOpts, x []*big.Int, y []*big.Int, newColor [][8]byte) (*TrixelsPixelChangedIterator, error) {

	var xRule []interface{}
	for _, xItem := range x {
		xRule = append(xRule, xItem)
	}
	var yRule []interface{}
	for _, yItem := range y {
		yRule = append(yRule, yItem)
	}

	var newColorRule []interface{}
	for _, newColorItem := range newColor {
		newColorRule = append(newColorRule, newColorItem)
	}

	logs, sub, err := _Trixels.contract.FilterLogs(opts, "PixelChanged", xRule, yRule, newColorRule)
	if err != nil {
		return nil, err
	}
	return &TrixelsPixelChangedIterator{contract: _Trixels.contract, event: "PixelChanged", logs: logs, sub: sub}, nil
}

// WatchPixelChanged is a free log subscription operation binding the contract event 0x4ed8888d7862eb69fdce4cb367605f01ea41de731b9f2f0293044aa598ae1097.
//
// Solidity: event PixelChanged(uint256 indexed x, uint256 indexed y, bytes3 oldColor, bytes8 indexed newColor, address user)
func (_Trixels *TrixelsFilterer) WatchPixelChanged(opts *bind.WatchOpts, sink chan<- *TrixelsPixelChanged, x []*big.Int, y []*big.Int, newColor [][8]byte) (event.Subscription, error) {

	var xRule []interface{}
	for _, xItem := range x {
		xRule = append(xRule, xItem)
	}
	var yRule []interface{}
	for _, yItem := range y {
		yRule = append(yRule, yItem)
	}

	var newColorRule []interface{}
	for _, newColorItem := range newColor {
		newColorRule = append(newColorRule, newColorItem)
	}

	logs, sub, err := _Trixels.contract.WatchLogs(opts, "PixelChanged", xRule, yRule, newColorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsPixelChanged)
				if err := _Trixels.contract.UnpackLog(event, "PixelChanged", log); err != nil {
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

// ParsePixelChanged is a log parse operation binding the contract event 0x4ed8888d7862eb69fdce4cb367605f01ea41de731b9f2f0293044aa598ae1097.
//
// Solidity: event PixelChanged(uint256 indexed x, uint256 indexed y, bytes3 oldColor, bytes8 indexed newColor, address user)
func (_Trixels *TrixelsFilterer) ParsePixelChanged(log types.Log) (*TrixelsPixelChanged, error) {
	event := new(TrixelsPixelChanged)
	if err := _Trixels.contract.UnpackLog(event, "PixelChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
