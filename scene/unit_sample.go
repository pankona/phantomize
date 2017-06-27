package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
)

type sampleUnit struct {
	*unitBase
	sprite simra.Sprite
}

func (u *sampleUnit) Initialize() {
	simra.GetInstance().AddSprite("player.png",
		image.Rect(0, 0, 384, 384),
		&u.sprite)
}

func (u *sampleUnit) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandSpawn:
		d, ok := c.data.(*sampleUnit)
		if !ok {
			// unhandled event. ignore.
			return
		}
		if u.id != d.GetID() {
			// this spawn event is not for me. nop.
			return
		}
		u.action = newAction(actionSpawn, d)

	default:
		// nop
	}
}

func (u *sampleUnit) DoAction() {
	a := u.action
	if a == nil {
		// idle
		return
	}

	switch a.actiontype {
	case actionSpawn:
		d := a.data.(*sampleUnit)
		u.sprite.W = 64
		u.sprite.H = 64
		u.sprite.X = (float32)(d.position.x)
		u.sprite.Y = (float32)(d.position.y)
		simra.LogDebug("@@@@@@ [SPAWN] i'm %s", u.GetID())
		u.action = nil

		// TODO start to lookup player
	default:
		// nop
	}
}
