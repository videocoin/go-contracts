// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RegistryRecord is an auto generated low-level Go binding around an user-defined struct.
type RegistryRecord struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Index       *big.Int
}

// RemoteBridgeABI is the input ABI used to generate the binding from.
const RemoteBridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"VersionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"versions\",\"type\":\"string\"}],\"name\":\"VersionUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"record\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structRegistry.Record\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"versions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RemoteBridgeBin is the compiled bytecode used for deploying new contracts.
var RemoteBridgeBin = "0x608060405260006100146100b760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100bf565b600033905090565b611d34806100ce6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80635adb1cb1116100665780635adb1cb114610148578063715018a6146101645780638da5cb5b1461016e5780638f32d59b1461018c578063f2fde38b146101aa57610093565b806301e647251461009857806344590a7e146100cc578063460e3ca2146100e8578063470bb62b14610118575b600080fd5b6100b260048036036100ad919081019061146e565b6101c6565b6040516100c39594939291906119ad565b60405180910390f35b6100e660048036036100e191908101906114c0565b610359565b005b61010260048036036100fd91908101906115ab565b6105a2565b60405161010f9190611a0e565b60405180910390f35b610132600480360361012d91908101906114c0565b61067e565b60405161013f9190611b07565b60405180910390f35b610162600480360361015d919081019061152c565b61096b565b005b61016c610e7f565b005b610176610f85565b6040516101839190611977565b60405180910390f35b610194610fae565b6040516101a19190611992565b60405180910390f35b6101c460048036036101bf9190810190611445565b61100c565b005b60016020528060005260406000206000915090508060000160009054906101000a900460ff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102855780601f1061025a57610100808354040283529160200191610285565b820191906000526020600020905b81548152906001019060200180831161026857829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103235780601f106102f857610100808354040283529160200191610323565b820191906000526020600020905b81548152906001019060200180831161030657829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905085565b610361610fae565b6103a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039790611ae7565b60405180910390fd5b6000825114156103e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103dc90611aa7565b60405180910390fd5b60008151141561042a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042190611a87565b60405180910390fd5b600060028383604051602001610441929190611a30565b60405160208183030381529060405260405161045d9190611949565b602060405180830381855afa15801561047a573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525061049d9190810190611497565b90506001600082815260200190815260200160002060000160009054906101000a900460ff161561059d5760006001600083815260200190815260200160002060040154905060016000838152602001908152602001600020600080820160006101000a81549060ff021916905560018201600061051b9190611195565b60028201600061052b9190611195565b6003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160009055505060028460405161056c9190611960565b9081526020016040518091039020818154811061058557fe5b90600052602060002001600061059b9190611195565b505b505050565b60028280516020810182018051848252602083016020850120818352809550505050505081815481106105d157fe5b90600052602060002001600091509150508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106765780601f1061064b57610100808354040283529160200191610676565b820191906000526020600020905b81548152906001019060200180831161065957829003601f168201915b505050505081565b6106866111dd565b6000835114156106cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c290611aa7565b60405180910390fd5b600082511415610710576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070790611a87565b60405180910390fd5b600060028484604051602001610727929190611a30565b6040516020818303038152906040526040516107439190611949565b602060405180830381855afa158015610760573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506107839190810190611497565b9050600160008281526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff16151515158152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108575780601f1061082c57610100808354040283529160200191610857565b820191906000526020600020905b81548152906001019060200180831161083a57829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108f95780601f106108ce576101008083540402835291602001916108f9565b820191906000526020600020905b8154815290600101906020018083116108dc57829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152505091505092915050565b610973610fae565b6109b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a990611ae7565b60405180910390fd5b6000835114156109f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ee90611aa7565b60405180910390fd5b600082511415610a3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3390611a87565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610aac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa390611ac7565b60405180910390fd5b600060028484604051602001610ac3929190611a30565b604051602081830303815290604052604051610adf9190611949565b602060405180830381855afa158015610afc573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250610b1f9190810190611497565b90506001600082815260200190815260200160002060000160009054906101000a900460ff16610d095760006001600286604051610b5d9190611960565b9081526020016040518091039020859080600181540180825580915050906001820390600052602060002001600090919290919091509080519060200190610ba6929190611224565b500390506040518060a001604052806001151581526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff168152602001828152506001600084815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190610c399291906112a4565b506040820151816002019080519060200190610c569291906112a4565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015590505083604051610cb99190611960565b604051809103902085604051610ccf9190611960565b60405180910390207f600be21ea77e400a162a88fb0bef7e7e4e0d1b28da89d643862828bc7a630cf660405160405180910390a350610e79565b6040518060a001604052806001151581526020018581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff16815260200160016000848152602001908152602001600020600401548152506001600083815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190610dae9291906112a4565b506040820151816002019080519060200190610dcb9291906112a4565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015590505082604051610e2e9190611960565b604051809103902084604051610e449190611960565b60405180910390207f22efc5f993dce37856b77dd72d7d7661032380c9728c4133f3c071c591bc6ca760405160405180910390a35b50505050565b610e87610fae565b610ec6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ebd90611ae7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610ff061105f565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b611014610fae565b611053576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104a90611ae7565b60405180910390fd5b61105c81611067565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156110d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ce90611a67565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b50805460018160011615610100020316600290046000825580601f106111bb57506111da565b601f0160209004906000526020600020908101906111d99190611324565b5b50565b6040518060a001604052806000151581526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061126557805160ff1916838001178555611293565b82800160010185558215611293579182015b82811115611292578251825591602001919060010190611277565b5b5090506112a09190611324565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112e557805160ff1916838001178555611313565b82800160010185558215611313579182015b828111156113125782518255916020019190600101906112f7565b5b5090506113209190611324565b5090565b61134691905b8082111561134257600081600090555060010161132a565b5090565b90565b60008135905061135881611cac565b92915050565b60008135905061136d81611cc3565b92915050565b60008151905061138281611cc3565b92915050565b600082601f83011261139957600080fd5b81356113ac6113a782611b56565b611b29565b915080825260208301602083018583830111156113c857600080fd5b6113d3838284611c59565b50505092915050565b600082601f8301126113ed57600080fd5b81356114006113fb82611b82565b611b29565b9150808252602083016020830185838301111561141c57600080fd5b611427838284611c59565b50505092915050565b60008135905061143f81611cda565b92915050565b60006020828403121561145757600080fd5b600061146584828501611349565b91505092915050565b60006020828403121561148057600080fd5b600061148e8482850161135e565b91505092915050565b6000602082840312156114a957600080fd5b60006114b784828501611373565b91505092915050565b600080604083850312156114d357600080fd5b600083013567ffffffffffffffff8111156114ed57600080fd5b6114f9858286016113dc565b925050602083013567ffffffffffffffff81111561151657600080fd5b611522858286016113dc565b9150509250929050565b60008060006060848603121561154157600080fd5b600084013567ffffffffffffffff81111561155b57600080fd5b611567868287016113dc565b935050602084013567ffffffffffffffff81111561158457600080fd5b611590868287016113dc565b92505060406115a186828701611349565b9150509250925092565b600080604083850312156115be57600080fd5b600083013567ffffffffffffffff8111156115d857600080fd5b6115e485828601611388565b92505060206115f585828601611430565b9150509250929050565b61160881611c07565b82525050565b61161781611c07565b82525050565b61162681611c19565b82525050565b61163581611c19565b82525050565b600061164682611bae565b6116508185611bcf565b9350611660818560208601611c68565b80840191505092915050565b600061167782611bc4565b6116818185611beb565b9350611691818560208601611c68565b61169a81611c9b565b840191505092915050565b60006116b082611bc4565b6116ba8185611bfc565b93506116ca818560208601611c68565b80840191505092915050565b60006116e182611bb9565b6116eb8185611bda565b93506116fb818560208601611c68565b61170481611c9b565b840191505092915050565b600061171a82611bb9565b6117248185611beb565b9350611734818560208601611c68565b61173d81611c9b565b840191505092915050565b6000611755602683611beb565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006117bb601d83611beb565b91507f52656769737472793a2076657273696f6e2069732072657175697265640000006000830152602082019050919050565b60006117fb601a83611beb565b91507f52656769737472793a206e616d652069732072657175697265640000000000006000830152602082019050919050565b600061183b601f83611beb565b91507f52656769737472793a206163636f756e742063616e2774206265207a65726f006000830152602082019050919050565b600061187b602083611beb565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b600060a0830160008301516118c6600086018261161d565b50602083015184820360208601526118de82826116d6565b915050604083015184820360408601526118f882826116d6565b915050606083015161190d60608601826115ff565b506080830151611920608086018261192b565b508091505092915050565b61193481611c4f565b82525050565b61194381611c4f565b82525050565b6000611955828461163b565b915081905092915050565b600061196c82846116a5565b915081905092915050565b600060208201905061198c600083018461160e565b92915050565b60006020820190506119a7600083018461162c565b92915050565b600060a0820190506119c2600083018861162c565b81810360208301526119d4818761170f565b905081810360408301526119e8818661170f565b90506119f7606083018561160e565b611a04608083018461193a565b9695505050505050565b60006020820190508181036000830152611a28818461170f565b905092915050565b60006040820190508181036000830152611a4a818561166c565b90508181036020830152611a5e818461166c565b90509392505050565b60006020820190508181036000830152611a8081611748565b9050919050565b60006020820190508181036000830152611aa0816117ae565b9050919050565b60006020820190508181036000830152611ac0816117ee565b9050919050565b60006020820190508181036000830152611ae08161182e565b9050919050565b60006020820190508181036000830152611b008161186e565b9050919050565b60006020820190508181036000830152611b2181846118ae565b905092915050565b6000604051905081810181811067ffffffffffffffff82111715611b4c57600080fd5b8060405250919050565b600067ffffffffffffffff821115611b6d57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115611b9957600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611c1282611c2f565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611c86578082015181840152602081019050611c6b565b83811115611c95576000848401525b50505050565b6000601f19601f8301169050919050565b611cb581611c07565b8114611cc057600080fd5b50565b611ccc81611c25565b8114611cd757600080fd5b50565b611ce381611c4f565b8114611cee57600080fd5b5056fea365627a7a723158201d6622d25e6bfb8e5d96587804fe72415824d926732c6cddac4acb91c01ad3c06c6578706572696d656e74616cf564736f6c634300050d0040"

