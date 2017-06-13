package protocol

import (
	"github.com/SmartPool/smartpool-client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type testStatRecorder struct {
}

func (self *testStatRecorder) RecordShare(status string, share smartpool.Share, rig smartpool.Rig) {
}

func (self *testStatRecorder) RecordClaim(status string, claim smartpool.Claim) {
}
func (self *testStatRecorder) RecordHashrate(hashrate hexutil.Uint64, id common.Hash, rig smartpool.Rig) {
}
func (self *testStatRecorder) ShareRestored(noshares uint64) {
}

func (self *testStatRecorder) OverallFarmStat() interface{} {
	return nil
}
func (self *testStatRecorder) FarmStat(start uint64, end uint64) interface{} {
	return nil
}
func (self *testStatRecorder) OverallRigStat(rig smartpool.Rig) interface{} {
	return nil
}
func (self *testStatRecorder) RigStat(rig smartpool.Rig, start uint64, end uint64) interface{} {
	return nil
}

func (self *testStatRecorder) Persist(storage smartpool.PersistentStorage) error {
	return nil
}
