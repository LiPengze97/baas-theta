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

// MainchainTFuelTokenBankMetaData contains all meta data concerning the MainchainTFuelTokenBank contract.
var MainchainTFuelTokenBankMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internaltype\":\"contractsubchainregistrar\",\"name\":\"subchainregistrar_\",\"type\":\"address\"}],\"statemutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"subchainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"mainchaintokensender\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"subchainvoucherreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"lockedamount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"tokenlocknonce\",\"type\":\"uint256\"}],\"name\":\"tfueltokenlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"subchainid\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"indexed\":false,\"internaltype\":\"address\",\"name\":\"mainchaintokenreceiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"unlockedamount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"subchainvoucherburnnonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"tokenunlocknonce\",\"type\":\"uint256\"}],\"name\":\"tfueltokenunlocked\",\"type\":\"event\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alldenoms\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allvouchers\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"denomtovoucherlookup\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"contractaddress\",\"type\":\"address\"},{\"internaltype\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"exists\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"voucheraddress\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"vouchercontractaddr\",\"type\":\"address\"}],\"name\":\"getdenom\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"}],\"name\":\"getvoucher\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenlocknonceonsubchain\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internaltype\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenlockvotingrecords\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"accumlatedshares\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenunlocknonceonsubchain\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totallockedamounts\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voucheraddresstodenomlookup\",\"outputs\":[{\"internaltype\":\"string\",\"name\":\"denom\",\"type\":\"string\"},{\"internaltype\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voucherburnnonceonsubchain\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internaltype\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voucherburnvotingrecords\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"accumlatedshares\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vouchermintnonceonsubchain\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"statemutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"subchainid\",\"type\":\"uint256\"},{\"internaltype\":\"address\",\"name\":\"subchainvoucherreceiver\",\"type\":\"address\"}],\"name\":\"locktokens\",\"outputs\":[],\"statemutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"subchainid\",\"type\":\"uint256\"},{\"internaltype\":\"addresspayable\",\"name\":\"mainchaintokenreceiver\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"unlockamount\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"dynasty\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"subchainvoucherburnnonce\",\"type\":\"uint256\"}],\"name\":\"unlocktokens\",\"outputs\":[],\"statemutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161195438038061195483398101604081905261002f91610057565b506001805460008290556001600160a01b031981166001600160a01b03909116179055610087565b60006020828403121561006957600080fd5b81516001600160a01b038116811461008057600080fd5b9392505050565b6118be806100966000396000f3fe6080604052600436106100fe5760003560e01c8063a2cc698111610095578063e70891e211610064578063e70891e214610382578063ebda9962146103af578063f6a3d24e146103cf578063feaff0521461040b578063ff248a441461044a57600080fd5b8063a2cc6981146102f3578063a92678f814610313578063aa68acde14610340578063df953bd51461035557600080fd5b8063261a323e116100d1578063261a323e1461023057806327ca4df114610260578063588b14081461029857806360569b5e146102c557600080fd5b80630fa04ea1146101035780631527b14d1461014357806319fd1a11146101af5780631eb78737146101dc575b600080fd5b34801561010f57600080fd5b5061013061011e366004611464565b60046020526000908152604090205481565b6040519081526020015b60405180910390f35b34801561014f57600080fd5b5061019061015e3660046113cf565b8051602081830181018051600a825292820191909301209152546001600160a01b03811690600160a01b900460ff1682565b604080516001600160a01b03909316835290151560208301520161013a565b3480156101bb57600080fd5b506101306101ca366004611464565b600e6020526000908152604090205481565b3480156101e857600080fd5b5061021b6101f73660046114f5565b60086020908152600092835260408084209091529082529020805460029091015482565b6040805192835260208301919091520161013a565b34801561023c57600080fd5b5061025061024b3660046113cf565b61045d565b604051901515815260200161013a565b34801561026c57600080fd5b5061028061027b366004611464565b61049d565b6040516001600160a01b03909116815260200161013a565b3480156102a457600080fd5b506102b86102b3366004611464565b6104c7565b60405161013a91906115b9565b3480156102d157600080fd5b506102e56102e03660046112c4565b610573565b60405161013a9291906115cc565b3480156102ff57600080fd5b5061028061030e3660046113cf565b61061a565b34801561031f57600080fd5b5061013061032e366004611464565b60026020526000908152604090205481565b61035361034e36600461147d565b61068b565b005b34801561036157600080fd5b50610130610370366004611464565b60036020526000908152604090205481565b34801561038e57600080fd5b5061013061039d366004611464565b60056020526000908152604090205481565b3480156103bb57600080fd5b506102b86103ca3660046112c4565b610886565b3480156103db57600080fd5b506102506103ea3660046112c4565b6001600160a01b03166000908152600b602052604090206001015460ff1690565b34801561041757600080fd5b5061021b6104263660046114f5565b60096020908152600092835260408084209091529082529020805460029091015482565b6103536104583660046114ad565b61097d565b60008061046983610bc4565b9050600a8160405161047b9190611543565b9081526040519081900360200190205460ff600160a01b909104169392505050565b600c81815481106104ad57600080fd5b6000918252602090912001546001600160a01b0316905081565b600d81815481106104d757600080fd5b9060005260206000200160009150905080546104f29061178a565b80601f016020809104026020016040519081016040528092919081815260200182805461051e9061178a565b801561056b5780601f106105405761010080835404028352916020019161056b565b820191906000526020600020905b81548152906001019060200180831161054e57829003601f168201915b505050505081565b600b6020526000908152604090208054819061058e9061178a565b80601f01602080910402602001604051908101604052809291908181526020018280546105ba9061178a565b80156106075780601f106105dc57610100808354040283529160200191610607565b820191906000526020600020905b8154815290600101906020018083116105ea57829003601f168201915b5050506001909301549192505060ff1682565b60008061062683610bc4565b90506000600a8260405161063a9190611543565b908152604080516020928190038301812081830190925290546001600160a01b0381168252600160a01b900460ff1615801592820192909252915061068157519392505050565b5060009392505050565b600260005414156106e35760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064015b60405180910390fd5b60026000556001546040516343b71f0560e01b8152600481018490526001600160a01b03909116906343b71f059060240160206040518083038186803b15801561072c57600080fd5b505afa158015610740573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076491906113ad565b6107b05760405162461bcd60e51b815260206004820152601d60248201527f737562636861696e4944206e6f7420796574207265676973746572656400000060448201526064016106da565b6000828152600e6020526040902054349033906107cd9083610bd5565b6000858152600e60205260409020556107e584610be8565b600061082346604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a815260200161185f602a9139610c0f565b60008681526002602052604090819020549051919250907f2d1283f885740325813766f94db1fd044767ac7998c0a5e64bd154d67d319d3990610871908890859087908a908a908890611637565b60405180910390a15050600160005550505050565b6001600160a01b0381166000908152600b602052604080822081518083019092528054606093929190829082906108bc9061178a565b80601f01602080910402602001604051908101604052809291908181526020018280546108e89061178a565b80156109355780601f1061090a57610100808354040283529160200191610935565b820191906000526020600020905b81548152906001019060200180831161091857829003601f168201915b50505091835250506001919091015460ff16151560209182015281015190915015610961575192915050565b5050604080516020810190915260008152919050565b50919050565b600260005414156109d05760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016106da565b60026000908155858152600e6020526040902054831115610a475760405162461bcd60e51b815260206004820152602b60248201527f43616e6e6f7420756e6c6f636b207468652072657175657374656420616d6f7560448201526a1b9d081bd98815119d595b60aa1b60648201526084016106da565b6040805160208101879052908101849052606085811b6bffffffffffffffffffffffff191690820152607481018390526094810182905260009060b40160408051601f19818403018152919052805160208201209091506000610aad8886848733610c4d565b90508015610bb5576000888152600e6020526040902054610ace9087610c6a565b6000898152600e60205260408082209290925590516001600160a01b0389169188156108fc02918991818181858888f19350505050158015610b14573d6000803e3d6000fd5b50610b1e88610c76565b6000610b5c46604051806040016040528060018152602001600360fc1b8152506040518060600160405280602a815260200161185f602a9139610c0f565b60008a81526003602052604090819020549051919250907f578cbc38022da4d2e744a006cd760f34e7c0dc9127c15210d5bc5442447b726190610baa908c9085908d908d908c9088906115f0565b60405180910390a150505b50506001600055505050505050565b6060610bcf82610c95565b92915050565b6000610be182846116d3565b9392505050565b6000818152600260205260408120805460019290610c079084906116d3565b909155505050565b6060610c45610c1d85610d0f565b8484604051602001610c319392919061155f565b604051602081830303815290604052610bc4565b949350505050565b6000610c60868686868660076009610e0d565b9695505050505050565b6000610be18284611743565b6000818152600360205260408120805460019290610c079084906116d3565b60608160005b8151811015610d0857610ccd828281518110610cb957610cb961181a565b01602001516001600160f81b0319166111f7565b828281518110610cdf57610cdf61181a565b60200101906001600160f81b031916908160001a90535080610d00816117bf565b915050610c9b565b5092915050565b606081610d335750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610d5d5780610d47816117bf565b9150610d569050600a83611710565b9150610d37565b60008167ffffffffffffffff811115610d7857610d78611830565b6040519080825280601f01601f191660200182016040528015610da2576020820181803683370190505b5090505b8415610c4557610db7600183611743565b9150610dc4600a866117da565b610dcf9060306116d3565b60f81b818381518110610de457610de461181a565b60200101906001600160f81b031916908160001a905350610e06600a86611710565b9450610da6565b600087815260208390526040812054610e279060016116d3565b8514610e35575060006111ec565b6001546040516343b71f0560e01b8152600481018a90526001600160a01b03909116906343b71f059060240160206040518083038186803b158015610e7957600080fd5b505afa158015610e8d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb191906113ad565b610efd5760405162461bcd60e51b815260206004820152601d60248201527f737562636861696e4944206e6f7420726567697374657265642079657400000060448201526064016106da565b60008881526020838152604080832089845290915280822060015491516343f27e4560e01b8152600481018c9052602481018b90529091839182916001600160a01b0316906343f27e459060440160006040518083038186803b158015610f6357600080fd5b505afa158015610f77573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f9f91908101906112e1565b9150915060005b82518110156110fc57886001600160a01b0316838281518110610fcb57610fcb61181a565b60200260200101516001600160a01b031614610fe6576110ea565b6001945060005b60018501548110156110875784600101818154811061100e5761100e61181a565b6000918252602090912001546001600160a01b038b8116911614156110755760405162461bcd60e51b815260206004820152601c60248201527f546869732076616c696461746f7220616c726561647920766f7465640000000060448201526064016106da565b8061107f816117bf565b915050610fed565b508b84556001808501805491820181556000908152602090200180546001600160a01b0319163317905581516110e4908390839081106110c9576110c961181a565b60200260200101518560020154610bd590919063ffffffff16565b60028501555b806110f4816117bf565b915050610fa6565b508361113c5760405162461bcd60e51b815260206004820152600f60248201526e2737ba1030903b30b634b230ba37b960891b60448201526064016106da565b6000805b82518110156111895761117583828151811061115e5761115e61181a565b602002602001015183610bd590919063ffffffff16565b915080611181816117bf565b915050611140565b50611195816002611246565b60028501546111a5906003611246565b106111e25760008d8152602089905260409020546111c49060016116d3565b60008e815260208a9052604090205550600194506111ec9350505050565b6000955050505050505b979650505050505050565b6000604160f81b6001600160f81b03198316108015906112255750602d60f91b6001600160f81b0319831611155b156112425761123960f883901c60206116eb565b60f81b92915050565b5090565b6000610be18284611724565b600082601f83011261126357600080fd5b81516020611278611273836116af565b61167e565b80838252828201915082860187848660051b890101111561129857600080fd5b60005b858110156112b75781518452928401929084019060010161129b565b5090979650505050505050565b6000602082840312156112d657600080fd5b8135610be181611846565b600080604083850312156112f457600080fd5b825167ffffffffffffffff8082111561130c57600080fd5b818501915085601f83011261132057600080fd5b81516020611330611273836116af565b8083825282820191508286018a848660051b890101111561135057600080fd5b600096505b8487101561137c57805161136881611846565b835260019690960195918301918301611355565b509188015191965090935050508082111561139657600080fd5b506113a385828601611252565b9150509250929050565b6000602082840312156113bf57600080fd5b81518015158114610be157600080fd5b600060208083850312156113e257600080fd5b823567ffffffffffffffff808211156113fa57600080fd5b818501915085601f83011261140e57600080fd5b81358181111561142057611420611830565b611432601f8201601f1916850161167e565b9150808252868482850101111561144857600080fd5b8084840185840137600090820190930192909252509392505050565b60006020828403121561147657600080fd5b5035919050565b6000806040838503121561149057600080fd5b8235915060208301356114a281611846565b809150509250929050565b600080600080600060a086880312156114c557600080fd5b8535945060208601356114d781611846565b94979496505050506040830135926060810135926080909101359150565b6000806040838503121561150857600080fd5b50508035926020909101359150565b6000815180845261152f81602086016020860161175a565b601f01601f19169290920160200192915050565b6000825161155581846020870161175a565b9190910192915050565b6000845161157181846020890161175a565b8083019050602f60f81b8082528551611591816001850160208a0161175a565b600192019182015283516115ac81600284016020880161175a565b0160020195945050505050565b602081526000610be16020830184611517565b6040815260006115df6040830185611517565b905082151560208301529392505050565b86815260c06020820152600061160960c0830188611517565b6001600160a01b03969096166040830152506060810193909352608083019190915260a09091015292915050565b86815260c06020820152600061165060c0830188611517565b6001600160a01b039687166040840152949095166060820152608081019290925260a0909101529392505050565b604051601f8201601f1916810167ffffffffffffffff811182821017156116a7576116a7611830565b604052919050565b600067ffffffffffffffff8211156116c9576116c9611830565b5060051b60200190565b600082198211156116e6576116e66117ee565b500190565b600060ff821660ff84168060ff03821115611708576117086117ee565b019392505050565b60008261171f5761171f611804565b500490565b600081600019048311821515161561173e5761173e6117ee565b500290565b600082821015611755576117556117ee565b500390565b60005b8381101561177557818101518382015260200161175d565b83811115611784576000848401525b50505050565b600181811c9082168061179e57607f821691505b6020821081141561097757634e487b7160e01b600052602260045260246000fd5b60006000198214156117d3576117d36117ee565b5060010190565b6000826117e9576117e9611804565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461185b57600080fd5b5056fe307830303030303030303030303030303030303030303030303030303030303030303030303030303030a26469706673582212201183ce0108f2c96b9f06b3ead4859abfbf47a7ec64a04c63f8415c1649c3cc5264736f6c63430008070033",
}

