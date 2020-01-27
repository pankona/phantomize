package scene

import (
	"image/color"

	"golang.org/x/mobile/asset"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// result represents a scene object for result
type result struct {
	simra        simra.Simraer
	text         simra.Spriter
	currentStage int
	again        simra.Spriter
	nextScene    simra.Driver
	bgm          simra.Audioer
	beep         asset.File
}

// Initialize initializes result scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (r *result) Initialize(sim simra.Simraer) {
	r.simra = sim

	r.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)
	r.initialize()
}

func (r *result) initialize() {
	r.text = r.simra.NewSprite()
	initTextSprite(r.simra, r.text, "Thank you for playing!",
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*4/6,
		60, color.RGBA{255, 0, 0, 255})

	r.again = r.simra.NewSprite()
	initTextSprite(r.simra, r.again, "try again?",
		config.ScreenWidth, 80, config.ScreenWidth/2, config.ScreenHeight*2/6,
		60, color.RGBA{255, 0, 0, 255})

	//simra.GetInstance().AddTouchListener(menu)
	r.again.AddTouchListener(&again{result: r})

	r.bgm = simra.NewAudio()
	resource, err := asset.Open("bgm3.mp3")
	if err != nil {
		panic(err.Error())
	}
	r.bgm.Play(resource, true, func(err error) {})

	r.beep, err = asset.Open("start_game.mp3")
	if err != nil {
		panic(err.Error())
	}
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (r *result) Drive() {
	if r.nextScene != nil {
		a := simra.NewAudio()
		a.Play(r.beep, false, func(err error) {})

		r.bgm.Stop()
		r.simra.SetScene(r.nextScene)
	}
}

// OnTouchBegin is called when result scene is Touched.
func (r *result) OnTouchBegin(x, y float32) {
	// nop
}

// OnTouchMove is called when result scene is Touched and moved.
func (r *result) OnTouchMove(x, y float32) {
	// nop
}

// OnTouchEnd is called when result scene is Touched and it is released.
func (r *result) OnTouchEnd(x, y float32) {
	// nop
}

type again struct {
	*result
}

func (a *again) OnTouchEnd(x, y float32) {
	a.result.nextScene = &game{currentStage: a.result.currentStage}
}

type next struct {
	*result
}

func (n *next) OnTouchEnd(x, y float32) {
	n.result.nextScene = &briefing{currentStage: n.result.currentStage + 1}
}
