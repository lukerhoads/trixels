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

// TrixelsAuctionHouseMetaData contains all meta data concerning the TrixelsAuctionHouse contract.
var TrixelsAuctionHouseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"trixelId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"extended\",\"type\":\"bool\"}],\"name\":\"AuctionBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"trixelId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"}],\"name\":\"AuctionMinBidIncrementPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"AuctionReservePriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"trixelId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"}],\"name\":\"AuctionTimeBufferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidIncrementPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeBuffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTrixelsToken\",\"name\":\"_trixels\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes23\",\"name\":\"skyNetID\",\"type\":\"bytes23\"}],\"name\":\"startAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"trixelId\",\"type\":\"uint256\"}],\"name\":\"createBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timeBuffer\",\"type\":\"uint256\"}],\"name\":\"setTimeBuffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"}],\"name\":\"setReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612944806100206000396000f3fe6080604052600436106100fe5760003560e01c8063715018a611610095578063b296024d11610064578063b296024d146102a2578063ce9c7c0d146102cd578063db2e1eed146102f6578063ec91f2a414610321578063f2fde38b1461034c576100fe565b8063715018a6146102325780638456cb59146102495780638da5cb5b14610260578063a4d0a17e1461028b576100fe565b80635c975abb116100d15780635c975abb14610199578063659dd2b4146101c45780636c28e349146101e05780637120334b14610209576100fe565b80630fb5a6b4146101035780632b107ca21461012e5780633f4ba83a146101575780633fc8cef31461016e575b600080fd5b34801561010f57600080fd5b50610118610375565b6040516101259190611b0e565b60405180910390f35b34801561013a57600080fd5b5061015560048036038101906101509190611b86565b61037b565b005b34801561016357600080fd5b5061016c610425565b005b34801561017a57600080fd5b506101836104ab565b6040516101909190611bf4565b60405180910390f35b3480156101a557600080fd5b506101ae6104d1565b6040516101bb9190611c2a565b60405180910390f35b6101de60048036038101906101d99190611c71565b6104e8565b005b3480156101ec57600080fd5b5061020760048036038101906102029190611d08565b61089d565b005b34801561021557600080fd5b50610230600480360381019061022b9190611c71565b6109e5565b005b34801561023e57600080fd5b50610247610aa2565b005b34801561025557600080fd5b5061025e610b2a565b005b34801561026c57600080fd5b50610275610bb0565b6040516102829190611bf4565b60405180910390f35b34801561029757600080fd5b506102a0610bda565b005b3480156102ae57600080fd5b506102b7610c82565b6040516102c49190611d9f565b60405180910390f35b3480156102d957600080fd5b506102f460048036038101906102ef9190611c71565b610c95565b005b34801561030257600080fd5b5061030b610d52565b6040516103189190611b0e565b60405180910390f35b34801561032d57600080fd5b50610336610d58565b6040516103439190611b0e565b60405180910390f35b34801561035857600080fd5b50610373600480360381019061036e9190611dba565b610d5e565b005b60c95481565b600260655414156103c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b890611e44565b60405180910390fd5b60026065819055506103d16104d1565b15610411576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040890611eb0565b60405180910390fd5b61041a81610e56565b600160658190555050565b61042d61102b565b73ffffffffffffffffffffffffffffffffffffffff1661044b610bb0565b73ffffffffffffffffffffffffffffffffffffffff16146104a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049890611f1c565b60405180910390fd5b6104a9611033565b565b60cd60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000603360009054906101000a900460ff16905090565b6002606554141561052e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052590611e44565b60405180910390fd5b6002606581905550600060ce6040518060c0016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460ff161515151581525050905081816000015114610626576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061d90611f88565b60405180910390fd5b8060600151421061066c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066390611ff4565b60405180910390fd5b60cc543410156106b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a890612060565b60405180910390fd5b606460cb60009054906101000a900460ff1660ff1682602001516106d591906120af565b6106df9190612138565b81602001516106ee9190612169565b341015610730576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072790612231565b60405180910390fd5b600081608001519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461077c5761077b8183602001516110d5565b5b3460ce600101819055503360ce60040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600060ca544284606001516107df9190612251565b10905080156108085760ca54426107f69190612169565b836060018181525060ce600301819055505b82600001517f1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea333348460405161084093929190612285565b60405180910390a2801561088f5782600001517f6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e84606001516040516108869190611b0e565b60405180910390a25b505050600160658190555050565b60006108a9600161120c565b905080156108cd576001600060016101000a81548160ff0219169083151502179055505b6108d56112fc565b6108dd611355565b6108e56113ae565b6108ed611407565b8160cd60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508560d360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360ca819055508460c9819055508260cc8190555080156109dd5760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516109d49190612301565b60405180910390a15b505050505050565b6109ed61102b565b73ffffffffffffffffffffffffffffffffffffffff16610a0b610bb0565b73ffffffffffffffffffffffffffffffffffffffff1614610a61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5890611f1c565b60405180910390fd5b8060ca819055507f1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d81604051610a979190611b0e565b60405180910390a150565b610aaa61102b565b73ffffffffffffffffffffffffffffffffffffffff16610ac8610bb0565b73ffffffffffffffffffffffffffffffffffffffff1614610b1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1590611f1c565b60405180910390fd5b610b2860006114aa565b565b610b3261102b565b73ffffffffffffffffffffffffffffffffffffffff16610b50610bb0565b73ffffffffffffffffffffffffffffffffffffffff1614610ba6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9d90611f1c565b60405180910390fd5b610bae611407565b565b6000609760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60026065541415610c20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1790611e44565b60405180910390fd5b6002606581905550610c306104d1565b15610c70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6790611eb0565b60405180910390fd5b610c78611570565b6001606581905550565b60cb60009054906101000a900460ff1681565b610c9d61102b565b73ffffffffffffffffffffffffffffffffffffffff16610cbb610bb0565b73ffffffffffffffffffffffffffffffffffffffff1614610d11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0890611f1c565b60405180910390fd5b8060cc819055507f6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b281604051610d479190611b0e565b60405180910390a150565b60cc5481565b60ca5481565b610d6661102b565b73ffffffffffffffffffffffffffffffffffffffff16610d84610bb0565b73ffffffffffffffffffffffffffffffffffffffff1614610dda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd190611f1c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610e4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e419061238e565b60405180910390fd5b610e53816114aa565b50565b600060d360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b487a841836040518263ffffffff1660e01b8152600401610eb391906123bd565b6020604051808303816000875af1158015610ed2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef691906123ed565b90506000429050600060c95482610f0d9190612169565b90506040518060c0016040528084815260200160008152602001838152602001828152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000151581525060ce6000820151816000015560208201518160010155604082015181600201556060820151816003015560808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160040160146101000a81548160ff021916908315150217905550905050827fd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca838360405161101d92919061241a565b60405180910390a250505050565b600033905090565b61103b6104d1565b61107a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110719061248f565b60405180910390fd5b6000603360006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6110be61102b565b6040516110cb9190611bf4565b60405180910390a1565b6110df82826118e1565b6112085760cd60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b15801561114d57600080fd5b505af1158015611161573d6000803e3d6000fd5b505050505060cd60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b81526004016111c39291906124af565b6020604051808303816000875af11580156111e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112069190612504565b505b5050565b60008060019054906101000a900460ff16156112835760018260ff1614801561123b5750611239306119ac565b155b61127a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611271906125a3565b60405180910390fd5b600090506112f7565b8160ff1660008054906101000a900460ff1660ff16106112d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112cf906125a3565b60405180910390fd5b816000806101000a81548160ff021916908360ff160217905550600190505b919050565b600060019054906101000a900460ff1661134b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161134290612635565b60405180910390fd5b6113536119cf565b565b600060019054906101000a900460ff166113a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161139b90612635565b60405180910390fd5b6113ac611a3b565b565b600060019054906101000a900460ff166113fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113f490612635565b60405180910390fd5b611405611a94565b565b61140f6104d1565b1561144f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161144690611eb0565b60405180910390fd5b6001603360006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861149361102b565b6040516114a09190611bf4565b60405180910390a1565b6000609760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081609760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600060ce6040518060c0016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160149054906101000a900460ff1615151515815250509050600081604001511415611662576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611659906126a1565b60405180910390fd5b8060a00151156116a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161169e9061270d565b60405180910390fd5b80606001514210156116ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e590612779565b60405180910390fd5b600160ce60040160146101000a81548160ff021916908315150217905550600073ffffffffffffffffffffffffffffffffffffffff16816080015173ffffffffffffffffffffffffffffffffffffffff1614156117db5760d360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166342966c6882600001516040518263ffffffff1660e01b81526004016117a49190611b0e565b600060405180830381600087803b1580156117be57600080fd5b505af11580156117d2573d6000803e3d6000fd5b50505050611875565b60d360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd30836080015184600001516040518463ffffffff1660e01b8152600401611842939291906127ee565b600060405180830381600087803b15801561185c57600080fd5b505af1158015611870573d6000803e3d6000fd5b505050505b6000816020015111156118985761189761188d610bb0565b82602001516110d5565b5b80600001517fc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99826080015183602001516040516118d6929190612825565b60405180910390a250565b6000808373ffffffffffffffffffffffffffffffffffffffff168361753090600067ffffffffffffffff81111561191b5761191a61284e565b5b6040519080825280601f01601f19166020018201604052801561194d5781602001600182028036833780820191505090505b5060405161195b91906128f7565b600060405180830381858888f193505050503d8060008114611999576040519150601f19603f3d011682016040523d82523d6000602084013e61199e565b606091505b505090508091505092915050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff16611a1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1590612635565b60405180910390fd5b6000603360006101000a81548160ff021916908315150217905550565b600060019054906101000a900460ff16611a8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a8190612635565b60405180910390fd5b6001606581905550565b600060019054906101000a900460ff16611ae3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ada90612635565b60405180910390fd5b611af3611aee61102b565b6114aa565b565b6000819050919050565b611b0881611af5565b82525050565b6000602082019050611b236000830184611aff565b92915050565b600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffff00000000000000000082169050919050565b611b6381611b2e565b8114611b6e57600080fd5b50565b600081359050611b8081611b5a565b92915050565b600060208284031215611b9c57611b9b611b29565b5b6000611baa84828501611b71565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611bde82611bb3565b9050919050565b611bee81611bd3565b82525050565b6000602082019050611c096000830184611be5565b92915050565b60008115159050919050565b611c2481611c0f565b82525050565b6000602082019050611c3f6000830184611c1b565b92915050565b611c4e81611af5565b8114611c5957600080fd5b50565b600081359050611c6b81611c45565b92915050565b600060208284031215611c8757611c86611b29565b5b6000611c9584828501611c5c565b91505092915050565b6000611ca982611bd3565b9050919050565b611cb981611c9e565b8114611cc457600080fd5b50565b600081359050611cd681611cb0565b92915050565b611ce581611bd3565b8114611cf057600080fd5b50565b600081359050611d0281611cdc565b92915050565b600080600080600060a08688031215611d2457611d23611b29565b5b6000611d3288828901611cc7565b9550506020611d4388828901611c5c565b9450506040611d5488828901611c5c565b9350506060611d6588828901611c5c565b9250506080611d7688828901611cf3565b9150509295509295909350565b600060ff82169050919050565b611d9981611d83565b82525050565b6000602082019050611db46000830184611d90565b92915050565b600060208284031215611dd057611dcf611b29565b5b6000611dde84828501611cf3565b91505092915050565b600082825260208201905092915050565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6000611e2e601f83611de7565b9150611e3982611df8565b602082019050919050565b60006020820190508181036000830152611e5d81611e21565b9050919050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b6000611e9a601083611de7565b9150611ea582611e64565b602082019050919050565b60006020820190508181036000830152611ec981611e8d565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611f06602083611de7565b9150611f1182611ed0565b602082019050919050565b60006020820190508181036000830152611f3581611ef9565b9050919050565b7f54726978656c206e6f7420757020666f722061756374696f6e00000000000000600082015250565b6000611f72601983611de7565b9150611f7d82611f3c565b602082019050919050565b60006020820190508181036000830152611fa181611f65565b9050919050565b7f41756374696f6e20657870697265640000000000000000000000000000000000600082015250565b6000611fde600f83611de7565b9150611fe982611fa8565b602082019050919050565b6000602082019050818103600083015261200d81611fd1565b9050919050565b7f4d7573742073656e64206174206c656173742072657365727665507269636500600082015250565b600061204a601f83611de7565b915061205582612014565b602082019050919050565b600060208201905081810360008301526120798161203d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006120ba82611af5565b91506120c583611af5565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156120fe576120fd612080565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061214382611af5565b915061214e83611af5565b92508261215e5761215d612109565b5b828204905092915050565b600061217482611af5565b915061217f83611af5565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156121b4576121b3612080565b5b828201905092915050565b7f4d7573742073656e64206d6f7265207468616e206c617374206269642062792060008201527f6d696e426964496e6372656d656e7450657263656e7461676520616d6f756e74602082015250565b600061221b604083611de7565b9150612226826121bf565b604082019050919050565b6000602082019050818103600083015261224a8161220e565b9050919050565b600061225c82611af5565b915061226783611af5565b92508282101561227a57612279612080565b5b828203905092915050565b600060608201905061229a6000830186611be5565b6122a76020830185611aff565b6122b46040830184611c1b565b949350505050565b6000819050919050565b6000819050919050565b60006122eb6122e66122e1846122bc565b6122c6565b611d83565b9050919050565b6122fb816122d0565b82525050565b600060208201905061231660008301846122f2565b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000612378602683611de7565b91506123838261231c565b604082019050919050565b600060208201905081810360008301526123a78161236b565b9050919050565b6123b781611b2e565b82525050565b60006020820190506123d260008301846123ae565b92915050565b6000815190506123e781611c45565b92915050565b60006020828403121561240357612402611b29565b5b6000612411848285016123d8565b91505092915050565b600060408201905061242f6000830185611aff565b61243c6020830184611aff565b9392505050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b6000612479601483611de7565b915061248482612443565b602082019050919050565b600060208201905081810360008301526124a88161246c565b9050919050565b60006040820190506124c46000830185611be5565b6124d16020830184611aff565b9392505050565b6124e181611c0f565b81146124ec57600080fd5b50565b6000815190506124fe816124d8565b92915050565b60006020828403121561251a57612519611b29565b5b6000612528848285016124ef565b91505092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600061258d602e83611de7565b915061259882612531565b604082019050919050565b600060208201905081810360008301526125bc81612580565b9050919050565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b600061261f602b83611de7565b915061262a826125c3565b604082019050919050565b6000602082019050818103600083015261264e81612612565b9050919050565b7f41756374696f6e206861736e277420626567756e000000000000000000000000600082015250565b600061268b601483611de7565b915061269682612655565b602082019050919050565b600060208201905081810360008301526126ba8161267e565b9050919050565b7f41756374696f6e2068617320616c7265616479206265656e20736574746c6564600082015250565b60006126f7602083611de7565b9150612702826126c1565b602082019050919050565b60006020820190508181036000830152612726816126ea565b9050919050565b7f41756374696f6e206861736e277420636f6d706c657465640000000000000000600082015250565b6000612763601883611de7565b915061276e8261272d565b602082019050919050565b6000602082019050818103600083015261279281612756565b9050919050565b60006127b46127af6127aa84611bb3565b6122c6565b611bb3565b9050919050565b60006127c682612799565b9050919050565b60006127d8826127bb565b9050919050565b6127e8816127cd565b82525050565b60006060820190506128036000830186611be5565b61281060208301856127df565b61281d6040830184611aff565b949350505050565b600060408201905061283a60008301856127df565b6128476020830184611aff565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600081519050919050565b600081905092915050565b60005b838110156128b1578082015181840152602081019050612896565b838111156128c0576000848401525b50505050565b60006128d18261287d565b6128db8185612888565b93506128eb818560208601612893565b80840191505092915050565b600061290382846128c6565b91508190509291505056fea2646970667358221220c6cc0bdc8bce18dc9e14d750480a05c4c4796ae411814a4476e4f08cbce6c76e64736f6c634300080a0033",
}

