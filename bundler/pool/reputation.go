package mempool

import (
	"bytes"
	"encoding/gob"
	"sync"
	"unsafe"

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
	Reputation map[common.Address]*ReputationParams //entity-> reputaion //for memory cached fast lookup
	db         ethdb.Database                       //local db
	mu         sync.RWMutex
}

func New(db *ethdb.Database) *ReputationManager {

	return &ReputationManager{
		Reputation: make(map[common.Address]*ReputationParams),
		db:         db,
	}
}

func (rm *ReputationManager) UpdateOpsSeen(entity common.Address) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	value, ok := rm.Reputation[entity]    //entity already cached in memory
	has, err := rm.db.Has(entity.Bytes()) //retrieves the entity from persistent database

	if err != nil {
		return err
	}

	if ok {
		//if entity reputation is already loaded in the memory
		value.OpsSeen++ //increment

		err := rm.db.Put(entity.Bytes(), value.Bytes()) //update db

		if err != nil {
			return err
		}
	}

	if has && !ok {
		//entity exists in the db but not loaded into memory
		rp, err := rm.db.Get(entity.Bytes())
		rm.Reputation[entity] = new(ReputationParams)

	}

	if !ok && !has {
		//first time the entity being seen
		rm.Reputation[entity] = new(ReputationParams)

		err := rm.db.Put(entity.Bytes(), new(ReputationParams).Bytes())

		if err != nil {
			return err
		}

	}

}

func (rp *ReputationParams) Bytes() []byte {
	buffer := bytes.NewBuffer(make([]byte, 0, unsafe.Sizeof(rp)))
	enc := gob.NewEncoder(buffer)
	enc.Encode(rp)

	return buffer.Bytes()

}

func decode(rp []byte) *ReputationParams {
	buffer := bytes.NewBuffer(rp)

	dec := gob.NewDecoder(buffer)
	rpStruct := new(ReputationParams)
	dec.Decode(rpStruct)

	return rpStruct
}
