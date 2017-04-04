package protocol

import (
	"math/big"
)

type testPersistentStorage struct{}

func (*testPersistentStorage) PersistLatestCounter(counter *big.Int) error {
	return nil
}
func (*testPersistentStorage) LoadLatestCounter() (*big.Int, error) {
	return big.NewInt(0), nil
}
