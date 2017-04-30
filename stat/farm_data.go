package stat

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

type PeriodFarmData struct {
	TotalSubmittedShare      uint64          `json:"total_submitted_share"`
	TotalAcceptedShare       uint64          `json:"total_accepted_share"`
	TotalRejectedShare       uint64          `json:"total_rejected_share"`
	TotalAcceptedDifficulty  *big.Int        `json:"total_accepted_difficulty"`
	AverageShareDifficulty   *big.Int        `json:"average_share_difficulty"`
	TotalSubmittedClaim      uint64          `json:"total_submitted_claim"`
	TotalAcceptedClaim       uint64          `json:"total_accepted_claim"`
	TotalRejectedClaim       uint64          `json:"total_rejected"`
	TotalHashrate            *big.Int        `json:"total_hashrate"`
	NoHashrateSubmission     uint64          `json:"no_hashrate_submission"`
	AverageReportedHashrate  *big.Int        `json:"average_reported_hashrate"`
	AverageEffectiveHashrate *big.Int        `json:"average_effective_hashrate"`
	Rigs                     map[string]bool `json:"rigs"`
	TotalBlockFound          uint64          `json:"total_block_found"`
	TimePeriod               uint64          `json:"time_period"`
	StartTime                *time.Time      `json:"start_time"`
}

