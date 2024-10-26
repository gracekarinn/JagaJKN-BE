package service

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"jagajkn/internal/models"
	"jagajkn/internal/repository"
)

type RecordService struct {
	repo *repository.RecordRepository
}

func NewRecordService(repo *repository.RecordRepository) *RecordService {
	return &RecordService{repo: repo}
}

func (s *RecordService) CreateRecord(userID string, input *models.RecordInput) (*models.RecordKesehatan, error) {
	existingRecords, err := s.repo.FindByUserID(userID)
	var previousHash *string
	if err == nil && len(existingRecords) > 0 {
		previousHash = &existingRecords[len(existingRecords)-1].HashCurrent
	}

	record := &models.RecordKesehatan{
		UserID:           userID,
		NoSEP:            input.NoSEP,
		JenisRawat:       input.JenisRawat,
		DiagnosaAwal:     input.DiagnosaAwal,
		DiagnosaPrimer:   input.DiagnosaPrimer,
		IcdX:             input.IcdX,
		Tindakan:         input.Tindakan,
		TanggalMasuk:     time.Now(),
		HashPrevious:     previousHash,
		RetentionYears:   5,
	}

	record.HashCurrent = calculateHash(record)

	if err := s.repo.Create(record); err != nil {
		return nil, err
	}

	return record, nil
}

func (s *RecordService) GetUserRecords(userID string) ([]models.RecordKesehatan, error) {
	return s.repo.FindByUserID(userID)
}

func (s *RecordService) GetRecord(id string) (*models.RecordKesehatan, error) {
	return s.repo.FindByID(id)
}

func calculateHash(record *models.RecordKesehatan) string {
	data := struct {
		UserID         string
		NoSEP          string
		TanggalMasuk   time.Time
		JenisRawat     models.JenisRawat
		DiagnosaAwal   string
		DiagnosaPrimer string
		IcdX           string
		Tindakan       string
		HashPrevious   *string
	}{
		UserID:         record.UserID,
		NoSEP:          record.NoSEP,
		TanggalMasuk:   record.TanggalMasuk,
		JenisRawat:     record.JenisRawat,
		DiagnosaAwal:   record.DiagnosaAwal,
		DiagnosaPrimer: record.DiagnosaPrimer,
		IcdX:           record.IcdX,
		Tindakan:       record.Tindakan,
		HashPrevious:   record.HashPrevious,
	}

	jsonData, _ := json.Marshal(data)
	hash := sha256.Sum256(jsonData)
	return hex.EncodeToString(hash[:])
}