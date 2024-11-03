package router

import (
	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"
	"jagajkn/internal/handler"
	"jagajkn/internal/middleware"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config, blockchainSvc *bService.BlockchainService) *gin.Engine {
    log.Printf("Setting up router...")

    r := gin.New()

    r.Use(gin.Recovery())
    r.Use(gin.Logger())

    corsConfig := cors.Config{
        AllowOrigins:     []string{"http://localhost:3000", "jagajkn.vercel.app"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * 3600, 
    }
    r.Use(cors.New(corsConfig))

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

    log.Println("Initializing handlers...")
    authHandler := handler.NewAuthHandler(db, blockchainSvc)
    recordHandler := handler.NewRecordHandler(db, blockchainSvc)
    adminHandler := handler.NewAdminHandler(db, blockchainSvc)
    faskesHandler := handler.NewFaskesHandler(db, blockchainSvc)

    r.GET("/", func(c *gin.Context) {
        log.Println("Handling root endpoint request")
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "message": "JagaJKN API is running",
            "version": "1.0.0",
        })
    })

    r.GET("/health", func(c *gin.Context) {
        log.Println("Handling health check request")
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "database": "connected",
            "blockchain": "initialized",
        })
    })

    log.Println("Setting up API routes...")
    
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
            userRoutes.GET("/profile", authHandler.GetProfile())
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
            faskesRoutes.GET("/profile", faskesHandler.GetProfile())
            faskesRoutes.PUT("/profile", faskesHandler.UpdateProfile())
            faskesRoutes.POST("/records", recordHandler.CreateRecord())
            faskesRoutes.GET("/records", recordHandler.GetFaskesRecords())
            faskesRoutes.GET("/records/:noSEP", recordHandler.GetRecord())
            faskesRoutes.POST("/transfer", faskesHandler.InitiateTransfer())
            faskesRoutes.GET("/transfers/pending", faskesHandler.GetPendingTransfers())
            faskesRoutes.POST("/transfers/:transferId/accept", faskesHandler.AcceptTransfer())
            faskesRoutes.GET("/transfers/history", faskesHandler.GetTransferHistory())
        }
    }

    r.NoRoute(func(c *gin.Context) {
        log.Printf("404 Not Found: %s %s", c.Request.Method, c.Request.URL.Path)
        c.JSON(http.StatusNotFound, gin.H{
            "status": "error",
            "message": "Route not found",
            "path": c.Request.URL.Path,
            "method": c.Request.Method,
        })
    })

    log.Println("Router setup completed successfully")
    return r
}