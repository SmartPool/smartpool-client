package stat

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"time"
)

type PeriodRigData struct {
	TotalSubmittedShare      uint64     `json:"total_submitted_share"`
	TotalAcceptedShare       uint64     `json:"total_accepted_share"`
	TotalAcceptedDifficulty  *big.Int   `json:"total_accepted_difficulty"`
	AverageShareDifficulty   *big.Int   `json:"average_share_difficulty"`
	TotalRejectedShare       uint64     `json:"total_rejected_share"`
	TotalHashrate            *big.Int   `json:"total_hashrate"`
	NoHashrateSubmission     uint64     `json:"no_hashrate_submission"`
	AverageReportedHashrate  *big.Int   `json:"average_reported_hashrate"`
	AverageEffectiveHashrate *big.Int   `json:"average_effective_hashrate"`
	TotalBlockFound          uint64     `json:"total_block_found"`
	TimePeriod               uint64     `json:"time_period"`
	StartTime                *time.Time `json:"start_time"`
}

func NewPeriodRigData(timePeriod uint64) *PeriodRigData {
	return &PeriodRigData{
		TotalHashrate:            big.NewInt(0),
		TotalAcceptedDifficulty:  big.NewInt(0),
		AverageShareDifficulty:   big.NewInt(0),
		AverageReportedHashrate:  big.NewInt(0),
		AverageEffectiveHashrate: big.NewInt(0),
		TimePeriod:               timePeriod,
	}
}

func (prd *PeriodRigData) updateAvgHashrate(t time.Time) {
	if prd.NoHashrateSubmission > 0 {
		prd.AverageReportedHashrate.Div(
			prd.TotalHashrate,
			big.NewInt(int64(prd.NoHashrateSubmission)),
		)
	}
}

func (prd *PeriodRigData) updateAvgEffHashrate(t time.Time) {
	duration := int64(t.Sub(*prd.StartTime).Seconds())
	if duration > 0 {
		prd.AverageEffectiveHashrate.Div(
			prd.TotalAcceptedDifficulty,
			big.NewInt(duration),
		)
	}
}

func (prd *PeriodRigData) updateAvgShareDifficulty(t time.Time) {
	if prd.TotalAcceptedShare > 0 {
		prd.AverageShareDifficulty.Div(
			prd.TotalAcceptedDifficulty,
			big.NewInt(int64(prd.TotalAcceptedShare)),
		)
	}
}

type OverallRigData struct {
	LastSubmittedShare       *time.Time `json:"last_submitted_share"`
	LastAcceptedShare        *time.Time `json:"last_accepted_share"`
	LastRejectedShare        *time.Time `json:"last_rejected_share"`
	LastBlock                *time.Time `json:"last_block"`
	TotalSubmittedShare      uint64     `json:"total_submitted_share"`
	TotalAcceptedShare       uint64     `json:"total_accepted_share"`
	TotalAcceptedDifficulty  *big.Int   `json:"total_accepted_difficulty"`
	AverageShareDifficulty   *big.Int   `json:"average_share_difficulty"`
	TotalRejectedShare       uint64     `json:"total_rejected_share"`
	TotalHashrate            *big.Int   `json:"total_hashrate"`
	NoHashrateSubmission     uint64     `json:"no_hashrate_submission"`
	AverageReportedHashrate  *big.Int   `json:"average_reported_hashrate"`
	AverageEffectiveHashrate *big.Int   `json:"average_effective_hashrate"`
	TotalBlockFound          uint64     `json:"total_block_found"`
	StartTime                *time.Time `json:"start_time"`
}

type RigData struct {
	Datas map[uint64]*PeriodRigData
	*OverallRigData
}

func NewRigData() *RigData {
	return &RigData{
		Datas: map[uint64]*PeriodRigData{},
		OverallRigData: &OverallRigData{
			TotalHashrate:            big.NewInt(0),
			TotalAcceptedDifficulty:  big.NewInt(0),
			AverageShareDifficulty:   big.NewInt(0),
			AverageReportedHashrate:  big.NewInt(0),
			AverageEffectiveHashrate: big.NewInt(0),
		},
	}
}

