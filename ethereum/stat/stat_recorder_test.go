package stat

import (
	"fmt"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"sync"
	"testing"
	"time"
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
	period := TimeToPeriod(time.Now())
	periodStats := recorder.FarmStat(period, period).(map[uint64]*PeriodFarmData)
	periodData := periodStats[period]
	if periodData.ReportedHashrate.Uint64() != 550 {
		fmt.Printf("got: %d, expected: %d\n", periodData.ReportedHashrate.Uint64(), 550)
		t.Fail()
	}
	rig2Stat := recorder.OverallRigStat(rig2).(*OverallRigData)
	if rig2Stat.AverageReportedHashrate.Uint64() != 400 {
		t.Fail()
	}
	periodRigStats := recorder.RigStat(rig2, period, period).(map[uint64]*PeriodRigData)
	periodRigData := periodRigStats[period]
	if periodRigData.AverageReportedHashrate.Uint64() != 400 {
		t.Fail()
	}
}
