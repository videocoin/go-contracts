// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payments

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

// PaymentManagerABI is the input ABI used to generate the binding from.
const PaymentManagerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"PendingTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"Retry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"local\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"foreign\",\"type\":\"bytes32\"}],\"name\":\"TxFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"local\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"foreign\",\"type\":\"bytes32\"}],\"name\":\"TxSuccess\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"requestRetry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"local\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"foreign\",\"type\":\"bytes32\"}],\"name\":\"submitFailed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"local\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"foreign\",\"type\":\"bytes32\"}],\"name\":\"submitPending\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"local\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"foreign\",\"type\":\"bytes32\"}],\"name\":\"submitSuccess\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"enumPaymentManager.State\",\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PaymentManagerBin is the compiled bytecode used for deploying new contracts.
var PaymentManagerBin = "0x608060405234801561001057600080fd5b5060006100216100c460201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100cc565b600033905090565b61134a806100db6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638f32d59b116100665780638f32d59b146101c9578063ab274005146101eb578063f16d47ae14610223578063f2fde38b14610251578063f5e078411461029557610093565b80633c64f04b14610098578063477c44571461013d578063715018a6146101755780638da5cb5b1461017f575b600080fd5b6100c4600480360360208110156100ae57600080fd5b8101908080359060200190929190505050610301565b604051808581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018367ffffffffffffffff1667ffffffffffffffff16815260200182600381111561012657fe5b60ff16815260200194505050505060405180910390f35b6101736004803603604081101561015357600080fd5b810190808035906020019092919080359060200190929190505050610372565b005b61017d610613565b005b61018761074c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101d1610775565b604051808215151515815260200191505060405180910390f35b6102216004803603604081101561020157600080fd5b8101908080359060200190929190803590602001909291905050506107d3565b005b61024f6004803603602081101561023957600080fd5b8101908080359060200190929190505050610a74565b005b6102936004803603602081101561026757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cba565b005b6102ff600480360360808110156102ab57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803567ffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610d40565b005b60016020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900467ffffffffffffffff169080600101601c9054906101000a900460ff16905084565b61037a610775565b6103ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000801b821415610465576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e76616c6964206c6f63616c2074782068617368000000000000000000000081525060200191505060405180910390fd5b6000801b8114156104de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f696e76616c696420666f726569676e207478206861736800000000000000000081525060200191505060405180910390fd5b600060038111156104eb57fe5b60016000848152602001908152602001600020600101601c9054906101000a900460ff16600381111561051a57fe5b141561058e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f7265636f726420697320756e696e697469616c697a656400000000000000000081525060200191505060405180910390fd5b600260016000848152602001908152602001600020600101601c6101000a81548160ff021916908360038111156105c157fe5b021790555080600160008481526020019081526020016000206000018190555080827f4408cad08d2a774cfff1d2d5a614fd6bb763a24423b429da80f91eba6bdcb53b60405160405180910390a35050565b61061b610775565b61068d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166107b76111a3565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b6107db610775565b61084d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000801b8214156108c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e76616c6964206c6f63616c2074782068617368000000000000000000000081525060200191505060405180910390fd5b6000801b81141561093f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f696e76616c696420666f726569676e207478206861736800000000000000000081525060200191505060405180910390fd5b6000600381111561094c57fe5b60016000848152602001908152602001600020600101601c9054906101000a900460ff16600381111561097b57fe5b14156109ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f7265636f726420697320756e696e697469616c697a656400000000000000000081525060200191505060405180910390fd5b600360016000848152602001908152602001600020600101601c6101000a81548160ff02191690836003811115610a2257fe5b021790555080600160008481526020019081526020016000206000018190555080827fa2ce775ebcb679a2560ac924613e60823422156883902e37acbe857d09dc030960405160405180910390a35050565b610a7c610775565b610aee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000801b811415610b67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f696e76616c69642074782068617368000000000000000000000000000000000081525060200191505060405180910390fd5b60026003811115610b7457fe5b60016000838152602001908152602001600020600101601c9054906101000a900460ff166003811115610ba357fe5b14610c16576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f6f6e6c79206661696c6564207265636f7264730000000000000000000000000081525060200191505060405180910390fd5b600160008281526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160146101000a81549067ffffffffffffffff021916905560018201601c6101000a81549060ff02191690555050807f42ea0d481e8a8b75278dcea1de885ff3b50763d9026c0556046be840ade3ce1b60405160405180910390a250565b610cc2610775565b610d34576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b610d3d816111ab565b50565b610d48610775565b610dba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610e5d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f696e76616c69642061646472657373000000000000000000000000000000000081525060200191505060405180910390fd5b6000801b821415610ed6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f696e76616c6964206c6f63616c2074782068617368000000000000000000000081525060200191505060405180910390fd5b6000801b811415610f4f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f696e76616c696420666f726569676e207478206861736800000000000000000081525060200191505060405180910390fd5b60006003811115610f5c57fe5b60016000848152602001908152602001600020600101601c9054906101000a900460ff166003811115610f8b57fe5b1480610fd0575060016003811115610f9f57fe5b60016000848152602001908152602001600020600101601c9054906101000a900460ff166003811115610fce57fe5b145b611042576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f7265636f72642069732066696e616c697a65640000000000000000000000000081525060200191505060405180910390fd5b60405180608001604052808281526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018467ffffffffffffffff1681526020016001600381111561108c57fe5b815250600160008481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160010160146101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550606082015181600101601c6101000a81548160ff0219169083600381111561114657fe5b0217905550905050818367ffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f798316e895cab65e6f329061235135f6f9440c7582bb336547a6d05185b131bc60405160405180910390a450505050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611231576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806112f06026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a265627a7a72315820a10537d05f033002e08817378b5a4c89fc6b13bd3b0794c9b011f3972c48597c64736f6c634300050d0032"

