package service

import (
	"jagajkn/internal/models"
	"jagajkn/internal/repository"
	"time"
)

type RecordService struct {
    repo *repository.RecordRepository
}

func NewRecordService(repo *repository.RecordRepository) *RecordService {
    return &RecordService{
        repo: repo,
    }
}

func (s *RecordService) CreateRecord(userNIK string, input *models.RecordInput) (*models.RecordKesehatan, error) {
    record := &models.RecordKesehatan{
        NoSEP:          input.NoSEP,
        UserNIK:        userNIK,
        JenisRawat:     input.JenisRawat,
        DiagnosaAwal:   input.DiagnosaAwal,
        DiagnosaPrimer: input.DiagnosaPrimer,
        IcdX:           input.IcdX,
        Tindakan:       input.Tindakan,
        TanggalMasuk:   time.Now(),
    }

    if err := s.repo.Create(record); err != nil {
        return nil, err
    }

    return record, nil
}

func (s *RecordService) GetUserRecords(userNIK string) ([]models.RecordKesehatan, error) {
    return s.repo.FindByUserNIK(userNIK)
}

func (s *RecordService) GetRecord(noSEP string) (*models.RecordKesehatan, error) {
    return s.repo.FindByNoSEP(noSEP)
}