// TrixelsAuctionHouseABI is the input ABI used to generate the binding from.
// Deprecated: Use TrixelsAuctionHouseMetaData.ABI instead.
var TrixelsAuctionHouseABI = TrixelsAuctionHouseMetaData.ABI

// TrixelsAuctionHouseBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TrixelsAuctionHouseMetaData.Bin instead.
var TrixelsAuctionHouseBin = TrixelsAuctionHouseMetaData.Bin

// DeployTrixelsAuctionHouse deploys a new Ethereum contract, binding an instance of TrixelsAuctionHouse to it.
func DeployTrixelsAuctionHouse(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TrixelsAuctionHouse, error) {
	parsed, err := TrixelsAuctionHouseMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TrixelsAuctionHouseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TrixelsAuctionHouse{TrixelsAuctionHouseCaller: TrixelsAuctionHouseCaller{contract: contract}, TrixelsAuctionHouseTransactor: TrixelsAuctionHouseTransactor{contract: contract}, TrixelsAuctionHouseFilterer: TrixelsAuctionHouseFilterer{contract: contract}}, nil
}

// TrixelsAuctionHouse is an auto generated Go binding around an Ethereum contract.
type TrixelsAuctionHouse struct {
	TrixelsAuctionHouseCaller     // Read-only binding to the contract
	TrixelsAuctionHouseTransactor // Write-only binding to the contract
	TrixelsAuctionHouseFilterer   // Log filterer for contract events
}

// TrixelsAuctionHouseCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrixelsAuctionHouseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrixelsAuctionHouseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrixelsAuctionHouseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrixelsAuctionHouseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrixelsAuctionHouseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrixelsAuctionHouseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrixelsAuctionHouseSession struct {
	Contract     *TrixelsAuctionHouse // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TrixelsAuctionHouseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrixelsAuctionHouseCallerSession struct {
	Contract *TrixelsAuctionHouseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TrixelsAuctionHouseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrixelsAuctionHouseTransactorSession struct {
	Contract     *TrixelsAuctionHouseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TrixelsAuctionHouseRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrixelsAuctionHouseRaw struct {
	Contract *TrixelsAuctionHouse // Generic contract binding to access the raw methods on
}

// TrixelsAuctionHouseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrixelsAuctionHouseCallerRaw struct {
	Contract *TrixelsAuctionHouseCaller // Generic read-only contract binding to access the raw methods on
}

// TrixelsAuctionHouseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrixelsAuctionHouseTransactorRaw struct {
	Contract *TrixelsAuctionHouseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrixelsAuctionHouse creates a new instance of TrixelsAuctionHouse, bound to a specific deployed contract.
func NewTrixelsAuctionHouse(address common.Address, backend bind.ContractBackend) (*TrixelsAuctionHouse, error) {
	contract, err := bindTrixelsAuctionHouse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouse{TrixelsAuctionHouseCaller: TrixelsAuctionHouseCaller{contract: contract}, TrixelsAuctionHouseTransactor: TrixelsAuctionHouseTransactor{contract: contract}, TrixelsAuctionHouseFilterer: TrixelsAuctionHouseFilterer{contract: contract}}, nil
}

