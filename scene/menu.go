package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// menu represents a scene object for menu
type menu struct {
	text simra.Sprite
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

func (menu *menu) initialize() {
	menu.text.W = config.ScreenWidth
	menu.text.H = 80
	menu.text.X = config.ScreenWidth / 2
	menu.text.Y = config.ScreenHeight / 2
	simra.GetInstance().AddTextSprite("menu",
		60, // fontsize
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(menu.text.W), int(menu.text.H)),
		&menu.text)

	simra.GetInstance().AddTouchListener(menu)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (menu *menu) Drive() {
}

// OnTouchBegin is called when menu scene is Touched.
func (menu *menu) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called when menu scene is Touched and moved.
func (menu *menu) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called when menu scene is Touched and it is released.
func (menu *menu) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	simra.GetInstance().SetScene(&Stage1{})
}
