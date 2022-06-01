// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DAO

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

// DAOMetaData contains all meta data concerning the DAO contract.
var DAOMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"inSupport\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"numProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"yay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nay\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"person\",\"type\":\"address\"}],\"name\":\"isMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_transactionData\",\"type\":\"bytes\"}],\"name\":\"makeProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_inSupport\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"}],\"name\":\"unVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_transactionData\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_transactionData\",\"type\":\"bytes\"}],\"name\":\"checkProposalHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"validHash\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001e1538038062001e158339818101604052810190620000379190620000e1565b80806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600060028190555050506200017b565b600081519050620000db8162000161565b92915050565b60008060408385031215620000fb57620000fa6200015c565b5b60006200010b85828601620000ca565b92505060206200011e85828601620000ca565b9150509250929050565b600062000135826200013c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600080fd5b6200016c8162000128565b81146200017857600080fd5b50565b611c8a806200018b6000396000f3fe60806040526004361061008a5760003560e01c8063909d02ad11610059578063909d02ad146101805780639fd6adb3146101bd578063a230c524146101e6578063c9d27afe14610223578063fc0c546a1461024c57610091565b8063013cf08b14610096578063237e9492146100db5780632891d6f214610118578063400e39491461015557610091565b3661009157005b600080fd5b3480156100a257600080fd5b506100bd60048036038101906100b89190611254565b610277565b6040516100d299989796959493929190611676565b60405180910390f35b3480156100e757600080fd5b5061010260048036038101906100fd9190611376565b61039a565b60405161010f919061165b565b60405180910390f35b34801561012457600080fd5b5061013f600480360381019061013a91906112ae565b61061c565b60405161014c919061165b565b60405180910390f35b34801561016157600080fd5b5061016a610673565b60405161017791906117e5565b60405180910390f35b34801561018c57600080fd5b506101a760048036038101906101a291906111ad565b610679565b6040516101b491906117e5565b60405180910390f35b3480156101c957600080fd5b506101e460048036038101906101df9190611254565b610837565b005b3480156101f257600080fd5b5061020d60048036038101906102089190611180565b610b82565b60405161021a919061165b565b60405180910390f35b34801561022f57600080fd5b5061024a60048036038101906102459190611336565b610c38565b005b34801561025857600080fd5b50610261610f35565b60405161026e919061170a565b60405180910390f35b60036020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030180546102cc90611969565b80601f01602080910402602001604051908101604052809291908181526020018280546102f890611969565b80156103455780601f1061031a57610100808354040283529160200191610345565b820191906000526020600020905b81548152906001019060200180831161032857829003601f168201915b5050505050908060040160009054906101000a900460ff16908060070154908060080154908060090160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600a0154905089565b60006103a533610b82565b6103e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103db906117c5565b60405180910390fd5b6000600360008681526020019081526020016000209050426212750082600a015461040f9190611827565b1061044f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610446906117a5565b60405180910390fd5b8060080154816007015411610499576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049090611725565b60405180910390fd5b8060040160009054906101000a900460ff16156104eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e290611785565b60405180910390fd5b8060020154471015610532576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052990611745565b60405180910390fd5b610567858260010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360020154878761061c565b6105a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059d90611765565b60405180910390fd5b6105da8160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682600201548686610f5b565b9150847f948f4a9cd986f1118c3fbd459f7a22b23c0693e1ca3ef06a6a8be5aa7d39cc038360405161060c919061165b565b60405180910390a2509392505050565b6000806003600088815260200190815260200160002090508585858560405160200161064b94939291906115ac565b6040516020818303038152906040528051906020012081600001541491505095945050505050565b60025481565b600061068433610b82565b6106c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ba906117c5565b60405180910390fd5b600060036000600260008154809291906106dc9061199b565b91905055815260200190815260200160002090508787858560405160200161070794939291906115ac565b604051602081830303815290604052805190602001208160000181905550878160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508681600201819055508585826003019190610784929190610fdd565b5060008160040160006101000a81548160ff021916908315150217905550338160090160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055504281600a0181905550817f646fec02522b41e7125cfc859a64fd4f4cefd5dc3b6237ca0abe251ded1fa88189898989604051610824949392919061161b565b60405180910390a2509695505050505050565b61084033610b82565b61087f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610876906117c5565b60405180910390fd5b60006003600083815260200190815260200160002090508060050160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610a0a57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b81526004016109459190611600565b60206040518083038186803b15801561095d57600080fd5b505afa158015610971573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109959190611281565b8160070160008282546109a8919061187d565b9250508190555060008160060160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b8060060160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615610b7e57600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401610ab99190611600565b60206040518083038186803b158015610ad157600080fd5b505afa158015610ae5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b099190611281565b816008016000828254610b1c919061187d565b9250508190555060008160060160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b5050565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231846040518263ffffffff1660e01b8152600401610be09190611600565b60206040518083038186803b158015610bf857600080fd5b505afa158015610c0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c309190611281565b119050919050565b610c4133610b82565b610c80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c77906117c5565b60405180910390fd5b60006003600084815260200190815260200160002090508115610dc157600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401610cf89190611600565b60206040518083038186803b158015610d1057600080fd5b505afa158015610d24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d489190611281565b816007016000828254610d5b9190611827565b9250508190555060018160050160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610ee1565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401610e1c9190611600565b60206040518083038186803b158015610e3457600080fd5b505afa158015610e48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6c9190611281565b816008016000828254610e7f9190611827565b9250508190555060018160060160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b3373ffffffffffffffffffffffffffffffffffffffff16837f86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae084604051610f28919061165b565b60405180910390a3505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000808573ffffffffffffffffffffffffffffffffffffffff1685617530908686604051610f8a9291906115e7565b600060405180830381858888f193505050503d8060008114610fc8576040519150601f19603f3d011682016040523d82523d6000602084013e610fcd565b606091505b5050905080915050949350505050565b828054610fe990611969565b90600052602060002090601f01602090048101928261100b5760008555611052565b82601f1061102457803560ff1916838001178555611052565b82800160010185558215611052579182015b82811115611051578235825591602001919060010190611036565b5b50905061105f9190611063565b5090565b5b8082111561107c576000816000905550600101611064565b5090565b60008135905061108f81611c0f565b92915050565b6000813590506110a481611c26565b92915050565b60008083601f8401126110c0576110bf611a75565b5b8235905067ffffffffffffffff8111156110dd576110dc611a70565b5b6020830191508360018202830111156110f9576110f8611a7a565b5b9250929050565b60008083601f84011261111657611115611a75565b5b8235905067ffffffffffffffff81111561113357611132611a70565b5b60208301915083600182028301111561114f5761114e611a7a565b5b9250929050565b60008135905061116581611c3d565b92915050565b60008151905061117a81611c3d565b92915050565b60006020828403121561119657611195611a84565b5b60006111a484828501611080565b91505092915050565b600080600080600080608087890312156111ca576111c9611a84565b5b60006111d889828a01611080565b96505060206111e989828a01611156565b955050604087013567ffffffffffffffff81111561120a57611209611a7f565b5b61121689828a01611100565b9450945050606087013567ffffffffffffffff81111561123957611238611a7f565b5b61124589828a016110aa565b92509250509295509295509295565b60006020828403121561126a57611269611a84565b5b600061127884828501611156565b91505092915050565b60006020828403121561129757611296611a84565b5b60006112a58482850161116b565b91505092915050565b6000806000806000608086880312156112ca576112c9611a84565b5b60006112d888828901611156565b95505060206112e988828901611080565b94505060406112fa88828901611156565b935050606086013567ffffffffffffffff81111561131b5761131a611a7f565b5b611327888289016110aa565b92509250509295509295909350565b6000806040838503121561134d5761134c611a84565b5b600061135b85828601611156565b925050602061136c85828601611095565b9150509250929050565b60008060006040848603121561138f5761138e611a84565b5b600061139d86828701611156565b935050602084013567ffffffffffffffff8111156113be576113bd611a7f565b5b6113ca868287016110aa565b92509250509250925092565b6113df816118b1565b82525050565b6113f66113f1826118b1565b6119e4565b82525050565b611405816118c3565b82525050565b611414816118cf565b82525050565b6000611426838561180b565b9350611433838584611927565b82840190509392505050565b61144881611903565b82525050565b600061145a8385611816565b9350611467838584611927565b61147083611a89565b840190509392505050565b600061148682611800565b6114908185611816565b93506114a0818560208601611936565b6114a981611a89565b840191505092915050565b60006114c1602183611816565b91506114cc82611aa7565b604082019050919050565b60006114e4602c83611816565b91506114ef82611af6565b604082019050919050565b6000611507601883611816565b915061151282611b45565b602082019050919050565b600061152a602083611816565b915061153582611b6e565b602082019050919050565b600061154d603483611816565b915061155882611b97565b604082019050919050565b6000611570600f83611816565b915061157b82611be6565b602082019050919050565b61158f816118f9565b82525050565b6115a66115a1826118f9565b611a08565b82525050565b60006115b882876113e5565b6014820191506115c88286611595565b6020820191506115d982848661141a565b915081905095945050505050565b60006115f482848661141a565b91508190509392505050565b600060208201905061161560008301846113d6565b92915050565b600060608201905061163060008301876113d6565b61163d6020830186611586565b818103604083015261165081848661144e565b905095945050505050565b600060208201905061167060008301846113fc565b92915050565b60006101208201905061168c600083018c61140b565b611699602083018b6113d6565b6116a6604083018a611586565b81810360608301526116b8818961147b565b90506116c760808301886113fc565b6116d460a0830187611586565b6116e160c0830186611586565b6116ee60e08301856113d6565b6116fc610100830184611586565b9a9950505050505050505050565b600060208201905061171f600083018461143f565b92915050565b6000602082019050818103600083015261173e816114b4565b9050919050565b6000602082019050818103600083015261175e816114d7565b9050919050565b6000602082019050818103600083015261177e816114fa565b9050919050565b6000602082019050818103600083015261179e8161151d565b9050919050565b600060208201905081810360008301526117be81611540565b9050919050565b600060208201905081810360008301526117de81611563565b9050919050565b60006020820190506117fa6000830184611586565b92915050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000611832826118f9565b915061183d836118f9565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561187257611871611a12565b5b828201905092915050565b6000611888826118f9565b9150611893836118f9565b9250828210156118a6576118a5611a12565b5b828203905092915050565b60006118bc826118d9565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061190e82611915565b9050919050565b6000611920826118d9565b9050919050565b82818337600083830152505050565b60005b83811015611954578082015181840152602081019050611939565b83811115611963576000848401525b50505050565b6000600282049050600182168061198157607f821691505b6020821081141561199557611994611a41565b5b50919050565b60006119a6826118f9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156119d9576119d8611a12565b5b600182019050919050565b60006119ef826119f6565b9050919050565b6000611a0182611a9a565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f59617920766f74657320646f206e6f7420657863656564204e617920766f746560008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b7f44414f20646f6573206e6f74206861766520656e6f75676820746f206578656360008201527f7574652070726f706f73616c0000000000000000000000000000000000000000602082015250565b7f496e76616c6964207472616e73616374696f6e20646174610000000000000000600082015250565b7f50726f706f73616c2068617320616c7265616479206265656e20706173736564600082015250565b7f43616e6e6f7420657865637574652070726f706f73616c2c20766f74696e672060008201527f706572696f6420686173206e6f7420656e646564000000000000000000000000602082015250565b7f4e6f742070617274206f662044414f0000000000000000000000000000000000600082015250565b611c18816118b1565b8114611c2357600080fd5b50565b611c2f816118c3565b8114611c3a57600080fd5b50565b611c46816118f9565b8114611c5157600080fd5b5056fea26469706673582212201c88177a8a341c93a49f3bd829c297ab8a137d8daca1a770f5580100989358d464736f6c63430008060033",
}

