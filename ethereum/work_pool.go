package ethereum

import (
	"github.com/SmartPool/smartpool-client"
	"time"
)

// workpool keeps track of pending works to ensure that each submitted solution
// can actually be accepted by a real pow work.
// workpool also implements ShareReceiver interface.
type WorkPool map[string]*Work

const (
	FullBlockSolution int = 2
	ValidShare        int = 1
	InvalidShare      int = 0
)

// AcceptSolution takes solution and find corresponding work and return
// associated share.
// It returns nil if the work is not found.
func (wp WorkPool) AcceptSolution(s smartpool.Solution) smartpool.Share {
	work := wp[s.WorkID()]
	if work == nil {
		smartpool.Output.Printf("work (%v) doesn't exist in workpool (len: %d)\n", s, len(wp))
		return nil
	}
	share := work.AcceptSolution(s).(*Share)
	if share.SolutionState == FullBlockSolution {
		delete(wp, s.WorkID())
	}
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

func (wp WorkPool) AddWork(w *Work) {
	wp[w.ID()] = w
}

func (wp WorkPool) Cleanning() {
	ticker := time.Tick(140 * time.Second)
	for _ = range ticker {
		count := 0
		for hash, work := range wp {
			if time.Since(work.createdAt) > 7*(12*time.Second) {
				delete(wp, hash)
				count += 1
			}
		}
		if count > 0 {
			smartpool.Output.Printf("Cleaned %d old works.\n", count)
		}
	}
}
