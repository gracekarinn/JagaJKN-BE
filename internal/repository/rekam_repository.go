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
    if err := r.db.Create(record).Error; err != nil {
        return err
    }
    return nil
}

func (r *RecordRepository) FindByUserNIK(userNIK string) ([]models.RecordKesehatan, error) {
    var records []models.RecordKesehatan
    if err := r.db.Where("user_nik = ?", userNIK).Find(&records).Error; err != nil {
        return nil, err
    }
    return records, nil
}

func (r *RecordRepository) FindByNoSEP(noSEP string) (*models.RecordKesehatan, error) {
    var record models.RecordKesehatan
    if err := r.db.Where("no_sep = ?", noSEP).First(&record).Error; err != nil {
        return nil, err
    }
    return &record, nil
}