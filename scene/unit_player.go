package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
)

type player struct {
	*unitBase
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
	case commandDead:
		if len(u.game.uniters) == 0 {
			// all enemies are eliminated
			simra.LogDebug("you won!")
			u.action = nil
			break
		}

	default:
		u.unitBase.onEvent(c)
	}
}

func (u *player) DoAction() {
	a := u.action
	if a == nil {
		// idle
		return
	}

	switch a.actiontype {
	case actionMoveToNearestTarget:
		u.target = nearestUnit(u.unitBase, u.game.uniters)
		if u.target == nil {
			break
		}
		moveToTarget(u.unitBase, u.target)

		if canAttackToTarget(u.unitBase, u.target) {
			u.action = newAction(actionAttack, u.target)
		}

	case actionDead:
		// i'm dead!
		u.sprite.W = 1
		u.sprite.H = 1
		u.SetPosition(-1, -1)
		simra.LogDebug("@@@@@@ [DEAD] i'm %s", u.GetID())
		u.action = nil
		u.isSpawned = false
		delete(u.game.players, u.GetID())

	default:
		u.unitBase.doAction(a)
	}
}
