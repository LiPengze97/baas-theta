package tx

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/ledger/types"
	wtypes "github.com/thetatoken/theta/wallet/types"
	"github.com/thetatoken/thetasubchain/cmd/thetasubcli/cmd/utils"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
	"github.com/thetatoken/thetasubchain/rpc"

	"github.com/ybbus/jsonrpc"
	rpcc "github.com/ybbus/jsonrpc"
)

// sendCmd represents the send command
// Example:
//		thetasubcli tx send --chain="privatenet" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --to=9F1233798E905E173560071255140b4A8aBd3Ec6 --tfuel=9 --seq=1
//		thetasubcli tx send --chain="privatenet" --path "m/44'/60'/0'/0/0" --to=9F1233798E905E173560071255140b4A8aBd3Ec6 --tfuel=9 --seq=1 --wallet=trezor
//		thetasubcli tx send --chain="privatenet" --path "m/44'/60'/0'/0" --to=9F1233798E905E173560071255140b4A8aBd3Ec6 --tfuel=9 --seq=1 --wallet=nano
var sendCmd = &cobra.Command{
	Use:     "send",
	Short:   "Send tokens",
	Example: `thetasubcli tx send --chain="privatenet" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --to=9F1233798E905E173560071255140b4A8aBd3Ec6 --tfuel=9 --seq=1`,
	Run:     doSendCmd,
}

var autoSendCmd = &cobra.Command{
	Use:     "autosend",
	Short:   "Send tokens",
	Example: `thetasubcli tx send --chain="privatenet" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --to=9F1233798E905E173560071255140b4A8aBd3Ec6 --tfuel=9 --seq=1`,
	Run:     doAutoSendCmd,
}

var autoQueueSendCmd = &cobra.Command{
	Use:     "autoqueue",
	Short:   "Send tokens",
	Example: `thetasubcli tx send --chain="privatenet" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --to=9F1233798E905E173560071255140b4A8aBd3Ec6 --tfuel=9 --seq=1`,
	Run:     doAutoQueueSendCmd,
}

var autoMultiQueueSendCmd = &cobra.Command{
	Use:     "automq",
	Short:   "Send tokens",
	Example: `thetasubcli tx send --chain="privatenet" --from=2E833968E5bB786Ae419c4d13189fB081Cc43bab --to=9F1233798E905E173560071255140b4A8aBd3Ec6 --tfuel=9 --seq=1`,
	Run:     doAutoQueueMultiSendCmd,
}

