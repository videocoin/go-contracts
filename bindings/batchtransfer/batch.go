// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package batchtransfer

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

// BatchTransferTransfer is an auto generated low-level Go binding around an user-defined struct.
type BatchTransferTransfer struct {
	To     common.Address
	Amount *big.Int
}

// BatchTransferABI is the input ABI used to generate the binding from.
const BatchTransferABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BatchedTransfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structBatchTransfer.Transfer[]\",\"name\":\"transfers\",\"type\":\"tuple[]\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BatchTransferBin is the compiled bytecode used for deploying new contracts.
var BatchTransferBin = "0x608060405234801561001057600080fd5b50610b52806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80636331687f14610030575b600080fd5b61004a6004803603610045919081019061069f565b61004c565b005b6100793330858773ffffffffffffffffffffffffffffffffffffffff166101b9909392919063ffffffff16565b60008090505b8282905081101561017d5761009261055c565b83838381811061009e57fe5b9050604002018036036100b4919081019061070b565b90506100e9816000015182602001518873ffffffffffffffffffffffffffffffffffffffff1661025b9092919063ffffffff16565b6101008160200151866102fa90919063ffffffff16565b9450806000015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f0c000455819b4d62986a32fd2a23a6915aad0986ddf1e3551cc7b0439eade9168360200151604051610167919061099b565b60405180910390a350808060010191505061007f565b5060008311156101b3576101b233848673ffffffffffffffffffffffffffffffffffffffff1661025b9092919063ffffffff16565b5b50505050565b610255848573ffffffffffffffffffffffffffffffffffffffff166323b872dd905060e01b8585856040516024016101f3939291906108b9565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610344565b50505050565b6102f5838473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb905060e01b84846040516024016102939291906108f0565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610344565b505050565b600061033c83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506104b6565b905092915050565b6103638273ffffffffffffffffffffffffffffffffffffffff16610511565b6103a2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103999061097b565b60405180910390fd5b600060608373ffffffffffffffffffffffffffffffffffffffff16836040516103cb91906108a2565b6000604051808303816000865af19150503d8060008114610408576040519150601f19603f3d011682016040523d82523d6000602084013e61040d565b606091505b509150915081610452576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104499061093b565b60405180910390fd5b6000815111156104b057808060200190516104709190810190610676565b6104af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a69061095b565b60405180910390fd5b5b50505050565b60008383111582906104fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f59190610919565b60405180910390fd5b5060008385039050809150509392505050565b60008060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f915080821415801561055357506000801b8214155b92505050919050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b60008135905061059b81610ab3565b92915050565b60008083601f8401126105b357600080fd5b8235905067ffffffffffffffff8111156105cc57600080fd5b6020830191508360408202830111156105e457600080fd5b9250929050565b6000815190506105fa81610aca565b92915050565b60008135905061060f81610ae1565b92915050565b60006040828403121561062757600080fd5b61063160406109b6565b905060006106418482850161058c565b600083015250602061065584828501610661565b60208301525092915050565b60008135905061067081610af8565b92915050565b60006020828403121561068857600080fd5b6000610696848285016105eb565b91505092915050565b600080600080606085870312156106b557600080fd5b60006106c387828801610600565b94505060206106d487828801610661565b935050604085013567ffffffffffffffff8111156106f157600080fd5b6106fd878288016105a1565b925092505092959194509250565b60006040828403121561071d57600080fd5b600061072b84828501610615565b91505092915050565b61073d81610a15565b82525050565b600061074e826109e3565b61075881856109f9565b9350610768818560208601610a6f565b80840191505092915050565b600061077f826109ee565b6107898185610a04565b9350610799818560208601610a6f565b6107a281610aa2565b840191505092915050565b60006107ba602083610a04565b91507f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65646000830152602082019050919050565b60006107fa602a83610a04565b91507f5361666545524332303a204552433230206f7065726174696f6e20646964206e60008301527f6f742073756363656564000000000000000000000000000000000000000000006020830152604082019050919050565b6000610860601f83610a04565b91507f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e7472616374006000830152602082019050919050565b61089c81610a65565b82525050565b60006108ae8284610743565b915081905092915050565b60006060820190506108ce6000830186610734565b6108db6020830185610734565b6108e86040830184610893565b949350505050565b60006040820190506109056000830185610734565b6109126020830184610893565b9392505050565b600060208201905081810360008301526109338184610774565b905092915050565b60006020820190508181036000830152610954816107ad565b9050919050565b60006020820190508181036000830152610974816107ed565b9050919050565b6000602082019050818103600083015261099481610853565b9050919050565b60006020820190506109b06000830184610893565b92915050565b6000604051905081810181811067ffffffffffffffff821117156109d957600080fd5b8060405250919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000610a2082610a45565b9050919050565b60008115159050919050565b6000610a3e82610a15565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015610a8d578082015181840152602081019050610a72565b83811115610a9c576000848401525b50505050565b6000601f19601f8301169050919050565b610abc81610a15565b8114610ac757600080fd5b50565b610ad381610a27565b8114610ade57600080fd5b50565b610aea81610a33565b8114610af557600080fd5b50565b610b0181610a65565b8114610b0c57600080fd5b5056fea365627a7a72315820e59ad87dc46efa4df0bacf9f434f24ec5e69b230e2fdbe1a0792327e90156d146c6578706572696d656e74616cf564736f6c634300050d0040"

