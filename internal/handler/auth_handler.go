package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	"jagajkn/internal/config"
	"jagajkn/internal/models"
)

type AuthHandler struct {
    db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
    return &AuthHandler{
        db: db,
    }
}

func (h *AuthHandler) Register() gin.HandlerFunc {
    return func(c *gin.Context) {
        var input models.UserSignupInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var existingUser models.User
        if err := h.db.Where("nik = ?", input.NIK).First(&existingUser).Error; err == nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "NIK already registered"})
            return
        }

        user := models.User{
            NIK:         input.NIK,
            NamaLengkap: input.NamaLengkap,
            NoTelp:      input.NoTelp,
            Email:       input.Email,
            Password:    input.Password, // Will be hashed by BeforeCreate hook
        }

        if err := h.db.Create(&user).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
            return
        }

        c.JSON(http.StatusCreated, gin.H{
            "message": "User registered successfully",
            "user":    user.ToJSON(),
        })
    }
}

func (h *AuthHandler) Login() gin.HandlerFunc {
    return func(c *gin.Context) {
        var input models.UserLoginInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        var user models.User
        if err := h.db.Where("nik = ?", input.NIK).First(&user).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if err := user.CheckPassword(input.Password); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        cfg := c.MustGet("config").(*config.Config)

        claims := jwt.MapClaims{
            "nik": user.NIK,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
        tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "token": tokenString,
            "user":  user.ToJSON(),
        })
    }
}