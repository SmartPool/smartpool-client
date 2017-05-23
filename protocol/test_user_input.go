package protocol

import (
	"math/big"
	"time"
)

type testUserInput struct {
}

func (self *testUserInput) RPCEndpoint() string {
	return "http://localhost:8545"
}
func (self *testUserInput) KeystorePath() string {
	return "keystore/path"
}
func (self *testUserInput) ShareThreshold() int {
	return 100
}
func (self *testUserInput) ShareDifficulty() *big.Int {
	return big.NewInt(1000000)
}
func (self *testUserInput) SubmitInterval() time.Duration {
	return time.Minute
}
func (self *testUserInput) ContractAddress() string {
	return "0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F"
}
func (self *testUserInput) MinerAddress() string {
	return "0x001aDBc838eDe392B5B054A47f8B8c28f2fA9F3F"
}
func (self *testUserInput) ExtraData() string {
	return "extra"
}
func (self *testUserInput) HotStop() bool {
	return true
}
