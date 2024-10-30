package main

import (
	"context"
	"log"
	"os"
	"time"

	"jagajkn/internal/blockchain/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    log.Println("Connecting to blockchain...")
    client, err := ethclient.Dial(os.Getenv("BLOCKCHAIN_PROVIDER"))
    if err != nil {
        log.Fatalf("Failed to connect to blockchain: %v", err)
    }

    // Get chain ID
    chainID, err := client.ChainID(context.Background())
    if err != nil {
        log.Fatalf("Failed to get chain ID: %v", err)
    }
    log.Printf("Connected to chain ID: %s", chainID.String())

    // Load private key
    privateKey, err := crypto.HexToECDSA(os.Getenv("BLOCKCHAIN_PRIVATE_KEY"))
    if err != nil {
        log.Fatalf("Failed to load private key: %v", err)
    }

    // Create auth
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        log.Fatalf("Failed to create auth: %v", err)
    }

    // Set gas price and limit
    auth.GasLimit = uint64(5000000)
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatalf("Failed to get gas price: %v", err)
    }
    auth.GasPrice = gasPrice

    log.Println("Deploying contract...")
    address, tx, instance, err := contracts.DeployContracts(auth, client)
    if err != nil {
        log.Fatalf("Failed to deploy contract: %v", err)
    }
    log.Printf("Contract deploying in transaction: %s", tx.Hash().Hex())

    // Wait for deployment
    ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
    defer cancel()
    _, err = bind.WaitMined(ctx, client, tx)
    if err != nil {
        log.Fatalf("Failed to wait for deployment: %v", err)
    }

    // Verify contract code
    code, err := client.CodeAt(context.Background(), address, nil)
    if err != nil {
        log.Fatalf("Failed to get contract code: %v", err)
    }
    log.Printf("Contract deployed to: %s", address.Hex())
    log.Printf("Contract code size: %d bytes", len(code))
    if len(code) == 0 {
        log.Fatal("No contract code at deployed address")
    }

    // Verify contract functionality
    log.Println("Verifying contract functionality...")
    if instance == nil {
        log.Fatal("Contract instance is nil")
    }

    log.Println("Deployment successful!")
}
