package repository

import (
	"jagajkn/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByNIK(nik string) (*models.User, error) {
	var user models.User
	err := r.db.Where("nik = ?", nik).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *RecordRepository) GetLastRecord(userID string) (*models.RecordKesehatan, error) {
	var record models.RecordKesehatan
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}