// NewTrixelsAuctionHouseCaller creates a new read-only instance of TrixelsAuctionHouse, bound to a specific deployed contract.
func NewTrixelsAuctionHouseCaller(address common.Address, caller bind.ContractCaller) (*TrixelsAuctionHouseCaller, error) {
	contract, err := bindTrixelsAuctionHouse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseCaller{contract: contract}, nil
}

// NewTrixelsAuctionHouseTransactor creates a new write-only instance of TrixelsAuctionHouse, bound to a specific deployed contract.
func NewTrixelsAuctionHouseTransactor(address common.Address, transactor bind.ContractTransactor) (*TrixelsAuctionHouseTransactor, error) {
	contract, err := bindTrixelsAuctionHouse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseTransactor{contract: contract}, nil
}

// NewTrixelsAuctionHouseFilterer creates a new log filterer instance of TrixelsAuctionHouse, bound to a specific deployed contract.
func NewTrixelsAuctionHouseFilterer(address common.Address, filterer bind.ContractFilterer) (*TrixelsAuctionHouseFilterer, error) {
	contract, err := bindTrixelsAuctionHouse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseFilterer{contract: contract}, nil
}

// bindTrixelsAuctionHouse binds a generic wrapper to an already deployed contract.
func bindTrixelsAuctionHouse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrixelsAuctionHouseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrixelsAuctionHouse *TrixelsAuctionHouseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrixelsAuctionHouse.Contract.TrixelsAuctionHouseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrixelsAuctionHouse *TrixelsAuctionHouseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.TrixelsAuctionHouseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrixelsAuctionHouse *TrixelsAuctionHouseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.TrixelsAuctionHouseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrixelsAuctionHouse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.contract.Transact(opts, method, params...)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrixelsAuctionHouse.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) Duration() (*big.Int, error) {
	return _TrixelsAuctionHouse.Contract.Duration(&_TrixelsAuctionHouse.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCallerSession) Duration() (*big.Int, error) {
	return _TrixelsAuctionHouse.Contract.Duration(&_TrixelsAuctionHouse.CallOpts)
}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCaller) MinBidIncrementPercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TrixelsAuctionHouse.contract.Call(opts, &out, "minBidIncrementPercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) MinBidIncrementPercentage() (uint8, error) {
	return _TrixelsAuctionHouse.Contract.MinBidIncrementPercentage(&_TrixelsAuctionHouse.CallOpts)
}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCallerSession) MinBidIncrementPercentage() (uint8, error) {
	return _TrixelsAuctionHouse.Contract.MinBidIncrementPercentage(&_TrixelsAuctionHouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrixelsAuctionHouse.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) Owner() (common.Address, error) {
	return _TrixelsAuctionHouse.Contract.Owner(&_TrixelsAuctionHouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCallerSession) Owner() (common.Address, error) {
	return _TrixelsAuctionHouse.Contract.Owner(&_TrixelsAuctionHouse.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TrixelsAuctionHouse.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) Paused() (bool, error) {
	return _TrixelsAuctionHouse.Contract.Paused(&_TrixelsAuctionHouse.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCallerSession) Paused() (bool, error) {
	return _TrixelsAuctionHouse.Contract.Paused(&_TrixelsAuctionHouse.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrixelsAuctionHouse.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) ReservePrice() (*big.Int, error) {
	return _TrixelsAuctionHouse.Contract.ReservePrice(&_TrixelsAuctionHouse.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCallerSession) ReservePrice() (*big.Int, error) {
	return _TrixelsAuctionHouse.Contract.ReservePrice(&_TrixelsAuctionHouse.CallOpts)
}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCaller) TimeBuffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrixelsAuctionHouse.contract.Call(opts, &out, "timeBuffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) TimeBuffer() (*big.Int, error) {
	return _TrixelsAuctionHouse.Contract.TimeBuffer(&_TrixelsAuctionHouse.CallOpts)
}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint256)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCallerSession) TimeBuffer() (*big.Int, error) {
	return _TrixelsAuctionHouse.Contract.TimeBuffer(&_TrixelsAuctionHouse.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrixelsAuctionHouse.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) Weth() (common.Address, error) {
	return _TrixelsAuctionHouse.Contract.Weth(&_TrixelsAuctionHouse.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseCallerSession) Weth() (common.Address, error) {
	return _TrixelsAuctionHouse.Contract.Weth(&_TrixelsAuctionHouse.CallOpts)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 trixelId) payable returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) CreateBid(opts *bind.TransactOpts, trixelId *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "createBid", trixelId)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 trixelId) payable returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) CreateBid(trixelId *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.CreateBid(&_TrixelsAuctionHouse.TransactOpts, trixelId)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 trixelId) payable returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) CreateBid(trixelId *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.CreateBid(&_TrixelsAuctionHouse.TransactOpts, trixelId)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c28e349.
//
// Solidity: function initialize(address _trixels, uint256 _duration, uint256 _timeBuffer, uint256 _reservePrice, address _weth) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) Initialize(opts *bind.TransactOpts, _trixels common.Address, _duration *big.Int, _timeBuffer *big.Int, _reservePrice *big.Int, _weth common.Address) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "initialize", _trixels, _duration, _timeBuffer, _reservePrice, _weth)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c28e349.
//
// Solidity: function initialize(address _trixels, uint256 _duration, uint256 _timeBuffer, uint256 _reservePrice, address _weth) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) Initialize(_trixels common.Address, _duration *big.Int, _timeBuffer *big.Int, _reservePrice *big.Int, _weth common.Address) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.Initialize(&_TrixelsAuctionHouse.TransactOpts, _trixels, _duration, _timeBuffer, _reservePrice, _weth)
}

// Initialize is a paid mutator transaction binding the contract method 0x6c28e349.
//
// Solidity: function initialize(address _trixels, uint256 _duration, uint256 _timeBuffer, uint256 _reservePrice, address _weth) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) Initialize(_trixels common.Address, _duration *big.Int, _timeBuffer *big.Int, _reservePrice *big.Int, _weth common.Address) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.Initialize(&_TrixelsAuctionHouse.TransactOpts, _trixels, _duration, _timeBuffer, _reservePrice, _weth)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) Pause() (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.Pause(&_TrixelsAuctionHouse.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) Pause() (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.Pause(&_TrixelsAuctionHouse.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) RenounceOwnership() (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.RenounceOwnership(&_TrixelsAuctionHouse.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.RenounceOwnership(&_TrixelsAuctionHouse.TransactOpts)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) SetReservePrice(opts *bind.TransactOpts, _reservePrice *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "setReservePrice", _reservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) SetReservePrice(_reservePrice *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.SetReservePrice(&_TrixelsAuctionHouse.TransactOpts, _reservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) SetReservePrice(_reservePrice *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.SetReservePrice(&_TrixelsAuctionHouse.TransactOpts, _reservePrice)
}

