package scene

import (
	"image"
	"math"

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

		// start moving to target
		u.action = newAction(actionMoveToNearestTarget, u)

	case actionMoveToNearestTarget:
		u.moveToTarget(u.game.player)

	default:
		// nop
	}
}

func (u *sampleUnit) moveToTarget(target uniter) {
	// get my position
	ux, uy := u.sprite.X, u.sprite.Y

	// get target (player's) position
	p := target.GetPosition()
	px, py := (float32)(p.x), (float32)(p.y)

	// calculate which way to go
	// move speed is temporary
	dx, dy := px-ux, py-uy
	newx := (float64)(u.moveSpeed) / math.Sqrt((float64)(dx*dx+dy*dy)) * (float64)(dx)
	newy := (float64)(u.moveSpeed) / math.Sqrt((float64)(dx*dx+dy*dy)) * (float64)(dy)
	u.sprite.X += (float32)(newx)
	u.sprite.Y += (float32)(newy)
}
