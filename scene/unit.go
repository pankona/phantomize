package scene

import (
	"fmt"

	"github.com/pankona/gomo-simra/simra"
)

// unit base implementation

// Uniter is an interface of unit
type Uniter interface {
	GetID() string
	GetBase() *unitBase
	simra.Subscriber
}

type unitBase struct {
	id string
	simra.Subscriber
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

type sampleUnit struct {
	*unitBase
}

func (u *sampleUnit) GetBase() *unitBase {
	return u.unitBase
}

func (u *sampleUnit) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case SPAWN:
		u, ok := c.data.(*sampleUnit)
		if !ok {
			panic("unexpected command received. fatal.")
		}
		// TODO: spawn myself
		fmt.Printf("@@@@@@ [SPAWN] i'm %s\n", u.GetID())
	default:
		// nop
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
