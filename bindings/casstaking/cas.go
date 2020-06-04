// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package casstaking

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CASStakingChange is an auto generated low-level Go binding around an user-defined struct.
type CASStakingChange struct {
	Delegator  common.Address
	Transcoder common.Address
	Amount     *big.Int
	Ctype      uint8
}

// CASStakingABI is the input ABI used to generate the binding from.
const CASStakingABI = "[{\"inputs\":[{\"internalType\":\"contractIStakingManager\",\"name\":\"_staking\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumCASStaking.ChangeType\",\"name\":\"ctype\",\"type\":\"uint8\"}],\"internalType\":\"structCASStaking.Change[]\",\"name\":\"changes\",\"type\":\"tuple[]\"}],\"name\":\"cas\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"processed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractIStakingManager\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CASStakingBin is the compiled bytecode used for deploying new contracts.
var CASStakingBin = "0x608060405234801561001057600080fd5b5060405162000f7138038062000f7183398181016040526100349190810190610146565b600061004461012960201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506101ca565b600033905090565b600081519050610140816101b3565b92915050565b60006020828403121561015857600080fd5b600061016684828501610131565b91505092915050565b600061017a82610193565b9050919050565b600061018c8261016f565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6101bc81610181565b81146101c757600080fd5b50565b610d9780620001da6000396000f3fe6080604052600436106100705760003560e01c8063715018a61161004e578063715018a6146100e45780638da5cb5b146100fb5780638f32d59b14610126578063f2fde38b1461015157610070565b80631a1517a6146100725780632ce5c2841461008e5780634cf088d9146100b9575b005b61008c6004803603610087919081019061098f565b61017a565b005b34801561009a57600080fd5b506100a3610451565b6040516100b09190610c29565b60405180910390f35b3480156100c557600080fd5b506100ce610457565b6040516100db9190610bae565b60405180910390f35b3480156100f057600080fd5b506100f961047d565b005b34801561010757600080fd5b50610110610583565b60405161011d9190610b18565b60405180910390f35b34801561013257600080fd5b5061013b6105ac565b6040516101489190610b93565b60405180910390f35b34801561015d57600080fd5b506101786004803603610173919081019061093d565b61060a565b005b6101826105ac565b6101c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b890610be9565b60405180910390fd5b8260015414610205576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fc90610c09565b60405180910390fd5b8160018190555060008090505b81518110156103ed57610223610793565b82828151811061022f57fe5b602002602001015190506000600181111561024657fe5b8160600151600181111561025657fe5b14156102fd57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315620cce8260400151836020015184600001516040518463ffffffff1660e01b81526004016102c6929190610b33565b6000604051808303818588803b1580156102df57600080fd5b505af11580156102f3573d6000803e3d6000fd5b50505050506103df565b60018081111561030957fe5b8160600151600181111561031957fe5b14156103de57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663254124fa8260200151836000015184604001516040518463ffffffff1660e01b815260040161038a93929190610b5c565b602060405180830381600087803b1580156103a457600080fd5b505af11580156103b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103dc9190810190610966565b505b5b508080600101915050610212565b503373ffffffffffffffffffffffffffffffffffffffff166108fc3073ffffffffffffffffffffffffffffffffffffffff16319081150290604051600060405180830381858888f1935050505015801561044b573d6000803e3d6000fd5b50505050565b60015481565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6104856105ac565b6104c4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bb90610be9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166105ee61065d565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b6106126105ac565b610651576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064890610be9565b60405180910390fd5b61065a81610665565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156106d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106cc90610bc9565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060800160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600060018111156107ec57fe5b81525090565b60008135905061080181610d16565b92915050565b600082601f83011261081857600080fd5b813561082b61082682610c71565b610c44565b9150818183526020840193506020810190508385608084028201111561085057600080fd5b60005b838110156108805781610866888261089f565b845260208401935060808301925050600181019050610853565b5050505092915050565b60008135905061089981610d2d565b92915050565b6000608082840312156108b157600080fd5b6108bb6080610c44565b905060006108cb848285016107f2565b60008301525060206108df848285016107f2565b60208301525060406108f384828501610913565b60408301525060606109078482850161088a565b60608301525092915050565b60008135905061092281610d3d565b92915050565b60008151905061093781610d3d565b92915050565b60006020828403121561094f57600080fd5b600061095d848285016107f2565b91505092915050565b60006020828403121561097857600080fd5b600061098684828501610928565b91505092915050565b6000806000606084860312156109a457600080fd5b60006109b286828701610913565b93505060206109c386828701610913565b925050604084013567ffffffffffffffff8111156109e057600080fd5b6109ec86828701610807565b9150509250925092565b6109ff81610caa565b82525050565b610a0e81610cbc565b82525050565b610a1d81610cf2565b82525050565b6000610a30602683610c99565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000610a96602083610c99565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000610ad6600b83610c99565b91507f434153206661696c7572650000000000000000000000000000000000000000006000830152602082019050919050565b610b1281610ce8565b82525050565b6000602082019050610b2d60008301846109f6565b92915050565b6000604082019050610b4860008301856109f6565b610b5560208301846109f6565b9392505050565b6000606082019050610b7160008301866109f6565b610b7e60208301856109f6565b610b8b6040830184610b09565b949350505050565b6000602082019050610ba86000830184610a05565b92915050565b6000602082019050610bc36000830184610a14565b92915050565b60006020820190508181036000830152610be281610a23565b9050919050565b60006020820190508181036000830152610c0281610a89565b9050919050565b60006020820190508181036000830152610c2281610ac9565b9050919050565b6000602082019050610c3e6000830184610b09565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610c6757600080fd5b8060405250919050565b600067ffffffffffffffff821115610c8857600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000610cb582610cc8565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610cfd82610d04565b9050919050565b6000610d0f82610cc8565b9050919050565b610d1f81610caa565b8114610d2a57600080fd5b50565b60028110610d3a57600080fd5b50565b610d4681610ce8565b8114610d5157600080fd5b5056fea365627a7a72315820b8274281977576308914eabae2c7a0f20dc6383a038438c0ca0d7de5986ae6aa6c6578706572696d656e74616cf564736f6c634300050d0040"