// DeployBatchTransfer deploys a new Ethereum contract, binding an instance of BatchTransfer to it.
func DeployBatchTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchTransfer, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchTransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BatchTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchTransfer{BatchTransferCaller: BatchTransferCaller{contract: contract}, BatchTransferTransactor: BatchTransferTransactor{contract: contract}, BatchTransferFilterer: BatchTransferFilterer{contract: contract}}, nil
}

// BatchTransfer is an auto generated Go binding around an Ethereum contract.
type BatchTransfer struct {
	BatchTransferCaller     // Read-only binding to the contract
	BatchTransferTransactor // Write-only binding to the contract
	BatchTransferFilterer   // Log filterer for contract events
}

// BatchTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchTransferSession struct {
	Contract     *BatchTransfer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchTransferCallerSession struct {
	Contract *BatchTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BatchTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchTransferTransactorSession struct {
	Contract     *BatchTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BatchTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchTransferRaw struct {
	Contract *BatchTransfer // Generic contract binding to access the raw methods on
}

// BatchTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchTransferCallerRaw struct {
	Contract *BatchTransferCaller // Generic read-only contract binding to access the raw methods on
}

// BatchTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchTransferTransactorRaw struct {
	Contract *BatchTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchTransfer creates a new instance of BatchTransfer, bound to a specific deployed contract.
func NewBatchTransfer(address common.Address, backend bind.ContractBackend) (*BatchTransfer, error) {
	contract, err := bindBatchTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchTransfer{BatchTransferCaller: BatchTransferCaller{contract: contract}, BatchTransferTransactor: BatchTransferTransactor{contract: contract}, BatchTransferFilterer: BatchTransferFilterer{contract: contract}}, nil
}

// NewBatchTransferCaller creates a new read-only instance of BatchTransfer, bound to a specific deployed contract.
func NewBatchTransferCaller(address common.Address, caller bind.ContractCaller) (*BatchTransferCaller, error) {
	contract, err := bindBatchTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchTransferCaller{contract: contract}, nil
}

// NewBatchTransferTransactor creates a new write-only instance of BatchTransfer, bound to a specific deployed contract.
func NewBatchTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchTransferTransactor, error) {
	contract, err := bindBatchTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchTransferTransactor{contract: contract}, nil
}

// NewBatchTransferFilterer creates a new log filterer instance of BatchTransfer, bound to a specific deployed contract.
func NewBatchTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchTransferFilterer, error) {
	contract, err := bindBatchTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchTransferFilterer{contract: contract}, nil
}

// bindBatchTransfer binds a generic wrapper to an already deployed contract.
func bindBatchTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchTransfer *BatchTransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BatchTransfer.Contract.BatchTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchTransfer *BatchTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchTransfer.Contract.BatchTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchTransfer *BatchTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchTransfer.Contract.BatchTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchTransfer *BatchTransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BatchTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchTransfer *BatchTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchTransfer *BatchTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchTransfer.Contract.contract.Transact(opts, method, params...)
}

