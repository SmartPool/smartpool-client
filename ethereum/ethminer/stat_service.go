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
	anHourAgo := stat.TimeToPeriod(t.Add(-time.Hour))
	threeHourAgo := stat.TimeToPeriod(t.Add(-3 * time.Hour))
	if method == "farm" {
		overall := SmartPool.StatRecorder.OverallFarmStat()
		hourAgo := SmartPool.StatRecorder.FarmStat(anHourAgo, curPeriod)
		threeHourAgo := SmartPool.StatRecorder.FarmStat(threeHourAgo, curPeriod)
		result := map[string]interface{}{
			"overall":         overall,
			"last_1_hour":     hourAgo,
			"last_3_hours":    threeHourAgo,
			"period_duration": stat.BaseTimePeriod,
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(&result)
	} else if method == "rig" {
		rigID := r.URL.Query().Get(":rig")
		rig := ethereum.NewRig(rigID)
		overall := SmartPool.StatRecorder.OverallRigStat(rig)
		hourAgo := SmartPool.StatRecorder.RigStat(rig, anHourAgo, curPeriod)
		threeHourAgo := SmartPool.StatRecorder.RigStat(rig, threeHourAgo, curPeriod)
		result := map[string]interface{}{
			"overall":         overall,
			"last_1_hour":     hourAgo,
			"last_3_hours":    threeHourAgo,
			"period_duration": stat.BaseTimePeriod,
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