// MainchainTFuelTokenBankABI is the input ABI used to generate the binding from.
// Deprecated: Use MainchainTFuelTokenBankMetaData.ABI instead.
var MainchainTFuelTokenBankABI = MainchainTFuelTokenBankMetaData.ABI

// MainchainTFuelTokenBankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainchainTFuelTokenBankMetaData.Bin instead.
var MainchainTFuelTokenBankBin = MainchainTFuelTokenBankMetaData.Bin

// DeployMainchainTFuelTokenBank deploys a new Ethereum contract, binding an instance of MainchainTFuelTokenBank to it.
func DeployMainchainTFuelTokenBank(auth *bind.TransactOpts, backend bind.ContractBackend, subchainregistrar_ common.Address) (common.Address, *types.Transaction, *MainchainTFuelTokenBank, error) {
	parsed, err := MainchainTFuelTokenBankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainchainTFuelTokenBankBin), backend, subchainregistrar_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MainchainTFuelTokenBank{MainchainTFuelTokenBankCaller: MainchainTFuelTokenBankCaller{contract: contract}, MainchainTFuelTokenBankTransactor: MainchainTFuelTokenBankTransactor{contract: contract}, MainchainTFuelTokenBankFilterer: MainchainTFuelTokenBankFilterer{contract: contract}}, nil
}

