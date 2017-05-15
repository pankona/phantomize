package scene

import (
	"image/color"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// result represents a scene object for result
type result struct {
	text         simra.Sprite
	currentStage int
	again        simra.Sprite
	next         simra.Sprite
	nextScene    simra.Driver
}

// Initialize initializes result scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (result *result) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	result.initialize()

	simra.LogDebug("[OUT]")
}

func (result *result) initialize() {
	initTextSprite(&result.text, "result",
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*4/6,
		60, color.RGBA{255, 0, 0, 255})
	initTextSprite(&result.again, "try again",
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*2/6,
		60, color.RGBA{255, 0, 0, 255})
	initTextSprite(&result.next, "go to next stage",
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*1/6,
		60, color.RGBA{255, 0, 0, 255})
	//simra.GetInstance().AddTouchListener(menu)
	result.again.AddTouchListener(&again{result: result})
	result.next.AddTouchListener(&next{result: result})

}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (result *result) Drive() {
	// nop
	if result.nextScene != nil {
		simra.GetInstance().SetScene(result.nextScene)
	}
}

// OnTouchBegin is called when result scene is Touched.
func (result *result) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when result scene is Touched and moved.
func (result *result) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when result scene is Touched and it is released.
func (result *result) OnTouchEnd(x, y float32) {
	// nop
}

type again struct {
	*result
}

func (again *again) OnTouchEnd(x, y float32) {
	again.result.nextScene = &game{currentStage: again.result.currentStage}
}

type next struct {
	*result
}

func (next *next) OnTouchEnd(x, y float32) {
	next.result.nextScene = &briefing{currentStage: next.result.currentStage + 1}
}
