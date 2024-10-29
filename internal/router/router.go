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
	"jagajkn/internal/models"
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
        auth.GET("/check-registration", authHandler.CheckUserRegistration())
        auth.GET("/contract-status", authHandler.VerifyContractStatus())
    }

    api := r.Group("/api/v1")
    api.Use(middleware.AuthMiddleware(cfg.JWTSecret))
    {
        userRoutes := api.Group("/user")
        userRoutes.Use(middleware.RequireRole(models.RoleUser, models.RoleAdmin, models.RoleFaskes))
        {
            userRoutes.GET("/records", recordHandler.GetUserRecords())
            userRoutes.GET("/records/:noSEP", recordHandler.GetRecord())
        }

        adminRoutes := api.Group("/admin")
        adminRoutes.Use(middleware.RequireRole(models.RoleAdmin))
        {
            adminRoutes.POST("/users/import", adminHandler.ImportUsersFromCSV())
            
            adminRoutes.POST("/faskes", adminHandler.CreateFaskes())
            adminRoutes.GET("/faskes", adminHandler.GetAllFaskes())
            adminRoutes.GET("/faskes/:id", adminHandler.GetFaskes())
            
            adminRoutes.GET("/users", adminHandler.GetAllUsers())
        }

        // faskesRoutes := api.Group("/faskes")
        // faskesRoutes.Use(middleware.RequireRole(models.RoleFaskes))
        // {
        //     faskesRoutes.POST("/records", recordHandler.CreateRecord())
        //     faskesRoutes.PUT("/records/:noSEP", recordHandler.UpdateRecord())
        //     faskesRoutes.GET("/records", recordHandler.GetFaskesRecords())
        // }
    }





    return r
}