package event

type ShareEventStorage interface {
	Add(rig smartpool.Rig)
	FarmStats(start uint64, end uint64) interface{}
	RigStats(rig Rig, start uint64, end uint64) interface{}
}