// DeployRemoteBridge deploys a new Ethereum contract, binding an instance of RemoteBridge to it.
func DeployRemoteBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RemoteBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(RemoteBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RemoteBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RemoteBridge{RemoteBridgeCaller: RemoteBridgeCaller{contract: contract}, RemoteBridgeTransactor: RemoteBridgeTransactor{contract: contract}, RemoteBridgeFilterer: RemoteBridgeFilterer{contract: contract}}, nil
}

// RemoteBridge is an auto generated Go binding around an Ethereum contract.
type RemoteBridge struct {
	RemoteBridgeCaller     // Read-only binding to the contract
	RemoteBridgeTransactor // Write-only binding to the contract
	RemoteBridgeFilterer   // Log filterer for contract events
}

// RemoteBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RemoteBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RemoteBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RemoteBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RemoteBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RemoteBridgeSession struct {
	Contract     *RemoteBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RemoteBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RemoteBridgeCallerSession struct {
	Contract *RemoteBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RemoteBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RemoteBridgeTransactorSession struct {
	Contract     *RemoteBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RemoteBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RemoteBridgeRaw struct {
	Contract *RemoteBridge // Generic contract binding to access the raw methods on
}

// RemoteBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RemoteBridgeCallerRaw struct {
	Contract *RemoteBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// RemoteBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RemoteBridgeTransactorRaw struct {
	Contract *RemoteBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRemoteBridge creates a new instance of RemoteBridge, bound to a specific deployed contract.
func NewRemoteBridge(address common.Address, backend bind.ContractBackend) (*RemoteBridge, error) {
	contract, err := bindRemoteBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RemoteBridge{RemoteBridgeCaller: RemoteBridgeCaller{contract: contract}, RemoteBridgeTransactor: RemoteBridgeTransactor{contract: contract}, RemoteBridgeFilterer: RemoteBridgeFilterer{contract: contract}}, nil
}

// NewRemoteBridgeCaller creates a new read-only instance of RemoteBridge, bound to a specific deployed contract.
func NewRemoteBridgeCaller(address common.Address, caller bind.ContractCaller) (*RemoteBridgeCaller, error) {
	contract, err := bindRemoteBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeCaller{contract: contract}, nil
}

// NewRemoteBridgeTransactor creates a new write-only instance of RemoteBridge, bound to a specific deployed contract.
func NewRemoteBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*RemoteBridgeTransactor, error) {
	contract, err := bindRemoteBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeTransactor{contract: contract}, nil
}

// NewRemoteBridgeFilterer creates a new log filterer instance of RemoteBridge, bound to a specific deployed contract.
func NewRemoteBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*RemoteBridgeFilterer, error) {
	contract, err := bindRemoteBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeFilterer{contract: contract}, nil
}

