package ethereum

import (
	"github.com/SmartPool/smartpool-client"
	"sync"
	"time"
)

// workpool keeps track of pending works to ensure that each submitted solution
// can actually be accepted by a real pow work.
// workpool also implements ShareReceiver interface.
type WorkPool struct {
	mu    sync.RWMutex
	works map[string]*Work
}

const (
	WORKPOOL_FILE     string = "workpool"
	FullBlockSolution int    = 2
	ValidShare        int    = 1
	InvalidShare      int    = 0
)

// AcceptSolution takes solution and find corresponding work and return
// associated share.
// It returns nil if the work is not found.
func (wp *WorkPool) AcceptSolution(s smartpool.Solution) smartpool.Share {
	wp.mu.RLock()
	defer wp.mu.RUnlock()
	work := wp.works[s.WorkID()]
	if work == nil {
		smartpool.Output.Printf("work (%v) doesn't exist in workpool (len: %d)\n", s, len(wp.works))
		return nil
	}
	share := work.AcceptSolution(s).(*Share)
	if share.SolutionState == InvalidShare {
		smartpool.Output.Printf("Solution (%v) is invalid\n", s)
		return nil
	} else {
		// smartpool.Output.Printf(
		// 	"Create share for work: ID: %s - createdAt: %s - timestamp: 0x%s\n",
		// 	work.ID(),
		// 	work.CreatedAt(),
		// 	work.BlockHeader().Time.Text(16),
		// )
		return share
	}
}

func (wp *WorkPool) AddWork(w *Work) {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	wp.works[w.ID()] = w
}

func (wp *WorkPool) RemoveWork(hash string) {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	delete(wp.works, hash)
}

func (wp *WorkPool) oldHashes() []string {
	wp.mu.RLock()
	defer wp.mu.RUnlock()
	result := []string{}
	for hash, work := range wp.works {
		if time.Since(work.CreatedAt) > 200*(12*time.Second) {
			result = append(result, hash)
		}
	}
	return result
}

func (wp *WorkPool) Clean() {
	oldHashes := wp.oldHashes()
	for _, hash := range oldHashes {
		wp.RemoveWork(hash)
	}
	if len(oldHashes) > 0 {
		smartpool.Output.Printf("Cleaned %d old works.\n", len(oldHashes))
	}
}

func (wp *WorkPool) RunCleaner() {
	ticker := time.Tick(140 * time.Second)
	for _ = range ticker {
		wp.Clean()
	}
}

func (wp *WorkPool) Persist(storage smartpool.PersistentStorage) error {
	wp.mu.RLock()
	defer wp.mu.RUnlock()
	smartpool.Output.Printf("Saving workpool to disk...\n")
	err := storage.Persist(&wp.works, WORKPOOL_FILE)
	if err == nil {
		smartpool.Output.Printf("Done.\n")
	} else {
		smartpool.Output.Printf("Failed. (%s)\n", err.Error())
	}
	return err
}

func NewWorkPool(storage smartpool.PersistentStorage) *WorkPool {
	wp, err := loadWorkPool(storage)
	if err != nil {
		smartpool.Output.Printf("Couldn't load workpool from last session (%s). Initialize with empty workpool.\n", err)
	}
	smartpool.Output.Printf("Loaded %d works from last session.\n", len(wp.works))
	return wp
}

func loadWorkPool(storage smartpool.PersistentStorage) (*WorkPool, error) {
	wp := &WorkPool{sync.RWMutex{}, map[string]*Work{}}
	works := map[string]*Work{}
	loadedWorks, err := storage.Load(&works, WORKPOOL_FILE)
	if err != nil {
		return wp, err
	}
	wp.works = *loadedWorks.(*map[string]*Work)
	return wp, err
}