// MainchainTFuelTokenBank is an auto generated Go binding around an Ethereum contract.
type MainchainTFuelTokenBank struct {
	MainchainTFuelTokenBankCaller     // Read-only binding to the contract
	MainchainTFuelTokenBankTransactor // Write-only binding to the contract
	MainchainTFuelTokenBankFilterer   // Log filterer for contract events
}

// MainchainTFuelTokenBankCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainchainTFuelTokenBankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTFuelTokenBankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainchainTFuelTokenBankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTFuelTokenBankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainchainTFuelTokenBankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainchainTFuelTokenBankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainchainTFuelTokenBankSession struct {
	Contract     *MainchainTFuelTokenBank // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MainchainTFuelTokenBankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainchainTFuelTokenBankCallerSession struct {
	Contract *MainchainTFuelTokenBankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// MainchainTFuelTokenBankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainchainTFuelTokenBankTransactorSession struct {
	Contract     *MainchainTFuelTokenBankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// MainchainTFuelTokenBankRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainchainTFuelTokenBankRaw struct {
	Contract *MainchainTFuelTokenBank // Generic contract binding to access the raw methods on
}

// MainchainTFuelTokenBankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainchainTFuelTokenBankCallerRaw struct {
	Contract *MainchainTFuelTokenBankCaller // Generic read-only contract binding to access the raw methods on
}

// MainchainTFuelTokenBankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainchainTFuelTokenBankTransactorRaw struct {
	Contract *MainchainTFuelTokenBankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainchainTFuelTokenBank creates a new instance of MainchainTFuelTokenBank, bound to a specific deployed contract.
func NewMainchainTFuelTokenBank(address common.Address, backend bind.ContractBackend) (*MainchainTFuelTokenBank, error) {
	contract, err := bindMainchainTFuelTokenBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBank{MainchainTFuelTokenBankCaller: MainchainTFuelTokenBankCaller{contract: contract}, MainchainTFuelTokenBankTransactor: MainchainTFuelTokenBankTransactor{contract: contract}, MainchainTFuelTokenBankFilterer: MainchainTFuelTokenBankFilterer{contract: contract}}, nil
}

// NewMainchainTFuelTokenBankCaller creates a new read-only instance of MainchainTFuelTokenBank, bound to a specific deployed contract.
func NewMainchainTFuelTokenBankCaller(address common.Address, caller bind.ContractCaller) (*MainchainTFuelTokenBankCaller, error) {
	contract, err := bindMainchainTFuelTokenBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankCaller{contract: contract}, nil
}

// NewMainchainTFuelTokenBankTransactor creates a new write-only instance of MainchainTFuelTokenBank, bound to a specific deployed contract.
func NewMainchainTFuelTokenBankTransactor(address common.Address, transactor bind.ContractTransactor) (*MainchainTFuelTokenBankTransactor, error) {
	contract, err := bindMainchainTFuelTokenBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankTransactor{contract: contract}, nil
}

// NewMainchainTFuelTokenBankFilterer creates a new log filterer instance of MainchainTFuelTokenBank, bound to a specific deployed contract.
func NewMainchainTFuelTokenBankFilterer(address common.Address, filterer bind.ContractFilterer) (*MainchainTFuelTokenBankFilterer, error) {
	contract, err := bindMainchainTFuelTokenBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankFilterer{contract: contract}, nil
}

