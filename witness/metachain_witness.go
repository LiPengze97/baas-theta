package witness

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	// "github.com/thetatoken/theta/crypto"

	scom "github.com/thetatoken/thetasubchain/common"
	scta "github.com/thetatoken/thetasubchain/contracts/accessors"
	score "github.com/thetatoken/thetasubchain/core"

	// "github.com/thetatoken/thetasubchain/eth/abi/bind"
	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/store"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

// SubchainRegistrarSendToSubchainEvent
var logger *log.Entry = log.WithFields(log.Fields{"prefix": "witness"})

type MetachainWitness struct {
	updateTicker   *time.Ticker
	updateInterval int

	// The main chain
	mainchainID                 *big.Int
	mainchainEthRpcUrl          string
	mainchainEthRpcClient       *ec.Client
	witnessedDynasty            *big.Int
	subchainRegistrar           *scta.SubchainRegistrar // the SubchainRegistrar contract deployed on the main chain
	mainchainTFuelTokenBankAddr common.Address
	mainchainTFuelTokenBank     *scta.MainchainTFuelTokenBank // the MainchainTFuelTokenBank contract deployed on the main chain
	mainchainTNT20TokenBankAddr common.Address
	mainchainTNT20TokenBank     *scta.MainchainTNT20TokenBank // the MainchainTNT20TokenBank contract deployed on the main chain
	mainchainBlockHeight        *big.Int
	lastQueryedMainChainHeight  *big.Int

	// The subchain
	subchainID                 *big.Int
	subchainEthRpcUrl          string
	subchainEthRpcClient       *ec.Client
	subchainTFuelTokenBankAddr common.Address
	subchainTFuelTokenBank     *scta.SubchainTFuelTokenBank // the SubchainTFuelTokenBank contract deployed on the subchain
	subchainTNT20TokenBankAddr common.Address
	subchainTNT20TokenBank     *scta.SubchainTNT20TokenBank // the SubchainTNT20TokenBank contract deployed on the subchain

	// Validator set
	validatorSetCache map[string]*score.ValidatorSet

	// Inter-chain messaging
	interChainEventCache *score.InterChainEventCache

	// Life cycle
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// NewMetachainWitness creates a new MetachainWitness
func NewMetachainWitness(updateInterval int, interChainEventCache *score.InterChainEventCache) *MetachainWitness {
	mainchainEthRpcURL := viper.GetString(scom.CfgMainchainEthRpcURL)
	mainchainEthRpcClient, err := ec.Dial(mainchainEthRpcURL)
	if err != nil {
		logger.Fatalf("the ETH client failed to connect to the mainchain ETH RPC %v\n", err)
	}
	mainchainID, err := mainchainEthRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Fatalf("failed to get the chainID of the main chain %v\n", err)
	}
	subchainRegistrarAddr := common.HexToAddress(viper.GetString(scom.CfgSubchainRegistrarContractAddress))
	subchainRegistrar, err := scta.NewSubchainRegistrar(subchainRegistrarAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create SubchainRegistrar contract %v\n", err)
	}
	mainchainTFuelTokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTFuelTokenBankContractAddress))
	mainchainTFuelTokenBank, err := scta.NewMainchainTFuelTokenBank(mainchainTFuelTokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTFuelTokenBank contract %v\n", err)
	}
	mainchainTNT20TokenBankAddr := common.HexToAddress(viper.GetString(scom.CfgMainchainTNT20TokenBankContractAddress))
	mainchainTNT20TokenBank, err := scta.NewMainchainTNT20TokenBank(mainchainTNT20TokenBankAddr, mainchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to create MainchainTNT20TokenBank contract %v\n", err)
	}

	subchainID := big.NewInt(viper.GetInt64(scom.CfgSubchainID))
	subchainEthRpcURL := viper.GetString(scom.CfgSubchainEthRpcURL)
	subchainEthRpcClient, err := ec.Dial(subchainEthRpcURL)
	if err != nil {
		logger.Fatalf("the ETH client failed to connect to the subchain ETH RPC %v\n", err)
	}

	mw := &MetachainWitness{
		updateInterval: updateInterval,

		mainchainID:                 mainchainID,
		mainchainEthRpcUrl:          mainchainEthRpcURL,
		mainchainEthRpcClient:       mainchainEthRpcClient,
		witnessedDynasty:            big.NewInt(0),
		subchainRegistrar:           subchainRegistrar,
		mainchainTFuelTokenBankAddr: mainchainTFuelTokenBankAddr,
		mainchainTFuelTokenBank:     mainchainTFuelTokenBank,
		mainchainTNT20TokenBankAddr: mainchainTNT20TokenBankAddr,
		mainchainTNT20TokenBank:     mainchainTNT20TokenBank,
		mainchainBlockHeight:        nil,
		lastQueryedMainChainHeight:  big.NewInt(0),

		subchainID:           subchainID,
		subchainEthRpcUrl:    subchainEthRpcURL,
		subchainEthRpcClient: subchainEthRpcClient,

		interChainEventCache: interChainEventCache,

		wg: &sync.WaitGroup{},
	}
	return mw
}

