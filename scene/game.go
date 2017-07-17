package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

const (
	ctrlPanelHeight = 220
	fps             = 60
)

// game represents a scene object for game
type game struct {
	currentStage     int
	nextScene        simra.Driver
	field            simra.Sprite
	ctrlPanel        simra.Sprite
	ctrlButton       []simra.Sprite
	currentFrame     int64
	players          map[string]uniter
	uniters          map[string]uniter
	unitPopTimeTable unitPopTimeTable
	pubsub           *simra.PubSub
	gameState        gameState
	currentRunLoop   func()
	eventqueue       chan *command
	selection        *selection
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
func (g *game) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	g.initialize()

	simra.LogDebug("[OUT]")
}

func (g *game) updateGameState(newState gameState) {
	g.gameState = newState
	switch newState {
	case gameStateInitial:
		g.currentRunLoop = g.initialRunLoop
	case gameStateRunning:
		g.currentRunLoop = g.runningRunLoop
	default:
		//nop
	}

	// reset frame
	g.currentFrame = 0
}

type selection struct {
	selecting *simra.Sprite
	class     string
}

func (s *selection) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandUpdateSelection:
		s.selecting = c.data.(*simra.Sprite)
		simra.LogDebug("selection updated: %v", s.selecting)
	case commandUnsetSelection:
		s.selecting = nil
		simra.LogDebug("selection updated: nil")
	default:
		// nop
	}
}

type ctrlButtonTouchListener struct {
	id   int
	game *game
}

func (c *ctrlButtonTouchListener) OnTouchBegin(x, y float32) {
	// nop
}

func (c *ctrlButtonTouchListener) OnTouchMove(x, y float32) {
	// nop
}

func (c *ctrlButtonTouchListener) OnTouchEnd(x, y float32) {
	c.game.eventqueue <- newCommand(commandUpdateSelection, &c.game.ctrlButton[c.id])
}

func (g *game) initCtrlPanel() {
	g.ctrlPanel.W = config.ScreenWidth
	g.ctrlPanel.H = ctrlPanelHeight
	g.ctrlPanel.X = config.ScreenWidth / 2
	g.ctrlPanel.Y = g.ctrlPanel.H / 2
	simra.GetInstance().AddSprite("ctrl_panel.png",
		image.Rect(0, 0, 1280, 240),
		&g.ctrlPanel)

	g.ctrlButton = make([]simra.Sprite, 9)
	for i := range g.ctrlButton {
		g.ctrlButton[i].W = 64
		g.ctrlButton[i].H = 64
		g.ctrlButton[i].X = (float32)(1010 + (64+45)*(i%3))
		g.ctrlButton[i].Y = (float32)(44 + (64+5)*(i/3))
		simra.GetInstance().AddSprite("player.png",
			image.Rect(0, 0, 384, 384),
			&g.ctrlButton[i])

		g.ctrlButton[i].AddTouchListener(&ctrlButtonTouchListener{id: i, game: g})
	}
}

func (g *game) initField() {
	g.field.W = config.ScreenWidth
	g.field.H = config.ScreenHeight
	g.field.X = config.ScreenWidth / 2
	g.field.Y = config.ScreenHeight / 2
	simra.GetInstance().AddSprite("field1.png",
		image.Rect(0, 0, 1280, 720),
		&g.field)
	g.field.AddTouchListener(&fieldTouchListener{game: g})
}

type fieldTouchListener struct {
	game *game
}

// OnTouchBegin is called when game scene is Touched.
func (f *fieldTouchListener) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when game scene is Touched and moved.
func (f *fieldTouchListener) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when game scene is Touched and it is released.
func (f *fieldTouchListener) OnTouchEnd(x, y float32) {
	if y <= ctrlPanelHeight {
		return
	}

	if f.game.selection.selecting == nil {
		return
	}

	s := f.game.selection.selecting
	var id string
	if s == &f.game.ctrlButton[6] {
		id = "player1"
	} else if s == &f.game.ctrlButton[7] {
		id = "player2"
	} else if s == &f.game.ctrlButton[8] {
		id = "player3"
	} else if s == &f.game.ctrlButton[3] {
		id = "player4"
	} else if s == &f.game.ctrlButton[4] {
		id = "player5"
	} else if s == &f.game.ctrlButton[5] {
		id = "player6"
	} else if s == &f.game.ctrlButton[0] {
		id = "player7"
	} else if s == &f.game.ctrlButton[1] {
		id = "player8"
	} else if s == &f.game.ctrlButton[2] {
		id = "player9"
	}

	if id != "" {
		p := newUnit(id, "player", f.game)
		f.game.players[id] = p
		f.game.pubsub.Subscribe(p.GetID(), p)
	}
}