func NewPeriodFarmData(timePeriod uint64) *PeriodFarmData {
	return &PeriodFarmData{
		TotalHashrate:            big.NewInt(0),
		TotalAcceptedDifficulty:  big.NewInt(0),
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
	duration := int64(t.Sub(*pfd.StartTime).Seconds())
	if duration > 0 {
		pfd.AverageEffectiveHashrate.Div(
			pfd.TotalAcceptedDifficulty,
			big.NewInt(duration),
		)
	}
}

func (pfd *PeriodFarmData) updateAvgShareDifficulty(t time.Time) {
	if pfd.TotalAcceptedShare > 0 {
		pfd.AverageShareDifficulty.Div(
			pfd.TotalAcceptedDifficulty,
			big.NewInt(int64(pfd.TotalAcceptedShare)),
		)
	}
}

type OverallFarmData struct {
	LastSubmittedShare       *time.Time      `json:"last_submitted_share"`
	LastAcceptedShare        *time.Time      `json:"last_accepted_share"`
	LastRejectedShare        *time.Time      `json:"last_rejected_share"`
	LastBlock                *time.Time      `json:"last_block"`
	TotalSubmittedShare      uint64          `json:"total_submitted_share"`
	TotalAcceptedShare       uint64          `json:"total_accepted_share"`
	TotalAcceptedDifficulty  *big.Int        `json:"total_accepted_difficulty"`
	AverageShareDifficulty   *big.Int        `json:"average_share_difficulty"`
	TotalRejectedShare       uint64          `json:"total_rejected_share"`
	LastSubmittedClaim       *time.Time      `json:"last_submitted_claim"`
	LastAcceptedClaim        *time.Time      `json:"last_accepted_claim"`
	LastRejectedClaim        *time.Time      `json:"last_rejected_claim"`
	TotalSubmittedClaim      uint64          `json:"total_submitted_claim"`
	TotalAcceptedClaim       uint64          `json:"total_accepted_claim"`
	TotalRejectedClaim       uint64          `json:"total_rejected_claim"`
	TotalHashrate            *big.Int        `json:"total_hashrate"`
	NoHashrateSubmission     uint64          `json:"no_hashrate_submission"`
	AverageReportedHashrate  *big.Int        `json:"average_reported_hashrate"`
	AverageEffectiveHashrate *big.Int        `json:"average_effective_hashrate"`
	Rigs                     map[string]bool `json:"rigs"`
	TotalBlockFound          uint64          `json:"total_block_found"`
	StartTime                *time.Time      `json:"start_time"`
}

type FarmData struct {
	Datas map[uint64]*PeriodFarmData
	*OverallFarmData
}

func NewFarmData() *FarmData {
	return &FarmData{
		Datas: map[uint64]*PeriodFarmData{},
		OverallFarmData: &OverallFarmData{
			TotalAcceptedDifficulty:  big.NewInt(0),
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
	if fd.StartTime == nil {
		fd.StartTime = &t
	}
	fd.Rigs[rig.ID()] = true
	fd.LastSubmittedShare = &t
	fd.TotalSubmittedShare++
	curPeriodData := fd.getData(t)
	if curPeriodData.StartTime == nil {
		curPeriodData.StartTime = &t
	}
	curPeriodData.TotalSubmittedShare++
	curPeriodData.Rigs[rig.ID()] = true
	if status == "accepted" {
		fd.LastAcceptedShare = &t
		fd.TotalAcceptedShare++
		fd.TotalAcceptedDifficulty.Add(fd.TotalAcceptedDifficulty, share.ShareDifficulty())
		fd.updateAvgShareDifficulty(t)
		fd.updateAvgEffHashrate(t)
		curPeriodData.TotalAcceptedShare++
		curPeriodData.TotalAcceptedDifficulty.Add(curPeriodData.TotalAcceptedDifficulty, share.ShareDifficulty())
		curPeriodData.updateAvgShareDifficulty(t)
		curPeriodData.updateAvgEffHashrate(t)
	} else if status == "rejected" {
		fd.LastRejectedShare = &t
		fd.TotalRejectedShare++
		curPeriodData.TotalRejectedShare++
	} else if status == "fullsolution" {
		fd.LastBlock = &t
		fd.LastAcceptedShare = &t
		fd.TotalAcceptedShare++
		fd.TotalAcceptedDifficulty.Add(fd.TotalAcceptedDifficulty, share.ShareDifficulty())
		fd.TotalBlockFound++
		fd.updateAvgShareDifficulty(t)
		fd.updateAvgEffHashrate(t)
		curPeriodData.TotalAcceptedShare++
		curPeriodData.TotalAcceptedDifficulty.Add(curPeriodData.TotalAcceptedDifficulty, share.ShareDifficulty())
		curPeriodData.updateAvgShareDifficulty(t)
		curPeriodData.updateAvgEffHashrate(t)
		curPeriodData.TotalBlockFound++
	}
}

func (fd *FarmData) AddClaim(status string, claim smartpool.Claim, t time.Time) {
	if fd.StartTime == nil {
		fd.StartTime = &t
	}
	curPeriodData := fd.getData(t)
	if curPeriodData.StartTime == nil {
		curPeriodData.StartTime = &t
	}
	if status == "submitted" {
		fd.LastSubmittedClaim = &t
		fd.TotalSubmittedClaim++
		curPeriodData.TotalSubmittedClaim++
	} else if status == "accepted" {
		fd.LastAcceptedClaim = &t
		fd.TotalAcceptedClaim++
		curPeriodData.TotalAcceptedClaim++
	} else if status == "rejected" {
		fd.LastRejectedClaim = &t
		fd.TotalRejectedClaim++
		curPeriodData.TotalRejectedClaim++
	}
}

func (fd *FarmData) AddHashrate(rig smartpool.Rig, hashrate hexutil.Uint64, id common.Hash, t time.Time) {
	if fd.StartTime == nil {
		fd.StartTime = &t
	}
	curPeriodData := fd.getData(t)
	if curPeriodData.StartTime == nil {
		curPeriodData.StartTime = &t
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
	duration := int64(t.Sub(*fd.StartTime).Seconds())
	if duration > 0 {
		fd.AverageEffectiveHashrate.Div(
			fd.TotalAcceptedDifficulty,
			big.NewInt(duration),
		)
	}
}

func (fd *FarmData) updateAvgShareDifficulty(t time.Time) {
	if fd.TotalAcceptedShare > 0 {
		fd.AverageShareDifficulty.Div(
			fd.TotalAcceptedDifficulty,
			big.NewInt(int64(fd.TotalAcceptedShare)),
		)
	}
}
