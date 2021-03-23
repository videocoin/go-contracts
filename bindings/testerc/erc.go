// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testerc

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

// TestERCABI is the input ABI used to generate the binding from.
const TestERCABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestERCBin is the compiled bytecode used for deploying new contracts.
var TestERCBin = "0x608060405262000024620000186200002a60201b60201c565b6200003260201b60201c565b62000257565b600033905090565b6200004d8160036200009360201b620012a81790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b620000a582826200017760201b60201c565b1562000119576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000200576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620018506022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6115e980620002676000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063983b2d5611610071578063983b2d56146102e7578063986502751461032b578063a457c2d714610335578063a9059cbb1461039b578063aa271e1a14610401578063dd62ed3e1461045d576100b4565b8063095ea7b3146100b957806318160ddd1461011f57806323b872dd1461013d57806339509351146101c357806340c10f191461022957806370a082311461028f575b600080fd5b610105600480360360408110156100cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104d5565b604051808215151515815260200191505060405180910390f35b6101276104f3565b6040518082815260200191505060405180910390f35b6101a96004803603606081101561015357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104fd565b604051808215151515815260200191505060405180910390f35b61020f600480360360408110156101d957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105d6565b604051808215151515815260200191505060405180910390f35b6102756004803603604081101561023f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610689565b604051808215151515815260200191505060405180910390f35b6102d1600480360360208110156102a557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610704565b6040518082815260200191505060405180910390f35b610329600480360360208110156102fd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061074c565b005b6103336107bd565b005b6103816004803603604081101561034b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506107cf565b604051808215151515815260200191505060405180910390f35b6103e7600480360360408110156103b157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061089c565b604051808215151515815260200191505060405180910390f35b6104436004803603602081101561041757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108ba565b604051808215151515815260200191505060405180910390f35b6104bf6004803603604081101561047357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108d7565b6040518082815260200191505060405180910390f35b60006104e96104e261095e565b8484610966565b6001905092915050565b6000600254905090565b600061050a848484610b5d565b6105cb8461051661095e565b6105c6856040518060600160405280602881526020016114fd60289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061057c61095e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e139092919063ffffffff16565b610966565b600190509392505050565b600061067f6105e361095e565b8461067a85600160006105f461095e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ed390919063ffffffff16565b610966565b6001905092915050565b600061069b61069661095e565b6108ba565b6106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806114ac6030913960400191505060405180910390fd5b6106fa8383610f5b565b6001905092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61075c61075761095e565b6108ba565b6107b1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001806114ac6030913960400191505060405180910390fd5b6107ba81611116565b50565b6107cd6107c861095e565b611170565b565b60006108926107dc61095e565b8461088d85604051806060016040528060258152602001611590602591396001600061080661095e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e139092919063ffffffff16565b610966565b6001905092915050565b60006108b06108a961095e565b8484610b5d565b6001905092915050565b60006108d08260036111ca90919063ffffffff16565b9050919050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156109ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602481526020018061156c6024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a72576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806114646022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610be3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806115476025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c69576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806114416023913960400191505060405180910390fd5b610cd481604051806060016040528060268152602001611486602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610e139092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d67816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ed390919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290610ec0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e85578082015181840152602081019050610e6a565b50505050905090810190601f168015610eb25780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600080828401905083811015610f51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ffe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b61101381600254610ed390919063ffffffff16565b60028190555061106a816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610ed390919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b61112a8160036112a890919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f660405160405180910390a250565b61118481600361138390919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669260405160405180910390a250565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611251576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806115256022913960400191505060405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6112b282826111ca565b15611325576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f526f6c65733a206163636f756e7420616c72656164792068617320726f6c650081525060200191505060405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b61138d82826111ca565b6113e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806114dc6021913960400191505060405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63654d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365526f6c65733a206163636f756e7420697320746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a72315820a4c4c1024437cc9ef0f372c7a4e11c35747d41ae6eeac4b9c518b189b6dbbd6364736f6c634300050d0032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployTestERC deploys a new Ethereum contract, binding an instance of TestERC to it.
func DeployTestERC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestERC, error) {
	parsed, err := abi.JSON(strings.NewReader(TestERCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestERCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestERC{TestERCCaller: TestERCCaller{contract: contract}, TestERCTransactor: TestERCTransactor{contract: contract}, TestERCFilterer: TestERCFilterer{contract: contract}}, nil
}

// TestERC is an auto generated Go binding around an Ethereum contract.
type TestERC struct {
	TestERCCaller     // Read-only binding to the contract
	TestERCTransactor // Write-only binding to the contract
	TestERCFilterer   // Log filterer for contract events
}

// TestERCCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestERCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestERCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestERCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestERCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestERCSession struct {
	Contract     *TestERC          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestERCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestERCCallerSession struct {
	Contract *TestERCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TestERCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestERCTransactorSession struct {
	Contract     *TestERCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestERCRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestERCRaw struct {
	Contract *TestERC // Generic contract binding to access the raw methods on
}

// TestERCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestERCCallerRaw struct {
	Contract *TestERCCaller // Generic read-only contract binding to access the raw methods on
}

// TestERCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestERCTransactorRaw struct {
	Contract *TestERCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestERC creates a new instance of TestERC, bound to a specific deployed contract.
func NewTestERC(address common.Address, backend bind.ContractBackend) (*TestERC, error) {
	contract, err := bindTestERC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestERC{TestERCCaller: TestERCCaller{contract: contract}, TestERCTransactor: TestERCTransactor{contract: contract}, TestERCFilterer: TestERCFilterer{contract: contract}}, nil
}

// NewTestERCCaller creates a new read-only instance of TestERC, bound to a specific deployed contract.
func NewTestERCCaller(address common.Address, caller bind.ContractCaller) (*TestERCCaller, error) {
	contract, err := bindTestERC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestERCCaller{contract: contract}, nil
}

// NewTestERCTransactor creates a new write-only instance of TestERC, bound to a specific deployed contract.
func NewTestERCTransactor(address common.Address, transactor bind.ContractTransactor) (*TestERCTransactor, error) {
	contract, err := bindTestERC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestERCTransactor{contract: contract}, nil
}

// NewTestERCFilterer creates a new log filterer instance of TestERC, bound to a specific deployed contract.
func NewTestERCFilterer(address common.Address, filterer bind.ContractFilterer) (*TestERCFilterer, error) {
	contract, err := bindTestERC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestERCFilterer{contract: contract}, nil
}

// bindTestERC binds a generic wrapper to an already deployed contract.
func bindTestERC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestERCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestERC *TestERCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestERC.Contract.TestERCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestERC *TestERCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC.Contract.TestERCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestERC *TestERCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestERC.Contract.TestERCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestERC *TestERCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestERC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestERC *TestERCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestERC *TestERCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestERC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestERC *TestERCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestERC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestERC *TestERCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestERC.Contract.Allowance(&_TestERC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestERC *TestERCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestERC.Contract.Allowance(&_TestERC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestERC *TestERCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestERC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestERC *TestERCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestERC.Contract.BalanceOf(&_TestERC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestERC *TestERCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestERC.Contract.BalanceOf(&_TestERC.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_TestERC *TestERCCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _TestERC.contract.Call(opts, &out, "isMinter", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_TestERC *TestERCSession) IsMinter(account common.Address) (bool, error) {
	return _TestERC.Contract.IsMinter(&_TestERC.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_TestERC *TestERCCallerSession) IsMinter(account common.Address) (bool, error) {
	return _TestERC.Contract.IsMinter(&_TestERC.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestERC *TestERCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestERC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestERC *TestERCSession) TotalSupply() (*big.Int, error) {
	return _TestERC.Contract.TotalSupply(&_TestERC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestERC *TestERCCallerSession) TotalSupply() (*big.Int, error) {
	return _TestERC.Contract.TotalSupply(&_TestERC.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_TestERC *TestERCTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TestERC.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_TestERC *TestERCSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _TestERC.Contract.AddMinter(&_TestERC.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_TestERC *TestERCTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _TestERC.Contract.AddMinter(&_TestERC.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestERC *TestERCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestERC *TestERCSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.Approve(&_TestERC.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestERC *TestERCTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.Approve(&_TestERC.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestERC *TestERCTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestERC.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestERC *TestERCSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.DecreaseAllowance(&_TestERC.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestERC *TestERCTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.DecreaseAllowance(&_TestERC.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestERC *TestERCTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestERC.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestERC *TestERCSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.IncreaseAllowance(&_TestERC.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestERC *TestERCTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.IncreaseAllowance(&_TestERC.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_TestERC *TestERCTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_TestERC *TestERCSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.Mint(&_TestERC.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_TestERC *TestERCTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.Mint(&_TestERC.TransactOpts, account, amount)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_TestERC *TestERCTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestERC.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_TestERC *TestERCSession) RenounceMinter() (*types.Transaction, error) {
	return _TestERC.Contract.RenounceMinter(&_TestERC.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_TestERC *TestERCTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _TestERC.Contract.RenounceMinter(&_TestERC.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestERC *TestERCTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestERC *TestERCSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.Transfer(&_TestERC.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TestERC *TestERCTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.Transfer(&_TestERC.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestERC *TestERCTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestERC *TestERCSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.TransferFrom(&_TestERC.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TestERC *TestERCTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestERC.Contract.TransferFrom(&_TestERC.TransactOpts, sender, recipient, amount)
}

// TestERCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestERC contract.
type TestERCApprovalIterator struct {
	Event *TestERCApproval // Event containing the contract specifics and raw log

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
func (it *TestERCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERCApproval)
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
		it.Event = new(TestERCApproval)
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
func (it *TestERCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERCApproval represents a Approval event raised by the TestERC contract.
type TestERCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestERC *TestERCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestERCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestERC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestERCApprovalIterator{contract: _TestERC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestERC *TestERCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestERCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestERC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERCApproval)
				if err := _TestERC.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestERC *TestERCFilterer) ParseApproval(log types.Log) (*TestERCApproval, error) {
	event := new(TestERCApproval)
	if err := _TestERC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERCMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the TestERC contract.
type TestERCMinterAddedIterator struct {
	Event *TestERCMinterAdded // Event containing the contract specifics and raw log

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
func (it *TestERCMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERCMinterAdded)
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
		it.Event = new(TestERCMinterAdded)
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
func (it *TestERCMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERCMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERCMinterAdded represents a MinterAdded event raised by the TestERC contract.
type TestERCMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_TestERC *TestERCFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*TestERCMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestERC.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &TestERCMinterAddedIterator{contract: _TestERC.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_TestERC *TestERCFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *TestERCMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestERC.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERCMinterAdded)
				if err := _TestERC.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_TestERC *TestERCFilterer) ParseMinterAdded(log types.Log) (*TestERCMinterAdded, error) {
	event := new(TestERCMinterAdded)
	if err := _TestERC.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERCMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the TestERC contract.
type TestERCMinterRemovedIterator struct {
	Event *TestERCMinterRemoved // Event containing the contract specifics and raw log

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
func (it *TestERCMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERCMinterRemoved)
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
		it.Event = new(TestERCMinterRemoved)
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
func (it *TestERCMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERCMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERCMinterRemoved represents a MinterRemoved event raised by the TestERC contract.
type TestERCMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_TestERC *TestERCFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*TestERCMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestERC.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &TestERCMinterRemovedIterator{contract: _TestERC.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_TestERC *TestERCFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *TestERCMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TestERC.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERCMinterRemoved)
				if err := _TestERC.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_TestERC *TestERCFilterer) ParseMinterRemoved(log types.Log) (*TestERCMinterRemoved, error) {
	event := new(TestERCMinterRemoved)
	if err := _TestERC.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestERCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestERC contract.
type TestERCTransferIterator struct {
	Event *TestERCTransfer // Event containing the contract specifics and raw log

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
func (it *TestERCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestERCTransfer)
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
		it.Event = new(TestERCTransfer)
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
func (it *TestERCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestERCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestERCTransfer represents a Transfer event raised by the TestERC contract.
type TestERCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestERC *TestERCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestERCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestERC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestERCTransferIterator{contract: _TestERC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestERC *TestERCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestERCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestERC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestERCTransfer)
				if err := _TestERC.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestERC *TestERCFilterer) ParseTransfer(log types.Log) (*TestERCTransfer, error) {
	event := new(TestERCTransfer)
	if err := _TestERC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
