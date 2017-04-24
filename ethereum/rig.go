package ethereum

type Rig struct {
	name string
}

func (r *Rig) ID() string {
	return r.name
}

func NewRig(name string) *Rig {
	return &Rig{name}
}