// SetTimeBuffer is a paid mutator transaction binding the contract method 0x7120334b.
//
// Solidity: function setTimeBuffer(uint256 _timeBuffer) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) SetTimeBuffer(opts *bind.TransactOpts, _timeBuffer *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "setTimeBuffer", _timeBuffer)
}

// SetTimeBuffer is a paid mutator transaction binding the contract method 0x7120334b.
//
// Solidity: function setTimeBuffer(uint256 _timeBuffer) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) SetTimeBuffer(_timeBuffer *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.SetTimeBuffer(&_TrixelsAuctionHouse.TransactOpts, _timeBuffer)
}

// SetTimeBuffer is a paid mutator transaction binding the contract method 0x7120334b.
//
// Solidity: function setTimeBuffer(uint256 _timeBuffer) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) SetTimeBuffer(_timeBuffer *big.Int) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.SetTimeBuffer(&_TrixelsAuctionHouse.TransactOpts, _timeBuffer)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xa4d0a17e.
//
// Solidity: function settleAuction() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) SettleAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "settleAuction")
}

// SettleAuction is a paid mutator transaction binding the contract method 0xa4d0a17e.
//
// Solidity: function settleAuction() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) SettleAuction() (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.SettleAuction(&_TrixelsAuctionHouse.TransactOpts)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xa4d0a17e.
//
// Solidity: function settleAuction() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) SettleAuction() (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.SettleAuction(&_TrixelsAuctionHouse.TransactOpts)
}

