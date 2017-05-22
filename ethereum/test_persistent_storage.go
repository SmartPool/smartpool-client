package ethereum

type testPersistentStorage struct{}

func (self *testPersistentStorage) Persist(data interface{}, id string) error {
	return nil
}
func (self *testPersistentStorage) Load(data interface{}, id string) (interface{}, error) {
	if id == ACTIVE_SHARE_FILE {
		return &map[string]gobShare{}, nil
	}
	return nil, nil
}
