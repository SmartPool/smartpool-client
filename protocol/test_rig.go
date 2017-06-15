package protocol

type testRig struct {
}

func (r *testRig) ID() string {
	return "rig"
}

func (r *testRig) IP() string {
	return "127.0.0.1"
}
