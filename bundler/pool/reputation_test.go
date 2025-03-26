package mempool

import (
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
)

/*
const (
	MIN_INCLUSION_RATE_DENOMINATOR = 2
	BAN_SLACK                      = 2
	THROTTLING_SLACK               = 1
)
*/

func TestNewReputationManager(t *testing.T) {
	db := memorydb.New()
	rm := New(db)

	if rm.db != db {
		t.Errorf("New() did not set the database correctly")
	}
	if rm.Reputation == nil {
		t.Errorf("New() did not initialize Reputation map")
	}
}

func TestReputationManager_UpdateOpsSeen(t *testing.T) {
	db := memorydb.New()
	rm := New(db)
	entity := common.HexToAddress("0x0102030405060708091011121314151617181920")

	err := rm.UpdateOpsSeen(entity)
	if err != nil {
		t.Fatalf("UpdateOpsSeen() failed: %v", err)
	}

	rp, err := rm.getEntity(entity)
	if err != nil {
		t.Fatalf("getEntity() failed: %v", err)
	}

	if rp.OpsSeen != 1 {
		t.Errorf("UpdateOpsSeen() did not increment OpsSeen, got %d, want %d", rp.OpsSeen, 1)
	}
	if rp.LastSeen.IsZero() {
		t.Errorf("UpdateOpsSeen() did not update LastSeen")
	}

	// Check persistence in db
	rm2 := New(db) // Create a new RM instance with the same DB to simulate reload
	rp2, err := rm2.getEntity(entity)
	if err != nil {
		t.Fatalf("getEntity() failed after reload: %v", err)
	}
	if rp2.OpsSeen != 1 {
		t.Errorf("OpsSeen not persisted, got %d, want %d", rp2.OpsSeen, 1)
	}
}

func TestReputationManager_UpdateOpsIncluded(t *testing.T) {
	db := memorydb.New()
	rm := New(db)
	entity := common.HexToAddress("0x0102030405060708091011121314151617181920")

	err := rm.UpdateOpsIncluded(entity)
	if err != nil {
		t.Fatalf("UpdateOpsIncluded() failed: %v", err)
	}

	rp, err := rm.getEntity(entity)
	if err != nil {
		t.Fatalf("getEntity() failed: %v", err)
	}

	if rp.OpsIncluded != 1 {
		t.Errorf("UpdateOpsIncluded() did not increment OpsIncluded, got %d, want %d", rp.OpsIncluded, 1)
	}
	if rp.LastIncluded.IsZero() {
		t.Errorf("UpdateOpsIncluded() did not update LastIncluded")
	}

	// Check persistence in db
	rm2 := New(db) // Create a new RM instance with the same DB to simulate reload
	rp2, err := rm2.getEntity(entity)
	if err != nil {
		t.Fatalf("getEntity() failed after reload: %v", err)
	}
	if rp2.OpsIncluded != 1 {
		t.Errorf("OpsIncluded not persisted, got %d, want %d", rp2.OpsIncluded, 1)
	}
}

func TestReputationManager_GetStatus(t *testing.T) {
	db := memorydb.New()
	rm := New(db)
	entity := common.HexToAddress("0x0102030405060708091011121314151617181920")

	// Initial status should be OK
	status, err := rm.GetStatus(entity)
	if err != nil {
		t.Fatalf("GetStatus() failed: %v", err)
	}
	if status != OK {
		t.Errorf("GetStatus() initial status got %d, want %d", status, OK)
	}

	// Throttled status
	rp, _ := rm.getEntity(entity)
	rp.OpsSeen = THROTTLING_SLACK*MIN_INCLUSION_RATE_DENOMINATOR + 15
	rp.OpsIncluded = 0
	db.Put(entity.Bytes(), rp.Bytes()) // persist to db
	rm.Reputation[entity] = rp         // update cache

	status, err = rm.GetStatus(entity)
	if err != nil {
		t.Fatalf("GetStatus() failed: %v", err)
	}
	if status != THROTTLED {
		t.Errorf("GetStatus() throttled status got %d, want %d", status, THROTTLED)
	}

	// Banned status
	rp, _ = rm.getEntity(entity) // get fresh from cache after last status update might change it in memory
	rp.OpsSeen = BAN_SLACK*MIN_INCLUSION_RATE_DENOMINATOR + 15
	rp.OpsIncluded = 0
	db.Put(entity.Bytes(), rp.Bytes()) // persist to db
	rm.Reputation[entity] = rp         // update cache

	status, err = rm.GetStatus(entity)
	if err != nil {
		t.Fatalf("GetStatus() failed: %v", err)
	}
	if status != BANNED {
		t.Errorf("GetStatus() banned status got %d, want %d", status, BANNED)
	}

	// OK status again after including ops
	rp, _ = rm.getEntity(entity) // get fresh from cache after last status update might change it in memory
	rp.OpsIncluded = BAN_SLACK + MIN_INCLUSION_RATE_DENOMINATOR + 1
	db.Put(entity.Bytes(), rp.Bytes()) // persist to db
	rm.Reputation[entity] = rp         // update cache

	status, err = rm.GetStatus(entity)
	if err != nil {
		t.Fatalf("GetStatus() failed: %v", err)
	}
	if status != OK {
		t.Errorf("GetStatus() OK status after include got %d, want %d", status, OK)
	}
}

