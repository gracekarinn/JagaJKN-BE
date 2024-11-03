package handler

import (
	"net/http"
	"strings"
	"time"

	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecordHandler struct {
    db            *gorm.DB
    blockchainSvc *bService.BlockchainService
}

func NewRecordHandler(db *gorm.DB, blockchainSvc *bService.BlockchainService) *RecordHandler {
    return &RecordHandler{
        db:            db,
        blockchainSvc: blockchainSvc,
    }
}

func (h *RecordHandler) CreateRecord() gin.HandlerFunc {
    return func(c *gin.Context) {
        faskesKode, exists := c.Get("faskes_kode")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Faskes not authenticated",
            })
            return
        }

        var input models.RecordInput
        if err := c.ShouldBindJSON(&input); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid input",
                "details": err.Error(),
            })
            return
        }

        record := models.RecordKesehatan{
            NoSEP:          input.NoSEP,
            UserNIK:        input.UserNIK,
            FaskesKode:     faskesKode.(string),
            JenisRawat:     input.JenisRawat,
            DiagnosaAwal:   input.DiagnosaAwal,
            DiagnosaPrimer: input.DiagnosaPrimer,
            IcdX:           input.IcdX,
            Tindakan:       input.Tindakan,
            TanggalMasuk:   time.Now(),
        }

        tx := h.db.Begin()

        if err := tx.Create(&record).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to create record",
                "details": err.Error(),
            })
            return
        }

        recordHash, err := h.blockchainSvc.CreateRecordHash(&record)
        if err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to create hash",
                "details": err.Error(),
            })
            return
        }

        hashStr := strings.TrimPrefix(recordHash, "0x")
        if len(hashStr) > 64 {
            hashStr = hashStr[:64]
        }
        record.HashCurrent = hashStr

        if err := tx.Save(&record).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to update record hash",
                "details": err.Error(),
            })
            return
        }

        if err := h.blockchainSvc.SaveMedicalRecord(c.Request.Context(), &record); err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to save to blockchain",
                "details": err.Error(),
            })
            return
        }

        if err := tx.Commit().Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to commit transaction",
                "details": err.Error(),
            })
            return
        }

        c.JSON(http.StatusCreated, gin.H{
            "message": "Record created successfully",
            "record": record,
        })
    }
}

func (h *RecordHandler) GetFaskesRecords() gin.HandlerFunc {
    return func(c *gin.Context) {
        faskesKode, _ := c.Get("faskes_kode")

        var records []models.RecordKesehatan
        result := h.db.
            Preload("User").  
            Where("faskes_kode = ?", faskesKode).
            Order("created_at DESC").
            Find(&records)

        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch records",
                "details": result.Error.Error(),
            })
            return
        }

        var response []map[string]interface{}
        for _, record := range records {
            verified, _ := h.blockchainSvc.VerifyMedicalRecord(c.Request.Context(), &record)
            recordData := record.ToBlockchainRecord()
            recordData["blockchainVerified"] = verified
            response = append(response, recordData)
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Records fetched successfully",
            "records": response,
            "count":   len(records),
        })
    }
}

func (h *RecordHandler) GetUserRecords() gin.HandlerFunc {
    return func(c *gin.Context) {
        userNIK, exists := c.Get("user_nik")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "User not authenticated",
            })
            return
        }

        var records []models.RecordKesehatan
        result := h.db.
            Where("user_nik = ?", userNIK).
            Order("created_at DESC").
            Find(&records)

        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to fetch records",
                "details": result.Error.Error(),
            })
            return
        }

        var response []map[string]interface{}
        for _, record := range records {
            verified, err := h.blockchainSvc.VerifyMedicalRecord(c.Request.Context(), &record)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{
                    "error": "Failed to verify record",
                    "details": err.Error(),
                })
                return
            }
            recordData := record.ToBlockchainRecord()
            recordData["blockchainVerified"] = verified
            response = append(response, recordData)
        }

        c.JSON(http.StatusOK, gin.H{
            "message": "Records fetched successfully",
            "records": response,
            "count":   len(records),
        })
    }
}

func (h *RecordHandler) GetRecord() gin.HandlerFunc {
    return func(c *gin.Context) {
        noSEP := c.Param("noSEP")

        var record models.RecordKesehatan
        if err := h.db.
            Preload("User").
            Where("no_sep = ?", noSEP).
            First(&record).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "Record not found",
            })
            return
        }

        if userNIK, exists := c.Get("user_nik"); exists {
            if record.UserNIK != userNIK.(string) {
                c.JSON(http.StatusForbidden, gin.H{
                    "error": "Access denied",
                })
                return
            }
        } else if faskesKode, exists := c.Get("faskes_kode"); exists {
            if record.FaskesKode != faskesKode.(string) {
                c.JSON(http.StatusForbidden, gin.H{
                    "error": "Access denied",
                })
                return
            }
        }

        verified, _ := h.blockchainSvc.VerifyMedicalRecord(c.Request.Context(), &record)
        recordData := record.ToBlockchainRecord()
        recordData["blockchainVerified"] = verified

        c.JSON(http.StatusOK, gin.H{
            "message": "Record fetched successfully",
            "record":  recordData,
        })
    }
}