package resolver

import (
	"fmt"

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
	registry, err := registry.NewRegistry(common.HexToAddress(e.registryAddress), e.client)
	if err != nil {
		return nil, err
	}

	versions, err := registry.Versions(nil, name)
	if err != nil {
		return nil, err
	}

	return versions, nil
}
