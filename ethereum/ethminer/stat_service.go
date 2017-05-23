package ethminer

import (
	"encoding/json"
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/SmartPool/smartpool-client/ethereum"
	"github.com/SmartPool/smartpool-client/ethereum/stat"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type StatService struct{}

func farmStat() map[string]interface{} {
	t := time.Now()
	curPeriod := stat.TimeToPeriod(t)
	shortWindow := stat.TimeToPeriod(t.Add(-time.Duration(stat.ShortWindow) * time.Second))
	longWindow := stat.TimeToPeriod(t.Add(-time.Duration(stat.LongWindow) * time.Second))
	overall := SmartPool.StatRecorder.OverallFarmStat()
	shortWindowStat := SmartPool.StatRecorder.FarmStat(shortWindow, curPeriod)
	longWindowStat := SmartPool.StatRecorder.FarmStat(longWindow, curPeriod)
	return map[string]interface{}{
		"overall":               overall,
		"short_window_sample":   shortWindowStat,
		"short_window_duration": stat.ShortWindow,
		"long_window_sample":    longWindowStat,
		"long_window_duration":  stat.LongWindow,
		"period_duration":       stat.BaseTimePeriod,
	}
}

func rigStat(rig smartpool.Rig) map[string]interface{} {
	t := time.Now()
	curPeriod := stat.TimeToPeriod(t)
	shortWindow := stat.TimeToPeriod(t.Add(-time.Duration(stat.ShortWindow) * time.Second))
	longWindow := stat.TimeToPeriod(t.Add(-time.Duration(stat.LongWindow) * time.Second))
	overall := SmartPool.StatRecorder.OverallRigStat(rig)
	shortWindowStat := SmartPool.StatRecorder.RigStat(rig, shortWindow, curPeriod)
	longWindowStat := SmartPool.StatRecorder.RigStat(rig, longWindow, curPeriod)
	return map[string]interface{}{
		"overall":               overall,
		"short_window_sample":   shortWindowStat,
		"short_window_duration": stat.ShortWindow,
		"long_window_sample":    longWindowStat,
		"long_window_duration":  stat.LongWindow,
		"period_duration":       stat.BaseTimePeriod,
	}
}

func (server *StatService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	method := r.URL.Query().Get(":method")
	fmt.Printf("method: %s\n", method)
	if method == "json" {
		scope := r.URL.Query().Get(":scope")
		if scope == "farm" {
			result := farmStat()
			encoder := json.NewEncoder(w)
			encoder.Encode(&result)
		} else if scope == "rig" {
			rigID := r.URL.Query().Get(":rig")
			rig := ethereum.NewRig(rigID)
			result := rigStat(rig)
			encoder := json.NewEncoder(w)
			encoder.Encode(&result)
		} else {
			http.Error(w, "Only /json/farm and /json/rig/:id are supported", 404)
		}
	} else if method == "ws" {
		if r.Header.Get("Origin") != "http://"+r.Host {
			http.Error(w, "Origin not allowed", 403)
			return
		}
		conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}
		defer conn.Close()
		server.handleMessages(conn)
	} else {
		http.Error(w, "Only /farm and /rig/:id are supported", 404)
	}
}

func (server *StatService) handleMessages(conn *websocket.Conn) {
	startTime := time.Now()
	for {
		if time.Since(startTime).Seconds() > 600 {
			break
		}
		m := make(map[string]string)
		err := conn.ReadJSON(&m)
		if err == nil {
			startTime = time.Now()
			if m["action"] == "getFarmInfo" {
				conn.WriteJSON(farmStat())
			} else if m["action"] == "getRigInfo" {
				rigID := m["rigId"]
				rig := ethereum.NewRig(rigID)
				conn.WriteJSON(rigStat(rig))
			}
		} else {
			break
		}
	}
}

func NewStatService() *StatService {
	return &StatService{}
}
