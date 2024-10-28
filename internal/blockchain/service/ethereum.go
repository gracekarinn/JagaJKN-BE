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
    Contract   *contracts.Contracts  
    contractAddr common.Address
}


func NewBlockchainService(cfg *config.BlockchainConfig) (*BlockchainService, error) {
    client, err := ethclient.Dial(cfg.ProviderURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to blockchain: %v", err)
    }

    privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
    if err != nil {
        return nil, fmt.Errorf("failed to load private key: %v", err)
    }

    var contract *contracts.Contracts
    var contractAddr common.Address

    if cfg.ContractAddress != "" {
        contractAddr = common.HexToAddress(cfg.ContractAddress)
        contract, err = contracts.NewContracts(contractAddr, client)
        if err != nil {
            return nil, fmt.Errorf("failed to load contract: %v", err)
        }
        log.Printf("Contract loaded at address: %s", contractAddr.Hex())
    }

    return &BlockchainService{
        client:       client,
        privateKey:   privateKey,
        Contract:     contract,          
        contractAddr: contractAddr,
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
    if s.Contract == nil {
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
    tx, err := s.Contract.AddRecord(auth, record.NoSEP, record.UserNIK, hash)
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
    if s.Contract == nil {
        return false, fmt.Errorf("contract not initialized")
    }

    hash := s.calculateRecordHash(record)

    return s.Contract.VerifyRecord(&bind.CallOpts{
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

func (s *BlockchainService) SaveUserRegistration(ctx context.Context, nik, userHash string) error {
    if s.Contract == nil {
        return fmt.Errorf("contract not initialized")
    }

    auth, err := s.createTransactOpts(ctx)
    if err != nil {
        return fmt.Errorf("failed to create transaction options: %w", err)
    }

    // Convert string hash to bytes32 using Keccak256
    hashData := []byte(userHash)
    var hashBytes [32]byte
    hash := crypto.Keccak256(hashData)
    copy(hashBytes[:], hash)

    // Log the hash being sent
    log.Printf("Sending hash to blockchain: %x", hashBytes)

    tx, err := s.Contract.AddUser(auth, nik, hashBytes)
    if err != nil {
        return fmt.Errorf("failed to save user to blockchain: %w", err)
    }

    _, err = bind.WaitMined(ctx, s.client, tx)
    if err != nil {
        return fmt.Errorf("failed waiting for transaction: %w", err)
    }

    return nil
}

func (s *BlockchainService) VerifyUserRegistration(ctx context.Context, nik, userHash string) (bool, error) {
    if s.Contract == nil {
        return false, fmt.Errorf("contract not initialized")
    }

    // Convert string hash to bytes32 using the same method as registration
    hashData := []byte(userHash)
    var hashBytes [32]byte
    hash := crypto.Keccak256(hashData)
    copy(hashBytes[:], hash)

    // Log the hash being verified
    log.Printf("Verifying hash: %x", hashBytes)

    return s.Contract.VerifyUser(&bind.CallOpts{
        Context: ctx,
    }, nik, hashBytes)
}

func (s *BlockchainService) GetContract() *contracts.Contracts {
    return s.Contract
}

func (s *BlockchainService) CheckContractStatus(ctx context.Context) (map[string]interface{}, error) {
    if s.Contract == nil {
        return nil, fmt.Errorf("contract not initialized")
    }


    chainID, err := s.client.ChainID(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to get chain ID: %v", err)
    }


    block, err := s.client.BlockByNumber(ctx, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to get latest block: %v", err)
    }


    code, err := s.client.CodeAt(ctx, s.contractAddr, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to get contract code: %v", err)
    }


    publicKey := s.privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("error casting public key to ECDSA")
    }
    callerAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    nonce, err := s.client.PendingNonceAt(ctx, callerAddress)
    if err != nil {
        return nil, fmt.Errorf("failed to get nonce: %v", err)
    }

    status := map[string]interface{}{
        "chainID":         chainID.String(),
        "latestBlock":     block.Number().String(),
        "contractAddress": s.contractAddr.Hex(),
        "hasCode":         len(code) > 0,
        "codeSize":        len(code),
        "callerAddress":   callerAddress.Hex(),
        "nonce":          nonce,
        "isConnected":    s.client != nil,
    }


    balance, err := s.client.BalanceAt(ctx, callerAddress, nil)
    if err != nil {
        status["balanceError"] = err.Error()
    } else {
        status["balance"] = balance.String()
    }

    gasPrice, err := s.client.SuggestGasPrice(ctx)
    if err != nil {
        status["gasPriceError"] = err.Error()
    } else {
        status["gasPrice"] = gasPrice.String()
    }

    return status, nil
}