// bindRemoteBridge binds a generic wrapper to an already deployed contract.
func bindRemoteBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RemoteBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RemoteBridge *RemoteBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RemoteBridge.Contract.RemoteBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RemoteBridge *RemoteBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RemoteBridge.Contract.RemoteBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RemoteBridge *RemoteBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RemoteBridge.Contract.RemoteBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RemoteBridge *RemoteBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RemoteBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RemoteBridge *RemoteBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RemoteBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RemoteBridge *RemoteBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RemoteBridge.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RemoteBridge *RemoteBridgeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RemoteBridge.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RemoteBridge *RemoteBridgeSession) IsOwner() (bool, error) {
	return _RemoteBridge.Contract.IsOwner(&_RemoteBridge.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_RemoteBridge *RemoteBridgeCallerSession) IsOwner() (bool, error) {
	return _RemoteBridge.Contract.IsOwner(&_RemoteBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RemoteBridge *RemoteBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RemoteBridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RemoteBridge *RemoteBridgeSession) Owner() (common.Address, error) {
	return _RemoteBridge.Contract.Owner(&_RemoteBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_RemoteBridge *RemoteBridgeCallerSession) Owner() (common.Address, error) {
	return _RemoteBridge.Contract.Owner(&_RemoteBridge.CallOpts)
}

// Record is a free data retrieval call binding the contract method 0x470bb62b.
//
// Solidity: function record(string name, string version) constant returns(RegistryRecord)
func (_RemoteBridge *RemoteBridgeCaller) Record(opts *bind.CallOpts, name string, version string) (RegistryRecord, error) {
	var (
		ret0 = new(RegistryRecord)
	)
	out := ret0
	err := _RemoteBridge.contract.Call(opts, out, "record", name, version)
	return *ret0, err
}

// Record is a free data retrieval call binding the contract method 0x470bb62b.
//
// Solidity: function record(string name, string version) constant returns(RegistryRecord)
func (_RemoteBridge *RemoteBridgeSession) Record(name string, version string) (RegistryRecord, error) {
	return _RemoteBridge.Contract.Record(&_RemoteBridge.CallOpts, name, version)
}

// Record is a free data retrieval call binding the contract method 0x470bb62b.
//
// Solidity: function record(string name, string version) constant returns(RegistryRecord)
func (_RemoteBridge *RemoteBridgeCallerSession) Record(name string, version string) (RegistryRecord, error) {
	return _RemoteBridge.Contract.Record(&_RemoteBridge.CallOpts, name, version)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, uint256 index)
func (_RemoteBridge *RemoteBridgeCaller) Records(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Index       *big.Int
}, error) {
	ret := new(struct {
		Initialized bool
		Name        string
		Version     string
		Account     common.Address
		Index       *big.Int
	})
	out := ret
	err := _RemoteBridge.contract.Call(opts, out, "records", arg0)
	return *ret, err
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, uint256 index)
func (_RemoteBridge *RemoteBridgeSession) Records(arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Index       *big.Int
}, error) {
	return _RemoteBridge.Contract.Records(&_RemoteBridge.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, uint256 index)
func (_RemoteBridge *RemoteBridgeCallerSession) Records(arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Index       *big.Int
}, error) {
	return _RemoteBridge.Contract.Records(&_RemoteBridge.CallOpts, arg0)
}

// Versions is a free data retrieval call binding the contract method 0x460e3ca2.
//
// Solidity: function versions(string , uint256 ) constant returns(string)
func (_RemoteBridge *RemoteBridgeCaller) Versions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RemoteBridge.contract.Call(opts, out, "versions", arg0, arg1)
	return *ret0, err
}

