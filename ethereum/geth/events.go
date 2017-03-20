package geth

import (
	"fmt"
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var RegisterEventTopic = common.HexToHash("0x1d759fb22634fe2d322d688a4b46aaf185dd0a3db78ccf01a9218f00ac3df03f").Big()
var SubmitClaimEventTopic = common.HexToHash("0x53ab9d877ae22286591454f9a8d58501caa34a07c99eac2c09bc0066c065400d").Big()
var VerifyClaimEventTopic = common.HexToHash("0x096caf97202169a068288f02e51ff9fcc85f98e1477f6ad9acbf6ebf25dbcd00").Big()
var SetEpochDataEventTopic = common.HexToHash("0x5cd723400be8430351b9cbaa5ea421b3fb2528c6a7650c493f895e7d97750da1").Big()

var ErrorMap = map[uint64][]string{
	0x80000000: []string{"register", "miner id is already in use", "miner id"},
	0x80000001: []string{"register", "payment address 0 is forbiden", ""},
	0x81000000: []string{"submitClaim", "miner is not registered", ""},
	0x81000001: []string{"submitClaim", "min counter is too low", "last submission counter value"},
	0x82000000: []string{"setEpochData", "only owner can set data", "msg.sender"},
	0x82000001: []string{"setEpochData", "epoch already set", "epoch number"},
	0x83000000: []string{"verifyExtraData", "miner id not as expected", "miner id"},
	0x83000001: []string{"verifyExtraData", "difficulty is not as expected", "encoded difficulty"},
	0x84000000: []string{"verifyClaim", "contract balance is too low to pay", "payment"},
	0x84000001: []string{"verifyClaim", "claim seed is 0", ""},
	0x84000002: []string{"verifyClaim", "share index is not as expected", "expected index"},
	0x84000003: []string{"verifyClaim", "there are no pending claims", ""},
	0x84000004: []string{"verifyClaim", "extra data not as expected", "extra data"},
	0x84000005: []string{"verifyClaim", "coinbase not as expected", "coinbase"},
	0x84000006: []string{"verifyClaim", "counter is smaller than min", "counter"},
	0x84000007: []string{"verifyClaim", "counter is smaller than max", "counter"},
	0x84000008: []string{"verifyClaim", "verification of augmented merkle tree failed", ""},
	0x84000009: []string{"verifyClaim", "ethash difficulty too low (or hashimoto verification failed)", "computed ethash value"},
	0x8400000a: []string{"verifyClaim", "epoch data was not set", "epoch number"},
}

func ErrorMsg(errCode, errInfo *big.Int) string {
	infos := ErrorMap[errCode.Uint64()]
	if len(infos) == 0 {
		smartpool.Output.Printf("Invalid errCode(0x%s)\n", errCode.Text(16))
		return ""
	} else if len(infos) != 3 {
		smartpool.Output.Printf("ErrorMap is not welformed for errCode(0x%s)\n", errCode.Text(16))
		return ""
	} else {
		msg := infos[1]
		function := infos[0]
		name := infos[2]
		return fmt.Sprintf(
			"Contract returned error code 0x%s. \n%s: %s. %s is 0x%s\n",
			errCode.Text(16),
			function,
			msg,
			name,
			errInfo.Text(16),
		)
	}
}
