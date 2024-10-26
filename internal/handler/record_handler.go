package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"jagajkn/internal/models"
	"jagajkn/internal/repository"
	"jagajkn/internal/service"
)

func CreateRecord(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(string)

		var input models.RecordInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recordRepo := repository.NewRecordRepository(db)
		recordService := service.NewRecordService(recordRepo)

		record, err := recordService.CreateRecord(userID, &input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, record)
	}
}

func GetUserRecords(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(string)

		recordRepo := repository.NewRecordRepository(db)
		recordService := service.NewRecordService(recordRepo)

		records, err := recordService.GetUserRecords(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, records)
	}
}

func GetRecord(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("userID").(string)
		recordID := c.Param("id")

		recordRepo := repository.NewRecordRepository(db)
		recordService := service.NewRecordService(recordRepo)

		record, err := recordService.GetRecord(recordID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
			return
		}

		if record.UserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			return
		}

		c.JSON(http.StatusOK, record)
	}
}