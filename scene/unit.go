package scene

import (
	"sync"

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
	DoAction()
}

type position struct {
	x int
	y int
}

type actiontype int

const (
	// SPAWN spawns an unit
	actionSpawn actiontype = iota
)

type action struct {
	actiontype actiontype
	data       interface{}
}

func newAction(a actiontype, d interface{}) *action {
	return &action{
		actiontype: a,
		data:       d,
	}
}

type unitBase struct {
	simra.Subscriber
	id       string
	position position
	action   *action
	mutex    sync.Mutex
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
	commandSpawn commandtype = iota
)

type command struct {
	commandtype commandtype
	data        interface{}
}

func newCommand() *command {
	return &command{}
}
