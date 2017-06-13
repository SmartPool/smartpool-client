package ethereum

type Rig struct {
	name string
	ip   string
}

func (r *Rig) ID() string {
	return r.name
}

func (r *Rig) IP() string {
	return r.ip
}

func NewRig(name string, ip string) *Rig {
	return &Rig{name, ip}
}