// DAOABI is the input ABI used to generate the binding from.
// Deprecated: Use DAOMetaData.ABI instead.
var DAOABI = DAOMetaData.ABI

// DAOBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DAOMetaData.Bin instead.
var DAOBin = DAOMetaData.Bin

// DeployDAO deploys a new Ethereum contract, binding an instance of DAO to it.
func DeployDAO(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _weth common.Address) (common.Address, *types.Transaction, *DAO, error) {
	parsed, err := DAOMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DAOBin), backend, _token, _weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DAO{DAOCaller: DAOCaller{contract: contract}, DAOTransactor: DAOTransactor{contract: contract}, DAOFilterer: DAOFilterer{contract: contract}}, nil
}

// DAO is an auto generated Go binding around an Ethereum contract.
type DAO struct {
	DAOCaller     // Read-only binding to the contract
	DAOTransactor // Write-only binding to the contract
	DAOFilterer   // Log filterer for contract events
}

// DAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAOSession struct {
	Contract     *DAO              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAOCallerSession struct {
	Contract *DAOCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAOTransactorSession struct {
	Contract     *DAOTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAORaw is an auto generated low-level Go binding around an Ethereum contract.
type DAORaw struct {
	Contract *DAO // Generic contract binding to access the raw methods on
}

// DAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAOCallerRaw struct {
	Contract *DAOCaller // Generic read-only contract binding to access the raw methods on
}

// DAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAOTransactorRaw struct {
	Contract *DAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAO creates a new instance of DAO, bound to a specific deployed contract.
func NewDAO(address common.Address, backend bind.ContractBackend) (*DAO, error) {
	contract, err := bindDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAO{DAOCaller: DAOCaller{contract: contract}, DAOTransactor: DAOTransactor{contract: contract}, DAOFilterer: DAOFilterer{contract: contract}}, nil
}

// NewDAOCaller creates a new read-only instance of DAO, bound to a specific deployed contract.
func NewDAOCaller(address common.Address, caller bind.ContractCaller) (*DAOCaller, error) {
	contract, err := bindDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAOCaller{contract: contract}, nil
}

// NewDAOTransactor creates a new write-only instance of DAO, bound to a specific deployed contract.
func NewDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*DAOTransactor, error) {
	contract, err := bindDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAOTransactor{contract: contract}, nil
}

