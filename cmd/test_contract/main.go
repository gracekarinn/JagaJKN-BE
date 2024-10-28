package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
        log.Fatalf("Failed to connect to blockchain: %v", err)
    }

    contractAddr := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
    contract, err := contracts.NewContracts(contractAddr, client)
    if err != nil {
        log.Fatalf("Failed to instantiate contract: %v", err)
    }

    log.Printf("Testing contract at address: %s", contractAddr.Hex())

    testNIK := "test123"
    exists, err := contract.IsUserRegistered(&bind.CallOpts{
        Context: context.Background(),
    }, testNIK)
    if err != nil {
        log.Printf("Error calling isUserRegistered: %v", err)
    } else {
        log.Printf("User %s exists: %v", testNIK, exists)
    }

    // INI TEST KARENA ERROR MULU
    var testHash [32]byte
    copy(testHash[:], []byte("test"))
    verified, err := contract.VerifyUser(&bind.CallOpts{
        Context: context.Background(),
    }, testNIK, testHash)
    if err != nil {
        log.Printf("Error calling verifyUser: %v", err)
    } else {
        log.Printf("User %s verified: %v", testNIK, verified)
    }

    log.Println("Contract test complete")
}