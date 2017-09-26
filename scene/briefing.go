package scene

import (
	"image"
	"image/color"
	"strconv"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// briefing represents a scene object for briefing
type briefing struct {
	simra        simra.Simraer
	text         simra.Spriter
	currentStage int
	nextScene    simra.Driver
}

// Initialize initializes briefing scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (briefing *briefing) Initialize(sim simra.Simraer) {
	briefing.simra = sim

	briefing.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)
	briefing.initialize()
}

func (briefing *briefing) initialize() {
	initTextSprite(briefing.simra, briefing.text, "briefing for stage "+strconv.Itoa(briefing.currentStage),
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*4/6,
		60, color.RGBA{255, 0, 0, 255})

	// temporary text (will be removed)
	temporary := briefing.simra.NewSprite()
	temporary.SetScale(config.ScreenWidth, 80)
	temporary.SetPosition(config.ScreenWidth/2, config.ScreenHeight*2/5)
	briefing.simra.AddSprite(temporary)
	tex := briefing.simra.NewTextTexture("(click to go to next scene)",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, temporary.GetScale().W, temporary.GetScale().H))
	temporary.ReplaceTexture(tex)

	briefing.simra.AddTouchListener(briefing)
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (briefing *briefing) Drive() {
	// nop
	if briefing.nextScene != nil {
		briefing.simra.SetScene(briefing.nextScene)
	}
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
	briefing.nextScene = &game{currentStage: briefing.currentStage}
}
