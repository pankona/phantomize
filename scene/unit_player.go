package scene

import (
	"fmt"
	"image"

	"github.com/pankona/gomo-simra/simra"
)

type player struct {
	*unitBase
}

func (u *player) Initialize() {
	assetName := u.game.assetNameByUnitType(u.unittype)
	simra.GetInstance().AddSprite(assetName,
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
	case commandAttack:
		d := c.data.(uniter)
		if u.GetID() != d.GetID() {
			// this is not for me. ignore
			break
		}
		// TODO: load in advance. don't do every time.
		texName := fmt.Sprintf("%s_atk.png", u.GetUnitType())
		tex := simra.NewImageTexture(texName, image.Rect(0, 0, 384, 384))
		u.sprite.ReplaceTexture2(tex)

		u.action = newAction(actionAttack, u.target)

	case commandAttackEnd:
		d := c.data.(uniter)
		if u.GetID() != d.GetID() {
			// this is not for me. ignore
			break
		}

		// TODO: load in advance. don't do every time.
		texName := fmt.Sprintf("%s.png", u.GetUnitType())
		tex := simra.NewImageTexture(texName, image.Rect(0, 0, 384, 384))
		u.sprite.ReplaceTexture2(tex)

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
			u.game.eventqueue <- newCommand(commandAttack, u)
		}

	case actionDead:
		// i'm dead!
		killUnit(u, u.game.players)

	default:
		u.unitBase.doAction(a)
	}
}
