package scene

import (
	"fmt"
	"image"

	"github.com/pankona/gomo-simra/simra"
)

// unit base implementation

type uniter interface {
	Initialize()
	GetID() string
	SetPosition(float32, float32)
	GetPosition() (float32, float32)
	IsSpawned() bool
	Dead()
	DoAction()
	GetUnitType() string
	GetCost() int
	GetTarget() uniter
	IsAlly() bool
	simra.Subscriber
}

type attackInfo struct {
	attackRange    int
	power          int
	cooltime       float32 // second
	lastAttackTime int64   // frame
}

type position struct {
	x float32
	y float32
}

type actiontype int

const (
	// SPAWN spawns an unit
	actionSpawn actiontype = iota
	actionDead
	actionMoveToNearestTarget
	actionAttack
)

type action struct {
	actiontype actiontype
	data       interface{}
}

func newAction(a actiontype, d interface{}) *action {
	return &action{
		actiontype: a,
		data:       d,
	}
}

type unitBase struct {
	simra.Subscriber
	sprite              simra.Sprite
	id                  string
	unittype            string
	action              *action
	game                *game
	moveSpeed           float32
	hp                  int
	attackinfo          *attackInfo
	target              uniter
	isSpawned           bool
	delayTimeToSummon   int64
	elapsedTimeToSummon int64
	isAlly              bool
	cost                int
}

func (u *unitBase) Initialize() {}

func (u *unitBase) GetID() string {
	return u.id
}

func (u *unitBase) SetPosition(x, y float32) {
	u.sprite.X = x
	u.sprite.Y = y
}

func (u *unitBase) GetPosition() (float32, float32) {
	return u.sprite.X, u.sprite.Y
}

func (u *unitBase) IsSpawned() bool {
	return u.isSpawned
}

// note that this is NOT delegate method.
// used by an object that composites unitBase.
func (u *unitBase) onEvent(c *command) {
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
		}

	case commandSpawned:
		d, ok := c.data.(uniter)
		if !ok {
			// unhandled event. ignore.
			return
		}

		if u.id != d.GetID() {
			// this spawn event is not for me.
			_, ok := d.(*sampleUnit)
			if ok {
				if u.isSpawned {
					simra.LogDebug("enemy's spawn %s is detected! kill them all!", d.GetID())

					// enemy's spawn. move to defeat.
					u.action = newAction(actionMoveToNearestTarget, nil)
				}
			}
			return
		}
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

		fmt.Printf("[unit][%s] ends attacking\n", u.GetID())

		// TODO: load in advance. don't do every time.
		texName := fmt.Sprintf("%s.png", u.GetUnitType())
		tex := simra.NewImageTexture(texName, image.Rect(0, 0, 384, 384))
		u.sprite.ReplaceTexture2(tex)

	case commandDamage:
		d, ok := c.data.(*damage)
		if !ok {
			return
		}
		if u.id != d.unit.GetID() {
			return
		}

		u.hp -= d.damage
		simra.LogDebug("[DAMAGE] i'm [%s], HP = %d", u.GetID(), u.hp)
		if u.hp <= 0 {
			simra.LogDebug("[DEAD] i'm %s", u.GetID())
			u.game.eventqueue <- newCommand(commandDead, u)
		}

	case commandDead:
		d := c.data.(uniter)
		if u.GetID() == d.GetID() {
			u.action = newAction(actionDead, nil)
		}
		if u.target == d.GetTarget() {
			fmt.Printf("target [%s] is down. [%s] stop attacking.\n", d.GetTarget().GetID(), u.GetID())
			u.game.eventqueue <- newCommand(commandAttackEnd, u)
		}
		if len(u.game.uniters) == 0 {
			u.action = nil
			break
		}

	default:
		// nop
	}
}

func (u *unitBase) DoAction() {}

func (u *unitBase) GetUnitType() string {
	return u.unittype
}

func (u *unitBase) GetTarget() uniter {
	return u.target
}

func (u *unitBase) GetCost() int {
	return u.cost
}

func (u *unitBase) IsAlly() bool {
	return u.isAlly
}

