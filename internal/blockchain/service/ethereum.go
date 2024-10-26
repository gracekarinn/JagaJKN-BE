package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"jagajkn/internal/config"
)

type BlockchainService struct {
    client     *ethclient.Client
    privateKey *ecdsa.PrivateKey
}

func NewBlockchainService(cfg *config.BlockchainConfig) (*BlockchainService, error) {
    if cfg.ProviderURL == "" {
        return nil, fmt.Errorf("provider URL is empty")
    }
    
    client, err := ethclient.Dial(cfg.ProviderURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to blockchain: %v", err)
    }

    privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
    if err != nil {
        return nil, fmt.Errorf("failed to load private key: %v", err)
    }

    return &BlockchainService{
        client:     client,
        privateKey: privateKey,
    }, nil
}

func (s *BlockchainService) TestConnection() error {
    blockNumber, err := s.client.BlockNumber(context.Background())
    if err != nil {
        return fmt.Errorf("failed to get block number: %v", err)
    }
    
    log.Printf("Connected to blockchain. Latest block: %d", blockNumber)
    return nil
}