func doSendCmd(cmd *cobra.Command, args []string) {
	walletType := getWalletType(cmd)
	if walletType == wtypes.WalletTypeSoft && len(fromFlag) == 0 {
		utils.Error("The from address cannot be empty") // we don't need to specify the "from address" for hardware wallets
		return
	}

	if len(toFlag) == 0 {
		utils.Error("The to address cannot be empty")
		return
	}
	if fromFlag == toFlag {
		utils.Error("The from and to address cannot be identical")
		return
	}

	wallet, fromAddress, err := walletUnlockWithPath(cmd, fromFlag, pathFlag, passwordFlag)
	if err != nil || wallet == nil {
		return
	}
	defer wallet.Lock(fromAddress)

	tfuel, ok := types.ParseCoinAmount(tfuelAmountFlag)
	if !ok {
		utils.Error("Failed to parse tfuel amount")
	}
	fee, ok := types.ParseCoinAmount(feeFlag)
	if !ok {
		utils.Error("Failed to parse fee")
	}
	inputs := []types.TxInput{{
		Address: fromAddress,
		Coins: types.Coins{
			TFuelWei: new(big.Int).Add(tfuel, fee),
			ThetaWei: new(big.Int).SetUint64(0),
		},
		Sequence: uint64(seqFlag),
	}}
	outputs := []types.TxOutput{{
		Address: common.HexToAddress(toFlag),
		Coins: types.Coins{
			TFuelWei: tfuel,
			ThetaWei: new(big.Int).SetUint64(0),
		},
	}}
	sendTx := &types.SendTx{
		Fee: types.Coins{
			ThetaWei: new(big.Int).SetUint64(0),
			TFuelWei: fee,
		},
		Inputs:  inputs,
		Outputs: outputs,
	}

	sig, err := wallet.Sign(fromAddress, sendTx.SignBytes(chainIDFlag))
	if err != nil {
		utils.Error("Failed to sign transaction: %v\n", err)
	}
	sendTx.SetSignature(fromAddress, sig)

	raw, err := stypes.TxToBytes(sendTx)
	if err != nil {
		utils.Error("Failed to encode transaction: %v\n", err)
	}
	signedTx := hex.EncodeToString(raw)
	// remoteRPCEndpoints := []string{"http://10.10.1.1:16900/rpc", "http://10.10.1.2:16900/rpc", "http://10.10.1.3:16900/rpc", "http://10.10.1.4:16900/rpc"}
	remoteRPCEndpoints := []string{"http://127.0.0.1:16930/rpc", "http://127.0.0.1:16910/rpc", "http://127.0.0.1:16920/rpc", "http://127.0.0.1:16900/rpc"}
	// remoteRPCEndpoints := []string{"http://127.0.0.1:16900/rpc", "http://127.0.0.1:16910/rpc", "http://127.0.0.1:16920/rpc"}
	var wg sync.WaitGroup

	for idx, remoteRPCEndpoint := range remoteRPCEndpoints {
		f := func(remoteRPCEndpoint string) {
			defer wg.Done()
			// if remoteRPCEndpoint == "http://127.0.0.1:16900/rpc" {
			// 	time.Sleep(time.Duration(3) * time.Second)
			// }
			client := rpcc.NewRPCClient(remoteRPCEndpoint)

			var res *jsonrpc.RPCResponse
			if asyncFlag {
				res, err = client.Call("theta.BroadcastRawTransactionAsync", rpc.BroadcastRawTransactionArgs{TxBytes: signedTx})
			} else {
				res, err = client.Call("theta.SendSoleRawTransaction", rpc.SendSoleRawTransactionArgs{TxBytes: signedTx})
			}

			if err != nil {
				utils.Error("Failed to send sole transaction: %v\n", err)
			}
			if res.Error != nil {
				utils.Error("Server returned error: %v\n", res.Error)
			}
			result := &rpc.SendSoleRawTransactionResult{}
			err = res.GetObject(result)
			if err != nil {
				utils.Error("Failed to parse server response: %v\n", err)
			}
			formatted, err := json.MarshalIndent(result, "", "    ")
			if err != nil {
				utils.Error("Failed to parse server response: %v\n", err)
			}
			fmt.Printf("Successfully send sole transaction to node %v:\n%s\n", idx+1, formatted)
		}
		wg.Add(1)
		go f(remoteRPCEndpoint)
	}
	wg.Wait()
}

func doAutoSendCmd(cmd *cobra.Command, args []string) {
	remoteRPCEndpoints := []string{"http://127.0.0.1:16930/rpc", "http://127.0.0.1:16910/rpc", "http://127.0.0.1:16920/rpc", "http://127.0.0.1:16900/rpc"}

	var wg sync.WaitGroup

	wallet, fromAddress, err := walletUnlockWithPath(cmd, fromFlag, pathFlag, passwordFlag)
	if err != nil || wallet == nil {
		return
	}
	defer wallet.Lock(fromAddress)
	tfuel, ok := types.ParseCoinAmount("1")
	if !ok {
		utils.Error("Failed to parse tfuel amount")
	}
	fee, ok := types.ParseCoinAmount(feeFlag)
	if !ok {
		utils.Error("Failed to parse fee")
	}
	outputs := []types.TxOutput{{
		Address: common.HexToAddress(toFlag),
		Coins: types.Coins{
			TFuelWei: tfuel,
			ThetaWei: new(big.Int).SetUint64(0),
		},
	}}
	for i := 1; i <= 10; i++ {
		inputs := []types.TxInput{{
			Address: fromAddress,
			Coins: types.Coins{
				TFuelWei: new(big.Int).Add(tfuel, fee),
				ThetaWei: new(big.Int).SetUint64(0),
			},
			Sequence: uint64(i),
		}}
		sendTx := &types.SendTx{
			Fee: types.Coins{
				ThetaWei: new(big.Int).SetUint64(0),
				TFuelWei: fee,
			},
			Inputs:  inputs,
			Outputs: outputs,
		}
		sig, err := wallet.Sign(fromAddress, sendTx.SignBytes(chainIDFlag))
		if err != nil {
			utils.Error("Failed to sign transaction: %v\n", err)
		}
		sendTx.SetSignature(fromAddress, sig)
		raw, err := stypes.TxToBytes(sendTx)
		if err != nil {
			utils.Error("Failed to encode transaction: %v\n", err)
		}
		signedTx := hex.EncodeToString(raw)
		for idx, remoteRPCEndpoint := range remoteRPCEndpoints {
			f := func(remoteRPCEndpoint string, index int) {
				defer wg.Done()
				if remoteRPCEndpoint == "http://127.0.0.1:16900/rpc" {
					time.Sleep(time.Duration(1) * time.Second)
				}
				client := rpcc.NewRPCClient(remoteRPCEndpoint)

				var res *jsonrpc.RPCResponse
				res, err = client.Call("theta.SendSoleRawTransaction", rpc.SendSoleRawTransactionArgs{TxBytes: signedTx})

				if err != nil {
					utils.Error("Failed to send sole transaction: %v\n", err)
				}
				if res.Error != nil {
					utils.Error("Server returned error: %v\n", res.Error)
				}
				result := &rpc.SendSoleRawTransactionResult{}
				err = res.GetObject(result)
				if err != nil {
					utils.Error("Failed to parse server response: %v\n", err)
				}
				formatted, err := json.MarshalIndent(result, "", "    ")
				if err != nil {
					utils.Error("Failed to parse server response: %v\n", err)
				}
				fmt.Printf("Successfully send sole transaction #%v to node %v\n", i, index+1)
				if index == 0 {
					fmt.Printf("%s", formatted)
				}
			}
			wg.Add(1)
			// time.Sleep(time.Duration(1000) * time.Millisecond)
			go f(remoteRPCEndpoint, idx)
		}
		wg.Wait()
	}
}

