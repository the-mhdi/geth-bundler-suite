package mempool

import (
	"bytes"
	"encoding/gob"
	"sync"
	"time"
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
	LastUpdate  time.Time
}

type ReputationManager struct {
	Reputation map[common.Address]*ReputationParams //entity-> reputaion //for memory cached fast lookup
	db         ethdb.Database                       //local db
	mu         sync.RWMutex
}

func New(db ethdb.Database) *ReputationManager {

	return &ReputationManager{
		Reputation: make(map[common.Address]*ReputationParams),
		db:         db,
	}
}

func (rm *ReputationManager) UpdateOpsSeen(entity common.Address) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rp, err := rm.getEntity(entity)

	if err != nil {
		return err
	}

	rp.OpsSeen++
	rp.LastUpdate = time.Now()

	err = rm.db.Put(entity.Bytes(), rp.Bytes())

	if err != nil {
		return err
	}

	return nil
}

func (rm *ReputationManager) UpdateOpsIncluded(entity common.Address) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rp, err := rm.getEntity(entity)

	if err != nil {
		return err
	}

	rp.OpsIncluded++
	rp.LastUpdate = time.Now()

	err = rm.db.Put(entity.Bytes(), rp.Bytes())

	if err != nil {
		return err
	}

	return nil
}

func (rm *ReputationManager) getEntity(entity common.Address) (*ReputationParams, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	value, ok := rm.Reputation[entity] //entity already cached in memory

	if ok {
		//if entity reputation is already loaded in the memory
		return value, nil
	}

	has, err := rm.db.Has(entity.Bytes()) //retrieves the entity from persistent database
	if err != nil {
		return nil, err
	}

	if has {
		//entity exists in the db but not loaded into memory
		rp, err := rm.db.Get(entity.Bytes())
		if err != nil {
			return nil, err
		}

		return decode(rp), nil
	}

	rm.Reputation[entity] = new(ReputationParams)

	err = rm.db.Put(entity.Bytes(), rm.Reputation[entity].Bytes())

	if err != nil {
		return nil, err
	}

	return rm.Reputation[entity], nil

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
