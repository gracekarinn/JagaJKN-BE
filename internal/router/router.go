package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"
	"jagajkn/internal/handler"
	"jagajkn/internal/middleware"
)

func SetupRouter(db *gorm.DB, cfg *config.Config, blockchainSvc *bService.BlockchainService) *gin.Engine {
    r := gin.Default()


    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:          12 * 60 * 60,
    }))

    r.Use(func(c *gin.Context) {
        if c.Request.URL.Path != "/" && c.Request.URL.Path[len(c.Request.URL.Path)-1] == '/' {
            c.Request.URL.Path = c.Request.URL.Path[:len(c.Request.URL.Path)-1]
        }
        c.Next()
    })

    r.Use(func(c *gin.Context) {
        c.Set("config", cfg)
        c.Next()
    })

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    authHandler := handler.NewAuthHandler(db, blockchainSvc)

    r.POST("/api/v1/auth/register", authHandler.Register())
    r.POST("/api/v1/auth/login", authHandler.Login())
    r.GET("/api/v1/auth/check-registration", authHandler.CheckUserRegistration())
    r.GET("/api/v1/auth/contract-status", authHandler.VerifyContractStatus())

    api := r.Group("/api/v1")
    api.Use(middleware.AuthMiddleware(cfg.JWTSecret))
    {
        // Nanti
    }

    return r
}