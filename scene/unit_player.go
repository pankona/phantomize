package scene

import (
	"image"
)

type player struct {
	*unitBase
	delayTimeToRecall   int
	elapsedTimeToRecall int
}

func (u *player) Initialize() {
	assetName := u.game.assetNameByUnitType(u.unittype)
	u.simra.AddSprite(u.sprite)
	tex := u.simra.NewImageTexture(assetName,
		image.Rect(0, 0, 384, 384))
	u.sprite.ReplaceTexture(tex)
}

func (u *player) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandRecalled:
		d := c.data.(uniter)
		if u.GetID() != d.GetID() {
			// this is not for me. ignore
			break
		}

		u.simra.RemoveSprite(u.sprite)
		killUnit(u, u.game.players)

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
			u.game.eventqueue <- newCommand(commandAttack, u)
		}

	case actionDead:
		// i'm dead!
		killUnit(u, u.game.players)

	case actionRecall:
		u.elapsedTimeToRecall++
		if u.elapsedTimeToRecall <= u.delayTimeToRecall {
			// still summoning...
			break
		}
		u.elapsedTimeToRecall = 0
		u.game.eventqueue <- newCommand(commandRecalled, u)

	default:
		u.unitBase.doAction(a)
	}
}
