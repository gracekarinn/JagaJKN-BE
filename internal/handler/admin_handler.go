package handler

import (
	"encoding/csv"
	"fmt"
	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminHandler struct {
    db            *gorm.DB
    blockchainSvc *bService.BlockchainService
}

func NewAdminHandler(db *gorm.DB, blockchainSvc *bService.BlockchainService) *AdminHandler {
    return &AdminHandler{
        db:            db,
        blockchainSvc: blockchainSvc,
    }
}

func (h *AdminHandler) GetAllUsers() gin.HandlerFunc {
    return func(c *gin.Context) {
        var users []models.User
        
        if err := h.db.Find(&users).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch users",
                "details": err.Error(),
            })
            return
        }

        var response []map[string]interface{}
        for _, user := range users {
            response = append(response, user.ToJSON())
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Users fetched successfully",
            "users":   response,
        })
    }
}

func (h *AdminHandler) ImportUsersFromCSV() gin.HandlerFunc {
    return func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "No file uploaded",
                "details": err.Error(),
            })
            return
        }

        openedFile, err := file.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Could not open file",
                "details": err.Error(),
            })
            return
        }
        defer openedFile.Close()

        reader := csv.NewReader(openedFile)
        _, err = reader.Read()
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid CSV format",
                "details": err.Error(),
            })
            return
        }

        var successCount, errorCount int
        var errors []string

        tx := h.db.Begin()

        for {
            record, err := reader.Read()
            if err != nil {
                break 
            }

            // CSV format: NIK,NamaLengkap,NoTelp,Email,Password
            if len(record) < 5 {
                errorCount++
                errors = append(errors, "Invalid row format")
                continue
            }

            user := models.User{
                NIK:         record[0],
                NamaLengkap: record[1],
                NoTelp:     record[2],
                Email:      &record[3],
                Password:   record[4],
                Role:       models.RoleUser, 
            }

            var existingUser models.User
            if err := h.db.Where("nik = ?", user.NIK).First(&existingUser).Error; err == nil {
                errorCount++
                errors = append(errors, fmt.Sprintf("User with NIK %s already exists", user.NIK))
                continue
            }

            if err := tx.Create(&user).Error; err != nil {
                errorCount++
                errors = append(errors, fmt.Sprintf("Failed to create user %s: %v", user.NIK, err))
                continue
            }

            userHash := calculateUserHash(&user)
            if err := h.blockchainSvc.SaveUserRegistration(c.Request.Context(), user.NIK, userHash); err != nil {
                errorCount++
                errors = append(errors, fmt.Sprintf("Failed to save user %s to blockchain: %v", user.NIK, err))
                continue
            }

            successCount++
        }

        if errorCount > 0 {
            tx.Rollback()
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Some users failed to import",
                "details": gin.H{
                    "success_count": successCount,
                    "error_count":   errorCount,
                    "errors":        errors,
                },
            })
            return
        }

        tx.Commit()
        c.JSON(http.StatusOK, gin.H{
            "message": "Users imported successfully",
            "details": gin.H{
                "total_imported": successCount,
            },
        })
    }
}

func (h *AdminHandler) CreateFaskes() gin.HandlerFunc {
    return func(c *gin.Context) {
        var input models.FaskesInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input",
                "details": err.Error(),
            })
            return
        }

        var existingFaskes models.Faskes
        if err := h.db.Where("kode_faskes = ?", input.KodeFaskes).First(&existingFaskes).Error; err == nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Faskes with this code already exists",
            })
            return
        }

        faskes := models.Faskes{
            KodeFaskes: input.KodeFaskes,
            Nama:       input.Nama,
            Alamat:     input.Alamat,
            NoTelp:     input.NoTelp,
            Tingkat:    input.Tingkat,
            Email:      input.Email,
            Password:   input.Password,
        }

        if err := h.db.Create(&faskes).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to create Faskes",
                "details": err.Error(),
            })
            return
        }

        c.JSON(http.StatusCreated, gin.H{
            "message": "Faskes created successfully",
            "faskes":  faskes.ToJSON(),
        })
    }
}

func (h *AdminHandler) GetAllFaskes() gin.HandlerFunc {
    return func(c *gin.Context) {
        var faskesList []models.Faskes
        
        if err := h.db.Find(&faskesList).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch Faskes list",
                "details": err.Error(),
            })
            return
        }

        var response []map[string]interface{}
        for _, faskes := range faskesList {
            response = append(response, faskes.ToJSON())
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Faskes list fetched successfully",
            "faskes":  response,
        })
    }
}

func (h *AdminHandler) GetFaskes() gin.HandlerFunc {
    return func(c *gin.Context) {
        kodeFaskes := c.Param("kodeFaskes")
        
        var faskes models.Faskes
        if err := h.db.Where("kode_faskes = ?", kodeFaskes).First(&faskes).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Faskes not found",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Faskes fetched successfully",
            "faskes":  faskes.ToJSON(),
        })
    }
}


func (h *AdminHandler) UpdateFaskes() gin.HandlerFunc {
    return func(c *gin.Context) {
        kodeFaskes := c.Param("kodeFaskes")
        
        var input struct {
            Nama    string          `json:"nama"`
            Alamat  string          `json:"alamat"`
            NoTelp  string          `json:"noTelp"`
            Tingkat models.TingkatFaskes `json:"tingkat"`
            Email   string          `json:"email"`
        }

        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input",
                "details": err.Error(),
            })
            return
        }

        var faskes models.Faskes
        if err := h.db.Where("kode_faskes = ?", kodeFaskes).First(&faskes).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Faskes not found",
            })
            return
        }


        updates := map[string]interface{}{
            "nama": input.Nama,
            "alamat": input.Alamat,
            "no_telp": input.NoTelp,
            "tingkat": input.Tingkat,
            "email": input.Email,
        }

        if err := h.db.Model(&faskes).Updates(updates).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to update Faskes",
                "details": err.Error(),
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Faskes updated successfully",
            "faskes":  faskes.ToJSON(),
        })
    }
}

func (h *AdminHandler) DeleteFaskes() gin.HandlerFunc {
    return func(c *gin.Context) {
        kodeFaskes := c.Param("kodeFaskes")
        
        if err := h.db.Where("kode_faskes = ?", kodeFaskes).Delete(&models.Faskes{}).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to delete Faskes",
                "details": err.Error(),
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Faskes deleted successfully",
        })
    }
}