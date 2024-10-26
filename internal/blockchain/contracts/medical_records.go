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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"userNIK\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RecordAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"userNIK\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"addRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"noSEP\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"verifyRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506109e28061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063026d4e5c14610038578063122a183a14610054575b5f5ffd5b610052600480360381019061004d91906103f0565b610084565b005b61006e60048036038101906100699190610478565b6101d8565b60405161007b91906104ec565b60405180910390f35b5f836040516100939190610557565b90815260200160405180910390206004015f9054906101000a900460ff16156100f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100e8906105c7565b60405180910390fd5b6040518060a00160405280848152602001838152602001828152602001428152602001600115158152505f8460405161012a9190610557565b90815260200160405180910390205f820151815f01908161014b91906107eb565b50602082015181600101908161016191906107eb565b5060408201518160020155606082015181600301556080820151816004015f6101000a81548160ff0219169083151502179055509050507f2fda3ce68e3316f21c26863e3b64c0bca4dfc4d6cb5777ef08d41c56482e22218383426040516101cb93929190610901565b60405180910390a1505050565b5f5f836040516101e89190610557565b90815260200160405180910390206004015f9054906101000a900460ff16610245576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023c9061098e565b60405180910390fd5b815f846040516102559190610557565b90815260200160405180910390206002015414905092915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102cf82610289565b810181811067ffffffffffffffff821117156102ee576102ed610299565b5b80604052505050565b5f610300610270565b905061030c82826102c6565b919050565b5f67ffffffffffffffff82111561032b5761032a610299565b5b61033482610289565b9050602081019050919050565b828183375f83830152505050565b5f61036161035c84610311565b6102f7565b90508281526020810184848401111561037d5761037c610285565b5b610388848285610341565b509392505050565b5f82601f8301126103a4576103a3610281565b5b81356103b484826020860161034f565b91505092915050565b5f819050919050565b6103cf816103bd565b81146103d9575f5ffd5b50565b5f813590506103ea816103c6565b92915050565b5f5f5f6060848603121561040757610406610279565b5b5f84013567ffffffffffffffff8111156104245761042361027d565b5b61043086828701610390565b935050602084013567ffffffffffffffff8111156104515761045061027d565b5b61045d86828701610390565b925050604061046e868287016103dc565b9150509250925092565b5f5f6040838503121561048e5761048d610279565b5b5f83013567ffffffffffffffff8111156104ab576104aa61027d565b5b6104b785828601610390565b92505060206104c8858286016103dc565b9150509250929050565b5f8115159050919050565b6104e6816104d2565b82525050565b5f6020820190506104ff5f8301846104dd565b92915050565b5f81519050919050565b5f81905092915050565b8281835e5f83830152505050565b5f61053182610505565b61053b818561050f565b935061054b818560208601610519565b80840191505092915050565b5f6105628284610527565b915081905092915050565b5f82825260208201905092915050565b7f5265636f726420616c72656164792065786973747300000000000000000000005f82015250565b5f6105b160158361056d565b91506105bc8261057d565b602082019050919050565b5f6020820190508181035f8301526105de816105a5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061062957607f821691505b60208210810361063c5761063b6105e5565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261069e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610663565b6106a88683610663565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6106ec6106e76106e2846106c0565b6106c9565b6106c0565b9050919050565b5f819050919050565b610705836106d2565b610719610711826106f3565b84845461066f565b825550505050565b5f5f905090565b610730610721565b61073b8184846106fc565b505050565b5b8181101561075e576107535f82610728565b600181019050610741565b5050565b601f8211156107a35761077481610642565b61077d84610654565b8101602085101561078c578190505b6107a061079885610654565b830182610740565b50505b505050565b5f82821c905092915050565b5f6107c35f19846008026107a8565b1980831691505092915050565b5f6107db83836107b4565b9150826002028217905092915050565b6107f482610505565b67ffffffffffffffff81111561080d5761080c610299565b5b6108178254610612565b610822828285610762565b5f60209050601f831160018114610853575f8415610841578287015190505b61084b85826107d0565b8655506108b2565b601f19841661086186610642565b5f5b8281101561088857848901518255600182019150602085019450602081019050610863565b868310156108a557848901516108a1601f8916826107b4565b8355505b6001600288020188555050505b505050505050565b5f6108c482610505565b6108ce818561056d565b93506108de818560208601610519565b6108e781610289565b840191505092915050565b6108fb816106c0565b82525050565b5f6060820190508181035f83015261091981866108ba565b9050818103602083015261092d81856108ba565b905061093c60408301846108f2565b949350505050565b7f5265636f726420646f6573206e6f7420657869737400000000000000000000005f82015250565b5f61097860158361056d565b915061098382610944565b602082019050919050565b5f6020820190508181035f8301526109a58161096c565b905091905056fea2646970667358221220928e7e5cfbd6ed2333aced71cebe0a9fc9c009061281ec64d86777776fcd590964736f6c634300081c0033",
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
