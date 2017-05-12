package scene

import (
	"image/color"
	"strconv"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// briefing represents a scene object for briefing
type briefing struct {
	text         simra.Sprite
	currentStage int
}

// Initialize initializes briefing scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (briefing *briefing) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	briefing.initialize()

	simra.LogDebug("[OUT]")
}

func (briefing *briefing) initialize() {
	initTextSprite(&briefing.text, "briefing for stage "+strconv.Itoa(briefing.currentStage),
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*4/6,
		60, color.RGBA{255, 0, 0, 255})
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (briefing *briefing) Drive() {
	// nop
}

// OnTouchBegin is called when briefing scene is Touched.
func (briefing *briefing) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when briefing scene is Touched and moved.
func (briefing *briefing) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when briefing scene is Touched and it is released.
func (briefing *briefing) OnTouchEnd(x, y float32) {
	// nop
}