func (rd *RigData) getData(t time.Time) *PeriodRigData {
	timePeriod := TimeToPeriod(t)
	data := rd.Datas[timePeriod]
	if data == nil {
		data = NewPeriodRigData(timePeriod)
		rd.Datas[timePeriod] = data
	}
	return data
}

func (rd *RigData) AddShare(status string, share smartpool.Share, t time.Time) {
	if rd.StartTime == nil {
		rd.StartTime = &t
	}
	rd.LastSubmittedShare = &t
	rd.TotalSubmittedShare++
	curPeriodData := rd.getData(t)
	if curPeriodData.StartTime == nil {
		curPeriodData.StartTime = &t
	}
	curPeriodData.TotalSubmittedShare++
	if status == "accepted" {
		rd.LastAcceptedShare = &t
		rd.TotalAcceptedShare++
		rd.TotalAcceptedDifficulty.Add(rd.TotalAcceptedDifficulty, share.ShareDifficulty())
		rd.updateAvgShareDifficulty(t)
		rd.updateAvgEffHashrate(t)
		curPeriodData.TotalAcceptedShare++
		curPeriodData.TotalAcceptedDifficulty.Add(curPeriodData.TotalAcceptedDifficulty, share.ShareDifficulty())
		curPeriodData.updateAvgShareDifficulty(t)
		curPeriodData.updateAvgEffHashrate(t)
	} else if status == "rejected" {
		rd.LastRejectedShare = &t
		rd.TotalRejectedShare++
		curPeriodData.TotalRejectedShare++
	} else if status == "fullsolution" {
		rd.LastBlock = &t
		rd.LastAcceptedShare = &t
		rd.TotalAcceptedShare++
		rd.TotalAcceptedDifficulty.Add(rd.TotalAcceptedDifficulty, share.ShareDifficulty())
		rd.updateAvgShareDifficulty(t)
		rd.updateAvgEffHashrate(t)
		rd.TotalBlockFound++
		curPeriodData.TotalAcceptedShare++
		curPeriodData.TotalAcceptedDifficulty.Add(curPeriodData.TotalAcceptedDifficulty, share.ShareDifficulty())
		curPeriodData.updateAvgShareDifficulty(t)
		curPeriodData.updateAvgEffHashrate(t)
		curPeriodData.TotalBlockFound++
	}
}

func (rd *RigData) AddHashrate(hashrate hexutil.Uint64, id common.Hash, t time.Time) {
	if rd.StartTime == nil {
		rd.StartTime = &t
	}
	curPeriodData := rd.getData(t)
	if curPeriodData.StartTime == nil {
		curPeriodData.StartTime = &t
	}
	rd.TotalHashrate.Add(rd.TotalHashrate, big.NewInt(int64(hashrate)))
	rd.NoHashrateSubmission++
	rd.updateAvgHashrate(t)
	curPeriodData.TotalHashrate.Add(curPeriodData.TotalHashrate, big.NewInt(int64(hashrate)))
	curPeriodData.NoHashrateSubmission++
	curPeriodData.updateAvgHashrate(t)
}

func (rd *RigData) updateAvgHashrate(t time.Time) {
	if rd.NoHashrateSubmission > 0 {
		rd.AverageReportedHashrate.Div(
			rd.TotalHashrate,
			big.NewInt(int64(rd.NoHashrateSubmission)),
		)
	}
}

func (rd *RigData) updateAvgEffHashrate(t time.Time) {
	duration := int64(t.Sub(*rd.StartTime).Seconds())
	if duration > 0 {
		rd.AverageEffectiveHashrate.Div(
			rd.TotalAcceptedDifficulty,
			big.NewInt(duration),
		)
	}
}

func (rd *RigData) updateAvgShareDifficulty(t time.Time) {
	if rd.TotalAcceptedShare > 0 {
		rd.AverageShareDifficulty.Div(
			rd.TotalAcceptedDifficulty,
			big.NewInt(int64(rd.TotalAcceptedShare)),
		)
	}
}
