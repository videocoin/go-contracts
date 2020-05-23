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
const RegistryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"RecordAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"versions\",\"type\":\"string\"}],\"name\":\"RecordUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"record\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"initialized\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structRegistry.ContractInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"versions\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RegistryBin is the compiled bytecode used for deploying new contracts.
var RegistryBin = "0x608060405260006100146100b760201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100bf565b600033905090565b61191d806100ce6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100ec5780638f32d59b1461010a578063b03205b014610128578063f2fde38b146101445761007d565b8063470bb62b146100825780635107f348146100b2578063715018a6146100e2575b600080fd5b61009c6004803603610097919081019061112c565b610160565b6040516100a99190611704565b60405180910390f35b6100cc60048036036100c791908101906110eb565b610499565b6040516100d991906115f0565b60405180910390f35b6100ea6105e5565b005b6100f46106eb565b60405161010191906115d5565b60405180910390f35b610112610714565b60405161011f9190611612565b60405180910390f35b610142600480360361013d9190810190611198565b610772565b005b61015e60048036036101599190810190611099565b610d10565b005b610168610e99565b6000835114156101ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a4906116a4565b60405180910390fd5b6000825114156101f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e990611684565b60405180910390fd5b60006002848460405160200161020992919061162d565b60405160208183030381529060405260405161022591906115a7565b602060405180830381855afa158015610242573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525061026591908101906110c2565b9050600160008281526020019081526020016000206040518060a00160405290816000820160009054906101000a900460ff16151515158152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103395780601f1061030e57610100808354040283529160200191610339565b820191906000526020600020905b81548152906001019060200180831161031c57829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103db5780601f106103b0576101008083540402835291602001916103db565b820191906000526020600020905b8154815290600101906020018083116103be57829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152505091505092915050565b60606000825114156104e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d7906116a4565b60405180910390fd5b6002826040516104f091906115be565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156105da578382906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105c65780601f1061059b576101008083540402835291602001916105c6565b820191906000526020600020905b8154815290600101906020018083116105a957829003601f168201915b50505050508152602001906001019061051e565b505050509050919050565b6105ed610714565b61062c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610623906116e4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610756610d63565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b61077a610714565b6107b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b0906116e4565b60405180910390fd5b6000845114156107fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f5906116a4565b60405180910390fd5b600083511415610843576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083a90611684565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108aa906116c4565b60405180910390fd5b6000600285856040516020016108ca92919061162d565b6040516020818303038152906040526040516108e691906115a7565b602060405180830381855afa158015610903573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525061092691908101906110c2565b90506001600082815260200190815260200160002060000160009054906101000a900460ff16610b5c5760028560405161096091906115be565b90815260200160405180910390208490806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906109a9929190610ef6565b50506040518060a001604052806001151581526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152506001600083815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190610a50929190610f76565b506040820151816002019080519060200190610a6d929190610f76565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505083604051610b0d91906115be565b604051809103902085604051610b2391906115be565b60405180910390207fbd0ee5ea7025d9c57d179b2ab864af8addb63b746cb7544d7e152d2f5943212f60405160405180910390a3610d09565b6040518060a001604052806001151581526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff168152506001600083815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001019080519060200190610c01929190610f76565b506040820151816002019080519060200190610c1e929190610f76565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505083604051610cbe91906115be565b604051809103902085604051610cd491906115be565b60405180910390207fa32b92311d4910d53fdddcf93a238419a815906305d579caf8d933aacd39f64060405160405180910390a35b5050505050565b610d18610714565b610d57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4e906116e4565b60405180910390fd5b610d6081610d6b565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ddb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd290611664565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060a001604052806000151581526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f3757805160ff1916838001178555610f65565b82800160010185558215610f65579182015b82811115610f64578251825591602001919060010190610f49565b5b509050610f729190610ff6565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610fb757805160ff1916838001178555610fe5565b82800160010185558215610fe5579182015b82811115610fe4578251825591602001919060010190610fc9565b5b509050610ff29190610ff6565b5090565b61101891905b80821115611014576000816000905550600101610ffc565b5090565b90565b60008135905061102a816118ac565b92915050565b60008151905061103f816118c3565b92915050565b600082601f83011261105657600080fd5b813561106961106482611753565b611726565b9150808252602083016020830185838301111561108557600080fd5b611090838284611859565b50505092915050565b6000602082840312156110ab57600080fd5b60006110b98482850161101b565b91505092915050565b6000602082840312156110d457600080fd5b60006110e284828501611030565b91505092915050565b6000602082840312156110fd57600080fd5b600082013567ffffffffffffffff81111561111757600080fd5b61112384828501611045565b91505092915050565b6000806040838503121561113f57600080fd5b600083013567ffffffffffffffff81111561115957600080fd5b61116585828601611045565b925050602083013567ffffffffffffffff81111561118257600080fd5b61118e85828601611045565b9150509250929050565b600080600080608085870312156111ae57600080fd5b600085013567ffffffffffffffff8111156111c857600080fd5b6111d487828801611045565b945050602085013567ffffffffffffffff8111156111f157600080fd5b6111fd87828801611045565b935050604061120e8782880161101b565b925050606061121f8782880161101b565b91505092959194509250565b6000611237838361138b565b905092915050565b61124881611811565b82525050565b61125781611811565b82525050565b60006112688261178f565b61127281856117c8565b9350836020820285016112848561177f565b8060005b858110156112c057848403895281516112a1858261122b565b94506112ac836117bb565b925060208a01995050600181019050611288565b50829750879550505050505092915050565b6112db81611823565b82525050565b6112ea81611823565b82525050565b60006112fb8261179a565b61130581856117d9565b9350611315818560208601611868565b80840191505092915050565b600061132c826117b0565b61133681856117f5565b9350611346818560208601611868565b61134f8161189b565b840191505092915050565b6000611365826117b0565b61136f8185611806565b935061137f818560208601611868565b80840191505092915050565b6000611396826117a5565b6113a081856117e4565b93506113b0818560208601611868565b6113b98161189b565b840191505092915050565b60006113d16026836117f5565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611437601d836117f5565b91507f52656769737472793a2076657273696f6e2069732072657175697265640000006000830152602082019050919050565b6000611477601a836117f5565b91507f52656769737472793a206e616d652069732072657175697265640000000000006000830152602082019050919050565b60006114b7601f836117f5565b91507f52656769737472793a206163636f756e742063616e2774206265207a65726f006000830152602082019050919050565b60006114f76020836117f5565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b600060a08301600083015161154260008601826112d2565b506020830151848203602086015261155a828261138b565b91505060408301518482036040860152611574828261138b565b9150506060830151611589606086018261123f565b50608083015161159c608086018261123f565b508091505092915050565b60006115b382846112f0565b915081905092915050565b60006115ca828461135a565b915081905092915050565b60006020820190506115ea600083018461124e565b92915050565b6000602082019050818103600083015261160a818461125d565b905092915050565b600060208201905061162760008301846112e1565b92915050565b600060408201905081810360008301526116478185611321565b9050818103602083015261165b8184611321565b90509392505050565b6000602082019050818103600083015261167d816113c4565b9050919050565b6000602082019050818103600083015261169d8161142a565b9050919050565b600060208201905081810360008301526116bd8161146a565b9050919050565b600060208201905081810360008301526116dd816114aa565b9050919050565b600060208201905081810360008301526116fd816114ea565b9050919050565b6000602082019050818103600083015261171e818461152a565b905092915050565b6000604051905081810181811067ffffffffffffffff8211171561174957600080fd5b8060405250919050565b600067ffffffffffffffff82111561176a57600080fd5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061181c82611839565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b82818337600083830152505050565b60005b8381101561188657808201518184015260208101905061186b565b83811115611895576000848401525b50505050565b6000601f19601f8301169050919050565b6118b581611811565b81146118c057600080fd5b50565b6118cc8161182f565b81146118d757600080fd5b5056fea365627a7a72315820ebb8ab1e8bdaeca587a70efd64e33548549e5b0ad1567cbf580dbe0eabdaed7f6c6578706572696d656e74616cf564736f6c634300050d0040"

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

// Versions is a free data retrieval call binding the contract method 0x5107f348.
//
// Solidity: function versions(string name) constant returns(string[])
func (_Registry *RegistryCaller) Versions(opts *bind.CallOpts, name string) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Registry.contract.Call(opts, out, "versions", name)
	return *ret0, err
}

// Versions is a free data retrieval call binding the contract method 0x5107f348.
//
// Solidity: function versions(string name) constant returns(string[])
func (_Registry *RegistrySession) Versions(name string) ([]string, error) {
	return _Registry.Contract.Versions(&_Registry.CallOpts, name)
}

// Versions is a free data retrieval call binding the contract method 0x5107f348.
//
// Solidity: function versions(string name) constant returns(string[])
func (_Registry *RegistryCallerSession) Versions(name string) ([]string, error) {
	return _Registry.Contract.Versions(&_Registry.CallOpts, name)
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
