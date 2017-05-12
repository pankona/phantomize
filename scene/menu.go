package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// menu represents a scene object for menu
type menu struct {
	menu      simra.Sprite
	start     simra.Sprite
	howto     simra.Sprite
	nextScene simra.Driver
}

// Initialize initializes menu scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (menu *menu) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	menu.initialize()

	simra.LogDebug("[OUT]")
}

func initTextSprite(sprite *simra.Sprite, text string, w, h, x, y float32, fontsize float64, color color.RGBA) {
	sprite.W, sprite.H, sprite.X, sprite.Y = w, h, x, y
	simra.GetInstance().AddTextSprite(text,
		fontsize, // fontsize
		color,
		image.Rect(0, 0, int(sprite.W), int(sprite.H)),
		sprite)
}

func (menu *menu) initialize() {
	initTextSprite(&menu.menu, "menu",
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*4/6,
		60, color.RGBA{255, 0, 0, 255})
	initTextSprite(&menu.start, "Start",
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*2/6,
		60, color.RGBA{255, 0, 0, 255})
	initTextSprite(&menu.howto, "How to play",
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*1/6,
		60, color.RGBA{255, 0, 0, 255})
	//simra.GetInstance().AddTouchListener(menu)
	menu.start.AddTouchListener(&startListener{menu: menu})
	menu.howto.AddTouchListener(&howToPlayListener{menu: menu})
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (menu *menu) Drive() {
	// nop
	if menu.nextScene != nil {
		simra.GetInstance().SetScene(menu.nextScene)
	}
}

// OnTouchBegin is called when menu scene is Touched.
func (menu *menu) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when menu scene is Touched and moved.
func (menu *menu) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when menu scene is Touched and it is released.
func (menu *menu) OnTouchEnd(x, y float32) {
	// nop
}

type startListener struct {
	*menu
}
type howToPlayListener struct {
	*menu
}

func (start *startListener) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	start.menu.nextScene = &briefing{currentStage: 1}
}

func (howto *howToPlayListener) OnTouchEnd(x, y float32) {
	howto.menu.nextScene = &howtoplay{}
}
