package scene

import (
	"image/color"
	"strconv"

	"golang.org/x/mobile/asset"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
	"github.com/pankona/gomo-simra/simra/simlog"
	"github.com/pankona/phantomize/scene/config"
)

const (
	framePerSec = 60
)

// game represents a scene object for game
type game struct {
	simra            simra.Simraer
	currentStage     int
	nextScene        simra.Driver
	field            simra.Spriter
	ctrlPanel        simra.Spriter
	ctrlButton       []simra.Spriter
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
	charainfo        *charainfo
	summonPipeline   int
	ongoingSummon    int
	playerID         int
	bgm              simra.Audioer
	sound            *sound
	instruction      *instruction
}

type gameState int

const (
	gameStateInitial gameState = iota
	gameStateRunning
	gameStateGameOver
)

// Initialize initializes game scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (g *game) Initialize(sim simra.Simraer) {
	g.simra = sim
	g.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)
	g.initialize()
}

func (g *game) updateGameState(newState gameState) {
	g.gameState = newState
	switch newState {
	case gameStateInitial:
		g.currentRunLoop = g.initialRunLoop
	case gameStateRunning:
		g.currentRunLoop = g.runningRunLoop
	case gameStateGameOver:
		g.currentRunLoop = g.gameoverRunLoop
	default:
		panic("unexpected game state")
	}

	// reset frame
	g.currentFrame = 0
}

func (g *game) initField() {
	g.field = g.simra.NewSprite()
	g.field.SetScale(config.ScreenWidth, config.ScreenHeight)
	g.field.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	g.simra.AddSprite(g.field)
	tex := g.simra.NewImageTexture("field1.png", image.Rect(0, 0, 1280, 720))
	g.field.ReplaceTexture(tex)
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

func (g *game) unitIDBySprite(s simra.Spriter) string {
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

	if unitID != "" {
		// TODO: every ally spawning occurs file I/O. lol
		// loading texture in advance is needed.
		if f.game.summonPipeline < f.game.ongoingSummon+1 {
			// too much to summon. abort summoning.
			f.game.eventqueue <- newCommand(commandShowMessage, "Another summon is ongoing. please wait.")
			return
		}
		u := getUnitByUnitType(f.game.simra, unitID)
		if u.GetCost() > f.game.resource.balance {
			// balance is not enough. abort spawning
			f.game.eventqueue <- newCommand(commandShowMessage, "Need more money!")
			return
		}

		id := strconv.Itoa(f.game.playerID)
		p := newUnit(id, unitID, f.game)
		f.game.ongoingSummon++
		f.game.playerID++

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
	unittype        string
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
	simlog.Errorf("%s is unknown unittype!", unittype)
	panic("unknown unittype!")
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
	e := &effect{simra: g.simra, game: g}
	e.initialize()
	g.pubsub.Subscribe("effect", e)
}

func (g *game) initialize() {
	g.pubsub = simra.NewPubSub()
	g.eventqueue = make(chan *command, 2048)
	g.initEffects()
	g.initField()
	g.initCtrlPanel()
	g.initPlayer()
	g.initUnits("") // TODO: input JSON string
	g.resource = &resource{
		simra:   g.simra,
		balance: 100,
		game:    g,
	}
	g.resource.initialize()
	g.message = &message{simra: g.simra, game: g}
	g.selection = &selection{simra: g.simra}
	g.selection.initialize(g)
	g.charainfo = &charainfo{simra: g.simra, game: g}
	g.charainfo.initialize()
	g.sound = &sound{game: g}
	g.instruction = &instruction{simra: g.simra, game: g}
	g.instruction.initialize()
	g.simra.AddTouchListener(g)
	g.pubsub.Subscribe("god", g)
	g.pubsub.Subscribe("selection", g.selection)
	g.pubsub.Subscribe("resource", g.resource)
	g.pubsub.Subscribe("message", g.message)
	g.pubsub.Subscribe("charainfo", g.charainfo)
	g.pubsub.Subscribe("sound", g.sound)
	g.pubsub.Subscribe("instruction", g.instruction)
	g.summonPipeline = 2
	g.updateGameState(gameStateInitial)

	g.bgm = simra.NewAudio()
	resource, err := asset.Open("bgm2.mp3")
	if err != nil {
		panic(err.Error())
	}
	g.bgm.Play(resource, true, func(err error) {})

	g.eventqueue <- newCommand(commandGameStarted, nil)
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

	if qlen > 1500 {
		panic("too many events enqueued!")
	}

	c := make([]*command, qlen)
	for i := 0; i < qlen; i++ {
		c[i] = <-g.eventqueue
	}
	return c
}

type gameoverTouchListener struct {
	game *game
}

// OnTouchBegin is called when game scene is Touched.
func (c *gameoverTouchListener) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when game scene is Touched and moved.
func (c *gameoverTouchListener) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when game scene is Touched and it is released.
func (c *gameoverTouchListener) OnTouchEnd(x, y float32) {
	c.game.nextScene = &result{}
}

func (g *game) showCongratulation() {
	sprite := g.simra.NewSprite()
	sprite.SetScale(config.ScreenWidth, 80)
	sprite.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	g.simra.AddSprite(sprite)
	tex := g.simra.NewTextTexture("You won! Congratulations!",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, float32(sprite.GetScale().W), float32(sprite.GetScale().H)))
	sprite.ReplaceTexture(tex)

	sprite.AddTouchListener(&gameoverTouchListener{game: g})
}

func (g *game) showLose() {
	sprite := g.simra.NewSprite()
	sprite.SetScale(config.ScreenWidth, 80)
	sprite.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	g.simra.AddSprite(sprite)
	tex := g.simra.NewTextTexture("You lose...",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, float32(sprite.GetScale().W), float32(sprite.GetScale().H)))
	sprite.ReplaceTexture(tex)

	sprite.AddTouchListener(&gameoverTouchListener{game: g})
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
	case commandWin:
		g.updateGameState(gameStateGameOver)
		g.showCongratulation()
	case commandLose:
		g.updateGameState(gameStateGameOver)
		g.showLose()
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

	if g.areAllEnemiesEliminated() {
		g.eventqueue <- newCommand(commandWin, g)
	}
	if g.areAllPlayersEliminated() {
		g.eventqueue <- newCommand(commandLose, g)
	}
}

func (g *game) gameoverRunLoop() {
	commands := g.eventFetch()
	for _, v := range commands {
		g.pubsub.Publish(v)
	}
}

func (g *game) areAllEnemiesEliminated() bool {
	if len(g.unitPopTimeTable) != 0 {
		// there's still enemies that are waiting or spawning
		return false
	}
	if len(g.uniters) != 0 {
		// there's still enemies that are on field
		return false
	}
	return true
}

func (g *game) areAllPlayersEliminated() bool {
	return len(g.players) == 0
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (g *game) Drive() {
	defer func() {
		g.currentFrame++
	}()

	if g.nextScene != nil {
		g.bgm.Stop()
		g.simra.SetScene(g.nextScene)
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
}