// DeployCASStaking deploys a new Ethereum contract, binding an instance of CASStaking to it.
func DeployCASStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _staking common.Address) (common.Address, *types.Transaction, *CASStaking, error) {
	parsed, err := abi.JSON(strings.NewReader(CASStakingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CASStakingBin), backend, _staking)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CASStaking{CASStakingCaller: CASStakingCaller{contract: contract}, CASStakingTransactor: CASStakingTransactor{contract: contract}, CASStakingFilterer: CASStakingFilterer{contract: contract}}, nil
}

// CASStaking is an auto generated Go binding around an Ethereum contract.
type CASStaking struct {
	CASStakingCaller     // Read-only binding to the contract
	CASStakingTransactor // Write-only binding to the contract
	CASStakingFilterer   // Log filterer for contract events
}

// CASStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CASStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CASStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CASStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CASStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CASStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CASStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CASStakingSession struct {
	Contract     *CASStaking       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CASStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CASStakingCallerSession struct {
	Contract *CASStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CASStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CASStakingTransactorSession struct {
	Contract     *CASStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CASStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CASStakingRaw struct {
	Contract *CASStaking // Generic contract binding to access the raw methods on
}

// CASStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CASStakingCallerRaw struct {
	Contract *CASStakingCaller // Generic read-only contract binding to access the raw methods on
}

// CASStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CASStakingTransactorRaw struct {
	Contract *CASStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCASStaking creates a new instance of CASStaking, bound to a specific deployed contract.
func NewCASStaking(address common.Address, backend bind.ContractBackend) (*CASStaking, error) {
	contract, err := bindCASStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CASStaking{CASStakingCaller: CASStakingCaller{contract: contract}, CASStakingTransactor: CASStakingTransactor{contract: contract}, CASStakingFilterer: CASStakingFilterer{contract: contract}}, nil
}

// NewCASStakingCaller creates a new read-only instance of CASStaking, bound to a specific deployed contract.
func NewCASStakingCaller(address common.Address, caller bind.ContractCaller) (*CASStakingCaller, error) {
	contract, err := bindCASStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CASStakingCaller{contract: contract}, nil
}

// NewCASStakingTransactor creates a new write-only instance of CASStaking, bound to a specific deployed contract.
func NewCASStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*CASStakingTransactor, error) {
	contract, err := bindCASStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CASStakingTransactor{contract: contract}, nil
}

// NewCASStakingFilterer creates a new log filterer instance of CASStaking, bound to a specific deployed contract.
func NewCASStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*CASStakingFilterer, error) {
	contract, err := bindCASStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CASStakingFilterer{contract: contract}, nil
}