func (mw *MetachainWitness) Start(ctx context.Context) {
	c, cancel := context.WithCancel(ctx)
	mw.ctx = c
	mw.cancel = cancel

	mw.wg.Add(1)
	go mw.mainloop(ctx)
}

func (mw *MetachainWitness) Stop() {
	if mw.updateTicker != nil {
		mw.updateTicker.Stop()
	}
	mw.cancel()
}

func (mw *MetachainWitness) Wait() {
	mw.wg.Wait()
}

func (mw *MetachainWitness) SetSubchainTokenBanks(ledger score.Ledger) {
	subchainTFuelTokenBankAddr, err := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTFuel)
	if subchainTFuelTokenBankAddr == nil || err != nil {
		logger.Fatalf("failed to obtain SubchainTFuelTokenBank contract address: %v\n", err)
	}
	mw.subchainTFuelTokenBankAddr = *subchainTFuelTokenBankAddr
	mw.subchainTFuelTokenBank, err = scta.NewSubchainTFuelTokenBank(*subchainTFuelTokenBankAddr, mw.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTFuelTokenBank contract: %v\n", err)
	}

	subchainTNT20TokenBankAddr, err := ledger.GetTokenBankContractAddress(score.CrossChainTokenTypeTNT20)
	if subchainTNT20TokenBankAddr == nil || err != nil {
		logger.Fatalf("failed to obtain SubchainTNT20TokenBank contract address: %v\n", err)
	}
	mw.subchainTNT20TokenBankAddr = *subchainTNT20TokenBankAddr
	mw.subchainTNT20TokenBank, err = scta.NewSubchainTNT20TokenBank(*subchainTNT20TokenBankAddr, mw.subchainEthRpcClient)
	if err != nil {
		logger.Fatalf("failed to set the SubchainTNT20TokenBankAddr contract: %v\n", err)
	}

}

// TODO: make sure the block number returned by the client.BlockNumber() call is the lastest *finalized* block number
func (mw *MetachainWitness) GetMainchainBlockHeight() (*big.Int, error) {
	if mw.mainchainBlockHeight == nil {
		return nil, fmt.Errorf("mainchain block height not been updated yet")
	}
	return mw.mainchainBlockHeight, nil
}

func (mw *MetachainWitness) GetValidatorSetByDynasty(dynasty *big.Int) (*score.ValidatorSet, error) {
	validatorSet, ok := mw.validatorSetCache[dynasty.String()]
	if ok && validatorSet != nil && validatorSet.Dynasty() == dynasty {
		return validatorSet, nil
	}

	var err error
	validatorSet, err = mw.updateValidatorSetCache(dynasty) // cache lazy update
	if err != nil {
		return nil, err
	}

	return validatorSet, nil
}

func (mw *MetachainWitness) mainloop(ctx context.Context) {
	mw.updateTicker = time.NewTicker(time.Duration(mw.updateInterval) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			return
		case <-mw.updateTicker.C:
			mw.update()
		}
	}
}

func (mw *MetachainWitness) update() {
	// Mainchain
	mw.updateMainchainBlockHeight()
	dynasty := scom.CalculateDynasty(mw.mainchainBlockHeight)
	if mw.witnessedDynasty == nil || dynasty.Cmp(mw.witnessedDynasty) > 0 { // needs to update the cache
		mw.updateValidatorSetCache(dynasty)
		mw.witnessedDynasty = dynasty
	}
	mw.collectInterChainMessageEventsOnMainchain()

	// Subchain
	mw.collectInterChainMessageEventsOnSubchain()
}

func (mw *MetachainWitness) updateMainchainBlockHeight() {
	mbh, err := mw.mainchainEthRpcClient.BlockNumber(context.Background())
	if err != nil {
		logger.Warnf("failed to get the mainchain block height %v\n", err)
		return
	}
	mw.mainchainBlockHeight = big.NewInt(int64(mbh))
}

func (mw *MetachainWitness) collectInterChainMessageEventsOnMainchain() {
	mw.collectInterChainMessageEventsOnChain(mw.mainchainID, mw.mainchainEthRpcUrl, mw.mainchainTFuelTokenBankAddr, mw.mainchainTNT20TokenBankAddr)
}

func (mw *MetachainWitness) collectInterChainMessageEventsOnSubchain() {
	mw.collectInterChainMessageEventsOnChain(mw.subchainID, mw.subchainEthRpcUrl, mw.subchainTFuelTokenBankAddr, mw.subchainTNT20TokenBankAddr)
}

