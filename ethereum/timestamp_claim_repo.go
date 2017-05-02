package ethereum

import (
	"errors"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/protocol"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"os"
	"sync"
)

const ACTIVE_SHARE_FILE string = "active_shares"

// TimestampClaimRepo only select shares that don't have most recent timestamp
// in order to make sure coming shares' counters are greater than selected
// shares
type TimestampClaimRepo struct {
	activeShares    map[string]*Share
	recentTimestamp *big.Int
	noShares        uint64
	noRecentShares  uint64
	mu              sync.Mutex
	storage         smartpool.PersistentStorage
	diff            *big.Int
	miner           string
}

func NewTimestampClaimRepo(diff *big.Int, miner, coinbase string, storage smartpool.PersistentStorage) *TimestampClaimRepo {
	shares, err := loadActiveShares(storage)
	if err != nil {
		smartpool.Output.Printf("Couldn't load active shares from last session (%s). Initialize with empty share pool.\n", err)
	}
	noShares := 0
	noRecentShares := 0
	currentTimestamp := big.NewInt(0)
	changedDiff := false
	changedMiner := false
	changedCoinbase := false
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
			if s.BlockHeader().Coinbase.Hex() != coinbase {
				changedCoinbase = true
			}
		}
	}
	var oneShare *Share
	if changedCoinbase {
		smartpool.Output.Printf("SmartPool contract address changed. Discarded %d shares from last session.\n", len(shares))
		shares = map[string]*Share{}
		noShares = 0
		noRecentShares = 0
		currentTimestamp = big.NewInt(0)
	} else if changedMiner {
		for _, s := range shares {
			oneShare = s
			break
		}
		fmt.Printf("You have %d shares from last session with miner %s that were not submitted to the contract.\n", len(shares), oneShare.BlockHeader().Coinbase.Hex())
		fmt.Printf("However you are going to run SmartPool with different miner %s.\n", miner)
		fmt.Printf("Please choose one of following options:\n")
		fmt.Printf("1. Discard those shares and continue running SmartPool with new miner.\n")
		fmt.Printf("2. Abort SmartPool and rerun it with --miner %s\n", oneShare.MinerAddress())
		var choice string
		for {
			fmt.Printf("Enter 1 or 2: ")
			fmt.Scanf("%s", &choice)
			if choice == "1" {
				shares = map[string]*Share{}
				noShares = 0
				noRecentShares = 0
				currentTimestamp = big.NewInt(0)
				smartpool.Output.Printf("You chose to discard the shares from last session.\n")
				break
			} else if choice == "2" {
				os.Exit(1)
			}
		}
	} else if changedDiff {
		for _, s := range shares {
			oneShare = s
			break
		}
		fmt.Printf("You have %d shares from last session with difficulty %s that were not submitted to the contract.\n", len(shares), oneShare.ShareDifficulty().Text(10))
		fmt.Printf("However you are going to run SmartPool with different share difficulty %s.\n", diff.Text(10))
		fmt.Printf("Please choose one of following options:\n")
		fmt.Printf("1. Discard those shares and continue running SmartPool with new difficulty.\n")
		fmt.Printf("2. Abort SmartPool and rerun it with --diff %s\n", oneShare.ShareDifficulty().Text(10))
		var choice string
		for {
			fmt.Printf("Enter 1 or 2: ")
			fmt.Scanf("%s", &choice)
			if choice == "1" {
				shares = map[string]*Share{}
				noShares = 0
				noRecentShares = 0
				currentTimestamp = big.NewInt(0)
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
	return &cr
}

type gobShare struct {
	BlockHeader     *types.Header
	Nonce           types.BlockNonce
	MixDigest       common.Hash
	ShareDifficulty *big.Int
	MinerAddress    string
	SolutionState   int
}

func loadActiveShares(storage smartpool.PersistentStorage) (map[string]*Share, error) {
	shares := map[string]*Share{}
	gobShares := map[string]gobShare{}
	loadedGobShares, err := storage.Load(&gobShares, ACTIVE_SHARE_FILE)
	gobShares = *loadedGobShares.(*map[string]gobShare)
	if err != nil {
		return shares, err
	}
	for k, gobShare := range gobShares {
		shares[k] = &Share{
			gobShare.BlockHeader,
			gobShare.Nonce,
			gobShare.MixDigest,
			gobShare.ShareDifficulty,
			gobShare.MinerAddress,
			gobShare.SolutionState,
			nil,
		}
	}
	return shares, nil
}

func (cr *TimestampClaimRepo) NoActiveShares() uint64 {
	return cr.noShares + cr.noRecentShares
}

func (cr *TimestampClaimRepo) Persist(storage smartpool.PersistentStorage) error {
	gobShares := map[string]gobShare{}
	var shareID string
	for _, s := range cr.activeShares {
		shareID = fmt.Sprintf(
			"%s-%s",
			s.BlockHeader().Hash().Hex(),
			s.Nonce())
		gobShares[shareID] = gobShare{
			s.BlockHeader(),
			s.nonce,
			s.mixDigest,
			s.shareDifficulty,
			s.minerAddress,
			s.SolutionState,
		}
	}
	smartpool.Output.Printf("Saving active shares to disk...")
	err := storage.Persist(&gobShares, ACTIVE_SHARE_FILE)
	if err == nil {
		smartpool.Output.Printf("Done.\n")
	} else {
		smartpool.Output.Printf("Failed. (%s)\n", err.Error())
	}
	return err
}

func (cr *TimestampClaimRepo) AddShare(s smartpool.Share) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	share := s.(*Share)
	shareID := fmt.Sprintf(
		"%s-%s",
		share.BlockHeader().Hash().Hex(),
		share.Nonce())
	if cr.activeShares[shareID] != nil {
		return errors.New("duplicated share")
	} else {
		cr.activeShares[shareID] = share
	}
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
	return nil
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
	newActiveShares := map[string]*Share{}
	for _, s := range cr.activeShares {
		if s.Timestamp().Cmp(cr.recentTimestamp) < 0 {
			c.AddShare(s)
		} else {
			shareID := fmt.Sprintf(
				"%s-%s",
				s.BlockHeader().Hash().Hex(),
				s.Nonce())
			newActiveShares[shareID] = s
		}
	}
	cr.activeShares = newActiveShares
	cr.noShares = 0
	cr.Persist(cr.storage)
	return c
}
