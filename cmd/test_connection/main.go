package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Get and verify environment variables
    providerURL := os.Getenv("BLOCKCHAIN_PROVIDER")
    privateKey := os.Getenv("BLOCKCHAIN_PRIVATE_KEY")
    
    fmt.Printf("Provider URL: %s\n", providerURL)
    fmt.Printf("Private Key Length: %d\n", len(privateKey))

    // Try to connect to Ganache
    client, err := ethclient.Dial(providerURL)
    if err != nil {
        log.Fatalf("Failed to connect to blockchain: %v", err)
    }

    // Try to parse private key
    _, err = crypto.HexToECDSA(privateKey)
    if err != nil {
        log.Fatalf("Failed to parse private key: %v", err)
    }

    // Get blockchain number to verify connection
    blockNumber, err := client.BlockNumber(context.Background())
    if err != nil {
        log.Fatalf("Failed to get block number: %v", err)
    }

    fmt.Println("Successfully connected to blockchain!")
    fmt.Printf("Current block number: %d\n", blockNumber)
}