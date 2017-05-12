package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

type howtoplay struct {
	text      simra.Sprite
	nextScene simra.Driver
}

// Initialize initializes howtoplay scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (howtoplay *howtoplay) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	howtoplay.initialize()

	simra.LogDebug("[OUT]")
}

func (howtoplay *howtoplay) initialize() {
	howtoplay.text.W = config.ScreenWidth
	howtoplay.text.H = 80
	howtoplay.text.X = config.ScreenWidth / 2
	howtoplay.text.Y = config.ScreenHeight / 2
	simra.GetInstance().AddTextSprite("How To Play",
		60, // fontsize
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(howtoplay.text.W), int(howtoplay.text.H)),
		&howtoplay.text)

	simra.GetInstance().AddTouchListener(howtoplay)

	// temporary text (will be removed)
	temporary := &simra.Sprite{}
	temporary.W = config.ScreenWidth
	temporary.H = 80
	temporary.X = config.ScreenWidth / 2
	temporary.Y = config.ScreenHeight * 2 / 5
	simra.GetInstance().AddTextSprite("(click to exit this page)",
		60, // fontsize
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(temporary.W), int(temporary.H)),
		temporary)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (howtoplay *howtoplay) Drive() {
	if howtoplay.nextScene != nil {
		simra.GetInstance().SetScene(howtoplay.nextScene)
	}
}

// OnTouchBegin is called when howtoplay scene is Touched.
func (howtoplay *howtoplay) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called when howtoplay scene is Touched and moved.
func (howtoplay *howtoplay) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called when howtoplay scene is Touched and it is released.
func (howtoplay *howtoplay) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	howtoplay.nextScene = &menu{}
}
