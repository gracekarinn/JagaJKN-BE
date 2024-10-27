package main

import (
	"fmt"
	"log"

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

    r := router.SetupRouter(db, cfg)

    serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
    log.Printf("Server starting on %s", serverAddr)
    if err := r.Run(serverAddr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}