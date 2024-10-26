package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"jagajkn/internal/blockchain/contracts"
)

func main() {
    // Load .env
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Ngonek ke blockchain
    client, err := ethclient.Dial(os.Getenv("BLOCKCHAIN_PROVIDER"))
    if err != nil {
        log.Fatal(err)
    }

    // Load private key
    privateKey, err := crypto.HexToECDSA(os.Getenv("BLOCKCHAIN_PRIVATE_KEY"))
    if err != nil {
        log.Fatal(err)
    }

    // Buat transactor
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
    if err != nil {
        log.Fatal(err)
    }

    // Set gas price and limit
    auth.GasPrice = big.NewInt(1000000000) // 1 gwei
    auth.GasLimit = uint64(3000000)        // 3 million gas

    // Deploy contract 
    address, tx, instance, err := contracts.DeployContracts(auth, client)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Contract deployment transaction sent: %s\n", tx.Hash().Hex())

    // Wait for deployment
    fmt.Println("Waiting for contract deployment transaction to be mined...")
    _, err = bind.WaitMined(context.Background(), client, tx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Contract deployed to: %s\n", address.Hex())
    fmt.Printf("Contract instance created: %v\n", instance != nil)

    // Save contract address to .env
    f, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    _, err = f.WriteString(fmt.Sprintf("\nCONTRACT_ADDRESS=%s", address.Hex()))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Contract address saved to .env file")
}