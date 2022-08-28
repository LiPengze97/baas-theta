// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accessors

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/thetatoken/thetasubchain/eth"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/thetasubchain/eth/core/types"
	"github.com/thetatoken/thetasubchain/eth/event"
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
)

// TokenBankMetaData contains all meta data concerning the TokenBank contract.
var TokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mainchainID_\",\"type\":\"uint256\"},{\"internalType\":\"contractChainRegistrar\",\"name\":\"chainRegistrar_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"mainchainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenLockVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenUnlockNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherBurnNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherBurnVotingRecords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumlatedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherMintNonceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedTokenLockNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"}],\"name\":\"getTokenLockEventHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxProcessedVoucherBurnNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eventNonce\",\"type\":\"uint256\"}],\"name\":\"getVoucherBurnEventHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161037438038061037483398101604081905261002f9161005e565b6001600081905591909155600280546001600160a01b0319166001600160a01b0390921691909117905561009b565b6000806040838503121561007157600080fd5b825160208401519092506001600160a01b038116811461009057600080fd5b809150509250929050565b6102ca806100aa6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638883931e116100715780638883931e1461017c578063ca2075691461019c578063ccf187c7146101bc578063dd17eb6d146101dc578063f95627ac14610207578063feaff0521461022757600080fd5b8063073b9502146100ae5780631eb78737146100ca5780636ac739b914610111578063740cb7f81461013c578063766f8fb01461015c575b600080fd5b6100b760015481565b6040519081526020015b60405180910390f35b6100fc6100d8366004610272565b600b6020908152600092835260408084209091529082529020805460029091015482565b604080519283526020830191909152016100c1565b6100b761011f366004610272565b600091825260086020908152604080842092845291905290205490565b6100b761014a366004610259565b60066020526000908152604090205481565b6100b761016a366004610259565b6000908152600a602052604090205490565b6100b761018a366004610259565b60036020526000908152604090205481565b6100b76101aa366004610259565b60056020526000908152604090205481565b6100b76101ca366004610259565b60046020526000908152604090205481565b6100b76101ea366004610272565b600091825260076020908152604080842092845291905290205490565b6100b7610215366004610259565b60009081526009602052604090205490565b6100fc610235366004610272565b600c6020908152600092835260408084209091529082529020805460029091015482565b60006020828403121561026b57600080fd5b5035919050565b6000806040838503121561028557600080fd5b5050803592602090910135915056fea26469706673582212207440b67d63b36a880b6acd4a4fbfdc981472fe35ba5afe35802db15d74e4a4b064736f6c63430008070033",
}

// TokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenBankMetaData.ABI instead.
var TokenBankABI = TokenBankMetaData.ABI

// TokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenBankMetaData.Bin instead.
var TokenBankBin = TokenBankMetaData.Bin

// DeployTokenBank deploys a new Ethereum contract, binding an instance of TokenBank to it.
func DeployTokenBank(auth *bind.TransactOpts, backend bind.ContractBackend, mainchainID_ *big.Int, chainRegistrar_ common.Address) (common.Address, *types.Transaction, *TokenBank, error) {
	parsed, err := TokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenBankBin), backend, mainchainID_, chainRegistrar_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenBank{TokenBankCaller: TokenBankCaller{contract: contract}, TokenBankTransactor: TokenBankTransactor{contract: contract}, TokenBankFilterer: TokenBankFilterer{contract: contract}}, nil
}

// TokenBank is an auto generated Go binding around an Ethereum contract.
type TokenBank struct {
	TokenBankCaller     // Read-only binding to the contract
	TokenBankTransactor // Write-only binding to the contract
	TokenBankFilterer   // Log filterer for contract events
}

// TokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenBankSession struct {
	Contract     *TokenBank        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenBankCallerSession struct {
	Contract *TokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenBankTransactorSession struct {
	Contract     *TokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenBankRaw struct {
	Contract *TokenBank // Generic contract binding to access the raw methods on
}

// TokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenBankCallerRaw struct {
	Contract *TokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// TokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenBankTransactorRaw struct {
	Contract *TokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenBank creates a new instance of TokenBank, bound to a specific deployed contract.
func NewTokenBank(address common.Address, backend bind.ContractBackend) (*TokenBank, error) {
	contract, err := bindTokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenBank{TokenBankCaller: TokenBankCaller{contract: contract}, TokenBankTransactor: TokenBankTransactor{contract: contract}, TokenBankFilterer: TokenBankFilterer{contract: contract}}, nil
}

// NewTokenBankCaller creates a new read-only instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankCaller(address common.Address, caller bind.ContractCaller) (*TokenBankCaller, error) {
	contract, err := bindTokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBankCaller{contract: contract}, nil
}

// NewTokenBankTransactor creates a new write-only instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenBankTransactor, error) {
	contract, err := bindTokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBankTransactor{contract: contract}, nil
}

