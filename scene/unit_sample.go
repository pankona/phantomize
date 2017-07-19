package scene

import (
	"image"
	"math"

	"github.com/pankona/gomo-simra/simra"
)

type sampleUnit struct {
	*unitBase
	hp         int
	attackinfo *attackInfo
	target     uniter
	isSpawned  bool
}

func (u *sampleUnit) IsSpawned() bool {
	return u.isSpawned
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
	case commandSpawn:
		d, ok := c.data.(uniter)
		if !ok {
			// unhandled event. ignore.
			return
		}
		if u.id == d.GetID() {
			// my spawn.
			u.action = newAction(actionSpawn, d)
			break
		} else if u.id != d.GetID() {
			// this spawn event is not for me.
			_, ok := d.(*player)
			if ok {
				if u.isSpawned {
					// player's spawn. move to defeat.
					u.action = newAction(actionMoveToNearestTarget, nil)
				}
			}
			return
		}

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
		if len(u.game.players) == 0 {
			// all players are eliminated
			simra.LogDebug("we won!")
			u.action = nil
			break
		}

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
		u.isSpawned = true

		// start moving to target
		u.action = newAction(actionMoveToNearestTarget, nil)

	case actionMoveToNearestTarget:
		u.target = u.nearestPlayer(u.game.players)
		if u.target == nil {
			break
		}
		u.moveToTarget(u.target)

		if u.canAttackToTarget(u.target) {
			u.action = newAction(actionAttack, u.target)
		}

	case actionAttack:
		// TODO: start animation

		target := a.data.(uniter)
		if !u.canAttackToTarget(target) {
			u.action = newAction(actionMoveToNearestTarget, nil)
			break
		}

		if u.game.currentFrame-u.attackinfo.lastAttackTime >=
			(int64)(u.attackinfo.cooltime*fps) {
			simra.LogDebug("[ATTACK] i'm %s", u.GetID())
			u.attackinfo.lastAttackTime = u.game.currentFrame

			u.game.eventqueue <- newCommand(commandDamage, &damage{target, u.attackinfo.power})
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
		// nop
	}
}

func (u *sampleUnit) nearestPlayer(players map[string]uniter) uniter {
	var (
		distance float64
		retID    string
	)
	for i, v := range players {
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
	return players[retID]
}

func (u *unitBase) moveToTarget(target uniter) {
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

func getDistanceBetweenUnit(u1, u2 uniter) float64 {
	ax, ay := u1.GetPosition()
	bx, by := u2.GetPosition()
	return getDistance(ax, ay, bx, by)
}
