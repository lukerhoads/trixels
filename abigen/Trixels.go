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

// TrixelsPixel is an auto generated low-level Go binding around an user-defined struct.
type TrixelsPixel struct {
	X          uint16
	Y          uint16
	Color      [8]byte
	LastEditor common.Address
}

// TrixelsMetaData contains all meta data concerning the Trixels contract.
var TrixelsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes8[]\",\"name\":\"newColors\",\"type\":\"bytes8[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"bytes8\",\"name\":\"newColor\",\"type\":\"bytes8\"}],\"name\":\"changePixel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"x\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"y\",\"type\":\"uint16\"},{\"internalType\":\"bytes8\",\"name\":\"color\",\"type\":\"bytes8\"},{\"internalType\":\"address\",\"name\":\"lastEditor\",\"type\":\"address\"}],\"internalType\":\"structTrixels.Pixel[]\",\"name\":\"newPixels\",\"type\":\"tuple[]\"}],\"name\":\"massUpdatePixels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"eab6402b": "changePixel(uint256,uint256,bytes8)",
		"452955fd": "massUpdatePixels((uint16,uint16,bytes8,address)[])",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516200078f3803806200078f8339810160408190526100319161019f565b600080546001600160a01b031916331781555b60c88110156101655760005b60c88210156101525760008161006760c885610279565b6100719190610298565b9050630233030360e41b60c88210156100a157848281518110610096576100966102b0565b602002602001015190505b8060018560c881106100b5576100b56102b0565b603202018460c881106100ca576100ca6102b0565b600491828204019190066008026101000a8154816001600160401b03021916908360c01c021790555060006127118560c88110610109576101096102b0565b60c802018460c8811061011e5761011e6102b0565b0180546001600160a01b0319166001600160a01b03929092169190911790555081905061014a816102c6565b915050610050565b508061015d816102c6565b915050610044565b50506102e1565b634e487b7160e01b600052604160045260246000fd5b80516001600160c01b03198116811461019a57600080fd5b919050565b600060208083850312156101b257600080fd5b82516001600160401b03808211156101c957600080fd5b818501915085601f8301126101dd57600080fd5b8151818111156101ef576101ef61016c565b8060051b604051601f19603f830116810181811085821117156102145761021461016c565b60405291825284820192508381018501918883111561023257600080fd5b938501935b828510156102575761024885610182565b84529385019392850192610237565b98975050505050505050565b634e487b7160e01b600052601160045260246000fd5b600081600019048311821515161561029357610293610263565b500290565b600082198211156102ab576102ab610263565b500190565b634e487b7160e01b600052603260045260246000fd5b60006000198214156102da576102da610263565b5060010190565b61049e80620002f16000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063452955fd1461003b578063eab6402b14610050575b600080fd5b61004e6100493660046102b7565b610063565b005b61004e61005e366004610349565b61018d565b6000546001600160a01b0316331461007a57600080fd5b60005b818110156101885760008383838181106100995761009961037e565b9050608002018036038101906100af91906103a6565b905080604001516001826000015161ffff1660c881106100d1576100d161037e565b60320201826020015161ffff1660c881106100ee576100ee61037e565b600491828204019190066008026101000a81548167ffffffffffffffff021916908360c01c02179055508060600151612711826000015161ffff1660c881106101395761013961037e565b60c80201826020015161ffff1660c881106101565761015661037e565b0180546001600160a01b0319166001600160a01b039290921691909117905550806101808161043f565b91505061007d565b505050565b60c883106101d45760405162461bcd60e51b815260206004820152600f60248201526e1e0818dbdbdc99081a5b9d985b1a59608a1b60448201526064015b60405180910390fd5b60c882106102165760405162461bcd60e51b815260206004820152600f60248201526e1e4818dbdbdc99081a5b9d985b1a59608a1b60448201526064016101cb565b8060018460c8811061022a5761022a61037e565b603202018360c8811061023f5761023f61037e565b600491828204019190066008026101000a81548167ffffffffffffffff021916908360c01c0217905550336127118460c8811061027e5761027e61037e565b60c802018360c881106102935761029361037e565b0180546001600160a01b0319166001600160a01b0392909216919091179055505050565b600080602083850312156102ca57600080fd5b823567ffffffffffffffff808211156102e257600080fd5b818501915085601f8301126102f657600080fd5b81358181111561030557600080fd5b8660208260071b850101111561031a57600080fd5b60209290920196919550909350505050565b80356001600160c01b03198116811461034457600080fd5b919050565b60008060006060848603121561035e57600080fd5b83359250602084013591506103756040850161032c565b90509250925092565b634e487b7160e01b600052603260045260246000fd5b803561ffff8116811461034457600080fd5b6000608082840312156103b857600080fd5b6040516080810181811067ffffffffffffffff821117156103e957634e487b7160e01b600052604160045260246000fd5b6040526103f583610394565b815261040360208401610394565b60208201526104146040840161032c565b604082015260608301356001600160a01b038116811461043357600080fd5b60608201529392505050565b600060001982141561046157634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220df0787fc76ef944f170709ac72cb8ccc5fde95bbaa94b876d549b3209a8c423764736f6c634300080a0033",
}

