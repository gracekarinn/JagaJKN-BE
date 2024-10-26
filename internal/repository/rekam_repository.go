package repository

import (
	"jagajkn/internal/models"

	"gorm.io/gorm"
)

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) *RecordRepository {
	return &RecordRepository{db: db}
}

func (r *RecordRepository) Create(record *models.RecordKesehatan) error {
	return r.db.Create(record).Error
}

func (r *RecordRepository) FindByID(id string) (*models.RecordKesehatan, error) {
	var record models.RecordKesehatan
	err := r.db.Preload("User").Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *RecordRepository) FindByUserID(userID string) ([]models.RecordKesehatan, error) {
	var records []models.RecordKesehatan
	err := r.db.Where("user_id = ?", userID).Find(&records).Error
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (r *RecordRepository) Update(record *models.RecordKesehatan) error {
	return r.db.Save(record).Error
}