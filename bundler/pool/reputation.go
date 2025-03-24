package mempool

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
)

const (
	OK        = 0
	THROTTLED = 1
	BANNED    = 2
)

type ReputationParams struct {
	OpsSeen     int64 //valid UserOperations seen by bundler
	OpsIncluded int64
	Status      int64
	Blacklisted bool
}

type ReputationManager struct {
	mu         sync.RWMutex
	Reputation map[common.Address]ReputationParams //entity-> reputaion //for memory cached fast lookup
	db         *ethdb.Database                     //local db
}

func New(db *ethdb.Database) *ReputationManager {

	return &ReputationManager{
		Reputation: make(map[common.Address]ReputationParams),
		db:         db,
	}
}

func (rm *ReputationManager) UpdateOpsSeen(entity common.Address) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if _, ok := rm.Reputation[entity]; !ok {

	}

}
