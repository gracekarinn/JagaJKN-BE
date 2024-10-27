package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/config"
	"jagajkn/internal/models"
)

type AuthHandler struct {
    db            *gorm.DB
    blockchainSvc *bService.BlockchainService
}

func NewAuthHandler(db *gorm.DB, blockchainSvc *bService.BlockchainService) *AuthHandler {
    return &AuthHandler{
        db:            db,
        blockchainSvc: blockchainSvc,
    }
}

func (h *AuthHandler) Register() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Starting registration process...")

        var input models.UserSignupInput
        if err := c.ShouldBindJSON(&input); err != nil {
            log.Printf("Input validation failed: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        log.Printf("Input validation passed for NIK: %s", input.NIK)

        var existingUser models.User
        if err := h.db.Where("nik = ?", input.NIK).First(&existingUser).Error; err == nil {
            log.Printf("User already exists with NIK: %s", input.NIK)
            c.JSON(http.StatusBadRequest, gin.H{"error": "NIK already registered"})
            return
        }

        user := models.User{
            NIK:         input.NIK,
            NamaLengkap: input.NamaLengkap,
            NoTelp:      input.NoTelp,
            Email:       input.Email,
            Password:    input.Password,
        }

        tx := h.db.Begin()
        log.Println("Starting database transaction...")

        if err := tx.Create(&user).Error; err != nil {
            tx.Rollback()
            log.Printf("Database save failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
            return
        }

        log.Println("User saved to database successfully")

        userHash := calculateUserHash(&user)
        log.Printf("Calculated user hash: %s", userHash)

        log.Println("Attempting to save to blockchain...")
        if err := h.blockchainSvc.SaveUserRegistration(c.Request.Context(), user.NIK, userHash); err != nil {
            tx.Rollback()
            log.Printf("Blockchain save failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user", "details": err.Error()})
            return
        }

        tx.Commit()
        log.Println("Transaction committed successfully")

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
            log.Printf("Login input validation failed: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        log.Printf("Attempting login for NIK: %s", input.NIK)

        var user models.User
        if err := h.db.Where("nik = ?", input.NIK).First(&user).Error; err != nil {
            log.Printf("Database lookup failed for NIK %s: %v", input.NIK, err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        log.Printf("User found in database: %s", user.NamaLengkap)

        if err := user.CheckPassword(input.Password); err != nil {
            log.Printf("Password verification failed for NIK %s: %v", input.NIK, err)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        log.Printf("Password verified successfully for NIK: %s", input.NIK)

        userHash := calculateUserHash(&user)
        log.Printf("Calculated user hash: %s", userHash)

        log.Printf("Attempting blockchain verification for NIK: %s", input.NIK)
        verified, err := h.blockchainSvc.VerifyUserRegistration(c.Request.Context(), user.NIK, userHash)
        if err != nil {
            log.Printf("Blockchain verification error for NIK %s: %v", input.NIK, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify user authenticity", "details": err.Error()})
            return
        }

        if !verified {
            log.Printf("Blockchain verification failed for NIK %s", input.NIK)
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User verification failed"})
            return
        }

        log.Printf("Blockchain verification successful for NIK: %s", input.NIK)

        cfg := c.MustGet("config").(*config.Config)
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "nik": user.NIK,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        })

        tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
        if err != nil {
            log.Printf("Token generation failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
            return
        }

        log.Printf("Login successful for NIK: %s", input.NIK)
        c.JSON(http.StatusOK, gin.H{
            "token": tokenString,
            "user":  user.ToJSON(),
        })
    }
}


func calculateUserHash(user *models.User) string {
    data := fmt.Sprintf("%s-%s-%s", user.NIK, user.NamaLengkap, user.NoTelp)
    if user.Email != nil {
        data += *user.Email
    }
    
    log.Printf("Calculating hash for data: %s", data)
    
    hash := sha256.Sum256([]byte(data))
    hashStr := hex.EncodeToString(hash[:])
    
    log.Printf("Generated hash: %s", hashStr)
    return hashStr
}

func (h *AuthHandler) VerifyUserRegistration() gin.HandlerFunc {
    return func(c *gin.Context) {
        nik := c.Query("nik")
        if nik == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "NIK is required"})
            return
        }

        var user models.User
        if err := h.db.Where("nik = ?", nik).First(&user).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        userHash := calculateUserHash(&user)

        verified, err := h.blockchainSvc.VerifyUserRegistration(c.Request.Context(), nik, userHash)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify registration"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "nik": nik,
            "verified": verified,
            "registrationDetails": gin.H{
                "namaLengkap": user.NamaLengkap,
                "email": user.Email,
                "noTelp": user.NoTelp,
            },
        })
    }
}

func (h *AuthHandler) VerifyContractStatus() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Starting contract status verification...")

        if h.blockchainSvc == nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Blockchain service not initialized",
            })
            return
        }

        status, err := h.blockchainSvc.CheckContractStatus(c.Request.Context())
        if err != nil {
            log.Printf("Contract status check failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Contract verification failed",
                "details": err.Error(),
            })
            return
        }

        log.Printf("Contract status check successful: %+v", status)

        status["timestamp"] = time.Now().Unix()
        status["serviceConnected"] = h.blockchainSvc != nil

        c.JSON(http.StatusOK, gin.H{
            "status": status,
            "message": "Contract status verified successfully",
        })
    }
}