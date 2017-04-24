package event

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"time"
)

type EventRecorder struct {
	accShareEvents ShareEventStorage
	rejShareEvents ShareEventStorage
	hashrateEvents HashrateEventStorage
}

func (er *EventRecorder) Record(eventName string, rig Rig, datas ...interface{}) {
	t := time.Now()
	if eventName == "accepted_share" {
		er.accShareEvents.Add(t, rig)
	} else if eventName == "rejected_share" {
		er.rejShareEvents.Add(t, rig)
	} else if eventName == "hashrate" {
		hashrate := datas[0].(hexutil.Uint64)
		id := datas[1].(common.Hash)
		er.hashrateEvents.Add(t, rig, hashrate, id)
	}
}

func (er *EventRecorder) FarmStats(start uint64, end uint64) interface{} {
	accShareStats := er.accShareEvents.FarmStats(start, end)
	rejShareStats := er.rejShareEvents.FarmStats(start, end)
	hashrateStats := er.hashrateEvents.FarmStats(start, end)
	totalShareStats := merge(accShareStats, rejShareStats)
	result := map[string]interface{}{}
	result["accepted_share"] = accShareStats
	result["rejected_share"] = rejShareStats
	result["total_share"] = totalShareStats
	result["hashrate"] = hashrateStats
	result["online_rig"] = er.hashrateEvents.OnlineRigs()
	result["start"] = start
	result["end"] = end
}

func (er *EventRecorder) RigStats(rig Rig, start uint64, end uint64) interface{} {
	accShareStats := er.accShareEvents.RigStats(rig, start, end)
	rejShareStats := er.rejShareEvents.RigStats(rig, start, end)
	hashrateStats := er.hashrateEvents.RigStats(rig, start, end)
	totalShareStats := merge(accShareStats, rejShareStats)
	result := map[string]interface{}{}
	result["accepted_share"] = accShareStats
	result["rejected_share"] = rejShareStats
	result["total_share"] = totalShareStats
	result["hashrate"] = hashrateStats
	result["rig"] = rig
	result["start"] = start
	result["end"] = end
	return result
}
