package scene

import (
	"image"
	"sync"

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
	runLoopMutex     sync.Mutex
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
	game.runLoopMutex.Lock()
	defer game.runLoopMutex.Unlock()

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
	p := NewUnit("player", "player")
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
	units["unit1"] = NewUnit("unit1", "")
	units["unit2"] = NewUnit("unit2", "")
	units["unit3"] = NewUnit("unit3", "")

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
	game.updateGameState(gameStateInitial)
}

func (game *game) eventFetch() {
eventFetch:
	for {
		select {
		case c := <-game.eventqueue:
			game.pubsub.Publish(c)
		default:
			break eventFetch
		}
	}
}

func (game *game) initialRunLoop() {
	game.eventFetch()
}

func (game *game) runningRunLoop() {
	poppedUnits := game.popUnits()
	for _, v := range poppedUnits {
		err := game.pubsub.Subscribe(v.GetID(), v)
		if err != nil {
			panic("failed to subscribe. fatal.")
		}

		// generate spawn command
		c := newCommand()
		c.commandtype = SPAWN
		c.data = v

		game.eventqueue <- c
	}

	// pop all events from event queue
	// this event should be done within 1 frame
	game.eventFetch()

	// generate command by each event

	// broadcast event to all units
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

	game.runLoopMutex.Lock()
	game.currentRunLoop()
	game.runLoopMutex.Unlock()
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
	if game.gameState == gameStateInitial {
		if y > 180 {
			// TODO: don't publish here. use event queue
			c := newCommand()
			c.commandtype = SPAWN
			game.player.SetPosition(position{(int)(x), (int)(y)})
			c.data = game.player

			game.eventqueue <- c
			game.updateGameState(gameStateRunning)
		}
	}
	//game.nextScene = &result{currentStage: game.currentStage}
}
