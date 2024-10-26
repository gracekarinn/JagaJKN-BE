package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DatabaseURL    string
	ServerPort     string
	JWTSecret      string
	AllowedOrigin  string
}


func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	config := &Config{
		DatabaseURL:    getEnvOrPanic("DATABASE_URL"),
		ServerPort:     getEnvOrDefault("SERVER_PORT", "8080"),
		JWTSecret:      getEnvOrPanic("JWT_SECRET"),
		AllowedOrigin:  getEnvOrDefault("ALLOWED_ORIGIN", "*"),
	}

	return config, nil
}


func (c *Config) ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}


	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}


	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}


	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}


func getEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Environment variable %s is required", key))
	}
	return value
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}