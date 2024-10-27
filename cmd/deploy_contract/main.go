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
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    client, err := ethclient.Dial(os.Getenv("BLOCKCHAIN_PROVIDER"))
    if err != nil {
        log.Fatal(err)
    }


    privateKey, err := crypto.HexToECDSA(os.Getenv("BLOCKCHAIN_PRIVATE_KEY"))
    if err != nil {
        log.Fatal(err)
    }


    chainID := big.NewInt(1337) 
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        log.Fatal(err)
    }


    auth.GasPrice = big.NewInt(1000000000)
    auth.GasLimit = uint64(3000000)


    address, tx, _, err := contracts.DeployContracts(auth, client)
    if err != nil {
        log.Fatalf("Failed to deploy contract: %v", err)
    }

    fmt.Printf("Contract deploying in transaction: %s\n", tx.Hash().Hex())


    _, err = bind.WaitMined(context.Background(), client, tx)
    if err != nil {
        log.Fatalf("Failed to wait for contract deployment: %v", err)
    }

    fmt.Printf("Contract deployed to: %s\n", address.Hex())
    
    f, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    if _, err := f.WriteString(fmt.Sprintf("\nCONTRACT_ADDRESS=%s", address.Hex())); err != nil {
        log.Fatal(err)
    }
}