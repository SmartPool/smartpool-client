package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func main() {

	fileHandler := http.StripPrefix("/stats/", http.FileServer(http.Dir("./statistic")))
	http.Handle("/stats/", fileHandler)

	//http.Handle("/stats", http.FileServer(http.Dir("./statistic")))
	http.HandleFunc("/ws/farm", wsHandlerFarm)
	http.HandleFunc("/ws/rig/his7chan", wsHandlerRig)
	http.HandleFunc("/api/ethminer/advanceinfo", getAdvanceInfo)
	//http.HandleFunc("/api/ethminer/configinfo", getConfigInfo)
	http.HandleFunc("/status", getConfigInfo)
	http.ListenAndServe(":8082", nil)
}

func getAdvanceInfo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["x"] = "1"
	response["y"] = "2"
	reponseJson, _ := json.Marshal(response)
	fmt.Fprintf(w, "%s", reponseJson)
}
func getConfigInfo(w http.ResponseWriter, r *http.Request) {
	responseJson := `{
					  "rpc_endpoint": "http://localhost:8545",
					  "share_threshold": 20,
					  "share_difficulty": 123456,
					  "contract_address": "0x9af93376af1ddd22fa2e94fd0a030b3dea96bb96",
					  "miner_address": "0xe034afdcc2ba0441ff215ee9ba0da3e86450108d",
					  "extra_data": "SmartPool-NsjdZFWvUone7w00000000",
					  "hotstop_mode": true
					}`
	fmt.Fprintf(w, "%s", responseJson)
}
func wsHandlerFarm(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	farmData := getData()
	conn.WriteMessage(1, []byte(farmData))
	//conn.WriteJSON(farmData)
	ticker := time.NewTicker(1000 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				farmData = getData()
				conn.WriteMessage(1, []byte(farmData))
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
func wsHandlerRig(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Origin") != "http://"+r.Host {
		http.Error(w, "Origin not allowed", 403)
		return
	}
	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}
	farmData := getData()
	conn.WriteMessage(1, []byte(farmData))
	ticker := time.NewTicker(1000 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				farmData = getData()
				conn.WriteMessage(1, []byte(farmData))
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func getData() string {
	return `
	{
  "long_window_duration": 86400,
  "long_window_sample": {
    "2491275": {
      "mined_share": 19,
      "valid_share": 17,
      "total_valid_difficulty": 74460000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 2,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 371110830,
      "effective_hashrate": 124100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169054112
        },
        "pcs7chan": {
          "ReportedHashrate": 202056718
        }
      },
      "block_found": 0,
      "time_period": 2491275,
      "start_time": "2017-05-14T05:34:36.6338567-07:00"
    },
    "2491276": {
      "mined_share": 86,
      "valid_share": 86,
      "total_valid_difficulty": 376680000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 904721339,
      "effective_hashrate": 627800000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 168914210
        },
        "his7chan": {
          "ReportedHashrate": 188344526
        },
        "pcs7chan": {
          "ReportedHashrate": 201960487
        },
        "pcsl1": {
          "ReportedHashrate": 172561287
        },
        "pcsl2": {
          "ReportedHashrate": 172940829
        }
      },
      "block_found": 0,
      "time_period": 2491276,
      "start_time": "2017-05-14T05:40:03.3741025-07:00"
    },
    "2491277": {
      "mined_share": 148,
      "valid_share": 148,
      "total_valid_difficulty": 648240000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1064103500,
      "effective_hashrate": 1080400000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169418557
        },
        "his7chan": {
          "ReportedHashrate": 189052686
        },
        "pcs1": {
          "ReportedHashrate": 162652406
        },
        "pcs7chan": {
          "ReportedHashrate": 196484794
        },
        "pcsl1": {
          "ReportedHashrate": 173276645
        },
        "pcsl2": {
          "ReportedHashrate": 173218412
        }
      },
      "block_found": 0,
      "time_period": 2491277,
      "start_time": "2017-05-14T05:50:00.1683215-07:00"
    },
    "2491278": {
      "mined_share": 153,
      "valid_share": 153,
      "total_valid_difficulty": 670140000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1070107892,
      "effective_hashrate": 1116900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169243393
        },
        "his7chan": {
          "ReportedHashrate": 189153339
        },
        "pcs1": {
          "ReportedHashrate": 163057151
        },
        "pcs7chan": {
          "ReportedHashrate": 201980485
        },
        "pcsl1": {
          "ReportedHashrate": 173220260
        },
        "pcsl2": {
          "ReportedHashrate": 173453264
        }
      },
      "block_found": 0,
      "time_period": 2491278,
      "start_time": "2017-05-14T06:00:00.5049633-07:00"
    },
    "2491279": {
      "mined_share": 115,
      "valid_share": 115,
      "total_valid_difficulty": 503700000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1064800219,
      "effective_hashrate": 839500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169508513
        },
        "his7chan": {
          "ReportedHashrate": 189179085
        },
        "pcs1": {
          "ReportedHashrate": 163031724
        },
        "pcs7chan": {
          "ReportedHashrate": 196381333
        },
        "pcsl1": {
          "ReportedHashrate": 173327890
        },
        "pcsl2": {
          "ReportedHashrate": 173371674
        }
      },
      "block_found": 0,
      "time_period": 2491279,
      "start_time": "2017-05-14T06:10:00.469375-07:00"
    },
    "2491280": {
      "mined_share": 50,
      "valid_share": 50,
      "total_valid_difficulty": 219000000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1023457417,
      "effective_hashrate": 365000000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 168894185
        },
        "his7chan": {
          "ReportedHashrate": 189585294
        },
        "pcs1": {
          "ReportedHashrate": 161340954
        },
        "pcs7chan": {
          "ReportedHashrate": 202348014
        },
        "pcsl1": {
          "ReportedHashrate": 173477648
        },
        "pcsl2": {
          "ReportedHashrate": 127811322
        }
      },
      "block_found": 0,
      "time_period": 2491280,
      "start_time": "2017-05-14T06:21:35.5249773-07:00"
    },
    "2491281": {
      "mined_share": 126,
      "valid_share": 126,
      "total_valid_difficulty": 551880000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897667956,
      "effective_hashrate": 919800000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169469538
        },
        "his7chan": {
          "ReportedHashrate": 189348989
        },
        "pcs1": {
          "ReportedHashrate": 163435937
        },
        "pcs7chan": {
          "ReportedHashrate": 202034382
        },
        "pcsl1": {
          "ReportedHashrate": 173379110
        }
      },
      "block_found": 0,
      "time_period": 2491281,
      "start_time": "2017-05-14T06:30:00.0108894-07:00"
    },
    "2491282": {
      "mined_share": 117,
      "valid_share": 117,
      "total_valid_difficulty": 512460000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 886726048,
      "effective_hashrate": 854100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169146274
        },
        "his7chan": {
          "ReportedHashrate": 189213376
        },
        "pcs1": {
          "ReportedHashrate": 163496572
        },
        "pcs7chan": {
          "ReportedHashrate": 191603996
        },
        "pcsl1": {
          "ReportedHashrate": 173265830
        }
      },
      "block_found": 0,
      "time_period": 2491282,
      "start_time": "2017-05-14T06:40:00.8399459-07:00"
    },
    "2491283": {
      "mined_share": 108,
      "valid_share": 108,
      "total_valid_difficulty": 473040000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897369859,
      "effective_hashrate": 788400000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169402875
        },
        "his7chan": {
          "ReportedHashrate": 189287419
        },
        "pcs1": {
          "ReportedHashrate": 163296876
        },
        "pcs7chan": {
          "ReportedHashrate": 202150209
        },
        "pcsl1": {
          "ReportedHashrate": 173232480
        }
      },
      "block_found": 0,
      "time_period": 2491283,
      "start_time": "2017-05-14T06:50:00.911515-07:00"
    },
    "2491284": {
      "mined_share": 102,
      "valid_share": 102,
      "total_valid_difficulty": 446760000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897377864,
      "effective_hashrate": 744600000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169550510
        },
        "his7chan": {
          "ReportedHashrate": 189341019
        },
        "pcs1": {
          "ReportedHashrate": 163045344
        },
        "pcs7chan": {
          "ReportedHashrate": 202060837
        },
        "pcsl1": {
          "ReportedHashrate": 173380154
        }
      },
      "block_found": 0,
      "time_period": 2491284,
      "start_time": "2017-05-14T07:00:00.0717704-07:00"
    },
    "2491285": {
      "mined_share": 127,
      "valid_share": 127,
      "total_valid_difficulty": 556260000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897006767,
      "effective_hashrate": 927100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169031475
        },
        "his7chan": {
          "ReportedHashrate": 189268182
        },
        "pcs1": {
          "ReportedHashrate": 163320972
        },
        "pcs7chan": {
          "ReportedHashrate": 202028961
        },
        "pcsl1": {
          "ReportedHashrate": 173357177
        }
      },
      "block_found": 0,
      "time_period": 2491285,
      "start_time": "2017-05-14T07:10:00.7344702-07:00"
    },
    "2491286": {
      "mined_share": 117,
      "valid_share": 117,
      "total_valid_difficulty": 512460000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897743922,
      "effective_hashrate": 854100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169568264
        },
        "his7chan": {
          "ReportedHashrate": 189419115
        },
        "pcs1": {
          "ReportedHashrate": 163560809
        },
        "pcs7chan": {
          "ReportedHashrate": 201921594
        },
        "pcsl1": {
          "ReportedHashrate": 173274140
        }
      },
      "block_found": 0,
      "time_period": 2491286,
      "start_time": "2017-05-14T07:20:01.1355951-07:00"
    },
    "2491287": {
      "mined_share": 137,
      "valid_share": 137,
      "total_valid_difficulty": 600060000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897453774,
      "effective_hashrate": 1000100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169663641
        },
        "his7chan": {
          "ReportedHashrate": 189244291
        },
        "pcs1": {
          "ReportedHashrate": 163346632
        },
        "pcs7chan": {
          "ReportedHashrate": 201912983
        },
        "pcsl1": {
          "ReportedHashrate": 173286227
        }
      },
      "block_found": 0,
      "time_period": 2491287,
      "start_time": "2017-05-14T07:30:03.0682818-07:00"
    },
    "2491288": {
      "mined_share": 105,
      "valid_share": 105,
      "total_valid_difficulty": 459900000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897294246,
      "effective_hashrate": 766500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169332580
        },
        "his7chan": {
          "ReportedHashrate": 189327946
        },
        "pcs1": {
          "ReportedHashrate": 163310949
        },
        "pcs7chan": {
          "ReportedHashrate": 201887082
        },
        "pcsl1": {
          "ReportedHashrate": 173435689
        }
      },
      "block_found": 0,
      "time_period": 2491288,
      "start_time": "2017-05-14T07:40:00.2923899-07:00"
    },
    "2491289": {
      "mined_share": 104,
      "valid_share": 104,
      "total_valid_difficulty": 455520000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897088893,
      "effective_hashrate": 759200000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169258715
        },
        "his7chan": {
          "ReportedHashrate": 189352568
        },
        "pcs1": {
          "ReportedHashrate": 163256268
        },
        "pcs7chan": {
          "ReportedHashrate": 201984534
        },
        "pcsl1": {
          "ReportedHashrate": 173236808
        }
      },
      "block_found": 0,
      "time_period": 2491289,
      "start_time": "2017-05-14T07:50:00.9078007-07:00"
    },
    "2491290": {
      "mined_share": 127,
      "valid_share": 127,
      "total_valid_difficulty": 556260000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 896939748,
      "effective_hashrate": 927100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169169276
        },
        "his7chan": {
          "ReportedHashrate": 189226934
        },
        "pcs1": {
          "ReportedHashrate": 163419154
        },
        "pcs7chan": {
          "ReportedHashrate": 201882526
        },
        "pcsl1": {
          "ReportedHashrate": 173241858
        }
      },
      "block_found": 0,
      "time_period": 2491290,
      "start_time": "2017-05-14T08:00:00.1156455-07:00"
    },
    "2491291": {
      "mined_share": 113,
      "valid_share": 113,
      "total_valid_difficulty": 494940000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897441757,
      "effective_hashrate": 824900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169341947
        },
        "his7chan": {
          "ReportedHashrate": 189295663
        },
        "pcs1": {
          "ReportedHashrate": 163616788
        },
        "pcs7chan": {
          "ReportedHashrate": 201815774
        },
        "pcsl1": {
          "ReportedHashrate": 173371585
        }
      },
      "block_found": 0,
      "time_period": 2491291,
      "start_time": "2017-05-14T08:10:00.3422497-07:00"
    },
    "2491292": {
      "mined_share": 121,
      "valid_share": 121,
      "total_valid_difficulty": 529980000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897260068,
      "effective_hashrate": 883300000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169556881
        },
        "his7chan": {
          "ReportedHashrate": 189123223
        },
        "pcs1": {
          "ReportedHashrate": 163291019
        },
        "pcs7chan": {
          "ReportedHashrate": 202013328
        },
        "pcsl1": {
          "ReportedHashrate": 173275617
        }
      },
      "block_found": 0,
      "time_period": 2491292,
      "start_time": "2017-05-14T08:20:01.8030558-07:00"
    },
    "2491293": {
      "mined_share": 129,
      "valid_share": 129,
      "total_valid_difficulty": 565020000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897532047,
      "effective_hashrate": 941700000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169450288
        },
        "his7chan": {
          "ReportedHashrate": 189308065
        },
        "pcs1": {
          "ReportedHashrate": 163525696
        },
        "pcs7chan": {
          "ReportedHashrate": 201887452
        },
        "pcsl1": {
          "ReportedHashrate": 173360546
        }
      },
      "block_found": 0,
      "time_period": 2491293,
      "start_time": "2017-05-14T08:30:00.8295219-07:00"
    },
    "2491294": {
      "mined_share": 95,
      "valid_share": 95,
      "total_valid_difficulty": 416100000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897350338,
      "effective_hashrate": 693500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169563523
        },
        "his7chan": {
          "ReportedHashrate": 189078600
        },
        "pcs1": {
          "ReportedHashrate": 163334462
        },
        "pcs7chan": {
          "ReportedHashrate": 201965746
        },
        "pcsl1": {
          "ReportedHashrate": 173408007
        }
      },
      "block_found": 0,
      "time_period": 2491294,
      "start_time": "2017-05-14T08:40:00.1863677-07:00"
    },
    "2491295": {
      "mined_share": 130,
      "valid_share": 130,
      "total_valid_difficulty": 569400000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897259192,
      "effective_hashrate": 949000000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169378478
        },
        "his7chan": {
          "ReportedHashrate": 189181616
        },
        "pcs1": {
          "ReportedHashrate": 163402852
        },
        "pcs7chan": {
          "ReportedHashrate": 202102995
        },
        "pcsl1": {
          "ReportedHashrate": 173193251
        }
      },
      "block_found": 0,
      "time_period": 2491295,
      "start_time": "2017-05-14T08:50:02.0962651-07:00"
    },
    "2491296": {
      "mined_share": 116,
      "valid_share": 116,
      "total_valid_difficulty": 508080000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897706279,
      "effective_hashrate": 846800000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169652624
        },
        "his7chan": {
          "ReportedHashrate": 189226228
        },
        "pcs1": {
          "ReportedHashrate": 163410522
        },
        "pcs7chan": {
          "ReportedHashrate": 202057845
        },
        "pcsl1": {
          "ReportedHashrate": 173359060
        }
      },
      "block_found": 0,
      "time_period": 2491296,
      "start_time": "2017-05-14T09:00:00.2098482-07:00"
    },
    "2491297": {
      "mined_share": 104,
      "valid_share": 104,
      "total_valid_difficulty": 455520000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897747713,
      "effective_hashrate": 759200000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169662002
        },
        "his7chan": {
          "ReportedHashrate": 189288691
        },
        "pcs1": {
          "ReportedHashrate": 163484254
        },
        "pcs7chan": {
          "ReportedHashrate": 202032781
        },
        "pcsl1": {
          "ReportedHashrate": 173279985
        }
      },
      "block_found": 0,
      "time_period": 2491297,
      "start_time": "2017-05-14T09:10:00.7581758-07:00"
    },
    "2491298": {
      "mined_share": 135,
      "valid_share": 135,
      "total_valid_difficulty": 591300000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897685090,
      "effective_hashrate": 985500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169650384
        },
        "his7chan": {
          "ReportedHashrate": 189292923
        },
        "pcs1": {
          "ReportedHashrate": 163515344
        },
        "pcs7chan": {
          "ReportedHashrate": 201939374
        },
        "pcsl1": {
          "ReportedHashrate": 173287065
        }
      },
      "block_found": 0,
      "time_period": 2491298,
      "start_time": "2017-05-14T09:20:00.5284705-07:00"
    },
    "2491299": {
      "mined_share": 119,
      "valid_share": 119,
      "total_valid_difficulty": 521220000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897133814,
      "effective_hashrate": 868700000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169186928
        },
        "his7chan": {
          "ReportedHashrate": 189364594
        },
        "pcs1": {
          "ReportedHashrate": 163523242
        },
        "pcs7chan": {
          "ReportedHashrate": 201819260
        },
        "pcsl1": {
          "ReportedHashrate": 173239790
        }
      },
      "block_found": 0,
      "time_period": 2491299,
      "start_time": "2017-05-14T09:30:00.624083-07:00"
    },
    "2491300": {
      "mined_share": 114,
      "valid_share": 114,
      "total_valid_difficulty": 499320000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897511057,
      "effective_hashrate": 832200000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169457394
        },
        "his7chan": {
          "ReportedHashrate": 189434411
        },
        "pcs1": {
          "ReportedHashrate": 163356600
        },
        "pcs7chan": {
          "ReportedHashrate": 202023359
        },
        "pcsl1": {
          "ReportedHashrate": 173239293
        }
      },
      "block_found": 0,
      "time_period": 2491300,
      "start_time": "2017-05-14T09:40:00.7674768-07:00"
    },
    "2491301": {
      "mined_share": 135,
      "valid_share": 135,
      "total_valid_difficulty": 591300000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 1,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 896927015,
      "effective_hashrate": 985500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169189709
        },
        "his7chan": {
          "ReportedHashrate": 189191471
        },
        "pcs1": {
          "ReportedHashrate": 163237116
        },
        "pcs7chan": {
          "ReportedHashrate": 202020392
        },
        "pcsl1": {
          "ReportedHashrate": 173288327
        }
      },
      "block_found": 0,
      "time_period": 2491301,
      "start_time": "2017-05-14T09:50:00.2812658-07:00"
    },
    "2491302": {
      "mined_share": 128,
      "valid_share": 128,
      "total_valid_difficulty": 560640000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897495634,
      "effective_hashrate": 934400000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169457560
        },
        "his7chan": {
          "ReportedHashrate": 189189435
        },
        "pcs1": {
          "ReportedHashrate": 163509006
        },
        "pcs7chan": {
          "ReportedHashrate": 202036636
        },
        "pcsl1": {
          "ReportedHashrate": 173302997
        }
      },
      "block_found": 0,
      "time_period": 2491302,
      "start_time": "2017-05-14T10:00:00.3695382-07:00"
    },
    "2491303": {
      "mined_share": 58,
      "valid_share": 56,
      "total_valid_difficulty": 245280000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 2,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 896800598,
      "effective_hashrate": 408800000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169370453
        },
        "his7chan": {
          "ReportedHashrate": 189253508
        },
        "pcs1": {
          "ReportedHashrate": 163056876
        },
        "pcs7chan": {
          "ReportedHashrate": 201831613
        },
        "pcsl1": {
          "ReportedHashrate": 173288148
        }
      },
      "block_found": 0,
      "time_period": 2491303,
      "start_time": "2017-05-14T10:10:01.5026609-07:00"
    },
    "2491304": {
      "mined_share": 123,
      "valid_share": 123,
      "total_valid_difficulty": 538740000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 896816573,
      "effective_hashrate": 897900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169070121
        },
        "his7chan": {
          "ReportedHashrate": 189302602
        },
        "pcs1": {
          "ReportedHashrate": 163356667
        },
        "pcs7chan": {
          "ReportedHashrate": 201828913
        },
        "pcsl1": {
          "ReportedHashrate": 173258270
        }
      },
      "block_found": 0,
      "time_period": 2491304,
      "start_time": "2017-05-14T10:20:01.6304922-07:00"
    },
    "2491305": {
      "mined_share": 142,
      "valid_share": 142,
      "total_valid_difficulty": 621960000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897347239,
      "effective_hashrate": 1036600000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169347069
        },
        "his7chan": {
          "ReportedHashrate": 189360648
        },
        "pcs1": {
          "ReportedHashrate": 163573194
        },
        "pcs7chan": {
          "ReportedHashrate": 201881184
        },
        "pcsl1": {
          "ReportedHashrate": 173185144
        }
      },
      "block_found": 0,
      "time_period": 2491305,
      "start_time": "2017-05-14T10:30:02.7795112-07:00"
    },
    "2491306": {
      "mined_share": 96,
      "valid_share": 96,
      "total_valid_difficulty": 420480000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897113970,
      "effective_hashrate": 700800000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 168947838
        },
        "his7chan": {
          "ReportedHashrate": 189263900
        },
        "pcs1": {
          "ReportedHashrate": 163518353
        },
        "pcs7chan": {
          "ReportedHashrate": 201977627
        },
        "pcsl1": {
          "ReportedHashrate": 173406252
        }
      },
      "block_found": 0,
      "time_period": 2491306,
      "start_time": "2017-05-14T10:40:00.4299116-07:00"
    },
    "2491307": {
      "mined_share": 137,
      "valid_share": 137,
      "total_valid_difficulty": 600060000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897352126,
      "effective_hashrate": 1000100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169378553
        },
        "his7chan": {
          "ReportedHashrate": 189239437
        },
        "pcs1": {
          "ReportedHashrate": 163476402
        },
        "pcs7chan": {
          "ReportedHashrate": 201919641
        },
        "pcsl1": {
          "ReportedHashrate": 173338093
        }
      },
      "block_found": 0,
      "time_period": 2491307,
      "start_time": "2017-05-14T10:50:01.4113846-07:00"
    },
    "2491308": {
      "mined_share": 102,
      "valid_share": 102,
      "total_valid_difficulty": 446760000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897350985,
      "effective_hashrate": 744600000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169355190
        },
        "his7chan": {
          "ReportedHashrate": 189146421
        },
        "pcs1": {
          "ReportedHashrate": 163486022
        },
        "pcs7chan": {
          "ReportedHashrate": 202008840
        },
        "pcsl1": {
          "ReportedHashrate": 173354512
        }
      },
      "block_found": 0,
      "time_period": 2491308,
      "start_time": "2017-05-14T11:00:01.7082426-07:00"
    },
    "2491309": {
      "mined_share": 122,
      "valid_share": 122,
      "total_valid_difficulty": 534360000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897257772,
      "effective_hashrate": 890600000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169274863
        },
        "his7chan": {
          "ReportedHashrate": 189329193
        },
        "pcs1": {
          "ReportedHashrate": 163504159
        },
        "pcs7chan": {
          "ReportedHashrate": 201981175
        },
        "pcsl1": {
          "ReportedHashrate": 173168382
        }
      },
      "block_found": 0,
      "time_period": 2491309,
      "start_time": "2017-05-14T11:10:00.0643069-07:00"
    },
    "2491310": {
      "mined_share": 117,
      "valid_share": 117,
      "total_valid_difficulty": 512460000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897476669,
      "effective_hashrate": 854100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169527629
        },
        "his7chan": {
          "ReportedHashrate": 189280068
        },
        "pcs1": {
          "ReportedHashrate": 163488682
        },
        "pcs7chan": {
          "ReportedHashrate": 201854221
        },
        "pcsl1": {
          "ReportedHashrate": 173326069
        }
      },
      "block_found": 0,
      "time_period": 2491310,
      "start_time": "2017-05-14T11:20:01.5228792-07:00"
    },
    "2491311": {
      "mined_share": 128,
      "valid_share": 128,
      "total_valid_difficulty": 560640000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897239792,
      "effective_hashrate": 934400000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169565374
        },
        "his7chan": {
          "ReportedHashrate": 189116268
        },
        "pcs1": {
          "ReportedHashrate": 163429588
        },
        "pcs7chan": {
          "ReportedHashrate": 201914965
        },
        "pcsl1": {
          "ReportedHashrate": 173213597
        }
      },
      "block_found": 0,
      "time_period": 2491311,
      "start_time": "2017-05-14T11:30:01.80894-07:00"
    },
    "2491312": {
      "mined_share": 126,
      "valid_share": 126,
      "total_valid_difficulty": 551880000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897193772,
      "effective_hashrate": 919800000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169527889
        },
        "his7chan": {
          "ReportedHashrate": 189207214
        },
        "pcs1": {
          "ReportedHashrate": 163468120
        },
        "pcs7chan": {
          "ReportedHashrate": 201809907
        },
        "pcsl1": {
          "ReportedHashrate": 173180642
        }
      },
      "block_found": 0,
      "time_period": 2491312,
      "start_time": "2017-05-14T11:40:00.3896995-07:00"
    },
    "2491313": {
      "mined_share": 128,
      "valid_share": 128,
      "total_valid_difficulty": 560640000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897662559,
      "effective_hashrate": 934400000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169433199
        },
        "his7chan": {
          "ReportedHashrate": 189396002
        },
        "pcs1": {
          "ReportedHashrate": 163551584
        },
        "pcs7chan": {
          "ReportedHashrate": 201958702
        },
        "pcsl1": {
          "ReportedHashrate": 173323072
        }
      },
      "block_found": 0,
      "time_period": 2491313,
      "start_time": "2017-05-14T11:50:01.4530674-07:00"
    },
    "2491314": {
      "mined_share": 106,
      "valid_share": 106,
      "total_valid_difficulty": 464280000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897393703,
      "effective_hashrate": 773800000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169144647
        },
        "his7chan": {
          "ReportedHashrate": 189231182
        },
        "pcs1": {
          "ReportedHashrate": 163687812
        },
        "pcs7chan": {
          "ReportedHashrate": 201971455
        },
        "pcsl1": {
          "ReportedHashrate": 173358607
        }
      },
      "block_found": 0,
      "time_period": 2491314,
      "start_time": "2017-05-14T12:00:00.7405653-07:00"
    },
    "2491315": {
      "mined_share": 112,
      "valid_share": 112,
      "total_valid_difficulty": 490560000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897517203,
      "effective_hashrate": 817600000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 168951926
        },
        "his7chan": {
          "ReportedHashrate": 189433374
        },
        "pcs1": {
          "ReportedHashrate": 163688613
        },
        "pcs7chan": {
          "ReportedHashrate": 202039526
        },
        "pcsl1": {
          "ReportedHashrate": 173403764
        }
      },
      "block_found": 0,
      "time_period": 2491315,
      "start_time": "2017-05-14T12:10:01.7117029-07:00"
    },
    "2491316": {
      "mined_share": 125,
      "valid_share": 125,
      "total_valid_difficulty": 547500000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897368252,
      "effective_hashrate": 912500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169348351
        },
        "his7chan": {
          "ReportedHashrate": 189031835
        },
        "pcs1": {
          "ReportedHashrate": 163759798
        },
        "pcs7chan": {
          "ReportedHashrate": 201830913
        },
        "pcsl1": {
          "ReportedHashrate": 173397355
        }
      },
      "block_found": 0,
      "time_period": 2491316,
      "start_time": "2017-05-14T12:20:00.0138614-07:00"
    },
    "2491317": {
      "mined_share": 141,
      "valid_share": 141,
      "total_valid_difficulty": 617580000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897116090,
      "effective_hashrate": 1029300000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169345106
        },
        "his7chan": {
          "ReportedHashrate": 189017386
        },
        "pcs1": {
          "ReportedHashrate": 163559836
        },
        "pcs7chan": {
          "ReportedHashrate": 201920447
        },
        "pcsl1": {
          "ReportedHashrate": 173273315
        }
      },
      "block_found": 0,
      "time_period": 2491317,
      "start_time": "2017-05-14T12:30:02.0607608-07:00"
    },
    "2491318": {
      "mined_share": 127,
      "valid_share": 127,
      "total_valid_difficulty": 556260000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897055052,
      "effective_hashrate": 927100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169455528
        },
        "his7chan": {
          "ReportedHashrate": 189207664
        },
        "pcs1": {
          "ReportedHashrate": 163507182
        },
        "pcs7chan": {
          "ReportedHashrate": 201769944
        },
        "pcsl1": {
          "ReportedHashrate": 173114734
        }
      },
      "block_found": 0,
      "time_period": 2491318,
      "start_time": "2017-05-14T12:40:01.1488994-07:00"
    },
    "2491319": {
      "mined_share": 139,
      "valid_share": 139,
      "total_valid_difficulty": 608820000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897586345,
      "effective_hashrate": 1014700000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169456006
        },
        "his7chan": {
          "ReportedHashrate": 189134670
        },
        "pcs1": {
          "ReportedHashrate": 163521072
        },
        "pcs7chan": {
          "ReportedHashrate": 202042410
        },
        "pcsl1": {
          "ReportedHashrate": 173432187
        }
      },
      "block_found": 0,
      "time_period": 2491319,
      "start_time": "2017-05-14T12:50:00.7508398-07:00"
    },
    "2491320": {
      "mined_share": 109,
      "valid_share": 109,
      "total_valid_difficulty": 477420000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 1,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 898145823,
      "effective_hashrate": 795700000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169657033
        },
        "his7chan": {
          "ReportedHashrate": 189260040
        },
        "pcs1": {
          "ReportedHashrate": 163784053
        },
        "pcs7chan": {
          "ReportedHashrate": 202085145
        },
        "pcsl1": {
          "ReportedHashrate": 173359552
        }
      },
      "block_found": 0,
      "time_period": 2491320,
      "start_time": "2017-05-14T13:00:02.2500991-07:00"
    },
    "2491321": {
      "mined_share": 110,
      "valid_share": 110,
      "total_valid_difficulty": 481800000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897509606,
      "effective_hashrate": 803000000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169248465
        },
        "his7chan": {
          "ReportedHashrate": 189191346
        },
        "pcs1": {
          "ReportedHashrate": 163698147
        },
        "pcs7chan": {
          "ReportedHashrate": 202014721
        },
        "pcsl1": {
          "ReportedHashrate": 173356927
        }
      },
      "block_found": 0,
      "time_period": 2491321,
      "start_time": "2017-05-14T13:10:00.0862668-07:00"
    },
    "2491322": {
      "mined_share": 123,
      "valid_share": 123,
      "total_valid_difficulty": 538740000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 1,
      "rejected_claim": 0,
      "reported_hashrate": 898034959,
      "effective_hashrate": 897900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169642328
        },
        "his7chan": {
          "ReportedHashrate": 189251349
        },
        "pcs1": {
          "ReportedHashrate": 163693804
        },
        "pcs7chan": {
          "ReportedHashrate": 202030110
        },
        "pcsl1": {
          "ReportedHashrate": 173417368
        }
      },
      "block_found": 0,
      "time_period": 2491322,
      "start_time": "2017-05-14T13:20:01.9435924-07:00"
    },
    "2491323": {
      "mined_share": 107,
      "valid_share": 107,
      "total_valid_difficulty": 468660000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897478098,
      "effective_hashrate": 781100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169435641
        },
        "his7chan": {
          "ReportedHashrate": 189365313
        },
        "pcs1": {
          "ReportedHashrate": 163579094
        },
        "pcs7chan": {
          "ReportedHashrate": 201917153
        },
        "pcsl1": {
          "ReportedHashrate": 173180897
        }
      },
      "block_found": 0,
      "time_period": 2491323,
      "start_time": "2017-05-14T13:30:00.5698411-07:00"
    },
    "2491324": {
      "mined_share": 94,
      "valid_share": 94,
      "total_valid_difficulty": 411720000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897855918,
      "effective_hashrate": 686200000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169456162
        },
        "his7chan": {
          "ReportedHashrate": 189434111
        },
        "pcs1": {
          "ReportedHashrate": 163691477
        },
        "pcs7chan": {
          "ReportedHashrate": 202046668
        },
        "pcsl1": {
          "ReportedHashrate": 173227500
        }
      },
      "block_found": 0,
      "time_period": 2491324,
      "start_time": "2017-05-14T13:40:00.7030087-07:00"
    },
    "2491325": {
      "mined_share": 123,
      "valid_share": 123,
      "total_valid_difficulty": 538740000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897114089,
      "effective_hashrate": 897900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169500428
        },
        "his7chan": {
          "ReportedHashrate": 189188805
        },
        "pcs1": {
          "ReportedHashrate": 163260384
        },
        "pcs7chan": {
          "ReportedHashrate": 201925194
        },
        "pcsl1": {
          "ReportedHashrate": 173239278
        }
      },
      "block_found": 0,
      "time_period": 2491325,
      "start_time": "2017-05-14T13:50:00.3772711-07:00"
    },
    "2491326": {
      "mined_share": 104,
      "valid_share": 104,
      "total_valid_difficulty": 455520000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897849807,
      "effective_hashrate": 759200000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169771692
        },
        "his7chan": {
          "ReportedHashrate": 189264488
        },
        "pcs1": {
          "ReportedHashrate": 163587590
        },
        "pcs7chan": {
          "ReportedHashrate": 201972577
        },
        "pcsl1": {
          "ReportedHashrate": 173253460
        }
      },
      "block_found": 0,
      "time_period": 2491326,
      "start_time": "2017-05-14T14:00:00.1216439-07:00"
    },
    "2491327": {
      "mined_share": 122,
      "valid_share": 122,
      "total_valid_difficulty": 534360000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897807802,
      "effective_hashrate": 890600000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169461989
        },
        "his7chan": {
          "ReportedHashrate": 189287930
        },
        "pcs1": {
          "ReportedHashrate": 163661839
        },
        "pcs7chan": {
          "ReportedHashrate": 202068323
        },
        "pcsl1": {
          "ReportedHashrate": 173327721
        }
      },
      "block_found": 0,
      "time_period": 2491327,
      "start_time": "2017-05-14T14:10:00.6560589-07:00"
    },
    "2491328": {
      "mined_share": 120,
      "valid_share": 120,
      "total_valid_difficulty": 525600000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897796520,
      "effective_hashrate": 876000000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169678460
        },
        "his7chan": {
          "ReportedHashrate": 189224799
        },
        "pcs1": {
          "ReportedHashrate": 163689539
        },
        "pcs7chan": {
          "ReportedHashrate": 201906890
        },
        "pcsl1": {
          "ReportedHashrate": 173296832
        }
      },
      "block_found": 0,
      "time_period": 2491328,
      "start_time": "2017-05-14T14:20:03.2192219-07:00"
    },
    "2491329": {
      "mined_share": 115,
      "valid_share": 115,
      "total_valid_difficulty": 503700000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897906054,
      "effective_hashrate": 839500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169781362
        },
        "his7chan": {
          "ReportedHashrate": 189390330
        },
        "pcs1": {
          "ReportedHashrate": 163575400
        },
        "pcs7chan": {
          "ReportedHashrate": 201873098
        },
        "pcsl1": {
          "ReportedHashrate": 173285864
        }
      },
      "block_found": 0,
      "time_period": 2491329,
      "start_time": "2017-05-14T14:30:01.5027682-07:00"
    },
    "2491330": {
      "mined_share": 114,
      "valid_share": 114,
      "total_valid_difficulty": 499320000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897211324,
      "effective_hashrate": 832200000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169259867
        },
        "his7chan": {
          "ReportedHashrate": 189196107
        },
        "pcs1": {
          "ReportedHashrate": 163513273
        },
        "pcs7chan": {
          "ReportedHashrate": 201958630
        },
        "pcsl1": {
          "ReportedHashrate": 173283447
        }
      },
      "block_found": 0,
      "time_period": 2491330,
      "start_time": "2017-05-14T14:40:02.4126874-07:00"
    },
    "2491331": {
      "mined_share": 136,
      "valid_share": 136,
      "total_valid_difficulty": 595680000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897774144,
      "effective_hashrate": 992800000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169680231
        },
        "his7chan": {
          "ReportedHashrate": 189327584
        },
        "pcs1": {
          "ReportedHashrate": 163526491
        },
        "pcs7chan": {
          "ReportedHashrate": 201967054
        },
        "pcsl1": {
          "ReportedHashrate": 173272784
        }
      },
      "block_found": 0,
      "time_period": 2491331,
      "start_time": "2017-05-14T14:50:00.0599269-07:00"
    },
    "2491332": {
      "mined_share": 115,
      "valid_share": 115,
      "total_valid_difficulty": 503700000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897727505,
      "effective_hashrate": 839500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169519806
        },
        "his7chan": {
          "ReportedHashrate": 189227836
        },
        "pcs1": {
          "ReportedHashrate": 163715863
        },
        "pcs7chan": {
          "ReportedHashrate": 201964663
        },
        "pcsl1": {
          "ReportedHashrate": 173299337
        }
      },
      "block_found": 0,
      "time_period": 2491332,
      "start_time": "2017-05-14T15:00:02.0570815-07:00"
    },
    "2491333": {
      "mined_share": 119,
      "valid_share": 119,
      "total_valid_difficulty": 521220000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1070157991,
      "effective_hashrate": 868700000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169660790
        },
        "his7chan": {
          "ReportedHashrate": 189157239
        },
        "pcs1": {
          "ReportedHashrate": 163544838
        },
        "pcs7chan": {
          "ReportedHashrate": 201994988
        },
        "pcsl1": {
          "ReportedHashrate": 173216614
        },
        "pcsl2": {
          "ReportedHashrate": 172583522
        }
      },
      "block_found": 1,
      "time_period": 2491333,
      "start_time": "2017-05-14T15:10:00.5377516-07:00"
    },
    "2491334": {
      "mined_share": 142,
      "valid_share": 142,
      "total_valid_difficulty": 621960000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1069939035,
      "effective_hashrate": 1036600000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169080715
        },
        "his7chan": {
          "ReportedHashrate": 189139050
        },
        "pcs1": {
          "ReportedHashrate": 163249901
        },
        "pcs7chan": {
          "ReportedHashrate": 201904530
        },
        "pcsl1": {
          "ReportedHashrate": 173327263
        },
        "pcsl2": {
          "ReportedHashrate": 173237576
        }
      },
      "block_found": 0,
      "time_period": 2491334,
      "start_time": "2017-05-14T15:20:00.5731701-07:00"
    },
    "2491335": {
      "mined_share": 153,
      "valid_share": 153,
      "total_valid_difficulty": 670140000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1069583223,
      "effective_hashrate": 1116900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169564710
        },
        "his7chan": {
          "ReportedHashrate": 189238043
        },
        "pcs1": {
          "ReportedHashrate": 163556272
        },
        "pcs7chan": {
          "ReportedHashrate": 201988266
        },
        "pcsl1": {
          "ReportedHashrate": 173224023
        },
        "pcsl2": {
          "ReportedHashrate": 172011909
        }
      },
      "block_found": 0,
      "time_period": 2491335,
      "start_time": "2017-05-14T15:30:01.6595749-07:00"
    },
    "2491336": {
      "mined_share": 125,
      "valid_share": 125,
      "total_valid_difficulty": 547500000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1069767602,
      "effective_hashrate": 912500000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169661521
        },
        "his7chan": {
          "ReportedHashrate": 189310949
        },
        "pcs1": {
          "ReportedHashrate": 161985008
        },
        "pcs7chan": {
          "ReportedHashrate": 201928718
        },
        "pcsl1": {
          "ReportedHashrate": 173398230
        },
        "pcsl2": {
          "ReportedHashrate": 173483176
        }
      },
      "block_found": 0,
      "time_period": 2491336,
      "start_time": "2017-05-14T15:40:00.3545871-07:00"
    },
    "2491337": {
      "mined_share": 154,
      "valid_share": 154,
      "total_valid_difficulty": 674520000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1070559437,
      "effective_hashrate": 1124200000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169460054
        },
        "his7chan": {
          "ReportedHashrate": 189093578
        },
        "pcs1": {
          "ReportedHashrate": 163497878
        },
        "pcs7chan": {
          "ReportedHashrate": 201870089
        },
        "pcsl1": {
          "ReportedHashrate": 173268257
        },
        "pcsl2": {
          "ReportedHashrate": 173369581
        }
      },
      "block_found": 0,
      "time_period": 2491337,
      "start_time": "2017-05-14T15:50:00.1072925-07:00"
    },
    "2491338": {
      "mined_share": 147,
      "valid_share": 147,
      "total_valid_difficulty": 643860000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1070434605,
      "effective_hashrate": 1073100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169176320
        },
        "his7chan": {
          "ReportedHashrate": 189293626
        },
        "pcs1": {
          "ReportedHashrate": 163353005
        },
        "pcs7chan": {
          "ReportedHashrate": 201896411
        },
        "pcsl1": {
          "ReportedHashrate": 173314302
        },
        "pcsl2": {
          "ReportedHashrate": 173400941
        }
      },
      "block_found": 0,
      "time_period": 2491338,
      "start_time": "2017-05-14T16:00:01.5320182-07:00"
    },
    "2491339": {
      "mined_share": 140,
      "valid_share": 140,
      "total_valid_difficulty": 613200000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1070438172,
      "effective_hashrate": 1022000000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169472763
        },
        "his7chan": {
          "ReportedHashrate": 189137351
        },
        "pcs1": {
          "ReportedHashrate": 163235367
        },
        "pcs7chan": {
          "ReportedHashrate": 201887827
        },
        "pcsl1": {
          "ReportedHashrate": 173288877
        },
        "pcsl2": {
          "ReportedHashrate": 173415987
        }
      },
      "block_found": 0,
      "time_period": 2491339,
      "start_time": "2017-05-14T16:10:04.099582-07:00"
    },
    "2491340": {
      "mined_share": 149,
      "valid_share": 149,
      "total_valid_difficulty": 652620000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1070860489,
      "effective_hashrate": 1087700000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169450412
        },
        "his7chan": {
          "ReportedHashrate": 189305993
        },
        "pcs1": {
          "ReportedHashrate": 163514305
        },
        "pcs7chan": {
          "ReportedHashrate": 202005891
        },
        "pcsl1": {
          "ReportedHashrate": 173350891
        },
        "pcsl2": {
          "ReportedHashrate": 173232997
        }
      },
      "block_found": 0,
      "time_period": 2491340,
      "start_time": "2017-05-14T16:20:01.2717378-07:00"
    },
    "2491341": {
      "mined_share": 148,
      "valid_share": 148,
      "total_valid_difficulty": 648240000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1070786985,
      "effective_hashrate": 1080400000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 169261395
        },
        "his7chan": {
          "ReportedHashrate": 189428566
        },
        "pcs1": {
          "ReportedHashrate": 163477047
        },
        "pcs7chan": {
          "ReportedHashrate": 201854339
        },
        "pcsl1": {
          "ReportedHashrate": 173419518
        },
        "pcsl2": {
          "ReportedHashrate": 173346120
        }
      },
      "block_found": 0,
      "time_period": 2491341,
      "start_time": "2017-05-14T16:30:01.3016175-07:00"
    },
    "2491342": {
      "mined_share": 117,
      "valid_share": 117,
      "total_valid_difficulty": 512460000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1069560613,
      "effective_hashrate": 854100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 168841234
        },
        "his7chan": {
          "ReportedHashrate": 188823892
        },
        "pcs1": {
          "ReportedHashrate": 163177016
        },
        "pcs7chan": {
          "ReportedHashrate": 202044347
        },
        "pcsl1": {
          "ReportedHashrate": 173341414
        },
        "pcsl2": {
          "ReportedHashrate": 173332710
        }
      },
      "block_found": 0,
      "time_period": 2491342,
      "start_time": "2017-05-14T16:40:04.4747471-07:00"
    },
    "2491356": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 140592274,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 140592274
        }
      },
      "block_found": 0,
      "time_period": 2491356,
      "start_time": "2017-05-14T19:09:33.7915023-07:00"
    },
    "2491357": {
      "mined_share": 43,
      "valid_share": 43,
      "total_valid_difficulty": 188340000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 330649897,
      "effective_hashrate": 313900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141357278
        },
        "his7chan": {
          "ReportedHashrate": 189292619
        }
      },
      "block_found": 0,
      "time_period": 2491357,
      "start_time": "2017-05-14T19:10:04.7928546-07:00"
    },
    "2491358": {
      "mined_share": 83,
      "valid_share": 83,
      "total_valid_difficulty": 363540000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 839047764,
      "effective_hashrate": 605900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141246383
        },
        "his7chan": {
          "ReportedHashrate": 189138652
        },
        "pcs1": {
          "ReportedHashrate": 162156054
        },
        "pcsl1": {
          "ReportedHashrate": 173449706
        },
        "pcsl2": {
          "ReportedHashrate": 173056969
        }
      },
      "block_found": 0,
      "time_period": 2491358,
      "start_time": "2017-05-14T19:20:00.9228357-07:00"
    },
    "2491359": {
      "mined_share": 104,
      "valid_share": 104,
      "total_valid_difficulty": 455520000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 1,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1041389975,
      "effective_hashrate": 759200000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141017428
        },
        "his7chan": {
          "ReportedHashrate": 189191103
        },
        "pcs1": {
          "ReportedHashrate": 162553395
        },
        "pcs7chan": {
          "ReportedHashrate": 202000556
        },
        "pcsl1": {
          "ReportedHashrate": 173112109
        },
        "pcsl2": {
          "ReportedHashrate": 173515384
        }
      },
      "block_found": 0,
      "time_period": 2491359,
      "start_time": "2017-05-14T19:30:02.0435114-07:00"
    },
    "2491360": {
      "mined_share": 97,
      "valid_share": 97,
      "total_valid_difficulty": 424860000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 1,
      "rejected_claim": 0,
      "reported_hashrate": 1038120465,
      "effective_hashrate": 708100000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 140482343
        },
        "his7chan": {
          "ReportedHashrate": 189054661
        },
        "pcs1": {
          "ReportedHashrate": 162225551
        },
        "pcs7chan": {
          "ReportedHashrate": 201027042
        },
        "pcsl1": {
          "ReportedHashrate": 172090286
        },
        "pcsl2": {
          "ReportedHashrate": 173240582
        }
      },
      "block_found": 0,
      "time_period": 2491360,
      "start_time": "2017-05-14T19:40:03.1135644-07:00"
    },
    "2491361": {
      "mined_share": 142,
      "valid_share": 142,
      "total_valid_difficulty": 621960000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1037928622,
      "effective_hashrate": 1036600000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141368603
        },
        "his7chan": {
          "ReportedHashrate": 189093784
        },
        "pcs1": {
          "ReportedHashrate": 162136308
        },
        "pcs7chan": {
          "ReportedHashrate": 198980659
        },
        "pcsl1": {
          "ReportedHashrate": 172964603
        },
        "pcsl2": {
          "ReportedHashrate": 173384665
        }
      },
      "block_found": 0,
      "time_period": 2491361,
      "start_time": "2017-05-14T19:50:00.4236614-07:00"
    },
    "2491362": {
      "mined_share": 133,
      "valid_share": 133,
      "total_valid_difficulty": 582540000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1038722554,
      "effective_hashrate": 970900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141105460
        },
        "his7chan": {
          "ReportedHashrate": 188654081
        },
        "pcs1": {
          "ReportedHashrate": 162345849
        },
        "pcs7chan": {
          "ReportedHashrate": 200246459
        },
        "pcsl1": {
          "ReportedHashrate": 173146859
        },
        "pcsl2": {
          "ReportedHashrate": 173223846
        }
      },
      "block_found": 0,
      "time_period": 2491362,
      "start_time": "2017-05-14T20:00:02.4032761-07:00"
    },
    "2491363": {
      "mined_share": 133,
      "valid_share": 133,
      "total_valid_difficulty": 582540000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1039517417,
      "effective_hashrate": 970900000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141079207
        },
        "his7chan": {
          "ReportedHashrate": 188651642
        },
        "pcs1": {
          "ReportedHashrate": 162378948
        },
        "pcs7chan": {
          "ReportedHashrate": 201400460
        },
        "pcsl1": {
          "ReportedHashrate": 172631233
        },
        "pcsl2": {
          "ReportedHashrate": 173375927
        }
      },
      "block_found": 0,
      "time_period": 2491363,
      "start_time": "2017-05-14T20:10:00.0239316-07:00"
    },
    "2491364": {
      "mined_share": 129,
      "valid_share": 129,
      "total_valid_difficulty": 565020000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1037397884,
      "effective_hashrate": 941700000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 140718994
        },
        "his7chan": {
          "ReportedHashrate": 189282909
        },
        "pcs1": {
          "ReportedHashrate": 161901463
        },
        "pcs7chan": {
          "ReportedHashrate": 199918344
        },
        "pcsl1": {
          "ReportedHashrate": 172322629
        },
        "pcsl2": {
          "ReportedHashrate": 173253545
        }
      },
      "block_found": 0,
      "time_period": 2491364,
      "start_time": "2017-05-14T20:20:00.1128665-07:00"
    },
    "2491365": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1007897962,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 139784853
        },
        "his480": {
          "ReportedHashrate": 170000956
        },
        "his7chan": {
          "ReportedHashrate": 189423661
        },
        "pcs1": {
          "ReportedHashrate": 161391200
        },
        "pcsl1": {
          "ReportedHashrate": 173559427
        },
        "pcsl2": {
          "ReportedHashrate": 173737865
        }
      },
      "block_found": 0,
      "time_period": 2491365,
      "start_time": "2017-05-14T20:30:48.0660399-07:00"
    },
    "2491366": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1196498204,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141487133
        },
        "his480": {
          "ReportedHashrate": 155723332
        },
        "his7chan": {
          "ReportedHashrate": 188895217
        },
        "pcs1": {
          "ReportedHashrate": 161924120
        },
        "pcs7chan": {
          "ReportedHashrate": 201347906
        },
        "pcsl1": {
          "ReportedHashrate": 173599793
        },
        "pcsl2": {
          "ReportedHashrate": 173520703
        }
      },
      "block_found": 0,
      "time_period": 2491366,
      "start_time": "2017-05-14T20:40:02.4407874-07:00"
    },
    "2491367": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1210669875,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141459837
        },
        "his480": {
          "ReportedHashrate": 170038156
        },
        "his7chan": {
          "ReportedHashrate": 188818567
        },
        "pcs1": {
          "ReportedHashrate": 161462702
        },
        "pcs7chan": {
          "ReportedHashrate": 202077036
        },
        "pcsl1": {
          "ReportedHashrate": 173226737
        },
        "pcsl2": {
          "ReportedHashrate": 173586840
        }
      },
      "block_found": 0,
      "time_period": 2491367,
      "start_time": "2017-05-14T20:50:05.8959447-07:00"
    },
    "2491368": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1210296237,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141484678
        },
        "his480": {
          "ReportedHashrate": 169247095
        },
        "his7chan": {
          "ReportedHashrate": 189618110
        },
        "pcs1": {
          "ReportedHashrate": 161800463
        },
        "pcs7chan": {
          "ReportedHashrate": 202293991
        },
        "pcsl1": {
          "ReportedHashrate": 173383729
        },
        "pcsl2": {
          "ReportedHashrate": 172468171
        }
      },
      "block_found": 0,
      "time_period": 2491368,
      "start_time": "2017-05-14T21:00:02.9808782-07:00"
    },
    "2491369": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1211026760,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141474094
        },
        "his480": {
          "ReportedHashrate": 169877709
        },
        "his7chan": {
          "ReportedHashrate": 189515029
        },
        "pcs1": {
          "ReportedHashrate": 160742078
        },
        "pcs7chan": {
          "ReportedHashrate": 202082190
        },
        "pcsl1": {
          "ReportedHashrate": 173633737
        },
        "pcsl2": {
          "ReportedHashrate": 173701923
        }
      },
      "block_found": 0,
      "time_period": 2491369,
      "start_time": "2017-05-14T21:10:15.1458449-07:00"
    },
    "2491370": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1209385032,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 139635579
        },
        "his480": {
          "ReportedHashrate": 170243648
        },
        "his7chan": {
          "ReportedHashrate": 189366365
        },
        "pcs1": {
          "ReportedHashrate": 161653744
        },
        "pcs7chan": {
          "ReportedHashrate": 201206632
        },
        "pcsl1": {
          "ReportedHashrate": 173555974
        },
        "pcsl2": {
          "ReportedHashrate": 173723090
        }
      },
      "block_found": 0,
      "time_period": 2491370,
      "start_time": "2017-05-14T21:20:07.8688942-07:00"
    },
    "2491371": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1195549709,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141465371
        },
        "his480": {
          "ReportedHashrate": 169892408
        },
        "his7chan": {
          "ReportedHashrate": 189478752
        },
        "pcs1": {
          "ReportedHashrate": 161781084
        },
        "pcs7chan": {
          "ReportedHashrate": 202095733
        },
        "pcsl1": {
          "ReportedHashrate": 157094594
        },
        "pcsl2": {
          "ReportedHashrate": 173741767
        }
      },
      "block_found": 0,
      "time_period": 2491371,
      "start_time": "2017-05-14T21:30:02.1835953-07:00"
    },
    "2491372": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1036602889,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141500165
        },
        "his480": {
          "ReportedHashrate": 169764166
        },
        "his7chan": {
          "ReportedHashrate": 189407533
        },
        "pcs1": {
          "ReportedHashrate": 161272679
        },
        "pcs7chan": {
          "ReportedHashrate": 200932693
        },
        "pcsl2": {
          "ReportedHashrate": 173725653
        }
      },
      "block_found": 0,
      "time_period": 2491372,
      "start_time": "2017-05-14T21:40:36.972272-07:00"
    },
    "2491373": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1036633673,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 140931087
        },
        "his480": {
          "ReportedHashrate": 169908402
        },
        "his7chan": {
          "ReportedHashrate": 188808572
        },
        "pcs1": {
          "ReportedHashrate": 161526926
        },
        "pcs7chan": {
          "ReportedHashrate": 202290289
        },
        "pcsl2": {
          "ReportedHashrate": 173168397
        }
      },
      "block_found": 0,
      "time_period": 2491373,
      "start_time": "2017-05-14T21:50:03.0191337-07:00"
    },
    "2491374": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1037584292,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141492215
        },
        "his480": {
          "ReportedHashrate": 169922422
        },
        "his7chan": {
          "ReportedHashrate": 189426928
        },
        "pcs1": {
          "ReportedHashrate": 161632834
        },
        "pcs7chan": {
          "ReportedHashrate": 201776694
        },
        "pcsl2": {
          "ReportedHashrate": 173333199
        }
      },
      "block_found": 0,
      "time_period": 2491374,
      "start_time": "2017-05-14T22:00:11.6590823-07:00"
    },
    "2491375": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1025042106,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 148023873
        },
        "his480": {
          "ReportedHashrate": 169024733
        },
        "his7chan": {
          "ReportedHashrate": 171169001
        },
        "pcs1": {
          "ReportedHashrate": 161084909
        },
        "pcs7chan": {
          "ReportedHashrate": 202249311
        },
        "pcsl2": {
          "ReportedHashrate": 173490279
        }
      },
      "block_found": 0,
      "time_period": 2491375,
      "start_time": "2017-05-14T22:10:18.254572-07:00"
    },
    "2491376": {
      "mined_share": 201,
      "valid_share": 201,
      "total_valid_difficulty": 880380000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897052453,
      "effective_hashrate": 1467300000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 0
        },
        "his480": {
          "ReportedHashrate": 170224558
        },
        "his7chan": {
          "ReportedHashrate": 188954772
        },
        "pcs1": {
          "ReportedHashrate": 162138459
        },
        "pcs7chan": {
          "ReportedHashrate": 202240531
        },
        "pcsl1": {
          "ReportedHashrate": 173494133
        },
        "pcsl2": {
          "ReportedHashrate": 0
        }
      },
      "block_found": 0,
      "time_period": 2491376,
      "start_time": "2017-05-14T22:20:08.4542167-07:00"
    },
    "2491377": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 161993290,
      "effective_hashrate": 0,
      "rigs": {
        "pcs1": {
          "ReportedHashrate": 161993290
        }
      },
      "block_found": 0,
      "time_period": 2491377,
      "start_time": "2017-05-14T22:30:01.0512512-07:00"
    }
  },
  "overall": {
    "last_mined_share": "2017-05-14T22:29:52.6250583-07:00",
    "last_valid_share": "2017-05-14T22:29:52.6250583-07:00",
    "last_rejected_share": "2017-05-14T10:15:16.0643921-07:00",
    "last_block": "2017-05-14T15:12:10.8435536-07:00",
    "mined_share": 9160,
    "valid_share": 9156,
    "total_valid_difficulty": 40103280000000,
    "average_share_difficulty": 4380000000,
    "rejected_share": 4,
    "last_submitted_claim": "2017-05-14T19:36:44.2517393-07:00",
    "last_accepted_claim": "2017-05-14T19:41:38.3658733-07:00",
    "last_rejected_claim": "0001-01-01T00:00:00Z",
    "total_submitted_claim": 3,
    "total_accepted_claim": 2,
    "total_rejected_claim": 1,
    "reported_hashrate": 1234767005,
    "effective_hashrate": 66838800000,
    "rigs": {
      "his470ss": {
        "ReportedHashrate": 165827772
      },
      "his480": {
        "ReportedHashrate": 168515512
      },
      "his7chan": {
        "ReportedHashrate": 189163484
      },
      "pcs1": {
        "ReportedHashrate": 163289200
      },
      "pcs7chan": {
        "ReportedHashrate": 201631073
      },
      "pcsl1": {
        "ReportedHashrate": 173203616
      },
      "pcsl2": {
        "ReportedHashrate": 173136348
      }
    },
    "total_block_found": 1,
    "pending_share": 881,
    "abandoned_share": 2928,
    "being_validated_share": 0,
    "verified_share": 5347,
    "bad_share": 0,
    "start_time": "2017-05-14T05:34:36.6338567-07:00"
  },
  "period_duration": 600,
  "short_window_duration": 7200,
  "short_window_sample": {
    "2491365": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1007897962,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 139784853
        },
        "his480": {
          "ReportedHashrate": 170000956
        },
        "his7chan": {
          "ReportedHashrate": 189423661
        },
        "pcs1": {
          "ReportedHashrate": 161391200
        },
        "pcsl1": {
          "ReportedHashrate": 173559427
        },
        "pcsl2": {
          "ReportedHashrate": 173737865
        }
      },
      "block_found": 0,
      "time_period": 2491365,
      "start_time": "2017-05-14T20:30:48.0660399-07:00"
    },
    "2491366": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1196498204,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141487133
        },
        "his480": {
          "ReportedHashrate": 155723332
        },
        "his7chan": {
          "ReportedHashrate": 188895217
        },
        "pcs1": {
          "ReportedHashrate": 161924120
        },
        "pcs7chan": {
          "ReportedHashrate": 201347906
        },
        "pcsl1": {
          "ReportedHashrate": 173599793
        },
        "pcsl2": {
          "ReportedHashrate": 173520703
        }
      },
      "block_found": 0,
      "time_period": 2491366,
      "start_time": "2017-05-14T20:40:02.4407874-07:00"
    },
    "2491367": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1210669875,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141459837
        },
        "his480": {
          "ReportedHashrate": 170038156
        },
        "his7chan": {
          "ReportedHashrate": 188818567
        },
        "pcs1": {
          "ReportedHashrate": 161462702
        },
        "pcs7chan": {
          "ReportedHashrate": 202077036
        },
        "pcsl1": {
          "ReportedHashrate": 173226737
        },
        "pcsl2": {
          "ReportedHashrate": 173586840
        }
      },
      "block_found": 0,
      "time_period": 2491367,
      "start_time": "2017-05-14T20:50:05.8959447-07:00"
    },
    "2491368": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1210296237,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141484678
        },
        "his480": {
          "ReportedHashrate": 169247095
        },
        "his7chan": {
          "ReportedHashrate": 189618110
        },
        "pcs1": {
          "ReportedHashrate": 161800463
        },
        "pcs7chan": {
          "ReportedHashrate": 202293991
        },
        "pcsl1": {
          "ReportedHashrate": 173383729
        },
        "pcsl2": {
          "ReportedHashrate": 172468171
        }
      },
      "block_found": 0,
      "time_period": 2491368,
      "start_time": "2017-05-14T21:00:02.9808782-07:00"
    },
    "2491369": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1211026760,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141474094
        },
        "his480": {
          "ReportedHashrate": 169877709
        },
        "his7chan": {
          "ReportedHashrate": 189515029
        },
        "pcs1": {
          "ReportedHashrate": 160742078
        },
        "pcs7chan": {
          "ReportedHashrate": 202082190
        },
        "pcsl1": {
          "ReportedHashrate": 173633737
        },
        "pcsl2": {
          "ReportedHashrate": 173701923
        }
      },
      "block_found": 0,
      "time_period": 2491369,
      "start_time": "2017-05-14T21:10:15.1458449-07:00"
    },
    "2491370": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1209385032,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 139635579
        },
        "his480": {
          "ReportedHashrate": 170243648
        },
        "his7chan": {
          "ReportedHashrate": 189366365
        },
        "pcs1": {
          "ReportedHashrate": 161653744
        },
        "pcs7chan": {
          "ReportedHashrate": 201206632
        },
        "pcsl1": {
          "ReportedHashrate": 173555974
        },
        "pcsl2": {
          "ReportedHashrate": 173723090
        }
      },
      "block_found": 0,
      "time_period": 2491370,
      "start_time": "2017-05-14T21:20:07.8688942-07:00"
    },
    "2491371": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1195549709,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141465371
        },
        "his480": {
          "ReportedHashrate": 169892408
        },
        "his7chan": {
          "ReportedHashrate": 189478752
        },
        "pcs1": {
          "ReportedHashrate": 161781084
        },
        "pcs7chan": {
          "ReportedHashrate": 202095733
        },
        "pcsl1": {
          "ReportedHashrate": 157094594
        },
        "pcsl2": {
          "ReportedHashrate": 173741767
        }
      },
      "block_found": 0,
      "time_period": 2491371,
      "start_time": "2017-05-14T21:30:02.1835953-07:00"
    },
    "2491372": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1036602889,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141500165
        },
        "his480": {
          "ReportedHashrate": 169764166
        },
        "his7chan": {
          "ReportedHashrate": 189407533
        },
        "pcs1": {
          "ReportedHashrate": 161272679
        },
        "pcs7chan": {
          "ReportedHashrate": 200932693
        },
        "pcsl2": {
          "ReportedHashrate": 173725653
        }
      },
      "block_found": 0,
      "time_period": 2491372,
      "start_time": "2017-05-14T21:40:36.972272-07:00"
    },
    "2491373": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1036633673,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 140931087
        },
        "his480": {
          "ReportedHashrate": 169908402
        },
        "his7chan": {
          "ReportedHashrate": 188808572
        },
        "pcs1": {
          "ReportedHashrate": 161526926
        },
        "pcs7chan": {
          "ReportedHashrate": 202290289
        },
        "pcsl2": {
          "ReportedHashrate": 173168397
        }
      },
      "block_found": 0,
      "time_period": 2491373,
      "start_time": "2017-05-14T21:50:03.0191337-07:00"
    },
    "2491374": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1037584292,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 141492215
        },
        "his480": {
          "ReportedHashrate": 169922422
        },
        "his7chan": {
          "ReportedHashrate": 189426928
        },
        "pcs1": {
          "ReportedHashrate": 161632834
        },
        "pcs7chan": {
          "ReportedHashrate": 201776694
        },
        "pcsl2": {
          "ReportedHashrate": 173333199
        }
      },
      "block_found": 0,
      "time_period": 2491374,
      "start_time": "2017-05-14T22:00:11.6590823-07:00"
    },
    "2491375": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 1025042106,
      "effective_hashrate": 0,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 148023873
        },
        "his480": {
          "ReportedHashrate": 169024733
        },
        "his7chan": {
          "ReportedHashrate": 171169001
        },
        "pcs1": {
          "ReportedHashrate": 161084909
        },
        "pcs7chan": {
          "ReportedHashrate": 202249311
        },
        "pcsl2": {
          "ReportedHashrate": 173490279
        }
      },
      "block_found": 0,
      "time_period": 2491375,
      "start_time": "2017-05-14T22:10:18.254572-07:00"
    },
    "2491376": {
      "mined_share": 201,
      "valid_share": 201,
      "total_valid_difficulty": 880380000000,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 897052453,
      "effective_hashrate": 1467300000,
      "rigs": {
        "his470ss": {
          "ReportedHashrate": 0
        },
        "his480": {
          "ReportedHashrate": 170224558
        },
        "his7chan": {
          "ReportedHashrate": 188954772
        },
        "pcs1": {
          "ReportedHashrate": 162138459
        },
        "pcs7chan": {
          "ReportedHashrate": 202240531
        },
        "pcsl1": {
          "ReportedHashrate": 173494133
        },
        "pcsl2": {
          "ReportedHashrate": 0
        }
      },
      "block_found": 0,
      "time_period": 2491376,
      "start_time": "2017-05-14T22:20:08.4542167-07:00"
    },
    "2491377": {
      "mined_share": 0,
      "valid_share": 0,
      "total_valid_difficulty": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "submitted_claim": 0,
      "accepted_claim": 0,
      "rejected_claim": 0,
      "reported_hashrate": 161993290,
      "effective_hashrate": 0,
      "rigs": {
        "pcs1": {
          "ReportedHashrate": 161993290
        }
      },
      "block_found": 0,
      "time_period": 2491377,
      "start_time": "2017-05-14T22:30:01.0512512-07:00"
    }
  }
}
	`
}
func getDataRig() string {
	return `
	{
  "last_1_hour": {
    "24905784": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149414090,
      "effective_hashrate": 243333333,
      "block_found": 0,
      "time_period": 24905784,
      "start_time": "2017-05-09T16:24:01.015801947Z"
    },
    "24905785": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150179006,
      "effective_hashrate": 82641509,
      "block_found": 0,
      "time_period": 24905785,
      "start_time": "2017-05-09T16:25:02.868159311Z"
    },
    "24905786": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149840523,
      "effective_hashrate": 389333333,
      "block_found": 0,
      "time_period": 24905786,
      "start_time": "2017-05-09T16:26:04.354419028Z"
    },
    "24905787": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149949542,
      "effective_hashrate": 156428571,
      "block_found": 0,
      "time_period": 24905787,
      "start_time": "2017-05-09T16:27:00.682953473Z"
    },
    "24905788": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150058186,
      "effective_hashrate": 171764705,
      "block_found": 0,
      "time_period": 24905788,
      "start_time": "2017-05-09T16:28:02.804729477Z"
    },
    "24905789": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150034848,
      "effective_hashrate": 398181818,
      "block_found": 0,
      "time_period": 24905789,
      "start_time": "2017-05-09T16:29:04.244924543Z"
    },
    "24905790": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150169720,
      "effective_hashrate": 320487804,
      "block_found": 0,
      "time_period": 24905790,
      "start_time": "2017-05-09T16:30:00.561643408Z"
    },
    "24905791": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149190733,
      "effective_hashrate": 292000000,
      "block_found": 0,
      "time_period": 24905791,
      "start_time": "2017-05-09T16:31:02.136949554Z"
    },
    "24905792": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149926527,
      "effective_hashrate": 125142857,
      "block_found": 0,
      "time_period": 24905792,
      "start_time": "2017-05-09T16:32:03.302626137Z"
    },
    "24905793": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149801308,
      "effective_hashrate": 178775510,
      "block_found": 0,
      "time_period": 24905793,
      "start_time": "2017-05-09T16:33:05.077724168Z"
    },
    "24905794": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149862821,
      "effective_hashrate": 171764705,
      "block_found": 0,
      "time_period": 24905794,
      "start_time": "2017-05-09T16:34:01.403604951Z"
    },
    "24905795": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 150027169,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905795,
      "start_time": "2017-05-09T16:35:02.672223793Z"
    },
    "24905796": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 148838767,
      "effective_hashrate": 226551724,
      "block_found": 0,
      "time_period": 24905796,
      "start_time": "2017-05-09T16:36:01.328651401Z"
    },
    "24905797": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149879547,
      "effective_hashrate": 75517241,
      "block_found": 0,
      "time_period": 24905797,
      "start_time": "2017-05-09T16:37:00.403607398Z"
    },
    "24905798": {
      "mined_share": 5,
      "valid_share": 5,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149487073,
      "effective_hashrate": 521428571,
      "block_found": 0,
      "time_period": 24905798,
      "start_time": "2017-05-09T16:38:01.496041334Z"
    },
    "24905799": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149832267,
      "effective_hashrate": 156428571,
      "block_found": 0,
      "time_period": 24905799,
      "start_time": "2017-05-09T16:39:03.13336059Z"
    },
    "24905800": {
      "mined_share": 6,
      "valid_share": 6,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149652402,
      "effective_hashrate": 486666666,
      "block_found": 0,
      "time_period": 24905800,
      "start_time": "2017-05-09T16:40:03.15538684Z"
    },
    "24905801": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149681799,
      "effective_hashrate": 257647058,
      "block_found": 0,
      "time_period": 24905801,
      "start_time": "2017-05-09T16:41:01.749104048Z"
    },
    "24905802": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149815321,
      "effective_hashrate": 625714285,
      "block_found": 0,
      "time_period": 24905802,
      "start_time": "2017-05-09T16:42:03.012391112Z"
    },
    "24905803": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149778509,
      "effective_hashrate": 292000000,
      "block_found": 0,
      "time_period": 24905803,
      "start_time": "2017-05-09T16:43:00.078834955Z"
    },
    "24905804": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149633727,
      "effective_hashrate": 238909090,
      "block_found": 0,
      "time_period": 24905804,
      "start_time": "2017-05-09T16:44:01.874319927Z"
    },
    "24905805": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149555651,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905805,
      "start_time": "2017-05-09T16:45:03.936010301Z"
    },
    "24905806": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149772875,
      "effective_hashrate": 486666666,
      "block_found": 0,
      "time_period": 24905806,
      "start_time": "2017-05-09T16:46:00.272641049Z"
    },
    "24905807": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149234376,
      "effective_hashrate": 153684210,
      "block_found": 0,
      "time_period": 24905807,
      "start_time": "2017-05-09T16:47:01.888109598Z"
    },
    "24905808": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149643713,
      "effective_hashrate": 312857142,
      "block_found": 0,
      "time_period": 24905808,
      "start_time": "2017-05-09T16:48:03.545967202Z"
    },
    "24905809": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149467453,
      "effective_hashrate": 547500000,
      "block_found": 0,
      "time_period": 24905809,
      "start_time": "2017-05-09T16:49:00.000557063Z"
    },
    "24905810": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149203244,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905810,
      "start_time": "2017-05-09T16:50:01.687095376Z"
    },
    "24905811": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149852085,
      "effective_hashrate": 84230769,
      "block_found": 0,
      "time_period": 24905811,
      "start_time": "2017-05-09T16:51:03.848908866Z"
    },
    "24905812": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149743399,
      "effective_hashrate": 112307692,
      "block_found": 0,
      "time_period": 24905812,
      "start_time": "2017-05-09T16:52:01.504950325Z"
    },
    "24905813": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149592633,
      "effective_hashrate": 298636363,
      "block_found": 0,
      "time_period": 24905813,
      "start_time": "2017-05-09T16:53:04.365886551Z"
    },
    "24905814": {
      "mined_share": 7,
      "valid_share": 7,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149679193,
      "effective_hashrate": 547500000,
      "block_found": 1,
      "time_period": 24905814,
      "start_time": "2017-05-09T16:54:00.98311751Z"
    },
    "24905815": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149805715,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905815,
      "start_time": "2017-05-09T16:55:04.36661599Z"
    },
    "24905816": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149674532,
      "effective_hashrate": 250285714,
      "block_found": 0,
      "time_period": 24905816,
      "start_time": "2017-05-09T16:56:01.739933648Z"
    },
    "24905817": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149768695,
      "effective_hashrate": 584000000,
      "block_found": 0,
      "time_period": 24905817,
      "start_time": "2017-05-09T16:57:03.896193415Z"
    },
    "24905818": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149805592,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905818,
      "start_time": "2017-05-09T16:58:01.541064989Z"
    },
    "24905819": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149899120,
      "effective_hashrate": 262800000,
      "block_found": 0,
      "time_period": 24905819,
      "start_time": "2017-05-09T16:59:05.039207578Z"
    },
    "24905820": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149927160,
      "effective_hashrate": 171764705,
      "block_found": 0,
      "time_period": 24905820,
      "start_time": "2017-05-09T17:00:02.864919173Z"
    },
    "24905821": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149787613,
      "effective_hashrate": 257647058,
      "block_found": 0,
      "time_period": 24905821,
      "start_time": "2017-05-09T17:01:00.587387399Z"
    },
    "24905822": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149792973,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905822,
      "start_time": "2017-05-09T17:02:03.573013126Z"
    },
    "24905823": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149914985,
      "effective_hashrate": 336923076,
      "block_found": 0,
      "time_period": 24905823,
      "start_time": "2017-05-09T17:03:01.258804189Z"
    },
    "24905824": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149820506,
      "effective_hashrate": 101860465,
      "block_found": 0,
      "time_period": 24905824,
      "start_time": "2017-05-09T17:04:04.321582459Z"
    },
    "24905825": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149795514,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905825,
      "start_time": "2017-05-09T17:05:02.087228015Z"
    },
    "24905826": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149866246,
      "effective_hashrate": 199090909,
      "block_found": 0,
      "time_period": 24905826,
      "start_time": "2017-05-09T17:06:04.102838333Z"
    },
    "24905827": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149942432,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905827,
      "start_time": "2017-05-09T17:07:00.868891928Z"
    },
    "24905828": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149788685,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905828,
      "start_time": "2017-05-09T17:08:02.237683864Z"
    },
    "24905829": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150071614,
      "effective_hashrate": 151034482,
      "block_found": 0,
      "time_period": 24905829,
      "start_time": "2017-05-09T17:09:03.864418531Z"
    },
    "24905830": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149381546,
      "effective_hashrate": 365000000,
      "block_found": 0,
      "time_period": 24905830,
      "start_time": "2017-05-09T17:10:00.441615499Z"
    },
    "24905831": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149341908,
      "effective_hashrate": 146000000,
      "block_found": 0,
      "time_period": 24905831,
      "start_time": "2017-05-09T17:11:02.363053509Z"
    },
    "24905832": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149527504,
      "effective_hashrate": 162222222,
      "block_found": 0,
      "time_period": 24905832,
      "start_time": "2017-05-09T17:12:03.835007241Z"
    },
    "24905833": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149410231,
      "effective_hashrate": 89387755,
      "block_found": 0,
      "time_period": 24905833,
      "start_time": "2017-05-09T17:13:00.232218978Z"
    },
    "24905834": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149981254,
      "effective_hashrate": 165283018,
      "block_found": 0,
      "time_period": 24905834,
      "start_time": "2017-05-09T17:14:02.448299519Z"
    },
    "24905835": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149836303,
      "effective_hashrate": 324444444,
      "block_found": 0,
      "time_period": 24905835,
      "start_time": "2017-05-09T17:15:03.929021946Z"
    },
    "24905836": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149368046,
      "effective_hashrate": 156428571,
      "block_found": 0,
      "time_period": 24905836,
      "start_time": "2017-05-09T17:16:04.962565117Z"
    },
    "24905837": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149900194,
      "effective_hashrate": 171764705,
      "block_found": 0,
      "time_period": 24905837,
      "start_time": "2017-05-09T17:17:01.354205011Z"
    },
    "24905838": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149857187,
      "effective_hashrate": 938571428,
      "block_found": 0,
      "time_period": 24905838,
      "start_time": "2017-05-09T17:18:00.992187217Z"
    },
    "24905839": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149699991,
      "effective_hashrate": 104285714,
      "block_found": 0,
      "time_period": 24905839,
      "start_time": "2017-05-09T17:19:00.230552106Z"
    },
    "24905840": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149708763,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905840,
      "start_time": "2017-05-09T17:20:02.057575079Z"
    },
    "24905841": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149610996,
      "effective_hashrate": 438000000,
      "block_found": 0,
      "time_period": 24905841,
      "start_time": "2017-05-09T17:21:03.526802898Z"
    },
    "24905842": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149336310,
      "effective_hashrate": 153684210,
      "block_found": 0,
      "time_period": 24905842,
      "start_time": "2017-05-09T17:22:00.561862253Z"
    },
    "24905843": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 148720380,
      "effective_hashrate": 547500000,
      "block_found": 0,
      "time_period": 24905843,
      "start_time": "2017-05-09T17:23:02.214242705Z"
    },
    "24905844": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149684285,
      "effective_hashrate": 772941176,
      "block_found": 0,
      "time_period": 24905844,
      "start_time": "2017-05-09T17:24:04.058779359Z"
    }
  },
  "last_3_hours": {
    "24905664": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150043538,
      "effective_hashrate": 438000000,
      "block_found": 0,
      "time_period": 24905664,
      "start_time": "2017-05-09T14:24:05.008004903Z"
    },
    "24905665": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149988706,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905665,
      "start_time": "2017-05-09T14:25:02.534479185Z"
    },
    "24905666": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150118045,
      "effective_hashrate": 175200000,
      "block_found": 0,
      "time_period": 24905666,
      "start_time": "2017-05-09T14:26:00.280513184Z"
    },
    "24905667": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149506282,
      "effective_hashrate": 224615384,
      "block_found": 0,
      "time_period": 24905667,
      "start_time": "2017-05-09T14:27:03.496329249Z"
    },
    "24905668": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149999805,
      "effective_hashrate": 324444444,
      "block_found": 0,
      "time_period": 24905668,
      "start_time": "2017-05-09T14:28:01.331944301Z"
    },
    "24905669": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149958026,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905669,
      "start_time": "2017-05-09T14:29:04.278656314Z"
    },
    "24905670": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150098310,
      "effective_hashrate": 162222222,
      "block_found": 0,
      "time_period": 24905670,
      "start_time": "2017-05-09T14:30:02.311768391Z"
    },
    "24905671": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150108864,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905671,
      "start_time": "2017-05-09T14:31:00.645510009Z"
    },
    "24905672": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149947283,
      "effective_hashrate": 350400000,
      "block_found": 0,
      "time_period": 24905672,
      "start_time": "2017-05-09T14:32:01.913050374Z"
    },
    "24905673": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149528777,
      "effective_hashrate": 410625000,
      "block_found": 0,
      "time_period": 24905673,
      "start_time": "2017-05-09T14:33:04.582453421Z"
    },
    "24905674": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149947542,
      "effective_hashrate": 298636363,
      "block_found": 0,
      "time_period": 24905674,
      "start_time": "2017-05-09T14:34:02.535338077Z"
    },
    "24905675": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150037364,
      "effective_hashrate": 165283018,
      "block_found": 0,
      "time_period": 24905675,
      "start_time": "2017-05-09T14:35:00.256214342Z"
    },
    "24905676": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150056287,
      "effective_hashrate": 257647058,
      "block_found": 0,
      "time_period": 24905676,
      "start_time": "2017-05-09T14:36:03.060544627Z"
    },
    "24905677": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150056800,
      "effective_hashrate": 380869565,
      "block_found": 0,
      "time_period": 24905677,
      "start_time": "2017-05-09T14:37:00.500564454Z"
    },
    "24905678": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149961244,
      "effective_hashrate": 365000000,
      "block_found": 0,
      "time_period": 24905678,
      "start_time": "2017-05-09T14:38:03.379427309Z"
    },
    "24905679": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150068020,
      "effective_hashrate": 162222222,
      "block_found": 0,
      "time_period": 24905679,
      "start_time": "2017-05-09T14:39:00.539155424Z"
    },
    "24905680": {
      "mined_share": 6,
      "valid_share": 6,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150004557,
      "effective_hashrate": 469285714,
      "block_found": 0,
      "time_period": 24905680,
      "start_time": "2017-05-09T14:40:03.352178663Z"
    },
    "24905681": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149849984,
      "effective_hashrate": 4380000000,
      "block_found": 0,
      "time_period": 24905681,
      "start_time": "2017-05-09T14:41:01.121435343Z"
    },
    "24905682": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149970319,
      "effective_hashrate": 438000000,
      "block_found": 0,
      "time_period": 24905682,
      "start_time": "2017-05-09T14:42:03.82808267Z"
    },
    "24905683": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150058789,
      "effective_hashrate": 115263157,
      "block_found": 0,
      "time_period": 24905683,
      "start_time": "2017-05-09T14:43:00.481313537Z"
    },
    "24905684": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 145720893,
      "effective_hashrate": 190434782,
      "block_found": 0,
      "time_period": 24905684,
      "start_time": "2017-05-09T14:44:03.820911635Z"
    },
    "24905685": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150064539,
      "effective_hashrate": 330566037,
      "block_found": 0,
      "time_period": 24905685,
      "start_time": "2017-05-09T14:45:01.816570499Z"
    },
    "24905686": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149822180,
      "effective_hashrate": 230526315,
      "block_found": 0,
      "time_period": 24905686,
      "start_time": "2017-05-09T14:46:04.81824174Z"
    },
    "24905687": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149284374,
      "effective_hashrate": 78214285,
      "block_found": 0,
      "time_period": 24905687,
      "start_time": "2017-05-09T14:47:02.553549072Z"
    },
    "24905688": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149828855,
      "effective_hashrate": 1095000000,
      "block_found": 0,
      "time_period": 24905688,
      "start_time": "2017-05-09T14:48:00.440528194Z"
    },
    "24905689": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149436666,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905689,
      "start_time": "2017-05-09T14:49:03.110238943Z"
    },
    "24905690": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149568631,
      "effective_hashrate": 95217391,
      "block_found": 0,
      "time_period": 24905690,
      "start_time": "2017-05-09T14:50:00.673505618Z"
    },
    "24905691": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150080840,
      "effective_hashrate": 252692307,
      "block_found": 0,
      "time_period": 24905691,
      "start_time": "2017-05-09T14:51:03.185912553Z"
    },
    "24905692": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149962289,
      "effective_hashrate": 168461538,
      "block_found": 0,
      "time_period": 24905692,
      "start_time": "2017-05-09T14:52:00.763404834Z"
    },
    "24905693": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150053066,
      "effective_hashrate": 273750000,
      "block_found": 0,
      "time_period": 24905693,
      "start_time": "2017-05-09T14:53:03.447733923Z"
    },
    "24905694": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149735422,
      "effective_hashrate": 226551724,
      "block_found": 0,
      "time_period": 24905694,
      "start_time": "2017-05-09T14:54:01.21709389Z"
    },
    "24905695": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149428078,
      "effective_hashrate": 87600000,
      "block_found": 0,
      "time_period": 24905695,
      "start_time": "2017-05-09T14:55:03.720390391Z"
    },
    "24905696": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149468556,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905696,
      "start_time": "2017-05-09T14:56:01.215379519Z"
    },
    "24905697": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149866942,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905697,
      "start_time": "2017-05-09T14:57:03.990730047Z"
    },
    "24905698": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149971208,
      "effective_hashrate": 730000000,
      "block_found": 0,
      "time_period": 24905698,
      "start_time": "2017-05-09T14:58:01.772113186Z"
    },
    "24905699": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149964053,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905699,
      "start_time": "2017-05-09T14:59:04.381534977Z"
    },
    "24905700": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149856767,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905700,
      "start_time": "2017-05-09T15:00:01.987497697Z"
    },
    "24905701": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149897061,
      "effective_hashrate": 417142857,
      "block_found": 0,
      "time_period": 24905701,
      "start_time": "2017-05-09T15:01:04.477080376Z"
    },
    "24905702": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149854635,
      "effective_hashrate": 262800000,
      "block_found": 0,
      "time_period": 24905702,
      "start_time": "2017-05-09T15:02:01.969074527Z"
    },
    "24905703": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 148780974,
      "effective_hashrate": 350400000,
      "block_found": 0,
      "time_period": 24905703,
      "start_time": "2017-05-09T15:03:00.666782263Z"
    },
    "24905704": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149876530,
      "effective_hashrate": 178775510,
      "block_found": 0,
      "time_period": 24905704,
      "start_time": "2017-05-09T15:04:02.654763554Z"
    },
    "24905705": {
      "mined_share": 7,
      "valid_share": 7,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149315541,
      "effective_hashrate": 537894736,
      "block_found": 0,
      "time_period": 24905705,
      "start_time": "2017-05-09T15:05:00.278031666Z"
    },
    "24905706": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149814974,
      "effective_hashrate": 156428571,
      "block_found": 0,
      "time_period": 24905706,
      "start_time": "2017-05-09T15:06:02.859149987Z"
    },
    "24905707": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149523817,
      "effective_hashrate": 312857142,
      "block_found": 0,
      "time_period": 24905707,
      "start_time": "2017-05-09T15:07:00.135939304Z"
    },
    "24905708": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149413003,
      "effective_hashrate": 175200000,
      "block_found": 0,
      "time_period": 24905708,
      "start_time": "2017-05-09T15:08:03.117801197Z"
    },
    "24905709": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149840343,
      "effective_hashrate": 312857142,
      "block_found": 0,
      "time_period": 24905709,
      "start_time": "2017-05-09T15:09:00.057311144Z"
    },
    "24905710": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149211708,
      "effective_hashrate": 243333333,
      "block_found": 0,
      "time_period": 24905710,
      "start_time": "2017-05-09T15:10:03.007474976Z"
    },
    "24905711": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149978682,
      "effective_hashrate": 365000000,
      "block_found": 0,
      "time_period": 24905711,
      "start_time": "2017-05-09T15:11:00.413424723Z"
    },
    "24905712": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149621224,
      "effective_hashrate": 292000000,
      "block_found": 0,
      "time_period": 24905712,
      "start_time": "2017-05-09T15:12:02.553958153Z"
    },
    "24905713": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149924835,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905713,
      "start_time": "2017-05-09T15:13:04.599323356Z"
    },
    "24905714": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149411194,
      "effective_hashrate": 257647058,
      "block_found": 0,
      "time_period": 24905714,
      "start_time": "2017-05-09T15:14:00.851660506Z"
    },
    "24905715": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149591760,
      "effective_hashrate": 262800000,
      "block_found": 0,
      "time_period": 24905715,
      "start_time": "2017-05-09T15:15:02.093241924Z"
    },
    "24905716": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149522553,
      "effective_hashrate": 318545454,
      "block_found": 0,
      "time_period": 24905716,
      "start_time": "2017-05-09T15:16:03.669255976Z"
    },
    "24905717": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150016849,
      "effective_hashrate": 178775510,
      "block_found": 0,
      "time_period": 24905717,
      "start_time": "2017-05-09T15:17:00.044904069Z"
    },
    "24905718": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150092591,
      "effective_hashrate": 312857142,
      "block_found": 0,
      "time_period": 24905718,
      "start_time": "2017-05-09T15:18:01.741524476Z"
    },
    "24905719": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149435859,
      "effective_hashrate": 386470588,
      "block_found": 0,
      "time_period": 24905719,
      "start_time": "2017-05-09T15:19:03.432673925Z"
    },
    "24905720": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149741671,
      "effective_hashrate": 257647058,
      "block_found": 0,
      "time_period": 24905720,
      "start_time": "2017-05-09T15:20:04.937319274Z"
    },
    "24905721": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149650946,
      "effective_hashrate": 190434782,
      "block_found": 0,
      "time_period": 24905721,
      "start_time": "2017-05-09T15:21:01.882353009Z"
    },
    "24905722": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149650414,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905722,
      "start_time": "2017-05-09T15:22:02.652465903Z"
    },
    "24905723": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149926377,
      "effective_hashrate": 78214285,
      "block_found": 0,
      "time_period": 24905723,
      "start_time": "2017-05-09T15:23:00.913222286Z"
    },
    "24905724": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149944296,
      "effective_hashrate": 438000000,
      "block_found": 0,
      "time_period": 24905724,
      "start_time": "2017-05-09T15:24:02.612870074Z"
    },
    "24905725": {
      "mined_share": 7,
      "valid_share": 7,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150012101,
      "effective_hashrate": 567777777,
      "block_found": 0,
      "time_period": 24905725,
      "start_time": "2017-05-09T15:25:00.064463137Z"
    },
    "24905726": {
      "mined_share": 5,
      "valid_share": 5,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150120355,
      "effective_hashrate": 413207547,
      "block_found": 0,
      "time_period": 24905726,
      "start_time": "2017-05-09T15:26:02.02187466Z"
    },
    "24905727": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149690013,
      "effective_hashrate": 279574468,
      "block_found": 0,
      "time_period": 24905727,
      "start_time": "2017-05-09T15:27:03.420557757Z"
    },
    "24905728": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150201935,
      "effective_hashrate": 87600000,
      "block_found": 0,
      "time_period": 24905728,
      "start_time": "2017-05-09T15:28:04.827951113Z"
    },
    "24905729": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150146404,
      "effective_hashrate": 305581395,
      "block_found": 0,
      "time_period": 24905729,
      "start_time": "2017-05-09T15:29:01.15465586Z"
    },
    "24905730": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150087368,
      "effective_hashrate": 318545454,
      "block_found": 0,
      "time_period": 24905730,
      "start_time": "2017-05-09T15:30:01.896779481Z"
    },
    "24905731": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150032663,
      "effective_hashrate": 97333333,
      "block_found": 0,
      "time_period": 24905731,
      "start_time": "2017-05-09T15:31:04.796946958Z"
    },
    "24905732": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150057419,
      "effective_hashrate": 230526315,
      "block_found": 0,
      "time_period": 24905732,
      "start_time": "2017-05-09T15:32:01.447723309Z"
    },
    "24905733": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150015151,
      "effective_hashrate": 250285714,
      "block_found": 0,
      "time_period": 24905733,
      "start_time": "2017-05-09T15:33:02.65816254Z"
    },
    "24905734": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 150055408,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905734,
      "start_time": "2017-05-09T15:34:04.883757792Z"
    },
    "24905735": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149343719,
      "effective_hashrate": 320487804,
      "block_found": 0,
      "time_period": 24905735,
      "start_time": "2017-05-09T15:35:01.011791263Z"
    },
    "24905736": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150017698,
      "effective_hashrate": 407441860,
      "block_found": 0,
      "time_period": 24905736,
      "start_time": "2017-05-09T15:36:02.82173246Z"
    },
    "24905737": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149905073,
      "effective_hashrate": 486666666,
      "block_found": 0,
      "time_period": 24905737,
      "start_time": "2017-05-09T15:37:04.525389837Z"
    },
    "24905738": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149485781,
      "effective_hashrate": 336923076,
      "block_found": 0,
      "time_period": 24905738,
      "start_time": "2017-05-09T15:38:01.059826328Z"
    },
    "24905739": {
      "mined_share": 5,
      "valid_share": 5,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149454258,
      "effective_hashrate": 405555555,
      "block_found": 0,
      "time_period": 24905739,
      "start_time": "2017-05-09T15:39:01.602614803Z"
    },
    "24905740": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150120915,
      "effective_hashrate": 273750000,
      "block_found": 0,
      "time_period": 24905740,
      "start_time": "2017-05-09T15:40:04.433071396Z"
    },
    "24905741": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 150161972,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905741,
      "start_time": "2017-05-09T15:41:01.19643226Z"
    },
    "24905742": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150021457,
      "effective_hashrate": 302068965,
      "block_found": 0,
      "time_period": 24905742,
      "start_time": "2017-05-09T15:42:02.637344807Z"
    },
    "24905743": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150143407,
      "effective_hashrate": 208571428,
      "block_found": 0,
      "time_period": 24905743,
      "start_time": "2017-05-09T15:43:04.329322022Z"
    },
    "24905744": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150221398,
      "effective_hashrate": 121666666,
      "block_found": 0,
      "time_period": 24905744,
      "start_time": "2017-05-09T15:44:01.035380315Z"
    },
    "24905745": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150120269,
      "effective_hashrate": 153684210,
      "block_found": 0,
      "time_period": 24905745,
      "start_time": "2017-05-09T15:45:02.799599617Z"
    },
    "24905746": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 150025581,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905746,
      "start_time": "2017-05-09T15:46:04.245572438Z"
    },
    "24905747": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149983161,
      "effective_hashrate": 178775510,
      "block_found": 0,
      "time_period": 24905747,
      "start_time": "2017-05-09T15:47:00.359187501Z"
    },
    "24905748": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149938880,
      "effective_hashrate": 876000000,
      "block_found": 0,
      "time_period": 24905748,
      "start_time": "2017-05-09T15:48:01.562089277Z"
    },
    "24905749": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150012327,
      "effective_hashrate": 2190000000,
      "block_found": 0,
      "time_period": 24905749,
      "start_time": "2017-05-09T15:49:03.438023183Z"
    },
    "24905750": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 150248601,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905750,
      "start_time": "2017-05-09T15:50:05.685493045Z"
    },
    "24905751": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149907227,
      "effective_hashrate": 89387755,
      "block_found": 0,
      "time_period": 24905751,
      "start_time": "2017-05-09T15:51:02.294321986Z"
    },
    "24905752": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149844633,
      "effective_hashrate": 730000000,
      "block_found": 0,
      "time_period": 24905752,
      "start_time": "2017-05-09T15:52:00.595994083Z"
    },
    "24905753": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149925515,
      "effective_hashrate": 243333333,
      "block_found": 0,
      "time_period": 24905753,
      "start_time": "2017-05-09T15:53:00.297858426Z"
    },
    "24905754": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150079818,
      "effective_hashrate": 186382978,
      "block_found": 0,
      "time_period": 24905754,
      "start_time": "2017-05-09T15:54:01.812753148Z"
    },
    "24905755": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149427865,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905755,
      "start_time": "2017-05-09T15:55:03.233484167Z"
    },
    "24905756": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149467504,
      "effective_hashrate": 178775510,
      "block_found": 0,
      "time_period": 24905756,
      "start_time": "2017-05-09T15:56:04.344359254Z"
    },
    "24905757": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150044520,
      "effective_hashrate": 97333333,
      "block_found": 0,
      "time_period": 24905757,
      "start_time": "2017-05-09T15:57:00.469532117Z"
    },
    "24905758": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149845362,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905758,
      "start_time": "2017-05-09T15:58:02.688769382Z"
    },
    "24905759": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149955094,
      "effective_hashrate": 262800000,
      "block_found": 0,
      "time_period": 24905759,
      "start_time": "2017-05-09T15:59:03.942892493Z"
    },
    "24905760": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 150107201,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905760,
      "start_time": "2017-05-09T16:00:00.159550948Z"
    },
    "24905761": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149918527,
      "effective_hashrate": 282580645,
      "block_found": 0,
      "time_period": 24905761,
      "start_time": "2017-05-09T16:01:01.531451781Z"
    },
    "24905762": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149751206,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905762,
      "start_time": "2017-05-09T16:02:02.57138612Z"
    },
    "24905763": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150032215,
      "effective_hashrate": 243333333,
      "block_found": 0,
      "time_period": 24905763,
      "start_time": "2017-05-09T16:03:03.945580156Z"
    },
    "24905764": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149993981,
      "effective_hashrate": 190434782,
      "block_found": 0,
      "time_period": 24905764,
      "start_time": "2017-05-09T16:04:00.368725027Z"
    },
    "24905765": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150000645,
      "effective_hashrate": 243333333,
      "block_found": 0,
      "time_period": 24905765,
      "start_time": "2017-05-09T16:05:02.189686753Z"
    },
    "24905766": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149417470,
      "effective_hashrate": 312857142,
      "block_found": 0,
      "time_period": 24905766,
      "start_time": "2017-05-09T16:06:03.37864691Z"
    },
    "24905767": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149251360,
      "effective_hashrate": 282580645,
      "block_found": 0,
      "time_period": 24905767,
      "start_time": "2017-05-09T16:07:00.004413103Z"
    },
    "24905768": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149967732,
      "effective_hashrate": 159272727,
      "block_found": 0,
      "time_period": 24905768,
      "start_time": "2017-05-09T16:08:01.628077865Z"
    },
    "24905769": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149991710,
      "effective_hashrate": 87600000,
      "block_found": 0,
      "time_period": 24905769,
      "start_time": "2017-05-09T16:09:02.444480192Z"
    },
    "24905770": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149430842,
      "effective_hashrate": 243333333,
      "block_found": 0,
      "time_period": 24905770,
      "start_time": "2017-05-09T16:10:03.333266188Z"
    },
    "24905771": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150111210,
      "effective_hashrate": 547500000,
      "block_found": 0,
      "time_period": 24905771,
      "start_time": "2017-05-09T16:11:04.463939395Z"
    },
    "24905772": {
      "mined_share": 5,
      "valid_share": 5,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149993684,
      "effective_hashrate": 438000000,
      "block_found": 0,
      "time_period": 24905772,
      "start_time": "2017-05-09T16:12:01.150001596Z"
    },
    "24905773": {
      "mined_share": 6,
      "valid_share": 6,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149862500,
      "effective_hashrate": 691578947,
      "block_found": 0,
      "time_period": 24905773,
      "start_time": "2017-05-09T16:13:02.430144615Z"
    },
    "24905774": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150074404,
      "effective_hashrate": 320487804,
      "block_found": 0,
      "time_period": 24905774,
      "start_time": "2017-05-09T16:14:03.633201791Z"
    },
    "24905775": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150023002,
      "effective_hashrate": 175200000,
      "block_found": 0,
      "time_period": 24905775,
      "start_time": "2017-05-09T16:15:04.646930108Z"
    },
    "24905776": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149995644,
      "effective_hashrate": 136875000,
      "block_found": 0,
      "time_period": 24905776,
      "start_time": "2017-05-09T16:16:01.430489434Z"
    },
    "24905777": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149580164,
      "effective_hashrate": 279574468,
      "block_found": 0,
      "time_period": 24905777,
      "start_time": "2017-05-09T16:17:04.090378751Z"
    },
    "24905778": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149794143,
      "effective_hashrate": 84230769,
      "block_found": 0,
      "time_period": 24905778,
      "start_time": "2017-05-09T16:18:00.928051797Z"
    },
    "24905779": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149935595,
      "effective_hashrate": 234642857,
      "block_found": 0,
      "time_period": 24905779,
      "start_time": "2017-05-09T16:19:01.978016966Z"
    },
    "24905780": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150011340,
      "effective_hashrate": 625714285,
      "block_found": 0,
      "time_period": 24905780,
      "start_time": "2017-05-09T16:20:04.473047079Z"
    },
    "24905781": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150067365,
      "effective_hashrate": 132727272,
      "block_found": 0,
      "time_period": 24905781,
      "start_time": "2017-05-09T16:21:01.082010712Z"
    },
    "24905782": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149906097,
      "effective_hashrate": 168461538,
      "block_found": 0,
      "time_period": 24905782,
      "start_time": "2017-05-09T16:22:02.755711331Z"
    },
    "24905783": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149954217,
      "effective_hashrate": 417142857,
      "block_found": 0,
      "time_period": 24905783,
      "start_time": "2017-05-09T16:23:04.469626068Z"
    },
    "24905784": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149414090,
      "effective_hashrate": 243333333,
      "block_found": 0,
      "time_period": 24905784,
      "start_time": "2017-05-09T16:24:01.015801947Z"
    },
    "24905785": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150179006,
      "effective_hashrate": 82641509,
      "block_found": 0,
      "time_period": 24905785,
      "start_time": "2017-05-09T16:25:02.868159311Z"
    },
    "24905786": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149840523,
      "effective_hashrate": 389333333,
      "block_found": 0,
      "time_period": 24905786,
      "start_time": "2017-05-09T16:26:04.354419028Z"
    },
    "24905787": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149949542,
      "effective_hashrate": 156428571,
      "block_found": 0,
      "time_period": 24905787,
      "start_time": "2017-05-09T16:27:00.682953473Z"
    },
    "24905788": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150058186,
      "effective_hashrate": 171764705,
      "block_found": 0,
      "time_period": 24905788,
      "start_time": "2017-05-09T16:28:02.804729477Z"
    },
    "24905789": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150034848,
      "effective_hashrate": 398181818,
      "block_found": 0,
      "time_period": 24905789,
      "start_time": "2017-05-09T16:29:04.244924543Z"
    },
    "24905790": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150169720,
      "effective_hashrate": 320487804,
      "block_found": 0,
      "time_period": 24905790,
      "start_time": "2017-05-09T16:30:00.561643408Z"
    },
    "24905791": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149190733,
      "effective_hashrate": 292000000,
      "block_found": 0,
      "time_period": 24905791,
      "start_time": "2017-05-09T16:31:02.136949554Z"
    },
    "24905792": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149926527,
      "effective_hashrate": 125142857,
      "block_found": 0,
      "time_period": 24905792,
      "start_time": "2017-05-09T16:32:03.302626137Z"
    },
    "24905793": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149801308,
      "effective_hashrate": 178775510,
      "block_found": 0,
      "time_period": 24905793,
      "start_time": "2017-05-09T16:33:05.077724168Z"
    },
    "24905794": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149862821,
      "effective_hashrate": 171764705,
      "block_found": 0,
      "time_period": 24905794,
      "start_time": "2017-05-09T16:34:01.403604951Z"
    },
    "24905795": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 150027169,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905795,
      "start_time": "2017-05-09T16:35:02.672223793Z"
    },
    "24905796": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 148838767,
      "effective_hashrate": 226551724,
      "block_found": 0,
      "time_period": 24905796,
      "start_time": "2017-05-09T16:36:01.328651401Z"
    },
    "24905797": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149879547,
      "effective_hashrate": 75517241,
      "block_found": 0,
      "time_period": 24905797,
      "start_time": "2017-05-09T16:37:00.403607398Z"
    },
    "24905798": {
      "mined_share": 5,
      "valid_share": 5,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149487073,
      "effective_hashrate": 521428571,
      "block_found": 0,
      "time_period": 24905798,
      "start_time": "2017-05-09T16:38:01.496041334Z"
    },
    "24905799": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149832267,
      "effective_hashrate": 156428571,
      "block_found": 0,
      "time_period": 24905799,
      "start_time": "2017-05-09T16:39:03.13336059Z"
    },
    "24905800": {
      "mined_share": 6,
      "valid_share": 6,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149652402,
      "effective_hashrate": 486666666,
      "block_found": 0,
      "time_period": 24905800,
      "start_time": "2017-05-09T16:40:03.15538684Z"
    },
    "24905801": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149681799,
      "effective_hashrate": 257647058,
      "block_found": 0,
      "time_period": 24905801,
      "start_time": "2017-05-09T16:41:01.749104048Z"
    },
    "24905802": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149815321,
      "effective_hashrate": 625714285,
      "block_found": 0,
      "time_period": 24905802,
      "start_time": "2017-05-09T16:42:03.012391112Z"
    },
    "24905803": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149778509,
      "effective_hashrate": 292000000,
      "block_found": 0,
      "time_period": 24905803,
      "start_time": "2017-05-09T16:43:00.078834955Z"
    },
    "24905804": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149633727,
      "effective_hashrate": 238909090,
      "block_found": 0,
      "time_period": 24905804,
      "start_time": "2017-05-09T16:44:01.874319927Z"
    },
    "24905805": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149555651,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905805,
      "start_time": "2017-05-09T16:45:03.936010301Z"
    },
    "24905806": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149772875,
      "effective_hashrate": 486666666,
      "block_found": 0,
      "time_period": 24905806,
      "start_time": "2017-05-09T16:46:00.272641049Z"
    },
    "24905807": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149234376,
      "effective_hashrate": 153684210,
      "block_found": 0,
      "time_period": 24905807,
      "start_time": "2017-05-09T16:47:01.888109598Z"
    },
    "24905808": {
      "mined_share": 4,
      "valid_share": 4,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149643713,
      "effective_hashrate": 312857142,
      "block_found": 0,
      "time_period": 24905808,
      "start_time": "2017-05-09T16:48:03.545967202Z"
    },
    "24905809": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149467453,
      "effective_hashrate": 547500000,
      "block_found": 0,
      "time_period": 24905809,
      "start_time": "2017-05-09T16:49:00.000557063Z"
    },
    "24905810": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149203244,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905810,
      "start_time": "2017-05-09T16:50:01.687095376Z"
    },
    "24905811": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149852085,
      "effective_hashrate": 84230769,
      "block_found": 0,
      "time_period": 24905811,
      "start_time": "2017-05-09T16:51:03.848908866Z"
    },
    "24905812": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149743399,
      "effective_hashrate": 112307692,
      "block_found": 0,
      "time_period": 24905812,
      "start_time": "2017-05-09T16:52:01.504950325Z"
    },
    "24905813": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149592633,
      "effective_hashrate": 298636363,
      "block_found": 0,
      "time_period": 24905813,
      "start_time": "2017-05-09T16:53:04.365886551Z"
    },
    "24905814": {
      "mined_share": 7,
      "valid_share": 7,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149679193,
      "effective_hashrate": 547500000,
      "block_found": 1,
      "time_period": 24905814,
      "start_time": "2017-05-09T16:54:00.98311751Z"
    },
    "24905815": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149805715,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905815,
      "start_time": "2017-05-09T16:55:04.36661599Z"
    },
    "24905816": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149674532,
      "effective_hashrate": 250285714,
      "block_found": 0,
      "time_period": 24905816,
      "start_time": "2017-05-09T16:56:01.739933648Z"
    },
    "24905817": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149768695,
      "effective_hashrate": 584000000,
      "block_found": 0,
      "time_period": 24905817,
      "start_time": "2017-05-09T16:57:03.896193415Z"
    },
    "24905818": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149805592,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905818,
      "start_time": "2017-05-09T16:58:01.541064989Z"
    },
    "24905819": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149899120,
      "effective_hashrate": 262800000,
      "block_found": 0,
      "time_period": 24905819,
      "start_time": "2017-05-09T16:59:05.039207578Z"
    },
    "24905820": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149927160,
      "effective_hashrate": 171764705,
      "block_found": 0,
      "time_period": 24905820,
      "start_time": "2017-05-09T17:00:02.864919173Z"
    },
    "24905821": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149787613,
      "effective_hashrate": 257647058,
      "block_found": 0,
      "time_period": 24905821,
      "start_time": "2017-05-09T17:01:00.587387399Z"
    },
    "24905822": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149792973,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905822,
      "start_time": "2017-05-09T17:02:03.573013126Z"
    },
    "24905823": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149914985,
      "effective_hashrate": 336923076,
      "block_found": 0,
      "time_period": 24905823,
      "start_time": "2017-05-09T17:03:01.258804189Z"
    },
    "24905824": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149820506,
      "effective_hashrate": 101860465,
      "block_found": 0,
      "time_period": 24905824,
      "start_time": "2017-05-09T17:04:04.321582459Z"
    },
    "24905825": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149795514,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905825,
      "start_time": "2017-05-09T17:05:02.087228015Z"
    },
    "24905826": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149866246,
      "effective_hashrate": 199090909,
      "block_found": 0,
      "time_period": 24905826,
      "start_time": "2017-05-09T17:06:04.102838333Z"
    },
    "24905827": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149942432,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905827,
      "start_time": "2017-05-09T17:07:00.868891928Z"
    },
    "24905828": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149788685,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905828,
      "start_time": "2017-05-09T17:08:02.237683864Z"
    },
    "24905829": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 150071614,
      "effective_hashrate": 151034482,
      "block_found": 0,
      "time_period": 24905829,
      "start_time": "2017-05-09T17:09:03.864418531Z"
    },
    "24905830": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149381546,
      "effective_hashrate": 365000000,
      "block_found": 0,
      "time_period": 24905830,
      "start_time": "2017-05-09T17:10:00.441615499Z"
    },
    "24905831": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149341908,
      "effective_hashrate": 146000000,
      "block_found": 0,
      "time_period": 24905831,
      "start_time": "2017-05-09T17:11:02.363053509Z"
    },
    "24905832": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149527504,
      "effective_hashrate": 162222222,
      "block_found": 0,
      "time_period": 24905832,
      "start_time": "2017-05-09T17:12:03.835007241Z"
    },
    "24905833": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149410231,
      "effective_hashrate": 89387755,
      "block_found": 0,
      "time_period": 24905833,
      "start_time": "2017-05-09T17:13:00.232218978Z"
    },
    "24905834": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149981254,
      "effective_hashrate": 165283018,
      "block_found": 0,
      "time_period": 24905834,
      "start_time": "2017-05-09T17:14:02.448299519Z"
    },
    "24905835": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149836303,
      "effective_hashrate": 324444444,
      "block_found": 0,
      "time_period": 24905835,
      "start_time": "2017-05-09T17:15:03.929021946Z"
    },
    "24905836": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149368046,
      "effective_hashrate": 156428571,
      "block_found": 0,
      "time_period": 24905836,
      "start_time": "2017-05-09T17:16:04.962565117Z"
    },
    "24905837": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149900194,
      "effective_hashrate": 171764705,
      "block_found": 0,
      "time_period": 24905837,
      "start_time": "2017-05-09T17:17:01.354205011Z"
    },
    "24905838": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149857187,
      "effective_hashrate": 938571428,
      "block_found": 0,
      "time_period": 24905838,
      "start_time": "2017-05-09T17:18:00.992187217Z"
    },
    "24905839": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149699991,
      "effective_hashrate": 104285714,
      "block_found": 0,
      "time_period": 24905839,
      "start_time": "2017-05-09T17:19:00.230552106Z"
    },
    "24905840": {
      "mined_share": 0,
      "valid_share": 0,
      "average_share_difficulty": 0,
      "rejected_share": 0,
      "reported_hashrate": 149708763,
      "effective_hashrate": 0,
      "block_found": 0,
      "time_period": 24905840,
      "start_time": "2017-05-09T17:20:02.057575079Z"
    },
    "24905841": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149610996,
      "effective_hashrate": 438000000,
      "block_found": 0,
      "time_period": 24905841,
      "start_time": "2017-05-09T17:21:03.526802898Z"
    },
    "24905842": {
      "mined_share": 2,
      "valid_share": 2,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149336310,
      "effective_hashrate": 153684210,
      "block_found": 0,
      "time_period": 24905842,
      "start_time": "2017-05-09T17:22:00.561862253Z"
    },
    "24905843": {
      "mined_share": 1,
      "valid_share": 1,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 148720380,
      "effective_hashrate": 547500000,
      "block_found": 0,
      "time_period": 24905843,
      "start_time": "2017-05-09T17:23:02.214242705Z"
    },
    "24905844": {
      "mined_share": 3,
      "valid_share": 3,
      "average_share_difficulty": 4380000000,
      "rejected_share": 0,
      "reported_hashrate": 149684285,
      "effective_hashrate": 772941176,
      "block_found": 0,
      "time_period": 24905844,
      "start_time": "2017-05-09T17:24:04.058779359Z"
    }
  },
  "overall": {
    "last_mined_share": "2017-05-09T17:24:21.571703772Z",
    "last_valid_share": "2017-05-09T17:24:21.57170827Z",
    "last_rejected_share": "2017-05-08T18:29:32.544949808Z",
    "last_block": "2017-05-09T16:54:00.9831363Z",
    "total_submitted_share": 5386,
    "total_accepted_share": 5385,
    "total_accepted_difficulty": 23586300000000,
    "average_share_difficulty": 4380000000,
    "total_rejected_share": 1,
    "total_hashrate": 4790657349842,
    "no_hashrate_submission": 32023,
    "average_reported_hashrate": 149600516,
    "average_effective_hashrate": 142831276,
    "total_block_found": 1,
    "start_time": "2017-05-07T19:32:07.530342074Z"
  },
  "period_duration": 60
}
	`
}