// NewTokenBankFilterer creates a new log filterer instance of TokenBank, bound to a specific deployed contract.
func NewTokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenBankFilterer, error) {
	contract, err := bindTokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenBankFilterer{contract: contract}, nil
}

// bindTokenBank binds a generic wrapper to an already deployed contract.
func bindTokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBank *TokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBank.Contract.TokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBank *TokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBank.Contract.TokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBank *TokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBank.Contract.TokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBank *TokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBank *TokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBank *TokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBank.Contract.contract.Transact(opts, method, params...)
}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankCaller) GetMaxProcessedTokenLockNonce(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getMaxProcessedTokenLockNonce", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankSession) GetMaxProcessedTokenLockNonce(chainID *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetMaxProcessedTokenLockNonce(&_TokenBank.CallOpts, chainID)
}

// GetMaxProcessedTokenLockNonce is a free data retrieval call binding the contract method 0xf95627ac.
//
// Solidity: function getMaxProcessedTokenLockNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) GetMaxProcessedTokenLockNonce(chainID *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetMaxProcessedTokenLockNonce(&_TokenBank.CallOpts, chainID)
}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankCaller) GetMaxProcessedVoucherBurnNonce(opts *bind.CallOpts, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getMaxProcessedVoucherBurnNonce", chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankSession) GetMaxProcessedVoucherBurnNonce(chainID *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetMaxProcessedVoucherBurnNonce(&_TokenBank.CallOpts, chainID)
}

// GetMaxProcessedVoucherBurnNonce is a free data retrieval call binding the contract method 0x766f8fb0.
//
// Solidity: function getMaxProcessedVoucherBurnNonce(uint256 chainID) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) GetMaxProcessedVoucherBurnNonce(chainID *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetMaxProcessedVoucherBurnNonce(&_TokenBank.CallOpts, chainID)
}

// GetTokenLockEventHeight is a free data retrieval call binding the contract method 0xdd17eb6d.
//
// Solidity: function getTokenLockEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankCaller) GetTokenLockEventHeight(opts *bind.CallOpts, chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getTokenLockEventHeight", chainID, eventNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenLockEventHeight is a free data retrieval call binding the contract method 0xdd17eb6d.
//
// Solidity: function getTokenLockEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankSession) GetTokenLockEventHeight(chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetTokenLockEventHeight(&_TokenBank.CallOpts, chainID, eventNonce)
}

// GetTokenLockEventHeight is a free data retrieval call binding the contract method 0xdd17eb6d.
//
// Solidity: function getTokenLockEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) GetTokenLockEventHeight(chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetTokenLockEventHeight(&_TokenBank.CallOpts, chainID, eventNonce)
}

// GetVoucherBurnEventHeight is a free data retrieval call binding the contract method 0x6ac739b9.
//
// Solidity: function getVoucherBurnEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankCaller) GetVoucherBurnEventHeight(opts *bind.CallOpts, chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "getVoucherBurnEventHeight", chainID, eventNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoucherBurnEventHeight is a free data retrieval call binding the contract method 0x6ac739b9.
//
// Solidity: function getVoucherBurnEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankSession) GetVoucherBurnEventHeight(chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetVoucherBurnEventHeight(&_TokenBank.CallOpts, chainID, eventNonce)
}

// GetVoucherBurnEventHeight is a free data retrieval call binding the contract method 0x6ac739b9.
//
// Solidity: function getVoucherBurnEventHeight(uint256 chainID, uint256 eventNonce) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) GetVoucherBurnEventHeight(chainID *big.Int, eventNonce *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.GetVoucherBurnEventHeight(&_TokenBank.CallOpts, chainID, eventNonce)
}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TokenBank *TokenBankCaller) MainchainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "mainchainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TokenBank *TokenBankSession) MainchainID() (*big.Int, error) {
	return _TokenBank.Contract.MainchainID(&_TokenBank.CallOpts)
}

