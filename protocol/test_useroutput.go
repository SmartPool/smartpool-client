package protocol

type testUserOutput struct{}

func (out *testUserOutput) Printf(format string, a ...interface{}) (n int, err error) {
	return 0, nil
}
