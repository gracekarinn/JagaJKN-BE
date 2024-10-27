// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userNIK\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RecordAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"userNIK\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"addRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"userHash\",\"type\":\"bytes32\"}],\"name\":\"addUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"}],\"name\":\"getUserRegistrationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"verifyRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"userHash\",\"type\":\"bytes32\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610e958061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c8063026d4e5c14610059578063122a183a1461007557806353d0e5bc146100a5578063a46417b8146100d5578063b2004ace14610105575b5f5ffd5b610073600480360381019061006e9190610745565b610121565b005b61008f600480360381019061008a91906107cd565b6102e2565b60405161009c9190610841565b60405180910390f35b6100bf60048036038101906100ba91906107cd565b61037a565b6040516100cc9190610841565b60405180910390f35b6100ef60048036038101906100ea919061085a565b610413565b6040516100fc91906108b9565b60405180910390f35b61011f600480360381019061011a91906107cd565b6104aa565b005b5f836040516101309190610924565b90815260200160405180910390206004015f9054906101000a900460ff161561018e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018590610994565b60405180910390fd5b60018260405161019e9190610924565b90815260200160405180910390206002015f9054906101000a900460ff166101fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f2906109fc565b60405180910390fd5b6040518060a00160405280848152602001838152602001828152602001428152602001600115158152505f846040516102349190610924565b90815260200160405180910390205f820151815f0190816102559190610c17565b50602082015181600101908161026b9190610c17565b5060408201518160020155606082015181600301556080820151816004015f6101000a81548160ff0219169083151502179055509050507f2fda3ce68e3316f21c26863e3b64c0bca4dfc4d6cb5777ef08d41c56482e22218383426040516102d593929190610d1e565b60405180910390a1505050565b5f5f836040516102f29190610924565b90815260200160405180910390206004015f9054906101000a900460ff1661034f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034690610dab565b60405180910390fd5b815f8460405161035f9190610924565b90815260200160405180910390206002015414905092915050565b5f60018360405161038b9190610924565b90815260200160405180910390206002015f9054906101000a900460ff166103e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103df906109fc565b60405180910390fd5b816001846040516103f99190610924565b90815260200160405180910390205f015414905092915050565b5f6001826040516104249190610924565b90815260200160405180910390206002015f9054906101000a900460ff16610481576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610478906109fc565b60405180910390fd5b6001826040516104919190610924565b9081526020016040518091039020600101549050919050565b6001826040516104ba9190610924565b90815260200160405180910390206002015f9054906101000a900460ff1615610518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050f90610e13565b60405180910390fd5b6040518060600160405280828152602001428152602001600115158152506001836040516105469190610924565b90815260200160405180910390205f820151815f0155602082015181600101556040820151816002015f6101000a81548160ff0219169083151502179055509050507fb48edca3bdc8658bb712b2fceff22de028026d6f971f9bf32aab68125fc49bc682426040516105b9929190610e31565b60405180910390a15050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610624826105de565b810181811067ffffffffffffffff82111715610643576106426105ee565b5b80604052505050565b5f6106556105c5565b9050610661828261061b565b919050565b5f67ffffffffffffffff8211156106805761067f6105ee565b5b610689826105de565b9050602081019050919050565b828183375f83830152505050565b5f6106b66106b184610666565b61064c565b9050828152602081018484840111156106d2576106d16105da565b5b6106dd848285610696565b509392505050565b5f82601f8301126106f9576106f86105d6565b5b81356107098482602086016106a4565b91505092915050565b5f819050919050565b61072481610712565b811461072e575f5ffd5b50565b5f8135905061073f8161071b565b92915050565b5f5f5f6060848603121561075c5761075b6105ce565b5b5f84013567ffffffffffffffff811115610779576107786105d2565b5b610785868287016106e5565b935050602084013567ffffffffffffffff8111156107a6576107a56105d2565b5b6107b2868287016106e5565b92505060406107c386828701610731565b9150509250925092565b5f5f604083850312156107e3576107e26105ce565b5b5f83013567ffffffffffffffff811115610800576107ff6105d2565b5b61080c858286016106e5565b925050602061081d85828601610731565b9150509250929050565b5f8115159050919050565b61083b81610827565b82525050565b5f6020820190506108545f830184610832565b92915050565b5f6020828403121561086f5761086e6105ce565b5b5f82013567ffffffffffffffff81111561088c5761088b6105d2565b5b610898848285016106e5565b91505092915050565b5f819050919050565b6108b3816108a1565b82525050565b5f6020820190506108cc5f8301846108aa565b92915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f6108fe826108d2565b61090881856108dc565b93506109188185602086016108e6565b80840191505092915050565b5f61092f82846108f4565b915081905092915050565b5f82825260208201905092915050565b7f5265636f726420616c72656164792065786973747300000000000000000000005f82015250565b5f61097e60158361093a565b91506109898261094a565b602082019050919050565b5f6020820190508181035f8301526109ab81610972565b9050919050565b7f55736572206e6f742072656769737465726564000000000000000000000000005f82015250565b5f6109e660138361093a565b91506109f1826109b2565b602082019050919050565b5f6020820190508181035f830152610a13816109da565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610a5e57607f821691505b602082108103610a7157610a70610a1a565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610ad37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a98565b610add8683610a98565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610b18610b13610b0e846108a1565b610af5565b6108a1565b9050919050565b5f819050919050565b610b3183610afe565b610b45610b3d82610b1f565b848454610aa4565b825550505050565b5f5f905090565b610b5c610b4d565b610b67818484610b28565b505050565b5b81811015610b8a57610b7f5f82610b54565b600181019050610b6d565b5050565b601f821115610bcf57610ba081610a77565b610ba984610a89565b81016020851015610bb8578190505b610bcc610bc485610a89565b830182610b6c565b50505b505050565b5f82821c905092915050565b5f610bef5f1984600802610bd4565b1980831691505092915050565b5f610c078383610be0565b9150826002028217905092915050565b610c20826108d2565b67ffffffffffffffff811115610c3957610c386105ee565b5b610c438254610a47565b610c4e828285610b8e565b5f60209050601f831160018114610c7f575f8415610c6d578287015190505b610c778582610bfc565b865550610cde565b601f198416610c8d86610a77565b5f5b82811015610cb457848901518255600182019150602085019450602081019050610c8f565b86831015610cd15784890151610ccd601f891682610be0565b8355505b6001600288020188555050505b505050505050565b5f610cf0826108d2565b610cfa818561093a565b9350610d0a8185602086016108e6565b610d13816105de565b840191505092915050565b5f6060820190508181035f830152610d368186610ce6565b90508181036020830152610d4a8185610ce6565b9050610d5960408301846108aa565b949350505050565b7f5265636f726420646f6573206e6f7420657869737400000000000000000000005f82015250565b5f610d9560158361093a565b9150610da082610d61565b602082019050919050565b5f6020820190508181035f830152610dc281610d89565b9050919050565b7f5573657220616c726561647920726567697374657265640000000000000000005f82015250565b5f610dfd60178361093a565b9150610e0882610dc9565b602082019050919050565b5f6020820190508181035f830152610e2a81610df1565b9050919050565b5f6040820190508181035f830152610e498185610ce6565b9050610e5860208301846108aa565b939250505056fea26469706673582212203dddc854d9d57483e9cebc7bdd7ee431a1b46e0d31ca1073139fe56184c6d85164736f6c634300081c0033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetUserRegistrationTime is a free data retrieval call binding the contract method 0xa46417b8.
//
// Solidity: function getUserRegistrationTime(string nik) view returns(uint256)
func (_Contracts *ContractsCaller) GetUserRegistrationTime(opts *bind.CallOpts, nik string) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getUserRegistrationTime", nik)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserRegistrationTime is a free data retrieval call binding the contract method 0xa46417b8.
//
// Solidity: function getUserRegistrationTime(string nik) view returns(uint256)
func (_Contracts *ContractsSession) GetUserRegistrationTime(nik string) (*big.Int, error) {
	return _Contracts.Contract.GetUserRegistrationTime(&_Contracts.CallOpts, nik)
}

