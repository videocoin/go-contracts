package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/videocoin/go-contracts/resolver"
)

type config struct {
	chainRPC string `mapstructure:"chain_uri"`

	registryAddr string `mapstructure:"contract_addr"`

	contractName string `mapstructure:"contract_name"`
}

func defaultConfig() *config {
	return &config{
		chainRPC:     "http://127.0.0.1:8545",
		registryAddr: "0x3E4B9Ff6223C974A998bA2E314C587c080543B8c",
		contractName: "StreamManager",
	}
}

func main() {
	config := defaultConfig()
	client, err := ethclient.Dial(config.chainRPC)
	if err != nil {
		log.Fatalf("Failed to dial %v", err)
	}
	defer client.Close()

	e := resolver.NewResolver(client, config.registryAddr)

	versions, err := e.ContractVersions(config.contractName)
	if err != nil {
		log.Fatalf("Failed to get records: %v", err)
	}

	fmt.Printf("Available versions for contract %s: %v\n", config.contractName, versions)

	for i := 0; i < len(versions); i++ {
		info, err := e.ContractInfo(config.contractName, versions[i])
		if err != nil {
			log.Fatalf("Failed to get record: %v", err)
		}

		fmt.Printf("%s version %s address %s\n", config.contractName, versions[i], common.Bytes2Hex(info.Account.Bytes()))
	}
}
