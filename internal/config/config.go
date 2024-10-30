package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
    DatabaseURL      string
    ServerPort       string
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
    if os.Getenv("RAILWAY_ENVIRONMENT") == "" {
        if err := godotenv.Load(); err != nil {
            log.Println("No .env file found, assuming environment variables are set")
        }
    }

    config := &Config{
        DatabaseURL:    os.Getenv("DATABASE_URL"),
        ServerPort:     getEnvOrDefault("PORT", "8080"),
        JWTSecret:     os.Getenv("JWT_SECRET"),
        AllowedOrigin: getEnvOrDefault("ALLOWED_ORIGIN", "*"),
        BlockchainConfig: &BlockchainConfig{
            ProviderURL:     os.Getenv("BLOCKCHAIN_PROVIDER"),
            PrivateKey:      os.Getenv("BLOCKCHAIN_PRIVATE_KEY"),
            ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
        },
    }

    if err := config.validate(); err != nil {
        return nil, err
    }

    return config, nil
}


func (c *Config) validate() error {
    missingVars := []string{}

    if c.DatabaseURL == "" {
        missingVars = append(missingVars, "DATABASE_URL")
    }
    if c.JWTSecret == "" {
        missingVars = append(missingVars, "JWT_SECRET")
    }
    if c.BlockchainConfig.ProviderURL == "" {
        missingVars = append(missingVars, "BLOCKCHAIN_PROVIDER")
    }
    if c.BlockchainConfig.PrivateKey == "" {
        missingVars = append(missingVars, "BLOCKCHAIN_PRIVATE_KEY")
    }

    if len(missingVars) > 0 {
        return fmt.Errorf("missing required environment variables: %v", missingVars)
    }

    return nil
}

func (c *Config) ConnectDB() (*gorm.DB, error) {
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold:             time.Second,
            LogLevel:                  logger.Info,
            IgnoreRecordNotFoundError: true,
            Colorful:                  true,
        },
    )

    var db *gorm.DB
    var err error
    maxRetries := 5
    retryDelay := time.Second * 5

    for i := 0; i < maxRetries; i++ {
        db, err = gorm.Open(postgres.Open(c.DatabaseURL), &gorm.Config{
            Logger: newLogger,
        })

        if err == nil {
            break
        }

        log.Printf("Failed to connect to database (attempt %d/%d): %v\n", i+1, maxRetries, err)
        if i < maxRetries-1 {
            time.Sleep(retryDelay)
        }
    }

    if err != nil {
        return nil, fmt.Errorf("failed to connect to database after %d attempts: %v", maxRetries, err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    if err := sqlDB.Ping(); err != nil {
        return nil, err
    }

    log.Println("Successfully connected to database")
    return db, nil
}

func (c *Config) GetBlockchainConfig() *BlockchainConfig {
    return c.BlockchainConfig
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}