// DeployPaymentManager deploys a new Ethereum contract, binding an instance of PaymentManager to it.
func DeployPaymentManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PaymentManager, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PaymentManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PaymentManager{PaymentManagerCaller: PaymentManagerCaller{contract: contract}, PaymentManagerTransactor: PaymentManagerTransactor{contract: contract}, PaymentManagerFilterer: PaymentManagerFilterer{contract: contract}}, nil
}

// PaymentManager is an auto generated Go binding around an Ethereum contract.
type PaymentManager struct {
	PaymentManagerCaller     // Read-only binding to the contract
	PaymentManagerTransactor // Write-only binding to the contract
	PaymentManagerFilterer   // Log filterer for contract events
}

// PaymentManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentManagerSession struct {
	Contract     *PaymentManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentManagerCallerSession struct {
	Contract *PaymentManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PaymentManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentManagerTransactorSession struct {
	Contract     *PaymentManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PaymentManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentManagerRaw struct {
	Contract *PaymentManager // Generic contract binding to access the raw methods on
}

// PaymentManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentManagerCallerRaw struct {
	Contract *PaymentManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentManagerTransactorRaw struct {
	Contract *PaymentManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentManager creates a new instance of PaymentManager, bound to a specific deployed contract.
func NewPaymentManager(address common.Address, backend bind.ContractBackend) (*PaymentManager, error) {
	contract, err := bindPaymentManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentManager{PaymentManagerCaller: PaymentManagerCaller{contract: contract}, PaymentManagerTransactor: PaymentManagerTransactor{contract: contract}, PaymentManagerFilterer: PaymentManagerFilterer{contract: contract}}, nil
}

// NewPaymentManagerCaller creates a new read-only instance of PaymentManager, bound to a specific deployed contract.
func NewPaymentManagerCaller(address common.Address, caller bind.ContractCaller) (*PaymentManagerCaller, error) {
	contract, err := bindPaymentManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentManagerCaller{contract: contract}, nil
}

// NewPaymentManagerTransactor creates a new write-only instance of PaymentManager, bound to a specific deployed contract.
func NewPaymentManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentManagerTransactor, error) {
	contract, err := bindPaymentManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentManagerTransactor{contract: contract}, nil
}

// NewPaymentManagerFilterer creates a new log filterer instance of PaymentManager, bound to a specific deployed contract.
func NewPaymentManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentManagerFilterer, error) {
	contract, err := bindPaymentManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentManagerFilterer{contract: contract}, nil
}

// bindPaymentManager binds a generic wrapper to an already deployed contract.
func bindPaymentManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentManager *PaymentManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentManager.Contract.PaymentManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentManager *PaymentManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentManager.Contract.PaymentManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentManager *PaymentManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentManager.Contract.PaymentManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentManager *PaymentManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentManager *PaymentManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentManager *PaymentManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentManager.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PaymentManager *PaymentManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PaymentManager.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PaymentManager *PaymentManagerSession) IsOwner() (bool, error) {
	return _PaymentManager.Contract.IsOwner(&_PaymentManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_PaymentManager *PaymentManagerCallerSession) IsOwner() (bool, error) {
	return _PaymentManager.Contract.IsOwner(&_PaymentManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentManager *PaymentManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentManager *PaymentManagerSession) Owner() (common.Address, error) {
	return _PaymentManager.Contract.Owner(&_PaymentManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentManager *PaymentManagerCallerSession) Owner() (common.Address, error) {
	return _PaymentManager.Contract.Owner(&_PaymentManager.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bytes32 hash, address signer, uint64 nonce, uint8 state)
func (_PaymentManager *PaymentManagerCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Hash   [32]byte
	Signer common.Address
	Nonce  uint64
	State  uint8
}, error) {
	var out []interface{}
	err := _PaymentManager.contract.Call(opts, &out, "transfers", arg0)

	outstruct := new(struct {
		Hash   [32]byte
		Signer common.Address
		Nonce  uint64
		State  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Signer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Nonce = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.State = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bytes32 hash, address signer, uint64 nonce, uint8 state)
func (_PaymentManager *PaymentManagerSession) Transfers(arg0 [32]byte) (struct {
	Hash   [32]byte
	Signer common.Address
	Nonce  uint64
	State  uint8
}, error) {
	return _PaymentManager.Contract.Transfers(&_PaymentManager.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bytes32 hash, address signer, uint64 nonce, uint8 state)
func (_PaymentManager *PaymentManagerCallerSession) Transfers(arg0 [32]byte) (struct {
	Hash   [32]byte
	Signer common.Address
	Nonce  uint64
	State  uint8
}, error) {
	return _PaymentManager.Contract.Transfers(&_PaymentManager.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentManager *PaymentManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentManager *PaymentManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentManager.Contract.RenounceOwnership(&_PaymentManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentManager *PaymentManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentManager.Contract.RenounceOwnership(&_PaymentManager.TransactOpts)
}

// RequestRetry is a paid mutator transaction binding the contract method 0xf16d47ae.
//
// Solidity: function requestRetry(bytes32 txHash) returns()
func (_PaymentManager *PaymentManagerTransactor) RequestRetry(opts *bind.TransactOpts, txHash [32]byte) (*types.Transaction, error) {
	return _PaymentManager.contract.Transact(opts, "requestRetry", txHash)
}

// RequestRetry is a paid mutator transaction binding the contract method 0xf16d47ae.
//
// Solidity: function requestRetry(bytes32 txHash) returns()
func (_PaymentManager *PaymentManagerSession) RequestRetry(txHash [32]byte) (*types.Transaction, error) {
	return _PaymentManager.Contract.RequestRetry(&_PaymentManager.TransactOpts, txHash)
}

// RequestRetry is a paid mutator transaction binding the contract method 0xf16d47ae.
//
// Solidity: function requestRetry(bytes32 txHash) returns()
func (_PaymentManager *PaymentManagerTransactorSession) RequestRetry(txHash [32]byte) (*types.Transaction, error) {
	return _PaymentManager.Contract.RequestRetry(&_PaymentManager.TransactOpts, txHash)
}

// SubmitFailed is a paid mutator transaction binding the contract method 0x477c4457.
//
// Solidity: function submitFailed(bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerTransactor) SubmitFailed(opts *bind.TransactOpts, local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.contract.Transact(opts, "submitFailed", local, foreign)
}

// SubmitFailed is a paid mutator transaction binding the contract method 0x477c4457.
//
// Solidity: function submitFailed(bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerSession) SubmitFailed(local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.Contract.SubmitFailed(&_PaymentManager.TransactOpts, local, foreign)
}

// SubmitFailed is a paid mutator transaction binding the contract method 0x477c4457.
//
// Solidity: function submitFailed(bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerTransactorSession) SubmitFailed(local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.Contract.SubmitFailed(&_PaymentManager.TransactOpts, local, foreign)
}

// SubmitPending is a paid mutator transaction binding the contract method 0xf5e07841.
//
// Solidity: function submitPending(address signer, uint64 nonce, bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerTransactor) SubmitPending(opts *bind.TransactOpts, signer common.Address, nonce uint64, local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.contract.Transact(opts, "submitPending", signer, nonce, local, foreign)
}

// SubmitPending is a paid mutator transaction binding the contract method 0xf5e07841.
//
// Solidity: function submitPending(address signer, uint64 nonce, bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerSession) SubmitPending(signer common.Address, nonce uint64, local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.Contract.SubmitPending(&_PaymentManager.TransactOpts, signer, nonce, local, foreign)
}

// SubmitPending is a paid mutator transaction binding the contract method 0xf5e07841.
//
// Solidity: function submitPending(address signer, uint64 nonce, bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerTransactorSession) SubmitPending(signer common.Address, nonce uint64, local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.Contract.SubmitPending(&_PaymentManager.TransactOpts, signer, nonce, local, foreign)
}

// SubmitSuccess is a paid mutator transaction binding the contract method 0xab274005.
//
// Solidity: function submitSuccess(bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerTransactor) SubmitSuccess(opts *bind.TransactOpts, local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.contract.Transact(opts, "submitSuccess", local, foreign)
}

// SubmitSuccess is a paid mutator transaction binding the contract method 0xab274005.
//
// Solidity: function submitSuccess(bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerSession) SubmitSuccess(local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.Contract.SubmitSuccess(&_PaymentManager.TransactOpts, local, foreign)
}

