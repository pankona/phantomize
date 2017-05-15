package scene

import (
	"image"
	"image/color"
	"strconv"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// game represents a scene object for game
type game struct {
	text         simra.Sprite
	currentStage int
	nextScene    simra.Driver
	ctrlPanel    simra.Sprite
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

func (game *game) initTempText() {
	// temporary text (will be removed)
	temporary := &simra.Sprite{}
	temporary.W = config.ScreenWidth
	temporary.H = 80
	temporary.X = config.ScreenWidth / 2
	temporary.Y = config.ScreenHeight * 2 / 5
	simra.GetInstance().AddTextSprite("(click to go to result scene)",
		60, // fontsize
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(temporary.W), int(temporary.H)),
		temporary)
}

func (game *game) initCtrlPanel() {
	game.ctrlPanel.W = config.ScreenWidth
	game.ctrlPanel.H = 240
	game.ctrlPanel.X = config.ScreenWidth / 2
	game.ctrlPanel.Y = 120
	simra.GetInstance().AddSprite("ctrl_panel.png",
		image.Rect(0, 0, 1280, 240),
		&game.ctrlPanel)
}

func (game *game) initialize() {
	initTextSprite(&game.text, "game for stage "+strconv.Itoa(game.currentStage),
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*4/6,
		60, color.RGBA{255, 0, 0, 255})
	// temporary text (will be removed)
	game.initTempText()
	game.initCtrlPanel()
	simra.GetInstance().AddTouchListener(game)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (game *game) Drive() {
	// nop
	if game.nextScene != nil {
		simra.GetInstance().SetScene(game.nextScene)
	}
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
	game.nextScene = &result{currentStage: game.currentStage}
}
