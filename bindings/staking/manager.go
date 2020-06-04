// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingManagerUnbondingRequest is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerUnbondingRequest struct {
	Transcoder common.Address
	Timestamp  *big.Int
	Amount     *big.Int
}

// StakingManagerABI is the input ABI used to generate the binding from.
const StakingManagerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_transcoderApprovalPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashRate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_slashPoolAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"Jailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"ManagerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"TranscoderRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"readiness\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnbondingRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"}],\"name\":\"Unjailed\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"delegateManaged\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pending\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"next\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"managed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"getSlashableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"getTrancoderSlashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"getTranscoderState\",\"outputs\":[{\"internalType\":\"enumStakingManager.TranscoderState\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unbondingID\",\"type\":\"uint256\"}],\"name\":\"getUnbondingRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"transcoder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingManager.UnbondingRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"isJailed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"}],\"name\":\"isManaged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"isManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minSelfStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingWithdrawalsExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"registerTranscoder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"}],\"name\":\"removeManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatorAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestUnbondingManaged\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setApprovalPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"}],\"name\":\"setCapacity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setSelfMinStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setSlashFundAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zone\",\"type\":\"uint256\"}],\"name\":\"setZone\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transcoderApprovalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"transcoders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zone\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"effectiveMinSelfStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transcodersArray\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transcodersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unbondingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"transcoderAddr\",\"type\":\"address\"}],\"name\":\"unjail\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAllPending\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawPending\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StakingManagerBin is the compiled bytecode used for deploying new contracts.
var StakingManagerBin = "0x60806040523480156200001157600080fd5b50604051620051d9380380620051d9833981810160405262000037919081019062000435565b6000620000496200016860201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350856001819055508460028190555083600381905550826004819055508160058190555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200015c336200017060201b60201c565b505050505050620006a0565b600033905090565b620001806200022360201b60201c565b620001c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001b990620005d9565b60405180910390fd5b620001dd81600a6200028960201b620030001790919060201c565b8073ffffffffffffffffffffffffffffffffffffffff167f3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a60405160405180910390a250565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166200026d6200016860201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b6200029b82826200033c60201b60201c565b15620002de576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002d590620005b7565b60405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620003b0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003a790620005fb565b60405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60008151905062000418816200066c565b92915050565b6000815190506200042f8162000686565b92915050565b60008060008060008060c087890312156200044f57600080fd5b60006200045f89828a016200041e565b96505060206200047289828a016200041e565b95505060406200048589828a016200041e565b94505060606200049889828a016200041e565b9350506080620004ab89828a016200041e565b92505060a0620004be89828a0162000407565b9150509295509295509295565b6000620004da601f836200061d565b91507f526f6c65733a206163636f756e7420616c72656164792068617320726f6c65006000830152602082019050919050565b60006200051c6020836200061d565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b60006200055e6022836200061d565b91507f526f6c65733a206163636f756e7420697320746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006020820190508181036000830152620005d281620004cb565b9050919050565b60006020820190508181036000830152620005f4816200050d565b9050919050565b6000602082019050818103600083015262000616816200054f565b9050919050565b600082825260208201905092915050565b60006200063b8262000642565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b62000677816200062e565b81146200068357600080fd5b50565b620006918162000662565b81146200069d57600080fd5b50565b614b2980620006b06000396000f3fe60806040526004361061023b5760003560e01c80638d23fc611161012e578063c96be4cb116100ab578063ec810a0a1161006f578063ec810a0a146108a8578063f2fde38b146108e5578063f3ae24151461090e578063fa7601541461094b578063fece707d146109765761023b565b8063c96be4cb14610782578063e2dc17f6146107bf578063e341181d14610803578063e71824bc1461082e578063ec69a1bb1461086b5761023b565b80639e79b122116100f25780639e79b1221461069d578063a79e7263146106da578063ac18de4314610703578063c57cc2001461072c578063c5f530af146107575761023b565b80638d23fc61146105a25780638da5cb5b146105e15780638efc97a11461060c5780638f32d59b1461064957806396fb4a9c146106745761023b565b8063399f57c0116101bc5780635c19a95c116101805780635c19a95c146104f05780636cf6d6751461050c5780636db2890914610537578063715018a6146105745780637edbceb11461058b5761023b565b8063399f57c014610421578063449ecfe61461044a5780635028e2e114610473578063503074ef146104b05780635afd2faa146104d95761023b565b8063254124fa11610203578063254124fa1461033e57806326e348ba1461037b5780632c9f0f2e146103a45780632d06177a146103cd5780633939e608146103f65761023b565b8063029859921461024057806314bfb5271461026b57806315620cce146102a85780631e7ff8f6146102c4578063220bb14e14610301575b600080fd5b34801561024c57600080fd5b506102556109b3565b60405161026291906148bc565b60405180910390f35b34801561027757600080fd5b50610292600480360361028d9190810190613cbb565b6109b9565b60405161029f91906145a9565b60405180910390f35b6102c260048036036102bd9190810190613d0d565b610ace565b005b3480156102d057600080fd5b506102eb60048036036102e69190810190613cbb565b610b85565b6040516102f891906148bc565b60405180910390f35b34801561030d57600080fd5b5061032860048036036103239190810190613cbb565b610c40565b60405161033591906145a9565b60405180910390f35b34801561034a57600080fd5b5061036560048036036103609190810190613d49565b610c9e565b60405161037291906148bc565b60405180910390f35b34801561038757600080fd5b506103a2600480360361039d9190810190613dd4565b610d91565b005b3480156103b057600080fd5b506103cb60048036036103c69190810190613dd4565b610def565b005b3480156103d957600080fd5b506103f460048036036103ef9190810190613cbb565b610e40565b005b34801561040257600080fd5b5061040b610ee1565b60405161041891906148bc565b60405180910390f35b34801561042d57600080fd5b5061044860048036036104439190810190613dd4565b610ee7565b005b34801561045657600080fd5b50610471600480360361046c9190810190613cbb565b611084565b005b34801561047f57600080fd5b5061049a60048036036104959190810190613d0d565b611229565b6040516104a791906148bc565b60405180910390f35b3480156104bc57600080fd5b506104d760048036036104d29190810190613d98565b6113b3565b005b3480156104e557600080fd5b506104ee611492565b005b61050a60048036036105059190810190613cbb565b611583565b005b34801561051857600080fd5b50610521611626565b60405161052e91906148bc565b60405180910390f35b34801561054357600080fd5b5061055e60048036036105599190810190613d98565b61162c565b60405161056b91906148bc565b60405180910390f35b34801561058057600080fd5b506105896116d6565b005b34801561059757600080fd5b506105a06117dc565b005b3480156105ae57600080fd5b506105c960048036036105c49190810190613cbb565b6118d9565b6040516105d893929190614900565b60405180910390f35b3480156105ed57600080fd5b506105f6611910565b604051610603919061458e565b60405180910390f35b34801561061857600080fd5b50610633600480360361062e9190810190613d0d565b611939565b60405161064091906148bc565b60405180910390f35b34801561065557600080fd5b5061065e611afe565b60405161066b91906145a9565b60405180910390f35b34801561068057600080fd5b5061069b60048036036106969190810190613d98565b611b5c565b005b3480156106a957600080fd5b506106c460048036036106bf9190810190613d98565b611c3b565b6040516106d191906148a1565b60405180910390f35b3480156106e657600080fd5b5061070160048036036106fc9190810190613ce4565b611d19565b005b34801561070f57600080fd5b5061072a60048036036107259190810190613cbb565b611e35565b005b34801561073857600080fd5b50610741611ed6565b60405161074e91906145a9565b60405180910390f35b34801561076357600080fd5b5061076c611f7c565b60405161077991906148bc565b60405180910390f35b34801561078e57600080fd5b506107a960048036036107a49190810190613cbb565b611f82565b6040516107b691906145a9565b60405180910390f35b3480156107cb57600080fd5b506107e660048036036107e19190810190613cbb565b61222b565b6040516107fa989796959493929190614937565b60405180910390f35b34801561080f57600080fd5b50610818612280565b60405161082591906148bc565b60405180910390f35b34801561083a57600080fd5b5061085560048036036108509190810190613cbb565b612286565b60405161086291906145c4565b60405180910390f35b34801561087757600080fd5b50610892600480360361088d9190810190613cbb565b61255e565b60405161089f91906148bc565b60405180910390f35b3480156108b457600080fd5b506108cf60048036036108ca9190810190613dd4565b612669565b6040516108dc919061458e565b60405180910390f35b3480156108f157600080fd5b5061090c60048036036109079190810190613cbb565b6126a5565b005b34801561091a57600080fd5b5061093560048036036109309190810190613cbb565b6126f8565b60405161094291906145a9565b60405180910390f35b34801561095757600080fd5b50610960612715565b60405161096d91906148bc565b60405180910390f35b34801561098257600080fd5b5061099d60048036036109989190810190613cbb565b612722565b6040516109aa91906148bc565b60405180910390f35b60015481565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610a2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a21906147e1565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411610ab4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aab906146c1565b60405180910390fd5b8060080160009054906101000a900460ff16915050919050565b610ad7336126f8565b610b16576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0d90614881565b60405180910390fd5b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160050160006101000a81548160ff021916908315150217905550610b808383612735565b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610bf6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bed906147e1565b60405180910390fd5b600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050919050565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060050160009054906101000a900460ff16915050919050565b6000610ca9336126f8565b610ce8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cdf90614881565b60405180910390fd5b6000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060050160009054906101000a900460ff16610d7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d73906147a1565b60405180910390fd5b610d87858585612a99565b9150509392505050565b610d99611afe565b610dd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcf90614741565b60405180910390fd5b60008111610de557600080fd5b8060028190555050565b610df7611afe565b610e36576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2d90614741565b60405180910390fd5b8060038190555050565b610e48611afe565b610e87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7e90614741565b60405180910390fd5b610e9b81600a61300090919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167f3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a60405160405180910390a250565b60035481565b60648110610f2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f21906147c1565b60405180910390fd5b60003390506000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015414610fb9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb0906146a1565b60405180910390fd5b428160010181905550828160020181905550600254816009018190555060098290806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550508173ffffffffffffffffffffffffffffffffffffffff167f6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b60405160405180910390a2505050565b61108c611afe565b6110cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c290614741565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561113b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611132906147e1565b60405180910390fd5b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010154116111c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111bc90614801565b60405180910390fd5b60008160080160006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167ffa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff60405160405180910390a25050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561129a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611291906147e1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561130a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611301906147e1565b60405180910390fd5b60006113168484611939565b90506113aa81600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546130a890919063ffffffff16565b91505092915050565b6113bb611afe565b6113fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113f190614741565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411611484576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147b906146c1565b60405180910390fd5b818160050181905550505050565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050806004015481600301541061151f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151690614821565b60405180910390fd5b6000816003015490505b816004015481101561157e57600061154182336130f2565b90508061155057505050611581565b6115686001846003015461344090919063ffffffff16565b8360030181905550508080600101915050611529565b50505b565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060050160009054906101000a900460ff1615611618576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161160f90614701565b60405180910390fd5b6116228233612735565b5050565b60045481565b600080600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060050160009054906101000a900460ff16156116c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116b990614701565b60405180910390fd5b6116cd843385612a99565b91505092915050565b6116de611afe565b61171d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161171490614741565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060040154816003015410611869576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161186090614821565b60405180910390fd5b6118778160030154336130f2565b6118b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118ad90614621565b60405180910390fd5b6118ce6001826003015461344090919063ffffffff16565b816003018190555050565b60086020528060005260406000206000915090508060030154908060040154908060050160009054906101000a900460ff16905083565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160010160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000836006018054905090506000809050818310611a2b578095505050505050611af8565b60008390505b82811015611aee576000866006018281548110611a4a57fe5b90600052602060002090600202016001015490506000611ab4828860000160008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461349590919063ffffffff16565b9050611aca60648261350590919063ffffffff16565b9050611adf818561344090919063ffffffff16565b93505050806001019050611a31565b5080955050505050505b92915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16611b4061354f565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b611b64611afe565b611ba3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b9a90614741565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411611c2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c24906146c1565b60405180910390fd5b818160040181905550505050565b611c43613c45565b6000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060020160008481526020019081526020016000206040518060600160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820154815260200160028201548152505091505092915050565b611d21611afe565b611d60576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d5790614741565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611df1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611de890614861565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611e3d611afe565b611e7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e7390614741565b60405180910390fd5b611e9081600a61355790919063ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff167fef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd3160405160405180910390a250565b600080600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816003015490505b8160040154811015611f7257600082600201600083815260200190815260200160002090506004548160010154420310611f645760019350505050611f79565b508080600101915050611f24565b5060009150505b90565b60025481565b6000611f8c611afe565b611fcb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fc290614741565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561207e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612075906147e1565b60405180910390fd5b60008160010154116120c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120bc90614801565b60405180910390fd5b60006120d084612286565b9050600160048111156120df57fe5b8160048111156120eb57fe5b1415801561211057506003600481111561210157fe5b81600481111561210d57fe5b14155b1561212057600092505050612226565b600061213b600554846000015461349590919063ffffffff16565b905061215160648261350590919063ffffffff16565b905061216a8184600001546130a890919063ffffffff16565b83600001819055508260060160405180604001604052804281526020016005548152509080600181540180825580915050906001820390600052602060002090600202016000909192909190915060008201518160000155602082015181600101555050506121d8856135fe565b6005548573ffffffffffffffffffffffffffffffffffffffff167f4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd60405160405180910390a3600193505050505b919050565b60076020528060005260406000206000915090508060000154908060010154908060020154908060030154908060040154908060050154908060080160009054906101000a900460ff16908060090154905088565b60055481565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156122f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122ee906147e1565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008260010154141561239557600492505050612559565b8160080160009054906101000a900460ff16156123b757600292505050612559565b60035482600101544203106125525781600901548160000160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541061241d57600192505050612559565b60008090506000826003015490505b82600401548110156124e757600083600201600083815260200190815260200160002090508673ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146124ae57506124da565b6004548160010154420310156124d8576124d581600201548461344090919063ffffffff16565b92505b505b808060010191505061242c565b5060025461253f8360000160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548361344090919063ffffffff16565b106125505760039350505050612559565b505b6000925050505b919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156125cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125c6906147e1565b60405180910390fd5b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816001015411612659576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612650906146c1565b60405180910390fd5b8060060180549050915050919050565b6009818154811061267657fe5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6126ad611afe565b6126ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126e390614741565b60405180910390fd5b6126f5816137a3565b50565b600061270e82600a6138d190919063ffffffff16565b9050919050565b6000600980549050905090565b600061272e8283611229565b9050919050565b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000349050600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161415612830576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161282790614841565b60405180910390fd5b600154811015612875576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161286c90614781565b60405180910390fd5b60008260000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054141561297557826007018490806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505082600601805490508260010160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b61297f8585613999565b61299681846000015461344090919063ffffffff16565b83600001819055506129f2818360000160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461344090919063ffffffff16565b8260000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b60405160405180910390a45050505050565b60008073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415612b0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b0190614841565b60405180910390fd5b6000600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050612b9a8686613999565b8060000160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054841115612c1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c1590614681565b60405180910390fd5b6000612c2987612286565b9050612c7f858360000160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546130a890919063ffffffff16565b8260000160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612cdb8584600001546130a890919063ffffffff16565b8360000181905550600082600401549050612d046001846004015461344090919063ffffffff16565b836004018190555060006004811115612d1957fe5b826004811115612d2557fe5b1480612d47575060026004811115612d3957fe5b826004811115612d4557fe5b145b80612d7e57508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614155b15612ee1578773ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16827f6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5428a604051612de39291906148d7565b60405180910390a460405180606001604052808973ffffffffffffffffffffffffffffffffffffffff168152602001600454420381526020018781525083600201600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001015560408201518160020155905050612e9d81886130f2565b612edc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ed390614621565b60405180910390fd5b612ff2565b8773ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16827f6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da560045442018a604051612f459291906148d7565b60405180910390a460405180606001604052808973ffffffffffffffffffffffffffffffffffffffff1681526020014281526020018781525083600201600083815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201559050505b809450505050509392505050565b61300a82826138d1565b1561304a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161304190614601565b60405180910390fd5b60018260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60006130ea83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613b89565b905092915050565b600080600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600201600086815260200190815260200160002090506000816002015414156131675760009250505061343a565b6004548160010154420310156131825760009250505061343a565b6000600760008360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008260010154905060006132096004548361344090919063ffffffff16565b9050600080905060008090505b84600601805490508110156132e057600085600601828154811061323657fe5b906000526020600020906002020190506000816000015490508581108061325c57508481115b156132685750506132d5565b600061328583600101548a6002015461349590919063ffffffff16565b905061329b60648261350590919063ffffffff16565b90506132b4818a600201546130a890919063ffffffff16565b89600201819055506132cf818661344090919063ffffffff16565b94505050505b806001019050613216565b50600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015613349573d6000803e3d6000fd5b50600085600201549050600086600201819055503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156133a3573d6000803e3d6000fd5b508560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168b7f544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f848460405161342691906148bc565b60405180910390a460019750505050505050505b92915050565b60008082840190508381101561348b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161348290614661565b60405180910390fd5b8091505092915050565b6000808314156134a857600090506134ff565b60008284029050828482816134b957fe5b04146134fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134f190614721565b60405180910390fd5b809150505b92915050565b600061354783836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250613be4565b905092915050565b600033905090565b61356182826138d1565b6135a0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613597906146e1565b60405180910390fd5b60008260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b613606611afe565b613645576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161363c90614741565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156136b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136ac906147e1565b60405180910390fd5b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600081600101541161373f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161373690614801565b60405180910390fd5b60018160080160006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e760405160405180910390a25050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415613813576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161380a90614641565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415613942576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161393990614761565b60405180910390fd5b8260000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000826006018054905090506000613a378686611939565b9050613a8d818460000160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546130a890919063ffffffff16565b8360000160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550818360010160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015613b80573d6000803e3d6000fd5b50505050505050565b6000838311158290613bd1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613bc891906145df565b60405180910390fd5b5060008385039050809150509392505050565b60008083118290613c2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c2291906145df565b60405180910390fd5b506000838581613c3757fe5b049050809150509392505050565b6040518060600160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b600081359050613c8b81614aa1565b92915050565b600081359050613ca081614ab8565b92915050565b600081359050613cb581614acf565b92915050565b600060208284031215613ccd57600080fd5b6000613cdb84828501613c7c565b91505092915050565b600060208284031215613cf657600080fd5b6000613d0484828501613c91565b91505092915050565b60008060408385031215613d2057600080fd5b6000613d2e85828601613c7c565b9250506020613d3f85828601613c7c565b9150509250929050565b600080600060608486031215613d5e57600080fd5b6000613d6c86828701613c7c565b9350506020613d7d86828701613c7c565b9250506040613d8e86828701613ca6565b9150509250925092565b60008060408385031215613dab57600080fd5b6000613db985828601613c7c565b9250506020613dca85828601613ca6565b9150509250929050565b600060208284031215613de657600080fd5b6000613df484828501613ca6565b91505092915050565b613e06816149d1565b82525050565b613e15816149d1565b82525050565b613e24816149f5565b82525050565b613e3381614a3e565b82525050565b6000613e44826149b5565b613e4e81856149c0565b9350613e5e818560208601614a50565b613e6781614a83565b840191505092915050565b6000613e7f601f836149c0565b91507f526f6c65733a206163636f756e7420616c72656164792068617320726f6c65006000830152602082019050919050565b6000613ebf6018836149c0565b91507f6661696c656420746f207769746864726177207374616b6500000000000000006000830152602082019050919050565b6000613eff6026836149c0565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000613f65601b836149c0565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b6000613fa56010836149c0565b91507f4e6f7420656e6f7567682066756e6473000000000000000000000000000000006000830152602082019050919050565b6000613fe5601d836149c0565b91507f5472616e73636f64657220616c726561647920726567697374657265640000006000830152602082019050919050565b60006140256019836149c0565b91507f5472616e73636f646572206e6f742072656769737465726564000000000000006000830152602082019050919050565b60006140656021836149c0565b91507f526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c60008301527f65000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006140cb6042836149c0565b91507f74686973206d6574686f642063616e277420626520757365642062792064656c60008301527f656761746f722074686174206465706f736974656420455243323020746f6b6560208301527f6e730000000000000000000000000000000000000000000000000000000000006040830152606082019050919050565b60006141576021836149c0565b91507f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f60008301527f77000000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006141bd6020836149c0565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b60006141fd6022836149c0565b91507f526f6c65733a206163636f756e7420697320746865207a65726f20616464726560008301527f73730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006142636023836149c0565b91507f4d757374206465706f736974206174206c65617374206d696e696d756d20766160008301527f6c756500000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006142c9604b836149c0565b91507f74686973206d6574686f642063616e206f6e6c792062652075736564206f6e6c60008301527f7920666f722064656c656761746f722074686174206465706f7369746564204560208301527f5243323020746f6b656e730000000000000000000000000000000000000000006040830152606082019050919050565b6000614355602b836149c0565b91507f52617465206d75737420626520612070657263656e746167652062657477656560008301527f6e203020616e64203130300000000000000000000000000000000000000000006020830152604082019050919050565b60006143bb6016836149c0565b91507f63616e277420757365207a65726f2061646472657373000000000000000000006000830152602082019050919050565b60006143fb601a836149c0565b91507f52656769737465726564207472616e73636f646572206f6e6c790000000000006000830152602082019050919050565b600061443b6013836149c0565b91507f6e6f2070656e64696e67207265717565737473000000000000000000000000006000830152602082019050919050565b600061447b6015836149c0565b91507f43616e60742075736520616464726573732030783000000000000000000000006000830152602082019050919050565b60006144bb601b836149c0565b91507f416c72656164792073657420746f2074686973206164647265737300000000006000830152602082019050919050565b60006144fb600d836149c0565b91507f6e6f742061206d616e61676572000000000000000000000000000000000000006000830152602082019050919050565b6060820160008201516145446000850182613dfd565b5060208201516145576020850182614570565b50604082015161456a6040850182614570565b50505050565b61457981614a34565b82525050565b61458881614a34565b82525050565b60006020820190506145a36000830184613e0c565b92915050565b60006020820190506145be6000830184613e1b565b92915050565b60006020820190506145d96000830184613e2a565b92915050565b600060208201905081810360008301526145f98184613e39565b905092915050565b6000602082019050818103600083015261461a81613e72565b9050919050565b6000602082019050818103600083015261463a81613eb2565b9050919050565b6000602082019050818103600083015261465a81613ef2565b9050919050565b6000602082019050818103600083015261467a81613f58565b9050919050565b6000602082019050818103600083015261469a81613f98565b9050919050565b600060208201905081810360008301526146ba81613fd8565b9050919050565b600060208201905081810360008301526146da81614018565b9050919050565b600060208201905081810360008301526146fa81614058565b9050919050565b6000602082019050818103600083015261471a816140be565b9050919050565b6000602082019050818103600083015261473a8161414a565b9050919050565b6000602082019050818103600083015261475a816141b0565b9050919050565b6000602082019050818103600083015261477a816141f0565b9050919050565b6000602082019050818103600083015261479a81614256565b9050919050565b600060208201905081810360008301526147ba816142bc565b9050919050565b600060208201905081810360008301526147da81614348565b9050919050565b600060208201905081810360008301526147fa816143ae565b9050919050565b6000602082019050818103600083015261481a816143ee565b9050919050565b6000602082019050818103600083015261483a8161442e565b9050919050565b6000602082019050818103600083015261485a8161446e565b9050919050565b6000602082019050818103600083015261487a816144ae565b9050919050565b6000602082019050818103600083015261489a816144ee565b9050919050565b60006060820190506148b6600083018461452e565b92915050565b60006020820190506148d1600083018461457f565b92915050565b60006040820190506148ec600083018561457f565b6148f9602083018461457f565b9392505050565b6000606082019050614915600083018661457f565b614922602083018561457f565b61492f6040830184613e1b565b949350505050565b60006101008201905061494d600083018b61457f565b61495a602083018a61457f565b614967604083018961457f565b614974606083018861457f565b614981608083018761457f565b61498e60a083018661457f565b61499b60c0830185613e1b565b6149a860e083018461457f565b9998505050505050505050565b600081519050919050565b600082825260208201905092915050565b60006149dc82614a14565b9050919050565b60006149ee82614a14565b9050919050565b60008115159050919050565b6000819050614a0f82614a94565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000614a4982614a01565b9050919050565b60005b83811015614a6e578082015181840152602081019050614a53565b83811115614a7d576000848401525b50505050565b6000601f19601f8301169050919050565b60058110614a9e57fe5b50565b614aaa816149d1565b8114614ab557600080fd5b50565b614ac1816149e3565b8114614acc57600080fd5b50565b614ad881614a34565b8114614ae357600080fd5b5056fea365627a7a723158204d6d1149e0140317652746cfa796d077a59eb6602310e1616fa2372b32aa21166c6578706572696d656e74616cf564736f6c634300050d0040"