func doAutoQueueSendCmd(cmd *cobra.Command, args []string) {

	remoteRPCEndpoints := []string{"http://127.0.0.1:16930/rpc", "http://127.0.0.1:16910/rpc", "http://127.0.0.1:16920/rpc", "http://127.0.0.1:16900/rpc"}
	// remoteRPCEndpoints := []string{"http://10.10.1.2:16900/rpc", "http://10.10.1.3:16900/rpc", "http://10.10.1.4:16900/rpc", "http://10.10.1.5:16900/rpc"}

	var wg sync.WaitGroup

	wallet, fromAddress, err := walletUnlockWithPath(cmd, fromFlag, pathFlag, passwordFlag)
	if err != nil || wallet == nil {
		return
	}
	defer wallet.Lock(fromAddress)
	tfuel, ok := types.ParseCoinAmount("1")
	if !ok {
		utils.Error("Failed to parse tfuel amount")
	}
	fee, ok := types.ParseCoinAmount(feeFlag)
	if !ok {
		utils.Error("Failed to parse fee")
	}
	outputs := []types.TxOutput{{
		Address: common.HexToAddress(toFlag),
		Coins: types.Coins{
			TFuelWei: tfuel,
			ThetaWei: new(big.Int).SetUint64(0),
		},
	}}
	txToSend := []string{}
	// s1 := time.Now()
	for i := 1; i <= int(sendTotalNumFlag); i++ {
		inputs := []types.TxInput{{
			Address: fromAddress,
			Coins: types.Coins{
				TFuelWei: new(big.Int).Add(tfuel, fee),
				ThetaWei: new(big.Int).SetUint64(0),
			},
			Sequence: uint64(i),
		}}
		sendTx := &types.SendTx{
			Fee: types.Coins{
				ThetaWei: new(big.Int).SetUint64(0),
				TFuelWei: fee,
			},
			Inputs:  inputs,
			Outputs: outputs,
		}
		sig, err := wallet.Sign(fromAddress, sendTx.SignBytes(chainIDFlag))
		if err != nil {
			utils.Error("Failed to sign transaction: %v\n", err)
		}
		sendTx.SetSignature(fromAddress, sig)
		raw, err := stypes.TxToBytes(sendTx)
		if err != nil {
			fmt.Printf("Failed to encode transaction: %v\n", err)
		}
		// fmt.Println("raw Tx size: ", len(raw))
		// hash := crypto.Keccak256Hash(raw)
		signedTx := hex.EncodeToString(raw)
		txToSend = append(txToSend, signedTx)
	}
	// fmt.Println("total time", time.Since(s1))

	for idx, remoteRPCEndpoint := range remoteRPCEndpoints {
		f := func(remoteRPCEndpoint string, index int) {
			defer wg.Done()
			// if remoteRPCEndpoint == "http://127.0.0.1:16900/rpc" {
			// 	time.Sleep(time.Duration(300) * time.Millisecond)
			// }

			client := rpcc.NewRPCClient(remoteRPCEndpoint)
			for tx_idx, signedTx := range txToSend {
				var res *jsonrpc.RPCResponse
				res, err = client.Call("theta.SendSoleRawTransaction", rpc.SendSoleRawTransactionArgs{TxBytes: signedTx})

				if err != nil {
					utils.Error("Failed to send sole transaction: %v\n", err)
				}
				if res.Error != nil {
					utils.Error("Server returned error: %v\n", res.Error)
				}
				if remoteRPCEndpoint == "http://10.10.1.2:16900/rpc" || remoteRPCEndpoint == "http://10.10.1.2:16920/rpc" {
					time.Sleep(time.Duration(15) * time.Millisecond)
				} else if remoteRPCEndpoint == "http://10.10.1.3:16900/rpc" || remoteRPCEndpoint == "http://10.10.1.2:16910/rpc" {
					time.Sleep(time.Duration(55) * time.Millisecond)
				} else if remoteRPCEndpoint == "http://10.10.1.4:16900/rpc" || remoteRPCEndpoint == "http://10.10.1.2:16900/rpc" {
					time.Sleep(time.Duration(155) * time.Millisecond)
					// time.Sleep(time.Duration(100) * time.Millisecond)
				} else {
					time.Sleep(time.Duration(296) * time.Millisecond)
					// time.Sleep(time.Duration(156) * time.Millisecond)
				}
				_ = &rpc.SendSoleRawTransactionResult{}
				/*
					result := &rpc.SendSoleRawTransactionResult{}
					err = res.GetObject(result)
					if err != nil {
						utils.Error("Failed to parse server response: %v\n", err)
					}
					// formatted, err := json.MarshalIndent(result, "", "    ")
					// if err != nil {
					// 	utils.Error("Failed to parse server response: %v\n", err)
					// }
					if (tx_idx+1)%10 == 0 {
						fmt.Printf("Successfully send sole transaction #%v to node %v\n", tx_idx, index+1)
					}

					// if index == 0 {
					// 	fmt.Printf("%s", formatted)
					// }
				*/
				if (tx_idx+1)%10 == 0 {
					fmt.Printf("Successfully send sole transaction #%v to node %v\n", tx_idx, index+1)
				}
				time.Sleep(time.Duration(sendIntervalMsFlag) * time.Millisecond)
			}
		}
		wg.Add(1)
		// time.Sleep(time.Duration(1000) * time.Millisecond)
		go f(remoteRPCEndpoint, idx)
	}
	wg.Wait()
}

