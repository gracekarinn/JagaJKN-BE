package main

import (
	"jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    blockchainService, err := service.NewBlockchainService(cfg.GetBlockchainConfig())
    if err != nil {
        log.Fatalf("Failed to create blockchain service: %v", err)
    }

    if err := blockchainService.TestConnection(); err != nil {
        log.Fatalf("Failed to connect to blockchain: %v", err)
    }

    log.Println("Successfully connected to blockchain!")
}