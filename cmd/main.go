package main

import (
	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"
	"jagajkn/internal/migrations"
	"jagajkn/internal/router"
	"log"
	"os"
)

var (
    infoLog  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
    errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
)

func main() {
    infoLog.Println("Loading configuration...")
    cfg, err := config.LoadConfig()
    if err != nil {
        errorLog.Fatalf("Failed to load config: %v", err)
    }

    infoLog.Println("Connecting to database...")
    db, err := cfg.ConnectDB()
    if err != nil {
        errorLog.Fatalf("Failed to connect to database: %v", err)
    }
    infoLog.Println("✅ Successfully connected to database")


    infoLog.Println("Starting database migrations...")
    if err := migrations.RunMigrations(db); err != nil {
        errorLog.Fatalf("Failed to run migrations: %v", err)
    }
    infoLog.Println("✅ Database migrations completed successfully")

    infoLog.Println("Initializing blockchain service...")
    blockchainSvc, err := bService.NewBlockchainService(cfg.GetBlockchainConfig())
    if err != nil {
        errorLog.Fatalf("Failed to initialize blockchain service: %v", err)
    }
    infoLog.Println("✅ Blockchain service initialized successfully")

    port := os.Getenv("PORT")
    if port == "" {
        port = cfg.ServerPort
    }

    infoLog.Println("Setting up router...")
    r := router.SetupRouter(db, cfg, blockchainSvc)
    infoLog.Println("✅ Router setup completed")

    infoLog.Printf("🚀 Server starting on port %s", port)
    if err := r.Run(":" + port); err != nil {
        errorLog.Fatalf("Failed to start server: %v", err)
    }
}