func doAutoQueueMultiSendCmd(cmd *cobra.Command, args []string) {

	remoteRPCEndpoints := []string{"http://127.0.0.1:16930/rpc", "http://127.0.0.1:16910/rpc", "http://127.0.0.1:16920/rpc", "http://127.0.0.1:16900/rpc"}
	// remoteRPCEndpoints := []string{"http://10.10.1.2:16900/rpc", "http://10.10.1.3:16900/rpc", "http://10.10.1.4:16900/rpc", "http://10.10.1.5:16900/rpc"}

	var wg sync.WaitGroup

	wallet, fromAddress, err := walletUnlockWithPath(cmd, "0x2E833968E5bB786Ae419c4d13189fB081Cc43bab", pathFlag, "qwertyuiop")
	if err != nil || wallet == nil {
		return
	}
	defer wallet.Lock(fromAddress)
	tfuel, ok := types.ParseCoinAmount("1")
	if !ok {
		utils.Error("Failed to parse tfuel amount")
	}
	fee, ok := types.ParseCoinAmount(feeFlag)
	if !ok {
		utils.Error("Failed to parse fee")
	}
	outputs := []types.TxOutput{{
		Address: common.HexToAddress("0x19E7E376E7C213B7E7e7e46cc70A5dD086DAff2A"),
		Coins: types.Coins{
			TFuelWei: tfuel,
			ThetaWei: new(big.Int).SetUint64(0),
		},
	}}

	realTxToSend := [][]string{}
	// s1 := time.Now()
	nonce := 0
	for j := 0; j < sendRoundFlag; j++ {
		txToSend := []string{}
		for i := 0; i < sendNumPerRound; i++ {
			inputs := []types.TxInput{{
				Address: fromAddress,
				Coins: types.Coins{
					TFuelWei: new(big.Int).Add(tfuel, fee),
					ThetaWei: new(big.Int).SetUint64(0),
				},
				Sequence: uint64(nonce),
			}}
			nonce++
			sendTx := &types.SendTx{
				Fee: types.Coins{
					ThetaWei: new(big.Int).SetUint64(0),
					TFuelWei: fee,
				},
				Inputs:  inputs,
				Outputs: outputs,
			}
			sig, err := wallet.Sign(fromAddress, sendTx.SignBytes("tsub360777"))
			if err != nil {
				utils.Error("Failed to sign transaction: %v\n", err)
			}
			sendTx.SetSignature(fromAddress, sig)
			raw, err := stypes.TxToBytes(sendTx)
			if err != nil {
				fmt.Printf("Failed to encode transaction: %v\n", err)
			}
			// fmt.Println("raw Tx size: ", len(raw))
			// hash := crypto.Keccak256Hash(raw)
			signedTx := hex.EncodeToString(raw)

			// fmt.Println("tx raw", signedTx, ", tx hash", hash.Hex())
			txToSend = append(txToSend, signedTx)
		}
		realTxToSend = append(realTxToSend, txToSend)
		// fmt.Println(j, ":", txToSend)
		// fmt.Println("total time", time.Since(s1))
	}

	for idx, remoteRPCEndpoint := range remoteRPCEndpoints {
		f := func(remoteRPCEndpoint string, index int) {
			defer wg.Done()
			client := rpcc.NewRPCClient(remoteRPCEndpoint)
			for tx_idx, signedTx := range realTxToSend {
				if index == 1 {
					fmt.Println("send msg", tx_idx, "to node", remoteRPCEndpoint)
				}
				var res *jsonrpc.RPCResponse
				res, err = client.Call("theta.SendSoleMultiRawTransaction", rpc.SendSoleMultiRawTransactionArgs{TxBytes: signedTx})
				if err != nil {
					utils.Error("Failed to send sole transaction: %v\n", err)
				}
				if res.Error != nil {
					utils.Error("Server returned error: %v\n", res.Error)
				}
				if remoteRPCEndpoint == "http://10.10.1.2:16900/rpc" || remoteRPCEndpoint == "http://10.10.1.2:16920/rpc" {
					time.Sleep(time.Duration(15) * time.Millisecond)
				} else if remoteRPCEndpoint == "http://10.10.1.3:16900/rpc" || remoteRPCEndpoint == "http://10.10.1.2:16910/rpc" {
					time.Sleep(time.Duration(55) * time.Millisecond)
				} else if remoteRPCEndpoint == "http://10.10.1.4:16900/rpc" || remoteRPCEndpoint == "http://10.10.1.2:16900/rpc" {
					time.Sleep(time.Duration(155) * time.Millisecond)
					// time.Sleep(time.Duration(100) * time.Millisecond)
				} else {
					time.Sleep(time.Duration(296) * time.Millisecond)
					// time.Sleep(time.Duration(156) * time.Millisecond)
				}
				fmt.Printf("Successfully send sole batch transactions #%v to node %v\n", tx_idx, index+1)
				time.Sleep(time.Duration(sendIntervalMsFlag) * time.Millisecond)
			}
		}
		wg.Add(1)
		go f(remoteRPCEndpoint, idx)
	}
	wg.Wait()
}

