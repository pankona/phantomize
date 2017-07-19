package scene

import (
	"image"
	"math"

	"github.com/pankona/gomo-simra/simra"
)

type sampleUnit struct {
	*unitBase
}

func (u *sampleUnit) Initialize() {
	simra.GetInstance().AddSprite("player.png",
		image.Rect(0, 0, 384, 384),
		&u.sprite)
	u.hp = 50
}

func (u *sampleUnit) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandDead:
		if len(u.game.players) == 0 {
			// all players are eliminated
			simra.LogDebug("we won!")
			u.action = nil
			break
		}

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
		delete(u.game.uniters, u.GetID())

	default:
		u.unitBase.doAction(a)
	}
}

// TODO: move to utility
func moveToTarget(u *unitBase, target uniter) {
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

// TODO: move to utility
func canAttackToTarget(u *unitBase, target uniter) bool {
	ux, uy := u.GetPosition()
	tx, ty := target.GetPosition()

	if (float64)(u.attackinfo.attackRange) >= getDistance(ux, uy, tx, ty) {
		return true
	}
	return false
}

// TODO: move to utility
func getDistance(ax, ay, bx, by float32) float64 {
	dx, dy := ax-bx, ay-by
	return math.Sqrt((float64)(dx*dx + dy*dy))
}

// TODO: move to utility
func getDistanceBetweenUnit(u1, u2 uniter) float64 {
	ax, ay := u1.GetPosition()
	bx, by := u2.GetPosition()
	return getDistance(ax, ay, bx, by)
}

// TODO: move to utility
func nearestUnit(u *unitBase, enemies map[string]uniter) uniter {
	var (
		distance float64
		retID    string
	)
	for i, v := range enemies {
		if !v.IsSpawned() {
			continue
		}
		d := getDistanceBetweenUnit(u, v)
		if distance == 0 {
			distance = d
			retID = i
			continue
		}
		if distance > d {
			distance = d
			retID = i
		}
	}

	if retID == "" {
		return nil
	}
	return enemies[retID]
}
