package repository

import (
	"context"
	"jagajkn/internal/models"
)

type RecordRepository interface {
    Create(ctx context.Context, record *models.RecordKesehatan) error
    GetByNoSEP(ctx context.Context, noSEP string) (*models.RecordKesehatan, error)
    GetByUserNIK(ctx context.Context, userNIK string) ([]*models.RecordKesehatan, error)
}