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

    authHandler := handler.NewAuthHandler(db, blockchainSvc)
    recordHandler := handler.NewRecordHandler(db, blockchainSvc)
    adminHandler := handler.NewAdminHandler(db, blockchainSvc)
    // faskesHandler := handler.NewFaskesHandler(db, blockchainSvc)

    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })

    auth := r.Group("/api/v1/auth")
    {
        auth.POST("/register", authHandler.Register())
        auth.POST("/login", authHandler.Login())
        
        auth.POST("/admin/login", authHandler.AdminLogin())
        
        auth.POST("/faskes/login", authHandler.FaskesLogin())
        
        auth.GET("/check-registration", authHandler.CheckUserRegistration())
        auth.GET("/contract-status", authHandler.VerifyContractStatus())
    }

    api := r.Group("/api/v1")
    {
        userRoutes := api.Group("/user")
        userRoutes.Use(middleware.UserAuthMiddleware(cfg.JWTSecret))
        {
            userRoutes.GET("/records", recordHandler.GetUserRecords())
            userRoutes.GET("/records/:noSEP", recordHandler.GetRecord())
        }

        adminRoutes := api.Group("/admin")
        adminRoutes.Use(middleware.AdminAuthMiddleware(cfg.JWTSecret))
        {
            adminRoutes.POST("/users/import", adminHandler.ImportUsersFromCSV())
            adminRoutes.GET("/users", adminHandler.GetAllUsers())
            
            adminRoutes.POST("/faskes", adminHandler.CreateFaskes())
            adminRoutes.GET("/faskes", adminHandler.GetAllFaskes())
            adminRoutes.GET("/faskes/:kodeFaskes", adminHandler.GetFaskes())
            adminRoutes.PUT("/faskes/:kodeFaskes", adminHandler.UpdateFaskes())
            adminRoutes.DELETE("/faskes/:kodeFaskes", adminHandler.DeleteFaskes())
        }
        faskesRoutes := api.Group("/faskes")
        faskesRoutes.Use(middleware.FaskesAuthMiddleware(cfg.JWTSecret))
        {
            faskesRoutes.POST("/records", recordHandler.CreateRecord())
            // faskesRoutes.GET("/records", recordHandler.GetFaskesRecords())
        }
    }

    return r
}