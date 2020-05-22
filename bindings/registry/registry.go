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
	Owner       common.Address
	Index       *big.Int
}

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"VersionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"versions\",\"type\":\"string\"}],\"name\":\"VersionUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"record\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"internalType\":\"structRegistry.Record\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"versions\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
var RegistryBin = "0x608060405260006100146100b760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100bf565b600033905090565b611ef1806100ce6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063715018a611610066578063715018a6146101495780638da5cb5b146101535780638f32d59b14610171578063b03205b01461018f578063f2fde38b146101ab57610093565b806301e647251461009857806344590a7e146100cd578063460e3ca2146100e9578063470bb62b14610119575b600080fd5b6100b260048036036100ad91908101906115f6565b6101c7565b6040516100c496959493929190611b5c565b60405180910390f35b6100e760048036036100e29190810190611648565b610380565b005b61010360048036036100fe9190810190611747565b6105f0565b6040516101109190611bcb565b60405180910390f35b610133600480360361012e9190810190611648565b6106cc565b6040516101409190611cc4565b60405180910390f35b610151610a0f565b005b61015b610b15565b6040516101689190611b26565b60405180910390f35b610179610b3e565b6040516101869190611b41565b60405180910390f35b6101a960048036036101a491908101906116b4565b610b9c565b005b6101c560048036036101c091908101906115cd565b611177565b005b60016020528060005260406000206000915090508060000160009054906101000a900460ff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102865780601f1061025b57610100808354040283529160200191610286565b820191906000526020600020905b81548152906001019060200180831161026957829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103245780601f106102f957610100808354040283529160200191610324565b820191906000526020600020905b81548152906001019060200180831161030757829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050154905086565b610388610b3e565b6103c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103be90611ca4565b60405180910390fd5b60008251141561040c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040390611c64565b60405180910390fd5b600081511415610451576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044890611c44565b60405180910390fd5b600060028383604051602001610468929190611bed565b6040516020818303038152906040526040516104849190611af8565b602060405180830381855afa1580156104a1573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506104c4919081019061161f565b90506001600082815260200190815260200160002060000160009054906101000a900460ff16156105eb5760006001600083815260200190815260200160002060050154905060016000838152602001908152602001600020600080820160006101000a81549060ff02191690556001820160006105429190611300565b6002820160006105529190611300565b6003820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556004820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600582016000905550506002846040516105ba9190611b0f565b908152602001604051809103902081815481106105d357fe5b9060005260206000200160006105e99190611300565b505b505050565b600282805160208101820180518482526020830160208501208183528095505050505050818154811061061f57fe5b90600052602060002001600091509150508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106c45780601f10610699576101008083540402835291602001916106c4565b820191906000526020600020905b8154815290600101906020018083116106a757829003601f168201915b505050505081565b6106d4611348565b600083511415610719576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071090611c64565b60405180910390fd5b60008251141561075e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075590611c44565b60405180910390fd5b600060028484604051602001610775929190611bed565b6040516020818303038152906040526040516107919190611af8565b602060405180830381855afa1580156107ae573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052506107d1919081019061161f565b9050600160008281526020019081526020016000206040518060c00160405290816000820160009054906101000a900460ff16151515158152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108a55780601f1061087a576101008083540402835291602001916108a5565b820191906000526020600020905b81548152906001019060200180831161088857829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109475780601f1061091c57610100808354040283529160200191610947565b820191906000526020600020905b81548152906001019060200180831161092a57829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160058201548152505091505092915050565b610a17610b3e565b610a56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4d90611ca4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610b806111ca565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b610ba4610b3e565b610be3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bda90611ca4565b60405180910390fd5b600084511415610c28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1f90611c64565b60405180910390fd5b600083511415610c6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c6490611c44565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610cdd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd490611c84565b60405180910390fd5b600060028585604051602001610cf4929190611bed565b604051602081830303815290604052604051610d109190611af8565b602060405180830381855afa158015610d2d573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250610d50919081019061161f565b90506001600082815260200190815260200160002060000160009054906101000a900460ff16610f9d5760006001600287604051610d8e9190611b0f565b9081526020016040518091039020869080600181540180825580915050906001820390600052602060002001600090919290919091509080519060200190610dd79291906113ac565b500390506040518060c001604052806001151581526020018781526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff168152602001828152506001600084815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190610e8692919061142c565b506040820151816002019080519060200190610ea392919061142c565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816005015590505084604051610f4d9190611b0f565b604051809103902086604051610f639190611b0f565b60405180910390207f600be21ea77e400a162a88fb0bef7e7e4e0d1b28da89d643862828bc7a630cf660405160405180910390a350611170565b6040518060c001604052806001151581526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff16815260200160016000848152602001908152602001600020600501548152506001600083815260200190815260200160002060008201518160000160006101000a81548160ff021916908315150217905550602082015181600101908051906020019061105e92919061142c565b50604082015181600201908051906020019061107b92919061142c565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050155905050836040516111259190611b0f565b60405180910390208560405161113b9190611b0f565b60405180910390207f22efc5f993dce37856b77dd72d7d7661032380c9728c4133f3c071c591bc6ca760405160405180910390a35b5050505050565b61117f610b3e565b6111be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111b590611ca4565b60405180910390fd5b6111c7816111d2565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611242576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123990611c24565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b50805460018160011615610100020316600290046000825580601f106113265750611345565b601f01602090049060005260206000209081019061134491906114ac565b5b50565b6040518060c001604052806000151581526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113ed57805160ff191683800117855561141b565b8280016001018555821561141b579182015b8281111561141a5782518255916020019190600101906113ff565b5b50905061142891906114ac565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061146d57805160ff191683800117855561149b565b8280016001018555821561149b579182015b8281111561149a57825182559160200191906001019061147f565b5b5090506114a891906114ac565b5090565b6114ce91905b808211156114ca5760008160009055506001016114b2565b5090565b90565b6000813590506114e081611e69565b92915050565b6000813590506114f581611e80565b92915050565b60008151905061150a81611e80565b92915050565b600082601f83011261152157600080fd5b813561153461152f82611d13565b611ce6565b9150808252602083016020830185838301111561155057600080fd5b61155b838284611e16565b50505092915050565b600082601f83011261157557600080fd5b813561158861158382611d3f565b611ce6565b915080825260208301602083018583830111156115a457600080fd5b6115af838284611e16565b50505092915050565b6000813590506115c781611e97565b92915050565b6000602082840312156115df57600080fd5b60006115ed848285016114d1565b91505092915050565b60006020828403121561160857600080fd5b6000611616848285016114e6565b91505092915050565b60006020828403121561163157600080fd5b600061163f848285016114fb565b91505092915050565b6000806040838503121561165b57600080fd5b600083013567ffffffffffffffff81111561167557600080fd5b61168185828601611564565b925050602083013567ffffffffffffffff81111561169e57600080fd5b6116aa85828601611564565b9150509250929050565b600080600080608085870312156116ca57600080fd5b600085013567ffffffffffffffff8111156116e457600080fd5b6116f087828801611564565b945050602085013567ffffffffffffffff81111561170d57600080fd5b61171987828801611564565b935050604061172a878288016114d1565b925050606061173b878288016114d1565b91505092959194509250565b6000806040838503121561175a57600080fd5b600083013567ffffffffffffffff81111561177457600080fd5b61178085828601611510565b9250506020611791858286016115b8565b9150509250929050565b6117a481611dc4565b82525050565b6117b381611dc4565b82525050565b6117c281611dd6565b82525050565b6117d181611dd6565b82525050565b60006117e282611d6b565b6117ec8185611d8c565b93506117fc818560208601611e25565b80840191505092915050565b600061181382611d81565b61181d8185611da8565b935061182d818560208601611e25565b61183681611e58565b840191505092915050565b600061184c82611d81565b6118568185611db9565b9350611866818560208601611e25565b80840191505092915050565b600061187d82611d76565b6118878185611d97565b9350611897818560208601611e25565b6118a081611e58565b840191505092915050565b60006118b682611d76565b6118c08185611da8565b93506118d0818560208601611e25565b6118d981611e58565b840191505092915050565b60006118f1602683611da8565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611957601d83611da8565b91507f52656769737472793a2076657273696f6e2069732072657175697265640000006000830152602082019050919050565b6000611997601a83611da8565b91507f52656769737472793a206e616d652069732072657175697265640000000000006000830152602082019050919050565b60006119d7601f83611da8565b91507f52656769737472793a206163636f756e742063616e2774206265207a65726f006000830152602082019050919050565b6000611a17602083611da8565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b600060c083016000830151611a6260008601826117b9565b5060208301518482036020860152611a7a8282611872565b91505060408301518482036040860152611a948282611872565b9150506060830151611aa9606086018261179b565b506080830151611abc608086018261179b565b5060a0830151611acf60a0860182611ada565b508091505092915050565b611ae381611e0c565b82525050565b611af281611e0c565b82525050565b6000611b0482846117d7565b915081905092915050565b6000611b1b8284611841565b915081905092915050565b6000602082019050611b3b60008301846117aa565b92915050565b6000602082019050611b5660008301846117c8565b92915050565b600060c082019050611b7160008301896117c8565b8181036020830152611b8381886118ab565b90508181036040830152611b9781876118ab565b9050611ba660608301866117aa565b611bb360808301856117aa565b611bc060a0830184611ae9565b979650505050505050565b60006020820190508181036000830152611be581846118ab565b905092915050565b60006040820190508181036000830152611c078185611808565b90508181036020830152611c1b8184611808565b90509392505050565b60006020820190508181036000830152611c3d816118e4565b9050919050565b60006020820190508181036000830152611c5d8161194a565b9050919050565b60006020820190508181036000830152611c7d8161198a565b9050919050565b60006020820190508181036000830152611c9d816119ca565b9050919050565b60006020820190508181036000830152611cbd81611a0a565b9050919050565b60006020820190508181036000830152611cde8184611a4a565b905092915050565b6000604051905081810181811067ffffffffffffffff82111715611d0957600080fd5b8060405250919050565b600067ffffffffffffffff821115611d2a57600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115611d5657600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611dcf82611dec565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611e43578082015181840152602081019050611e28565b83811115611e52576000848401525b50505050565b6000601f19601f8301169050919050565b611e7281611dc4565b8114611e7d57600080fd5b50565b611e8981611de2565b8114611e9457600080fd5b50565b611ea081611e0c565b8114611eab57600080fd5b5056fea365627a7a72315820fa12361d908435ec798d4d9f05e3936502d695e4992f949a417eb743655d0e926c6578706572696d656e74616cf564736f6c634300050d0040"

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistryCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistrySession) IsOwner() (bool, error) {
	return _Registry.Contract.IsOwner(&_Registry.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Registry *RegistryCallerSession) IsOwner() (bool, error) {
	return _Registry.Contract.IsOwner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Record is a free data retrieval call binding the contract method 0x470bb62b.
//
// Solidity: function record(string name, string version) constant returns(RegistryRecord)
func (_Registry *RegistryCaller) Record(opts *bind.CallOpts, name string, version string) (RegistryRecord, error) {
	var (
		ret0 = new(RegistryRecord)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "record", name, version)
	return *ret0, err
}

// Record is a free data retrieval call binding the contract method 0x470bb62b.
//
// Solidity: function record(string name, string version) constant returns(RegistryRecord)
func (_Registry *RegistrySession) Record(name string, version string) (RegistryRecord, error) {
	return _Registry.Contract.Record(&_Registry.CallOpts, name, version)
}

// Record is a free data retrieval call binding the contract method 0x470bb62b.
//
// Solidity: function record(string name, string version) constant returns(RegistryRecord)
func (_Registry *RegistryCallerSession) Record(name string, version string) (RegistryRecord, error) {
	return _Registry.Contract.Record(&_Registry.CallOpts, name, version)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, address owner, uint256 index)
func (_Registry *RegistryCaller) Records(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Owner       common.Address
	Index       *big.Int
}, error) {
	ret := new(struct {
		Initialized bool
		Name        string
		Version     string
		Account     common.Address
		Owner       common.Address
		Index       *big.Int
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "records", arg0)
	return *ret, err
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, address owner, uint256 index)
func (_Registry *RegistrySession) Records(arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Owner       common.Address
	Index       *big.Int
}, error) {
	return _Registry.Contract.Records(&_Registry.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, address owner, uint256 index)
func (_Registry *RegistryCallerSession) Records(arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Owner       common.Address
	Index       *big.Int
}, error) {
	return _Registry.Contract.Records(&_Registry.CallOpts, arg0)
}

// Versions is a free data retrieval call binding the contract method 0x460e3ca2.
//
// Solidity: function versions(string , uint256 ) constant returns(string)
func (_Registry *RegistryCaller) Versions(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "versions", arg0, arg1)
	return *ret0, err
}

