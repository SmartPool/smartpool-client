package protocol

type testRig struct {
}

func (r *testRig) ID() string {
	return "rig-127.0.0.1"
}

func (r *testRig) IP() string {
	return "127.0.0.1"
}

func (r *testRig) Name() string {
	return "rig"
}
