package scene

import (
	"fmt"
	"image"
)

type sampleUnit struct {
	*unitBase
}

func (u *sampleUnit) Initialize() {
	assetName := u.game.assetNameByUnitType(u.unittype)
	u.simra.AddSprite(u.sprite)
	tex := u.simra.NewImageTexture(assetName,
		image.Rect(0, 0, 384, 384))
	u.sprite.ReplaceTexture(tex)

}

func (u *sampleUnit) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	default:
		u.unitBase.onEvent(c)
	}
}

func (u *sampleUnit) DoAction() {
	a := u.action
	if a == nil {
		// idle
		return
	}

	switch a.actiontype {
	case actionMoveToNearestTarget:
		u.target = nearestUnit(u.unitBase, u.game.players)
		if u.target == nil {
			break
		}
		moveToTarget(u.unitBase, u.target)

		if canAttackToTarget(u.unitBase, u.target) {
			u.game.eventqueue <- newCommand(commandAttack, u)
		}

	case actionDead:
		// i'm dead!
		fmt.Println("@@@@@@@ dead!!")
		u.game.pubsub.Unsubscribe(u.GetID())
		killUnit(u, u.game.uniters)

	default:
		u.unitBase.doAction(a)
	}
}
