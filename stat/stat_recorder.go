package stat

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
	"time"
)

type StatRecorder struct {
	mu sync.Mutex

	RigDatas map[string]*RigData
	FarmData *FarmData
}

func NewStatRecorder() *StatRecorder {
	return &StatRecorder{
		mu:       sync.Mutex{},
		RigDatas: map[string]*RigData{},
		FarmData: NewFarmData(),
	}
}

func (sr *StatRecorder) getRigData(rig smartpool.Rig) *RigData {
	data := sr.RigDatas[rig.ID()]
	if data == nil {
		data = NewRigData()
		sr.RigDatas[rig.ID()] = data
	}
	return data
}

func (sr *StatRecorder) RecordShare(status string, share smartpool.Share, rig smartpool.Rig) {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	t := time.Now()
	rigData := sr.getRigData(rig)
	rigData.AddShare(status, share, t)
	sr.FarmData.AddShare(rig, status, share, t)
}

func (sr *StatRecorder) RecordClaim(status string, claim smartpool.Claim) {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	t := time.Now()
	sr.FarmData.AddClaim(status, claim, t)
}

func (sr *StatRecorder) RecordHashrate(hashrate hexutil.Uint64, id common.Hash, rig smartpool.Rig) {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	t := time.Now()
	rigData := sr.getRigData(rig)
	rigData.AddHashrate(hashrate, id, t)
	sr.FarmData.AddHashrate(rig, hashrate, id, t)
}

func (sr *StatRecorder) OverallFarmStat() interface{} {
	return sr.FarmData.OverallFarmData
}

func (sr *StatRecorder) FarmStat(start uint64, end uint64) interface{} {
	result := map[uint64]*PeriodFarmData{}
	for period, data := range sr.FarmData.Datas {
		if start <= period && period <= end {
			result[period] = data
		}
	}
	return result
}

func (sr *StatRecorder) OverallRigStat(rig smartpool.Rig) interface{} {
	rigData := sr.getRigData(rig)
	return rigData.OverallRigData
}

func (sr *StatRecorder) RigStat(rig smartpool.Rig, start, end uint64) interface{} {
	rigData := sr.getRigData(rig)
	result := map[uint64]*PeriodRigData{}
	for period, data := range rigData.Datas {
		if start <= period && period <= end {
			result[period] = data
		}
	}
	return result
}

// func (er *EventRecorder) FarmStats(start uint64, end uint64) interface{} {
// 	accShareStats := er.accShareEvents.FarmStats(start, end)
// 	rejShareStats := er.rejShareEvents.FarmStats(start, end)
// 	hashrateStats := er.hashrateEvents.FarmStats(start, end)
// 	totalShareStats := merge(accShareStats, rejShareStats)
// 	result := map[string]interface{}{}
// 	result["accepted_share"] = accShareStats
// 	result["rejected_share"] = rejShareStats
// 	result["total_share"] = totalShareStats
// 	result["hashrate"] = hashrateStats
// 	result["online_rig"] = er.hashrateEvents.OnlineRigs()
// 	result["start"] = start
// 	result["end"] = end
// }
//
// func (er *EventRecorder) RigStats(rig Rig, start uint64, end uint64) interface{} {
// 	accShareStats := er.accShareEvents.RigStats(rig, start, end)
// 	rejShareStats := er.rejShareEvents.RigStats(rig, start, end)
// 	hashrateStats := er.hashrateEvents.RigStats(rig, start, end)
// 	totalShareStats := merge(accShareStats, rejShareStats)
// 	result := map[string]interface{}{}
// 	result["accepted_share"] = accShareStats
// 	result["rejected_share"] = rejShareStats
// 	result["total_share"] = totalShareStats
// 	result["hashrate"] = hashrateStats
// 	result["rig"] = rig
// 	result["start"] = start
// 	result["end"] = end
// 	return result
// }