// DeployStakingManager deploys a new Ethereum contract, binding an instance of StakingManager to it.
func DeployStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, _minDelegation *big.Int, _minSelfStake *big.Int, _transcoderApprovalPeriod *big.Int, _unbondingPeriod *big.Int, _slashRate *big.Int, _slashPoolAddress common.Address) (common.Address, *types.Transaction, *StakingManager, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakingManagerBin), backend, _minDelegation, _minSelfStake, _transcoderApprovalPeriod, _unbondingPeriod, _slashRate, _slashPoolAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerCaller) Delegators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	ret := new(struct {
		Pending *big.Int
		Next    *big.Int
		Managed bool
	})
	out := ret
	err := _StakingManager.contract.Call(opts, out, "delegators", arg0)
	return *ret, err
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerSession) Delegators(arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	return _StakingManager.Contract.Delegators(&_StakingManager.CallOpts, arg0)
}

// Delegators is a free data retrieval call binding the contract method 0x8d23fc61.
//
// Solidity: function delegators(address ) view returns(uint256 pending, uint256 next, bool managed)
func (_StakingManager *StakingManagerCallerSession) Delegators(arg0 common.Address) (struct {
	Pending *big.Int
	Next    *big.Int
	Managed bool
}, error) {
	return _StakingManager.Contract.Delegators(&_StakingManager.CallOpts, arg0)
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetDelegatorStake(opts *bind.CallOpts, transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getDelegatorStake", transcoderAddr, delegAddr)
	return *ret0, err
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetDelegatorStake(transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetDelegatorStake(&_StakingManager.CallOpts, transcoderAddr, delegAddr)
}

// GetDelegatorStake is a free data retrieval call binding the contract method 0x5028e2e1.
//
// Solidity: function getDelegatorStake(address transcoderAddr, address delegAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetDelegatorStake(transcoderAddr common.Address, delegAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetDelegatorStake(&_StakingManager.CallOpts, transcoderAddr, delegAddr)
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetSelfStake(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getSelfStake", _addr)
	return *ret0, err
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetSelfStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSelfStake(&_StakingManager.CallOpts, _addr)
}

// GetSelfStake is a free data retrieval call binding the contract method 0xfece707d.
//
// Solidity: function getSelfStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetSelfStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSelfStake(&_StakingManager.CallOpts, _addr)
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetSlashableAmount(opts *bind.CallOpts, transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getSlashableAmount", transcoderAddr, delegatorAddr)
	return *ret0, err
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetSlashableAmount(transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSlashableAmount(&_StakingManager.CallOpts, transcoderAddr, delegatorAddr)
}

// GetSlashableAmount is a free data retrieval call binding the contract method 0x8efc97a1.
//
// Solidity: function getSlashableAmount(address transcoderAddr, address delegatorAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetSlashableAmount(transcoderAddr common.Address, delegatorAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetSlashableAmount(&_StakingManager.CallOpts, transcoderAddr, delegatorAddr)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTotalStake(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTotalStake", _addr)
	return *ret0, err
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetTotalStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts, _addr)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _addr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTotalStake(_addr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTotalStake(&_StakingManager.CallOpts, _addr)
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetTrancoderSlashes(opts *bind.CallOpts, transcoderAddr common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTrancoderSlashes", transcoderAddr)
	return *ret0, err
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerSession) GetTrancoderSlashes(transcoderAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTrancoderSlashes(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTrancoderSlashes is a free data retrieval call binding the contract method 0xec69a1bb.
//
// Solidity: function getTrancoderSlashes(address transcoderAddr) view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetTrancoderSlashes(transcoderAddr common.Address) (*big.Int, error) {
	return _StakingManager.Contract.GetTrancoderSlashes(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerCaller) GetTranscoderState(opts *bind.CallOpts, transcoderAddr common.Address) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getTranscoderState", transcoderAddr)
	return *ret0, err
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerSession) GetTranscoderState(transcoderAddr common.Address) (uint8, error) {
	return _StakingManager.Contract.GetTranscoderState(&_StakingManager.CallOpts, transcoderAddr)
}

// GetTranscoderState is a free data retrieval call binding the contract method 0xe71824bc.
//
// Solidity: function getTranscoderState(address transcoderAddr) view returns(uint8)
func (_StakingManager *StakingManagerCallerSession) GetTranscoderState(transcoderAddr common.Address) (uint8, error) {
	return _StakingManager.Contract.GetTranscoderState(&_StakingManager.CallOpts, transcoderAddr)
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerCaller) GetUnbondingRequest(opts *bind.CallOpts, delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	var (
		ret0 = new(StakingManagerUnbondingRequest)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "getUnbondingRequest", delegatorAddr, unbondingID)
	return *ret0, err
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerSession) GetUnbondingRequest(delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	return _StakingManager.Contract.GetUnbondingRequest(&_StakingManager.CallOpts, delegatorAddr, unbondingID)
}

// GetUnbondingRequest is a free data retrieval call binding the contract method 0x9e79b122.
//
// Solidity: function getUnbondingRequest(address delegatorAddr, uint256 unbondingID) view returns(StakingManagerUnbondingRequest)
func (_StakingManager *StakingManagerCallerSession) GetUnbondingRequest(delegatorAddr common.Address, unbondingID *big.Int) (StakingManagerUnbondingRequest, error) {
	return _StakingManager.Contract.GetUnbondingRequest(&_StakingManager.CallOpts, delegatorAddr, unbondingID)
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsJailed(opts *bind.CallOpts, transcoderAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isJailed", transcoderAddr)
	return *ret0, err
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerSession) IsJailed(transcoderAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsJailed(&_StakingManager.CallOpts, transcoderAddr)
}

// IsJailed is a free data retrieval call binding the contract method 0x14bfb527.
//
// Solidity: function isJailed(address transcoderAddr) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsJailed(transcoderAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsJailed(&_StakingManager.CallOpts, transcoderAddr)
}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsManaged(opts *bind.CallOpts, delegatorAddr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isManaged", delegatorAddr)
	return *ret0, err
}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerSession) IsManaged(delegatorAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsManaged(&_StakingManager.CallOpts, delegatorAddr)
}

// IsManaged is a free data retrieval call binding the contract method 0x220bb14e.
//
// Solidity: function isManaged(address delegatorAddr) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsManaged(delegatorAddr common.Address) (bool, error) {
	return _StakingManager.Contract.IsManaged(&_StakingManager.CallOpts, delegatorAddr)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerCaller) IsManager(opts *bind.CallOpts, v common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isManager", v)
	return *ret0, err
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerSession) IsManager(v common.Address) (bool, error) {
	return _StakingManager.Contract.IsManager(&_StakingManager.CallOpts, v)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address v) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsManager(v common.Address) (bool, error) {
	return _StakingManager.Contract.IsManager(&_StakingManager.CallOpts, v)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakingManager *StakingManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakingManager *StakingManagerSession) IsOwner() (bool, error) {
	return _StakingManager.Contract.IsOwner(&_StakingManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) IsOwner() (bool, error) {
	return _StakingManager.Contract.IsOwner(&_StakingManager.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "minDelegation")
	return *ret0, err
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinDelegation() (*big.Int, error) {
	return _StakingManager.Contract.MinDelegation(&_StakingManager.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinDelegation() (*big.Int, error) {
	return _StakingManager.Contract.MinDelegation(&_StakingManager.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerCaller) MinSelfStake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "minSelfStake")
	return *ret0, err
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerSession) MinSelfStake() (*big.Int, error) {
	return _StakingManager.Contract.MinSelfStake(&_StakingManager.CallOpts)
}

// MinSelfStake is a free data retrieval call binding the contract method 0xc5f530af.
//
// Solidity: function minSelfStake() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) MinSelfStake() (*big.Int, error) {
	return _StakingManager.Contract.MinSelfStake(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingManager *StakingManagerCallerSession) Owner() (common.Address, error) {
	return _StakingManager.Contract.Owner(&_StakingManager.CallOpts)
}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerCaller) PendingWithdrawalsExist(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "pendingWithdrawalsExist")
	return *ret0, err
}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerSession) PendingWithdrawalsExist() (bool, error) {
	return _StakingManager.Contract.PendingWithdrawalsExist(&_StakingManager.CallOpts)
}

// PendingWithdrawalsExist is a free data retrieval call binding the contract method 0xc57cc200.
//
// Solidity: function pendingWithdrawalsExist() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) PendingWithdrawalsExist() (bool, error) {
	return _StakingManager.Contract.PendingWithdrawalsExist(&_StakingManager.CallOpts)
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerCaller) SlashRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "slashRate")
	return *ret0, err
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerSession) SlashRate() (*big.Int, error) {
	return _StakingManager.Contract.SlashRate(&_StakingManager.CallOpts)
}

