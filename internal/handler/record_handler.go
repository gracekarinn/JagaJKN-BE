package handler

import (
	"log"
	"net/http"

	bService "jagajkn/internal/blockchain/service"
	"jagajkn/internal/models"
	"jagajkn/internal/repository"
	"jagajkn/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecordHandler struct {
    recordService *service.RecordService
}

func NewRecordHandler(db *gorm.DB, blockchainSvc *bService.BlockchainService) *RecordHandler {
    recordRepo := repository.NewRecordRepository(db)
    recordService := service.NewRecordService(recordRepo, blockchainSvc)
    return &RecordHandler{
        recordService: recordService,
    }
}

func (h *RecordHandler) CreateRecord() gin.HandlerFunc {
    return func(c *gin.Context) {
        userNIK := c.MustGet("userNIK").(string)
        log.Printf("Creating record for user: %s", userNIK)

        var input models.RecordInput
        if err := c.ShouldBindJSON(&input); err != nil {
            log.Printf("Invalid input: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
            return
        }
        log.Printf("Input data: %+v", input)

        result, err := h.recordService.CreateRecord(c.Request.Context(), &input, userNIK)
        if err != nil {
            log.Printf("Error creating record: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
            return
        }

        log.Printf("Record created successfully: %+v", result)
        c.JSON(http.StatusCreated, gin.H{
            "message": "Record created successfully",
            "record":  result,
        })
    }
}

func (h *RecordHandler) GetUserRecords() gin.HandlerFunc {
    return func(c *gin.Context) {
        userNIK := c.MustGet("userNIK").(string)
        log.Printf("Fetching records for user: %s", userNIK)

        records, err := h.recordService.GetUserRecords(c.Request.Context(), userNIK)
        if err != nil {
            log.Printf("Error fetching user records: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        if records == nil {
            records = []*models.RecordKesehatan{}
        }

        log.Printf("Fetched %d records for user: %s", len(records), userNIK)
        c.JSON(http.StatusOK, gin.H{
            "status": "success",
            "data": gin.H{
                "records": records,
                "count":   len(records),
            },
        })
    }
}

func (h *RecordHandler) GetRecord() gin.HandlerFunc {
    return func(c *gin.Context) {
        userNIK := c.MustGet("userNIK").(string)
        noSEP := c.Param("noSEP")
        log.Printf("Fetching record %s for user: %s", noSEP, userNIK)

        record, err := h.recordService.GetRecord(c.Request.Context(), noSEP)
        if err != nil {
            log.Printf("Error fetching record: %v", err)
            c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
            return
        }

        if record.UserNIK != userNIK {
            log.Printf("Access denied for user: %s to record: %s", userNIK, noSEP)
            c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
            return
        }

        log.Printf("Fetched record: %+v", record)
        c.JSON(http.StatusOK, record)
    }
}