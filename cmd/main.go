package main

import (
	"fmt"
	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"
	"jagajkn/internal/migrations"
	"jagajkn/internal/router"
	"log"
	"os"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("Starting application...")

    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }
    log.Println("Configuration loaded successfully")

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

    log.Println("Setting up router...")
    r := router.SetupRouter(db, cfg, blockchainSvc)
    log.Println("Router setup completed")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }

    routes := r.Routes()
    log.Println("Registered routes:")
    for _, route := range routes {
        log.Printf("%s %s", route.Method, route.Path)
    }

    addr := fmt.Sprintf("0.0.0.0:%s", port)
    log.Printf("Server starting on %s", addr)
    if err := r.Run(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}