// SlashRate is a free data retrieval call binding the contract method 0xe341181d.
//
// Solidity: function slashRate() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) SlashRate() (*big.Int, error) {
	return _StakingManager.Contract.SlashRate(&_StakingManager.CallOpts)
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TranscoderApprovalPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "transcoderApprovalPeriod")
	return *ret0, err
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerSession) TranscoderApprovalPeriod() (*big.Int, error) {
	return _StakingManager.Contract.TranscoderApprovalPeriod(&_StakingManager.CallOpts)
}

// TranscoderApprovalPeriod is a free data retrieval call binding the contract method 0x3939e608.
//
// Solidity: function transcoderApprovalPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TranscoderApprovalPeriod() (*big.Int, error) {
	return _StakingManager.Contract.TranscoderApprovalPeriod(&_StakingManager.CallOpts)
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerCaller) Transcoders(opts *bind.CallOpts, arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	ret := new(struct {
		Total                 *big.Int
		Timestamp             *big.Int
		RewardRate            *big.Int
		Rewards               *big.Int
		Zone                  *big.Int
		Capacity              *big.Int
		Jailed                bool
		EffectiveMinSelfStake *big.Int
	})
	out := ret
	err := _StakingManager.contract.Call(opts, out, "transcoders", arg0)
	return *ret, err
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerSession) Transcoders(arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	return _StakingManager.Contract.Transcoders(&_StakingManager.CallOpts, arg0)
}