// bindCASStaking binds a generic wrapper to an already deployed contract.
func bindCASStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CASStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CASStaking *CASStakingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CASStaking.Contract.CASStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CASStaking *CASStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CASStaking.Contract.CASStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CASStaking *CASStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CASStaking.Contract.CASStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CASStaking *CASStakingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CASStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CASStaking *CASStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CASStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CASStaking *CASStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CASStaking.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_CASStaking *CASStakingCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CASStaking.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_CASStaking *CASStakingSession) IsOwner() (bool, error) {
	return _CASStaking.Contract.IsOwner(&_CASStaking.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_CASStaking *CASStakingCallerSession) IsOwner() (bool, error) {
	return _CASStaking.Contract.IsOwner(&_CASStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CASStaking *CASStakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CASStaking.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CASStaking *CASStakingSession) Owner() (common.Address, error) {
	return _CASStaking.Contract.Owner(&_CASStaking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CASStaking *CASStakingCallerSession) Owner() (common.Address, error) {
	return _CASStaking.Contract.Owner(&_CASStaking.CallOpts)
}

// Processed is a free data retrieval call binding the contract method 0x2ce5c284.
//
// Solidity: function processed() view returns(uint256)
func (_CASStaking *CASStakingCaller) Processed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CASStaking.contract.Call(opts, out, "processed")
	return *ret0, err
}

// Processed is a free data retrieval call binding the contract method 0x2ce5c284.
//
// Solidity: function processed() view returns(uint256)
func (_CASStaking *CASStakingSession) Processed() (*big.Int, error) {
	return _CASStaking.Contract.Processed(&_CASStaking.CallOpts)
}

// Processed is a free data retrieval call binding the contract method 0x2ce5c284.
//
// Solidity: function processed() view returns(uint256)
func (_CASStaking *CASStakingCallerSession) Processed() (*big.Int, error) {
	return _CASStaking.Contract.Processed(&_CASStaking.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_CASStaking *CASStakingCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CASStaking.contract.Call(opts, out, "staking")
	return *ret0, err
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_CASStaking *CASStakingSession) Staking() (common.Address, error) {
	return _CASStaking.Contract.Staking(&_CASStaking.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_CASStaking *CASStakingCallerSession) Staking() (common.Address, error) {
	return _CASStaking.Contract.Staking(&_CASStaking.CallOpts)
}

// Cas is a paid mutator transaction binding the contract method 0x1a1517a6.
//
// Solidity: function cas(uint256 from, uint256 to, []CASStakingChange changes) payable returns()
func (_CASStaking *CASStakingTransactor) Cas(opts *bind.TransactOpts, from *big.Int, to *big.Int, changes []CASStakingChange) (*types.Transaction, error) {
	return _CASStaking.contract.Transact(opts, "cas", from, to, changes)
}

// Cas is a paid mutator transaction binding the contract method 0x1a1517a6.
//
// Solidity: function cas(uint256 from, uint256 to, []CASStakingChange changes) payable returns()
func (_CASStaking *CASStakingSession) Cas(from *big.Int, to *big.Int, changes []CASStakingChange) (*types.Transaction, error) {
	return _CASStaking.Contract.Cas(&_CASStaking.TransactOpts, from, to, changes)
}

// Cas is a paid mutator transaction binding the contract method 0x1a1517a6.
//
// Solidity: function cas(uint256 from, uint256 to, []CASStakingChange changes) payable returns()
func (_CASStaking *CASStakingTransactorSession) Cas(from *big.Int, to *big.Int, changes []CASStakingChange) (*types.Transaction, error) {
	return _CASStaking.Contract.Cas(&_CASStaking.TransactOpts, from, to, changes)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CASStaking *CASStakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CASStaking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CASStaking *CASStakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _CASStaking.Contract.RenounceOwnership(&_CASStaking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CASStaking *CASStakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CASStaking.Contract.RenounceOwnership(&_CASStaking.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CASStaking *CASStakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CASStaking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CASStaking *CASStakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CASStaking.Contract.TransferOwnership(&_CASStaking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CASStaking *CASStakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CASStaking.Contract.TransferOwnership(&_CASStaking.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CASStaking *CASStakingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CASStaking.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CASStaking *CASStakingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CASStaking.Contract.Fallback(&_CASStaking.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CASStaking *CASStakingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CASStaking.Contract.Fallback(&_CASStaking.TransactOpts, calldata)
}

// CASStakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CASStaking contract.
type CASStakingOwnershipTransferredIterator struct {
	Event *CASStakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CASStakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CASStakingOwnershipTransferred)
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
		it.Event = new(CASStakingOwnershipTransferred)
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
func (it *CASStakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CASStakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CASStakingOwnershipTransferred represents a OwnershipTransferred event raised by the CASStaking contract.
type CASStakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CASStaking *CASStakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CASStakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CASStaking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CASStakingOwnershipTransferredIterator{contract: _CASStaking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CASStaking *CASStakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CASStakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CASStaking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CASStakingOwnershipTransferred)
				if err := _CASStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CASStaking *CASStakingFilterer) ParseOwnershipTransferred(log types.Log) (*CASStakingOwnershipTransferred, error) {
	event := new(CASStakingOwnershipTransferred)
	if err := _CASStaking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
