package repository

import (
	"context"
	"jagajkn/internal/models"

	"gorm.io/gorm"
)


type recordRepositoryImpl struct {
    db *gorm.DB
}


func NewRecordRepository(db *gorm.DB) RecordRepository {
    return &recordRepositoryImpl{
        db: db,
    }
}


func (r *recordRepositoryImpl) GetByNoSEP(ctx context.Context, noSEP string) (*models.RecordKesehatan, error) {
    var record models.RecordKesehatan
    err := r.db.WithContext(ctx).
        Preload("User").  
        Where("no_sep = ?", noSEP).
        First(&record).Error
    if err != nil {
        return nil, err
    }
    return &record, nil
}

func (r *recordRepositoryImpl) GetByUserNIK(ctx context.Context, userNIK string) ([]*models.RecordKesehatan, error) {
    var records []*models.RecordKesehatan
    err := r.db.WithContext(ctx).
        Preload("User").  
        Where("user_nik = ?", userNIK).
        Find(&records).Error
    if err != nil {
        return nil, err
    }
    return records, nil
}

func (r *recordRepositoryImpl) Create(ctx context.Context, record *models.RecordKesehatan) error {
    return r.db.WithContext(ctx).Create(record).Error
}