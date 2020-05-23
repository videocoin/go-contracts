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

// RegistryContractInfo is an auto generated low-level Go binding around an user-defined struct.
type RegistryContractInfo struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Owner       common.Address
}

// RegistryABI is the input ABI used to generate the binding from.
const RegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"RecordAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"versions\",\"type\":\"string\"}],\"name\":\"RecordUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"record\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structRegistry.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"records\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"versions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
var RegistryBin = "0x608060405260006100146100b760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100bf565b600033905090565b611dad806100ce6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063715018a611610066578063715018a61461015c5780638da5cb5b146101665780638f32d59b14610184578063b03205b0146101a2578063f2fde38b146101be57610093565b806301e64725146100985780631f1e00dc146100cc578063470bb62b146100fc5780635107f3481461012c575b600080fd5b6100b260048036036100ad9190810190611426565b6101da565b6040516100c39594939291906119eb565b60405180910390f35b6100e660048036036100e191908101906115b8565b61038d565b6040516100f39190611a4c565b60405180910390f35b610116600480360361011191908101906114b9565b610511565b6040516101239190611b65565b60405180910390f35b61014660048036036101419190810190611478565b61084a565b6040516101539190611b87565b60405180910390f35b61016461087e565b005b61016e610984565b60405161017b91906119b5565b60405180910390f35b61018c6109ad565b60405161019991906119d0565b60405180910390f35b6101bc60048036036101b79190810190611525565b610a0b565b005b6101d860048036036101d391908101906113fd565b610ff6565b005b60016020528060005260406000206000915090508060000160009054906101000a900460ff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102995780601f1061026e57610100808354040283529160200191610299565b820191906000526020600020905b81548152906001019060200180831161027c57829003601f168201915b505050505090806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103375780601f1061030c57610100808354040283529160200191610337565b820191906000526020600020905b81548152906001019060200180831161031a57829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905085565b60606000835114156103d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103cb90611b05565b60405180910390fd5b6002836040516103e4919061199e565b908152602001604051809103902060010180549050821061043a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043190611ae5565b60405180910390fd5b60028360405161044a919061199e565b9081526020016040518091039020600101828154811061046657fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105045780601f106104d957610100808354040283529160200191610504565b820191906000526020600020905b8154815290600101906020018083116104e757829003601f168201915b5050505050905092915050565b61051961117f565b60008351141561055e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055590611b05565b60405180910390fd5b6000825114156105a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059a90611ac5565b60405180910390fd5b6000600284846040516020016105ba929190611a6e565b6040516020818303038152906040526040516105d69190611987565b602060405180830381855afa1580156105f3573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250610616919081019061144f565b9050600160008281526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff16151515158152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106ea5780601f106106bf576101008083540402835291602001916106ea565b820191906000526020600020905b8154815290600101906020018083116106cd57829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561078c5780601f106107615761010080835404028352916020019161078c565b820191906000526020600020905b81548152906001019060200180831161076f57829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505091505092915050565b6002818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154905081565b6108866109ad565b6108c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108bc90611b45565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166109ef611049565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b610a136109ad565b610a52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4990611b45565b60405180910390fd5b600084511415610a97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8e90611b05565b60405180910390fd5b600083511415610adc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad390611ac5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b4c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b4390611b25565b60405180910390fd5b600060028585604051602001610b63929190611a6e565b604051602081830303815290604052604051610b7f9190611987565b602060405180830381855afa158015610b9c573d6000803e3d6000fd5b5050506040513d601f19601f82011682018060405250610bbf919081019061144f565b90506001600082815260200190815260200160002060000160009054906101000a900460ff16610e4257600285604051610bf9919061199e565b9081526020016040518091039020600101849080600181540180825580915050906001820390600052602060002001600090919290919091509080519060200190610c459291906111dc565b5050600285604051610c57919061199e565b908152602001604051809103902060010180549050600286604051610c7c919061199e565b9081526020016040518091039020600001819055506040518060a001604052806001151581526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152506001600083815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190610d3692919061125c565b506040820151816002019080519060200190610d5392919061125c565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505083604051610df3919061199e565b604051809103902085604051610e09919061199e565b60405180910390207fbd0ee5ea7025d9c57d179b2ab864af8addb63b746cb7544d7e152d2f5943212f60405160405180910390a3610fef565b6040518060a001604052806001151581526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152506001600083815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190610ee792919061125c565b506040820151816002019080519060200190610f0492919061125c565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505083604051610fa4919061199e565b604051809103902085604051610fba919061199e565b60405180910390207fa32b92311d4910d53fdddcf93a238419a815906305d579caf8d933aacd39f64060405160405180910390a35b5050505050565b610ffe6109ad565b61103d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161103490611b45565b60405180910390fd5b61104681611051565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156110c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110b890611aa5565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060a001604052806000151581526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061121d57805160ff191683800117855561124b565b8280016001018555821561124b579182015b8281111561124a57825182559160200191906001019061122f565b5b50905061125891906112dc565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061129d57805160ff19168380011785556112cb565b828001600101855582156112cb579182015b828111156112ca5782518255916020019190600101906112af565b5b5090506112d891906112dc565b5090565b6112fe91905b808211156112fa5760008160009055506001016112e2565b5090565b90565b60008135905061131081611d25565b92915050565b60008135905061132581611d3c565b92915050565b60008151905061133a81611d3c565b92915050565b600082601f83011261135157600080fd5b813561136461135f82611bcf565b611ba2565b9150808252602083016020830185838301111561138057600080fd5b61138b838284611cd2565b50505092915050565b600082601f8301126113a557600080fd5b81356113b86113b382611bfb565b611ba2565b915080825260208301602083018583830111156113d457600080fd5b6113df838284611cd2565b50505092915050565b6000813590506113f781611d53565b92915050565b60006020828403121561140f57600080fd5b600061141d84828501611301565b91505092915050565b60006020828403121561143857600080fd5b600061144684828501611316565b91505092915050565b60006020828403121561146157600080fd5b600061146f8482850161132b565b91505092915050565b60006020828403121561148a57600080fd5b600082013567ffffffffffffffff8111156114a457600080fd5b6114b084828501611340565b91505092915050565b600080604083850312156114cc57600080fd5b600083013567ffffffffffffffff8111156114e657600080fd5b6114f285828601611394565b925050602083013567ffffffffffffffff81111561150f57600080fd5b61151b85828601611394565b9150509250929050565b6000806000806080858703121561153b57600080fd5b600085013567ffffffffffffffff81111561155557600080fd5b61156187828801611394565b945050602085013567ffffffffffffffff81111561157e57600080fd5b61158a87828801611394565b935050604061159b87828801611301565b92505060606115ac87828801611301565b91505092959194509250565b600080604083850312156115cb57600080fd5b600083013567ffffffffffffffff8111156115e557600080fd5b6115f185828601611394565b9250506020611602858286016113e8565b9150509250929050565b61161581611c80565b82525050565b61162481611c80565b82525050565b61163381611c92565b82525050565b61164281611c92565b82525050565b600061165382611c27565b61165d8185611c48565b935061166d818560208601611ce1565b80840191505092915050565b600061168482611c3d565b61168e8185611c64565b935061169e818560208601611ce1565b6116a781611d14565b840191505092915050565b60006116bd82611c3d565b6116c78185611c75565b93506116d7818560208601611ce1565b80840191505092915050565b60006116ee82611c32565b6116f88185611c53565b9350611708818560208601611ce1565b61171181611d14565b840191505092915050565b600061172782611c32565b6117318185611c64565b9350611741818560208601611ce1565b61174a81611d14565b840191505092915050565b6000611762602683611c64565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006117c8601d83611c64565b91507f52656769737472793a2076657273696f6e2069732072657175697265640000006000830152602082019050919050565b6000611808601683611c64565b91507f52656769737472793a206f7574206f662072616e6765000000000000000000006000830152602082019050919050565b6000611848601a83611c64565b91507f52656769737472793a206e616d652069732072657175697265640000000000006000830152602082019050919050565b6000611888601f83611c64565b91507f52656769737472793a206163636f756e742063616e2774206265207a65726f006000830152602082019050919050565b60006118c8602083611c64565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b600060a083016000830151611913600086018261162a565b506020830151848203602086015261192b82826116e3565b9150506040830151848203604086015261194582826116e3565b915050606083015161195a606086018261160c565b50608083015161196d608086018261160c565b508091505092915050565b61198181611cc8565b82525050565b60006119938284611648565b915081905092915050565b60006119aa82846116b2565b915081905092915050565b60006020820190506119ca600083018461161b565b92915050565b60006020820190506119e56000830184611639565b92915050565b600060a082019050611a006000830188611639565b8181036020830152611a12818761171c565b90508181036040830152611a26818661171c565b9050611a35606083018561161b565b611a42608083018461161b565b9695505050505050565b60006020820190508181036000830152611a668184611679565b905092915050565b60006040820190508181036000830152611a888185611679565b90508181036020830152611a9c8184611679565b90509392505050565b60006020820190508181036000830152611abe81611755565b9050919050565b60006020820190508181036000830152611ade816117bb565b9050919050565b60006020820190508181036000830152611afe816117fb565b9050919050565b60006020820190508181036000830152611b1e8161183b565b9050919050565b60006020820190508181036000830152611b3e8161187b565b9050919050565b60006020820190508181036000830152611b5e816118bb565b9050919050565b60006020820190508181036000830152611b7f81846118fb565b905092915050565b6000602082019050611b9c6000830184611978565b92915050565b6000604051905081810181811067ffffffffffffffff82111715611bc557600080fd5b8060405250919050565b600067ffffffffffffffff821115611be657600080fd5b601f19601f8301169050602081019050919050565b600067ffffffffffffffff821115611c1257600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611c8b82611ca8565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611cff578082015181840152602081019050611ce4565b83811115611d0e576000848401525b50505050565b6000601f19601f8301169050919050565b611d2e81611c80565b8114611d3957600080fd5b50565b611d4581611c9e565b8114611d5057600080fd5b50565b611d5c81611cc8565b8114611d6757600080fd5b5056fea365627a7a72315820e24cef7d56c3c00b3cbe8ab078ccaebc27aa60eff37a8bd04300d5c6333b6d546c6578706572696d656e74616cf564736f6c634300050d0040"

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
// Solidity: function record(string name, string version) constant returns(RegistryContractInfo)
func (_Registry *RegistryCaller) Record(opts *bind.CallOpts, name string, version string) (RegistryContractInfo, error) {
	var (
		ret0 = new(RegistryContractInfo)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "record", name, version)
	return *ret0, err
}