// GetUserRegistrationTime is a free data retrieval call binding the contract method 0xa46417b8.
//
// Solidity: function getUserRegistrationTime(string nik) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetUserRegistrationTime(nik string) (*big.Int, error) {
	return _Contracts.Contract.GetUserRegistrationTime(&_Contracts.CallOpts, nik)
}

// VerifyRecord is a free data retrieval call binding the contract method 0x122a183a.
//
// Solidity: function verifyRecord(string noSEP, bytes32 dataHash) view returns(bool)
func (_Contracts *ContractsCaller) VerifyRecord(opts *bind.CallOpts, noSEP string, dataHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "verifyRecord", noSEP, dataHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyRecord is a free data retrieval call binding the contract method 0x122a183a.
//
// Solidity: function verifyRecord(string noSEP, bytes32 dataHash) view returns(bool)
func (_Contracts *ContractsSession) VerifyRecord(noSEP string, dataHash [32]byte) (bool, error) {
	return _Contracts.Contract.VerifyRecord(&_Contracts.CallOpts, noSEP, dataHash)
}

// VerifyRecord is a free data retrieval call binding the contract method 0x122a183a.
//
// Solidity: function verifyRecord(string noSEP, bytes32 dataHash) view returns(bool)
func (_Contracts *ContractsCallerSession) VerifyRecord(noSEP string, dataHash [32]byte) (bool, error) {
	return _Contracts.Contract.VerifyRecord(&_Contracts.CallOpts, noSEP, dataHash)
}

// VerifyUser is a free data retrieval call binding the contract method 0x53d0e5bc.
//
// Solidity: function verifyUser(string nik, bytes32 userHash) view returns(bool)
func (_Contracts *ContractsCaller) VerifyUser(opts *bind.CallOpts, nik string, userHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "verifyUser", nik, userHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyUser is a free data retrieval call binding the contract method 0x53d0e5bc.
//
// Solidity: function verifyUser(string nik, bytes32 userHash) view returns(bool)
func (_Contracts *ContractsSession) VerifyUser(nik string, userHash [32]byte) (bool, error) {
	return _Contracts.Contract.VerifyUser(&_Contracts.CallOpts, nik, userHash)
}

// VerifyUser is a free data retrieval call binding the contract method 0x53d0e5bc.
//
// Solidity: function verifyUser(string nik, bytes32 userHash) view returns(bool)
func (_Contracts *ContractsCallerSession) VerifyUser(nik string, userHash [32]byte) (bool, error) {
	return _Contracts.Contract.VerifyUser(&_Contracts.CallOpts, nik, userHash)
}

// AddRecord is a paid mutator transaction binding the contract method 0x026d4e5c.
//
// Solidity: function addRecord(string noSEP, string userNIK, bytes32 dataHash) returns()
func (_Contracts *ContractsTransactor) AddRecord(opts *bind.TransactOpts, noSEP string, userNIK string, dataHash [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addRecord", noSEP, userNIK, dataHash)
}

// AddRecord is a paid mutator transaction binding the contract method 0x026d4e5c.
//
// Solidity: function addRecord(string noSEP, string userNIK, bytes32 dataHash) returns()
func (_Contracts *ContractsSession) AddRecord(noSEP string, userNIK string, dataHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.AddRecord(&_Contracts.TransactOpts, noSEP, userNIK, dataHash)
}

// AddRecord is a paid mutator transaction binding the contract method 0x026d4e5c.
//
// Solidity: function addRecord(string noSEP, string userNIK, bytes32 dataHash) returns()
func (_Contracts *ContractsTransactorSession) AddRecord(noSEP string, userNIK string, dataHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.AddRecord(&_Contracts.TransactOpts, noSEP, userNIK, dataHash)
}

// AddUser is a paid mutator transaction binding the contract method 0xb2004ace.
//
// Solidity: function addUser(string nik, bytes32 userHash) returns()
func (_Contracts *ContractsTransactor) AddUser(opts *bind.TransactOpts, nik string, userHash [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addUser", nik, userHash)
}

// AddUser is a paid mutator transaction binding the contract method 0xb2004ace.
//
// Solidity: function addUser(string nik, bytes32 userHash) returns()
func (_Contracts *ContractsSession) AddUser(nik string, userHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.AddUser(&_Contracts.TransactOpts, nik, userHash)
}

// AddUser is a paid mutator transaction binding the contract method 0xb2004ace.
//
// Solidity: function addUser(string nik, bytes32 userHash) returns()
func (_Contracts *ContractsTransactorSession) AddUser(nik string, userHash [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.AddUser(&_Contracts.TransactOpts, nik, userHash)
}

// ContractsRecordAddedIterator is returned from FilterRecordAdded and is used to iterate over the raw logs and unpacked data for RecordAdded events raised by the Contracts contract.
type ContractsRecordAddedIterator struct {
	Event *ContractsRecordAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsRecordAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRecordAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsRecordAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsRecordAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRecordAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRecordAdded represents a RecordAdded event raised by the Contracts contract.
type ContractsRecordAdded struct {
	NoSEP     string
	UserNIK   string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecordAdded is a free log retrieval operation binding the contract event 0x2fda3ce68e3316f21c26863e3b64c0bca4dfc4d6cb5777ef08d41c56482e2221.
//
// Solidity: event RecordAdded(string noSEP, string userNIK, uint256 timestamp)
func (_Contracts *ContractsFilterer) FilterRecordAdded(opts *bind.FilterOpts) (*ContractsRecordAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "RecordAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsRecordAddedIterator{contract: _Contracts.contract, event: "RecordAdded", logs: logs, sub: sub}, nil
}

// WatchRecordAdded is a free log subscription operation binding the contract event 0x2fda3ce68e3316f21c26863e3b64c0bca4dfc4d6cb5777ef08d41c56482e2221.
//
// Solidity: event RecordAdded(string noSEP, string userNIK, uint256 timestamp)
func (_Contracts *ContractsFilterer) WatchRecordAdded(opts *bind.WatchOpts, sink chan<- *ContractsRecordAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "RecordAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRecordAdded)
				if err := _Contracts.contract.UnpackLog(event, "RecordAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRecordAdded is a log parse operation binding the contract event 0x2fda3ce68e3316f21c26863e3b64c0bca4dfc4d6cb5777ef08d41c56482e2221.
//
// Solidity: event RecordAdded(string noSEP, string userNIK, uint256 timestamp)
func (_Contracts *ContractsFilterer) ParseRecordAdded(log types.Log) (*ContractsRecordAdded, error) {
	event := new(ContractsRecordAdded)
	if err := _Contracts.contract.UnpackLog(event, "RecordAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUserRegisteredIterator is returned from FilterUserRegistered and is used to iterate over the raw logs and unpacked data for UserRegistered events raised by the Contracts contract.
type ContractsUserRegisteredIterator struct {
	Event *ContractsUserRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUserRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsUserRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUserRegistered represents a UserRegistered event raised by the Contracts contract.
type ContractsUserRegistered struct {
	Nik       string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUserRegistered is a free log retrieval operation binding the contract event 0xb48edca3bdc8658bb712b2fceff22de028026d6f971f9bf32aab68125fc49bc6.
//
// Solidity: event UserRegistered(string nik, uint256 timestamp)
func (_Contracts *ContractsFilterer) FilterUserRegistered(opts *bind.FilterOpts) (*ContractsUserRegisteredIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "UserRegistered")
	if err != nil {
		return nil, err
	}
	return &ContractsUserRegisteredIterator{contract: _Contracts.contract, event: "UserRegistered", logs: logs, sub: sub}, nil
}

// WatchUserRegistered is a free log subscription operation binding the contract event 0xb48edca3bdc8658bb712b2fceff22de028026d6f971f9bf32aab68125fc49bc6.
//
// Solidity: event UserRegistered(string nik, uint256 timestamp)
func (_Contracts *ContractsFilterer) WatchUserRegistered(opts *bind.WatchOpts, sink chan<- *ContractsUserRegistered) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "UserRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUserRegistered)
				if err := _Contracts.contract.UnpackLog(event, "UserRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserRegistered is a log parse operation binding the contract event 0xb48edca3bdc8658bb712b2fceff22de028026d6f971f9bf32aab68125fc49bc6.
//
// Solidity: event UserRegistered(string nik, uint256 timestamp)
func (_Contracts *ContractsFilterer) ParseUserRegistered(log types.Log) (*ContractsUserRegistered, error) {
	event := new(ContractsUserRegistered)
	if err := _Contracts.contract.UnpackLog(event, "UserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
