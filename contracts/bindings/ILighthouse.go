// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ILighthousePayload is an auto generated low-level Go binding around an user-defined struct.
type ILighthousePayload struct {
	Price     *big.Int
	Nonce     *big.Int
	TxHashes  [][32]byte
	Bidder    common.Address
	Signature []byte
}

// ILighthouseRevertPayload is an auto generated low-level Go binding around an user-defined struct.
type ILighthouseRevertPayload struct {
	Price     *big.Int
	TxHashes  [][32]byte
	RollupId  string
	AuctionId string
	Reason    uint8
	Bidder    common.Address
	Round     uint8
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AuctionAlreadySettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceStillLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidderAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"InvalidBidderSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRollupOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupInsufficientRevenue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RollupNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ServerInsufficientRevenue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expected\",\"type\":\"address\"}],\"name\":\"UnauthorizedLighthouse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalNotAvailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"auctionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BidSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"BidderReigstered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BidderWithdrawalReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BidderWithdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDeposited\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"auctionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RevertPayloadSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"}],\"name\":\"RollupRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RollupWithdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"server\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ServerRevenueWithdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"auctionId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumILighthouse.SkippedReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"SkippedPayload\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"auctionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"txHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structILighthouse.Payload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"_handleOnePayload\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"auctionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"txHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structILighthouse.Payload[]\",\"name\":\"payloads\",\"type\":\"tuple[]\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"}],\"name\":\"getBidderBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"}],\"name\":\"getRollupRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getServerBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"}],\"name\":\"isWithdrawAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"registerBidder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minBaseFee\",\"type\":\"uint256\"}],\"name\":\"registerRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveBidderWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lighthouse\",\"type\":\"address\"}],\"name\":\"setLighthouse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"auctionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"txHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"auctionId\",\"type\":\"string\"},{\"internalType\":\"enumILighthouse.SkippedReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"round\",\"type\":\"uint8\"}],\"internalType\":\"structILighthouse.RevertPayload\",\"name\":\"payload\",\"type\":\"tuple\"}],\"name\":\"submitRevertPayload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawBidderDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"rollupId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawRollupRevenue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// GetBidderBalance is a free data retrieval call binding the contract method 0x9bb7307c.
//
// Solidity: function getBidderBalance(address bidderAddress) view returns(uint256)
func (_Bindings *BindingsCaller) GetBidderBalance(opts *bind.CallOpts, bidderAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getBidderBalance", bidderAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidderBalance is a free data retrieval call binding the contract method 0x9bb7307c.
//
// Solidity: function getBidderBalance(address bidderAddress) view returns(uint256)
func (_Bindings *BindingsSession) GetBidderBalance(bidderAddress common.Address) (*big.Int, error) {
	return _Bindings.Contract.GetBidderBalance(&_Bindings.CallOpts, bidderAddress)
}

// GetBidderBalance is a free data retrieval call binding the contract method 0x9bb7307c.
//
// Solidity: function getBidderBalance(address bidderAddress) view returns(uint256)
func (_Bindings *BindingsCallerSession) GetBidderBalance(bidderAddress common.Address) (*big.Int, error) {
	return _Bindings.Contract.GetBidderBalance(&_Bindings.CallOpts, bidderAddress)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address bidder) view returns(uint256)
func (_Bindings *BindingsCaller) GetNonce(opts *bind.CallOpts, bidder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getNonce", bidder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address bidder) view returns(uint256)
func (_Bindings *BindingsSession) GetNonce(bidder common.Address) (*big.Int, error) {
	return _Bindings.Contract.GetNonce(&_Bindings.CallOpts, bidder)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address bidder) view returns(uint256)
func (_Bindings *BindingsCallerSession) GetNonce(bidder common.Address) (*big.Int, error) {
	return _Bindings.Contract.GetNonce(&_Bindings.CallOpts, bidder)
}

// GetRollupRevenue is a free data retrieval call binding the contract method 0xde3e1528.
//
// Solidity: function getRollupRevenue(string rollupId) view returns(uint256)
func (_Bindings *BindingsCaller) GetRollupRevenue(opts *bind.CallOpts, rollupId string) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getRollupRevenue", rollupId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRollupRevenue is a free data retrieval call binding the contract method 0xde3e1528.
//
// Solidity: function getRollupRevenue(string rollupId) view returns(uint256)
func (_Bindings *BindingsSession) GetRollupRevenue(rollupId string) (*big.Int, error) {
	return _Bindings.Contract.GetRollupRevenue(&_Bindings.CallOpts, rollupId)
}

// GetRollupRevenue is a free data retrieval call binding the contract method 0xde3e1528.
//
// Solidity: function getRollupRevenue(string rollupId) view returns(uint256)
func (_Bindings *BindingsCallerSession) GetRollupRevenue(rollupId string) (*big.Int, error) {
	return _Bindings.Contract.GetRollupRevenue(&_Bindings.CallOpts, rollupId)
}

// GetServerBalance is a free data retrieval call binding the contract method 0xd072c326.
//
// Solidity: function getServerBalance() view returns(uint256)
func (_Bindings *BindingsCaller) GetServerBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getServerBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetServerBalance is a free data retrieval call binding the contract method 0xd072c326.
//
// Solidity: function getServerBalance() view returns(uint256)
func (_Bindings *BindingsSession) GetServerBalance() (*big.Int, error) {
	return _Bindings.Contract.GetServerBalance(&_Bindings.CallOpts)
}

// GetServerBalance is a free data retrieval call binding the contract method 0xd072c326.
//
// Solidity: function getServerBalance() view returns(uint256)
func (_Bindings *BindingsCallerSession) GetServerBalance() (*big.Int, error) {
	return _Bindings.Contract.GetServerBalance(&_Bindings.CallOpts)
}

// IsWithdrawAvailable is a free data retrieval call binding the contract method 0x5b9522e2.
//
// Solidity: function isWithdrawAvailable(address bidderAddress) view returns(bool)
func (_Bindings *BindingsCaller) IsWithdrawAvailable(opts *bind.CallOpts, bidderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "isWithdrawAvailable", bidderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawAvailable is a free data retrieval call binding the contract method 0x5b9522e2.
//
// Solidity: function isWithdrawAvailable(address bidderAddress) view returns(bool)
func (_Bindings *BindingsSession) IsWithdrawAvailable(bidderAddress common.Address) (bool, error) {
	return _Bindings.Contract.IsWithdrawAvailable(&_Bindings.CallOpts, bidderAddress)
}

// IsWithdrawAvailable is a free data retrieval call binding the contract method 0x5b9522e2.
//
// Solidity: function isWithdrawAvailable(address bidderAddress) view returns(bool)
func (_Bindings *BindingsCallerSession) IsWithdrawAvailable(bidderAddress common.Address) (bool, error) {
	return _Bindings.Contract.IsWithdrawAvailable(&_Bindings.CallOpts, bidderAddress)
}

// HandleOnePayload is a paid mutator transaction binding the contract method 0x1e883bbb.
//
// Solidity: function _handleOnePayload(string rollupId, string auctionId, (uint256,uint256,bytes32[],address,bytes) payload) returns(bool)
func (_Bindings *BindingsTransactor) HandleOnePayload(opts *bind.TransactOpts, rollupId string, auctionId string, payload ILighthousePayload) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "_handleOnePayload", rollupId, auctionId, payload)
}

// HandleOnePayload is a paid mutator transaction binding the contract method 0x1e883bbb.
//
// Solidity: function _handleOnePayload(string rollupId, string auctionId, (uint256,uint256,bytes32[],address,bytes) payload) returns(bool)
func (_Bindings *BindingsSession) HandleOnePayload(rollupId string, auctionId string, payload ILighthousePayload) (*types.Transaction, error) {
	return _Bindings.Contract.HandleOnePayload(&_Bindings.TransactOpts, rollupId, auctionId, payload)
}

// HandleOnePayload is a paid mutator transaction binding the contract method 0x1e883bbb.
//
// Solidity: function _handleOnePayload(string rollupId, string auctionId, (uint256,uint256,bytes32[],address,bytes) payload) returns(bool)
func (_Bindings *BindingsTransactorSession) HandleOnePayload(rollupId string, auctionId string, payload ILighthousePayload) (*types.Transaction, error) {
	return _Bindings.Contract.HandleOnePayload(&_Bindings.TransactOpts, rollupId, auctionId, payload)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x5ebd0789.
//
// Solidity: function closeAuction(string rollupId, string auctionId, (uint256,uint256,bytes32[],address,bytes)[] payloads) returns()
func (_Bindings *BindingsTransactor) CloseAuction(opts *bind.TransactOpts, rollupId string, auctionId string, payloads []ILighthousePayload) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "closeAuction", rollupId, auctionId, payloads)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x5ebd0789.
//
// Solidity: function closeAuction(string rollupId, string auctionId, (uint256,uint256,bytes32[],address,bytes)[] payloads) returns()
func (_Bindings *BindingsSession) CloseAuction(rollupId string, auctionId string, payloads []ILighthousePayload) (*types.Transaction, error) {
	return _Bindings.Contract.CloseAuction(&_Bindings.TransactOpts, rollupId, auctionId, payloads)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x5ebd0789.
//
// Solidity: function closeAuction(string rollupId, string auctionId, (uint256,uint256,bytes32[],address,bytes)[] payloads) returns()
func (_Bindings *BindingsTransactorSession) CloseAuction(rollupId string, auctionId string, payloads []ILighthousePayload) (*types.Transaction, error) {
	return _Bindings.Contract.CloseAuction(&_Bindings.TransactOpts, rollupId, auctionId, payloads)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Bindings *BindingsTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Bindings *BindingsSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Deposit(&_Bindings.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_Bindings *BindingsTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Deposit(&_Bindings.TransactOpts, amount)
}

// RegisterBidder is a paid mutator transaction binding the contract method 0x1435cd2c.
//
// Solidity: function registerBidder(uint256 balance) returns()
func (_Bindings *BindingsTransactor) RegisterBidder(opts *bind.TransactOpts, balance *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerBidder", balance)
}

// RegisterBidder is a paid mutator transaction binding the contract method 0x1435cd2c.
//
// Solidity: function registerBidder(uint256 balance) returns()
func (_Bindings *BindingsSession) RegisterBidder(balance *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterBidder(&_Bindings.TransactOpts, balance)
}

// RegisterBidder is a paid mutator transaction binding the contract method 0x1435cd2c.
//
// Solidity: function registerBidder(uint256 balance) returns()
func (_Bindings *BindingsTransactorSession) RegisterBidder(balance *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterBidder(&_Bindings.TransactOpts, balance)
}

// RegisterRollup is a paid mutator transaction binding the contract method 0x6e10cc01.
//
// Solidity: function registerRollup(string rollupId, uint256 minBaseFee) returns()
func (_Bindings *BindingsTransactor) RegisterRollup(opts *bind.TransactOpts, rollupId string, minBaseFee *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "registerRollup", rollupId, minBaseFee)
}

// RegisterRollup is a paid mutator transaction binding the contract method 0x6e10cc01.
//
// Solidity: function registerRollup(string rollupId, uint256 minBaseFee) returns()
func (_Bindings *BindingsSession) RegisterRollup(rollupId string, minBaseFee *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterRollup(&_Bindings.TransactOpts, rollupId, minBaseFee)
}

// RegisterRollup is a paid mutator transaction binding the contract method 0x6e10cc01.
//
// Solidity: function registerRollup(string rollupId, uint256 minBaseFee) returns()
func (_Bindings *BindingsTransactorSession) RegisterRollup(rollupId string, minBaseFee *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.RegisterRollup(&_Bindings.TransactOpts, rollupId, minBaseFee)
}

// ReserveBidderWithdrawal is a paid mutator transaction binding the contract method 0xa8f55583.
//
// Solidity: function reserveBidderWithdrawal() returns()
func (_Bindings *BindingsTransactor) ReserveBidderWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "reserveBidderWithdrawal")
}

