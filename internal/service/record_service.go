package service

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"jagajkn/internal/blockchain/service"
	"jagajkn/internal/models"
	"jagajkn/internal/repository"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"gorm.io/gorm"
)

type RecordService struct {
    recordRepo    repository.RecordRepository
    blockchainSvc *service.BlockchainService
    db           *gorm.DB
}

func NewRecordService(
    db *gorm.DB,   
    recordRepo repository.RecordRepository,
    blockchainSvc *service.BlockchainService,
) *RecordService {
    return &RecordService{
        db:            db,
        recordRepo:    recordRepo,
        blockchainSvc: blockchainSvc,
    }
}

func (s *RecordService) CreateRecord(ctx context.Context, input *models.RecordInput, userNIK string) (*models.RecordKesehatan, error) {
    var user models.User
    if err := s.db.Where("nik = ?", userNIK).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user not found: %w", err)
    }

    record := &models.RecordKesehatan{
        NoSEP:          input.NoSEP,
        UserNIK:        userNIK,
        User:           user,  
        TanggalMasuk:   time.Now(),
        JenisRawat:     input.JenisRawat,
        DiagnosaAwal:   input.DiagnosaAwal,
        DiagnosaPrimer: input.DiagnosaPrimer,
        IcdX:           input.IcdX,
        Tindakan:       input.Tindakan,
        RetentionYears: 5, 
    }


    hashData := fmt.Sprintf("%s-%s-%s-%s-%s-%s",
        record.NoSEP,
        record.UserNIK,
        record.DiagnosaAwal,
        record.DiagnosaPrimer,
        record.IcdX,
        record.Tindakan,
    )
    record.HashCurrent = hex.EncodeToString(crypto.Keccak256([]byte(hashData)))


    if err := s.recordRepo.Create(ctx, record); err != nil {
        return nil, fmt.Errorf("failed to create record: %w", err)
    }

    if err := s.blockchainSvc.SaveMedicalRecord(ctx, record); err != nil {
        return nil, fmt.Errorf("failed to save to blockchain: %w", err)
    }

    return s.recordRepo.GetByNoSEP(ctx, record.NoSEP)
}


func (s *RecordService) GetRecord(ctx context.Context, noSEP string) (*models.RecordKesehatan, error) {
    record, err := s.recordRepo.GetByNoSEP(ctx, noSEP)
    if err != nil {
        return nil, err
    }

    hashData := fmt.Sprintf("%s-%s-%s-%s-%s-%s",
        record.NoSEP,
        record.UserNIK,
        record.DiagnosaAwal,
        record.DiagnosaPrimer,
        record.IcdX,
        record.Tindakan,
    )
    hash := crypto.Keccak256Hash([]byte(hashData))

    verified, err := s.blockchainSvc.Contract.VerifyRecord(&bind.CallOpts{
        Context: ctx,
    }, record.NoSEP, hash)
    
    if err != nil {
        log.Printf("Blockchain verification error: %v", err)
    } else {
        log.Printf("Record blockchain verification: %v", verified)
    }

    record.BlockchainVerified = verified

    return record, nil
}

func (s *RecordService) GetUserRecords(ctx context.Context, userNIK string) ([]*models.RecordKesehatan, error) {
    return s.recordRepo.GetByUserNIK(ctx, userNIK)
}