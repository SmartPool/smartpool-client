package ethereum

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/protocol"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
)

// TimestampClaimRepo only select shares that don't have most recent timestamp
// in order to make sure coming shares' counters are greater than selected
// shares
type TimestampClaimRepo struct {
	activeShares    []*Share
	recentTimestamp *big.Int
	noShares        uint64
	noRecentShares  uint64
	mu              sync.Mutex
}

// TODO: Load from persistent storage
func NewTimestampClaimRepo() *TimestampClaimRepo {
	return &TimestampClaimRepo{
		[]*Share{},
		big.NewInt(0),
		0,
		0,
		sync.Mutex{},
	}
}

func (cr *TimestampClaimRepo) AddShare(s smartpool.Share) {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	share := s.(*Share)
	cr.activeShares = append(cr.activeShares, share)
	if share.Timestamp().Cmp(cr.recentTimestamp) == 0 {
		cr.noRecentShares++
	} else if share.Timestamp().Cmp(cr.recentTimestamp) < 0 {
		cr.noShares++
	} else if share.Timestamp().Cmp(cr.recentTimestamp) > 0 {
		cr.noShares += cr.noRecentShares
		cr.noRecentShares = 1
		cr.recentTimestamp = big.NewInt(0)
		cr.recentTimestamp.Add(share.Timestamp(), common.Big0)
	}
}

func (cr *TimestampClaimRepo) GetCurrentClaim(threshold int) smartpool.Claim {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	smartpool.Output.Printf("Have %d eligible shares\n", cr.noShares)
	smartpool.Output.Printf("Current timestamp: 0x%s\n", cr.recentTimestamp.Text(16))
	smartpool.Output.Printf("Shares with current timestamp: %d\n", cr.noRecentShares)
	if cr.noShares < uint64(threshold) {
		return nil
	}
	c := protocol.NewClaim()
	newActiveShares := []*Share{}
	for _, s := range cr.activeShares {
		if s.Timestamp().Cmp(cr.recentTimestamp) < 0 {
			c.AddShare(s)
		} else {
			newActiveShares = append(newActiveShares, s)
		}
	}
	cr.activeShares = newActiveShares
	cr.noShares = 0
	return c
}
