package service

import (
	"context"
	"fmt"
	"time"

	"jagajkn/internal/blockchain/service"
	"jagajkn/internal/models"
	"jagajkn/internal/repository"
)

type RecordService struct {
    recordRepo    repository.RecordRepository
    blockchainSvc *service.BlockchainService
}

func NewRecordService(
    recordRepo repository.RecordRepository,
    blockchainSvc *service.BlockchainService,
) *RecordService {
    return &RecordService{
        recordRepo:    recordRepo,
        blockchainSvc: blockchainSvc,
    }
}

func (s *RecordService) CreateRecord(ctx context.Context, input *models.RecordInput, userNIK string) (*models.RecordKesehatan, error) {

    record := &models.RecordKesehatan{
        NoSEP:          input.NoSEP,
        UserNIK:        userNIK,
        TanggalMasuk:   time.Now(),
        JenisRawat:     input.JenisRawat,
        DiagnosaAwal:   input.DiagnosaAwal,
        DiagnosaPrimer: input.DiagnosaPrimer,
        IcdX:           input.IcdX,
        Tindakan:       input.Tindakan,
    }


    if err := s.recordRepo.Create(ctx, record); err != nil {
        return nil, err
    }


    if err := s.blockchainSvc.SaveMedicalRecord(ctx, record); err != nil {
        return nil, err
    }

    return record, nil
}


func (s *RecordService) GetRecord(ctx context.Context, noSEP string) (*models.RecordKesehatan, error) {
    record, err := s.recordRepo.GetByNoSEP(ctx, noSEP)
    if err != nil {
        return nil, err
    }


    verified, err := s.blockchainSvc.VerifyMedicalRecord(ctx, record)
    if err != nil {
        return nil, err
    }

    if !verified {
        return nil, fmt.Errorf("record integrity verification failed")
    }

    return record, nil
}


func (s *RecordService) GetUserRecords(ctx context.Context, userNIK string) ([]*models.RecordKesehatan, error) {
    return s.recordRepo.GetByUserNIK(ctx, userNIK)
}