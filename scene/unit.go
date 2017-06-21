package scene

import (
	"github.com/pankona/gomo-simra/simra"
)

// unit base implementation

// Uniter is an interface of unit
type Uniter interface {
	Initialize()
	GetID() string
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

func (u *unitBase) SetPosition(p position) {
	u.position = p
}

func (u *unitBase) GetPosition() position {
	return u.position
}

// NewUnit returns a uniter
func NewUnit(id, unittype string) Uniter {
	// TODO: sample unit implemenation
	// unit type should be specified and switch here
	var u Uniter
	switch unittype {
	case "player":
		u = &player{unitBase: &unitBase{id: id}}
	default:
		// TODO: remove later
		u = &sampleUnit{unitBase: &unitBase{id: id}}
	}
	u.Initialize()
	return u
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
