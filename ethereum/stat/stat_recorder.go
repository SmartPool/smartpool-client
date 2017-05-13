package stat

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
	"time"
)

const STATRECORDER_FILE string = "stat_recorder"

type StatRecorder struct {
	mu sync.RWMutex

	RigDatas map[string]*RigData
	FarmData *FarmData
}

func loadStatRecorder(storage smartpool.PersistentStorage) (*StatRecorder, error) {
	result := &StatRecorder{
		mu:       sync.RWMutex{},
		RigDatas: map[string]*RigData{},
		FarmData: NewFarmData(),
	}
	loadedStats, err := storage.Load(result, STATRECORDER_FILE)
	result = loadedStats.(*StatRecorder)
	return result, err
}

func NewStatRecorder(storage smartpool.PersistentStorage) *StatRecorder {
	statRecorder, err := loadStatRecorder(storage)
	if err != nil {
		smartpool.Output.Printf("Couldn't load stat from last session. Initialize with fresh stat recorder.\n")
	}
	smartpool.Output.Printf("Stat persister is running...\n")
	go func(sr *StatRecorder, storage smartpool.PersistentStorage) {
		for {
			smartpool.Output.Printf("Truncating stat datas...\n")
			err := sr.truncateData(storage)
			if err == nil {
				smartpool.Output.Printf("Done truncating stat.\n")
			} else {
				smartpool.Output.Printf("Failed truncating stat. (%s)\n", err.Error())
			}
			time.Sleep(time.Minute)
		}
	}(statRecorder, storage)
	smartpool.Output.Printf("Stat truncator is running...\n")
	return statRecorder
}

func (sr *StatRecorder) Persist(storage smartpool.PersistentStorage) error {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	smartpool.Output.Printf("Saving stats to disk...")
	err := storage.Persist(sr, STATRECORDER_FILE)
	if err == nil {
		smartpool.Output.Printf("Done.\n")
	} else {
		smartpool.Output.Printf("Failed. (%s)\n", err.Error())
	}
	return err
}

func (sr *StatRecorder) truncateData(storage smartpool.PersistentStorage) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	var err error
	if err = sr.FarmData.TruncateData(storage); err != nil {
		return err
	}
	for _, rigData := range sr.RigDatas {
		if err = rigData.TruncateData(storage); err != nil {
			return err
		}
	}
	return nil
}

func (sr *StatRecorder) getRigData(rig smartpool.Rig) *RigData {
	data := sr.RigDatas[rig.ID()]
	if data == nil {
		data = NewRigData(rig.ID())
		sr.RigDatas[rig.ID()] = data
	}
	return data
}

func (sr *StatRecorder) ShareRestored(noShares uint64) {
	sr.mu.Lock()
	defer sr.mu.Unlock()
	sr.FarmData.ShareRestored(noShares)
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
	fmt.Printf("going to update rig hashrate:\n")
	fmt.Printf("rig: %v\n", rig)
	fmt.Printf("reported hashrate: %v\n", rigData.AverageReportedHashrate.Uint64())
	fmt.Printf("effective hashrate: %v\n", rigData.AverageEffectiveHashrate.Uint64())
	fmt.Printf("period reported hashrate: %v\n", rigData.PeriodReportedHashrate(t).Uint64())
	fmt.Printf("period effective hashrate: %v\n", rigData.PeriodEffectiveHashrate(t).Uint64())
	sr.FarmData.UpdateRigHashrate(
		rig, rigData.AverageReportedHashrate, rigData.AverageEffectiveHashrate,
		rigData.PeriodReportedHashrate(t), rigData.PeriodEffectiveHashrate(t),
		t,
	)
}

func (sr *StatRecorder) OverallFarmStat() interface{} {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	return sr.FarmData.OverallFarmData
}

func (sr *StatRecorder) FarmStat(start uint64, end uint64) interface{} {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	result := map[uint64]*PeriodFarmData{}
	for period, data := range sr.FarmData.Datas {
		if start <= period && period <= end {
			result[period] = data
		}
	}
	return result
}

func (sr *StatRecorder) OverallRigStat(rig smartpool.Rig) interface{} {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	rigData := sr.getRigData(rig)
	return rigData.OverallRigData
}

func (sr *StatRecorder) RigStat(rig smartpool.Rig, start, end uint64) interface{} {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	rigData := sr.getRigData(rig)
	result := map[uint64]*PeriodRigData{}
	for period, data := range rigData.Datas {
		if start <= period && period <= end {
			result[period] = data
		}
	}
	return result
}
