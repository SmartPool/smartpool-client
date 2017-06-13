package protocol

import (
	"math/big"
)

type testPersistentStorage struct{}

func (self *testPersistentStorage) Persist(data interface{}, id string) error {
	return nil
}
func (self *testPersistentStorage) Load(data interface{}, id string) (interface{}, error) {
	if id == COUNTER_FILE {
		return big.NewInt(0), nil
	}
	return nil, nil
}
