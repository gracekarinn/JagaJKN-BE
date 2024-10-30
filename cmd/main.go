package main

import (
	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"
	"jagajkn/internal/migrations"
	"jagajkn/internal/router"
	"log"
	"os"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    log.Println("Connecting to database...")
    db, err := cfg.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Successfully connected to database")

    log.Println("Starting database migrations...")
    if err := migrations.RunMigrations(db); err != nil {
        log.Fatalf("Failed to run migrations: %v", err)
    }
    log.Println("Database migrations completed successfully")

    log.Println("Initializing blockchain service...")
    blockchainSvc, err := bService.NewBlockchainService(cfg.GetBlockchainConfig())
    if err != nil {
        log.Fatalf("Failed to initialize blockchain service: %v", err)
    }
    log.Println("Blockchain service initialized successfully")

    r := router.SetupRouter(db, cfg, blockchainSvc)

    port := os.Getenv("PORT")
    if port == "" {
        port = cfg.ServerPort 
    }

    log.Printf("Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}