// TrixelsABI is the input ABI used to generate the binding from.
// Deprecated: Use TrixelsMetaData.ABI instead.
var TrixelsABI = TrixelsMetaData.ABI

// Deprecated: Use TrixelsMetaData.Sigs instead.
// TrixelsFuncSigs maps the 4-byte function signature to its string representation.
var TrixelsFuncSigs = TrixelsMetaData.Sigs

// TrixelsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TrixelsMetaData.Bin instead.
var TrixelsBin = TrixelsMetaData.Bin

// DeployTrixels deploys a new Ethereum contract, binding an instance of Trixels to it.
func DeployTrixels(auth *bind.TransactOpts, backend bind.ContractBackend, newColors [][8]byte) (common.Address, *types.Transaction, *Trixels, error) {
	parsed, err := TrixelsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TrixelsBin), backend, newColors)
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

// ChangePixel is a paid mutator transaction binding the contract method 0xeab6402b.
//
// Solidity: function changePixel(uint256 x, uint256 y, bytes8 newColor) returns()
func (_Trixels *TrixelsTransactor) ChangePixel(opts *bind.TransactOpts, x *big.Int, y *big.Int, newColor [8]byte) (*types.Transaction, error) {
	return _Trixels.contract.Transact(opts, "changePixel", x, y, newColor)
}

// ChangePixel is a paid mutator transaction binding the contract method 0xeab6402b.
//
// Solidity: function changePixel(uint256 x, uint256 y, bytes8 newColor) returns()
func (_Trixels *TrixelsSession) ChangePixel(x *big.Int, y *big.Int, newColor [8]byte) (*types.Transaction, error) {
	return _Trixels.Contract.ChangePixel(&_Trixels.TransactOpts, x, y, newColor)
}

// ChangePixel is a paid mutator transaction binding the contract method 0xeab6402b.
//
// Solidity: function changePixel(uint256 x, uint256 y, bytes8 newColor) returns()
func (_Trixels *TrixelsTransactorSession) ChangePixel(x *big.Int, y *big.Int, newColor [8]byte) (*types.Transaction, error) {
	return _Trixels.Contract.ChangePixel(&_Trixels.TransactOpts, x, y, newColor)
}

// MassUpdatePixels is a paid mutator transaction binding the contract method 0x452955fd.
//
// Solidity: function massUpdatePixels((uint16,uint16,bytes8,address)[] newPixels) returns()
func (_Trixels *TrixelsTransactor) MassUpdatePixels(opts *bind.TransactOpts, newPixels []TrixelsPixel) (*types.Transaction, error) {
	return _Trixels.contract.Transact(opts, "massUpdatePixels", newPixels)
}

// MassUpdatePixels is a paid mutator transaction binding the contract method 0x452955fd.
//
// Solidity: function massUpdatePixels((uint16,uint16,bytes8,address)[] newPixels) returns()
func (_Trixels *TrixelsSession) MassUpdatePixels(newPixels []TrixelsPixel) (*types.Transaction, error) {
	return _Trixels.Contract.MassUpdatePixels(&_Trixels.TransactOpts, newPixels)
}

// MassUpdatePixels is a paid mutator transaction binding the contract method 0x452955fd.
//
// Solidity: function massUpdatePixels((uint16,uint16,bytes8,address)[] newPixels) returns()
func (_Trixels *TrixelsTransactorSession) MassUpdatePixels(newPixels []TrixelsPixel) (*types.Transaction, error) {
	return _Trixels.Contract.MassUpdatePixels(&_Trixels.TransactOpts, newPixels)
}
