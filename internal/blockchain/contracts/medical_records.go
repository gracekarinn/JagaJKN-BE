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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userNIK\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RecordAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"userNIK\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"addRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"userHash\",\"type\":\"bytes32\"}],\"name\":\"addUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"}],\"name\":\"isUserRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"verifyRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nik\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"userHash\",\"type\":\"bytes32\"}],\"name\":\"verifyUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e93806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063026d4e5c1461005c578063122a183a1461007857806353d0e5bc146100a8578063a8cc7784146100d8578063b2004ace14610108575b600080fd5b61007660048036038101906100719190610709565b610124565b005b610092600480360381019061008d9190610794565b6102ec565b60405161009f919061080b565b60405180910390f35b6100c260048036038101906100bd9190610794565b610388565b6040516100cf919061080b565b60405180910390f35b6100f260048036038101906100ed9190610826565b610423565b6040516100ff919061080b565b60405180910390f35b610122600480360381019061011d9190610794565b61045a565b005b60018360405161013491906108e0565b908152602001604051809103902060040160009054906101000a900460ff1615610193576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161018a90610954565b60405180910390fd5b6000826040516101a391906108e0565b908152602001604051809103902060020160009054906101000a900460ff16610201576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f8906109c0565b60405180910390fd5b6040518060a001604052808481526020018381526020018281526020014281526020016001151581525060018460405161023b91906108e0565b9081526020016040518091039020600082015181600001908161025e9190610bf6565b5060208201518160010190816102749190610bf6565b50604082015181600201556060820151816003015560808201518160040160006101000a81548160ff0219169083151502179055509050507f2fda3ce68e3316f21c26863e3b64c0bca4dfc4d6cb5777ef08d41c56482e22218383426040516102df93929190610d10565b60405180910390a1505050565b60006001836040516102fe91906108e0565b908152602001604051809103902060040160009054906101000a900460ff1661035c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035390610da1565b60405180910390fd5b8160018460405161036d91906108e0565b90815260200160405180910390206002015414905092915050565b6000808360405161039991906108e0565b908152602001604051809103902060020160009054906101000a900460ff166103f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ee906109c0565b60405180910390fd5b8160008460405161040891906108e0565b90815260200160405180910390206000015414905092915050565b6000808260405161043491906108e0565b908152602001604051809103902060020160009054906101000a900460ff169050919050565b60008260405161046a91906108e0565b908152602001604051809103902060020160009054906101000a900460ff16156104c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104c090610e0d565b60405180910390fd5b6040518060600160405280828152602001428152602001600115158152506000836040516104f791906108e0565b9081526020016040518091039020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff0219169083151502179055509050507fb48edca3bdc8658bb712b2fceff22de028026d6f971f9bf32aab68125fc49bc6824260405161056d929190610e2d565b60405180910390a15050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6105e082610597565b810181811067ffffffffffffffff821117156105ff576105fe6105a8565b5b80604052505050565b6000610612610579565b905061061e82826105d7565b919050565b600067ffffffffffffffff82111561063e5761063d6105a8565b5b61064782610597565b9050602081019050919050565b82818337600083830152505050565b600061067661067184610623565b610608565b90508281526020810184848401111561069257610691610592565b5b61069d848285610654565b509392505050565b600082601f8301126106ba576106b961058d565b5b81356106ca848260208601610663565b91505092915050565b6000819050919050565b6106e6816106d3565b81146106f157600080fd5b50565b600081359050610703816106dd565b92915050565b60008060006060848603121561072257610721610583565b5b600084013567ffffffffffffffff8111156107405761073f610588565b5b61074c868287016106a5565b935050602084013567ffffffffffffffff81111561076d5761076c610588565b5b610779868287016106a5565b925050604061078a868287016106f4565b9150509250925092565b600080604083850312156107ab576107aa610583565b5b600083013567ffffffffffffffff8111156107c9576107c8610588565b5b6107d5858286016106a5565b92505060206107e6858286016106f4565b9150509250929050565b60008115159050919050565b610805816107f0565b82525050565b600060208201905061082060008301846107fc565b92915050565b60006020828403121561083c5761083b610583565b5b600082013567ffffffffffffffff81111561085a57610859610588565b5b610866848285016106a5565b91505092915050565b600081519050919050565b600081905092915050565b60005b838110156108a3578082015181840152602081019050610888565b60008484015250505050565b60006108ba8261086f565b6108c4818561087a565b93506108d4818560208601610885565b80840191505092915050565b60006108ec82846108af565b915081905092915050565b600082825260208201905092915050565b7f5265636f726420616c7265616479206578697374730000000000000000000000600082015250565b600061093e6015836108f7565b915061094982610908565b602082019050919050565b6000602082019050818103600083015261096d81610931565b9050919050565b7f55736572206e6f74207265676973746572656400000000000000000000000000600082015250565b60006109aa6013836108f7565b91506109b582610974565b602082019050919050565b600060208201905081810360008301526109d98161099d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610a2757607f821691505b602082108103610a3a57610a396109e0565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610aa27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a65565b610aac8683610a65565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610af3610aee610ae984610ac4565b610ace565b610ac4565b9050919050565b6000819050919050565b610b0d83610ad8565b610b21610b1982610afa565b848454610a72565b825550505050565b600090565b610b36610b29565b610b41818484610b04565b505050565b5b81811015610b6557610b5a600082610b2e565b600181019050610b47565b5050565b601f821115610baa57610b7b81610a40565b610b8484610a55565b81016020851015610b93578190505b610ba7610b9f85610a55565b830182610b46565b50505b505050565b600082821c905092915050565b6000610bcd60001984600802610baf565b1980831691505092915050565b6000610be68383610bbc565b9150826002028217905092915050565b610bff8261086f565b67ffffffffffffffff811115610c1857610c176105a8565b5b610c228254610a0f565b610c2d828285610b69565b600060209050601f831160018114610c605760008415610c4e578287015190505b610c588582610bda565b865550610cc0565b601f198416610c6e86610a40565b60005b82811015610c9657848901518255600182019150602085019450602081019050610c71565b86831015610cb35784890151610caf601f891682610bbc565b8355505b6001600288020188555050505b505050505050565b6000610cd38261086f565b610cdd81856108f7565b9350610ced818560208601610885565b610cf681610597565b840191505092915050565b610d0a81610ac4565b82525050565b60006060820190508181036000830152610d2a8186610cc8565b90508181036020830152610d3e8185610cc8565b9050610d4d6040830184610d01565b949350505050565b7f5265636f726420646f6573206e6f742065786973740000000000000000000000600082015250565b6000610d8b6015836108f7565b9150610d9682610d55565b602082019050919050565b60006020820190508181036000830152610dba81610d7e565b9050919050565b7f5573657220616c72656164792072656769737465726564000000000000000000600082015250565b6000610df76017836108f7565b9150610e0282610dc1565b602082019050919050565b60006020820190508181036000830152610e2681610dea565b9050919050565b60006040820190508181036000830152610e478185610cc8565b9050610e566020830184610d01565b939250505056fea2646970667358221220ec9e315f4ff53870032b58cffca1ae84ae630a779d4412252da28f330b055a2764736f6c63430008130033",
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

// IsUserRegistered is a free data retrieval call binding the contract method 0xa8cc7784.
//
// Solidity: function isUserRegistered(string nik) view returns(bool)
func (_Contracts *ContractsCaller) IsUserRegistered(opts *bind.CallOpts, nik string) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isUserRegistered", nik)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserRegistered is a free data retrieval call binding the contract method 0xa8cc7784.
//
// Solidity: function isUserRegistered(string nik) view returns(bool)
func (_Contracts *ContractsSession) IsUserRegistered(nik string) (bool, error) {
	return _Contracts.Contract.IsUserRegistered(&_Contracts.CallOpts, nik)
}

// IsUserRegistered is a free data retrieval call binding the contract method 0xa8cc7784.
//
// Solidity: function isUserRegistered(string nik) view returns(bool)
func (_Contracts *ContractsCallerSession) IsUserRegistered(nik string) (bool, error) {
	return _Contracts.Contract.IsUserRegistered(&_Contracts.CallOpts, nik)
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