// StartAuction is a paid mutator transaction binding the contract method 0x2b107ca2.
//
// Solidity: function startAuction(bytes23 skyNetID) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) StartAuction(opts *bind.TransactOpts, skyNetID [23]byte) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "startAuction", skyNetID)
}

// StartAuction is a paid mutator transaction binding the contract method 0x2b107ca2.
//
// Solidity: function startAuction(bytes23 skyNetID) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) StartAuction(skyNetID [23]byte) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.StartAuction(&_TrixelsAuctionHouse.TransactOpts, skyNetID)
}

// StartAuction is a paid mutator transaction binding the contract method 0x2b107ca2.
//
// Solidity: function startAuction(bytes23 skyNetID) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) StartAuction(skyNetID [23]byte) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.StartAuction(&_TrixelsAuctionHouse.TransactOpts, skyNetID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.TransferOwnership(&_TrixelsAuctionHouse.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.TransferOwnership(&_TrixelsAuctionHouse.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrixelsAuctionHouse.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseSession) Unpause() (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.Unpause(&_TrixelsAuctionHouse.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TrixelsAuctionHouse *TrixelsAuctionHouseTransactorSession) Unpause() (*types.Transaction, error) {
	return _TrixelsAuctionHouse.Contract.Unpause(&_TrixelsAuctionHouse.TransactOpts)
}

