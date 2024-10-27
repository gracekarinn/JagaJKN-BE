package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"jagajkn/internal/blockchain/contracts"
	"jagajkn/internal/config"
	"jagajkn/internal/models"
)

type BlockchainService struct {
    client     *ethclient.Client
    privateKey *ecdsa.PrivateKey
    contract   *contracts.Contracts  
}


func NewBlockchainService(cfg *config.BlockchainConfig) (*BlockchainService, error) {
    // Connect ke blockchain
    client, err := ethclient.Dial(cfg.ProviderURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to blockchain: %v", err)
    }

    // Load private key
    privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
    if err != nil {
        return nil, fmt.Errorf("failed to load private key: %v", err)
    }

    // Load contract
    var contract *contracts.Contracts
    if cfg.ContractAddress != "" {
        contract, err = contracts.NewContracts(
            common.HexToAddress(cfg.ContractAddress),
            client,
        )
        if err != nil {
            return nil, fmt.Errorf("failed to load contract: %v", err)
        }
    }

    return &BlockchainService{
        client:     client,
        privateKey: privateKey,
        contract:   contract,
    }, nil
}


func (s *BlockchainService) TestConnection() error {
    blockNumber, err := s.client.BlockNumber(context.Background())
    if err != nil {
        return fmt.Errorf("failed to get block number: %v", err)
    }
    log.Printf("Connected to blockchain. Latest block: %d", blockNumber)
    return nil
}


func (s *BlockchainService) SaveMedicalRecord(ctx context.Context, record *models.RecordKesehatan) error {
    if s.contract == nil {
        return fmt.Errorf("contract not initialized")
    }

    // Buat transaction options
    auth, err := s.createTransactOpts(ctx)
    if err != nil {
        return err
    }

    // Calculate hash dari record
    hash := s.calculateRecordHash(record)

    // Add record to blockchain
    tx, err := s.contract.AddRecord(auth, record.NoSEP, record.UserNIK, hash)
    if err != nil {
        return fmt.Errorf("failed to add record: %v", err)
    }

    // Nunggu transaksi selesai
    _, err = bind.WaitMined(ctx, s.client, tx)
    if err != nil {
        return fmt.Errorf("failed to wait for transaction: %v", err)
    }

    return nil
}


func (s *BlockchainService) VerifyMedicalRecord(ctx context.Context, record *models.RecordKesehatan) (bool, error) {
    if s.contract == nil {
        return false, fmt.Errorf("contract not initialized")
    }

    hash := s.calculateRecordHash(record)

    // Call contract's verify function
    return s.contract.VerifyRecord(&bind.CallOpts{
        Context: ctx,
    }, record.NoSEP, hash)
}


func (s *BlockchainService) calculateRecordHash(record *models.RecordKesehatan) [32]byte {
    data := fmt.Sprintf("%s-%s-%s-%s-%s-%s",
        record.NoSEP,
        record.UserNIK,
        record.DiagnosaAwal,
        record.DiagnosaPrimer,
        record.IcdX,
        record.Tindakan,
    )
    return crypto.Keccak256Hash([]byte(data))
}


func (s *BlockchainService) createTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
    nonce, err := s.client.PendingNonceAt(ctx, crypto.PubkeyToAddress(s.privateKey.PublicKey))
    if err != nil {
        return nil, err
    }

    gasPrice, err := s.client.SuggestGasPrice(ctx)
    if err != nil {
        return nil, err
    }

    chainID, err := s.client.ChainID(ctx)
    if err != nil {
        return nil, err
    }

    auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, chainID)
    if err != nil {
        return nil, err
    }

    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)      
    auth.GasLimit = uint64(3000000)
    auth.GasPrice = gasPrice

    return auth, nil
}