// Transfer is a paid mutator transaction binding the contract method 0x6331687f.
//
// Solidity: function transfer(address token, uint256 total, []BatchTransferTransfer transfers) returns()
func (_BatchTransfer *BatchTransferTransactor) Transfer(opts *bind.TransactOpts, token common.Address, total *big.Int, transfers []BatchTransferTransfer) (*types.Transaction, error) {
	return _BatchTransfer.contract.Transact(opts, "transfer", token, total, transfers)
}

// Transfer is a paid mutator transaction binding the contract method 0x6331687f.
//
// Solidity: function transfer(address token, uint256 total, []BatchTransferTransfer transfers) returns()
func (_BatchTransfer *BatchTransferSession) Transfer(token common.Address, total *big.Int, transfers []BatchTransferTransfer) (*types.Transaction, error) {
	return _BatchTransfer.Contract.Transfer(&_BatchTransfer.TransactOpts, token, total, transfers)
}

// Transfer is a paid mutator transaction binding the contract method 0x6331687f.
//
// Solidity: function transfer(address token, uint256 total, []BatchTransferTransfer transfers) returns()
func (_BatchTransfer *BatchTransferTransactorSession) Transfer(token common.Address, total *big.Int, transfers []BatchTransferTransfer) (*types.Transaction, error) {
	return _BatchTransfer.Contract.Transfer(&_BatchTransfer.TransactOpts, token, total, transfers)
}

// BatchTransferBatchedTransferIterator is returned from FilterBatchedTransfer and is used to iterate over the raw logs and unpacked data for BatchedTransfer events raised by the BatchTransfer contract.
type BatchTransferBatchedTransferIterator struct {
	Event *BatchTransferBatchedTransfer // Event containing the contract specifics and raw log

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
func (it *BatchTransferBatchedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchTransferBatchedTransfer)
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
		it.Event = new(BatchTransferBatchedTransfer)
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
func (it *BatchTransferBatchedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchTransferBatchedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchTransferBatchedTransfer represents a BatchedTransfer event raised by the BatchTransfer contract.
type BatchTransferBatchedTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchedTransfer is a free log retrieval operation binding the contract event 0x0c000455819b4d62986a32fd2a23a6915aad0986ddf1e3551cc7b0439eade916.
//
// Solidity: event BatchedTransfer(address indexed from, address indexed to, uint256 amount)
func (_BatchTransfer *BatchTransferFilterer) FilterBatchedTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BatchTransferBatchedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BatchTransfer.contract.FilterLogs(opts, "BatchedTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BatchTransferBatchedTransferIterator{contract: _BatchTransfer.contract, event: "BatchedTransfer", logs: logs, sub: sub}, nil
}

// WatchBatchedTransfer is a free log subscription operation binding the contract event 0x0c000455819b4d62986a32fd2a23a6915aad0986ddf1e3551cc7b0439eade916.
//
// Solidity: event BatchedTransfer(address indexed from, address indexed to, uint256 amount)
func (_BatchTransfer *BatchTransferFilterer) WatchBatchedTransfer(opts *bind.WatchOpts, sink chan<- *BatchTransferBatchedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BatchTransfer.contract.WatchLogs(opts, "BatchedTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchTransferBatchedTransfer)
				if err := _BatchTransfer.contract.UnpackLog(event, "BatchedTransfer", log); err != nil {
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

// ParseBatchedTransfer is a log parse operation binding the contract event 0x0c000455819b4d62986a32fd2a23a6915aad0986ddf1e3551cc7b0439eade916.
//
// Solidity: event BatchedTransfer(address indexed from, address indexed to, uint256 amount)
func (_BatchTransfer *BatchTransferFilterer) ParseBatchedTransfer(log types.Log) (*BatchTransferBatchedTransfer, error) {
	event := new(BatchTransferBatchedTransfer)
	if err := _BatchTransfer.contract.UnpackLog(event, "BatchedTransfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
