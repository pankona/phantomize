package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// game represents a scene object for game
type game struct {
	currentStage     int
	nextScene        simra.Driver
	field            simra.Sprite
	ctrlPanel        simra.Sprite
	player           simra.Sprite
	currentFrame     int64
	uniters          map[string]Uniter
	unitPopTimeTable unitPopTimeTable
	pubsub           *simra.PubSub
}

// Initialize initializes game scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (game *game) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	game.initialize()

	simra.LogDebug("[OUT]")
}

func (game *game) initCtrlPanel() {
	game.ctrlPanel.W = config.ScreenWidth
	game.ctrlPanel.H = 180
	game.ctrlPanel.X = config.ScreenWidth / 2
	game.ctrlPanel.Y = game.ctrlPanel.H / 2
	simra.GetInstance().AddSprite("ctrl_panel.png",
		image.Rect(0, 0, 1280, 240),
		&game.ctrlPanel)
}

func (game *game) initField() {
	game.field.W = config.ScreenWidth
	game.field.H = config.ScreenHeight
	game.field.X = config.ScreenWidth / 2
	game.field.Y = config.ScreenHeight / 2
	simra.GetInstance().AddSprite("field1.png",
		image.Rect(0, 0, 1280, 720),
		&game.field)
}

func (game *game) initPlayer() {
	simra.GetInstance().AddSprite("player.png",
		image.Rect(0, 0, 384, 384),
		&game.player)
}

type unitPopTime struct {
	unitID          string
	popTime         int64
	initialPosition position
}

type unitPopTimeTable []*unitPopTime

const (
	fps = 60
)

func (game *game) initUnits(json string) {
	// TODO: implement
	units := make(map[string]Uniter)
	units["unit1"] = NewUnit("unit1", "")
	units["unit2"] = NewUnit("unit2", "")
	units["unit3"] = NewUnit("unit3", "")

	// TODO: unitpopTimeTable should be sorted by popTime
	game.unitPopTimeTable = append(game.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit1",
			popTime:         3 * fps,
			initialPosition: position{50, 50},
		})
	game.unitPopTimeTable = append(game.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit2",
			popTime:         5 * fps,
			initialPosition: position{150, 50},
		})
	game.unitPopTimeTable = append(game.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit3",
			popTime:         7 * fps,
			initialPosition: position{250, 50},
		})

	game.uniters = units
}

func (game *game) summonPlayer(x, y float32) {
	game.player.W = 64
	game.player.H = 64
	game.player.X = x
	game.player.Y = y
}

func (game *game) popUnits() []Uniter {
	poppedUnits := make([]Uniter, 0)
	for _, v := range game.unitPopTimeTable {
		if v.popTime <= game.currentFrame {
			// pop unit
			u := game.uniters[v.unitID]
			u.SetPosition(v.initialPosition)
			poppedUnits = append(poppedUnits, u)
			continue
		}

	}

	if len(poppedUnits) != 0 {
		// remove popped units from unitPopTimeTable
		game.unitPopTimeTable = game.unitPopTimeTable[len(poppedUnits):]
	}

	return poppedUnits
}

func (game *game) initialize() {
	game.initField()
	game.initCtrlPanel()
	game.initPlayer()
	game.initUnits("") // TODO: input JSON string
	game.pubsub = simra.NewPubSub()
	simra.GetInstance().AddTouchListener(game)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (game *game) Drive() {
	defer func() {
		game.currentFrame++
	}()

	if game.nextScene != nil {
		simra.GetInstance().SetScene(game.nextScene)
	}

	poppedUnits := game.popUnits()
	for _, v := range poppedUnits {
		err := game.pubsub.Subscribe(v.GetID(), v)
		if err != nil {
			panic("failed to subscribe. fatal.")
		}

		// generate unit pop command
		c := newCommand()
		c.commandtype = SPAWN
		c.data = v

		// publish unit pop command
		game.pubsub.Publish(c)
	}

	// pop all events from event queue
	// this event should be done within 1 frame

	// generate command by each event

	// broadcast event to all units
}

type eventer interface {
	// TODO: implement
}

func (game *game) eventCallback(event eventer) {
}

// OnTouchBegin is called when game scene is Touched.
func (game *game) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when game scene is Touched and moved.
func (game *game) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when game scene is Touched and it is released.
func (game *game) OnTouchEnd(x, y float32) {
	if y > 180 {
		game.summonPlayer(x, y)
	}
	//game.nextScene = &result{currentStage: game.currentStage}
}
