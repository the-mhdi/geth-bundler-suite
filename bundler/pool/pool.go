package pool

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/the-mhdi/geth-bundler-suite/types"
)

// https://eips.ethereum.org/EIPS/eip-7562#validation-rules-types
type Config struct {
	PoolType                 string           `json:"PoolType"`                 //public - private
	Locals                   []common.Address `json:"Locals"`                   // Addresses that should be treated by default as local
	NoLocals                 bool             `json:"NoLocals"`                 // Whether local transaction handling should be disabled
	Lifetime                 time.Duration    `json:"Lifetime"`                 // Maximum amount of time transaction are queued
	MaxPendingUserOperations uint64           `json:"maxPendingUserOperations"` // Maximum number of pending UserOperations in the mempool.
}

type Pool struct {
	conf                  Config
	pendingUserOperations map[common.Hash]*types.UserOperation // Pool for pending UserOperations, using hash as key for efficient lookup.
	pendingMutex          sync.RWMutex                         // Mutex to protect concurrent access to pendingUserOperations.

	// TODO: Add Queues for incoming UserOperations (e.g., incoming queue, validation queue, etc.) if needed for more complex processing pipeline.
	// TODO: Add 'Ready' UserOperation Pool for UserOperations that are ready to be included in bundles.

	// TODO: Implement persistence if required (e.g., saving mempool to disk and loading on startup).
}

func New(config Config) *Pool {
	return &Pool{
		conf:                  config,
		pendingUserOperations: make(map[common.Hash]*types.UserOperation), // Initialize the pending UserOperations map.
		pendingMutex:          sync.RWMutex{},                             // Initialize the mutex.
	}
}

func (m *Pool) AddUserOperation(userOp *types.UserOperation) error {
	m.pendingMutex.Lock() // Acquire write lock for pendingUserOperations.
	defer m.pendingMutex.Unlock()

	ValidationPhase(userOp)
}

func (pool *Pool) loop() {}

func ValidationPhase(userOp UserOperation) (bool, error) {}

func ExecutionPhase() {}