/* broadcast version
func doSendCmd(cmd *cobra.Command, args []string) {
	walletType := getWalletType(cmd)
	if walletType == wtypes.WalletTypeSoft && len(fromFlag) == 0 {
		utils.Error("The from address cannot be empty") // we don't need to specify the "from address" for hardware wallets
		return
	}

	if len(toFlag) == 0 {
		utils.Error("The to address cannot be empty")
		return
	}
	if fromFlag == toFlag {
		utils.Error("The from and to address cannot be identical")
		return
	}

	wallet, fromAddress, err := walletUnlockWithPath(cmd, fromFlag, pathFlag, passwordFlag)
	if err != nil || wallet == nil {
		return
	}
	defer wallet.Lock(fromAddress)

	tfuel, ok := types.ParseCoinAmount(tfuelAmountFlag)
	if !ok {
		utils.Error("Failed to parse tfuel amount")
	}
	fee, ok := types.ParseCoinAmount(feeFlag)
	if !ok {
		utils.Error("Failed to parse fee")
	}
	inputs := []types.TxInput{{
		Address: fromAddress,
		Coins: types.Coins{
			TFuelWei: new(big.Int).Add(tfuel, fee),
			ThetaWei: new(big.Int).SetUint64(0),
		},
		Sequence: uint64(seqFlag),
	}}
	outputs := []types.TxOutput{{
		Address: common.HexToAddress(toFlag),
		Coins: types.Coins{
			TFuelWei: tfuel,
			ThetaWei: new(big.Int).SetUint64(0),
		},
	}}
	sendTx := &types.SendTx{
		Fee: types.Coins{
			ThetaWei: new(big.Int).SetUint64(0),
			TFuelWei: fee,
		},
		Inputs:  inputs,
		Outputs: outputs,
	}

	sig, err := wallet.Sign(fromAddress, sendTx.SignBytes(chainIDFlag))
	if err != nil {
		utils.Error("Failed to sign transaction: %v\n", err)
	}
	sendTx.SetSignature(fromAddress, sig)

	raw, err := stypes.TxToBytes(sendTx)
	if err != nil {
		utils.Error("Failed to encode transaction: %v\n", err)
	}
	signedTx := hex.EncodeToString(raw)

	client := rpcc.NewRPCClient(viper.GetString(utils.CfgRemoteRPCEndpoint))

	var res *jsonrpc.RPCResponse
	if asyncFlag {
		res, err = client.Call("theta.BroadcastRawTransactionAsync", rpc.BroadcastRawTransactionArgs{TxBytes: signedTx})
	} else {
		// res, err = client.Call("theta.BroadcastRawTransaction", rpc.BroadcastRawTransactionArgs{TxBytes: signedTx})
		res, err = client.Call("theta.SendSoleRawTransaction", rpc.SendSoleRawTransactionArgs{TxBytes: signedTx})
	}

	if err != nil {
		utils.Error("Failed to broadcast transaction: %v\n", err)
	}
	if res.Error != nil {
		utils.Error("Server returned error: %v\n", res.Error)
	}
	result := &rpc.BroadcastRawTransactionResult{}
	err = res.GetObject(result)
	if err != nil {
		utils.Error("Failed to parse server response: %v\n", err)
	}
	formatted, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		utils.Error("Failed to parse server response: %v\n", err)
	}
	fmt.Printf("Successfully broadcasted transaction:\n%s\n", formatted)
}
*/

