package main

import (
	"log"

	"jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"

	"github.com/joho/godotenv"
)

func main() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    cfg := config.NewBlockchainConfig()
    
    // Buat blockchain service
    blockchainService, err := service.NewBlockchainService(cfg)
    if err != nil {
        log.Fatalf("Failed to create blockchain service: %v", err)
    }
    
    // Test koneksi
    if err := blockchainService.TestConnection(); err != nil {
        log.Fatalf("Failed to connect to blockchain: %v", err)
    }
    
    log.Println("Successfully connected to blockchain!")
}