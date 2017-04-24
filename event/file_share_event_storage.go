package event

import (
	"math/big"
	"sync"
	"time"
)

type DataPoint struct {
	UnixTime uint64 `json:"time"`
	Count    uint64 `json:"count"`
}

func (dp *DataPoint) Inc() {
	dp.Count++
}

type FileShareEventStorage struct {
	mux        sync.Mutex
	datas      map[smartpool.Rig][]DataPoint
	lastShare  time.Time
	totalShare uint64
	startTime  time.Time
	endTime    time.Time
}

func (ses *FileShareEventStorage) rigDatas(rig smartpool.Rig) []DataPoint {
	result, ok := ses.datas[rig]
	if ok {
		return result
	} else {
		ses.datas[rig] = []DataPoint{}
		return ses.datas[rig]
	}
}

func (ses *FileShareEventStorage) Add(t time.Time, rig smartpool.Rig) {
	ses.mux.Lock()
	defer ses.mux.Unlock()
	tick := uint64(t.Unix() / 60)
	if ses.totalShare == 0 {
		ses.startTime = t
	}
	ses.totalShare++
	ses.endTime = t
	rigDatas := ses.rigDatas(rig)
	if len(rigDatas) == 0 {
		rigDatas = append(rigDatas, DataPoint{tick, 1})
	} else {
		data := rigDatas[len(rigDatas-1)]
		if data.UnixTime == tick {
			data.Inc()
		} else {
			rigDatas = append(rigDatas, DataPoint{tick, 1})
		}
	}
	ses.lastShare = t
}

type FarmShareStat struct {
	LastShare  time.Time `json:"last_share"`
	TotalShare uint64    `json:"total_share"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
}

func (ses *FileShareEventStorage) FarmStats() interface{} {
	return FarmShareStat{
		ses.lastShare,
		ses.totalShare,
		ses.startTime,
		ses.endTime,
	}
}

type FarmShareRangeStat struct {
	Datas        []DataPoint `json:"data"`
	TotalShare   uint64      `json:"total_share"`
	AverageShare float64     `json:"average_share"`
	StartTick    uint64      `json:"start_tick"`
	EndTick      uint64      `json:"end_tick"`
}

func (ses *FileShareEventStorage) FarmRangeStats(start uint64, end uint64) interface{} {
	startTick := start / 60
	endTick := end / 60
	result := FarmShareRangeStat{}
	result.StartTick = startTick
	result.EndTick = endTick
	samples := map[uint64]DataPoint{}
	for rig, rigData := range ses.datas {
		for i := len(rigData) - 1; i >= 0; i-- {
			data := rigData[i]
			if data.UnixTime < startTick {
				break
			} else if data.UnixTime <= endTick {
				result.TotalShare += data.Count
				sample, ok := samples[data.UnixTime]
				if ok {
					sample.Count += data.Count
				} else {
					samples[data.UnixTime] = DataPoint{data.UnixTime, data.Count}
				}
			}
		}
	}
	result.AverageShare = float64(result.TotalShare) / float64(result.EndTick-result.StartTick+1)
	result.Datas = []DataPoint{}
	for _, datapoint := range samples {
		result.Datas = append(result.Datas, datapoint)
	}
	return result
}
