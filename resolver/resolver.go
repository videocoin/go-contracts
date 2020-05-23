package resolver

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/videocoin/go-contracts/bindings/registry"

	"github.com/videocoin/common/logging"
)

// NewResolver creates resolver object
func NewResolver(client bind.ContractBackend, registryAddress string) Resolver {
	return enforcer{
		client:          client,
		registryAddress: registryAddress,
	}
}

// Resolver is a contract wrapper to read contract data from registry
type Resolver interface {
	ContractInfo(name string, version string) (*registry.RegistryContractInfo, error)
	ContractVersions(name string) ([]string, error)
}

type enforcer struct {
	log             logging.Interface
	client          bind.ContractBackend
	registryAddress string
}

func (e enforcer) ContractInfo(name string, version string) (*registry.RegistryContractInfo, error) {
	registry, err := registry.NewRegistry(common.HexToAddress(e.registryAddress), e.client)
	if err != nil {
		return nil, err
	}

	record, err := registry.Record(nil, name, version)
	if err != nil {
		return nil, err
	}

	if !record.Initialized {
		return nil, fmt.Errorf("contract %s with version %s is not registered", name, version)
	}

	return &record, nil
}

func (e enforcer) ContractVersions(name string) ([]string, error) {
	var ret []string
	registry, err := registry.NewRegistry(common.HexToAddress(e.registryAddress), e.client)
	if err != nil {
		return ret, err
	}

	length, err := registry.Versions(nil, name)
	if err != nil {
		return ret, err
	}

	var i int64
	for i = 0; i < length.Int64(); i++ {
		version, err := registry.Version(nil, name, new(big.Int).SetInt64(i))
		if err != nil {
			return ret, err
		}
		ret = append(ret, version)
	}

	return ret, nil
}