// ReserveBidderWithdrawal is a paid mutator transaction binding the contract method 0xa8f55583.
//
// Solidity: function reserveBidderWithdrawal() returns()
func (_Bindings *BindingsSession) ReserveBidderWithdrawal() (*types.Transaction, error) {
	return _Bindings.Contract.ReserveBidderWithdrawal(&_Bindings.TransactOpts)
}

// ReserveBidderWithdrawal is a paid mutator transaction binding the contract method 0xa8f55583.
//
// Solidity: function reserveBidderWithdrawal() returns()
func (_Bindings *BindingsTransactorSession) ReserveBidderWithdrawal() (*types.Transaction, error) {
	return _Bindings.Contract.ReserveBidderWithdrawal(&_Bindings.TransactOpts)
}

// SetLighthouse is a paid mutator transaction binding the contract method 0x0393658c.
//
// Solidity: function setLighthouse(address lighthouse) returns()
func (_Bindings *BindingsTransactor) SetLighthouse(opts *bind.TransactOpts, lighthouse common.Address) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "setLighthouse", lighthouse)
}

// SetLighthouse is a paid mutator transaction binding the contract method 0x0393658c.
//
// Solidity: function setLighthouse(address lighthouse) returns()
func (_Bindings *BindingsSession) SetLighthouse(lighthouse common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetLighthouse(&_Bindings.TransactOpts, lighthouse)
}

