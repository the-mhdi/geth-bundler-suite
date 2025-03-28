package mempool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/event"
	"github.com/the-mhdi/geth-bundler-suite/bundler/types"
)

type Pool interface {
	// Has returns an indicator whether txpool has a transaction
	// cached with the given hash.
	Has(hash common.Hash) bool

	// Get retrieves the transaction from local txpool with given
	// tx hash.
	Get(hash common.Hash) *types.Transaction

	// GetRLP retrieves the RLP-encoded transaction from local txpool
	// with given tx hash.
	GetRLP(hash common.Hash) []byte

	// Add should add the given transactions to the pool.
	Add(txs []*types.Transaction, sync bool) []error

	// Pending should return pending transactions.
	// The slice should be modifiable by the caller.
	Pending(filter txpool.PendingFilter) map[common.Address][]*txpool.LazyTransaction

	// SubscribeTransactions subscribes to new transaction events. The subscriber
	// can decide whether to receive notifications only for newly seen transactions
	// or also for reorged out ones.
	SubscribeTransactions(ch chan<- core.NewTxsEvent, reorgs bool) event.Subscription
}

// memPool implements Pool interface, it's the eth.txPool of geth client as specified in its source code

type memPool struct {
	Reputation *ReputationManager
	Validation *ValidationManager
}

func (p *memPool) SubmitTopool(up *types.UserOperation) error {

}
