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
		scope := r.URL.Query().Get(":scope")
		wsQuit := make(chan struct{})
		fmt.Printf("scope: %s\n", scope)
		if scope == "farm" {
			server.handleWSForFarm(conn, wsQuit)
		} else if scope == "rig" {
			rigID := r.URL.Query().Get(":rig")
			rig := ethereum.NewRig(rigID)
			server.handleWSForRig(conn, rig, wsQuit)
		}
	} else {
		http.Error(w, "Only /farm and /rig/:id are supported", 404)
	}
}

func (server *StatService) handleWSForFarm(conn *websocket.Conn, quit chan struct{}) {
	conn.WriteJSON(farmStat())
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				conn.WriteJSON(farmStat())
			case <-quit:
				ticker.Stop()
				break
			}
		}
	}()
}

func (server *StatService) handleWSForRig(conn *websocket.Conn, rig smartpool.Rig, quit chan struct{}) {
	conn.WriteJSON(rigStat(rig))
	ticker := time.NewTicker(10 * time.Second)
	fmt.Printf("handle for rig\n")
	go func(rig smartpool.Rig) {
		for {
			select {
			case <-ticker.C:
				conn.WriteJSON(rigStat(rig))
			case <-quit:
				ticker.Stop()
				break
			}
		}
	}(rig)
}

func NewStatService() *StatService {
	return &StatService{}
}
