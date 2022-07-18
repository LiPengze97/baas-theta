package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/thetatoken/theta/common"
	score "github.com/thetatoken/thetasubchain/core"
	"github.com/thetatoken/thetasubchain/eth/abi"
	"github.com/thetatoken/thetasubchain/eth/ethclient"
	ct "github.com/thetatoken/thetasubchain/integration/tools/go_new/accessors"
	scta "github.com/thetatoken/thetasubchain/interchain/contracts/accessors"
	// rg "chainRegistrarOnMainchain" // for demo
)

// func main1() {
// 	//registerAndStake()
// 	AccountsInit()
// 	client, err := ethclient.Dial("http://localhost:18888/rpc")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	subchainID := big.NewInt(9988)
// 	var dec18 = new(big.Int)
// 	dec18.SetString("1000000000000000000", 10)
// 	user := accountList[6].fromAddress

// 	instanceTNT20VoucherContract, err := ct.NewTNT20VoucherContract(TNT20VoucherContractAddress, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	instanceTNT20TokenBank, err := ct.NewTNT20TokenBank(TNT20TokenBankAddress, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	authAccount0 := SelectAccount(client, 0)
// 	_, err = instanceTNT20VoucherContract.Mint(authAccount0, user, big.NewInt(30))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	authUser := SelectAccount(client, 6)
// 	_, err = instanceTNT20VoucherContract.Approve(authUser, TNT20TokenBankAddress, big.NewInt(30))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	authUser = SelectAccount(client, 6)
// 	_, err = instanceTNT20TokenBank.LockTokens(authUser, subchainID, TNT20VoucherContractAddress, user, big.NewInt(20))
// 	height, _ := client.BlockNumber(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("ok!")
// 	fmt.Println(height)
// 	//validatorStakingAmount := new(big.Int).Mul(dec18, big.NewInt(100000))
// 	//validator1 := accountList[1].fromAddress
// 	//WthetaAddress := common.HexToAddress("0x17deB845AAcCA09873D1984784EbdC59DBB38846")
// 	//WthetaAddress := common.HexToAddress("0x5395019D1d4794eb5Ac0AC51976ee48995bdA694")

