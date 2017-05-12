package stat

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"math/big"
	"time"
)

type RigHashrate struct {
	ReportedHashrate  *big.Int
	EffectiveHashrate *big.Int
}

func NewRigHashrate() *RigHashrate {
	return &RigHashrate{
		ReportedHashrate:  big.NewInt(0),
		EffectiveHashrate: big.NewInt(0),
	}
}

type PeriodFarmData struct {
	MinedShare             uint64                  `json:"mined_share"`
	ValidShare             uint64                  `json:"valid_share"`
	TotalValidDifficulty   *big.Int                `json:"-"`
	AverageShareDifficulty *big.Int                `json:"average_share_difficulty"`
	RejectedShare          uint64                  `json:"rejected_share"`
	SubmittedClaim         uint64                  `json:"submitted_claim"`
	AcceptedClaim          uint64                  `json:"accepted_claim"`
	RejectedClaim          uint64                  `json:"rejected_claim"`
	ReportedHashrate       *big.Int                `json:"reported_hashrate"`
	EffectiveHashrate      *big.Int                `json:"effective_hashrate"`
	Rigs                   map[string]*RigHashrate `json:"rigs"`
	BlockFound             uint64                  `json:"block_found"`
	TimePeriod             uint64                  `json:"time_period"`
	StartTime              time.Time               `json:"start_time"`
}

func NewPeriodFarmData(timePeriod uint64) *PeriodFarmData {
	return &PeriodFarmData{
		TotalValidDifficulty:   big.NewInt(0),
		AverageShareDifficulty: big.NewInt(0),
		ReportedHashrate:       big.NewInt(0),
		EffectiveHashrate:      big.NewInt(0),
		Rigs:                   map[string]*RigHashrate{},
		TimePeriod:             timePeriod,
	}
}

