package handler

import (
	"jagajkn/internal/config"
	"jagajkn/internal/models"
	"jagajkn/internal/repository"
	"jagajkn/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req models.UserLoginInput
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
            return
        }

        var user models.User
        if err := db.Where("nik = ?", req.NIK).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        cfg, exists := c.Get("config")
        if !exists {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Configuration not found"})
            return
        }

        config := cfg.(*config.Config)

        if !repository.CheckPasswordHash(req.Password, user.Password) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        token, err := repository.GenerateJWT(user.NIK, config.JWTSecret)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": token})
    }
}

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