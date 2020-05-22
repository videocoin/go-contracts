package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/videocoin/go-contracts/bindings/registry"
)

type config struct {
	chainRPC string `mapstructure:"chain_uri"`

	registryAddr string `mapstructure:"contract_addr"`
}

func defaultConfig() *config {
	return &config{
		chainRPC:     "http://127.0.0.1:8545",
		registryAddr: "0xd67abbdC587549AF6Bf8F62361c1d91116aE6D77",
	}
}

func main() {
	config := defaultConfig()
	client, err := ethclient.Dial(config.chainRPC)
	if err != nil {
		log.Fatalf("Failed to dial %v", err)
	}
	defer client.Close()

	registry, err := registry.NewRegistry(common.HexToAddress(config.registryAddr), client)

	version, err := registry.Versions(&bind.CallOpts{}, "StreamManager", new(big.Int).SetUint64(0))
	if err != nil {
		log.Fatalf("Failed to get version %v", err)
	}

	fmt.Printf("Version is: %s\n", version)

	record, err := registry.Record(nil, "StreamManager", version)
	if err != nil {
		log.Fatalf("Failed to get records: %v", err)
	}

	fmt.Printf("Address %s\n", common.Bytes2Hex(record.Account.Bytes()))
	fmt.Printf("  Owner %s\n", common.Bytes2Hex(record.Owner.Bytes()))
}