// 	//GovernanceTokenAddress := common.HexToAddress("0x6c1d3aBA75Fea74000a10BbfbE86be30f20a2F3D")
// 	//TNT20TokenBankAddress := common.HexToAddress("0x1f629139b3b4A03799c6e6655b7F59a1F01598E7")
// 	// clientNew, _ := SelectAccount(client, 3) //chainGuarantor
// 	// instance, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, clientNew)
// 	// //height, _ := client.BlockNumber(context.Background())
// 	// //height1 := big.NewInt(int64(height))
// 	// tx, err := instance.GetStakeSnapshotHeights(nil, subchainID)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(tx)
// 	//instanceChainRegistrarValidator3, err := ct.NewChainRegistrarOnMainchain(RegisterOnMainchainAddress, client)
// 	// clientValidator3, authValidator3 = chooseAccount(client, 3) //Validator1
// 	// _, err = instanceChainRegistrarValidator3.DepositStake(authValidator3, subchainID, accountList[3].fromAddress, validatorStakingAmount)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// time.Sleep(2 * time.Second)
// 	// tx, err := instanceChainRegistrarValidator3.GetValidatorSet(nil, subchainID, big.NewInt(4))
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(tx)
// }
func main2() {
	//registerAndStake()
	AccountsInit()
	//client, err := ethclient.Dial("http://localhost:18888/rpc")
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	//subchainAddress:=common.
	if err != nil {
		log.Fatal(err)
	}
	subchainID := big.NewInt(9988)
	var dec18 = new(big.Int)
	dec18.SetString("1000000000000000000", 10)
	user := accountList[6].fromAddress

	instanceTNT20VoucherContract, err := ct.NewTNT20VoucherContract(TNT20VoucherContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	instanceTNT20TokenBank, err := ct.NewTNT20TokenBank(TNT20TokenBankAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 := SelectAccount(client, 0)
	fmt.Println(instanceTNT20VoucherContract.BalanceOf(nil, accountList[6].fromAddress))
	_, err = instanceTNT20VoucherContract.Mint(authAccount0, user, big.NewInt(30))
	if err != nil {
		log.Fatal(err)
	}
	authAccount0 = SelectAccount(client, 0)
	fmt.Println(instanceTNT20VoucherContract.BalanceOf(nil, accountList[6].fromAddress))

	authUser := SelectAccount(client, 6)
	_, err = instanceTNT20VoucherContract.Approve(authUser, TNT20TokenBankAddress, big.NewInt(30))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instanceTNT20TokenBank.GetMaxProcessedTokenLockNonce(nil, subchainID))
	authUser = SelectAccount(client, 6)
	_, err = instanceTNT20TokenBank.LockTokens(authUser, subchainID, TNT20VoucherContractAddress, user, big.NewInt(20))
	// height, _ := client.BlockNumber(context.Background())
	time.Sleep(2 * time.Second)
	fmt.Println(instanceTNT20TokenBank.GetMaxProcessedTokenLockNonce(nil, subchainID))
	fmt.Println(instanceTNT20VoucherContract.BalanceOf(nil, accountList[6].fromAddress))

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("ok!")
	// fmt.Println(height)
}
func extract() {
	Data := "0x0000000000000000000000000000000000000000000000000000000000000120000000000000000000000000ae72a48c1a36bd18af168541c53037965d26e4a80000000000000000000000000000000000000000000000000000000000002704000000000000000000000000ae72a48c1a36bd18af168541c53037965d26e4a80000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000d00000000000000000000000000000000000000000000000000000000000000313336362f32302f3078353766346230386530623363633561396362633838383633396462653231373162313430383732320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002323000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000023230000000000000000000000000000000000000000000000000000000000000"
	data, _ := hex.DecodeString(Data[2:])
	var tma score.CrossChainTNT20TokenLockedEvent
	contractAbi, _ := abi.JSON(strings.NewReader(string(scta.TNT20TokenBankABI)))
	contractAbi.UnpackIntoInterface(&tma, "TNT20TokenLocked", data)

	fmt.Println(tma.TargetChainID)
}
func extract2() {
	RawABI := scta.TNT20TokenBankABI
	parsed, _ := abi.JSON(strings.NewReader(RawABI))
	// var res struct {
	// 	Receiver []common.Address // 返回值名称
	// 	Values   []*big.Int       // 返回值名称
	// }
	var tma score.CrossChainTNT20TokenLockedEvent
	// {"Receiver":["0x80819b3f30e9d77de6be3df9d6efaa88261dff9c"],"Values":[10]}
	raw := common.Hex2Bytes("0x0000000000000000000000000000000000000000000000000000000000000120000000000000000000000000ae72a48c1a36bd18af168541c53037965d26e4a80000000000000000000000000000000000000000000000000000000000002704000000000000000000000000ae72a48c1a36bd18af168541c53037965d26e4a80000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000d00000000000000000000000000000000000000000000000000000000000000313336362f32302f3078353766346230386530623363633561396362633838383633396462653231373162313430383732320000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002323000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000023230000000000000000000000000000000000000000000000000000000000000")
	parsed.UnpackIntoInterface(&tma, "TNT20TokenLocked", raw)
	fmt.Println(tma.TokenLockNonce)

}

func main3() {
	//registerAndStake()
	AccountsInit()
	//client, err := ethclient.Dial("http://localhost:18888/rpc")
	client, err := ethclient.Dial("http://localhost:19888/rpc")
	subchainAddress := common.HexToAddress("0x47e9fbef8c83a1714f1951f142132e6e90f5fa5d")
	if err != nil {
		log.Fatal(err)
	}
	instance, err := ct.NewTNT20TokenBank(subchainAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instance.GetMaxProcessedTokenLockNonce(nil, big.NewInt(11)))
}
func main() {
	main3()
	//extract()
}