package execution

import (
	"fmt"
	"math/big"

	"github.com/thetatoken/theta/common"
	"github.com/thetatoken/theta/common/result"
	"github.com/thetatoken/theta/ledger/types"
	"github.com/thetatoken/theta/store/database"

	sbc "github.com/thetatoken/thetasubchain/blockchain"
	score "github.com/thetatoken/thetasubchain/core"
	slst "github.com/thetatoken/thetasubchain/ledger/state"
	stypes "github.com/thetatoken/thetasubchain/ledger/types"
	"github.com/thetatoken/thetasubchain/witness"
)

var _ TxExecutor = (*SubchainValidatorSetUpdateTxExecutor)(nil)

// ------------------------------- SubchainValidatorSet Transaction -----------------------------------

// SubchainValidatorSetUpdateTxExecutor implements the TxExecutor interface
type SubchainValidatorSetUpdateTxExecutor struct {
	db               database.Database
	chain            *sbc.Chain
	state            *slst.LedgerState
	consensus        score.ConsensusEngine
	valMgr           score.ValidatorManager
	mainchainWitness witness.ChainWitness
}

// NewSubchainValidatorSetUpdateTxExecutor creates a new instance of SubchainValidatorSetUpdateTxExecutor
func NewSubchainValidatorSetUpdateTxExecutor(db database.Database, chain *sbc.Chain, state *slst.LedgerState, consensus score.ConsensusEngine,
	valMgr score.ValidatorManager, mainchainWitness witness.ChainWitness) *SubchainValidatorSetUpdateTxExecutor {
	return &SubchainValidatorSetUpdateTxExecutor{
		db:               db,
		chain:            chain,
		state:            state,
		consensus:        consensus,
		valMgr:           valMgr,
		mainchainWitness: mainchainWitness,
	}
}

func (exec *SubchainValidatorSetUpdateTxExecutor) sanityCheck(chainID string, view *slst.StoreView, transaction types.Tx) result.Result {
	tx := transaction.(*stypes.SubchainValidatorSetUpdateTx)
	validatorSet := getValidatorSet(exec.consensus.GetLedger(), exec.valMgr)
	validatorAddresses := getValidatorAddresses(validatorSet)

	// Validate proposer, basic
	res := tx.Proposer.ValidateBasic()
	if res.IsError() {
		return res
	}

	// verify that at most one subchain validator set update transaction is processed for each block
	if view.SubchainValidatorSetTransactionProcessed() {
		return result.Error("Another subchain validator set update transaction has been processed for the current block")
	}

	// verify the proposer is one of the validators
	res = isAValidator(tx.Proposer.Address, validatorAddresses)
	if res.IsError() {
		return res
	}

	proposerAccount, res := getOrMakeInput(view, tx.Proposer)
	if res.IsError() {
		return res
	}

	// verify the proposer's signature
	signBytes := tx.SignBytes(chainID)
	if !tx.Proposer.Signature.Verify(signBytes, proposerAccount.Address) {
		return result.Error("SignBytes: %X", signBytes)
	}

	return result.OK
}

func (exec *SubchainValidatorSetUpdateTxExecutor) process(chainID string, view *slst.StoreView, transaction types.Tx) (common.Hash, result.Result) {
	tx := transaction.(*stypes.SubchainValidatorSetUpdateTx)

	if view.SubchainValidatorSetTransactionProcessed() {
		return common.Hash{}, result.Error("Another subchain validator set update transaction has been processed for the current block")
	}

	// new validator set and dynasty from the transaction
	newDynasty := tx.Dynasty
	newValidatorSet := score.NewValidatorSet(newDynasty)
	newValidatorSet.SetValidators(tx.Validators)

	currentDynasty := view.GetDynasty()
	if newDynasty.Cmp(currentDynasty) <= 0 {
		// This must be an invalid block, since dynasty needs to strictly increase. Without this check,
		// a malicious proposer may attempt to install a validator set of a previous dynasty.
		return common.Hash{}, result.Error(fmt.Sprintf("new dynasty needs to be strictly larger than the current dynasty (new: %v, current: %v)", newDynasty, currentDynasty))
	}

	witnessedValidatorSet, err := exec.mainchainWitness.GetValidatorSetByDynasty(newDynasty)
	if err != nil {
		return common.Hash{}, result.UndecidedWith(result.Info{"newDynasty": newDynasty, "err": err})
	}

	if !newValidatorSet.Equals(witnessedValidatorSet) {
		return common.Hash{}, result.Error("validator set mismatch")
	}

	// update the dynasty and the subchain validator set
	view.UpdateValidatorSet(newValidatorSet)
	view.SetSubchainValidatorSetTransactionProcessed(true)
	txHash := types.TxID(chainID, tx)
	return txHash, result.OK
}

func (exec *SubchainValidatorSetUpdateTxExecutor) getTxInfo(transaction types.Tx) *score.TxInfo {
	return &score.TxInfo{
		EffectiveGasPrice: exec.calculateEffectiveGasPrice(transaction),
	}
}

func (exec *SubchainValidatorSetUpdateTxExecutor) calculateEffectiveGasPrice(transaction types.Tx) *big.Int {
	return new(big.Int).SetUint64(0)
}