// SubmitSuccess is a paid mutator transaction binding the contract method 0xab274005.
//
// Solidity: function submitSuccess(bytes32 local, bytes32 foreign) returns()
func (_PaymentManager *PaymentManagerTransactorSession) SubmitSuccess(local [32]byte, foreign [32]byte) (*types.Transaction, error) {
	return _PaymentManager.Contract.SubmitSuccess(&_PaymentManager.TransactOpts, local, foreign)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentManager *PaymentManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PaymentManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentManager *PaymentManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaymentManager.Contract.TransferOwnership(&_PaymentManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentManager *PaymentManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaymentManager.Contract.TransferOwnership(&_PaymentManager.TransactOpts, newOwner)
}

// PaymentManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PaymentManager contract.
type PaymentManagerOwnershipTransferredIterator struct {
	Event *PaymentManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentManagerOwnershipTransferred)
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
		it.Event = new(PaymentManagerOwnershipTransferred)
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
func (it *PaymentManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentManagerOwnershipTransferred represents a OwnershipTransferred event raised by the PaymentManager contract.
type PaymentManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaymentManager *PaymentManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentManagerOwnershipTransferredIterator{contract: _PaymentManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaymentManager *PaymentManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentManagerOwnershipTransferred)
				if err := _PaymentManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PaymentManager *PaymentManagerFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentManagerOwnershipTransferred, error) {
	event := new(PaymentManagerOwnershipTransferred)
	if err := _PaymentManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentManagerPendingTransferIterator is returned from FilterPendingTransfer and is used to iterate over the raw logs and unpacked data for PendingTransfer events raised by the PaymentManager contract.
type PaymentManagerPendingTransferIterator struct {
	Event *PaymentManagerPendingTransfer // Event containing the contract specifics and raw log

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
func (it *PaymentManagerPendingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentManagerPendingTransfer)
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
		it.Event = new(PaymentManagerPendingTransfer)
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
func (it *PaymentManagerPendingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentManagerPendingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentManagerPendingTransfer represents a PendingTransfer event raised by the PaymentManager contract.
type PaymentManagerPendingTransfer struct {
	Signer common.Address
	Nonce  uint64
	TxHash [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPendingTransfer is a free log retrieval operation binding the contract event 0x798316e895cab65e6f329061235135f6f9440c7582bb336547a6d05185b131bc.
//
// Solidity: event PendingTransfer(address indexed signer, uint64 indexed nonce, bytes32 indexed txHash)
func (_PaymentManager *PaymentManagerFilterer) FilterPendingTransfer(opts *bind.FilterOpts, signer []common.Address, nonce []uint64, txHash [][32]byte) (*PaymentManagerPendingTransferIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentManager.contract.FilterLogs(opts, "PendingTransfer", signerRule, nonceRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentManagerPendingTransferIterator{contract: _PaymentManager.contract, event: "PendingTransfer", logs: logs, sub: sub}, nil
}

// WatchPendingTransfer is a free log subscription operation binding the contract event 0x798316e895cab65e6f329061235135f6f9440c7582bb336547a6d05185b131bc.
//
// Solidity: event PendingTransfer(address indexed signer, uint64 indexed nonce, bytes32 indexed txHash)
func (_PaymentManager *PaymentManagerFilterer) WatchPendingTransfer(opts *bind.WatchOpts, sink chan<- *PaymentManagerPendingTransfer, signer []common.Address, nonce []uint64, txHash [][32]byte) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentManager.contract.WatchLogs(opts, "PendingTransfer", signerRule, nonceRule, txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentManagerPendingTransfer)
				if err := _PaymentManager.contract.UnpackLog(event, "PendingTransfer", log); err != nil {
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

// ParsePendingTransfer is a log parse operation binding the contract event 0x798316e895cab65e6f329061235135f6f9440c7582bb336547a6d05185b131bc.
//
// Solidity: event PendingTransfer(address indexed signer, uint64 indexed nonce, bytes32 indexed txHash)
func (_PaymentManager *PaymentManagerFilterer) ParsePendingTransfer(log types.Log) (*PaymentManagerPendingTransfer, error) {
	event := new(PaymentManagerPendingTransfer)
	if err := _PaymentManager.contract.UnpackLog(event, "PendingTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentManagerRetryIterator is returned from FilterRetry and is used to iterate over the raw logs and unpacked data for Retry events raised by the PaymentManager contract.
type PaymentManagerRetryIterator struct {
	Event *PaymentManagerRetry // Event containing the contract specifics and raw log

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
func (it *PaymentManagerRetryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentManagerRetry)
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
		it.Event = new(PaymentManagerRetry)
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
func (it *PaymentManagerRetryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentManagerRetryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentManagerRetry represents a Retry event raised by the PaymentManager contract.
type PaymentManagerRetry struct {
	TxHash [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRetry is a free log retrieval operation binding the contract event 0x42ea0d481e8a8b75278dcea1de885ff3b50763d9026c0556046be840ade3ce1b.
//
// Solidity: event Retry(bytes32 indexed txHash)
func (_PaymentManager *PaymentManagerFilterer) FilterRetry(opts *bind.FilterOpts, txHash [][32]byte) (*PaymentManagerRetryIterator, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentManager.contract.FilterLogs(opts, "Retry", txHashRule)
	if err != nil {
		return nil, err
	}
	return &PaymentManagerRetryIterator{contract: _PaymentManager.contract, event: "Retry", logs: logs, sub: sub}, nil
}

// WatchRetry is a free log subscription operation binding the contract event 0x42ea0d481e8a8b75278dcea1de885ff3b50763d9026c0556046be840ade3ce1b.
//
// Solidity: event Retry(bytes32 indexed txHash)
func (_PaymentManager *PaymentManagerFilterer) WatchRetry(opts *bind.WatchOpts, sink chan<- *PaymentManagerRetry, txHash [][32]byte) (event.Subscription, error) {

	var txHashRule []interface{}
	for _, txHashItem := range txHash {
		txHashRule = append(txHashRule, txHashItem)
	}

	logs, sub, err := _PaymentManager.contract.WatchLogs(opts, "Retry", txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentManagerRetry)
				if err := _PaymentManager.contract.UnpackLog(event, "Retry", log); err != nil {
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

// ParseRetry is a log parse operation binding the contract event 0x42ea0d481e8a8b75278dcea1de885ff3b50763d9026c0556046be840ade3ce1b.
//
// Solidity: event Retry(bytes32 indexed txHash)
func (_PaymentManager *PaymentManagerFilterer) ParseRetry(log types.Log) (*PaymentManagerRetry, error) {
	event := new(PaymentManagerRetry)
	if err := _PaymentManager.contract.UnpackLog(event, "Retry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentManagerTxFailedIterator is returned from FilterTxFailed and is used to iterate over the raw logs and unpacked data for TxFailed events raised by the PaymentManager contract.
type PaymentManagerTxFailedIterator struct {
	Event *PaymentManagerTxFailed // Event containing the contract specifics and raw log

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
func (it *PaymentManagerTxFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentManagerTxFailed)
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
		it.Event = new(PaymentManagerTxFailed)
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
func (it *PaymentManagerTxFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentManagerTxFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentManagerTxFailed represents a TxFailed event raised by the PaymentManager contract.
type PaymentManagerTxFailed struct {
	Local   [32]byte
	Foreign [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTxFailed is a free log retrieval operation binding the contract event 0x4408cad08d2a774cfff1d2d5a614fd6bb763a24423b429da80f91eba6bdcb53b.
//
// Solidity: event TxFailed(bytes32 indexed local, bytes32 indexed foreign)
func (_PaymentManager *PaymentManagerFilterer) FilterTxFailed(opts *bind.FilterOpts, local [][32]byte, foreign [][32]byte) (*PaymentManagerTxFailedIterator, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var foreignRule []interface{}
	for _, foreignItem := range foreign {
		foreignRule = append(foreignRule, foreignItem)
	}

	logs, sub, err := _PaymentManager.contract.FilterLogs(opts, "TxFailed", localRule, foreignRule)
	if err != nil {
		return nil, err
	}
	return &PaymentManagerTxFailedIterator{contract: _PaymentManager.contract, event: "TxFailed", logs: logs, sub: sub}, nil
}

// WatchTxFailed is a free log subscription operation binding the contract event 0x4408cad08d2a774cfff1d2d5a614fd6bb763a24423b429da80f91eba6bdcb53b.
//
// Solidity: event TxFailed(bytes32 indexed local, bytes32 indexed foreign)
func (_PaymentManager *PaymentManagerFilterer) WatchTxFailed(opts *bind.WatchOpts, sink chan<- *PaymentManagerTxFailed, local [][32]byte, foreign [][32]byte) (event.Subscription, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var foreignRule []interface{}
	for _, foreignItem := range foreign {
		foreignRule = append(foreignRule, foreignItem)
	}

	logs, sub, err := _PaymentManager.contract.WatchLogs(opts, "TxFailed", localRule, foreignRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentManagerTxFailed)
				if err := _PaymentManager.contract.UnpackLog(event, "TxFailed", log); err != nil {
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

// ParseTxFailed is a log parse operation binding the contract event 0x4408cad08d2a774cfff1d2d5a614fd6bb763a24423b429da80f91eba6bdcb53b.
//
// Solidity: event TxFailed(bytes32 indexed local, bytes32 indexed foreign)
func (_PaymentManager *PaymentManagerFilterer) ParseTxFailed(log types.Log) (*PaymentManagerTxFailed, error) {
	event := new(PaymentManagerTxFailed)
	if err := _PaymentManager.contract.UnpackLog(event, "TxFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentManagerTxSuccessIterator is returned from FilterTxSuccess and is used to iterate over the raw logs and unpacked data for TxSuccess events raised by the PaymentManager contract.
type PaymentManagerTxSuccessIterator struct {
	Event *PaymentManagerTxSuccess // Event containing the contract specifics and raw log

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
func (it *PaymentManagerTxSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentManagerTxSuccess)
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
		it.Event = new(PaymentManagerTxSuccess)
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
func (it *PaymentManagerTxSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentManagerTxSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentManagerTxSuccess represents a TxSuccess event raised by the PaymentManager contract.
type PaymentManagerTxSuccess struct {
	Local   [32]byte
	Foreign [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTxSuccess is a free log retrieval operation binding the contract event 0xa2ce775ebcb679a2560ac924613e60823422156883902e37acbe857d09dc0309.
//
// Solidity: event TxSuccess(bytes32 indexed local, bytes32 indexed foreign)
func (_PaymentManager *PaymentManagerFilterer) FilterTxSuccess(opts *bind.FilterOpts, local [][32]byte, foreign [][32]byte) (*PaymentManagerTxSuccessIterator, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var foreignRule []interface{}
	for _, foreignItem := range foreign {
		foreignRule = append(foreignRule, foreignItem)
	}

	logs, sub, err := _PaymentManager.contract.FilterLogs(opts, "TxSuccess", localRule, foreignRule)
	if err != nil {
		return nil, err
	}
	return &PaymentManagerTxSuccessIterator{contract: _PaymentManager.contract, event: "TxSuccess", logs: logs, sub: sub}, nil
}

// WatchTxSuccess is a free log subscription operation binding the contract event 0xa2ce775ebcb679a2560ac924613e60823422156883902e37acbe857d09dc0309.
//
// Solidity: event TxSuccess(bytes32 indexed local, bytes32 indexed foreign)
func (_PaymentManager *PaymentManagerFilterer) WatchTxSuccess(opts *bind.WatchOpts, sink chan<- *PaymentManagerTxSuccess, local [][32]byte, foreign [][32]byte) (event.Subscription, error) {

	var localRule []interface{}
	for _, localItem := range local {
		localRule = append(localRule, localItem)
	}
	var foreignRule []interface{}
	for _, foreignItem := range foreign {
		foreignRule = append(foreignRule, foreignItem)
	}

	logs, sub, err := _PaymentManager.contract.WatchLogs(opts, "TxSuccess", localRule, foreignRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentManagerTxSuccess)
				if err := _PaymentManager.contract.UnpackLog(event, "TxSuccess", log); err != nil {
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

// ParseTxSuccess is a log parse operation binding the contract event 0xa2ce775ebcb679a2560ac924613e60823422156883902e37acbe857d09dc0309.
//
// Solidity: event TxSuccess(bytes32 indexed local, bytes32 indexed foreign)
func (_PaymentManager *PaymentManagerFilterer) ParseTxSuccess(log types.Log) (*PaymentManagerTxSuccess, error) {
	event := new(PaymentManagerTxSuccess)
	if err := _PaymentManager.contract.UnpackLog(event, "TxSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