func (u *unitBase) doAction(a *action) {
	switch a.actiontype {
	case actionSpawn:
		u.elapsedTimeToSummon++
		if u.elapsedTimeToSummon <= u.delayTimeToSummon {
			// still summoning...
			break
		}
		u.elapsedTimeToSummon = 0

		d := a.data.(uniter)
		u.sprite.W = 64
		u.sprite.H = 64
		u.SetPosition(d.GetPosition())
		simra.LogDebug("@@@@@@ [SPAWN] i'm %s", u.GetID())
		u.isSpawned = true

		// start moving to target
		u.game.eventqueue <- newCommand(commandSpawned, u)
		u.action = newAction(actionMoveToNearestTarget, nil)

	case actionAttack:
		target := a.data.(uniter)
		if !canAttackToTarget(u, target) {
			u.game.eventqueue <- newCommand(commandAttackEnd, u)
			u.action = newAction(actionMoveToNearestTarget, nil)
			break
		}

		if u.game.currentFrame-u.attackinfo.lastAttackTime >=
			(int64)(u.attackinfo.cooltime*fps) {
			simra.LogDebug("[ATTACK] i'm %s", u.GetID())
			u.attackinfo.lastAttackTime = u.game.currentFrame

			u.game.eventqueue <- newCommand(commandDamage, &damage{target, u.attackinfo.power})
		}

	default:
		// nop
	}
}

func (u *unitBase) Dead() {
	u.sprite.W = 1
	u.sprite.H = 1
	u.SetPosition(-1, -1)
	u.action = nil
	u.target = nil
	u.isSpawned = false
}

func newUnit(id, unittype string, game *game) uniter {
	// TODO: sample unit implemenation
	// unit type should be specified and switch here
	var u uniter
	switch unittype {
	case "player1":
		u = &player{
			unitBase: &unitBase{
				id:        id,
				unittype:  unittype,
				game:      game,
				moveSpeed: 1.5,
				hp:        50,
				attackinfo: &attackInfo{
					attackRange: 50,
					power:       15,
					cooltime:    2,
				},
				delayTimeToSummon: 5 * fps,
				isAlly:            true,
				cost:              10,
			},
		}
	case "player2":
		u = &player{
			unitBase: &unitBase{
				id:        id,
				unittype:  unittype,
				game:      game,
				moveSpeed: 1.0,
				hp:        75,
				attackinfo: &attackInfo{
					attackRange: 50,
					power:       20,
					cooltime:    3,
				},
				delayTimeToSummon: 5 * fps,
				isAlly:            true,
				cost:              20,
			},
		}

	case "player3":
		u = &player{
			unitBase: &unitBase{
				id:        id,
				unittype:  unittype,
				game:      game,
				moveSpeed: 0.5,
				hp:        30,
				attackinfo: &attackInfo{
					attackRange: 200,
					power:       20,
					cooltime:    3,
				},
				delayTimeToSummon: 5 * fps,
				isAlly:            true,
				cost:              25,
			},
		}

	case "enemy1":
		u = &sampleUnit{
			unitBase: &unitBase{
				id:        id,
				unittype:  unittype,
				game:      game,
				moveSpeed: 0.5,
				attackinfo: &attackInfo{
					attackRange: 50,
					power:       15,
					cooltime:    2,
				},
				isAlly: false,
			},
		}

	case "enemy2":
		u = &sampleUnit{
			unitBase: &unitBase{
				id:        id,
				unittype:  unittype,
				game:      game,
				moveSpeed: 0.5,
				attackinfo: &attackInfo{
					attackRange: 50,
					power:       15,
					cooltime:    2,
				},
				isAlly: false,
			},
		}
	}

	// call each unit's initialize function
	u.Initialize()
	return u
}

type commandtype int

const (
	commandSpawn commandtype = iota
	commandSpawned
	commandAttack
	commandAttackEnd
	commandDamage
	commandDead
	commandGoToInitialState
	commandGoToRunningState
	commandUpdateSelection
	commandUnsetSelection
)

type command struct {
	commandtype commandtype
	data        interface{}
}

type damage struct {
	unit   uniter
	damage int
}

func newCommand(c commandtype, d interface{}) *command {
	return &command{commandtype: c, data: d}
}

func killUnit(u uniter, umap map[string]uniter) {
	u.Dead()
	delete(umap, u.GetID())
}
