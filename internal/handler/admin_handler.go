package handler

import (
	"encoding/csv"
	"fmt"
	"io"
	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/models"
	"net/http"
	"strings"
	"time"

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

type ImportedData struct {
    User           models.User
    MedicalRecord  models.RecordKesehatan
    HasMedicalData bool
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

func truncateString(s string, maxLen int) string {
    if len(s) > maxLen {
        return s[:maxLen]
    }
    return s
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
        
        headers, err := reader.Read()
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid CSV format",
                "details": err.Error(),
            })
            return
        }

        columnMap := make(map[string]int)
        for i, header := range headers {
            columnMap[header] = i
        }

        requiredColumns := []string{"NIK", "NamaLengkap", "NoTelp", "Password"}
        for _, col := range requiredColumns {
            if _, exists := columnMap[col]; !exists {
                c.JSON(http.StatusBadRequest, gin.H{
                    "error": "Missing required column",
                    "details": fmt.Sprintf("Column %s is required", col),
                })
                return
            }
        }

        var successCount, errorCount int
        var errors []string

        for {
            record, err := reader.Read()
            if err == io.EOF {
                break
            }
            if err != nil {
                errorCount++
                errors = append(errors, fmt.Sprintf("Error reading row: %v", err))
                continue
            }

            tx := h.db.Begin()
            if tx.Error != nil {
                errorCount++
                errors = append(errors, fmt.Sprintf("Failed to begin transaction: %v", tx.Error))
                continue
            }

            user := models.User{
                NIK:         truncateString(record[columnMap["NIK"]], 16),
                NamaLengkap: truncateString(record[columnMap["NamaLengkap"]], 255),
                NoTelp:     truncateString(record[columnMap["NoTelp"]], 20),
                Password:   record[columnMap["Password"]],
            }

            if err := tx.Create(&user).Error; err != nil {
                if rbErr := tx.Rollback().Error; rbErr != nil {
                    errors = append(errors, fmt.Sprintf("Rollback failed: %v", rbErr))
                }
                errorCount++
                errors = append(errors, fmt.Sprintf("Failed to create user %s: %v", user.NIK, err))
                continue
            }

            if sepIdx, exists := columnMap["NoSEP"]; exists && sepIdx < len(record) && record[sepIdx] != "" {
                medRecord := models.RecordKesehatan{
                    NoSEP:          truncateString(record[columnMap["NoSEP"]], 20),
                    UserNIK:        user.NIK,
                    JenisRawat:     models.JenisRawat(record[columnMap["JenisRawat"]]),
                    DiagnosaAwal:   record[columnMap["DiagnosaAwal"]], 
                    DiagnosaPrimer: record[columnMap["DiagnosaPrimer"]], 
                    IcdX:           truncateString(record[columnMap["IcdX"]], 10),
                    Tindakan:       record[columnMap["Tindakan"]], 
                    TanggalMasuk:   time.Now(),
                }

                recordHash, err := h.blockchainSvc.CreateRecordHash(&medRecord)
                if err != nil {
                    if rbErr := tx.Rollback().Error; rbErr != nil {
                        errors = append(errors, fmt.Sprintf("Rollback failed: %v", rbErr))
                    }
                    errorCount++
                    errors = append(errors, fmt.Sprintf("Failed to create hash for medical record: %v", err))
                    continue
                }

                medRecord.HashCurrent = truncateString(strings.TrimPrefix(recordHash, "0x"), 64)

                if err := tx.Create(&medRecord).Error; err != nil {
                    if rbErr := tx.Rollback().Error; rbErr != nil {
                        errors = append(errors, fmt.Sprintf("Rollback failed: %v", rbErr))
                    }
                    errorCount++
                    errors = append(errors, fmt.Sprintf("Failed to create medical record for user %s: %v", user.NIK, err))
                    continue
                }

                if err := h.blockchainSvc.SaveMedicalRecord(c.Request.Context(), &medRecord); err != nil {
                    if rbErr := tx.Rollback().Error; rbErr != nil {
                        errors = append(errors, fmt.Sprintf("Rollback failed: %v", rbErr))
                    }
                    errorCount++
                    errors = append(errors, fmt.Sprintf("Failed to save medical record to blockchain: %v", err))
                    continue
                }
            }

            userHash, err := h.blockchainSvc.CreateUserHash(&user)
            if err != nil {
                if rbErr := tx.Rollback().Error; rbErr != nil {
                    errors = append(errors, fmt.Sprintf("Rollback failed: %v", rbErr))
                }
                errorCount++
                errors = append(errors, fmt.Sprintf("Failed to create hash for user %s: %v", user.NIK, err))
                continue
            }

            formattedUserHash := truncateString(strings.TrimPrefix(userHash, "0x"), 64)

            if err := h.blockchainSvc.SaveUserRegistration(c.Request.Context(), user.NIK, formattedUserHash); err != nil {
                if rbErr := tx.Rollback().Error; rbErr != nil {
                    errors = append(errors, fmt.Sprintf("Rollback failed: %v", rbErr))
                }
                errorCount++
                errors = append(errors, fmt.Sprintf("Failed to save user %s to blockchain: %v", user.NIK, err))
                continue
            }

            if err := tx.Commit().Error; err != nil {
                if rbErr := tx.Rollback().Error; rbErr != nil {
                    errors = append(errors, fmt.Sprintf("Rollback failed: %v", rbErr))
                }
                errorCount++
                errors = append(errors, fmt.Sprintf("Failed to commit transaction for user %s: %v", user.NIK, err))
                continue
            }

            successCount++
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Import process completed",
            "details": gin.H{
                "success_count": successCount,
                "error_count":   errorCount,
                "errors":        errors,
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