func TestReputationManager_getEntity(t *testing.T) {
	db := memorydb.New()
	rm := New(db)
	entity := common.HexToAddress("0x0102030405060708091011121314151617181920")

	// Get entity for the first time (should create new)
	rp1, err := rm.getEntity(entity)
	if err != nil {
		t.Fatalf("getEntity() failed for new entity: %v", err)
	}
	if rp1 == nil {
		t.Fatalf("getEntity() returned nil ReputationParams for new entity")
	}

	// Get entity again (should be from cache)
	rp2, err := rm.getEntity(entity)
	if err != nil {
		t.Fatalf("getEntity() failed for cached entity: %v", err)
	}
	if rp1 != rp2 { // Pointer comparison to check if it's the same object from cache
		t.Errorf("getEntity() did not return cached entity")
	}

	// Simulate retrieving from DB after cache eviction (not explicitly evicted in this test, but simulates reload)
	rm.Reputation = make(map[common.Address]*ReputationParams) // Clear cache
	rp3, err := rm.getEntity(entity)
	if err != nil {
		t.Fatalf("getEntity() failed for db entity: %v", err)
	}
	if rp1 == rp3 {
		t.Errorf("getEntity() did not return entity from DB correctly, got %+v, want %+v", rp3, rp1)
	}
}

func TestReputationParams_refresh(t *testing.T) {
	rp := &ReputationParams{
		OpsSeen:      24,
		OpsIncluded:  24,
		LastSeen:     time.Now().Add(time.Hour * 2),
		LastIncluded: time.Now().Add(time.Hour * 2),
	}

	rp.refresh()

	if rp.OpsSeen != ((24*23/24)*23)/24 { // decay twice since 2 hours passed
		t.Errorf("refresh() OpsSeen decay incorrect, got %d, want %d", rp.OpsSeen, ((24*23/24)*23)/24)
	}
	if rp.OpsIncluded != ((24*23/24)*23)/24 { // decay twice since 2 hours passed
		t.Errorf("refresh() OpsIncluded decay incorrect, got %d, want %d", rp.OpsIncluded, ((24*23/24)*23)/24)
	}

	rpNoDecay := &ReputationParams{
		OpsSeen:      10,
		OpsIncluded:  10,
		LastSeen:     time.Now().Add(time.Minute * 30),
		LastIncluded: time.Now().Add(time.Minute * 30),
	}

	opsSeenBefore := rpNoDecay.OpsSeen
	opsIncludedBefore := rpNoDecay.OpsIncluded
	rpNoDecay.refresh()

	if rpNoDecay.OpsSeen != opsSeenBefore {
		t.Errorf("refresh() OpsSeen decayed when it shouldn't, got %d, want %d", rpNoDecay.OpsSeen, opsSeenBefore)
	}

	if rpNoDecay.OpsIncluded != opsIncludedBefore {
		t.Errorf("refresh() OpsIncluded decayed when it shouldn't, got %d, want %d", rpNoDecay.OpsIncluded, opsIncludedBefore)
	}
}

// --- Concurrent tests to check mutex usage ---

func TestReputationManager_ConcurrentUpdates(t *testing.T) {
	db := memorydb.New()
	rm := New(db)
	entity := common.HexToAddress("0x0102030405060708091011121314151617181920")

	var wg sync.WaitGroup
	numUpdates := 100

	wg.Add(numUpdates * 2) // UpdateOpsSeen and UpdateOpsIncluded

	for i := 0; i < numUpdates; i++ {
		go func() {
			defer wg.Done()
			rm.UpdateOpsSeen(entity)
		}()
		go func() {
			defer wg.Done()
			rm.UpdateOpsIncluded(entity)
		}()
	}
	wg.Wait()

	rp, _ := rm.getEntity(entity)
	if rp.OpsSeen != int64(numUpdates) {
		t.Errorf("Concurrent UpdateOpsSeen OpsSeen count incorrect, got %d, want %d", rp.OpsSeen, numUpdates)
	}
	if rp.OpsIncluded != int64(numUpdates) {
		t.Errorf("Concurrent UpdateOpsIncluded OpsIncluded count incorrect, got %d, want %d", rp.OpsIncluded, numUpdates)
	}
}

func TestReputationManager_ConcurrentGetStatus(t *testing.T) {
	db := memorydb.New()
	rm := New(db)
	entity := common.HexToAddress("0x0102030405060708091011121314151617181920")

	var wg sync.WaitGroup
	numReaders := 100

	wg.Add(numReaders)
	for i := 0; i < numReaders; i++ {
		go func() {
			defer wg.Done()
			rm.GetStatus(entity) // Just read status concurrently
		}()
	}
	wg.Wait()
	// No assertions needed here, just checking for race conditions during concurrent reads.
	// If there's a race, the test might panic or produce unexpected results (though unlikely in this simple read-only scenario).
}