func (mw *MetachainWitness) collectInterChainMessageEventsOnChain(queriedChainID *big.Int, ethRpcUrl string,
	tfuelTokenBankAddr common.Address, tnt20TokenBankAddr common.Address) {
	fromBlock, err := mw.interChainEventCache.GetLastQueryedHeightForType(queriedChainID, score.IMCEventTypeCrossChainTokenLock)
	if err == store.ErrKeyNotFound {
		mw.interChainEventCache.SetLastQueryedHeightForType(queriedChainID, score.IMCEventTypeCrossChainTokenLock, common.Big0)
	} else if err != nil {
		logger.Warnf("failed to get the last queryed height %v\n", err)
	}
	toBlock := mw.calculateToBlock(fromBlock)
	logger.Infof("Query inter-chain message events from block height %v to %v on chain %v", queriedChainID.String(), fromBlock.String(), toBlock.String())

	queryTypes := append(score.LockTypes, score.UnlockTypes...)
	for _, imceType := range queryTypes {
		var events []*score.InterChainMessageEvent
		switch imceType {
		case score.IMCEventTypeCrossChainTokenLockTFuel, score.IMCEventTypeCrossChainTokenUnlockTFuel:
			events = score.QueryInterChainEventLog(queriedChainID, fromBlock, toBlock, tfuelTokenBankAddr, imceType, ethRpcUrl)
		case score.IMCEventTypeCrossChainTokenLockTNT20, score.IMCEventTypeCrossChainTokenUnlockTNT20:
			events = score.QueryInterChainEventLog(queriedChainID, fromBlock, toBlock, tnt20TokenBankAddr, imceType, ethRpcUrl)
		}
		if len(events) == 0 {
			continue
		}
		if imceType == score.IMCEventTypeCrossChainTokenUnlockTFuel || imceType == score.IMCEventTypeCrossChainTokenUnlockTNT20 || imceType == score.IMCEventTypeCrossChainTokenUnlockTNT721 {
			mw.updateVoucherBurnStatus(events)
		}
		err = mw.interChainEventCache.InsertList(events)
		if err != nil { // should not happen
			logger.Panicf("failed to insert events into cache")
		}
	}
	mw.interChainEventCache.SetLastQueryedHeightForType(queriedChainID, score.IMCEventTypeCrossChainTokenLock, toBlock)
}

func (mw *MetachainWitness) updateVoucherBurnStatus(events []*score.InterChainMessageEvent) {
	for _, e := range events {
		sourceChainID := e.SourceChainID
		statusExists, err := mw.interChainEventCache.VoucherBurnNonceExists(sourceChainID, e.Type, e.Nonce)
		if !statusExists && err == nil {
			break
		} else {
			// Should not happen. Since statusExists
			logger.Panic(err)
		}
		eventStatus, err := mw.interChainEventCache.GetVoucherBurnStatus(sourceChainID, e.Type, e.Nonce)
		if err == nil {
			// Should not happen. Since statusExists
			logger.Panic(err)
		}
		eventStatus.Status = score.VoucherBurnEventStatusFinalized
		mw.interChainEventCache.SetVoucherBurnStatus(sourceChainID, eventStatus)
	}
}

func (mw *MetachainWitness) calculateToBlock(fromBlock *big.Int) *big.Int {
	toBlock, err := mw.GetMainchainBlockHeight()
	if err != nil {
		return fromBlock
	}
	maxBlockRange := int64(1000) // block range query allows at most 5000 blocks, here we intentionally use a much smaller range to limit cpu/mem resource usage
	minBlockGap := int64(10)     // tentative, to ensure the chain has enough time to finalize the event
	if new(big.Int).Sub(toBlock, fromBlock).Cmp(big.NewInt(maxBlockRange)) > 0 {
		// catch-up phase, gap is over maxBlockRange， catch-up at full speed
		toBlock = new(big.Int).Add(fromBlock, big.NewInt(maxBlockRange))
	} else {
		// steady phase, gap is between minBlockGap and maxBlockRange
		toBlock = new(big.Int).Sub(toBlock, big.NewInt(minBlockGap))
	}
	return toBlock
}

func (mw *MetachainWitness) updateValidatorSetCache(dynasty *big.Int) (*score.ValidatorSet, error) {
	queryBlockHeight := big.NewInt(1).Mul(dynasty, big.NewInt(1).SetInt64(scom.NumMainchainBlocksPerDynasty))
	queryBlockHeight = big.NewInt(0).Add(queryBlockHeight, big.NewInt(1)) // increment by one to make sure the query block height falls into the dynasty
	vs, err := mw.subchainRegistrar.Getvalidatorset(nil, mw.subchainID, queryBlockHeight)
	validatorAddrs := vs.Validators
	validatorStakes := vs.Shareamounts

	if err != nil {
		return nil, err
	}

	if len(validatorAddrs) != len(validatorStakes) {
		return nil, fmt.Errorf("the length of validatorAddrs and validatorStakes are not equal")
	}

	validatorSet := score.NewValidatorSet(dynasty)
	for i := 0; i < len(validatorAddrs); i++ {
		validator := score.NewValidator(validatorAddrs[i].Hex(), validatorStakes[i])
		validatorSet.AddValidator(validator)
	}

	mw.validatorSetCache[dynasty.String()] = validatorSet

	return validatorSet, nil
}

func (mw *MetachainWitness) GetInterChainEventCache() *score.InterChainEventCache {
	return mw.interChainEventCache
}
