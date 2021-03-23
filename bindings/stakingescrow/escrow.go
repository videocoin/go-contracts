// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingescrow

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

// StakingEscrowABI is the input ABI used to generate the binding from.
const StakingEscrowABI = "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StakingEscrowBin is the compiled bytecode used for deploying new contracts.
var StakingEscrowBin = "0x608060405234801561001057600080fd5b50604051610edb380380610edb8339818101604052602081101561003357600080fd5b810190808051906020019092919050505080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610e46806100956000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806323b872dd14610046578063a9059cbb146100cc578063db20266f14610132575b600080fd5b6100b26004803603606081101561005c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506101aa565b604051808215151515815260200191505060405180910390f35b610118600480360360408110156100e257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610560565b604051808215151515815260200191505060405180910390f35b6101946004803603604081101561014857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107a2565b6040518082815260200191505060405180910390f35b6000808211610221576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f63616e2774207769746864726177207a65726f0000000000000000000000000081525060200191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146102a5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180610dc76021913960400191505060405180910390fd5b816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610396576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601d8152602001807f726571756573746564206d6f7265207468616e20617661696c61626c6500000081525060200191505060405180910390fd5b610424826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546107c790919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506104f08383600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108119092919063ffffffff16565b8373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fe6e0ef9cd056ca98561ca60e347ada61e1ede2f1142a078951b7a52e1e508e60846040518082815260200191505060405180910390a3600190509392505050565b60008082116105d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f63616e2774206465706f736974207a65726f000000000000000000000000000081525060200191505060405180910390fd5b610626333084600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108e2909392919063ffffffff16565b6106b4826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546109e890919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f989eaa915cbb416ea3d6f9a63b1a3de51770c7674b11fe21ecdf76b4e1d13910846040518082815260200191505060405180910390a36001905092915050565b6000602052816000526040600020602052806000526040600020600091509150505481565b600061080983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610a70565b905092915050565b6108dd838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb905060e01b8484604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610b30565b505050565b6109e2848573ffffffffffffffffffffffffffffffffffffffff166323b872dd905060e01b858585604051602401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610b30565b50505050565b600080828401905083811015610a66576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b6000838311158290610b1d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610ae2578082015181840152602081019050610ac7565b50505050905090810190601f168015610b0f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b610b4f8273ffffffffffffffffffffffffffffffffffffffff16610d7b565b610bc1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e74726163740081525060200191505060405180910390fd5b600060608373ffffffffffffffffffffffffffffffffffffffff16836040518082805190602001908083835b60208310610c105780518252602082019150602081019050602083039250610bed565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610c72576040519150601f19603f3d011682016040523d82523d6000602084013e610c77565b606091505b509150915081610cef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656481525060200191505060405180910390fd5b600081511115610d7557808060200190516020811015610d0e57600080fd5b8101908080519060200190929190505050610d74576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610de8602a913960400191505060405180910390fd5b5b50505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f9150808214158015610dbd57506000801b8214155b9250505091905056fe63616e206265206578656375746564206f6e6c7920627920726563697069656e745361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a265627a7a72315820d5292e932df385d5b23cfdf588355e5ef07c199381a588904594f09d1c92152064736f6c634300050d0032"

// DeployStakingEscrow deploys a new Ethereum contract, binding an instance of StakingEscrow to it.
func DeployStakingEscrow(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address) (common.Address, *types.Transaction, *StakingEscrow, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingEscrowABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakingEscrowBin), backend, token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingEscrow{StakingEscrowCaller: StakingEscrowCaller{contract: contract}, StakingEscrowTransactor: StakingEscrowTransactor{contract: contract}, StakingEscrowFilterer: StakingEscrowFilterer{contract: contract}}, nil
}

// StakingEscrow is an auto generated Go binding around an Ethereum contract.
type StakingEscrow struct {
	StakingEscrowCaller     // Read-only binding to the contract
	StakingEscrowTransactor // Write-only binding to the contract
	StakingEscrowFilterer   // Log filterer for contract events
}

// StakingEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingEscrowSession struct {
	Contract     *StakingEscrow    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingEscrowCallerSession struct {
	Contract *StakingEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakingEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingEscrowTransactorSession struct {
	Contract     *StakingEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingEscrowRaw struct {
	Contract *StakingEscrow // Generic contract binding to access the raw methods on
}

// StakingEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingEscrowCallerRaw struct {
	Contract *StakingEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// StakingEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingEscrowTransactorRaw struct {
	Contract *StakingEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingEscrow creates a new instance of StakingEscrow, bound to a specific deployed contract.
func NewStakingEscrow(address common.Address, backend bind.ContractBackend) (*StakingEscrow, error) {
	contract, err := bindStakingEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingEscrow{StakingEscrowCaller: StakingEscrowCaller{contract: contract}, StakingEscrowTransactor: StakingEscrowTransactor{contract: contract}, StakingEscrowFilterer: StakingEscrowFilterer{contract: contract}}, nil
}

// NewStakingEscrowCaller creates a new read-only instance of StakingEscrow, bound to a specific deployed contract.
func NewStakingEscrowCaller(address common.Address, caller bind.ContractCaller) (*StakingEscrowCaller, error) {
	contract, err := bindStakingEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingEscrowCaller{contract: contract}, nil
}

// NewStakingEscrowTransactor creates a new write-only instance of StakingEscrow, bound to a specific deployed contract.
func NewStakingEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingEscrowTransactor, error) {
	contract, err := bindStakingEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingEscrowTransactor{contract: contract}, nil
}

// NewStakingEscrowFilterer creates a new log filterer instance of StakingEscrow, bound to a specific deployed contract.
func NewStakingEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingEscrowFilterer, error) {
	contract, err := bindStakingEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingEscrowFilterer{contract: contract}, nil
}