// NewDAOFilterer creates a new log filterer instance of DAO, bound to a specific deployed contract.
func NewDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*DAOFilterer, error) {
	contract, err := bindDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAOFilterer{contract: contract}, nil
}

// bindDAO binds a generic wrapper to an already deployed contract.
func bindDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DAOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAO *DAORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAO.Contract.DAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAO *DAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAO.Contract.DAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAO *DAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAO.Contract.DAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAO *DAOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAO *DAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAO *DAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAO.Contract.contract.Transact(opts, method, params...)
}

// CheckProposalHash is a free data retrieval call binding the contract method 0x2891d6f2.
//
// Solidity: function checkProposalHash(uint256 _proposalID, address _recipient, uint256 _amount, bytes _transactionData) view returns(bool validHash)
func (_DAO *DAOCaller) CheckProposalHash(opts *bind.CallOpts, _proposalID *big.Int, _recipient common.Address, _amount *big.Int, _transactionData []byte) (bool, error) {
	var out []interface{}
	err := _DAO.contract.Call(opts, &out, "checkProposalHash", _proposalID, _recipient, _amount, _transactionData)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProposalHash is a free data retrieval call binding the contract method 0x2891d6f2.
//
// Solidity: function checkProposalHash(uint256 _proposalID, address _recipient, uint256 _amount, bytes _transactionData) view returns(bool validHash)
func (_DAO *DAOSession) CheckProposalHash(_proposalID *big.Int, _recipient common.Address, _amount *big.Int, _transactionData []byte) (bool, error) {
	return _DAO.Contract.CheckProposalHash(&_DAO.CallOpts, _proposalID, _recipient, _amount, _transactionData)
}

// CheckProposalHash is a free data retrieval call binding the contract method 0x2891d6f2.
//
// Solidity: function checkProposalHash(uint256 _proposalID, address _recipient, uint256 _amount, bytes _transactionData) view returns(bool validHash)
func (_DAO *DAOCallerSession) CheckProposalHash(_proposalID *big.Int, _recipient common.Address, _amount *big.Int, _transactionData []byte) (bool, error) {
	return _DAO.Contract.CheckProposalHash(&_DAO.CallOpts, _proposalID, _recipient, _amount, _transactionData)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address person) view returns(bool)
func (_DAO *DAOCaller) IsMember(opts *bind.CallOpts, person common.Address) (bool, error) {
	var out []interface{}
	err := _DAO.contract.Call(opts, &out, "isMember", person)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address person) view returns(bool)
func (_DAO *DAOSession) IsMember(person common.Address) (bool, error) {
	return _DAO.Contract.IsMember(&_DAO.CallOpts, person)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address person) view returns(bool)
func (_DAO *DAOCallerSession) IsMember(person common.Address) (bool, error) {
	return _DAO.Contract.IsMember(&_DAO.CallOpts, person)
}

// NumProposals is a free data retrieval call binding the contract method 0x400e3949.
//
// Solidity: function numProposals() view returns(uint256)
func (_DAO *DAOCaller) NumProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DAO.contract.Call(opts, &out, "numProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumProposals is a free data retrieval call binding the contract method 0x400e3949.
//
// Solidity: function numProposals() view returns(uint256)
func (_DAO *DAOSession) NumProposals() (*big.Int, error) {
	return _DAO.Contract.NumProposals(&_DAO.CallOpts)
}

// NumProposals is a free data retrieval call binding the contract method 0x400e3949.
//
// Solidity: function numProposals() view returns(uint256)
func (_DAO *DAOCallerSession) NumProposals() (*big.Int, error) {
	return _DAO.Contract.NumProposals(&_DAO.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 proposalHash, address recipient, uint256 amount, string description, bool passed, uint256 yay, uint256 nay, address creator, uint256 createdAt)
func (_DAO *DAOCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _DAO.contract.Call(opts, &out, "proposals", arg0)

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
func (_DAO *DAOSession) Proposals(arg0 *big.Int) (struct {
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
	return _DAO.Contract.Proposals(&_DAO.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 proposalHash, address recipient, uint256 amount, string description, bool passed, uint256 yay, uint256 nay, address creator, uint256 createdAt)
func (_DAO *DAOCallerSession) Proposals(arg0 *big.Int) (struct {
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
	return _DAO.Contract.Proposals(&_DAO.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DAO *DAOCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DAO.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DAO *DAOSession) Token() (common.Address, error) {
	return _DAO.Contract.Token(&_DAO.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_DAO *DAOCallerSession) Token() (common.Address, error) {
	return _DAO.Contract.Token(&_DAO.CallOpts)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 _proposalID, bytes _transactionData) returns(bool success)
func (_DAO *DAOTransactor) ExecuteProposal(opts *bind.TransactOpts, _proposalID *big.Int, _transactionData []byte) (*types.Transaction, error) {
	return _DAO.contract.Transact(opts, "executeProposal", _proposalID, _transactionData)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 _proposalID, bytes _transactionData) returns(bool success)
func (_DAO *DAOSession) ExecuteProposal(_proposalID *big.Int, _transactionData []byte) (*types.Transaction, error) {
	return _DAO.Contract.ExecuteProposal(&_DAO.TransactOpts, _proposalID, _transactionData)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x237e9492.
//
// Solidity: function executeProposal(uint256 _proposalID, bytes _transactionData) returns(bool success)
func (_DAO *DAOTransactorSession) ExecuteProposal(_proposalID *big.Int, _transactionData []byte) (*types.Transaction, error) {
	return _DAO.Contract.ExecuteProposal(&_DAO.TransactOpts, _proposalID, _transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_DAO *DAOTransactor) MakeProposal(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _DAO.contract.Transact(opts, "makeProposal", _recipient, _amount, _description, _transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_DAO *DAOSession) MakeProposal(_recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _DAO.Contract.MakeProposal(&_DAO.TransactOpts, _recipient, _amount, _description, _transactionData)
}

// MakeProposal is a paid mutator transaction binding the contract method 0x909d02ad.
//
// Solidity: function makeProposal(address _recipient, uint256 _amount, string _description, bytes _transactionData) returns(uint256 proposalID)
func (_DAO *DAOTransactorSession) MakeProposal(_recipient common.Address, _amount *big.Int, _description string, _transactionData []byte) (*types.Transaction, error) {
	return _DAO.Contract.MakeProposal(&_DAO.TransactOpts, _recipient, _amount, _description, _transactionData)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_DAO *DAOTransactor) UnVote(opts *bind.TransactOpts, _proposalID *big.Int) (*types.Transaction, error) {
	return _DAO.contract.Transact(opts, "unVote", _proposalID)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_DAO *DAOSession) UnVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _DAO.Contract.UnVote(&_DAO.TransactOpts, _proposalID)
}

// UnVote is a paid mutator transaction binding the contract method 0x9fd6adb3.
//
// Solidity: function unVote(uint256 _proposalID) returns()
func (_DAO *DAOTransactorSession) UnVote(_proposalID *big.Int) (*types.Transaction, error) {
	return _DAO.Contract.UnVote(&_DAO.TransactOpts, _proposalID)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_DAO *DAOTransactor) Vote(opts *bind.TransactOpts, _proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _DAO.contract.Transact(opts, "vote", _proposalID, _inSupport)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_DAO *DAOSession) Vote(_proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _DAO.Contract.Vote(&_DAO.TransactOpts, _proposalID, _inSupport)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(uint256 _proposalID, bool _inSupport) returns()
func (_DAO *DAOTransactorSession) Vote(_proposalID *big.Int, _inSupport bool) (*types.Transaction, error) {
	return _DAO.Contract.Vote(&_DAO.TransactOpts, _proposalID, _inSupport)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAO *DAOTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAO.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAO *DAOSession) Receive() (*types.Transaction, error) {
	return _DAO.Contract.Receive(&_DAO.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DAO *DAOTransactorSession) Receive() (*types.Transaction, error) {
	return _DAO.Contract.Receive(&_DAO.TransactOpts)
}

// DAOProposalAddedIterator is returned from FilterProposalAdded and is used to iterate over the raw logs and unpacked data for ProposalAdded events raised by the DAO contract.
type DAOProposalAddedIterator struct {
	Event *DAOProposalAdded // Event containing the contract specifics and raw log

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
func (it *DAOProposalAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOProposalAdded)
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
		it.Event = new(DAOProposalAdded)
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
func (it *DAOProposalAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOProposalAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOProposalAdded represents a ProposalAdded event raised by the DAO contract.
type DAOProposalAdded struct {
	ProposalID  *big.Int
	Recipient   common.Address
	Amount      *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalAdded is a free log retrieval operation binding the contract event 0x646fec02522b41e7125cfc859a64fd4f4cefd5dc3b6237ca0abe251ded1fa881.
//
// Solidity: event ProposalAdded(uint256 indexed proposalID, address recipient, uint256 amount, string description)
func (_DAO *DAOFilterer) FilterProposalAdded(opts *bind.FilterOpts, proposalID []*big.Int) (*DAOProposalAddedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _DAO.contract.FilterLogs(opts, "ProposalAdded", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &DAOProposalAddedIterator{contract: _DAO.contract, event: "ProposalAdded", logs: logs, sub: sub}, nil
}

// WatchProposalAdded is a free log subscription operation binding the contract event 0x646fec02522b41e7125cfc859a64fd4f4cefd5dc3b6237ca0abe251ded1fa881.
//
// Solidity: event ProposalAdded(uint256 indexed proposalID, address recipient, uint256 amount, string description)
func (_DAO *DAOFilterer) WatchProposalAdded(opts *bind.WatchOpts, sink chan<- *DAOProposalAdded, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _DAO.contract.WatchLogs(opts, "ProposalAdded", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOProposalAdded)
				if err := _DAO.contract.UnpackLog(event, "ProposalAdded", log); err != nil {
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
func (_DAO *DAOFilterer) ParseProposalAdded(log types.Log) (*DAOProposalAdded, error) {
	event := new(DAOProposalAdded)
	if err := _DAO.contract.UnpackLog(event, "ProposalAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the DAO contract.
type DAOProposalExecutedIterator struct {
	Event *DAOProposalExecuted // Event containing the contract specifics and raw log

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
func (it *DAOProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOProposalExecuted)
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
		it.Event = new(DAOProposalExecuted)
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
func (it *DAOProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOProposalExecuted represents a ProposalExecuted event raised by the DAO contract.
type DAOProposalExecuted struct {
	ProposalID *big.Int
	Result     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x948f4a9cd986f1118c3fbd459f7a22b23c0693e1ca3ef06a6a8be5aa7d39cc03.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, bool result)
func (_DAO *DAOFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalID []*big.Int) (*DAOProposalExecutedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _DAO.contract.FilterLogs(opts, "ProposalExecuted", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return &DAOProposalExecutedIterator{contract: _DAO.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x948f4a9cd986f1118c3fbd459f7a22b23c0693e1ca3ef06a6a8be5aa7d39cc03.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalID, bool result)
func (_DAO *DAOFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *DAOProposalExecuted, proposalID []*big.Int) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	logs, sub, err := _DAO.contract.WatchLogs(opts, "ProposalExecuted", proposalIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOProposalExecuted)
				if err := _DAO.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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
func (_DAO *DAOFilterer) ParseProposalExecuted(log types.Log) (*DAOProposalExecuted, error) {
	event := new(DAOProposalExecuted)
	if err := _DAO.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DAOVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the DAO contract.
type DAOVotedIterator struct {
	Event *DAOVoted // Event containing the contract specifics and raw log

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
func (it *DAOVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DAOVoted)
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
		it.Event = new(DAOVoted)
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
func (it *DAOVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DAOVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DAOVoted represents a Voted event raised by the DAO contract.
type DAOVoted struct {
	ProposalID *big.Int
	InSupport  bool
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed proposalID, bool inSupport, address indexed voter)
func (_DAO *DAOFilterer) FilterVoted(opts *bind.FilterOpts, proposalID []*big.Int, voter []common.Address) (*DAOVotedIterator, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DAO.contract.FilterLogs(opts, "Voted", proposalIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &DAOVotedIterator{contract: _DAO.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: event Voted(uint256 indexed proposalID, bool inSupport, address indexed voter)
func (_DAO *DAOFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *DAOVoted, proposalID []*big.Int, voter []common.Address) (event.Subscription, error) {

	var proposalIDRule []interface{}
	for _, proposalIDItem := range proposalID {
		proposalIDRule = append(proposalIDRule, proposalIDItem)
	}

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _DAO.contract.WatchLogs(opts, "Voted", proposalIDRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DAOVoted)
				if err := _DAO.contract.UnpackLog(event, "Voted", log); err != nil {
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
func (_DAO *DAOFilterer) ParseVoted(log types.Log) (*DAOVoted, error) {
	event := new(DAOVoted)
	if err := _DAO.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