// TrixelsAuctionHouseAuctionBidIterator is returned from FilterAuctionBid and is used to iterate over the raw logs and unpacked data for AuctionBid events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionBidIterator struct {
	Event *TrixelsAuctionHouseAuctionBid // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseAuctionBid)
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
		it.Event = new(TrixelsAuctionHouseAuctionBid)
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
func (it *TrixelsAuctionHouseAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseAuctionBid represents a AuctionBid event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionBid struct {
	TrixelId *big.Int
	Sender   common.Address
	Value    *big.Int
	Extended bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuctionBid is a free log retrieval operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed trixelId, address sender, uint256 value, bool extended)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterAuctionBid(opts *bind.FilterOpts, trixelId []*big.Int) (*TrixelsAuctionHouseAuctionBidIterator, error) {

	var trixelIdRule []interface{}
	for _, trixelIdItem := range trixelId {
		trixelIdRule = append(trixelIdRule, trixelIdItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "AuctionBid", trixelIdRule)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseAuctionBidIterator{contract: _TrixelsAuctionHouse.contract, event: "AuctionBid", logs: logs, sub: sub}, nil
}

// WatchAuctionBid is a free log subscription operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed trixelId, address sender, uint256 value, bool extended)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchAuctionBid(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseAuctionBid, trixelId []*big.Int) (event.Subscription, error) {

	var trixelIdRule []interface{}
	for _, trixelIdItem := range trixelId {
		trixelIdRule = append(trixelIdRule, trixelIdItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "AuctionBid", trixelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseAuctionBid)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionBid", log); err != nil {
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

// ParseAuctionBid is a log parse operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed trixelId, address sender, uint256 value, bool extended)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseAuctionBid(log types.Log) (*TrixelsAuctionHouseAuctionBid, error) {
	event := new(TrixelsAuctionHouseAuctionBid)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionCreatedIterator struct {
	Event *TrixelsAuctionHouseAuctionCreated // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseAuctionCreated)
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
		it.Event = new(TrixelsAuctionHouseAuctionCreated)
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
func (it *TrixelsAuctionHouseAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseAuctionCreated represents a AuctionCreated event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionCreated struct {
	TrixelId  *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed trixelId, uint256 startTime, uint256 endTime)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterAuctionCreated(opts *bind.FilterOpts, trixelId []*big.Int) (*TrixelsAuctionHouseAuctionCreatedIterator, error) {

	var trixelIdRule []interface{}
	for _, trixelIdItem := range trixelId {
		trixelIdRule = append(trixelIdRule, trixelIdItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "AuctionCreated", trixelIdRule)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseAuctionCreatedIterator{contract: _TrixelsAuctionHouse.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed trixelId, uint256 startTime, uint256 endTime)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseAuctionCreated, trixelId []*big.Int) (event.Subscription, error) {

	var trixelIdRule []interface{}
	for _, trixelIdItem := range trixelId {
		trixelIdRule = append(trixelIdRule, trixelIdItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "AuctionCreated", trixelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseAuctionCreated)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed trixelId, uint256 startTime, uint256 endTime)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseAuctionCreated(log types.Log) (*TrixelsAuctionHouseAuctionCreated, error) {
	event := new(TrixelsAuctionHouseAuctionCreated)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseAuctionExtendedIterator is returned from FilterAuctionExtended and is used to iterate over the raw logs and unpacked data for AuctionExtended events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionExtendedIterator struct {
	Event *TrixelsAuctionHouseAuctionExtended // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseAuctionExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseAuctionExtended)
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
		it.Event = new(TrixelsAuctionHouseAuctionExtended)
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
func (it *TrixelsAuctionHouseAuctionExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseAuctionExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseAuctionExtended represents a AuctionExtended event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionExtended struct {
	NounId  *big.Int
	EndTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionExtended is a free log retrieval operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed nounId, uint256 endTime)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterAuctionExtended(opts *bind.FilterOpts, nounId []*big.Int) (*TrixelsAuctionHouseAuctionExtendedIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "AuctionExtended", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseAuctionExtendedIterator{contract: _TrixelsAuctionHouse.contract, event: "AuctionExtended", logs: logs, sub: sub}, nil
}

// WatchAuctionExtended is a free log subscription operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed nounId, uint256 endTime)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchAuctionExtended(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseAuctionExtended, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "AuctionExtended", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseAuctionExtended)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionExtended", log); err != nil {
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

// ParseAuctionExtended is a log parse operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed nounId, uint256 endTime)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseAuctionExtended(log types.Log) (*TrixelsAuctionHouseAuctionExtended, error) {
	event := new(TrixelsAuctionHouseAuctionExtended)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator is returned from FilterAuctionMinBidIncrementPercentageUpdated and is used to iterate over the raw logs and unpacked data for AuctionMinBidIncrementPercentageUpdated events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator struct {
	Event *TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated)
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
		it.Event = new(TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated)
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
func (it *TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated represents a AuctionMinBidIncrementPercentageUpdated event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated struct {
	MinBidIncrementPercentage *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterAuctionMinBidIncrementPercentageUpdated is a free log retrieval operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterAuctionMinBidIncrementPercentageUpdated(opts *bind.FilterOpts) (*TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "AuctionMinBidIncrementPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator{contract: _TrixelsAuctionHouse.contract, event: "AuctionMinBidIncrementPercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionMinBidIncrementPercentageUpdated is a free log subscription operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchAuctionMinBidIncrementPercentageUpdated(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "AuctionMinBidIncrementPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionMinBidIncrementPercentageUpdated", log); err != nil {
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

// ParseAuctionMinBidIncrementPercentageUpdated is a log parse operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseAuctionMinBidIncrementPercentageUpdated(log types.Log) (*TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated, error) {
	event := new(TrixelsAuctionHouseAuctionMinBidIncrementPercentageUpdated)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionMinBidIncrementPercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseAuctionReservePriceUpdatedIterator is returned from FilterAuctionReservePriceUpdated and is used to iterate over the raw logs and unpacked data for AuctionReservePriceUpdated events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionReservePriceUpdatedIterator struct {
	Event *TrixelsAuctionHouseAuctionReservePriceUpdated // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseAuctionReservePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseAuctionReservePriceUpdated)
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
		it.Event = new(TrixelsAuctionHouseAuctionReservePriceUpdated)
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
func (it *TrixelsAuctionHouseAuctionReservePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseAuctionReservePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseAuctionReservePriceUpdated represents a AuctionReservePriceUpdated event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionReservePriceUpdated struct {
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionReservePriceUpdated is a free log retrieval operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterAuctionReservePriceUpdated(opts *bind.FilterOpts) (*TrixelsAuctionHouseAuctionReservePriceUpdatedIterator, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "AuctionReservePriceUpdated")
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseAuctionReservePriceUpdatedIterator{contract: _TrixelsAuctionHouse.contract, event: "AuctionReservePriceUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionReservePriceUpdated is a free log subscription operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchAuctionReservePriceUpdated(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseAuctionReservePriceUpdated) (event.Subscription, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "AuctionReservePriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseAuctionReservePriceUpdated)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionReservePriceUpdated", log); err != nil {
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

// ParseAuctionReservePriceUpdated is a log parse operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseAuctionReservePriceUpdated(log types.Log) (*TrixelsAuctionHouseAuctionReservePriceUpdated, error) {
	event := new(TrixelsAuctionHouseAuctionReservePriceUpdated)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionReservePriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionSettledIterator struct {
	Event *TrixelsAuctionHouseAuctionSettled // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseAuctionSettled)
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
		it.Event = new(TrixelsAuctionHouseAuctionSettled)
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
func (it *TrixelsAuctionHouseAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseAuctionSettled represents a AuctionSettled event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionSettled struct {
	TrixelId *big.Int
	Winner   common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed trixelId, address winner, uint256 amount)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterAuctionSettled(opts *bind.FilterOpts, trixelId []*big.Int) (*TrixelsAuctionHouseAuctionSettledIterator, error) {

	var trixelIdRule []interface{}
	for _, trixelIdItem := range trixelId {
		trixelIdRule = append(trixelIdRule, trixelIdItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "AuctionSettled", trixelIdRule)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseAuctionSettledIterator{contract: _TrixelsAuctionHouse.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed trixelId, address winner, uint256 amount)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseAuctionSettled, trixelId []*big.Int) (event.Subscription, error) {

	var trixelIdRule []interface{}
	for _, trixelIdItem := range trixelId {
		trixelIdRule = append(trixelIdRule, trixelIdItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "AuctionSettled", trixelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseAuctionSettled)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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

// ParseAuctionSettled is a log parse operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed trixelId, address winner, uint256 amount)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseAuctionSettled(log types.Log) (*TrixelsAuctionHouseAuctionSettled, error) {
	event := new(TrixelsAuctionHouseAuctionSettled)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseAuctionTimeBufferUpdatedIterator is returned from FilterAuctionTimeBufferUpdated and is used to iterate over the raw logs and unpacked data for AuctionTimeBufferUpdated events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionTimeBufferUpdatedIterator struct {
	Event *TrixelsAuctionHouseAuctionTimeBufferUpdated // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseAuctionTimeBufferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseAuctionTimeBufferUpdated)
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
		it.Event = new(TrixelsAuctionHouseAuctionTimeBufferUpdated)
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
func (it *TrixelsAuctionHouseAuctionTimeBufferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseAuctionTimeBufferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseAuctionTimeBufferUpdated represents a AuctionTimeBufferUpdated event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseAuctionTimeBufferUpdated struct {
	TimeBuffer *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionTimeBufferUpdated is a free log retrieval operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterAuctionTimeBufferUpdated(opts *bind.FilterOpts) (*TrixelsAuctionHouseAuctionTimeBufferUpdatedIterator, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "AuctionTimeBufferUpdated")
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseAuctionTimeBufferUpdatedIterator{contract: _TrixelsAuctionHouse.contract, event: "AuctionTimeBufferUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionTimeBufferUpdated is a free log subscription operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchAuctionTimeBufferUpdated(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseAuctionTimeBufferUpdated) (event.Subscription, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "AuctionTimeBufferUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseAuctionTimeBufferUpdated)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionTimeBufferUpdated", log); err != nil {
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

// ParseAuctionTimeBufferUpdated is a log parse operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseAuctionTimeBufferUpdated(log types.Log) (*TrixelsAuctionHouseAuctionTimeBufferUpdated, error) {
	event := new(TrixelsAuctionHouseAuctionTimeBufferUpdated)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "AuctionTimeBufferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseInitializedIterator struct {
	Event *TrixelsAuctionHouseInitialized // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseInitialized)
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
		it.Event = new(TrixelsAuctionHouseInitialized)
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
func (it *TrixelsAuctionHouseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseInitialized represents a Initialized event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterInitialized(opts *bind.FilterOpts) (*TrixelsAuctionHouseInitializedIterator, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseInitializedIterator{contract: _TrixelsAuctionHouse.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseInitialized) (event.Subscription, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseInitialized)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseInitialized(log types.Log) (*TrixelsAuctionHouseInitialized, error) {
	event := new(TrixelsAuctionHouseInitialized)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseOwnershipTransferredIterator struct {
	Event *TrixelsAuctionHouseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseOwnershipTransferred)
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
		it.Event = new(TrixelsAuctionHouseOwnershipTransferred)
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
func (it *TrixelsAuctionHouseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseOwnershipTransferred represents a OwnershipTransferred event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TrixelsAuctionHouseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseOwnershipTransferredIterator{contract: _TrixelsAuctionHouse.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseOwnershipTransferred)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseOwnershipTransferred(log types.Log) (*TrixelsAuctionHouseOwnershipTransferred, error) {
	event := new(TrixelsAuctionHouseOwnershipTransferred)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHousePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHousePausedIterator struct {
	Event *TrixelsAuctionHousePaused // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHousePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHousePaused)
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
		it.Event = new(TrixelsAuctionHousePaused)
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
func (it *TrixelsAuctionHousePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHousePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHousePaused represents a Paused event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHousePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterPaused(opts *bind.FilterOpts) (*TrixelsAuctionHousePausedIterator, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHousePausedIterator{contract: _TrixelsAuctionHouse.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHousePaused) (event.Subscription, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHousePaused)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParsePaused(log types.Log) (*TrixelsAuctionHousePaused, error) {
	event := new(TrixelsAuctionHousePaused)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrixelsAuctionHouseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseUnpausedIterator struct {
	Event *TrixelsAuctionHouseUnpaused // Event containing the contract specifics and raw log

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
func (it *TrixelsAuctionHouseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrixelsAuctionHouseUnpaused)
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
		it.Event = new(TrixelsAuctionHouseUnpaused)
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
func (it *TrixelsAuctionHouseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrixelsAuctionHouseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrixelsAuctionHouseUnpaused represents a Unpaused event raised by the TrixelsAuctionHouse contract.
type TrixelsAuctionHouseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TrixelsAuctionHouseUnpausedIterator, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TrixelsAuctionHouseUnpausedIterator{contract: _TrixelsAuctionHouse.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TrixelsAuctionHouseUnpaused) (event.Subscription, error) {

	logs, sub, err := _TrixelsAuctionHouse.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrixelsAuctionHouseUnpaused)
				if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TrixelsAuctionHouse *TrixelsAuctionHouseFilterer) ParseUnpaused(log types.Log) (*TrixelsAuctionHouseUnpaused, error) {
	event := new(TrixelsAuctionHouseUnpaused)
	if err := _TrixelsAuctionHouse.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
