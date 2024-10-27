package main

import (
	"fmt"
	"log"

	"jagajkn/internal/config"

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

    fmt.Println("Configuration loaded successfully:")
    fmt.Printf("Blockchain Provider: %s\n", cfg.BlockchainConfig.ProviderURL)
    fmt.Printf("Contract Address: %s\n", cfg.BlockchainConfig.ContractAddress)
    fmt.Println("JWT Secret: [HIDDEN]")
    fmt.Println("Private Key: [HIDDEN]")
}