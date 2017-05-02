package stat

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

type PeriodFarmData struct {
	MinedShare               uint64          `json:"mined_share"`
	ValidShare               uint64          `json:"valid_share"`
	TotalValidDifficulty     *big.Int        `json:"-"`
	AverageShareDifficulty   *big.Int        `json:"average_share_difficulty"`
	RejectedShare            uint64          `json:"rejected_share"`
	SubmittedClaim           uint64          `json:"submitted_claim"`
	AcceptedClaim            uint64          `json:"accepted_claim"`
	RejectedClaim            uint64          `json:"rejected_claim"`
	TotalHashrate            *big.Int        `json:"-"`
	NoHashrateSubmission     uint64          `json:"-"`
	AverageReportedHashrate  *big.Int        `json:"reported_hashrate"`
	AverageEffectiveHashrate *big.Int        `json:"effective_hashrate"`
	Rigs                     map[string]bool `json:"rigs"`
	BlockFound               uint64          `json:"block_found"`
	TimePeriod               uint64          `json:"time_period"`
	StartTime                time.Time       `json:"start_time"`
}

func NewPeriodFarmData(timePeriod uint64) *PeriodFarmData {
	return &PeriodFarmData{
		TotalHashrate:            big.NewInt(0),
		TotalValidDifficulty:     big.NewInt(0),
		AverageShareDifficulty:   big.NewInt(0),
		AverageReportedHashrate:  big.NewInt(0),
		AverageEffectiveHashrate: big.NewInt(0),
		Rigs:       map[string]bool{},
		TimePeriod: timePeriod,
	}
}

func (pfd *PeriodFarmData) updateAvgHashrate(t time.Time) {
	if pfd.NoHashrateSubmission > 0 {
		pfd.AverageReportedHashrate.Div(
			pfd.TotalHashrate,
			big.NewInt(int64(pfd.NoHashrateSubmission)),
		)
	}
}

func (pfd *PeriodFarmData) updateAvgEffHashrate(t time.Time) {
	duration := int64(t.Sub(pfd.StartTime).Seconds())
	if duration > 0 {
		pfd.AverageEffectiveHashrate.Div(
			pfd.TotalValidDifficulty,
			big.NewInt(duration),
		)
	}
}

func (pfd *PeriodFarmData) updateAvgShareDifficulty(t time.Time) {
	if pfd.ValidShare > 0 {
		pfd.AverageShareDifficulty.Div(
			pfd.TotalValidDifficulty,
			big.NewInt(int64(pfd.ValidShare)),
		)
	}
}

type OverallFarmData struct {
	LastMinedShare           time.Time       `json:"last_mined_share"`
	LastValidShare           time.Time       `json:"last_valid_share"`
	LastRejectedShare        time.Time       `json:"last_rejected_share"`
	LastBlock                time.Time       `json:"last_block"`
	MinedShare               uint64          `json:"mined_share"`
	ValidShare               uint64          `json:"valid_share"`
	TotalValidDifficulty     *big.Int        `json:"-"`
	AverageShareDifficulty   *big.Int        `json:"average_share_difficulty"`
	RejectedShare            uint64          `json:"rejected_share"`
	LastSubmittedClaim       time.Time       `json:"last_submitted_claim"`
	LastAcceptedClaim        time.Time       `json:"last_accepted_claim"`
	LastRejectedClaim        time.Time       `json:"last_rejected_claim"`
	SubmittedClaim           uint64          `json:"total_submitted_claim"`
	AcceptedClaim            uint64          `json:"total_accepted_claim"`
	RejectedClaim            uint64          `json:"total_rejected_claim"`
	TotalHashrate            *big.Int        `json:"-"`
	NoHashrateSubmission     uint64          `json:"-"`
	AverageReportedHashrate  *big.Int        `json:"reported_hashrate"`
	AverageEffectiveHashrate *big.Int        `json:"effective_hashrate"`
	Rigs                     map[string]bool `json:"rigs"`
	BlockFound               uint64          `json:"total_block_found"`
	// A share is pending when it is not included in any claim
	PendingShare uint64 `json:"pending_share"`
	// A share is considered as abandoned when the claim it belongs to was
	// rejected by the contract or it is discarded while being restored
	// from last running session because of configuration changes
	AbandonedShare      uint64    `json:"abandoned_share"`
	BeingValidatedShare uint64    `json:"being_validated_share"`
	VerifiedShare       uint64    `json:"verified_share"`
	StartTime           time.Time `json:"start_time"`
}

type FarmData struct {
	Datas map[uint64]*PeriodFarmData
	*OverallFarmData
}

func NewFarmData() *FarmData {
	return &FarmData{
		Datas: map[uint64]*PeriodFarmData{},
		OverallFarmData: &OverallFarmData{
			TotalValidDifficulty:     big.NewInt(0),
			TotalHashrate:            big.NewInt(0),
			AverageShareDifficulty:   big.NewInt(0),
			AverageReportedHashrate:  big.NewInt(0),
			AverageEffectiveHashrate: big.NewInt(0),
			Rigs: map[string]bool{},
		},
	}
}

func (fd *FarmData) getData(t time.Time) *PeriodFarmData {
	timePeriod := TimeToPeriod(t)
	data := fd.Datas[timePeriod]
	if data == nil {
		data = NewPeriodFarmData(timePeriod)
		fd.Datas[timePeriod] = data
	}
	return data
}