// Transcoders is a free data retrieval call binding the contract method 0xe2dc17f6.
//
// Solidity: function transcoders(address ) view returns(uint256 total, uint256 timestamp, uint256 rewardRate, uint256 rewards, uint256 zone, uint256 capacity, bool jailed, uint256 effectiveMinSelfStake)
func (_StakingManager *StakingManagerCallerSession) Transcoders(arg0 common.Address) (struct {
	Total                 *big.Int
	Timestamp             *big.Int
	RewardRate            *big.Int
	Rewards               *big.Int
	Zone                  *big.Int
	Capacity              *big.Int
	Jailed                bool
	EffectiveMinSelfStake *big.Int
}, error) {
	return _StakingManager.Contract.Transcoders(&_StakingManager.CallOpts, arg0)
}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerCaller) TranscodersArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "transcodersArray", arg0)
	return *ret0, err
}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerSession) TranscodersArray(arg0 *big.Int) (common.Address, error) {
	return _StakingManager.Contract.TranscodersArray(&_StakingManager.CallOpts, arg0)
}

// TranscodersArray is a free data retrieval call binding the contract method 0xec810a0a.
//
// Solidity: function transcodersArray(uint256 ) view returns(address)
func (_StakingManager *StakingManagerCallerSession) TranscodersArray(arg0 *big.Int) (common.Address, error) {
	return _StakingManager.Contract.TranscodersArray(&_StakingManager.CallOpts, arg0)
}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerCaller) TranscodersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "transcodersCount")
	return *ret0, err
}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerSession) TranscodersCount() (*big.Int, error) {
	return _StakingManager.Contract.TranscodersCount(&_StakingManager.CallOpts)
}

