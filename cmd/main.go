package main

import (
	"log"

	"jagajkn/internal/config"
	"jagajkn/internal/migrations"
	"jagajkn/internal/router"
	"jagajkn/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := cfg.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}


	if err := migrations.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	r := router.SetupRouter(db, cfg)


	server := server.NewServer(r, cfg.ServerPort)
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := server.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}