func (g *game) initPlayer() {
	g.players = make(map[string]uniter)
	p := newUnit("player", "player", g)
	g.players["player"] = p
	g.pubsub.Subscribe(p.GetID(), p)
}

type unitPopTime struct {
	unitID          string
	popTime         int64
	initialPosition position
}

type unitPopTimeTable []*unitPopTime

func (g *game) initUnits(json string) {
	// TODO: load from json file
	units := make(map[string]uniter)
	units["unit1"] = newUnit("unit1", "", g)
	units["unit2"] = newUnit("unit2", "", g)
	units["unit3"] = newUnit("unit3", "", g)

	// TODO: unitpopTimeTable should be sorted by popTime
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit1",
			popTime:         3 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5},
		})
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit2",
			popTime:         4 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4},
		})
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "unit3",
			popTime:         5 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 3},
		})

	g.uniters = units
}

func (g *game) popUnits() []uniter {
	poppedUnits := make([]uniter, 0)
	for _, v := range g.unitPopTimeTable {
		if v.popTime <= g.currentFrame {
			// pop unit
			u := g.uniters[v.unitID]
			u.SetPosition(v.initialPosition.x, v.initialPosition.y)
			poppedUnits = append(poppedUnits, u)
			continue
		}
	}

	if len(poppedUnits) != 0 {
		// remove popped units from unitPopTimeTable
		g.unitPopTimeTable = g.unitPopTimeTable[len(poppedUnits):]
	}

	return poppedUnits
}

func (g *game) initialize() {
	g.pubsub = simra.NewPubSub()
	g.eventqueue = make(chan *command, 256)
	g.initField()
	g.initCtrlPanel()
	g.initPlayer()
	g.initUnits("") // TODO: input JSON string
	g.selection = &selection{}
	simra.GetInstance().AddTouchListener(g)
	g.pubsub.Subscribe("god", g)
	g.pubsub.Subscribe("selection", g.selection)
	g.updateGameState(gameStateInitial)
}

func (g *game) eventFetch() []*command {
	// note:
	// if new events are pushed while fetching,
	// they should be fetched next run loop to
	// avoid inifinite event fetching.

	qlen := len(g.eventqueue)
	if qlen == 0 {
		return nil
	}

	c := make([]*command, qlen)
	for i := 0; i < qlen; i++ {
		c[i] = <-g.eventqueue
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

func (g *game) initialRunLoop() {
	commands := g.eventFetch()
	for _, v := range commands {
		g.pubsub.Publish(v)
	}
	g.players["player"].DoAction()
}

func (g *game) runningRunLoop() {
	poppedUnits := g.popUnits()
	for _, v := range poppedUnits {
		err := g.pubsub.Subscribe(v.GetID(), v)
		if err != nil {
			panic("failed to subscribe. fatal.")
		}

		// generate spawn command
		g.eventqueue <- newCommand(commandSpawn, v)
	}

	// event fetch and publish to all subscribers
	commands := g.eventFetch()
	for _, v := range commands {
		g.pubsub.Publish(v)
	}

	// invoke action for all units
	for _, v := range g.uniters {
		v.DoAction()
	}
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (g *game) Drive() {
	defer func() {
		g.currentFrame++
	}()

	if g.nextScene != nil {
		simra.GetInstance().SetScene(g.nextScene)
	}
	g.currentRunLoop()
}

// OnTouchBegin is called when game scene is Touched.
func (g *game) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when game scene is Touched and moved.
func (g *game) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when game scene is Touched and it is released.
func (g *game) OnTouchEnd(x, y float32) {
	if g.gameState == gameStateInitial {
		if y > ctrlPanelHeight {
			g.players["player"].SetPosition(x, y)

			c := newCommand(commandSpawn, g.players["player"])
			g.eventqueue <- c

			c = newCommand(commandGoToRunningState, g)
			g.eventqueue <- c
		}
	}
	//g.nextScene = &result{currentStage: g.currentStage}
}