// MainchainID is a free data retrieval call binding the contract method 0x073b9502.
//
// Solidity: function mainchainID() view returns(uint256)
func (_TokenBank *TokenBankCallerSession) MainchainID() (*big.Int, error) {
	return _TokenBank.Contract.MainchainID(&_TokenBank.CallOpts)
}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCaller) TokenLockNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "tokenLockNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankSession) TokenLockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.TokenLockNonceMap(&_TokenBank.CallOpts, arg0)
}

// TokenLockNonceMap is a free data retrieval call binding the contract method 0x8883931e.
//
// Solidity: function tokenLockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) TokenLockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.TokenLockNonceMap(&_TokenBank.CallOpts, arg0)
}

// TokenLockVotingRecords is a free data retrieval call binding the contract method 0x1eb78737.
//
// Solidity: function tokenLockVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankCaller) TokenLockVotingRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "tokenLockVotingRecords", arg0, arg1)

	outstruct := new(struct {
		Dynasty          *big.Int
		AccumlatedShares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dynasty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccumlatedShares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenLockVotingRecords is a free data retrieval call binding the contract method 0x1eb78737.
//
// Solidity: function tokenLockVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankSession) TokenLockVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TokenBank.Contract.TokenLockVotingRecords(&_TokenBank.CallOpts, arg0, arg1)
}

// TokenLockVotingRecords is a free data retrieval call binding the contract method 0x1eb78737.
//
// Solidity: function tokenLockVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankCallerSession) TokenLockVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TokenBank.Contract.TokenLockVotingRecords(&_TokenBank.CallOpts, arg0, arg1)
}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCaller) TokenUnlockNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "tokenUnlockNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankSession) TokenUnlockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.TokenUnlockNonceMap(&_TokenBank.CallOpts, arg0)
}

// TokenUnlockNonceMap is a free data retrieval call binding the contract method 0xccf187c7.
//
// Solidity: function tokenUnlockNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) TokenUnlockNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.TokenUnlockNonceMap(&_TokenBank.CallOpts, arg0)
}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCaller) VoucherBurnNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "voucherBurnNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankSession) VoucherBurnNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.VoucherBurnNonceMap(&_TokenBank.CallOpts, arg0)
}

// VoucherBurnNonceMap is a free data retrieval call binding the contract method 0xca207569.
//
// Solidity: function voucherBurnNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) VoucherBurnNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.VoucherBurnNonceMap(&_TokenBank.CallOpts, arg0)
}

// VoucherBurnVotingRecords is a free data retrieval call binding the contract method 0xfeaff052.
//
// Solidity: function voucherBurnVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankCaller) VoucherBurnVotingRecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "voucherBurnVotingRecords", arg0, arg1)

	outstruct := new(struct {
		Dynasty          *big.Int
		AccumlatedShares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dynasty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccumlatedShares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VoucherBurnVotingRecords is a free data retrieval call binding the contract method 0xfeaff052.
//
// Solidity: function voucherBurnVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankSession) VoucherBurnVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TokenBank.Contract.VoucherBurnVotingRecords(&_TokenBank.CallOpts, arg0, arg1)
}

// VoucherBurnVotingRecords is a free data retrieval call binding the contract method 0xfeaff052.
//
// Solidity: function voucherBurnVotingRecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedShares)
func (_TokenBank *TokenBankCallerSession) VoucherBurnVotingRecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	AccumlatedShares *big.Int
}, error) {
	return _TokenBank.Contract.VoucherBurnVotingRecords(&_TokenBank.CallOpts, arg0, arg1)
}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCaller) VoucherMintNonceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenBank.contract.Call(opts, &out, "voucherMintNonceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankSession) VoucherMintNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.VoucherMintNonceMap(&_TokenBank.CallOpts, arg0)
}

// VoucherMintNonceMap is a free data retrieval call binding the contract method 0x740cb7f8.
//
// Solidity: function voucherMintNonceMap(uint256 ) view returns(uint256)
func (_TokenBank *TokenBankCallerSession) VoucherMintNonceMap(arg0 *big.Int) (*big.Int, error) {
	return _TokenBank.Contract.VoucherMintNonceMap(&_TokenBank.CallOpts, arg0)
}