// bindStakingEscrow binds a generic wrapper to an already deployed contract.
func bindStakingEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingEscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingEscrow *StakingEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingEscrow.Contract.StakingEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingEscrow *StakingEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingEscrow.Contract.StakingEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingEscrow *StakingEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingEscrow.Contract.StakingEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingEscrow *StakingEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingEscrow *StakingEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingEscrow *StakingEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingEscrow.Contract.contract.Transact(opts, method, params...)
}

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address , address ) view returns(uint256)
func (_StakingEscrow *StakingEscrowCaller) Locked(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingEscrow.contract.Call(opts, &out, "locked", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address , address ) view returns(uint256)
func (_StakingEscrow *StakingEscrowSession) Locked(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StakingEscrow.Contract.Locked(&_StakingEscrow.CallOpts, arg0, arg1)
}

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address , address ) view returns(uint256)
func (_StakingEscrow *StakingEscrowCallerSession) Locked(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StakingEscrow.Contract.Locked(&_StakingEscrow.CallOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakingEscrow *StakingEscrowTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingEscrow.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakingEscrow *StakingEscrowSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingEscrow.Contract.Transfer(&_StakingEscrow.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StakingEscrow *StakingEscrowTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingEscrow.Contract.Transfer(&_StakingEscrow.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakingEscrow *StakingEscrowTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingEscrow.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakingEscrow *StakingEscrowSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingEscrow.Contract.TransferFrom(&_StakingEscrow.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StakingEscrow *StakingEscrowTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingEscrow.Contract.TransferFrom(&_StakingEscrow.TransactOpts, sender, recipient, amount)
}

// StakingEscrowLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the StakingEscrow contract.
type StakingEscrowLockedIterator struct {
	Event *StakingEscrowLocked // Event containing the contract specifics and raw log

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
func (it *StakingEscrowLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingEscrowLocked)
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
		it.Event = new(StakingEscrowLocked)
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
func (it *StakingEscrowLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingEscrowLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingEscrowLocked represents a Locked event raised by the StakingEscrow contract.
type StakingEscrowLocked struct {
	Delegator common.Address
	Delegatee common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x989eaa915cbb416ea3d6f9a63b1a3de51770c7674b11fe21ecdf76b4e1d13910.
//
// Solidity: event Locked(address indexed delegator, address indexed delegatee, uint256 value)
func (_StakingEscrow *StakingEscrowFilterer) FilterLocked(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*StakingEscrowLockedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _StakingEscrow.contract.FilterLogs(opts, "Locked", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StakingEscrowLockedIterator{contract: _StakingEscrow.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x989eaa915cbb416ea3d6f9a63b1a3de51770c7674b11fe21ecdf76b4e1d13910.
//
// Solidity: event Locked(address indexed delegator, address indexed delegatee, uint256 value)
func (_StakingEscrow *StakingEscrowFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *StakingEscrowLocked, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _StakingEscrow.contract.WatchLogs(opts, "Locked", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingEscrowLocked)
				if err := _StakingEscrow.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x989eaa915cbb416ea3d6f9a63b1a3de51770c7674b11fe21ecdf76b4e1d13910.
//
// Solidity: event Locked(address indexed delegator, address indexed delegatee, uint256 value)
func (_StakingEscrow *StakingEscrowFilterer) ParseLocked(log types.Log) (*StakingEscrowLocked, error) {
	event := new(StakingEscrowLocked)
	if err := _StakingEscrow.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingEscrowUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the StakingEscrow contract.
type StakingEscrowUnlockedIterator struct {
	Event *StakingEscrowUnlocked // Event containing the contract specifics and raw log

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
func (it *StakingEscrowUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingEscrowUnlocked)
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
		it.Event = new(StakingEscrowUnlocked)
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
func (it *StakingEscrowUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingEscrowUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingEscrowUnlocked represents a Unlocked event raised by the StakingEscrow contract.
type StakingEscrowUnlocked struct {
	Delegator common.Address
	Delegatee common.Address
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0xe6e0ef9cd056ca98561ca60e347ada61e1ede2f1142a078951b7a52e1e508e60.
//
// Solidity: event Unlocked(address indexed delegator, address indexed delegatee, uint256 value)
func (_StakingEscrow *StakingEscrowFilterer) FilterUnlocked(opts *bind.FilterOpts, delegator []common.Address, delegatee []common.Address) (*StakingEscrowUnlockedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _StakingEscrow.contract.FilterLogs(opts, "Unlocked", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &StakingEscrowUnlockedIterator{contract: _StakingEscrow.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0xe6e0ef9cd056ca98561ca60e347ada61e1ede2f1142a078951b7a52e1e508e60.
//
// Solidity: event Unlocked(address indexed delegator, address indexed delegatee, uint256 value)
func (_StakingEscrow *StakingEscrowFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *StakingEscrowUnlocked, delegator []common.Address, delegatee []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _StakingEscrow.contract.WatchLogs(opts, "Unlocked", delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingEscrowUnlocked)
				if err := _StakingEscrow.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0xe6e0ef9cd056ca98561ca60e347ada61e1ede2f1142a078951b7a52e1e508e60.
//
// Solidity: event Unlocked(address indexed delegator, address indexed delegatee, uint256 value)
func (_StakingEscrow *StakingEscrowFilterer) ParseUnlocked(log types.Log) (*StakingEscrowUnlocked, error) {
	event := new(StakingEscrowUnlocked)
	if err := _StakingEscrow.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
