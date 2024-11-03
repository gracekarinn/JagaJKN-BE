package handler

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
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
    data := fmt.Sprintf("%s:%s:%s", 
        user.NIK,
        user.NamaLengkap,
        user.NoTelp,
    )
    if user.Email != nil {
        data += ":" + *user.Email
    }
    
    log.Printf("Calculating hash for data: %s", data)
    
    hash := crypto.Keccak256([]byte(data))
    hashStr := hex.EncodeToString(hash)
    
    log.Printf("Generated hash: %s", hashStr)
    return hashStr
}

func (h *AuthHandler) GetProfile() gin.HandlerFunc {
    return func(c *gin.Context) {

        claims, _ := c.Get("claims")
        userClaims := claims.(jwt.MapClaims)
        nik := userClaims["nik"].(string)

        var user models.User
        if err := h.db.Where("nik = ?", nik).First(&user).Error; err != nil {
            log.Printf("User not found: %v", err)
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"user": user.ToJSON()})
    }
}

func (h *AuthHandler) CheckUserRegistration() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("Starting registration check...")
        
        nik := c.Query("nik")
        if nik == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "NIK is required"})
            return
        }

        log.Printf("Checking NIK: %s", nik)

        var user models.User
        if err := h.db.Where("nik = ?", nik).First(&user).Error; err != nil {
            log.Printf("User not found in database: %v", err)
            c.JSON(http.StatusNotFound, gin.H{"error": "User not found in database"})
            return
        }

        log.Printf("User found in database: %s", user.NamaLengkap)

        userHash := calculateUserHash(&user)
        log.Printf("Calculated user hash: %s", userHash)

        var hashBytes [32]byte
        copy(hashBytes[:], []byte(userHash))

        opts := &bind.CallOpts{
            Context: c.Request.Context(),
        }

        contract := h.blockchainSvc.GetContract()
        if contract == nil {
            log.Println("Contract not initialized")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Contract not initialized"})
            return
        }

        log.Println("Checking blockchain registration...")
        isRegistered, err := contract.IsUserRegistered(opts, nik)
        if err != nil {
            log.Printf("Blockchain check failed: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to check blockchain registration",
                "details": err.Error(),
            })
            return
        }

        log.Printf("Blockchain registration status: %v", isRegistered)

        var verificationStatus *bool
        if isRegistered {
            verified, err := contract.VerifyUser(opts, nik, hashBytes)
            if err != nil {
                log.Printf("Hash verification failed: %v", err)
            } else {
                verificationStatus = &verified
                log.Printf("Hash verification status: %v", verified)
            }
        }

        c.JSON(http.StatusOK, gin.H{
            "database_status": "found",
            "blockchain_status": isRegistered,
            "hash_verified": verificationStatus,
            "calculated_hash": userHash,
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

func (h *AuthHandler) FaskesLogin() gin.HandlerFunc {
    return func(c *gin.Context) {
        var input models.FaskesLoginInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input",
                "details": err.Error(),
            })
            return
        }

        var faskes models.Faskes
        if err := h.db.Where("kode_faskes = ?", input.KodeFaskes).First(&faskes).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if err := faskes.CheckPassword(input.Password); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        cfg := c.MustGet("config").(*config.Config)
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "kode_faskes": faskes.KodeFaskes,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        })

        tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "token": tokenString,
            "faskes": faskes.ToJSON(),
        })
    }
}

func (h *AuthHandler) AdminLogin() gin.HandlerFunc {
    return func(c *gin.Context) {
        var input models.AdminLoginInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input",
                "details": err.Error(),
            })
            return
        }

        var admin models.Admin
        if err := h.db.Where("email = ?", input.Email).First(&admin).Error; err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        if err := admin.CheckPassword(input.Password); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        cfg := c.MustGet("config").(*config.Config)
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
            "admin_id": admin.ID,
            "email": admin.Email,
            "exp": time.Now().Add(time.Hour * 24).Unix(),
        })

        tokenString, err := token.SignedString([]byte(cfg.JWTSecret))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "token": tokenString,
            "admin": admin.ToJSON(),
        })
    }
}