// SetLighthouse is a paid mutator transaction binding the contract method 0x0393658c.
//
// Solidity: function setLighthouse(address lighthouse) returns()
func (_Bindings *BindingsTransactorSession) SetLighthouse(lighthouse common.Address) (*types.Transaction, error) {
	return _Bindings.Contract.SetLighthouse(&_Bindings.TransactOpts, lighthouse)
}

// SubmitRevertPayload is a paid mutator transaction binding the contract method 0x79016e13.
//
// Solidity: function submitRevertPayload(string rollupId, string auctionId, (uint256,bytes32[],string,string,uint8,address,uint8) payload) returns()
func (_Bindings *BindingsTransactor) SubmitRevertPayload(opts *bind.TransactOpts, rollupId string, auctionId string, payload ILighthouseRevertPayload) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "submitRevertPayload", rollupId, auctionId, payload)
}

// SubmitRevertPayload is a paid mutator transaction binding the contract method 0x79016e13.
//
// Solidity: function submitRevertPayload(string rollupId, string auctionId, (uint256,bytes32[],string,string,uint8,address,uint8) payload) returns()
func (_Bindings *BindingsSession) SubmitRevertPayload(rollupId string, auctionId string, payload ILighthouseRevertPayload) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitRevertPayload(&_Bindings.TransactOpts, rollupId, auctionId, payload)
}

// SubmitRevertPayload is a paid mutator transaction binding the contract method 0x79016e13.
//
// Solidity: function submitRevertPayload(string rollupId, string auctionId, (uint256,bytes32[],string,string,uint8,address,uint8) payload) returns()
func (_Bindings *BindingsTransactorSession) SubmitRevertPayload(rollupId string, auctionId string, payload ILighthouseRevertPayload) (*types.Transaction, error) {
	return _Bindings.Contract.SubmitRevertPayload(&_Bindings.TransactOpts, rollupId, auctionId, payload)
}

// WithdrawBidderDeposit is a paid mutator transaction binding the contract method 0x8cb94c8d.
//
// Solidity: function withdrawBidderDeposit() returns()
func (_Bindings *BindingsTransactor) WithdrawBidderDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "withdrawBidderDeposit")
}

// WithdrawBidderDeposit is a paid mutator transaction binding the contract method 0x8cb94c8d.
//
// Solidity: function withdrawBidderDeposit() returns()
func (_Bindings *BindingsSession) WithdrawBidderDeposit() (*types.Transaction, error) {
	return _Bindings.Contract.WithdrawBidderDeposit(&_Bindings.TransactOpts)
}

// WithdrawBidderDeposit is a paid mutator transaction binding the contract method 0x8cb94c8d.
//
// Solidity: function withdrawBidderDeposit() returns()
func (_Bindings *BindingsTransactorSession) WithdrawBidderDeposit() (*types.Transaction, error) {
	return _Bindings.Contract.WithdrawBidderDeposit(&_Bindings.TransactOpts)
}

