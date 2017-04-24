package event

type HashrateEventStorage interface {
	Add(rig smartpool.Rig, hashrate hexutil.Uint64, id common.Hash)
}
