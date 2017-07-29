package scene

import (
	"image"
	"strconv"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

const (
	fps = 60
)

// game represents a scene object for game
type game struct {
	currentStage     int
	nextScene        simra.Driver
	field            simra.Sprite
	ctrlPanel        simra.Sprite
	ctrlButton       []*simra.Sprite
	currentFrame     int64
	players          map[string]uniter
	uniters          map[string]uniter
	unitPopTimeTable unitPopTimeTable
	pubsub           *simra.PubSub
	gameState        gameState
	currentRunLoop   func()
	eventqueue       chan *command
	selection        *selection
	resource         *resource
	message          *message
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

func (g *game) unitIDBySprite(s *simra.Sprite) string {
	var unitID string
	if s == g.ctrlButton[0] {
		unitID = "player1"
	} else if s == g.ctrlButton[1] {
		unitID = "player2"
	} else if s == g.ctrlButton[2] {
		unitID = "player3"
	}
	return unitID
}

// OnTouchEnd is called when game scene is Touched and it is released.
func (f *fieldTouchListener) OnTouchEnd(x, y float32) {
	if y <= ctrlPanelHeight {
		return
	}

	if f.game.selection.selecting == nil {
		return
	}

	unitID := f.game.unitIDBySprite(f.game.selection.selecting)
	id := strconv.Itoa(len(f.game.players))

	if unitID != "" {
		// TODO: every ally spawning occurs file I/O. lol
		// loading texture in advance is needed.
		p := newUnit(id, unitID, f.game)
		if p.GetCost() > f.game.resource.balance {
			// balance is not enough. abort spawning
			f.game.eventqueue <- newCommand(commandShowMessage, "Need more money!")
			return
		}
		p.SetPosition(x, y)
		f.game.players[id] = p
		f.game.pubsub.Subscribe(p.GetID(), p)
		f.game.eventqueue <- newCommand(commandSpawn, p)
		f.game.eventqueue <- newCommand(commandUnsetSelection, nil)
	}
}

func (g *game) initPlayer() {
	g.players = make(map[string]uniter)
}

type unitPopTime struct {
	unitID          string
	popTime         int64
	initialPosition position
}

type unitPopTimeTable []*unitPopTime

func (g *game) assetNameByUnitType(unittype string) string {
	switch unittype {
	case "player1":
		return "player1.png"
	case "player2":
		return "player2.png"
	case "player3":
		return "player3.png"
	case "enemy1":
		return "enemy1.png"
	case "enemy2":
		return "enemy2.png"
	}

	simra.LogError("%s is unknown unittype!", unittype)
	panic("unknown unittype!")
}

func (g *game) initUnits(json string) {
	// TODO: load from json file
	units := make(map[string]uniter)
	units["e1"] = newUnit("e1", "enemy1", g)
	units["e2"] = newUnit("e2", "enemy1", g)
	units["e3"] = newUnit("e3", "enemy1", g)

	// TODO: unitpopTimeTable should be sorted by popTime
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "e1",
			popTime:         3 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 5},
		})
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "e2",
			popTime:         4 * fps,
			initialPosition: position{config.ScreenWidth - 32, config.ScreenHeight / 6 * 4},
		})
	g.unitPopTimeTable = append(g.unitPopTimeTable,
		&unitPopTime{
			unitID:          "e3",
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

func (g *game) initEffects() {
	e := &effect{game: g}
	e.initialize()
	g.pubsub.Subscribe("effect", e)
}

func (g *game) initialize() {
	g.pubsub = simra.NewPubSub()
	g.eventqueue = make(chan *command, 256)
	g.initEffects()
	g.initField()
	g.initCtrlPanel()
	g.initPlayer()
	g.initUnits("") // TODO: input JSON string
	g.resource = &resource{
		balance: 100,
		game:    g,
	}
	g.resource.initialize()
	g.message = &message{game: g}
	g.selection = &selection{}
	g.selection.initialize(g)
	simra.GetInstance().AddTouchListener(g)
	g.pubsub.Subscribe("god", g)
	g.pubsub.Subscribe("selection", g.selection)
	g.pubsub.Subscribe("resource", g.resource)
	g.pubsub.Subscribe("message", g.message)
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
	for _, v := range g.players {
		v.DoAction()
	}
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
	for _, v := range g.players {
		v.DoAction()
	}
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
	// nop
	//g.nextScene = &result{currentStage: g.currentStage}
}
