package main

import (
	"log"

	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"
	"jagajkn/internal/router"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    db, err := cfg.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    blockchainSvc, err := bService.NewBlockchainService(cfg.GetBlockchainConfig())
    if err != nil {
        log.Fatalf("Failed to initialize blockchain service: %v", err)
    }

    r := router.SetupRouter(db, cfg, blockchainSvc)

    log.Printf("Server starting on :%s", cfg.ServerPort)
    if err := r.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}