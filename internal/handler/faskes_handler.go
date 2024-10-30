package handler

import (
	"fmt"
	"net/http"
	"time"

	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FaskesHandler struct {
    db            *gorm.DB
    blockchainSvc *bService.BlockchainService
}

func NewFaskesHandler(db *gorm.DB, blockchainSvc *bService.BlockchainService) *FaskesHandler {
    return &FaskesHandler{
        db:            db,
        blockchainSvc: blockchainSvc,
    }
}

func (h *FaskesHandler) GetProfile() gin.HandlerFunc {
    return func(c *gin.Context) {
        faskesKode, _ := c.Get("faskes_kode")
        
        var faskes models.Faskes
        if err := h.db.Where("kode_faskes = ?", faskesKode).First(&faskes).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Faskes not found",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Faskes profile fetched successfully",
            "faskes":  faskes.ToJSON(),
        })
    }
}

func (h *FaskesHandler) UpdateProfile() gin.HandlerFunc {
    return func(c *gin.Context) {
        faskesKode, _ := c.Get("faskes_kode")
        
        var input struct {
            Nama    string    `json:"nama"`
            Alamat  string    `json:"alamat"`
            NoTelp  string    `json:"noTelp"`
            Email   string    `json:"email"`
        }

        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input",
                "details": err.Error(),
            })
            return
        }

        updates := map[string]interface{}{
            "nama": input.Nama,
            "alamat": input.Alamat,
            "no_telp": input.NoTelp,
            "email": input.Email,
            "updated_at": time.Now(),
        }

        result := h.db.Model(&models.Faskes{}).
            Where("kode_faskes = ?", faskesKode).
            Updates(updates)

        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to update profile",
                "details": result.Error.Error(),
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Profile updated successfully",
        })
    }
}

func (h *FaskesHandler) InitiateTransfer() gin.HandlerFunc {
    return func(c *gin.Context) {
        var input struct {
            NoSEP           string `json:"noSEP" binding:"required"`
            DestinationCode string `json:"destinationCode" binding:"required"`
            Reason          string `json:"reason" binding:"required"`
        }

        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input",
                "details": err.Error(),
            })
            return
        }

        sourceFaskes, _ := c.Get("faskes_kode")

        tx := h.db.Begin()

        var record models.RecordKesehatan
        if err := tx.Preload("User").Where("no_sep = ?", input.NoSEP).First(&record).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Record not found",
            })
            return
        }

        transfer := models.RekamMedisTransfer{
            ID:               fmt.Sprintf("TRF-%s-%s", record.NoSEP, time.Now().Format("20060102150405")),
            NoSEP:           record.NoSEP,
            RecordKesehatan: &record,
            SourceFaskes:    sourceFaskes.(string),
            DestinationFaskes: input.DestinationCode,
            TransferReason:   input.Reason,
            TransferTime:     time.Now(),
            Status:          "PENDING",
        }

        if err := tx.Create(&transfer).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to create transfer record",
                "details": err.Error(),
            })
            return
        }

        if err := tx.Model(&record).Updates(map[string]interface{}{
            "status_pulang": models.Rujuk,
            "tanggal_keluar": time.Now(),
        }).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to update record status",
            })
            return
        }

        if err := tx.Commit().Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to commit transaction",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Transfer initiated successfully",
            "transfer": transfer.ToJSON(),
        })
    }
}

func (h *FaskesHandler) GetPendingTransfers() gin.HandlerFunc {
    return func(c *gin.Context) {
        faskesKode, _ := c.Get("faskes_kode")

        var transfers []models.RekamMedisTransfer
        result := h.db.
            Preload("RecordKesehatan").
            Preload("RecordKesehatan.User").
            Where("destination_faskes = ? AND status = ?", faskesKode, "PENDING").
            Find(&transfers)

        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch transfers",
                "details": result.Error.Error(),
            })
            return
        }

        var response []map[string]interface{}
        for _, transfer := range transfers {
            response = append(response, transfer.ToJSON())
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Pending transfers fetched successfully",
            "transfers": response,
        })
    }
}

func (h *FaskesHandler) AcceptTransfer() gin.HandlerFunc {
    return func(c *gin.Context) {
        transferID := c.Param("transferId")
        faskesKode, _ := c.Get("faskes_kode")

        tx := h.db.Begin()

        var transfer models.RekamMedisTransfer
        if err := tx.Preload("RecordKesehatan").
            Where("id = ? AND destination_faskes = ? AND status = ?", 
                transferID, faskesKode, "PENDING").
            First(&transfer).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Transfer not found",
            })
            return
        }

        if err := tx.Model(&transfer).Update("status", "ACCEPTED").Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to update transfer status",
            })
            return
        }

        if err := tx.Commit().Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to commit transaction",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Transfer accepted successfully",
            "transfer": transfer.ToJSON(),
        })
    }
}

func (h *FaskesHandler) GetTransferHistory() gin.HandlerFunc {
    return func(c *gin.Context) {
        faskesKode, _ := c.Get("faskes_kode")

        var transfers []models.RekamMedisTransfer
        result := h.db.
            Preload("RecordKesehatan").
            Preload("RecordKesehatan.User").
            Where("source_faskes = ? OR destination_faskes = ?", faskesKode, faskesKode).
            Order("created_at DESC").
            Find(&transfers)

        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch transfer history",
                "details": result.Error.Error(),
            })
            return
        }

        var response []map[string]interface{}
        for _, transfer := range transfers {
            response = append(response, transfer.ToJSON())
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Transfer history fetched successfully",
            "transfers": response,
        })
    }
}