func (fd *FarmData) AddShare(rig smartpool.Rig, status string, share smartpool.Share, t time.Time) {
	if fd.StartTime.IsZero() {
		fd.StartTime = t
	}
	fd.Rigs[rig.ID()] = true
	curPeriodData := fd.getData(t)
	if curPeriodData.StartTime.IsZero() {
		curPeriodData.StartTime = t
	}
	curPeriodData.Rigs[rig.ID()] = true
	if status == "submitted" {
		fd.LastMinedShare = t
		fd.MinedShare++
		curPeriodData.MinedShare++
	} else if status == "accepted" {
		fd.LastValidShare = t
		fd.ValidShare++
		fd.TotalValidDifficulty.Add(fd.TotalValidDifficulty, share.ShareDifficulty())
		fd.updateAvgShareDifficulty(t)
		fd.updateAvgEffHashrate(t)
		fd.PendingShare++
		curPeriodData.ValidShare++
		curPeriodData.TotalValidDifficulty.Add(curPeriodData.TotalValidDifficulty, share.ShareDifficulty())
		curPeriodData.updateAvgShareDifficulty(t)
		curPeriodData.updateAvgEffHashrate(t)
	} else if status == "rejected" {
		fd.LastRejectedShare = t
		fd.RejectedShare++
		curPeriodData.RejectedShare++
	} else if status == "fullsolution" {
		fd.LastBlock = t
		fd.LastValidShare = t
		fd.ValidShare++
		fd.TotalValidDifficulty.Add(fd.TotalValidDifficulty, share.ShareDifficulty())
		fd.BlockFound++
		fd.updateAvgShareDifficulty(t)
		fd.updateAvgEffHashrate(t)
		fd.PendingShare++
		curPeriodData.ValidShare++
		curPeriodData.TotalValidDifficulty.Add(curPeriodData.TotalValidDifficulty, share.ShareDifficulty())
		curPeriodData.updateAvgShareDifficulty(t)
		curPeriodData.updateAvgEffHashrate(t)
		curPeriodData.BlockFound++
	}
}

func (fd *FarmData) AddClaim(status string, claim smartpool.Claim, t time.Time) {
	if fd.StartTime.IsZero() {
		fd.StartTime = t
	}
	curPeriodData := fd.getData(t)
	if curPeriodData.StartTime.IsZero() {
		curPeriodData.StartTime = t
	}
	if status == "submitted" {
		fd.LastSubmittedClaim = t
		fd.SubmittedClaim++
		fd.PendingShare -= claim.NumShares().Uint64()
		fd.BeingValidatedShare += claim.NumShares().Uint64()
		curPeriodData.SubmittedClaim++
	} else if status == "accepted" {
		fd.LastAcceptedClaim = t
		fd.AcceptedClaim++
		fd.VerifiedShare += claim.NumShares().Uint64()
		fd.BeingValidatedShare -= claim.NumShares().Uint64()
		curPeriodData.AcceptedClaim++
	} else if status == "rejected" {
		fd.LastRejectedClaim = t
		fd.RejectedClaim++
		fd.BeingValidatedShare -= claim.NumShares().Uint64()
		fd.AbandonedShare += claim.NumShares().Uint64()
		curPeriodData.RejectedClaim++
	}
}

func (fd *FarmData) ShareRestored(noShares uint64) {
	numAbandoned := fd.PendingShare - noShares
	fd.PendingShare -= numAbandoned
	fd.AbandonedShare += numAbandoned
}

func (fd *FarmData) AddHashrate(rig smartpool.Rig, hashrate hexutil.Uint64, id common.Hash, t time.Time) {
	if fd.StartTime.IsZero() {
		fd.StartTime = t
	}
	curPeriodData := fd.getData(t)
	if curPeriodData.StartTime.IsZero() {
		curPeriodData.StartTime = t
	}
	fd.TotalHashrate.Add(fd.TotalHashrate, big.NewInt(int64(hashrate)))
	fd.NoHashrateSubmission++
	fd.updateAvgHashrate(t)
	curPeriodData.TotalHashrate.Add(curPeriodData.TotalHashrate, big.NewInt(int64(hashrate)))
	curPeriodData.NoHashrateSubmission++
	curPeriodData.updateAvgHashrate(t)
}

func (fd *FarmData) updateAvgHashrate(t time.Time) {
	if fd.NoHashrateSubmission > 0 {
		fd.AverageReportedHashrate.Div(
			fd.TotalHashrate,
			big.NewInt(int64(fd.NoHashrateSubmission)),
		)
	}
}

func (fd *FarmData) updateAvgEffHashrate(t time.Time) {
	duration := int64(t.Sub(fd.StartTime).Seconds())
	if duration > 0 {
		fd.AverageEffectiveHashrate.Div(
			fd.TotalValidDifficulty,
			big.NewInt(duration),
		)
	}
}

func (fd *FarmData) updateAvgShareDifficulty(t time.Time) {
	if fd.ValidShare > 0 {
		fd.AverageShareDifficulty.Div(
			fd.TotalValidDifficulty,
			big.NewInt(int64(fd.ValidShare)),
		)
	}
}