// Record is a free data retrieval call binding the contract method 0x470bb62b.
//
// Solidity: function record(string name, string version) constant returns(RegistryContractInfo)
func (_Registry *RegistrySession) Record(name string, version string) (RegistryContractInfo, error) {
	return _Registry.Contract.Record(&_Registry.CallOpts, name, version)
}

// Record is a free data retrieval call binding the contract method 0x470bb62b.
//
// Solidity: function record(string name, string version) constant returns(RegistryContractInfo)
func (_Registry *RegistryCallerSession) Record(name string, version string) (RegistryContractInfo, error) {
	return _Registry.Contract.Record(&_Registry.CallOpts, name, version)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, address owner)
func (_Registry *RegistryCaller) Records(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Owner       common.Address
}, error) {
	ret := new(struct {
		Initialized bool
		Name        string
		Version     string
		Account     common.Address
		Owner       common.Address
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "records", arg0)
	return *ret, err
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, address owner)
func (_Registry *RegistrySession) Records(arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Owner       common.Address
}, error) {
	return _Registry.Contract.Records(&_Registry.CallOpts, arg0)
}

// Records is a free data retrieval call binding the contract method 0x01e64725.
//
// Solidity: function records(bytes32 ) constant returns(bool initialized, string name, string version, address account, address owner)
func (_Registry *RegistryCallerSession) Records(arg0 [32]byte) (struct {
	Initialized bool
	Name        string
	Version     string
	Account     common.Address
	Owner       common.Address
}, error) {
	return _Registry.Contract.Records(&_Registry.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x1f1e00dc.
//
// Solidity: function version(string name, uint256 index) constant returns(string)
func (_Registry *RegistryCaller) Version(opts *bind.CallOpts, name string, index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "version", name, index)
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x1f1e00dc.
//
// Solidity: function version(string name, uint256 index) constant returns(string)
func (_Registry *RegistrySession) Version(name string, index *big.Int) (string, error) {
	return _Registry.Contract.Version(&_Registry.CallOpts, name, index)
}

// Version is a free data retrieval call binding the contract method 0x1f1e00dc.
//
// Solidity: function version(string name, uint256 index) constant returns(string)
func (_Registry *RegistryCallerSession) Version(name string, index *big.Int) (string, error) {
	return _Registry.Contract.Version(&_Registry.CallOpts, name, index)
}

// Versions is a free data retrieval call binding the contract method 0x5107f348.
//
// Solidity: function versions(string ) constant returns(uint256 length)
func (_Registry *RegistryCaller) Versions(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "versions", arg0)
	return *ret0, err
}

// Versions is a free data retrieval call binding the contract method 0x5107f348.
//
// Solidity: function versions(string ) constant returns(uint256 length)
func (_Registry *RegistrySession) Versions(arg0 string) (*big.Int, error) {
	return _Registry.Contract.Versions(&_Registry.CallOpts, arg0)
}

// Versions is a free data retrieval call binding the contract method 0x5107f348.
//
// Solidity: function versions(string ) constant returns(uint256 length)
func (_Registry *RegistryCallerSession) Versions(arg0 string) (*big.Int, error) {
	return _Registry.Contract.Versions(&_Registry.CallOpts, arg0)
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

// RegistryRecordAddedIterator is returned from FilterRecordAdded and is used to iterate over the raw logs and unpacked data for RecordAdded events raised by the Registry contract.
type RegistryRecordAddedIterator struct {
	Event *RegistryRecordAdded // Event containing the contract specifics and raw log

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
func (it *RegistryRecordAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRecordAdded)
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
		it.Event = new(RegistryRecordAdded)
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
func (it *RegistryRecordAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRecordAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRecordAdded represents a RecordAdded event raised by the Registry contract.
type RegistryRecordAdded struct {
	Name    common.Hash
	Version common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecordAdded is a free log retrieval operation binding the contract event 0xbd0ee5ea7025d9c57d179b2ab864af8addb63b746cb7544d7e152d2f5943212f.
//
// Solidity: event RecordAdded(string indexed name, string indexed version)
func (_Registry *RegistryFilterer) FilterRecordAdded(opts *bind.FilterOpts, name []string, version []string) (*RegistryRecordAddedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RecordAdded", nameRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRecordAddedIterator{contract: _Registry.contract, event: "RecordAdded", logs: logs, sub: sub}, nil
}

// WatchRecordAdded is a free log subscription operation binding the contract event 0xbd0ee5ea7025d9c57d179b2ab864af8addb63b746cb7544d7e152d2f5943212f.
//
// Solidity: event RecordAdded(string indexed name, string indexed version)
func (_Registry *RegistryFilterer) WatchRecordAdded(opts *bind.WatchOpts, sink chan<- *RegistryRecordAdded, name []string, version []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RecordAdded", nameRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRecordAdded)
				if err := _Registry.contract.UnpackLog(event, "RecordAdded", log); err != nil {
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

// ParseRecordAdded is a log parse operation binding the contract event 0xbd0ee5ea7025d9c57d179b2ab864af8addb63b746cb7544d7e152d2f5943212f.
//
// Solidity: event RecordAdded(string indexed name, string indexed version)
func (_Registry *RegistryFilterer) ParseRecordAdded(log types.Log) (*RegistryRecordAdded, error) {
	event := new(RegistryRecordAdded)
	if err := _Registry.contract.UnpackLog(event, "RecordAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RegistryRecordUpdatedIterator is returned from FilterRecordUpdated and is used to iterate over the raw logs and unpacked data for RecordUpdated events raised by the Registry contract.
type RegistryRecordUpdatedIterator struct {
	Event *RegistryRecordUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryRecordUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRecordUpdated)
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
		it.Event = new(RegistryRecordUpdated)
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
func (it *RegistryRecordUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRecordUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRecordUpdated represents a RecordUpdated event raised by the Registry contract.
type RegistryRecordUpdated struct {
	Name     common.Hash
	Versions common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRecordUpdated is a free log retrieval operation binding the contract event 0xa32b92311d4910d53fdddcf93a238419a815906305d579caf8d933aacd39f640.
//
// Solidity: event RecordUpdated(string indexed name, string indexed versions)
func (_Registry *RegistryFilterer) FilterRecordUpdated(opts *bind.FilterOpts, name []string, versions []string) (*RegistryRecordUpdatedIterator, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionsRule []interface{}
	for _, versionsItem := range versions {
		versionsRule = append(versionsRule, versionsItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RecordUpdated", nameRule, versionsRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRecordUpdatedIterator{contract: _Registry.contract, event: "RecordUpdated", logs: logs, sub: sub}, nil
}

// WatchRecordUpdated is a free log subscription operation binding the contract event 0xa32b92311d4910d53fdddcf93a238419a815906305d579caf8d933aacd39f640.
//
// Solidity: event RecordUpdated(string indexed name, string indexed versions)
func (_Registry *RegistryFilterer) WatchRecordUpdated(opts *bind.WatchOpts, sink chan<- *RegistryRecordUpdated, name []string, versions []string) (event.Subscription, error) {

	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var versionsRule []interface{}
	for _, versionsItem := range versions {
		versionsRule = append(versionsRule, versionsItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RecordUpdated", nameRule, versionsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRecordUpdated)
				if err := _Registry.contract.UnpackLog(event, "RecordUpdated", log); err != nil {
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

// ParseRecordUpdated is a log parse operation binding the contract event 0xa32b92311d4910d53fdddcf93a238419a815906305d579caf8d933aacd39f640.
//
// Solidity: event RecordUpdated(string indexed name, string indexed versions)
func (_Registry *RegistryFilterer) ParseRecordUpdated(log types.Log) (*RegistryRecordUpdated, error) {
	event := new(RegistryRecordUpdated)
	if err := _Registry.contract.UnpackLog(event, "RecordUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