// TranscodersCount is a free data retrieval call binding the contract method 0xfa760154.
//
// Solidity: function transcodersCount() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) TranscodersCount() (*big.Int, error) {
	return _StakingManager.Contract.TranscodersCount(&_StakingManager.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StakingManager.contract.Call(opts, out, "unbondingPeriod")
	return *ret0, err
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerSession) UnbondingPeriod() (*big.Int, error) {
	return _StakingManager.Contract.UnbondingPeriod(&_StakingManager.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _StakingManager.Contract.UnbondingPeriod(&_StakingManager.CallOpts)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerTransactor) AddManager(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "addManager", v)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerSession) AddManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.AddManager(&_StakingManager.TransactOpts, v)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address v) returns()
func (_StakingManager *StakingManagerTransactorSession) AddManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.AddManager(&_StakingManager.TransactOpts, v)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) payable returns()
func (_StakingManager *StakingManagerTransactor) Delegate(opts *bind.TransactOpts, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "delegate", transcoderAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) payable returns()
func (_StakingManager *StakingManagerSession) Delegate(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Delegate(&_StakingManager.TransactOpts, transcoderAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address transcoderAddr) payable returns()
func (_StakingManager *StakingManagerTransactorSession) Delegate(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Delegate(&_StakingManager.TransactOpts, transcoderAddr)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x15620cce.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr) payable returns()
func (_StakingManager *StakingManagerTransactor) DelegateManaged(opts *bind.TransactOpts, transcoderAddr common.Address, delegatorAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "delegateManaged", transcoderAddr, delegatorAddr)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x15620cce.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr) payable returns()
func (_StakingManager *StakingManagerSession) DelegateManaged(transcoderAddr common.Address, delegatorAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.DelegateManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr)
}

