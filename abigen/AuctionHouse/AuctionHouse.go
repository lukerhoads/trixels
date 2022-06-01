// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AuctionHouse

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

// AuctionHouseMetaData contains all meta data concerning the AuctionHouse contract.
var AuctionHouseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_distributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_minBidIncrementPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_daoCut\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"name\":\"AuctionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidPlaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"contractIDAO\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoCut\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"internalType\":\"contractIDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidIncrementPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startAuction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"}],\"name\":\"placeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"setDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"}],\"name\":\"setReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minBidIncrementPercentage\",\"type\":\"uint8\"}],\"name\":\"setMinIncrementPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620029ad380380620029ad833981810160405281019062000037919062000397565b83620000586200004c6200028660201b60201c565b6200028e60201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505086600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600b8190555081600d60006101000a81548160ff021916908360ff16021790555080600d60016101000a81548160ff021916908360ff1602179055506040518060c00160405280600081526020016000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160011515815250600560008201518160000155602082015181600101556040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a08201518160050160006101000a81548160ff02191690831515021790555090505050505050505050620004e8565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008151905062000363816200049a565b92915050565b6000815190506200037a81620004b4565b92915050565b6000815190506200039181620004ce565b92915050565b600080600080600080600060e0888a031215620003b957620003b862000495565b5b6000620003c98a828b0162000352565b9750506020620003dc8a828b0162000352565b9650506040620003ef8a828b0162000352565b9550506060620004028a828b0162000352565b9450506080620004158a828b0162000369565b93505060a0620004288a828b0162000380565b92505060c06200043b8a828b0162000380565b91505092959891949750929550565b600062000457826200045e565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600080fd5b620004a5816200044a565b8114620004b157600080fd5b50565b620004bf816200047e565b8114620004cb57600080fd5b50565b620004d98162000488565b8114620004e557600080fd5b50565b6124b580620004f86000396000f3fe6080604052600436106101095760003560e01c8063b296024d11610095578063ee33c37411610064578063ee33c3741461032f578063f2fde38b14610358578063f6be71d114610381578063fc0c546a146103aa578063fe67a54b146103d557610109565b8063b296024d14610285578063bfe10928146102b0578063ce9c7c0d146102db578063db2e1eed1461030457610109565b80636b64c769116100dc5780636b64c769146101cc578063715018a6146101f75780637d9f6db51461020e5780638da5cb5b1461023e5780639979ef451461026957610109565b80630fb5a6b41461010e578063150b7a02146101395780631ab3989a146101765780634162169f146101a1575b600080fd5b34801561011a57600080fd5b506101236103ec565b6040516101309190611e17565b60405180910390f35b34801561014557600080fd5b50610160600480360381019061015b91906117d9565b6103f2565b60405161016d9190611c2b565b60405180910390f35b34801561018257600080fd5b5061018b610407565b6040516101989190611eca565b60405180910390f35b3480156101ad57600080fd5b506101b661041a565b6040516101c39190611c46565b60405180910390f35b3480156101d857600080fd5b506101e1610440565b6040516101ee9190611e17565b60405180910390f35b34801561020357600080fd5b5061020c610795565b005b34801561021a57600080fd5b5061022361081d565b60405161023596959493929190611e69565b60405180910390f35b34801561024a57600080fd5b50610253610874565b6040516102609190611b87565b60405180910390f35b610283600480360381019061027e919061188e565b61089d565b005b34801561029157600080fd5b5061029a610c2b565b6040516102a79190611eca565b60405180910390f35b3480156102bc57600080fd5b506102c5610c3e565b6040516102d29190611c61565b60405180910390f35b3480156102e757600080fd5b5061030260048036038101906102fd919061188e565b610c64565b005b34801561031057600080fd5b50610319610cea565b6040516103269190611e17565b60405180910390f35b34801561033b57600080fd5b50610356600480360381019061035191906118e8565b610cf0565b005b34801561036457600080fd5b5061037f600480360381019061037a91906117ac565b610d8a565b005b34801561038d57600080fd5b506103a860048036038101906103a3919061188e565b610e82565b005b3480156103b657600080fd5b506103bf610f08565b6040516103cc9190611c7c565b60405180910390f35b3480156103e157600080fd5b506103ea610f2e565b005b600b5481565b600063150b7a0260e01b905095945050505050565b600d60019054906101000a900460ff1681565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061044a61145f565b73ffffffffffffffffffffffffffffffffffffffff16610468610874565b73ffffffffffffffffffffffffffffffffffffffff16146104be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b590611d77565b60405180910390fd5b600060056040518060c00160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820160009054906101000a900460ff16151515158152505090508060a001516105ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105a390611db7565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636a627842306040518263ffffffff1660e01b81526004016106099190611b87565b602060405180830381600087803b15801561062357600080fd5b505af1158015610637573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065b91906118bb565b905060004290506000600b54826106729190611f01565b90506040518060c0016040528083815260200182815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200184815260200160001515815250600560008201518160000155602082015181600101556040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a08201518160050160006101000a81548160ff0219169083151502179055509050507f44c53be110c6aa83aa83cd02e351ed172359268272ee1b5d31c0fe48db35c6c783838360405161078393929190611e32565b60405180910390a18294505050505090565b61079d61145f565b73ffffffffffffffffffffffffffffffffffffffff166107bb610874565b73ffffffffffffffffffffffffffffffffffffffff1614610811576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161080890611d77565b60405180910390fd5b61081b6000611467565b565b60058060000154908060010154908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050160009054906101000a900460ff16905086565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600060056040518060c00160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820160009054906101000a900460ff16151515158152505090508060800151821461098d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098490611cd7565b60405180910390fd5b600c5434116109d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c890611dd7565b60405180910390fd5b6064600d60009054906101000a900460ff1660ff1682604001516109f59190611f88565b6109ff9190611f57565b8160400151610a0e9190611f01565b341015610a50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4790611d37565b60405180910390fd5b806060015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610ac3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aba90611c97565b60405180910390fd5b80602001514210610b09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0090611df7565b60405180910390fd5b8060a0015115610b4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4590611d97565b60405180910390fd5b600081606001519050600073ffffffffffffffffffffffffffffffffffffffff16826060015173ffffffffffffffffffffffffffffffffffffffff1614610b9f57610b9e8160056002015461152b565b5b3460056002018190555033600560030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d3334604051610c1e929190611c02565b60405180910390a1505050565b600d60009054906101000a900460ff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610c6c61145f565b73ffffffffffffffffffffffffffffffffffffffff16610c8a610874565b73ffffffffffffffffffffffffffffffffffffffff1614610ce0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd790611d77565b60405180910390fd5b80600c8190555050565b600c5481565b610cf861145f565b73ffffffffffffffffffffffffffffffffffffffff16610d16610874565b73ffffffffffffffffffffffffffffffffffffffff1614610d6c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6390611d77565b60405180910390fd5b80600d60006101000a81548160ff021916908360ff16021790555050565b610d9261145f565b73ffffffffffffffffffffffffffffffffffffffff16610db0610874565b73ffffffffffffffffffffffffffffffffffffffff1614610e06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dfd90611d77565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610e76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6d90611cf7565b60405180910390fd5b610e7f81611467565b50565b610e8a61145f565b73ffffffffffffffffffffffffffffffffffffffff16610ea8610874565b73ffffffffffffffffffffffffffffffffffffffff1614610efe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ef590611d77565b60405180910390fd5b80600b8190555050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610f3661145f565b73ffffffffffffffffffffffffffffffffffffffff16610f54610874565b73ffffffffffffffffffffffffffffffffffffffff1614610faa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa190611d77565b60405180910390fd5b600060056040518060c00160405290816000820154815260200160018201548152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481526020016005820160009054906101000a900460ff161515151581525050905060008160000151141561109c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161109390611d57565b60405180910390fd5b8060a00151156110e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110d890611d97565b60405180910390fd5b80602001514211611127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111e90611cb7565b60405180910390fd5b60016005800160006101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff16600560030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561123457600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342966c6882608001516040518263ffffffff1660e01b81526004016111fd9190611e17565b600060405180830381600087803b15801561121757600080fd5b505af115801561122b573d6000803e3d6000fd5b505050506112ce565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd30836060015184608001516040518463ffffffff1660e01b815260040161129b93929190611bcb565b600060405180830381600087803b1580156112b557600080fd5b505af11580156112c9573d6000803e3d6000fd5b505050505b60008160400151111561141b5760006064600d60019054906101000a900460ff1660ff1683604001516113019190611f88565b61130b9190611f57565b9050600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636f27f69082846040015161135b9190611fe2565b84608001516040518363ffffffff1660e01b815260040161137c9190611e17565b6000604051808303818588803b15801561139557600080fd5b505af11580156113a9573d6000803e3d6000fd5b50505050506113da600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682611671565b611419576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141090611d17565b60405180910390fd5b505b7fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda81606001518260400151604051611454929190611ba2565b60405180910390a150565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6115358282611671565b61166d57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b1580156115a357600080fd5b505af11580156115b7573d6000803e3d6000fd5b5050505050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401611619929190611c02565b602060405180830381600087803b15801561163357600080fd5b505af1158015611647573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061166b9190611861565b505b5050565b6000808373ffffffffffffffffffffffffffffffffffffffff16836175309060405161169c90611b72565b600060405180830381858888f193505050503d80600081146116da576040519150601f19603f3d011682016040523d82523d6000602084013e6116df565b606091505b505090508091505092915050565b6000813590506116fc81612423565b92915050565b6000815190506117118161243a565b92915050565b60008083601f84011261172d5761172c6121ae565b5b8235905067ffffffffffffffff81111561174a576117496121a9565b5b602083019150836001820283011115611766576117656121b3565b5b9250929050565b60008135905061177c81612451565b92915050565b60008151905061179181612451565b92915050565b6000813590506117a681612468565b92915050565b6000602082840312156117c2576117c16121bd565b5b60006117d0848285016116ed565b91505092915050565b6000806000806000608086880312156117f5576117f46121bd565b5b6000611803888289016116ed565b9550506020611814888289016116ed565b94505060406118258882890161176d565b935050606086013567ffffffffffffffff811115611846576118456121b8565b5b61185288828901611717565b92509250509295509295909350565b600060208284031215611877576118766121bd565b5b600061188584828501611702565b91505092915050565b6000602082840312156118a4576118a36121bd565b5b60006118b28482850161176d565b91505092915050565b6000602082840312156118d1576118d06121bd565b5b60006118df84828501611782565b91505092915050565b6000602082840312156118fe576118fd6121bd565b5b600061190c84828501611797565b91505092915050565b61191e816120a9565b82525050565b61192d81612028565b82525050565b61193c81612016565b82525050565b61194b8161203a565b82525050565b61195a81612046565b82525050565b611969816120bb565b82525050565b611978816120df565b82525050565b61198781612103565b82525050565b600061199a601c83611ef0565b91506119a5826121c2565b602082019050919050565b60006119bd601b83611ef0565b91506119c8826121eb565b602082019050919050565b60006119e0601883611ef0565b91506119eb82612214565b602082019050919050565b6000611a03602683611ef0565b9150611a0e8261223d565b604082019050919050565b6000611a26602283611ef0565b9150611a318261228c565b604082019050919050565b6000611a49601783611ef0565b9150611a54826122db565b602082019050919050565b6000611a6c601683611ef0565b9150611a7782612304565b602082019050919050565b6000611a8f602083611ef0565b9150611a9a8261232d565b602082019050919050565b6000611ab2602083611ef0565b9150611abd82612356565b602082019050919050565b6000611ad5602883611ef0565b9150611ae08261237f565b604082019050919050565b6000611af8600083611ee5565b9150611b03826123ce565b600082019050919050565b6000611b1b601f83611ef0565b9150611b26826123d1565b602082019050919050565b6000611b3e601383611ef0565b9150611b49826123fa565b602082019050919050565b611b5d81612092565b82525050565b611b6c8161209c565b82525050565b6000611b7d82611aeb565b9150819050919050565b6000602082019050611b9c6000830184611933565b92915050565b6000604082019050611bb76000830185611915565b611bc46020830184611b54565b9392505050565b6000606082019050611be06000830186611933565b611bed6020830185611915565b611bfa6040830184611b54565b949350505050565b6000604082019050611c176000830185611933565b611c246020830184611b54565b9392505050565b6000602082019050611c406000830184611951565b92915050565b6000602082019050611c5b6000830184611960565b92915050565b6000602082019050611c76600083018461196f565b92915050565b6000602082019050611c91600083018461197e565b92915050565b60006020820190508181036000830152611cb08161198d565b9050919050565b60006020820190508181036000830152611cd0816119b0565b9050919050565b60006020820190508181036000830152611cf0816119d3565b9050919050565b60006020820190508181036000830152611d10816119f6565b9050919050565b60006020820190508181036000830152611d3081611a19565b9050919050565b60006020820190508181036000830152611d5081611a3c565b9050919050565b60006020820190508181036000830152611d7081611a5f565b9050919050565b60006020820190508181036000830152611d9081611a82565b9050919050565b60006020820190508181036000830152611db081611aa5565b9050919050565b60006020820190508181036000830152611dd081611ac8565b9050919050565b60006020820190508181036000830152611df081611b0e565b9050919050565b60006020820190508181036000830152611e1081611b31565b9050919050565b6000602082019050611e2c6000830184611b54565b92915050565b6000606082019050611e476000830186611b54565b611e546020830185611b54565b611e616040830184611b54565b949350505050565b600060c082019050611e7e6000830189611b54565b611e8b6020830188611b54565b611e986040830187611b54565b611ea56060830186611924565b611eb26080830185611b54565b611ebf60a0830184611942565b979650505050505050565b6000602082019050611edf6000830184611b63565b92915050565b600081905092915050565b600082825260208201905092915050565b6000611f0c82612092565b9150611f1783612092565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611f4c57611f4b61214b565b5b828201905092915050565b6000611f6282612092565b9150611f6d83612092565b925082611f7d57611f7c61217a565b5b828204905092915050565b6000611f9382612092565b9150611f9e83612092565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611fd757611fd661214b565b5b828202905092915050565b6000611fed82612092565b9150611ff883612092565b92508282101561200b5761200a61214b565b5b828203905092915050565b600061202182612072565b9050919050565b600061203382612072565b9050919050565b60008115159050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60006120b482612127565b9050919050565b60006120c6826120cd565b9050919050565b60006120d882612072565b9050919050565b60006120ea826120f1565b9050919050565b60006120fc82612072565b9050919050565b600061210e82612115565b9050919050565b600061212082612072565b9050919050565b600061213282612139565b9050919050565b600061214482612072565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b7f416c726561647920686176652074686520686967686573742062696400000000600082015250565b7f41756374696f6e2063616e6e6f7420626520656e646564207965740000000000600082015250565b7f546f6b656e206e6f7420757020666f722061756374696f6e0000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f436f756c64206e6f74207472616e7366657220746f2044414f20636f6e74726160008201527f6374000000000000000000000000000000000000000000000000000000000000602082015250565b7f42696420616d6f756e7420696e73756666696369656e74000000000000000000600082015250565b7f41756374696f6e206861736e2774207374617274656400000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f41756374696f6e2068617320616c7265616479206265656e20736574746c6564600082015250565b7f416e6f746865722061756374696f6e2068617320616c7265616479206265656e60008201527f2073746172746564000000000000000000000000000000000000000000000000602082015250565b50565b7f42696420646f6573206e6f74206d656574207265736572766520707269636500600082015250565b7f41756374696f6e20686173206578706972656400000000000000000000000000600082015250565b61242c81612016565b811461243757600080fd5b50565b6124438161203a565b811461244e57600080fd5b50565b61245a81612092565b811461246557600080fd5b50565b6124718161209c565b811461247c57600080fd5b5056fea2646970667358221220c67955979ce11a63117bebb674a62863ea7900b2c548108a0752ee6c559379d564736f6c63430008060033",
}

// AuctionHouseABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionHouseMetaData.ABI instead.
var AuctionHouseABI = AuctionHouseMetaData.ABI

// AuctionHouseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionHouseMetaData.Bin instead.
var AuctionHouseBin = AuctionHouseMetaData.Bin

// DeployAuctionHouse deploys a new Ethereum contract, binding an instance of AuctionHouse to it.
func DeployAuctionHouse(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _dao common.Address, _distributor common.Address, _weth common.Address, _duration *big.Int, _minBidIncrementPercentage uint8, _daoCut uint8) (common.Address, *types.Transaction, *AuctionHouse, error) {
	parsed, err := AuctionHouseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionHouseBin), backend, _token, _dao, _distributor, _weth, _duration, _minBidIncrementPercentage, _daoCut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuctionHouse{AuctionHouseCaller: AuctionHouseCaller{contract: contract}, AuctionHouseTransactor: AuctionHouseTransactor{contract: contract}, AuctionHouseFilterer: AuctionHouseFilterer{contract: contract}}, nil
}

// AuctionHouse is an auto generated Go binding around an Ethereum contract.
type AuctionHouse struct {
	AuctionHouseCaller     // Read-only binding to the contract
	AuctionHouseTransactor // Write-only binding to the contract
	AuctionHouseFilterer   // Log filterer for contract events
}

// AuctionHouseCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionHouseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionHouseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionHouseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionHouseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionHouseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionHouseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionHouseSession struct {
	Contract     *AuctionHouse     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionHouseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionHouseCallerSession struct {
	Contract *AuctionHouseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AuctionHouseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionHouseTransactorSession struct {
	Contract     *AuctionHouseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AuctionHouseRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionHouseRaw struct {
	Contract *AuctionHouse // Generic contract binding to access the raw methods on
}

// AuctionHouseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionHouseCallerRaw struct {
	Contract *AuctionHouseCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionHouseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionHouseTransactorRaw struct {
	Contract *AuctionHouseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionHouse creates a new instance of AuctionHouse, bound to a specific deployed contract.
func NewAuctionHouse(address common.Address, backend bind.ContractBackend) (*AuctionHouse, error) {
	contract, err := bindAuctionHouse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionHouse{AuctionHouseCaller: AuctionHouseCaller{contract: contract}, AuctionHouseTransactor: AuctionHouseTransactor{contract: contract}, AuctionHouseFilterer: AuctionHouseFilterer{contract: contract}}, nil
}

// NewAuctionHouseCaller creates a new read-only instance of AuctionHouse, bound to a specific deployed contract.
func NewAuctionHouseCaller(address common.Address, caller bind.ContractCaller) (*AuctionHouseCaller, error) {
	contract, err := bindAuctionHouse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionHouseCaller{contract: contract}, nil
}

// NewAuctionHouseTransactor creates a new write-only instance of AuctionHouse, bound to a specific deployed contract.
func NewAuctionHouseTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionHouseTransactor, error) {
	contract, err := bindAuctionHouse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionHouseTransactor{contract: contract}, nil
}

// NewAuctionHouseFilterer creates a new log filterer instance of AuctionHouse, bound to a specific deployed contract.
func NewAuctionHouseFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionHouseFilterer, error) {
	contract, err := bindAuctionHouse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionHouseFilterer{contract: contract}, nil
}

// bindAuctionHouse binds a generic wrapper to an already deployed contract.
func bindAuctionHouse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionHouseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionHouse *AuctionHouseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionHouse.Contract.AuctionHouseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionHouse *AuctionHouseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHouse.Contract.AuctionHouseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionHouse *AuctionHouseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionHouse.Contract.AuctionHouseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionHouse *AuctionHouseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionHouse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionHouse *AuctionHouseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHouse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionHouse *AuctionHouseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionHouse.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 startDate, uint256 endDate, uint256 highestBid, address highestBidder, uint256 tokenId, bool settled)
func (_AuctionHouse *AuctionHouseCaller) Auction(opts *bind.CallOpts) (struct {
	StartDate     *big.Int
	EndDate       *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenId       *big.Int
	Settled       bool
}, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "auction")

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
func (_AuctionHouse *AuctionHouseSession) Auction() (struct {
	StartDate     *big.Int
	EndDate       *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenId       *big.Int
	Settled       bool
}, error) {
	return _AuctionHouse.Contract.Auction(&_AuctionHouse.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 startDate, uint256 endDate, uint256 highestBid, address highestBidder, uint256 tokenId, bool settled)
func (_AuctionHouse *AuctionHouseCallerSession) Auction() (struct {
	StartDate     *big.Int
	EndDate       *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	TokenId       *big.Int
	Settled       bool
}, error) {
	return _AuctionHouse.Contract.Auction(&_AuctionHouse.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_AuctionHouse *AuctionHouseCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_AuctionHouse *AuctionHouseSession) Dao() (common.Address, error) {
	return _AuctionHouse.Contract.Dao(&_AuctionHouse.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_AuctionHouse *AuctionHouseCallerSession) Dao() (common.Address, error) {
	return _AuctionHouse.Contract.Dao(&_AuctionHouse.CallOpts)
}

// DaoCut is a free data retrieval call binding the contract method 0x1ab3989a.
//
// Solidity: function daoCut() view returns(uint8)
func (_AuctionHouse *AuctionHouseCaller) DaoCut(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "daoCut")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DaoCut is a free data retrieval call binding the contract method 0x1ab3989a.
//
// Solidity: function daoCut() view returns(uint8)
func (_AuctionHouse *AuctionHouseSession) DaoCut() (uint8, error) {
	return _AuctionHouse.Contract.DaoCut(&_AuctionHouse.CallOpts)
}

// DaoCut is a free data retrieval call binding the contract method 0x1ab3989a.
//
// Solidity: function daoCut() view returns(uint8)
func (_AuctionHouse *AuctionHouseCallerSession) DaoCut() (uint8, error) {
	return _AuctionHouse.Contract.DaoCut(&_AuctionHouse.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_AuctionHouse *AuctionHouseCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_AuctionHouse *AuctionHouseSession) Distributor() (common.Address, error) {
	return _AuctionHouse.Contract.Distributor(&_AuctionHouse.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_AuctionHouse *AuctionHouseCallerSession) Distributor() (common.Address, error) {
	return _AuctionHouse.Contract.Distributor(&_AuctionHouse.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_AuctionHouse *AuctionHouseCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_AuctionHouse *AuctionHouseSession) Duration() (*big.Int, error) {
	return _AuctionHouse.Contract.Duration(&_AuctionHouse.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_AuctionHouse *AuctionHouseCallerSession) Duration() (*big.Int, error) {
	return _AuctionHouse.Contract.Duration(&_AuctionHouse.CallOpts)
}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_AuctionHouse *AuctionHouseCaller) MinBidIncrementPercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "minBidIncrementPercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_AuctionHouse *AuctionHouseSession) MinBidIncrementPercentage() (uint8, error) {
	return _AuctionHouse.Contract.MinBidIncrementPercentage(&_AuctionHouse.CallOpts)
}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_AuctionHouse *AuctionHouseCallerSession) MinBidIncrementPercentage() (uint8, error) {
	return _AuctionHouse.Contract.MinBidIncrementPercentage(&_AuctionHouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionHouse *AuctionHouseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionHouse *AuctionHouseSession) Owner() (common.Address, error) {
	return _AuctionHouse.Contract.Owner(&_AuctionHouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionHouse *AuctionHouseCallerSession) Owner() (common.Address, error) {
	return _AuctionHouse.Contract.Owner(&_AuctionHouse.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_AuctionHouse *AuctionHouseCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_AuctionHouse *AuctionHouseSession) ReservePrice() (*big.Int, error) {
	return _AuctionHouse.Contract.ReservePrice(&_AuctionHouse.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_AuctionHouse *AuctionHouseCallerSession) ReservePrice() (*big.Int, error) {
	return _AuctionHouse.Contract.ReservePrice(&_AuctionHouse.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AuctionHouse *AuctionHouseCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionHouse.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AuctionHouse *AuctionHouseSession) Token() (common.Address, error) {
	return _AuctionHouse.Contract.Token(&_AuctionHouse.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_AuctionHouse *AuctionHouseCallerSession) Token() (common.Address, error) {
	return _AuctionHouse.Contract.Token(&_AuctionHouse.CallOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_AuctionHouse *AuctionHouseTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_AuctionHouse *AuctionHouseSession) EndAuction() (*types.Transaction, error) {
	return _AuctionHouse.Contract.EndAuction(&_AuctionHouse.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_AuctionHouse *AuctionHouseTransactorSession) EndAuction() (*types.Transaction, error) {
	return _AuctionHouse.Contract.EndAuction(&_AuctionHouse.TransactOpts)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_AuctionHouse *AuctionHouseTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_AuctionHouse *AuctionHouseSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AuctionHouse.Contract.OnERC721Received(&_AuctionHouse.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_AuctionHouse *AuctionHouseTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _AuctionHouse.Contract.OnERC721Received(&_AuctionHouse.TransactOpts, operator, from, tokenId, data)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenID) payable returns()
func (_AuctionHouse *AuctionHouseTransactor) PlaceBid(opts *bind.TransactOpts, _tokenID *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "placeBid", _tokenID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenID) payable returns()
func (_AuctionHouse *AuctionHouseSession) PlaceBid(_tokenID *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.Contract.PlaceBid(&_AuctionHouse.TransactOpts, _tokenID)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 _tokenID) payable returns()
func (_AuctionHouse *AuctionHouseTransactorSession) PlaceBid(_tokenID *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.Contract.PlaceBid(&_AuctionHouse.TransactOpts, _tokenID)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionHouse *AuctionHouseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionHouse *AuctionHouseSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionHouse.Contract.RenounceOwnership(&_AuctionHouse.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionHouse *AuctionHouseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionHouse.Contract.RenounceOwnership(&_AuctionHouse.TransactOpts)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_AuctionHouse *AuctionHouseTransactor) SetDuration(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "setDuration", _duration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_AuctionHouse *AuctionHouseSession) SetDuration(_duration *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.Contract.SetDuration(&_AuctionHouse.TransactOpts, _duration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _duration) returns()
func (_AuctionHouse *AuctionHouseTransactorSession) SetDuration(_duration *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.Contract.SetDuration(&_AuctionHouse.TransactOpts, _duration)
}

// SetMinIncrementPercentage is a paid mutator transaction binding the contract method 0xee33c374.
//
// Solidity: function setMinIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_AuctionHouse *AuctionHouseTransactor) SetMinIncrementPercentage(opts *bind.TransactOpts, _minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "setMinIncrementPercentage", _minBidIncrementPercentage)
}

// SetMinIncrementPercentage is a paid mutator transaction binding the contract method 0xee33c374.
//
// Solidity: function setMinIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_AuctionHouse *AuctionHouseSession) SetMinIncrementPercentage(_minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _AuctionHouse.Contract.SetMinIncrementPercentage(&_AuctionHouse.TransactOpts, _minBidIncrementPercentage)
}

// SetMinIncrementPercentage is a paid mutator transaction binding the contract method 0xee33c374.
//
// Solidity: function setMinIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_AuctionHouse *AuctionHouseTransactorSession) SetMinIncrementPercentage(_minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _AuctionHouse.Contract.SetMinIncrementPercentage(&_AuctionHouse.TransactOpts, _minBidIncrementPercentage)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_AuctionHouse *AuctionHouseTransactor) SetReservePrice(opts *bind.TransactOpts, _reservePrice *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "setReservePrice", _reservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_AuctionHouse *AuctionHouseSession) SetReservePrice(_reservePrice *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.Contract.SetReservePrice(&_AuctionHouse.TransactOpts, _reservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_AuctionHouse *AuctionHouseTransactorSession) SetReservePrice(_reservePrice *big.Int) (*types.Transaction, error) {
	return _AuctionHouse.Contract.SetReservePrice(&_AuctionHouse.TransactOpts, _reservePrice)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_AuctionHouse *AuctionHouseTransactor) StartAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "startAuction")
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_AuctionHouse *AuctionHouseSession) StartAuction() (*types.Transaction, error) {
	return _AuctionHouse.Contract.StartAuction(&_AuctionHouse.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x6b64c769.
//
// Solidity: function startAuction() returns(uint256)
func (_AuctionHouse *AuctionHouseTransactorSession) StartAuction() (*types.Transaction, error) {
	return _AuctionHouse.Contract.StartAuction(&_AuctionHouse.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionHouse *AuctionHouseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionHouse.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionHouse *AuctionHouseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionHouse.Contract.TransferOwnership(&_AuctionHouse.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionHouse *AuctionHouseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionHouse.Contract.TransferOwnership(&_AuctionHouse.TransactOpts, newOwner)
}

// AuctionHouseAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the AuctionHouse contract.
type AuctionHouseAuctionEndedIterator struct {
	Event *AuctionHouseAuctionEnded // Event containing the contract specifics and raw log

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
func (it *AuctionHouseAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHouseAuctionEnded)
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
		it.Event = new(AuctionHouseAuctionEnded)
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
func (it *AuctionHouseAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHouseAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHouseAuctionEnded represents a AuctionEnded event raised by the AuctionHouse contract.
type AuctionHouseAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_AuctionHouse *AuctionHouseFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*AuctionHouseAuctionEndedIterator, error) {

	logs, sub, err := _AuctionHouse.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &AuctionHouseAuctionEndedIterator{contract: _AuctionHouse.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_AuctionHouse *AuctionHouseFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *AuctionHouseAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _AuctionHouse.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHouseAuctionEnded)
				if err := _AuctionHouse.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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
func (_AuctionHouse *AuctionHouseFilterer) ParseAuctionEnded(log types.Log) (*AuctionHouseAuctionEnded, error) {
	event := new(AuctionHouseAuctionEnded)
	if err := _AuctionHouse.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHouseAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the AuctionHouse contract.
type AuctionHouseAuctionStartedIterator struct {
	Event *AuctionHouseAuctionStarted // Event containing the contract specifics and raw log

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
func (it *AuctionHouseAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHouseAuctionStarted)
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
		it.Event = new(AuctionHouseAuctionStarted)
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
func (it *AuctionHouseAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHouseAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHouseAuctionStarted represents a AuctionStarted event raised by the AuctionHouse contract.
type AuctionHouseAuctionStarted struct {
	TokenId   *big.Int
	StartDate *big.Int
	EndDate   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0x44c53be110c6aa83aa83cd02e351ed172359268272ee1b5d31c0fe48db35c6c7.
//
// Solidity: event AuctionStarted(uint256 tokenId, uint256 startDate, uint256 endDate)
func (_AuctionHouse *AuctionHouseFilterer) FilterAuctionStarted(opts *bind.FilterOpts) (*AuctionHouseAuctionStartedIterator, error) {

	logs, sub, err := _AuctionHouse.contract.FilterLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return &AuctionHouseAuctionStartedIterator{contract: _AuctionHouse.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0x44c53be110c6aa83aa83cd02e351ed172359268272ee1b5d31c0fe48db35c6c7.
//
// Solidity: event AuctionStarted(uint256 tokenId, uint256 startDate, uint256 endDate)
func (_AuctionHouse *AuctionHouseFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *AuctionHouseAuctionStarted) (event.Subscription, error) {

	logs, sub, err := _AuctionHouse.contract.WatchLogs(opts, "AuctionStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHouseAuctionStarted)
				if err := _AuctionHouse.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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
func (_AuctionHouse *AuctionHouseFilterer) ParseAuctionStarted(log types.Log) (*AuctionHouseAuctionStarted, error) {
	event := new(AuctionHouseAuctionStarted)
	if err := _AuctionHouse.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHouseBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the AuctionHouse contract.
type AuctionHouseBidPlacedIterator struct {
	Event *AuctionHouseBidPlaced // Event containing the contract specifics and raw log

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
func (it *AuctionHouseBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHouseBidPlaced)
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
		it.Event = new(AuctionHouseBidPlaced)
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
func (it *AuctionHouseBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHouseBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHouseBidPlaced represents a BidPlaced event raised by the AuctionHouse contract.
type AuctionHouseBidPlaced struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address bidder, uint256 amount)
func (_AuctionHouse *AuctionHouseFilterer) FilterBidPlaced(opts *bind.FilterOpts) (*AuctionHouseBidPlacedIterator, error) {

	logs, sub, err := _AuctionHouse.contract.FilterLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return &AuctionHouseBidPlacedIterator{contract: _AuctionHouse.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0x3fabff0a9c3ecd6814702e247fa9733e5d0aa69e3a38590f92cb18f623a2254d.
//
// Solidity: event BidPlaced(address bidder, uint256 amount)
func (_AuctionHouse *AuctionHouseFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *AuctionHouseBidPlaced) (event.Subscription, error) {

	logs, sub, err := _AuctionHouse.contract.WatchLogs(opts, "BidPlaced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHouseBidPlaced)
				if err := _AuctionHouse.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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
func (_AuctionHouse *AuctionHouseFilterer) ParseBidPlaced(log types.Log) (*AuctionHouseBidPlaced, error) {
	event := new(AuctionHouseBidPlaced)
	if err := _AuctionHouse.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionHouseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuctionHouse contract.
type AuctionHouseOwnershipTransferredIterator struct {
	Event *AuctionHouseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionHouseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionHouseOwnershipTransferred)
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
		it.Event = new(AuctionHouseOwnershipTransferred)
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
func (it *AuctionHouseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionHouseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionHouseOwnershipTransferred represents a OwnershipTransferred event raised by the AuctionHouse contract.
type AuctionHouseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionHouse *AuctionHouseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionHouseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionHouse.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionHouseOwnershipTransferredIterator{contract: _AuctionHouse.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionHouse *AuctionHouseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionHouseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionHouse.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionHouseOwnershipTransferred)
				if err := _AuctionHouse.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuctionHouse *AuctionHouseFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionHouseOwnershipTransferred, error) {
	event := new(AuctionHouseOwnershipTransferred)
	if err := _AuctionHouse.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
