package stat

import (
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
	"testing"
)

var rig *ethereum.Rig = ethereum.NewRig("testrig")
var rig2 *ethereum.Rig = ethereum.NewRig("anotherrig")

func newStatRecorder() *StatRecorder {
	return &StatRecorder{
		mu:       sync.RWMutex{},
		RigDatas: map[string]*RigData{},
		FarmData: NewFarmData(),
	}
}

func TestRecordHashrate(t *testing.T) {
	recorder := newStatRecorder()
	recorder.RecordHashrate(
		hexutil.Uint64(100),
		common.Hash{},
		rig,
	)
	recorder.RecordHashrate(
		hexutil.Uint64(200),
		common.Hash{},
		rig,
	)
	overall := recorder.OverallFarmStat().(*OverallFarmData)
	if overall.ReportedHashrate.Uint64() != 150 {
		t.Fail()
	}
	if overall.Rigs[rig.ID()].ReportedHashrate.Uint64() != 150 {
		t.Fail()
	}
	rig := recorder.OverallRigStat(rig).(*OverallRigData)
	if rig.AverageReportedHashrate.Uint64() != 150 {
		t.Fail()
	}
}

func TestRecordHashrateFromMultipleRigs(t *testing.T) {
	recorder := newStatRecorder()
	recorder.RecordHashrate(hexutil.Uint64(100), common.Hash{}, rig)
	recorder.RecordHashrate(hexutil.Uint64(200), common.Hash{}, rig)
	recorder.RecordHashrate(hexutil.Uint64(500), common.Hash{}, rig2)
	recorder.RecordHashrate(hexutil.Uint64(300), common.Hash{}, rig2)
	overall := recorder.OverallFarmStat().(*OverallFarmData)
	if overall.ReportedHashrate.Uint64() != 550 {
		t.Fail()
	}
	if overall.Rigs[rig2.ID()].ReportedHashrate.Uint64() != 400 {
		t.Fail()
	}
	rig2 := recorder.OverallRigStat(rig2).(*OverallRigData)
	if rig2.AverageReportedHashrate.Uint64() != 400 {
		t.Fail()
	}
}
