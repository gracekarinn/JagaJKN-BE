package router

import (
	"log"
	"net/http"

	"jagajkn/internal/config"
	"jagajkn/internal/handler"
	"jagajkn/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * 60 * 60,
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

    r.POST("/api/v1/auth/register", handler.Register(db))
    r.POST("/api/v1/auth/login", handler.Login(db))


    api := r.Group("/api/v1")
    {
        records := api.Group("/records")
        records.Use(middleware.AuthMiddleware(cfg.JWTSecret))
        {
            records.POST("", handler.CreateRecord(db))      
            records.POST("/", handler.CreateRecord(db))      
            records.GET("", handler.GetUserRecords(db))     
            records.GET("/", handler.GetUserRecords(db))    
            records.GET("/:id", handler.GetRecord(db))
        }
    }

    r.NoRoute(func(c *gin.Context) {
        log.Printf("404 for path: %s", c.Request.URL.Path)
        c.JSON(http.StatusNotFound, gin.H{"error": "Route not found"})
    })

    return r
}