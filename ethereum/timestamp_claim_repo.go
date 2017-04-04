package ethereum

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/protocol"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"os"
	"sync"
	"time"
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
	storage         *FileStorage
	diff            *big.Int
	miner           string
}

func NewTimestampClaimRepo(diff *big.Int, miner string) *TimestampClaimRepo {
	storage := NewFileStorage()
	shares, err := storage.LoadActiveShares()
	if err != nil {
		smartpool.Output.Printf("Couldn't load active shares from last session (%s). Initialize with empty share pool.\n", err)
	}
	noShares := 0
	noRecentShares := 0
	currentTimestamp := big.NewInt(0)
	changedDiff := false
	changedMiner := false
	if len(shares) > 0 {
		for _, s := range shares {
			if currentTimestamp.Cmp(s.Timestamp()) < 0 {
				currentTimestamp.Add(s.Timestamp(), common.Big0)
			}
		}
		for _, s := range shares {
			if s.Timestamp().Cmp(currentTimestamp) == 0 {
				noRecentShares++
			} else {
				noShares++
			}
			if s.ShareDifficulty().Cmp(diff) != 0 {
				changedDiff = true
			}
			if s.MinerAddress() != miner {
				changedMiner = true
			}
		}
	}
	if changedMiner {
		fmt.Printf("You have %d shares from last session with miner %s that were not submitted to the contract.\n", len(shares), shares[0].BlockHeader().Coinbase.Hex())
		fmt.Printf("However you are going to run SmartPool with different miner %s.\n", miner)
		fmt.Printf("Please choose one of following options:\n")
		fmt.Printf("1. Discard those shares and continue running SmartPool with new miner.\n")
		fmt.Printf("2. Abort SmartPool and rerun it with --miner %s\n", shares[0].MinerAddress())
		var choice string
		for {
			fmt.Printf("Enter 1 or 2: ")
			fmt.Scanf("%s", &choice)
			if choice == "1" {
				shares = []*Share{}
				smartpool.Output.Printf("You chose to discard the shares from last session.\n")
				break
			} else if choice == "2" {
				os.Exit(1)
			}
		}
	}
	if changedDiff {
		fmt.Printf("You have %d shares from last session with difficulty %s that were not submitted to the contract.\n", len(shares), shares[0].ShareDifficulty().Text(10))
		fmt.Printf("However you are going to run SmartPool with different share difficulty %s.\n", diff.Text(10))
		fmt.Printf("Please choose one of following options:\n")
		fmt.Printf("1. Discard those shares and continue running SmartPool with new difficulty.\n")
		fmt.Printf("2. Abort SmartPool and rerun it with --diff %s\n", shares[0].ShareDifficulty().Text(10))
		var choice string
		for {
			fmt.Printf("Enter 1 or 2: ")
			fmt.Scanf("%s", &choice)
			if choice == "1" {
				shares = []*Share{}
				smartpool.Output.Printf("You chose to discard the shares from last session.\n")
				break
			} else if choice == "2" {
				os.Exit(1)
			}
		}
	}
	cr := TimestampClaimRepo{
		shares,
		currentTimestamp,
		uint64(noShares),
		uint64(noRecentShares),
		sync.Mutex{},
		storage,
		diff,
		miner,
	}
	smartpool.Output.Printf("Loaded %d valid shares\n", noShares)
	smartpool.Output.Printf("Loaded timestamp: 0x%s\n", currentTimestamp.Text(16))
	smartpool.Output.Printf("Loaded %d shares with current timestamp\n", noRecentShares)
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			smartpool.Output.Printf("Saving active shares to disk...")
			err = cr.Persist()
			smartpool.Output.Printf("Done. (%s)\n", err)
		}
	}()
	smartpool.Output.Printf("Share persister is running...\n")
	return &cr
}

func (cr *TimestampClaimRepo) Persist() error {
	return cr.storage.PersistActiveShares(cr.activeShares)
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
	smartpool.Output.Printf("Have %d valid shares\n", cr.noShares)
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
