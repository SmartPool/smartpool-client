package ethereum

import (
	"fmt"
)

type Rig struct {
	name string
	ip   string
}

func (r *Rig) ID() string {
	return fmt.Sprintf("%s-%s", r.name, r.ip)
}

func (r *Rig) IP() string {
	return r.ip
}

func (r *Rig) Name() string {
	return r.name
}

func NewRig(name string, ip string) *Rig {
	return &Rig{name, ip}
}
