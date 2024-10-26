package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"jagajkn/internal/config"
	"jagajkn/internal/models"
	"jagajkn/internal/repository"
	"jagajkn/internal/service"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.UserSignupInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userRepo := repository.NewUserRepository(db)
		userService := service.NewUserService(userRepo, "")

		user, err := userService.Register(&input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user.ToJSON())
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.UserLoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userRepo := repository.NewUserRepository(db)
		userService := service.NewUserService(userRepo, c.MustGet("config").(*config.Config).JWTSecret)

		token, err := userService.Login(&input)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}