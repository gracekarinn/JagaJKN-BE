package config

import (
	"os"
)

type BlockchainConfig struct {
    ProviderURL string
    PrivateKey  string
}

func NewBlockchainConfig() *BlockchainConfig {
    return &BlockchainConfig{
        ProviderURL: os.Getenv("BLOCKCHAIN_PROVIDER"),
        PrivateKey:  os.Getenv("BLOCKCHAIN_PRIVATE_KEY"),
    }
}