// DelegateManaged is a paid mutator transaction binding the contract method 0x15620cce.
//
// Solidity: function delegateManaged(address transcoderAddr, address delegatorAddr) payable returns()
func (_StakingManager *StakingManagerTransactorSession) DelegateManaged(transcoderAddr common.Address, delegatorAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.DelegateManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactor) RegisterTranscoder(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "registerTranscoder", rate)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerSession) RegisterTranscoder(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTranscoder(&_StakingManager.TransactOpts, rate)
}

// RegisterTranscoder is a paid mutator transaction binding the contract method 0x399f57c0.
//
// Solidity: function registerTranscoder(uint256 rate) returns()
func (_StakingManager *StakingManagerTransactorSession) RegisterTranscoder(rate *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RegisterTranscoder(&_StakingManager.TransactOpts, rate)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerTransactor) RemoveManager(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "removeManager", v)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerSession) RemoveManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RemoveManager(&_StakingManager.TransactOpts, v)
}

// RemoveManager is a paid mutator transaction binding the contract method 0xac18de43.
//
// Solidity: function removeManager(address v) returns()
func (_StakingManager *StakingManagerTransactorSession) RemoveManager(v common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RemoveManager(&_StakingManager.TransactOpts, v)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceOwnership(&_StakingManager.TransactOpts)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) RequestUnbonding(opts *bind.TransactOpts, transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "requestUnbonding", transcoderAddr, amount)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerSession) RequestUnbonding(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbonding(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// RequestUnbonding is a paid mutator transaction binding the contract method 0x6db28909.
//
// Solidity: function requestUnbonding(address transcoderAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) RequestUnbonding(transcoderAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbonding(&_StakingManager.TransactOpts, transcoderAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactor) RequestUnbondingManaged(opts *bind.TransactOpts, transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "requestUnbondingManaged", transcoderAddr, delegatorAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerSession) RequestUnbondingManaged(transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbondingManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr, amount)
}

// RequestUnbondingManaged is a paid mutator transaction binding the contract method 0x254124fa.
//
// Solidity: function requestUnbondingManaged(address transcoderAddr, address delegatorAddr, uint256 amount) returns(uint256)
func (_StakingManager *StakingManagerTransactorSession) RequestUnbondingManaged(transcoderAddr common.Address, delegatorAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnbondingManaged(&_StakingManager.TransactOpts, transcoderAddr, delegatorAddr, amount)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactor) SetApprovalPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setApprovalPeriod", period)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerSession) SetApprovalPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetApprovalPeriod(&_StakingManager.TransactOpts, period)
}

// SetApprovalPeriod is a paid mutator transaction binding the contract method 0x2c9f0f2e.
//
// Solidity: function setApprovalPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactorSession) SetApprovalPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetApprovalPeriod(&_StakingManager.TransactOpts, period)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerTransactor) SetCapacity(opts *bind.TransactOpts, addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setCapacity", addr, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerSession) SetCapacity(addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetCapacity(&_StakingManager.TransactOpts, addr, capacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x503074ef.
//
// Solidity: function setCapacity(address addr, uint256 capacity) returns()
func (_StakingManager *StakingManagerTransactorSession) SetCapacity(addr common.Address, capacity *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetCapacity(&_StakingManager.TransactOpts, addr, capacity)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerTransactor) SetSelfMinStake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSelfMinStake", amount)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerSession) SetSelfMinStake(amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSelfMinStake(&_StakingManager.TransactOpts, amount)
}

// SetSelfMinStake is a paid mutator transaction binding the contract method 0x26e348ba.
//
// Solidity: function setSelfMinStake(uint256 amount) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSelfMinStake(amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSelfMinStake(&_StakingManager.TransactOpts, amount)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerTransactor) SetSlashFundAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSlashFundAddress", addr)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerSession) SetSlashFundAddress(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashFundAddress(&_StakingManager.TransactOpts, addr)
}

// SetSlashFundAddress is a paid mutator transaction binding the contract method 0xa79e7263.
//
// Solidity: function setSlashFundAddress(address addr) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSlashFundAddress(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlashFundAddress(&_StakingManager.TransactOpts, addr)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerTransactor) SetZone(opts *bind.TransactOpts, addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setZone", addr, zone)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerSession) SetZone(addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetZone(&_StakingManager.TransactOpts, addr, zone)
}

// SetZone is a paid mutator transaction binding the contract method 0x96fb4a9c.
//
// Solidity: function setZone(address addr, uint256 zone) returns()
func (_StakingManager *StakingManagerTransactorSession) SetZone(addr common.Address, zone *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetZone(&_StakingManager.TransactOpts, addr, zone)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerTransactor) Slash(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "slash", addr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerSession) Slash(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, addr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address addr) returns(bool)
func (_StakingManager *StakingManagerTransactorSession) Slash(addr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingManager *StakingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.TransferOwnership(&_StakingManager.TransactOpts, newOwner)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactor) Unjail(opts *bind.TransactOpts, transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unjail", transcoderAddr)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerSession) Unjail(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Unjail(&_StakingManager.TransactOpts, transcoderAddr)
}

// Unjail is a paid mutator transaction binding the contract method 0x449ecfe6.
//
// Solidity: function unjail(address transcoderAddr) returns()
func (_StakingManager *StakingManagerTransactorSession) Unjail(transcoderAddr common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Unjail(&_StakingManager.TransactOpts, transcoderAddr)
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerTransactor) WithdrawAllPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawAllPending")
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerSession) WithdrawAllPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawAllPending(&_StakingManager.TransactOpts)
}

// WithdrawAllPending is a paid mutator transaction binding the contract method 0x5afd2faa.
//
// Solidity: function withdrawAllPending() returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawAllPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawAllPending(&_StakingManager.TransactOpts)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerTransactor) WithdrawPending(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdrawPending")
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerSession) WithdrawPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawPending(&_StakingManager.TransactOpts)
}