// Versions is a free data retrieval call binding the contract method 0x460e3ca2.
//
// Solidity: function versions(string , uint256 ) constant returns(string)
func (_Registry *RegistrySession) Versions(arg0 string, arg1 *big.Int) (string, error) {
	return _Registry.Contract.Versions(&_Registry.CallOpts, arg0, arg1)
}

// Versions is a free data retrieval call binding the contract method 0x460e3ca2.
//
// Solidity: function versions(string , uint256 ) constant returns(string)
func (_Registry *RegistryCallerSession) Versions(arg0 string, arg1 *big.Int) (string, error) {
	return _Registry.Contract.Versions(&_Registry.CallOpts, arg0, arg1)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string name, string version) returns()
func (_Registry *RegistryTransactor) Remove(opts *bind.TransactOpts, name string, version string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "remove", name, version)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string name, string version) returns()
func (_Registry *RegistrySession) Remove(name string, version string) (*types.Transaction, error) {
	return _Registry.Contract.Remove(&_Registry.TransactOpts, name, version)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string name, string version) returns()
func (_Registry *RegistryTransactorSession) Remove(name string, version string) (*types.Transaction, error) {
	return _Registry.Contract.Remove(&_Registry.TransactOpts, name, version)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0xb03205b0.
//
// Solidity: function update(string name, string version, address account, address owner) returns()
func (_Registry *RegistryTransactor) Update(opts *bind.TransactOpts, name string, version string, account common.Address, owner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "update", name, version, account, owner)
}

