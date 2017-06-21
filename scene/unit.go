package scene

import (
	"github.com/pankona/gomo-simra/simra"
)

// unit base implementation

// Uniter is an interface of unit
type Uniter interface {
	GetID() string
	GetBase() *unitBase
	simra.Subscriber
	SetPosition(p position)
	GetPosition() position
}

type position struct {
	x int
	y int
}

type unitBase struct {
	id string
	simra.Subscriber
	position position
}

func (u *unitBase) GetID() string {
	return u.id
}

func (u *unitBase) GetBase() *unitBase {
	// nop
	return nil
}

// NewUnit returns a uniter
func NewUnit(id, unittype string) Uniter {
	// TODO: sample unit implemenation
	// unit type should be specified and switch here
	switch unittype {
	default:
		// TODO: remove later
		return &sampleUnit{unitBase: &unitBase{id: id}}
	}
}

type commandtype int

const (
	// SPAWN spawns an unit
	SPAWN commandtype = iota
)

type command struct {
	commandtype commandtype
	data        interface{}
}

func newCommand() *command {
	return &command{}
}
