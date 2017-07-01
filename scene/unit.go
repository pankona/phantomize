package scene

import (
	"github.com/pankona/gomo-simra/simra"
)

// unit base implementation

type uniter interface {
	Initialize()
	GetID() string
	simra.Subscriber
	SetPosition(float32, float32)
	GetPosition() (float32, float32)
	DoAction()
}

type attackInfo struct {
	attackRange int
	power       int
	speed       int
}

type position struct {
	x float32
	y float32
}

type actiontype int

const (
	// SPAWN spawns an unit
	actionSpawn actiontype = iota
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
	sprite    simra.Sprite
	id        string
	action    *action
	game      *game
	moveSpeed float32
}

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

// NewUnit returns a uniter
func NewUnit(id, unittype string, game *game) uniter {
	// TODO: sample unit implemenation
	// unit type should be specified and switch here
	var u uniter
	switch unittype {
	case "player":
		u = &player{
			unitBase: &unitBase{id: id, game: game, moveSpeed: 0},
		}
	default:
		// TODO: remove later
		u = &sampleUnit{
			unitBase:   &unitBase{id: id, game: game, moveSpeed: 0.5},
			attackinfo: &attackInfo{attackRange: 5, power: 5, speed: 10},
		}
	}

	// call each unit's initialize function
	u.Initialize()
	return u
}

type commandtype int

const (
	// SPAWN spawns an unit
	commandSpawn commandtype = iota
	commandGoToInitialState
	commandGoToRunningState
)

type command struct {
	commandtype commandtype
	data        interface{}
}

func newCommand(c commandtype, d interface{}) *command {
	return &command{commandtype: c, data: d}
}