// Update is a paid mutator transaction binding the contract method 0xb03205b0.
//
// Solidity: function update(string name, string version, address account, address owner) returns()
func (_Registry *RegistrySession) Update(name string, version string, account common.Address, owner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Update(&_Registry.TransactOpts, name, version, account, owner)
}

// Update is a paid mutator transaction binding the contract method 0xb03205b0.
//
// Solidity: function update(string name, string version, address account, address owner) returns()
func (_Registry *RegistryTransactorSession) Update(name string, version string, account common.Address, owner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Update(&_Registry.TransactOpts, name, version, account, owner)
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryVersionAddedIterator is returned from FilterVersionAdded and is used to iterate over the raw logs and unpacked data for VersionAdded events raised by the Registry contract.
type RegistryVersionAddedIterator struct {
	Event *RegistryVersionAdded // Event containing the contract specifics and raw log

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
func (it *RegistryVersionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryVersionAdded)
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
		it.Event = new(RegistryVersionAdded)
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
func (it *RegistryVersionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryVersionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryVersionAdded represents a VersionAdded event raised by the Registry contract.
type RegistryVersionAdded struct {
	Name    common.Hash
	Version common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVersionAdded is a free log retrieval operation binding the contract event 0x600be21ea77e400a162a88fb0bef7e7e4e0d1b28da89d643862828bc7a630cf6.
//
// Solidity: event VersionAdded(string indexed name, string indexed version)
func (_Registry *RegistryFilterer) FilterVersionAdded(opts *bind.FilterOpts, name []string, version []string) (*RegistryVersionAddedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "VersionAdded", nameRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &RegistryVersionAddedIterator{contract: _Registry.contract, event: "VersionAdded", logs: logs, sub: sub}, nil
}

// WatchVersionAdded is a free log subscription operation binding the contract event 0x600be21ea77e400a162a88fb0bef7e7e4e0d1b28da89d643862828bc7a630cf6.
//
// Solidity: event VersionAdded(string indexed name, string indexed version)
func (_Registry *RegistryFilterer) WatchVersionAdded(opts *bind.WatchOpts, sink chan<- *RegistryVersionAdded, name []string, version []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "VersionAdded", nameRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryVersionAdded)
				if err := _Registry.contract.UnpackLog(event, "VersionAdded", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseVersionAdded(log types.Log) (*RegistryVersionAdded, error) {
	event := new(RegistryVersionAdded)
	if err := _Registry.contract.UnpackLog(event, "VersionAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryVersionUpdatedIterator is returned from FilterVersionUpdated and is used to iterate over the raw logs and unpacked data for VersionUpdated events raised by the Registry contract.
type RegistryVersionUpdatedIterator struct {
	Event *RegistryVersionUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryVersionUpdated)
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
		it.Event = new(RegistryVersionUpdated)
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
func (it *RegistryVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryVersionUpdated represents a VersionUpdated event raised by the Registry contract.
type RegistryVersionUpdated struct {
	Name     common.Hash
	Versions common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVersionUpdated is a free log retrieval operation binding the contract event 0x22efc5f993dce37856b77dd72d7d7661032380c9728c4133f3c071c591bc6ca7.
//
// Solidity: event VersionUpdated(string indexed name, string indexed versions)
func (_Registry *RegistryFilterer) FilterVersionUpdated(opts *bind.FilterOpts, name []string, versions []string) (*RegistryVersionUpdatedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionsRule []interface{}
	for _, versionsItem := range versions {
		versionsRule = append(versionsRule, versionsItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "VersionUpdated", nameRule, versionsRule)
	if err != nil {
		return nil, err
	}
	return &RegistryVersionUpdatedIterator{contract: _Registry.contract, event: "VersionUpdated", logs: logs, sub: sub}, nil
}

// WatchVersionUpdated is a free log subscription operation binding the contract event 0x22efc5f993dce37856b77dd72d7d7661032380c9728c4133f3c071c591bc6ca7.
//
// Solidity: event VersionUpdated(string indexed name, string indexed versions)
func (_Registry *RegistryFilterer) WatchVersionUpdated(opts *bind.WatchOpts, sink chan<- *RegistryVersionUpdated, name []string, versions []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionsRule []interface{}
	for _, versionsItem := range versions {
		versionsRule = append(versionsRule, versionsItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "VersionUpdated", nameRule, versionsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryVersionUpdated)
				if err := _Registry.contract.UnpackLog(event, "VersionUpdated", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseVersionUpdated(log types.Log) (*RegistryVersionUpdated, error) {
	event := new(RegistryVersionUpdated)
	if err := _Registry.contract.UnpackLog(event, "VersionUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
