package scene

import (
	"image"
	"math"

	"github.com/pankona/gomo-simra/simra"
)

type sampleUnit struct {
	*unitBase
	attackinfo *attackInfo
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

	case commandDead:
		_, ok := c.data.(*player)
		if !ok {
			return
		}
		simra.LogDebug("i'm %s, we won!", u.GetID())
		u.action = nil
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
		u.SetPosition(d.GetPosition())
		simra.LogDebug("@@@@@@ [SPAWN] i'm %s", u.GetID())

		// start moving to target
		u.action = newAction(actionMoveToNearestTarget, nil)

	case actionMoveToNearestTarget:
		// TODO: lookup nearest target
		u.moveToTarget(u.game.player)

		if u.canAttackToTarget(u.game.player) {
			u.action = newAction(actionAttack, u.game.player)
		}

	case actionAttack:
		// TODO: start animation

		if !u.canAttackToTarget(u.game.player) {
			u.action = newAction(actionMoveToNearestTarget, nil)
			break
		}

		if u.game.currentFrame-u.attackinfo.lastAttackTime >=
			(int64)(u.attackinfo.cooltime*fps) {
			simra.LogDebug("[ATTACK] i'm %s", u.GetID())
			u.attackinfo.lastAttackTime = u.game.currentFrame

			u.game.eventqueue <- newCommand(commandDamage, &damage{u.game.player, u.attackinfo.power})

		}
	default:
		// nop
	}
}

func (u *sampleUnit) moveToTarget(target uniter) {
	ux, uy := u.GetPosition()
	tx, ty := target.GetPosition()

	// calculate which way to go
	// move speed is temporary
	dx, dy := tx-ux, ty-uy
	newx := (float64)(u.moveSpeed) / getDistance(ux, uy, tx, ty) * (float64)(dx)
	newy := (float64)(u.moveSpeed) / getDistance(ux, uy, tx, ty) * (float64)(dy)
	u.sprite.X += (float32)(newx)
	u.sprite.Y += (float32)(newy)
}

func (u *sampleUnit) canAttackToTarget(target uniter) bool {
	ux, uy := u.GetPosition()
	tx, ty := target.GetPosition()

	if (float64)(u.attackinfo.attackRange) >= getDistance(ux, uy, tx, ty) {
		return true
	}
	return false
}

func getDistance(ax, ay, bx, by float32) float64 {
	dx, dy := ax-bx, ay-by
	return math.Sqrt((float64)(dx*dx + dy*dy))
}
