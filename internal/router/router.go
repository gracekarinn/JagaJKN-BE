package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"jagajkn/internal/config"
	"jagajkn/internal/handler"
	"jagajkn/internal/middleware"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) http.Handler {

	if gin.Mode() == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()


	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.AllowedOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:          12 * 60 * 60, // 12 hours
	}))


	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})


	v1 := router.Group("/api/v1")
	{

		auth := v1.Group("/auth")
		{
			auth.POST("/register", handler.Register(db))
			auth.POST("/login", handler.Login(db))
		}


		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			records := protected.Group("/records")
			{
				records.POST("/", handler.CreateRecord(db))
				records.GET("/", handler.GetUserRecords(db))
				records.GET("/:id", handler.GetRecord(db))
			}
		}
	}

	return router
}