// WithdrawRollupRevenue is a paid mutator transaction binding the contract method 0xba8c61ca.
//
// Solidity: function withdrawRollupRevenue(string rollupId, uint256 amount) returns()
func (_Bindings *BindingsTransactor) WithdrawRollupRevenue(opts *bind.TransactOpts, rollupId string, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "withdrawRollupRevenue", rollupId, amount)
}

// WithdrawRollupRevenue is a paid mutator transaction binding the contract method 0xba8c61ca.
//
// Solidity: function withdrawRollupRevenue(string rollupId, uint256 amount) returns()
func (_Bindings *BindingsSession) WithdrawRollupRevenue(rollupId string, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.WithdrawRollupRevenue(&_Bindings.TransactOpts, rollupId, amount)
}

// WithdrawRollupRevenue is a paid mutator transaction binding the contract method 0xba8c61ca.
//
// Solidity: function withdrawRollupRevenue(string rollupId, uint256 amount) returns()
func (_Bindings *BindingsTransactorSession) WithdrawRollupRevenue(rollupId string, amount *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.WithdrawRollupRevenue(&_Bindings.TransactOpts, rollupId, amount)
}

// BindingsBidSubmittedIterator is returned from FilterBidSubmitted and is used to iterate over the raw logs and unpacked data for BidSubmitted events raised by the Bindings contract.
type BindingsBidSubmittedIterator struct {
	Event *BindingsBidSubmitted // Event containing the contract specifics and raw log

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
func (it *BindingsBidSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBidSubmitted)
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
		it.Event = new(BindingsBidSubmitted)
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
func (it *BindingsBidSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBidSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBidSubmitted represents a BidSubmitted event raised by the Bindings contract.
type BindingsBidSubmitted struct {
	RollupId  common.Hash
	AuctionId common.Hash
	Buyer     common.Address
	Amount    *big.Int
	Nonce     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidSubmitted is a free log retrieval operation binding the contract event 0x70ada55029fb8e229b0ba6acc04690895590ae6a9035c6bcbf0449b7d471179e.
//
// Solidity: event BidSubmitted(string indexed rollupId, string indexed auctionId, address buyer, uint256 amount, uint256 nonce)
func (_Bindings *BindingsFilterer) FilterBidSubmitted(opts *bind.FilterOpts, rollupId []string, auctionId []string) (*BindingsBidSubmittedIterator, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BidSubmitted", rollupIdRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBidSubmittedIterator{contract: _Bindings.contract, event: "BidSubmitted", logs: logs, sub: sub}, nil
}

// WatchBidSubmitted is a free log subscription operation binding the contract event 0x70ada55029fb8e229b0ba6acc04690895590ae6a9035c6bcbf0449b7d471179e.
//
// Solidity: event BidSubmitted(string indexed rollupId, string indexed auctionId, address buyer, uint256 amount, uint256 nonce)
func (_Bindings *BindingsFilterer) WatchBidSubmitted(opts *bind.WatchOpts, sink chan<- *BindingsBidSubmitted, rollupId []string, auctionId []string) (event.Subscription, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BidSubmitted", rollupIdRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBidSubmitted)
				if err := _Bindings.contract.UnpackLog(event, "BidSubmitted", log); err != nil {
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

// ParseBidSubmitted is a log parse operation binding the contract event 0x70ada55029fb8e229b0ba6acc04690895590ae6a9035c6bcbf0449b7d471179e.
//
// Solidity: event BidSubmitted(string indexed rollupId, string indexed auctionId, address buyer, uint256 amount, uint256 nonce)
func (_Bindings *BindingsFilterer) ParseBidSubmitted(log types.Log) (*BindingsBidSubmitted, error) {
	event := new(BindingsBidSubmitted)
	if err := _Bindings.contract.UnpackLog(event, "BidSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsBidderReigsteredIterator is returned from FilterBidderReigstered and is used to iterate over the raw logs and unpacked data for BidderReigstered events raised by the Bindings contract.
type BindingsBidderReigsteredIterator struct {
	Event *BindingsBidderReigstered // Event containing the contract specifics and raw log

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
func (it *BindingsBidderReigsteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBidderReigstered)
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
		it.Event = new(BindingsBidderReigstered)
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
func (it *BindingsBidderReigsteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBidderReigsteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBidderReigstered represents a BidderReigstered event raised by the Bindings contract.
type BindingsBidderReigstered struct {
	BidderAddress common.Address
	Balance       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBidderReigstered is a free log retrieval operation binding the contract event 0x24d69a2d631fb085e314b9db7b8ea6bd7f18c9f6f6d68f652ff58ae80a9dbbdf.
//
// Solidity: event BidderReigstered(address indexed bidderAddress, uint256 balance)
func (_Bindings *BindingsFilterer) FilterBidderReigstered(opts *bind.FilterOpts, bidderAddress []common.Address) (*BindingsBidderReigsteredIterator, error) {

	var bidderAddressRule []interface{}
	for _, bidderAddressItem := range bidderAddress {
		bidderAddressRule = append(bidderAddressRule, bidderAddressItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BidderReigstered", bidderAddressRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBidderReigsteredIterator{contract: _Bindings.contract, event: "BidderReigstered", logs: logs, sub: sub}, nil
}

// WatchBidderReigstered is a free log subscription operation binding the contract event 0x24d69a2d631fb085e314b9db7b8ea6bd7f18c9f6f6d68f652ff58ae80a9dbbdf.
//
// Solidity: event BidderReigstered(address indexed bidderAddress, uint256 balance)
func (_Bindings *BindingsFilterer) WatchBidderReigstered(opts *bind.WatchOpts, sink chan<- *BindingsBidderReigstered, bidderAddress []common.Address) (event.Subscription, error) {

	var bidderAddressRule []interface{}
	for _, bidderAddressItem := range bidderAddress {
		bidderAddressRule = append(bidderAddressRule, bidderAddressItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BidderReigstered", bidderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBidderReigstered)
				if err := _Bindings.contract.UnpackLog(event, "BidderReigstered", log); err != nil {
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

// ParseBidderReigstered is a log parse operation binding the contract event 0x24d69a2d631fb085e314b9db7b8ea6bd7f18c9f6f6d68f652ff58ae80a9dbbdf.
//
// Solidity: event BidderReigstered(address indexed bidderAddress, uint256 balance)
func (_Bindings *BindingsFilterer) ParseBidderReigstered(log types.Log) (*BindingsBidderReigstered, error) {
	event := new(BindingsBidderReigstered)
	if err := _Bindings.contract.UnpackLog(event, "BidderReigstered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsBidderWithdrawalReservedIterator is returned from FilterBidderWithdrawalReserved and is used to iterate over the raw logs and unpacked data for BidderWithdrawalReserved events raised by the Bindings contract.
type BindingsBidderWithdrawalReservedIterator struct {
	Event *BindingsBidderWithdrawalReserved // Event containing the contract specifics and raw log

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
func (it *BindingsBidderWithdrawalReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBidderWithdrawalReserved)
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
		it.Event = new(BindingsBidderWithdrawalReserved)
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
func (it *BindingsBidderWithdrawalReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBidderWithdrawalReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBidderWithdrawalReserved represents a BidderWithdrawalReserved event raised by the Bindings contract.
type BindingsBidderWithdrawalReserved struct {
	BidderAddress common.Address
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBidderWithdrawalReserved is a free log retrieval operation binding the contract event 0x954cebc1763790b4bcb63aa8502de6d2c509f4b9af114d21d904e2bbb889d9d8.
//
// Solidity: event BidderWithdrawalReserved(address indexed bidderAddress, uint256 timestamp)
func (_Bindings *BindingsFilterer) FilterBidderWithdrawalReserved(opts *bind.FilterOpts, bidderAddress []common.Address) (*BindingsBidderWithdrawalReservedIterator, error) {

	var bidderAddressRule []interface{}
	for _, bidderAddressItem := range bidderAddress {
		bidderAddressRule = append(bidderAddressRule, bidderAddressItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BidderWithdrawalReserved", bidderAddressRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBidderWithdrawalReservedIterator{contract: _Bindings.contract, event: "BidderWithdrawalReserved", logs: logs, sub: sub}, nil
}

// WatchBidderWithdrawalReserved is a free log subscription operation binding the contract event 0x954cebc1763790b4bcb63aa8502de6d2c509f4b9af114d21d904e2bbb889d9d8.
//
// Solidity: event BidderWithdrawalReserved(address indexed bidderAddress, uint256 timestamp)
func (_Bindings *BindingsFilterer) WatchBidderWithdrawalReserved(opts *bind.WatchOpts, sink chan<- *BindingsBidderWithdrawalReserved, bidderAddress []common.Address) (event.Subscription, error) {

	var bidderAddressRule []interface{}
	for _, bidderAddressItem := range bidderAddress {
		bidderAddressRule = append(bidderAddressRule, bidderAddressItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BidderWithdrawalReserved", bidderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBidderWithdrawalReserved)
				if err := _Bindings.contract.UnpackLog(event, "BidderWithdrawalReserved", log); err != nil {
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

// ParseBidderWithdrawalReserved is a log parse operation binding the contract event 0x954cebc1763790b4bcb63aa8502de6d2c509f4b9af114d21d904e2bbb889d9d8.
//
// Solidity: event BidderWithdrawalReserved(address indexed bidderAddress, uint256 timestamp)
func (_Bindings *BindingsFilterer) ParseBidderWithdrawalReserved(log types.Log) (*BindingsBidderWithdrawalReserved, error) {
	event := new(BindingsBidderWithdrawalReserved)
	if err := _Bindings.contract.UnpackLog(event, "BidderWithdrawalReserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsBidderWithdrewIterator is returned from FilterBidderWithdrew and is used to iterate over the raw logs and unpacked data for BidderWithdrew events raised by the Bindings contract.
type BindingsBidderWithdrewIterator struct {
	Event *BindingsBidderWithdrew // Event containing the contract specifics and raw log

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
func (it *BindingsBidderWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsBidderWithdrew)
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
		it.Event = new(BindingsBidderWithdrew)
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
func (it *BindingsBidderWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsBidderWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsBidderWithdrew represents a BidderWithdrew event raised by the Bindings contract.
type BindingsBidderWithdrew struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidderWithdrew is a free log retrieval operation binding the contract event 0x9d06c1dd15c31c6eda24b3385ad1dc60a402059b3d2d799675a4a2af5b29af2e.
//
// Solidity: event BidderWithdrew(address indexed bidder, uint256 amount)
func (_Bindings *BindingsFilterer) FilterBidderWithdrew(opts *bind.FilterOpts, bidder []common.Address) (*BindingsBidderWithdrewIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "BidderWithdrew", bidderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsBidderWithdrewIterator{contract: _Bindings.contract, event: "BidderWithdrew", logs: logs, sub: sub}, nil
}

// WatchBidderWithdrew is a free log subscription operation binding the contract event 0x9d06c1dd15c31c6eda24b3385ad1dc60a402059b3d2d799675a4a2af5b29af2e.
//
// Solidity: event BidderWithdrew(address indexed bidder, uint256 amount)
func (_Bindings *BindingsFilterer) WatchBidderWithdrew(opts *bind.WatchOpts, sink chan<- *BindingsBidderWithdrew, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "BidderWithdrew", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsBidderWithdrew)
				if err := _Bindings.contract.UnpackLog(event, "BidderWithdrew", log); err != nil {
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

// ParseBidderWithdrew is a log parse operation binding the contract event 0x9d06c1dd15c31c6eda24b3385ad1dc60a402059b3d2d799675a4a2af5b29af2e.
//
// Solidity: event BidderWithdrew(address indexed bidder, uint256 amount)
func (_Bindings *BindingsFilterer) ParseBidderWithdrew(log types.Log) (*BindingsBidderWithdrew, error) {
	event := new(BindingsBidderWithdrew)
	if err := _Bindings.contract.UnpackLog(event, "BidderWithdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Bindings contract.
type BindingsDepositedIterator struct {
	Event *BindingsDeposited // Event containing the contract specifics and raw log

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
func (it *BindingsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsDeposited)
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
		it.Event = new(BindingsDeposited)
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
func (it *BindingsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsDeposited represents a Deposited event raised by the Bindings contract.
type BindingsDeposited struct {
	Bidder         common.Address
	Amount         *big.Int
	TotalDeposited *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed bidder, uint256 amount, uint256 totalDeposited)
func (_Bindings *BindingsFilterer) FilterDeposited(opts *bind.FilterOpts, bidder []common.Address) (*BindingsDepositedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "Deposited", bidderRule)
	if err != nil {
		return nil, err
	}
	return &BindingsDepositedIterator{contract: _Bindings.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed bidder, uint256 amount, uint256 totalDeposited)
func (_Bindings *BindingsFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *BindingsDeposited, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "Deposited", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsDeposited)
				if err := _Bindings.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x73a19dd210f1a7f902193214c0ee91dd35ee5b4d920cba8d519eca65a7b488ca.
//
// Solidity: event Deposited(address indexed bidder, uint256 amount, uint256 totalDeposited)
func (_Bindings *BindingsFilterer) ParseDeposited(log types.Log) (*BindingsDeposited, error) {
	event := new(BindingsDeposited)
	if err := _Bindings.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRevertPayloadSubmittedIterator is returned from FilterRevertPayloadSubmitted and is used to iterate over the raw logs and unpacked data for RevertPayloadSubmitted events raised by the Bindings contract.
type BindingsRevertPayloadSubmittedIterator struct {
	Event *BindingsRevertPayloadSubmitted // Event containing the contract specifics and raw log

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
func (it *BindingsRevertPayloadSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRevertPayloadSubmitted)
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
		it.Event = new(BindingsRevertPayloadSubmitted)
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
func (it *BindingsRevertPayloadSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRevertPayloadSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRevertPayloadSubmitted represents a RevertPayloadSubmitted event raised by the Bindings contract.
type BindingsRevertPayloadSubmitted struct {
	RollupId  common.Hash
	AuctionId common.Hash
	Buyer     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevertPayloadSubmitted is a free log retrieval operation binding the contract event 0x71b7b8510a2397a2174898abc0eeddbf718fa0daf69856fef026423181743766.
//
// Solidity: event RevertPayloadSubmitted(string indexed rollupId, string indexed auctionId, address buyer, uint256 amount)
func (_Bindings *BindingsFilterer) FilterRevertPayloadSubmitted(opts *bind.FilterOpts, rollupId []string, auctionId []string) (*BindingsRevertPayloadSubmittedIterator, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RevertPayloadSubmitted", rollupIdRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRevertPayloadSubmittedIterator{contract: _Bindings.contract, event: "RevertPayloadSubmitted", logs: logs, sub: sub}, nil
}

// WatchRevertPayloadSubmitted is a free log subscription operation binding the contract event 0x71b7b8510a2397a2174898abc0eeddbf718fa0daf69856fef026423181743766.
//
// Solidity: event RevertPayloadSubmitted(string indexed rollupId, string indexed auctionId, address buyer, uint256 amount)
func (_Bindings *BindingsFilterer) WatchRevertPayloadSubmitted(opts *bind.WatchOpts, sink chan<- *BindingsRevertPayloadSubmitted, rollupId []string, auctionId []string) (event.Subscription, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RevertPayloadSubmitted", rollupIdRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRevertPayloadSubmitted)
				if err := _Bindings.contract.UnpackLog(event, "RevertPayloadSubmitted", log); err != nil {
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

// ParseRevertPayloadSubmitted is a log parse operation binding the contract event 0x71b7b8510a2397a2174898abc0eeddbf718fa0daf69856fef026423181743766.
//
// Solidity: event RevertPayloadSubmitted(string indexed rollupId, string indexed auctionId, address buyer, uint256 amount)
func (_Bindings *BindingsFilterer) ParseRevertPayloadSubmitted(log types.Log) (*BindingsRevertPayloadSubmitted, error) {
	event := new(BindingsRevertPayloadSubmitted)
	if err := _Bindings.contract.UnpackLog(event, "RevertPayloadSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRollupRegisteredIterator is returned from FilterRollupRegistered and is used to iterate over the raw logs and unpacked data for RollupRegistered events raised by the Bindings contract.
type BindingsRollupRegisteredIterator struct {
	Event *BindingsRollupRegistered // Event containing the contract specifics and raw log

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
func (it *BindingsRollupRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRollupRegistered)
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
		it.Event = new(BindingsRollupRegistered)
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
func (it *BindingsRollupRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRollupRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRollupRegistered represents a RollupRegistered event raised by the Bindings contract.
type BindingsRollupRegistered struct {
	RollupAddress common.Address
	RollupId      string
	MinBaseFee    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollupRegistered is a free log retrieval operation binding the contract event 0xdde4dc2aa3cea284e6fe54b8de3c1387843d5099282bf41d895ecf713ac589cb.
//
// Solidity: event RollupRegistered(address indexed rollupAddress, string rollupId, uint256 minBaseFee)
func (_Bindings *BindingsFilterer) FilterRollupRegistered(opts *bind.FilterOpts, rollupAddress []common.Address) (*BindingsRollupRegisteredIterator, error) {

	var rollupAddressRule []interface{}
	for _, rollupAddressItem := range rollupAddress {
		rollupAddressRule = append(rollupAddressRule, rollupAddressItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RollupRegistered", rollupAddressRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRollupRegisteredIterator{contract: _Bindings.contract, event: "RollupRegistered", logs: logs, sub: sub}, nil
}

// WatchRollupRegistered is a free log subscription operation binding the contract event 0xdde4dc2aa3cea284e6fe54b8de3c1387843d5099282bf41d895ecf713ac589cb.
//
// Solidity: event RollupRegistered(address indexed rollupAddress, string rollupId, uint256 minBaseFee)
func (_Bindings *BindingsFilterer) WatchRollupRegistered(opts *bind.WatchOpts, sink chan<- *BindingsRollupRegistered, rollupAddress []common.Address) (event.Subscription, error) {

	var rollupAddressRule []interface{}
	for _, rollupAddressItem := range rollupAddress {
		rollupAddressRule = append(rollupAddressRule, rollupAddressItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RollupRegistered", rollupAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRollupRegistered)
				if err := _Bindings.contract.UnpackLog(event, "RollupRegistered", log); err != nil {
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

// ParseRollupRegistered is a log parse operation binding the contract event 0xdde4dc2aa3cea284e6fe54b8de3c1387843d5099282bf41d895ecf713ac589cb.
//
// Solidity: event RollupRegistered(address indexed rollupAddress, string rollupId, uint256 minBaseFee)
func (_Bindings *BindingsFilterer) ParseRollupRegistered(log types.Log) (*BindingsRollupRegistered, error) {
	event := new(BindingsRollupRegistered)
	if err := _Bindings.contract.UnpackLog(event, "RollupRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsRollupWithdrewIterator is returned from FilterRollupWithdrew and is used to iterate over the raw logs and unpacked data for RollupWithdrew events raised by the Bindings contract.
type BindingsRollupWithdrewIterator struct {
	Event *BindingsRollupWithdrew // Event containing the contract specifics and raw log

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
func (it *BindingsRollupWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsRollupWithdrew)
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
		it.Event = new(BindingsRollupWithdrew)
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
func (it *BindingsRollupWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsRollupWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsRollupWithdrew represents a RollupWithdrew event raised by the Bindings contract.
type BindingsRollupWithdrew struct {
	RollupId common.Hash
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollupWithdrew is a free log retrieval operation binding the contract event 0x51791aabfd71634ef8d41d86d9ce54bc3a871c08d12c786e2f2ed6eff96fa3fe.
//
// Solidity: event RollupWithdrew(string indexed rollupId, uint256 amount)
func (_Bindings *BindingsFilterer) FilterRollupWithdrew(opts *bind.FilterOpts, rollupId []string) (*BindingsRollupWithdrewIterator, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "RollupWithdrew", rollupIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsRollupWithdrewIterator{contract: _Bindings.contract, event: "RollupWithdrew", logs: logs, sub: sub}, nil
}

// WatchRollupWithdrew is a free log subscription operation binding the contract event 0x51791aabfd71634ef8d41d86d9ce54bc3a871c08d12c786e2f2ed6eff96fa3fe.
//
// Solidity: event RollupWithdrew(string indexed rollupId, uint256 amount)
func (_Bindings *BindingsFilterer) WatchRollupWithdrew(opts *bind.WatchOpts, sink chan<- *BindingsRollupWithdrew, rollupId []string) (event.Subscription, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "RollupWithdrew", rollupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsRollupWithdrew)
				if err := _Bindings.contract.UnpackLog(event, "RollupWithdrew", log); err != nil {
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

// ParseRollupWithdrew is a log parse operation binding the contract event 0x51791aabfd71634ef8d41d86d9ce54bc3a871c08d12c786e2f2ed6eff96fa3fe.
//
// Solidity: event RollupWithdrew(string indexed rollupId, uint256 amount)
func (_Bindings *BindingsFilterer) ParseRollupWithdrew(log types.Log) (*BindingsRollupWithdrew, error) {
	event := new(BindingsRollupWithdrew)
	if err := _Bindings.contract.UnpackLog(event, "RollupWithdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsServerRevenueWithdrewIterator is returned from FilterServerRevenueWithdrew and is used to iterate over the raw logs and unpacked data for ServerRevenueWithdrew events raised by the Bindings contract.
type BindingsServerRevenueWithdrewIterator struct {
	Event *BindingsServerRevenueWithdrew // Event containing the contract specifics and raw log

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
func (it *BindingsServerRevenueWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsServerRevenueWithdrew)
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
		it.Event = new(BindingsServerRevenueWithdrew)
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
func (it *BindingsServerRevenueWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsServerRevenueWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsServerRevenueWithdrew represents a ServerRevenueWithdrew event raised by the Bindings contract.
type BindingsServerRevenueWithdrew struct {
	Server common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterServerRevenueWithdrew is a free log retrieval operation binding the contract event 0xda0800af85a06097ca48951d86feba8789603ed7b36cc0693ead2b3b58e6bfb6.
//
// Solidity: event ServerRevenueWithdrew(address indexed server, uint256 amount)
func (_Bindings *BindingsFilterer) FilterServerRevenueWithdrew(opts *bind.FilterOpts, server []common.Address) (*BindingsServerRevenueWithdrewIterator, error) {

	var serverRule []interface{}
	for _, serverItem := range server {
		serverRule = append(serverRule, serverItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "ServerRevenueWithdrew", serverRule)
	if err != nil {
		return nil, err
	}
	return &BindingsServerRevenueWithdrewIterator{contract: _Bindings.contract, event: "ServerRevenueWithdrew", logs: logs, sub: sub}, nil
}

// WatchServerRevenueWithdrew is a free log subscription operation binding the contract event 0xda0800af85a06097ca48951d86feba8789603ed7b36cc0693ead2b3b58e6bfb6.
//
// Solidity: event ServerRevenueWithdrew(address indexed server, uint256 amount)
func (_Bindings *BindingsFilterer) WatchServerRevenueWithdrew(opts *bind.WatchOpts, sink chan<- *BindingsServerRevenueWithdrew, server []common.Address) (event.Subscription, error) {

	var serverRule []interface{}
	for _, serverItem := range server {
		serverRule = append(serverRule, serverItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "ServerRevenueWithdrew", serverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsServerRevenueWithdrew)
				if err := _Bindings.contract.UnpackLog(event, "ServerRevenueWithdrew", log); err != nil {
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

// ParseServerRevenueWithdrew is a log parse operation binding the contract event 0xda0800af85a06097ca48951d86feba8789603ed7b36cc0693ead2b3b58e6bfb6.
//
// Solidity: event ServerRevenueWithdrew(address indexed server, uint256 amount)
func (_Bindings *BindingsFilterer) ParseServerRevenueWithdrew(log types.Log) (*BindingsServerRevenueWithdrew, error) {
	event := new(BindingsServerRevenueWithdrew)
	if err := _Bindings.contract.UnpackLog(event, "ServerRevenueWithdrew", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingsSkippedPayloadIterator is returned from FilterSkippedPayload and is used to iterate over the raw logs and unpacked data for SkippedPayload events raised by the Bindings contract.
type BindingsSkippedPayloadIterator struct {
	Event *BindingsSkippedPayload // Event containing the contract specifics and raw log

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
func (it *BindingsSkippedPayloadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsSkippedPayload)
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
		it.Event = new(BindingsSkippedPayload)
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
func (it *BindingsSkippedPayloadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsSkippedPayloadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsSkippedPayload represents a SkippedPayload event raised by the Bindings contract.
type BindingsSkippedPayload struct {
	Bidder    common.Address
	RollupId  common.Hash
	AuctionId string
	Reason    uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSkippedPayload is a free log retrieval operation binding the contract event 0x84f0c0521613bf5a01a52deb10058d6647bcfe2e4f62547b4d90ec2f47ccfe67.
//
// Solidity: event SkippedPayload(address indexed bidder, string indexed rollupId, string auctionId, uint8 reason)
func (_Bindings *BindingsFilterer) FilterSkippedPayload(opts *bind.FilterOpts, bidder []common.Address, rollupId []string) (*BindingsSkippedPayloadIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "SkippedPayload", bidderRule, rollupIdRule)
	if err != nil {
		return nil, err
	}
	return &BindingsSkippedPayloadIterator{contract: _Bindings.contract, event: "SkippedPayload", logs: logs, sub: sub}, nil
}

// WatchSkippedPayload is a free log subscription operation binding the contract event 0x84f0c0521613bf5a01a52deb10058d6647bcfe2e4f62547b4d90ec2f47ccfe67.
//
// Solidity: event SkippedPayload(address indexed bidder, string indexed rollupId, string auctionId, uint8 reason)
func (_Bindings *BindingsFilterer) WatchSkippedPayload(opts *bind.WatchOpts, sink chan<- *BindingsSkippedPayload, bidder []common.Address, rollupId []string) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "SkippedPayload", bidderRule, rollupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsSkippedPayload)
				if err := _Bindings.contract.UnpackLog(event, "SkippedPayload", log); err != nil {
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

// ParseSkippedPayload is a log parse operation binding the contract event 0x84f0c0521613bf5a01a52deb10058d6647bcfe2e4f62547b4d90ec2f47ccfe67.
//
// Solidity: event SkippedPayload(address indexed bidder, string indexed rollupId, string auctionId, uint8 reason)
func (_Bindings *BindingsFilterer) ParseSkippedPayload(log types.Log) (*BindingsSkippedPayload, error) {
	event := new(BindingsSkippedPayload)
	if err := _Bindings.contract.UnpackLog(event, "SkippedPayload", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