func init() {
	sendCmd.Flags().StringVar(&chainIDFlag, "chain", "", "Chain ID")
	sendCmd.Flags().StringVar(&fromFlag, "from", "", "Address to send from")
	sendCmd.Flags().StringVar(&toFlag, "to", "", "Address to send to")
	sendCmd.Flags().StringVar(&pathFlag, "path", "", "Wallet derivation path")
	sendCmd.Flags().Uint64Var(&seqFlag, "seq", 0, "Sequence number of the transaction")
	sendCmd.Flags().StringVar(&tfuelAmountFlag, "tfuel", "0", "TFuel amount")
	sendCmd.Flags().StringVar(&feeFlag, "fee", fmt.Sprintf("%dwei", types.MinimumTransactionFeeTFuelWeiJune2021), "Fee")
	sendCmd.Flags().StringVar(&walletFlag, "wallet", "soft", "Wallet type (soft|nano|trezor)")
	sendCmd.Flags().BoolVar(&asyncFlag, "async", false, "block until tx has been included in the blockchain")
	sendCmd.Flags().StringVar(&passwordFlag, "password", "", "password to unlock the wallet")

	sendCmd.MarkFlagRequired("chain")
	//sendCmd.MarkFlagRequired("from")
	sendCmd.MarkFlagRequired("to")
	sendCmd.MarkFlagRequired("seq")

	autoSendCmd.Flags().StringVar(&chainIDFlag, "chain", "", "Chain ID")
	autoSendCmd.Flags().StringVar(&fromFlag, "from", "", "Address to send from")
	autoSendCmd.Flags().StringVar(&toFlag, "to", "", "Address to send to")
	autoSendCmd.Flags().StringVar(&pathFlag, "path", "", "Wallet derivation path")
	autoSendCmd.Flags().Uint64Var(&seqFlag, "seq", 0, "Sequence number of the transaction")
	autoSendCmd.Flags().StringVar(&tfuelAmountFlag, "tfuel", "0", "TFuel amount")
	autoSendCmd.Flags().StringVar(&feeFlag, "fee", fmt.Sprintf("%dwei", types.MinimumTransactionFeeTFuelWeiJune2021), "Fee")
	autoSendCmd.Flags().StringVar(&walletFlag, "wallet", "soft", "Wallet type (soft|nano|trezor)")
	autoSendCmd.Flags().BoolVar(&asyncFlag, "async", false, "block until tx has been included in the blockchain")
	autoSendCmd.Flags().StringVar(&passwordFlag, "password", "", "password to unlock the wallet")

	autoQueueSendCmd.Flags().StringVar(&chainIDFlag, "chain", "", "Chain ID")
	autoQueueSendCmd.Flags().StringVar(&fromFlag, "from", "", "Address to send from")
	autoQueueSendCmd.Flags().StringVar(&toFlag, "to", "", "Address to send to")
	autoQueueSendCmd.Flags().StringVar(&pathFlag, "path", "", "Wallet derivation path")
	autoQueueSendCmd.Flags().Uint64Var(&seqFlag, "seq", 0, "Sequence number of the transaction")
	autoQueueSendCmd.Flags().StringVar(&tfuelAmountFlag, "tfuel", "0", "TFuel amount")
	autoQueueSendCmd.Flags().StringVar(&feeFlag, "fee", fmt.Sprintf("%dwei", types.MinimumTransactionFeeTFuelWeiJune2021), "Fee")
	autoQueueSendCmd.Flags().StringVar(&walletFlag, "wallet", "soft", "Wallet type (soft|nano|trezor)")
	autoQueueSendCmd.Flags().BoolVar(&asyncFlag, "async", false, "block until tx has been included in the blockchain")
	autoQueueSendCmd.Flags().StringVar(&passwordFlag, "password", "", "password to unlock the wallet")
	autoQueueSendCmd.Flags().Uint64Var(&sendTotalNumFlag, "num", 1000, "total number of message")
	autoQueueSendCmd.Flags().Uint64Var(&sendIntervalMsFlag, "interval", 1000, "sending rate")

	autoMultiQueueSendCmd.Flags().StringVar(&chainIDFlag, "chain", "", "Chain ID")
	autoMultiQueueSendCmd.Flags().StringVar(&fromFlag, "from", "", "Address to send from")
	autoMultiQueueSendCmd.Flags().StringVar(&toFlag, "to", "", "Address to send to")
	autoMultiQueueSendCmd.Flags().StringVar(&pathFlag, "path", "", "Wallet derivation path")
	autoMultiQueueSendCmd.Flags().Uint64Var(&seqFlag, "seq", 0, "Sequence number of the transaction")
	autoMultiQueueSendCmd.Flags().StringVar(&tfuelAmountFlag, "tfuel", "0", "TFuel amount")
	autoMultiQueueSendCmd.Flags().StringVar(&feeFlag, "fee", fmt.Sprintf("%dwei", types.MinimumTransactionFeeTFuelWeiJune2021), "Fee")
	autoMultiQueueSendCmd.Flags().StringVar(&walletFlag, "wallet", "soft", "Wallet type (soft|nano|trezor)")
	autoMultiQueueSendCmd.Flags().BoolVar(&asyncFlag, "async", false, "block until tx has been included in the blockchain")
	autoMultiQueueSendCmd.Flags().StringVar(&passwordFlag, "password", "", "password to unlock the wallet")
	autoMultiQueueSendCmd.Flags().Uint64Var(&sendTotalNumFlag, "num", 1000, "total number of message")
	autoMultiQueueSendCmd.Flags().Uint64Var(&sendIntervalMsFlag, "interval", 1000, "sending rate")
	autoMultiQueueSendCmd.Flags().IntVar(&sendNumPerRound, "txnum", 1, "txs in one round")
	autoMultiQueueSendCmd.Flags().IntVar(&sendRoundFlag, "round", 1, "number of rounds")
}
