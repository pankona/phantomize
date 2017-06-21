package scene

import (
	"fmt"
)

type sampleUnit struct {
	*unitBase
}

func (u *sampleUnit) GetBase() *unitBase {
	return u.unitBase
}

func (u *sampleUnit) SetPosition(p position) {
	u.position = p
}

func (u *sampleUnit) GetPosition() position {
	return u.position
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