// bindMainchainTFuelTokenBank binds a generic wrapper to an already deployed contract.
func bindMainchainTFuelTokenBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainchainTFuelTokenBankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainTFuelTokenBank.Contract.MainchainTFuelTokenBankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.MainchainTFuelTokenBankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.MainchainTFuelTokenBankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainchainTFuelTokenBank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.contract.Transact(opts, method, params...)
}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Alldenoms(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "alldenoms", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Alldenoms(arg0 *big.Int) (string, error) {
	return _MainchainTFuelTokenBank.Contract.Alldenoms(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Alldenoms is a free data retrieval call binding the contract method 0xbc15cedc.
//
// Solidity: function alldenoms(uint256 ) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Alldenoms(arg0 *big.Int) (string, error) {
	return _MainchainTFuelTokenBank.Contract.Alldenoms(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Allvouchers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "allvouchers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Allvouchers(arg0 *big.Int) (common.Address, error) {
	return _MainchainTFuelTokenBank.Contract.Allvouchers(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Allvouchers is a free data retrieval call binding the contract method 0x11418b8e.
//
// Solidity: function allvouchers(uint256 ) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Allvouchers(arg0 *big.Int) (common.Address, error) {
	return _MainchainTFuelTokenBank.Contract.Allvouchers(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Denomtovoucherlookup(opts *bind.CallOpts, arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "denomtovoucherlookup", arg0)

	outstruct := new(struct {
		Contractaddress common.Address
		Exists          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Contractaddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Denomtovoucherlookup(arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	return _MainchainTFuelTokenBank.Contract.Denomtovoucherlookup(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Denomtovoucherlookup is a free data retrieval call binding the contract method 0x4c68012b.
//
// Solidity: function denomtovoucherlookup(string ) view returns(address contractaddress, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Denomtovoucherlookup(arg0 string) (struct {
	Contractaddress common.Address
	Exists          bool
}, error) {
	return _MainchainTFuelTokenBank.Contract.Denomtovoucherlookup(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Exists(opts *bind.CallOpts, denom string) (bool, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "exists", denom)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Exists(denom string) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.Exists(&_MainchainTFuelTokenBank.CallOpts, denom)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string denom) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Exists(denom string) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.Exists(&_MainchainTFuelTokenBank.CallOpts, denom)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Exists0(opts *bind.CallOpts, voucheraddress common.Address) (bool, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "exists0", voucheraddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Exists0(voucheraddress common.Address) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.Exists0(&_MainchainTFuelTokenBank.CallOpts, voucheraddress)
}

// Exists0 is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address voucheraddress) view returns(bool)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Exists0(voucheraddress common.Address) (bool, error) {
	return _MainchainTFuelTokenBank.Contract.Exists0(&_MainchainTFuelTokenBank.CallOpts, voucheraddress)
}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Getdenom(opts *bind.CallOpts, vouchercontractaddr common.Address) (string, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "getdenom", vouchercontractaddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Getdenom(vouchercontractaddr common.Address) (string, error) {
	return _MainchainTFuelTokenBank.Contract.Getdenom(&_MainchainTFuelTokenBank.CallOpts, vouchercontractaddr)
}

// Getdenom is a free data retrieval call binding the contract method 0xaf46078f.
//
// Solidity: function getdenom(address vouchercontractaddr) view returns(string)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Getdenom(vouchercontractaddr common.Address) (string, error) {
	return _MainchainTFuelTokenBank.Contract.Getdenom(&_MainchainTFuelTokenBank.CallOpts, vouchercontractaddr)
}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Getvoucher(opts *bind.CallOpts, denom string) (common.Address, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "getvoucher", denom)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Getvoucher(denom string) (common.Address, error) {
	return _MainchainTFuelTokenBank.Contract.Getvoucher(&_MainchainTFuelTokenBank.CallOpts, denom)
}

// Getvoucher is a free data retrieval call binding the contract method 0xd2990e7d.
//
// Solidity: function getvoucher(string denom) view returns(address)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Getvoucher(denom string) (common.Address, error) {
	return _MainchainTFuelTokenBank.Contract.Getvoucher(&_MainchainTFuelTokenBank.CallOpts, denom)
}

// Tokenlocknonceonsubchain is a free data retrieval call binding the contract method 0x3bf4ecc2.
//
// Solidity: function tokenlocknonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Tokenlocknonceonsubchain(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "tokenlocknonceonsubchain", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokenlocknonceonsubchain is a free data retrieval call binding the contract method 0x3bf4ecc2.
//
// Solidity: function tokenlocknonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Tokenlocknonceonsubchain(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Tokenlocknonceonsubchain(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Tokenlocknonceonsubchain is a free data retrieval call binding the contract method 0x3bf4ecc2.
//
// Solidity: function tokenlocknonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Tokenlocknonceonsubchain(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Tokenlocknonceonsubchain(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Tokenlockvotingrecords is a free data retrieval call binding the contract method 0x1b637e32.
//
// Solidity: function tokenlockvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Tokenlockvotingrecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "tokenlockvotingrecords", arg0, arg1)

	outstruct := new(struct {
		Dynasty          *big.Int
		Accumlatedshares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dynasty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Accumlatedshares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tokenlockvotingrecords is a free data retrieval call binding the contract method 0x1b637e32.
//
// Solidity: function tokenlockvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Tokenlockvotingrecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	return _MainchainTFuelTokenBank.Contract.Tokenlockvotingrecords(&_MainchainTFuelTokenBank.CallOpts, arg0, arg1)
}

// Tokenlockvotingrecords is a free data retrieval call binding the contract method 0x1b637e32.
//
// Solidity: function tokenlockvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Tokenlockvotingrecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	return _MainchainTFuelTokenBank.Contract.Tokenlockvotingrecords(&_MainchainTFuelTokenBank.CallOpts, arg0, arg1)
}

// Tokenunlocknonceonsubchain is a free data retrieval call binding the contract method 0xbfbf56c4.
//
// Solidity: function tokenunlocknonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Tokenunlocknonceonsubchain(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "tokenunlocknonceonsubchain", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tokenunlocknonceonsubchain is a free data retrieval call binding the contract method 0xbfbf56c4.
//
// Solidity: function tokenunlocknonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Tokenunlocknonceonsubchain(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Tokenunlocknonceonsubchain(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Tokenunlocknonceonsubchain is a free data retrieval call binding the contract method 0xbfbf56c4.
//
// Solidity: function tokenunlocknonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Tokenunlocknonceonsubchain(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Tokenunlocknonceonsubchain(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Totallockedamounts is a free data retrieval call binding the contract method 0x5a35e496.
//
// Solidity: function totallockedamounts(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Totallockedamounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "totallockedamounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Totallockedamounts is a free data retrieval call binding the contract method 0x5a35e496.
//
// Solidity: function totallockedamounts(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Totallockedamounts(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Totallockedamounts(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Totallockedamounts is a free data retrieval call binding the contract method 0x5a35e496.
//
// Solidity: function totallockedamounts(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Totallockedamounts(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Totallockedamounts(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Voucheraddresstodenomlookup(opts *bind.CallOpts, arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "voucheraddresstodenomlookup", arg0)

	outstruct := new(struct {
		Denom  string
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Denom = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Voucheraddresstodenomlookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _MainchainTFuelTokenBank.Contract.Voucheraddresstodenomlookup(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Voucheraddresstodenomlookup is a free data retrieval call binding the contract method 0xcd22e450.
//
// Solidity: function voucheraddresstodenomlookup(address ) view returns(string denom, bool exists)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Voucheraddresstodenomlookup(arg0 common.Address) (struct {
	Denom  string
	Exists bool
}, error) {
	return _MainchainTFuelTokenBank.Contract.Voucheraddresstodenomlookup(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Voucherburnnonceonsubchain is a free data retrieval call binding the contract method 0x6bb4e464.
//
// Solidity: function voucherburnnonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Voucherburnnonceonsubchain(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "voucherburnnonceonsubchain", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Voucherburnnonceonsubchain is a free data retrieval call binding the contract method 0x6bb4e464.
//
// Solidity: function voucherburnnonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Voucherburnnonceonsubchain(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Voucherburnnonceonsubchain(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Voucherburnnonceonsubchain is a free data retrieval call binding the contract method 0x6bb4e464.
//
// Solidity: function voucherburnnonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Voucherburnnonceonsubchain(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Voucherburnnonceonsubchain(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Voucherburnvotingrecords is a free data retrieval call binding the contract method 0x147ea516.
//
// Solidity: function voucherburnvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Voucherburnvotingrecords(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "voucherburnvotingrecords", arg0, arg1)

	outstruct := new(struct {
		Dynasty          *big.Int
		Accumlatedshares *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dynasty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Accumlatedshares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voucherburnvotingrecords is a free data retrieval call binding the contract method 0x147ea516.
//
// Solidity: function voucherburnvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Voucherburnvotingrecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	return _MainchainTFuelTokenBank.Contract.Voucherburnvotingrecords(&_MainchainTFuelTokenBank.CallOpts, arg0, arg1)
}

// Voucherburnvotingrecords is a free data retrieval call binding the contract method 0x147ea516.
//
// Solidity: function voucherburnvotingrecords(uint256 , bytes32 ) view returns(uint256 dynasty, uint256 accumlatedshares)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Voucherburnvotingrecords(arg0 *big.Int, arg1 [32]byte) (struct {
	Dynasty          *big.Int
	Accumlatedshares *big.Int
}, error) {
	return _MainchainTFuelTokenBank.Contract.Voucherburnvotingrecords(&_MainchainTFuelTokenBank.CallOpts, arg0, arg1)
}

// Vouchermintnonceonsubchain is a free data retrieval call binding the contract method 0x35cccd8c.
//
// Solidity: function vouchermintnonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCaller) Vouchermintnonceonsubchain(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MainchainTFuelTokenBank.contract.Call(opts, &out, "vouchermintnonceonsubchain", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vouchermintnonceonsubchain is a free data retrieval call binding the contract method 0x35cccd8c.
//
// Solidity: function vouchermintnonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Vouchermintnonceonsubchain(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Vouchermintnonceonsubchain(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Vouchermintnonceonsubchain is a free data retrieval call binding the contract method 0x35cccd8c.
//
// Solidity: function vouchermintnonceonsubchain(uint256 ) view returns(uint256)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankCallerSession) Vouchermintnonceonsubchain(arg0 *big.Int) (*big.Int, error) {
	return _MainchainTFuelTokenBank.Contract.Vouchermintnonceonsubchain(&_MainchainTFuelTokenBank.CallOpts, arg0)
}

// Locktokens is a paid mutator transaction binding the contract method 0xd9a3958f.
//
// Solidity: function locktokens(uint256 subchainid, address subchainvoucherreceiver) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactor) Locktokens(opts *bind.TransactOpts, subchainid *big.Int, subchainvoucherreceiver common.Address) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.contract.Transact(opts, "locktokens", subchainid, subchainvoucherreceiver)
}

// Locktokens is a paid mutator transaction binding the contract method 0xd9a3958f.
//
// Solidity: function locktokens(uint256 subchainid, address subchainvoucherreceiver) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Locktokens(subchainid *big.Int, subchainvoucherreceiver common.Address) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.Locktokens(&_MainchainTFuelTokenBank.TransactOpts, subchainid, subchainvoucherreceiver)
}

// Locktokens is a paid mutator transaction binding the contract method 0xd9a3958f.
//
// Solidity: function locktokens(uint256 subchainid, address subchainvoucherreceiver) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactorSession) Locktokens(subchainid *big.Int, subchainvoucherreceiver common.Address) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.Locktokens(&_MainchainTFuelTokenBank.TransactOpts, subchainid, subchainvoucherreceiver)
}

// Unlocktokens is a paid mutator transaction binding the contract method 0x9fc3fc8a.
//
// Solidity: function unlocktokens(uint256 subchainid, address mainchaintokenreceiver, uint256 unlockamount, uint256 dynasty, uint256 subchainvoucherburnnonce) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactor) Unlocktokens(opts *bind.TransactOpts, subchainid *big.Int, mainchaintokenreceiver common.Address, unlockamount *big.Int, dynasty *big.Int, subchainvoucherburnnonce *big.Int) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.contract.Transact(opts, "unlocktokens", subchainid, mainchaintokenreceiver, unlockamount, dynasty, subchainvoucherburnnonce)
}

// Unlocktokens is a paid mutator transaction binding the contract method 0x9fc3fc8a.
//
// Solidity: function unlocktokens(uint256 subchainid, address mainchaintokenreceiver, uint256 unlockamount, uint256 dynasty, uint256 subchainvoucherburnnonce) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankSession) Unlocktokens(subchainid *big.Int, mainchaintokenreceiver common.Address, unlockamount *big.Int, dynasty *big.Int, subchainvoucherburnnonce *big.Int) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.Unlocktokens(&_MainchainTFuelTokenBank.TransactOpts, subchainid, mainchaintokenreceiver, unlockamount, dynasty, subchainvoucherburnnonce)
}

// Unlocktokens is a paid mutator transaction binding the contract method 0x9fc3fc8a.
//
// Solidity: function unlocktokens(uint256 subchainid, address mainchaintokenreceiver, uint256 unlockamount, uint256 dynasty, uint256 subchainvoucherburnnonce) payable returns()
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankTransactorSession) Unlocktokens(subchainid *big.Int, mainchaintokenreceiver common.Address, unlockamount *big.Int, dynasty *big.Int, subchainvoucherburnnonce *big.Int) (*types.Transaction, error) {
	return _MainchainTFuelTokenBank.Contract.Unlocktokens(&_MainchainTFuelTokenBank.TransactOpts, subchainid, mainchaintokenreceiver, unlockamount, dynasty, subchainvoucherburnnonce)
}

// MainchainTFuelTokenBankTfueltokenlockedIterator is returned from FilterTfueltokenlocked and is used to iterate over the raw logs and unpacked data for Tfueltokenlocked events raised by the MainchainTFuelTokenBank contract.
type MainchainTFuelTokenBankTfueltokenlockedIterator struct {
	Event *MainchainTFuelTokenBankTfueltokenlocked // Event containing the contract specifics and raw log

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
func (it *MainchainTFuelTokenBankTfueltokenlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainTFuelTokenBankTfueltokenlocked)
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
		it.Event = new(MainchainTFuelTokenBankTfueltokenlocked)
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
func (it *MainchainTFuelTokenBankTfueltokenlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainTFuelTokenBankTfueltokenlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainTFuelTokenBankTfueltokenlocked represents a Tfueltokenlocked event raised by the MainchainTFuelTokenBank contract.
type MainchainTFuelTokenBankTfueltokenlocked struct {
	Subchainid              *big.Int
	Denom                   string
	Mainchaintokensender    common.Address
	Subchainvoucherreceiver common.Address
	Lockedamount            *big.Int
	Tokenlocknonce          *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTfueltokenlocked is a free log retrieval operation binding the contract event 0x3c846a66063619d8beffa327e943e3620e804def6524603a44f18c74d879e530.
//
// Solidity: event tfueltokenlocked(uint256 subchainid, string denom, address mainchaintokensender, address subchainvoucherreceiver, uint256 lockedamount, uint256 tokenlocknonce)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) FilterTfueltokenlocked(opts *bind.FilterOpts) (*MainchainTFuelTokenBankTfueltokenlockedIterator, error) {

	logs, sub, err := _MainchainTFuelTokenBank.contract.FilterLogs(opts, "tfueltokenlocked")
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankTfueltokenlockedIterator{contract: _MainchainTFuelTokenBank.contract, event: "tfueltokenlocked", logs: logs, sub: sub}, nil
}

// WatchTfueltokenlocked is a free log subscription operation binding the contract event 0x3c846a66063619d8beffa327e943e3620e804def6524603a44f18c74d879e530.
//
// Solidity: event tfueltokenlocked(uint256 subchainid, string denom, address mainchaintokensender, address subchainvoucherreceiver, uint256 lockedamount, uint256 tokenlocknonce)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) WatchTfueltokenlocked(opts *bind.WatchOpts, sink chan<- *MainchainTFuelTokenBankTfueltokenlocked) (event.Subscription, error) {

	logs, sub, err := _MainchainTFuelTokenBank.contract.WatchLogs(opts, "tfueltokenlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainTFuelTokenBankTfueltokenlocked)
				if err := _MainchainTFuelTokenBank.contract.UnpackLog(event, "tfueltokenlocked", log); err != nil {
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

// ParseTfueltokenlocked is a log parse operation binding the contract event 0x3c846a66063619d8beffa327e943e3620e804def6524603a44f18c74d879e530.
//
// Solidity: event tfueltokenlocked(uint256 subchainid, string denom, address mainchaintokensender, address subchainvoucherreceiver, uint256 lockedamount, uint256 tokenlocknonce)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) ParseTfueltokenlocked(log types.Log) (*MainchainTFuelTokenBankTfueltokenlocked, error) {
	event := new(MainchainTFuelTokenBankTfueltokenlocked)
	if err := _MainchainTFuelTokenBank.contract.UnpackLog(event, "tfueltokenlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainchainTFuelTokenBankTfueltokenunlockedIterator is returned from FilterTfueltokenunlocked and is used to iterate over the raw logs and unpacked data for Tfueltokenunlocked events raised by the MainchainTFuelTokenBank contract.
type MainchainTFuelTokenBankTfueltokenunlockedIterator struct {
	Event *MainchainTFuelTokenBankTfueltokenunlocked // Event containing the contract specifics and raw log

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
func (it *MainchainTFuelTokenBankTfueltokenunlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainchainTFuelTokenBankTfueltokenunlocked)
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
		it.Event = new(MainchainTFuelTokenBankTfueltokenunlocked)
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
func (it *MainchainTFuelTokenBankTfueltokenunlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainchainTFuelTokenBankTfueltokenunlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainchainTFuelTokenBankTfueltokenunlocked represents a Tfueltokenunlocked event raised by the MainchainTFuelTokenBank contract.
type MainchainTFuelTokenBankTfueltokenunlocked struct {
	Subchainid               *big.Int
	Denom                    string
	Mainchaintokenreceiver   common.Address
	Unlockedamount           *big.Int
	Subchainvoucherburnnonce *big.Int
	Tokenunlocknonce         *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterTfueltokenunlocked is a free log retrieval operation binding the contract event 0x44512cf6e5b61650c4a6360296ced8edae2b1e431e28a3047adfd4221264c4bc.
//
// Solidity: event tfueltokenunlocked(uint256 subchainid, string denom, address mainchaintokenreceiver, uint256 unlockedamount, uint256 subchainvoucherburnnonce, uint256 tokenunlocknonce)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) FilterTfueltokenunlocked(opts *bind.FilterOpts) (*MainchainTFuelTokenBankTfueltokenunlockedIterator, error) {

	logs, sub, err := _MainchainTFuelTokenBank.contract.FilterLogs(opts, "tfueltokenunlocked")
	if err != nil {
		return nil, err
	}
	return &MainchainTFuelTokenBankTfueltokenunlockedIterator{contract: _MainchainTFuelTokenBank.contract, event: "tfueltokenunlocked", logs: logs, sub: sub}, nil
}

// WatchTfueltokenunlocked is a free log subscription operation binding the contract event 0x44512cf6e5b61650c4a6360296ced8edae2b1e431e28a3047adfd4221264c4bc.
//
// Solidity: event tfueltokenunlocked(uint256 subchainid, string denom, address mainchaintokenreceiver, uint256 unlockedamount, uint256 subchainvoucherburnnonce, uint256 tokenunlocknonce)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) WatchTfueltokenunlocked(opts *bind.WatchOpts, sink chan<- *MainchainTFuelTokenBankTfueltokenunlocked) (event.Subscription, error) {

	logs, sub, err := _MainchainTFuelTokenBank.contract.WatchLogs(opts, "tfueltokenunlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainchainTFuelTokenBankTfueltokenunlocked)
				if err := _MainchainTFuelTokenBank.contract.UnpackLog(event, "tfueltokenunlocked", log); err != nil {
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

// ParseTfueltokenunlocked is a log parse operation binding the contract event 0x44512cf6e5b61650c4a6360296ced8edae2b1e431e28a3047adfd4221264c4bc.
//
// Solidity: event tfueltokenunlocked(uint256 subchainid, string denom, address mainchaintokenreceiver, uint256 unlockedamount, uint256 subchainvoucherburnnonce, uint256 tokenunlocknonce)
func (_MainchainTFuelTokenBank *MainchainTFuelTokenBankFilterer) ParseTfueltokenunlocked(log types.Log) (*MainchainTFuelTokenBankTfueltokenunlocked, error) {
	event := new(MainchainTFuelTokenBankTfueltokenunlocked)
	if err := _MainchainTFuelTokenBank.contract.UnpackLog(event, "tfueltokenunlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