// Versions is a free data retrieval call binding the contract method 0x460e3ca2.
//
// Solidity: function versions(string , uint256 ) constant returns(string)
func (_RemoteBridge *RemoteBridgeSession) Versions(arg0 string, arg1 *big.Int) (string, error) {
	return _RemoteBridge.Contract.Versions(&_RemoteBridge.CallOpts, arg0, arg1)
}

// Versions is a free data retrieval call binding the contract method 0x460e3ca2.
//
// Solidity: function versions(string , uint256 ) constant returns(string)
func (_RemoteBridge *RemoteBridgeCallerSession) Versions(arg0 string, arg1 *big.Int) (string, error) {
	return _RemoteBridge.Contract.Versions(&_RemoteBridge.CallOpts, arg0, arg1)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string name, string version) returns()
func (_RemoteBridge *RemoteBridgeTransactor) Remove(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "remove", name, version)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string name, string version) returns()
func (_RemoteBridge *RemoteBridgeSession) Remove(name string, version string) (*types.Transaction, error) {
	return _RemoteBridge.Contract.Remove(&_RemoteBridge.TransactOpts, name, version)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string name, string version) returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) Remove(name string, version string) (*types.Transaction, error) {
	return _RemoteBridge.Contract.Remove(&_RemoteBridge.TransactOpts, name, version)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RemoteBridge *RemoteBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RemoteBridge *RemoteBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _RemoteBridge.Contract.RenounceOwnership(&_RemoteBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RemoteBridge.Contract.RenounceOwnership(&_RemoteBridge.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RemoteBridge *RemoteBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RemoteBridge *RemoteBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RemoteBridge.Contract.TransferOwnership(&_RemoteBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RemoteBridge.Contract.TransferOwnership(&_RemoteBridge.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x5adb1cb1.
//
// Solidity: function update(string name, string version, address account) returns()
func (_RemoteBridge *RemoteBridgeTransactor) Update(opts *bind.TransactOpts, name string, version string, account common.Address) (*types.Transaction, error) {
	return _RemoteBridge.contract.Transact(opts, "update", name, version, account)
}

// Update is a paid mutator transaction binding the contract method 0x5adb1cb1.
//
// Solidity: function update(string name, string version, address account) returns()
func (_RemoteBridge *RemoteBridgeSession) Update(name string, version string, account common.Address) (*types.Transaction, error) {
	return _RemoteBridge.Contract.Update(&_RemoteBridge.TransactOpts, name, version, account)
}

// Update is a paid mutator transaction binding the contract method 0x5adb1cb1.
//
// Solidity: function update(string name, string version, address account) returns()
func (_RemoteBridge *RemoteBridgeTransactorSession) Update(name string, version string, account common.Address) (*types.Transaction, error) {
	return _RemoteBridge.Contract.Update(&_RemoteBridge.TransactOpts, name, version, account)
}

// RemoteBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RemoteBridge contract.
type RemoteBridgeOwnershipTransferredIterator struct {
	Event *RemoteBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RemoteBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RemoteBridgeOwnershipTransferred)
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
		it.Event = new(RemoteBridgeOwnershipTransferred)
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
func (it *RemoteBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RemoteBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RemoteBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the RemoteBridge contract.
type RemoteBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RemoteBridge *RemoteBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RemoteBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RemoteBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeOwnershipTransferredIterator{contract: _RemoteBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RemoteBridge *RemoteBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RemoteBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RemoteBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RemoteBridgeOwnershipTransferred)
				if err := _RemoteBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RemoteBridge *RemoteBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*RemoteBridgeOwnershipTransferred, error) {
	event := new(RemoteBridgeOwnershipTransferred)
	if err := _RemoteBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RemoteBridgeVersionAddedIterator is returned from FilterVersionAdded and is used to iterate over the raw logs and unpacked data for VersionAdded events raised by the RemoteBridge contract.
type RemoteBridgeVersionAddedIterator struct {
	Event *RemoteBridgeVersionAdded // Event containing the contract specifics and raw log

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
func (it *RemoteBridgeVersionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RemoteBridgeVersionAdded)
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
		it.Event = new(RemoteBridgeVersionAdded)
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
func (it *RemoteBridgeVersionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RemoteBridgeVersionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RemoteBridgeVersionAdded represents a VersionAdded event raised by the RemoteBridge contract.
type RemoteBridgeVersionAdded struct {
	Name    common.Hash
	Version common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVersionAdded is a free log retrieval operation binding the contract event 0x600be21ea77e400a162a88fb0bef7e7e4e0d1b28da89d643862828bc7a630cf6.
//
// Solidity: event VersionAdded(string indexed name, string indexed version)
func (_RemoteBridge *RemoteBridgeFilterer) FilterVersionAdded(opts *bind.FilterOpts, name []string, version []string) (*RemoteBridgeVersionAddedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _RemoteBridge.contract.FilterLogs(opts, "VersionAdded", nameRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeVersionAddedIterator{contract: _RemoteBridge.contract, event: "VersionAdded", logs: logs, sub: sub}, nil
}

// WatchVersionAdded is a free log subscription operation binding the contract event 0x600be21ea77e400a162a88fb0bef7e7e4e0d1b28da89d643862828bc7a630cf6.
//
// Solidity: event VersionAdded(string indexed name, string indexed version)
func (_RemoteBridge *RemoteBridgeFilterer) WatchVersionAdded(opts *bind.WatchOpts, sink chan<- *RemoteBridgeVersionAdded, name []string, version []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _RemoteBridge.contract.WatchLogs(opts, "VersionAdded", nameRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RemoteBridgeVersionAdded)
				if err := _RemoteBridge.contract.UnpackLog(event, "VersionAdded", log); err != nil {
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

// ParseVersionAdded is a log parse operation binding the contract event 0x600be21ea77e400a162a88fb0bef7e7e4e0d1b28da89d643862828bc7a630cf6.
//
// Solidity: event VersionAdded(string indexed name, string indexed version)
func (_RemoteBridge *RemoteBridgeFilterer) ParseVersionAdded(log types.Log) (*RemoteBridgeVersionAdded, error) {
	event := new(RemoteBridgeVersionAdded)
	if err := _RemoteBridge.contract.UnpackLog(event, "VersionAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RemoteBridgeVersionUpdatedIterator is returned from FilterVersionUpdated and is used to iterate over the raw logs and unpacked data for VersionUpdated events raised by the RemoteBridge contract.
type RemoteBridgeVersionUpdatedIterator struct {
	Event *RemoteBridgeVersionUpdated // Event containing the contract specifics and raw log

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
func (it *RemoteBridgeVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RemoteBridgeVersionUpdated)
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
		it.Event = new(RemoteBridgeVersionUpdated)
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
func (it *RemoteBridgeVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RemoteBridgeVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RemoteBridgeVersionUpdated represents a VersionUpdated event raised by the RemoteBridge contract.
type RemoteBridgeVersionUpdated struct {
	Name     common.Hash
	Versions common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVersionUpdated is a free log retrieval operation binding the contract event 0x22efc5f993dce37856b77dd72d7d7661032380c9728c4133f3c071c591bc6ca7.
//
// Solidity: event VersionUpdated(string indexed name, string indexed versions)
func (_RemoteBridge *RemoteBridgeFilterer) FilterVersionUpdated(opts *bind.FilterOpts, name []string, versions []string) (*RemoteBridgeVersionUpdatedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionsRule []interface{}
	for _, versionsItem := range versions {
		versionsRule = append(versionsRule, versionsItem)
	}

	logs, sub, err := _RemoteBridge.contract.FilterLogs(opts, "VersionUpdated", nameRule, versionsRule)
	if err != nil {
		return nil, err
	}
	return &RemoteBridgeVersionUpdatedIterator{contract: _RemoteBridge.contract, event: "VersionUpdated", logs: logs, sub: sub}, nil
}

// WatchVersionUpdated is a free log subscription operation binding the contract event 0x22efc5f993dce37856b77dd72d7d7661032380c9728c4133f3c071c591bc6ca7.
//
// Solidity: event VersionUpdated(string indexed name, string indexed versions)
func (_RemoteBridge *RemoteBridgeFilterer) WatchVersionUpdated(opts *bind.WatchOpts, sink chan<- *RemoteBridgeVersionUpdated, name []string, versions []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionsRule []interface{}
	for _, versionsItem := range versions {
		versionsRule = append(versionsRule, versionsItem)
	}

	logs, sub, err := _RemoteBridge.contract.WatchLogs(opts, "VersionUpdated", nameRule, versionsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RemoteBridgeVersionUpdated)
				if err := _RemoteBridge.contract.UnpackLog(event, "VersionUpdated", log); err != nil {
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

// ParseVersionUpdated is a log parse operation binding the contract event 0x22efc5f993dce37856b77dd72d7d7661032380c9728c4133f3c071c591bc6ca7.
//
// Solidity: event VersionUpdated(string indexed name, string indexed versions)
func (_RemoteBridge *RemoteBridgeFilterer) ParseVersionUpdated(log types.Log) (*RemoteBridgeVersionUpdated, error) {
	event := new(RemoteBridgeVersionUpdated)
	if err := _RemoteBridge.contract.UnpackLog(event, "VersionUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