func (pfd *PeriodFarmData) updateAvgEffHashrate(t time.Time) {
	duration := int64(t.Sub(pfd.StartTime).Seconds())
	if duration > 0 {
		pfd.EffectiveHashrate.Div(
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

func (pfd *PeriodFarmData) getRigHashrate(id string) *RigHashrate {
	if _, exist := pfd.Rigs[id]; !exist {
		pfd.Rigs[id] = NewRigHashrate()
	}
	return pfd.Rigs[id]
}

type OverallFarmData struct {
	LastMinedShare         time.Time               `json:"last_mined_share"`
	LastValidShare         time.Time               `json:"last_valid_share"`
	LastRejectedShare      time.Time               `json:"last_rejected_share"`
	LastBlock              time.Time               `json:"last_block"`
	MinedShare             uint64                  `json:"mined_share"`
	ValidShare             uint64                  `json:"valid_share"`
	TotalValidDifficulty   *big.Int                `json:"-"`
	AverageShareDifficulty *big.Int                `json:"average_share_difficulty"`
	RejectedShare          uint64                  `json:"rejected_share"`
	LastSubmittedClaim     time.Time               `json:"last_submitted_claim"`
	LastAcceptedClaim      time.Time               `json:"last_accepted_claim"`
	LastRejectedClaim      time.Time               `json:"last_rejected_claim"`
	SubmittedClaim         uint64                  `json:"total_submitted_claim"`
	AcceptedClaim          uint64                  `json:"total_accepted_claim"`
	RejectedClaim          uint64                  `json:"total_rejected_claim"`
	ReportedHashrate       *big.Int                `json:"reported_hashrate"`
	EffectiveHashrate      *big.Int                `json:"effective_hashrate"`
	Rigs                   map[string]*RigHashrate `json:"rigs"`
	BlockFound             uint64                  `json:"total_block_found"`
	// A share is pending when it is not included in any claim
	PendingShare uint64 `json:"pending_share"`
	// A share is considered as abandoned when it is discarded while being
	// restored from last running session because of configuration changes
	// In case when the client was shutdown while waiting to verify a claim,
	// its shares will also be abandoned and will be counted to
	// AbandonedShare too.
	AbandonedShare      uint64    `json:"abandoned_share"`
	BeingValidatedShare uint64    `json:"being_validated_share"`
	VerifiedShare       uint64    `json:"verified_share"`
	BadShare            uint64    `json:"bad_share"`
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
			TotalValidDifficulty:   big.NewInt(0),
			AverageShareDifficulty: big.NewInt(0),
			ReportedHashrate:       big.NewInt(0),
			EffectiveHashrate:      big.NewInt(0),
			Rigs:                   map[string]*RigHashrate{},
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

func (fd *FarmData) TruncateData(storage smartpool.PersistentStorage) error {
	curPeriod := TimeToPeriod(time.Now())
	var err error
	for period, farmData := range fd.Datas {
		if int64(curPeriod-period) > LongWindow/BaseTimePeriod {
			if err = storage.Persist(farmData, fmt.Sprintf("farm-data-%d", period)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (fd *FarmData) AddShare(rig smartpool.Rig, status string, share smartpool.Share, t time.Time) {
	if fd.StartTime.IsZero() {
		fd.StartTime = t
	}
	if _, exist := fd.Rigs[rig.ID()]; !exist {
		fd.Rigs[rig.ID()] = NewRigHashrate()
	}
	curPeriodData := fd.getData(t)
	if curPeriodData.StartTime.IsZero() {
		curPeriodData.StartTime = t
	}
	if _, exist := curPeriodData.Rigs[rig.ID()]; !exist {
		curPeriodData.Rigs[rig.ID()] = NewRigHashrate()
	}
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
		fd.BadShare += claim.NumShares().Uint64()
		curPeriodData.RejectedClaim++
	}
}

func (fd *FarmData) ShareRestored(noShares uint64) {
	numAbandoned := fd.PendingShare - noShares
	fd.PendingShare -= numAbandoned
	fd.AbandonedShare += numAbandoned
	// TODO: We currently discard last unverified claim
	// on client startup.
	// So we increase RejectedClaim and consider all its shares
	// as AbandonedShare
	// NOTE: There is a case that the verification tx was
	// already sent out. It would be likely to be mined and the
	// claim is actually verified and its share shouldn't be
	// consider as abandoned.
	// SOLUTION:
	// 1. Recover claim which is waiting to verify
	// 2. Prevent the client from shutting down when its waiting
	// to verify last claim
	if fd.BeingValidatedShare != 0 {
		fd.RejectedClaim++
	}
	fd.AbandonedShare += fd.BeingValidatedShare
	fd.BeingValidatedShare = 0
}

func (fd *FarmData) UpdateRigHashrate(
	rig smartpool.Rig, reportedHashrate *big.Int, effectiveHashrate *big.Int,
	periodReportedHashrate *big.Int, periodEffectiveHashrate *big.Int,
	t time.Time) {
	if fd.StartTime.IsZero() {
		fd.StartTime = t
	}
	curPeriodData := fd.getData(t)
	if curPeriodData.StartTime.IsZero() {
		curPeriodData.StartTime = t
	}
	rigHashrate := fd.getRigHashrate(rig.ID())
	changedReportedHashrate := reportedHashrate.Sub(reportedHashrate, rigHashrate.ReportedHashrate)
	rigHashrate.ReportedHashrate = reportedHashrate
	changedEffectiveHashrate := effectiveHashrate.Sub(effectiveHashrate, rigHashrate.EffectiveHashrate)
	fd.ReportedHashrate.Add(fd.ReportedHashrate, changedReportedHashrate)
	fd.EffectiveHashrate.Add(fd.EffectiveHashrate, changedEffectiveHashrate)
	rigHashrate = curPeriodData.getRigHashrate(rig.ID())
	changedReportedHashrate = periodReportedHashrate.Sub(periodReportedHashrate, rigHashrate.ReportedHashrate)
	curPeriodData.ReportedHashrate.Add(curPeriodData.ReportedHashrate, changedReportedHashrate)
	changedEffectiveHashrate = periodEffectiveHashrate.Sub(periodEffectiveHashrate, rigHashrate.EffectiveHashrate)
	curPeriodData.EffectiveHashrate.Add(curPeriodData.EffectiveHashrate, changedEffectiveHashrate)
}

func (fd *FarmData) getRigHashrate(id string) *RigHashrate {
	if _, exist := fd.Rigs[id]; !exist {
		fd.Rigs[id] = NewRigHashrate()
	}
	return fd.Rigs[id]
}

func (fd *FarmData) updateAvgEffHashrate(t time.Time) {
	duration := int64(t.Sub(fd.StartTime).Seconds())
	if duration > 0 {
		fd.EffectiveHashrate.Div(
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
