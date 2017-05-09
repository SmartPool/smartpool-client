package ethminer

import (
	"encoding/json"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/SmartPool/smartpool-client/ethereum/stat"
	"net/http"
	"time"
)

type StatService struct{}

func (server *StatService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	method := r.URL.Query().Get(":method")
	t := time.Now()
	curPeriod := stat.TimeToPeriod(t)
	shortWindow := stat.TimeToPeriod(t.Add(-time.Duration(stat.ShortWindow) * time.Second))
	longWindow := stat.TimeToPeriod(t.Add(-time.Duration(stat.LongWindow) * time.Second))
	if method == "farm" {
		overall := SmartPool.StatRecorder.OverallFarmStat()
		shortWindowStat := SmartPool.StatRecorder.FarmStat(shortWindow, curPeriod)
		longWindowStat := SmartPool.StatRecorder.FarmStat(longWindow, curPeriod)
		result := map[string]interface{}{
			"overall":               overall,
			"short_window_sample":   shortWindowStat,
			"short_window_duration": stat.ShortWindow,
			"long_window_sample":    longWindowStat,
			"long_window_duration":  stat.LongWindow,
			"period_duration":       stat.BaseTimePeriod,
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(&result)
	} else if method == "rig" {
		rigID := r.URL.Query().Get(":rig")
		rig := ethereum.NewRig(rigID)
		overall := SmartPool.StatRecorder.OverallRigStat(rig)
		shortWindowStat := SmartPool.StatRecorder.RigStat(rig, shortWindow, curPeriod)
		longWindowStat := SmartPool.StatRecorder.RigStat(rig, longWindow, curPeriod)
		result := map[string]interface{}{
			"overall":               overall,
			"short_window_sample":   shortWindowStat,
			"short_window_duration": stat.ShortWindow,
			"long_window_sample":    longWindowStat,
			"long_window_duration":  stat.LongWindow,
			"period_duration":       stat.BaseTimePeriod,
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(&result)
	} else {
		http.Error(w, "Only /farm and /rig/:id are supported", 404)
	}
}

func NewStatService() *StatService {
	return &StatService{}
}
