package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
)

type player struct {
	*unitBase
	hp int
}

func (u *player) Initialize() {
	simra.GetInstance().AddSprite("player.png",
		image.Rect(0, 0, 384, 384),
		&u.sprite)
	u.hp = 100
}

func (u *player) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandSpawn:
		d, ok := c.data.(*player)
		if !ok {
			// unhandled event. ignore
			return
		}
		if u.id != d.GetID() {
			// this spawn event is not for me. nop.
			return
		}
		u.action = newAction(actionSpawn, d)

	case commandDamage:
		d, ok := c.data.(*damage)
		if !ok {
			return
		}
		if u.id != d.unit.GetID() {
			return
		}

		// TODO: reduce HP of unit
		u.hp -= d.damage
		simra.LogDebug("[DAMAGE] i'm [%s], HP = %d", u.GetID(), u.hp)
		if u.hp <= 0 {
			simra.LogDebug("[DEAD] i'm %s", u.GetID())
			u.game.eventqueue <- newCommand(commandDead, u)
			u.action = newAction(actionDead, nil)
		}

	case commandDead:
		d, ok := c.data.(*player)
		if !ok {
			// unhandled event. ignore
			return
		}
		if u.id != d.GetID() {
			// this dead event is not for me. nop.
			return
		}

	default:
		// nop
	}
}

func (u *player) DoAction() {
	a := u.action
	if a == nil {
		// idle
		return
	}

	switch a.actiontype {
	case actionSpawn:
		d := a.data.(*player)
		u.sprite.W = 64
		u.sprite.H = 64
		u.SetPosition(d.GetPosition())
		simra.LogDebug("@@@@@@ [SPAWN] i'm %s", u.GetID())
		u.action = nil

	case actionDead:
		// i'm dead!
		u.sprite.W = 1
		u.sprite.H = 1
		u.SetPosition(-1, -1)
		simra.LogDebug("@@@@@@ [DEAD] i'm %s", u.GetID())
		u.action = nil
		delete(u.game.players, u.GetID())

	default:
		// nop
	}
}
