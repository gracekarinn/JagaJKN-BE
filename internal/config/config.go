package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
    DatabaseURL     string
    ServerPort      string
    JWTSecret       string
    AllowedOrigin   string
    BlockchainConfig *BlockchainConfig
}

type BlockchainConfig struct {
    ProviderURL     string
    PrivateKey      string
    ContractAddress string
}

func LoadConfig() (*Config, error) {
    config := &Config{
        DatabaseURL:    getEnvOrPanic("DATABASE_URL"),
        ServerPort:     getEnvOrDefault("SERVER_PORT", "8080"),
        JWTSecret:      getEnvOrPanic("JWT_SECRET"),
        AllowedOrigin:  getEnvOrDefault("ALLOWED_ORIGIN", "*"),
        BlockchainConfig: &BlockchainConfig{
            ProviderURL:     getEnvOrPanic("BLOCKCHAIN_PROVIDER"),
            PrivateKey:      getEnvOrPanic("BLOCKCHAIN_PRIVATE_KEY"),
            ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
        },
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

func (c *Config) GetBlockchainConfig() *BlockchainConfig {
    return c.BlockchainConfig
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