// WithdrawPending is a paid mutator transaction binding the contract method 0x7edbceb1.
//
// Solidity: function withdrawPending() returns()
func (_StakingManager *StakingManagerTransactorSession) WithdrawPending() (*types.Transaction, error) {
	return _StakingManager.Contract.WithdrawPending(&_StakingManager.TransactOpts)
}

// StakingManagerDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the StakingManager contract.
type StakingManagerDelegatedIterator struct {
	Event *StakingManagerDelegated // Event containing the contract specifics and raw log

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
func (it *StakingManagerDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerDelegated)
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
		it.Event = new(StakingManagerDelegated)
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
func (it *StakingManagerDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerDelegated represents a Delegated event raised by the StakingManager contract.
type StakingManagerDelegated struct {
	Transcoder common.Address
	Delegator  common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) FilterDelegated(opts *bind.FilterOpts, transcoder []common.Address, delegator []common.Address, amount []*big.Int) (*StakingManagerDelegatedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Delegated", transcoderRule, delegatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerDelegatedIterator{contract: _StakingManager.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *StakingManagerDelegated, transcoder []common.Address, delegator []common.Address, amount []*big.Int) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Delegated", transcoderRule, delegatorRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerDelegated)
				if err := _StakingManager.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed transcoder, address indexed delegator, uint256 indexed amount)
func (_StakingManager *StakingManagerFilterer) ParseDelegated(log types.Log) (*StakingManagerDelegated, error) {
	event := new(StakingManagerDelegated)
	if err := _StakingManager.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerJailedIterator is returned from FilterJailed and is used to iterate over the raw logs and unpacked data for Jailed events raised by the StakingManager contract.
type StakingManagerJailedIterator struct {
	Event *StakingManagerJailed // Event containing the contract specifics and raw log

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
func (it *StakingManagerJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerJailed)
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
		it.Event = new(StakingManagerJailed)
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
func (it *StakingManagerJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerJailed represents a Jailed event raised by the StakingManager contract.
type StakingManagerJailed struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterJailed is a free log retrieval operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterJailed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerJailedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Jailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerJailedIterator{contract: _StakingManager.contract, event: "Jailed", logs: logs, sub: sub}, nil
}

// WatchJailed is a free log subscription operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchJailed(opts *bind.WatchOpts, sink chan<- *StakingManagerJailed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Jailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerJailed)
				if err := _StakingManager.contract.UnpackLog(event, "Jailed", log); err != nil {
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

// ParseJailed is a log parse operation binding the contract event 0x519ec2af7b403e5bfa116afc87904cd6aa3e97a09cae81b522551191195674e7.
//
// Solidity: event Jailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseJailed(log types.Log) (*StakingManagerJailed, error) {
	event := new(StakingManagerJailed)
	if err := _StakingManager.contract.UnpackLog(event, "Jailed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerManagerAddedIterator is returned from FilterManagerAdded and is used to iterate over the raw logs and unpacked data for ManagerAdded events raised by the StakingManager contract.
type StakingManagerManagerAddedIterator struct {
	Event *StakingManagerManagerAdded // Event containing the contract specifics and raw log

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
func (it *StakingManagerManagerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerManagerAdded)
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
		it.Event = new(StakingManagerManagerAdded)
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
func (it *StakingManagerManagerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerManagerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerManagerAdded represents a ManagerAdded event raised by the StakingManager contract.
type StakingManagerManagerAdded struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerAdded is a free log retrieval operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) FilterManagerAdded(opts *bind.FilterOpts, manager []common.Address) (*StakingManagerManagerAddedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerManagerAddedIterator{contract: _StakingManager.contract, event: "ManagerAdded", logs: logs, sub: sub}, nil
}

// WatchManagerAdded is a free log subscription operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) WatchManagerAdded(opts *bind.WatchOpts, sink chan<- *StakingManagerManagerAdded, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ManagerAdded", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerManagerAdded)
				if err := _StakingManager.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
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

// ParseManagerAdded is a log parse operation binding the contract event 0x3b4a40cccf2058c593542587329dd385be4f0b588db5471fbd9598e56dd7093a.
//
// Solidity: event ManagerAdded(address indexed manager)
func (_StakingManager *StakingManagerFilterer) ParseManagerAdded(log types.Log) (*StakingManagerManagerAdded, error) {
	event := new(StakingManagerManagerAdded)
	if err := _StakingManager.contract.UnpackLog(event, "ManagerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerManagerRemovedIterator is returned from FilterManagerRemoved and is used to iterate over the raw logs and unpacked data for ManagerRemoved events raised by the StakingManager contract.
type StakingManagerManagerRemovedIterator struct {
	Event *StakingManagerManagerRemoved // Event containing the contract specifics and raw log

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
func (it *StakingManagerManagerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerManagerRemoved)
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
		it.Event = new(StakingManagerManagerRemoved)
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
func (it *StakingManagerManagerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerManagerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerManagerRemoved represents a ManagerRemoved event raised by the StakingManager contract.
type StakingManagerManagerRemoved struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerRemoved is a free log retrieval operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) FilterManagerRemoved(opts *bind.FilterOpts, manager []common.Address) (*StakingManagerManagerRemovedIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "ManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerManagerRemovedIterator{contract: _StakingManager.contract, event: "ManagerRemoved", logs: logs, sub: sub}, nil
}

// WatchManagerRemoved is a free log subscription operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) WatchManagerRemoved(opts *bind.WatchOpts, sink chan<- *StakingManagerManagerRemoved, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "ManagerRemoved", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerManagerRemoved)
				if err := _StakingManager.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
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

// ParseManagerRemoved is a log parse operation binding the contract event 0xef69f7d97228658c92417be1b16b19058315de71fecb435d07b7d23728b6bd31.
//
// Solidity: event ManagerRemoved(address indexed manager)
func (_StakingManager *StakingManagerFilterer) ParseManagerRemoved(log types.Log) (*StakingManagerManagerRemoved, error) {
	event := new(StakingManagerManagerRemoved)
	if err := _StakingManager.contract.UnpackLog(event, "ManagerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingManager contract.
type StakingManagerOwnershipTransferredIterator struct {
	Event *StakingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerOwnershipTransferred)
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
		it.Event = new(StakingManagerOwnershipTransferred)
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
func (it *StakingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the StakingManager contract.
type StakingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerOwnershipTransferredIterator{contract: _StakingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingManager *StakingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerOwnershipTransferred)
				if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*StakingManagerOwnershipTransferred, error) {
	event := new(StakingManagerOwnershipTransferred)
	if err := _StakingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the StakingManager contract.
type StakingManagerSlashedIterator struct {
	Event *StakingManagerSlashed // Event containing the contract specifics and raw log

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
func (it *StakingManagerSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerSlashed)
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
		it.Event = new(StakingManagerSlashed)
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
func (it *StakingManagerSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerSlashed represents a Slashed event raised by the StakingManager contract.
type StakingManagerSlashed struct {
	Transcoder common.Address
	Rate       *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) FilterSlashed(opts *bind.FilterOpts, transcoder []common.Address, rate []*big.Int) (*StakingManagerSlashedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Slashed", transcoderRule, rateRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerSlashedIterator{contract: _StakingManager.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakingManagerSlashed, transcoder []common.Address, rate []*big.Int) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}
	var rateRule []interface{}
	for _, rateItem := range rate {
		rateRule = append(rateRule, rateItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Slashed", transcoderRule, rateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerSlashed)
				if err := _StakingManager.contract.UnpackLog(event, "Slashed", log); err != nil {
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

// ParseSlashed is a log parse operation binding the contract event 0x4ed05e9673c26d2ed44f7ef6a7f2942df0ee3b5e1e17db4b99f9dcd261a339cd.
//
// Solidity: event Slashed(address indexed transcoder, uint256 indexed rate)
func (_StakingManager *StakingManagerFilterer) ParseSlashed(log types.Log) (*StakingManagerSlashed, error) {
	event := new(StakingManagerSlashed)
	if err := _StakingManager.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerStakeWithdrawalIterator is returned from FilterStakeWithdrawal and is used to iterate over the raw logs and unpacked data for StakeWithdrawal events raised by the StakingManager contract.
type StakingManagerStakeWithdrawalIterator struct {
	Event *StakingManagerStakeWithdrawal // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeWithdrawal)
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
		it.Event = new(StakingManagerStakeWithdrawal)
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
func (it *StakingManagerStakeWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeWithdrawal represents a StakeWithdrawal event raised by the StakingManager contract.
type StakingManagerStakeWithdrawal struct {
	UnbondingID *big.Int
	Delegator   common.Address
	Transcoder  common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawal is a free log retrieval operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterStakeWithdrawal(opts *bind.FilterOpts, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (*StakingManagerStakeWithdrawalIterator, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeWithdrawal", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeWithdrawalIterator{contract: _StakingManager.contract, event: "StakeWithdrawal", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawal is a free log subscription operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchStakeWithdrawal(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeWithdrawal, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (event.Subscription, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeWithdrawal", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeWithdrawal)
				if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
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

// ParseStakeWithdrawal is a log parse operation binding the contract event 0x544ab8c284dc3fe11e91e1be98918875679b41f64ade1d34b53fbfaab5e14f84.
//
// Solidity: event StakeWithdrawal(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseStakeWithdrawal(log types.Log) (*StakingManagerStakeWithdrawal, error) {
	event := new(StakingManagerStakeWithdrawal)
	if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerTranscoderRegisteredIterator is returned from FilterTranscoderRegistered and is used to iterate over the raw logs and unpacked data for TranscoderRegistered events raised by the StakingManager contract.
type StakingManagerTranscoderRegisteredIterator struct {
	Event *StakingManagerTranscoderRegistered // Event containing the contract specifics and raw log

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
func (it *StakingManagerTranscoderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerTranscoderRegistered)
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
		it.Event = new(StakingManagerTranscoderRegistered)
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
func (it *StakingManagerTranscoderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerTranscoderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerTranscoderRegistered represents a TranscoderRegistered event raised by the StakingManager contract.
type StakingManagerTranscoderRegistered struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTranscoderRegistered is a free log retrieval operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterTranscoderRegistered(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerTranscoderRegisteredIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "TranscoderRegistered", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTranscoderRegisteredIterator{contract: _StakingManager.contract, event: "TranscoderRegistered", logs: logs, sub: sub}, nil
}

// WatchTranscoderRegistered is a free log subscription operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchTranscoderRegistered(opts *bind.WatchOpts, sink chan<- *StakingManagerTranscoderRegistered, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "TranscoderRegistered", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerTranscoderRegistered)
				if err := _StakingManager.contract.UnpackLog(event, "TranscoderRegistered", log); err != nil {
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

// ParseTranscoderRegistered is a log parse operation binding the contract event 0x6fbcf0f12b438f90175bebf725f86a4a74d12525d5d2c144a68e400696bce58b.
//
// Solidity: event TranscoderRegistered(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseTranscoderRegistered(log types.Log) (*StakingManagerTranscoderRegistered, error) {
	event := new(StakingManagerTranscoderRegistered)
	if err := _StakingManager.contract.UnpackLog(event, "TranscoderRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerUnbondingRequestedIterator is returned from FilterUnbondingRequested and is used to iterate over the raw logs and unpacked data for UnbondingRequested events raised by the StakingManager contract.
type StakingManagerUnbondingRequestedIterator struct {
	Event *StakingManagerUnbondingRequested // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnbondingRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnbondingRequested)
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
		it.Event = new(StakingManagerUnbondingRequested)
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
func (it *StakingManagerUnbondingRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnbondingRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnbondingRequested represents a UnbondingRequested event raised by the StakingManager contract.
type StakingManagerUnbondingRequested struct {
	UnbondingID *big.Int
	Delegator   common.Address
	Transcoder  common.Address
	Readiness   *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUnbondingRequested is a free log retrieval operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterUnbondingRequested(opts *bind.FilterOpts, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (*StakingManagerUnbondingRequestedIterator, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnbondingRequested", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnbondingRequestedIterator{contract: _StakingManager.contract, event: "UnbondingRequested", logs: logs, sub: sub}, nil
}

// WatchUnbondingRequested is a free log subscription operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchUnbondingRequested(opts *bind.WatchOpts, sink chan<- *StakingManagerUnbondingRequested, unbondingID []*big.Int, delegator []common.Address, transcoder []common.Address) (event.Subscription, error) {

	var unbondingIDRule []interface{}
	for _, unbondingIDItem := range unbondingID {
		unbondingIDRule = append(unbondingIDRule, unbondingIDItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnbondingRequested", unbondingIDRule, delegatorRule, transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnbondingRequested)
				if err := _StakingManager.contract.UnpackLog(event, "UnbondingRequested", log); err != nil {
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

// ParseUnbondingRequested is a log parse operation binding the contract event 0x6f76dbe74f79ce91bd34ace12f43ca1064eeb3107135a5203603e1b7b6791da5.
//
// Solidity: event UnbondingRequested(uint256 indexed unbondingID, address indexed delegator, address indexed transcoder, uint256 readiness, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseUnbondingRequested(log types.Log) (*StakingManagerUnbondingRequested, error) {
	event := new(StakingManagerUnbondingRequested)
	if err := _StakingManager.contract.UnpackLog(event, "UnbondingRequested", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingManagerUnjailedIterator is returned from FilterUnjailed and is used to iterate over the raw logs and unpacked data for Unjailed events raised by the StakingManager contract.
type StakingManagerUnjailedIterator struct {
	Event *StakingManagerUnjailed // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnjailed)
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
		it.Event = new(StakingManagerUnjailed)
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
func (it *StakingManagerUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnjailed represents a Unjailed event raised by the StakingManager contract.
type StakingManagerUnjailed struct {
	Transcoder common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnjailed is a free log retrieval operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) FilterUnjailed(opts *bind.FilterOpts, transcoder []common.Address) (*StakingManagerUnjailedIterator, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Unjailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnjailedIterator{contract: _StakingManager.contract, event: "Unjailed", logs: logs, sub: sub}, nil
}

// WatchUnjailed is a free log subscription operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) WatchUnjailed(opts *bind.WatchOpts, sink chan<- *StakingManagerUnjailed, transcoder []common.Address) (event.Subscription, error) {

	var transcoderRule []interface{}
	for _, transcoderItem := range transcoder {
		transcoderRule = append(transcoderRule, transcoderItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Unjailed", transcoderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnjailed)
				if err := _StakingManager.contract.UnpackLog(event, "Unjailed", log); err != nil {
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

// ParseUnjailed is a log parse operation binding the contract event 0xfa5039497ad9ba11f0eb5239b2614e925541bbcc0cf3476dd68e1927c86d33ff.
//
// Solidity: event Unjailed(address indexed transcoder)
func (_StakingManager *StakingManagerFilterer) ParseUnjailed(log types.Log) (*StakingManagerUnjailed, error) {
	event := new(StakingManagerUnjailed)
	if err := _StakingManager.contract.UnpackLog(event, "Unjailed", log); err != nil {
		return nil, err
	}
	return event, nil
}
