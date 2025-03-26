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
const (
	anHour int64 = 3600
)

type ReputationParams struct {
	OpsSeen      int64 //valid UserOperations seen by bundler
	OpsIncluded  int64
	Status       int64
	Blacklisted  bool
	LastSeen     time.Time
	LastIncluded time.Time
}

type ReputationManager struct {
	Reputation map[common.Address]*ReputationParams //entity-> reputaion //for memory cached fast lookup
	db         ethdb.KeyValueStore                  //ethdb.Database                       //local db
	mu         sync.RWMutex
}

func New(db ethdb.KeyValueStore) *ReputationManager {

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

	rp.refresh()
	rp.OpsSeen++
	rp.LastSeen = time.Now()

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

	rp.refresh()
	rp.OpsIncluded++
	rp.LastIncluded = time.Now()

	err = rm.db.Put(entity.Bytes(), rp.Bytes())

	if err != nil {
		return err
	}

	return nil
}

func (rm *ReputationManager) GetStatus(entity common.Address) (int64, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rp, err := rm.getEntity(entity)

	if err != nil {
		return rp.Status, err
	}

	max_seen := rp.OpsSeen / MIN_INCLUSION_RATE_DENOMINATOR

	if max_seen > (rp.OpsIncluded + BAN_SLACK) {

		rp.Status = BANNED

		return rp.Status, nil
	}

	if max_seen > (rp.OpsIncluded + THROTTLING_SLACK) {

		rp.Status = THROTTLED

		return rp.Status, nil
	}

	rp.Status = OK
	return rp.Status, nil

}

func (rm *ReputationManager) getEntity(entity common.Address) (*ReputationParams, error) {

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

	rm.Reputation[entity] = &ReputationParams{
		OpsSeen:      0,
		OpsIncluded:  0,
		Status:       0,
		Blacklisted:  false,
		LastSeen:     time.Now(),
		LastIncluded: time.Now(),
	}

	err = rm.db.Put(entity.Bytes(), rm.Reputation[entity].Bytes())

	if err != nil {
		return nil, err
	}

	return rm.Reputation[entity], nil

}

func (rp *ReputationParams) refresh() {
	now := time.Now().Unix()

	PastSeen := (rp.LastSeen.Unix() - now)
	PastIncluded := (rp.LastIncluded.Unix() - now)

	if PastSeen >= anHour {

		rate := (PastSeen / anHour)

		for i := rate; i > 0; i-- {

			rp.OpsSeen = (rp.OpsSeen * 23) / 24

		}

	}

	if PastIncluded >= anHour {

		rate := (PastIncluded / anHour)

		for i := rate; i > 0; i-- {

			rp.OpsIncluded = (rp.OpsIncluded * 23) / 24

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
