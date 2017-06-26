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
	player           Uniter
	currentFrame     int64
	uniters          map[string]Uniter
	unitPopTimeTable unitPopTimeTable
	pubsub           *simra.PubSub
	gameState        gameState
	currentRunLoop   func()
	eventqueue       chan *command
}

type gameState int

const (
	gameStateInitial gameState = iota
	gameStateRunning
)

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

func (game *game) updateGameState(newState gameState) {
	game.gameState = newState
	switch newState {
	case gameStateInitial:
		game.currentRunLoop = game.initialRunLoop
	case gameStateRunning:
		game.currentRunLoop = game.runningRunLoop
	default:
		//nop
	}

	// reset frame
	game.currentFrame = 0
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
	p := NewUnit("player", "player", game)
	game.player = p
	game.pubsub.Subscribe(p.GetID(), p)
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
	units["unit1"] = NewUnit("unit1", "", game)
	units["unit2"] = NewUnit("unit2", "", game)
	units["unit3"] = NewUnit("unit3", "", game)

	// TODO: unitpopTimeTable should be sorted by popTime
	game.unitPopTimeTable = append(game.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit1",
			popTime:         3 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5},
		})
	game.unitPopTimeTable = append(game.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit2",
			popTime:         3 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4},
		})
	game.unitPopTimeTable = append(game.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit3",
			popTime:         3 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3},
		})

	game.uniters = units
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
	game.pubsub = simra.NewPubSub()
	game.eventqueue = make(chan *command, 256)
	game.initField()
	game.initCtrlPanel()
	game.initPlayer()
	game.initUnits("") // TODO: input JSON string
	simra.GetInstance().AddTouchListener(game)
	game.pubsub.Subscribe("god", game)
	game.updateGameState(gameStateInitial)
}

func (game *game) eventFetch() []*command {
	// note:
	// if new events are pushed while fetching,
	// they should be fetched next run loop to
	// avoid inifinite event fetching.

	qlen := len(game.eventqueue)
	if qlen == 0 {
		return nil
	}

	c := make([]*command, qlen)
	for i := 0; i < qlen; i++ {
		c[i] = <-game.eventqueue
	}
	return c
}

func (g *game) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		// should be a command. ignore.
		return
	}

	_, ok = c.data.(*game)
	if !ok {
		// this is not for me
		return
	}

	switch c.commandtype {
	case commandGoToInitialState:
		g.updateGameState(gameStateInitial)
	case commandGoToRunningState:
		g.updateGameState(gameStateRunning)
	}
}

func (game *game) initialRunLoop() {
	commands := game.eventFetch()
	for _, v := range commands {
		game.pubsub.Publish(v)
	}
	game.player.DoAction()
}

func (game *game) runningRunLoop() {
	poppedUnits := game.popUnits()
	for _, v := range poppedUnits {
		err := game.pubsub.Subscribe(v.GetID(), v)
		if err != nil {
			panic("failed to subscribe. fatal.")
		}

		// generate spawn command
		c := newCommand(commandSpawn, v)
		c.data = v

		game.eventqueue <- c
	}

	// event fetch and publish to all subscribers
	commands := game.eventFetch()
	for _, v := range commands {
		game.pubsub.Publish(v)
	}

	// invoke action for all units
	for _, v := range game.uniters {
		v.DoAction()
	}
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
	game.currentRunLoop()
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
	if game.gameState == gameStateInitial {
		if y > 180 {
			game.player.SetPosition(position{(int)(x), (int)(y)})

			c := newCommand(commandSpawn, game.player)
			game.eventqueue <- c

			c = newCommand(commandGoToRunningState, game)
			game.eventqueue <- c
		}
	}
	//game.nextScene = &result{currentStage: game.currentStage}
}
