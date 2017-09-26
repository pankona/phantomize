package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

type howtoplay struct {
	simra     simra.Simraer
	text      simra.Spriter
	nextScene simra.Driver
}

// Initialize initializes howtoplay scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (howtoplay *howtoplay) Initialize(sim simra.Simraer) {
	howtoplay.simra = sim

	howtoplay.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)
	howtoplay.initialize()
}

func (howtoplay *howtoplay) initialize() {
	howtoplay.text.SetScale(config.ScreenWidth, 80)
	howtoplay.text.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	howtoplay.simra.AddSprite(howtoplay.text)

	howtoplay.simra.AddTouchListener(howtoplay)

	// temporary text (will be removed)
	temporary := howtoplay.simra.NewSprite()
	temporary.SetScale(config.ScreenWidth, 80)
	temporary.SetPosition(config.ScreenWidth/2, config.ScreenHeight*2/5)
	howtoplay.simra.AddSprite(temporary)

	var tex *simra.Texture
	tex = howtoplay.simra.NewTextTexture("How To Play",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(howtoplay.text.GetScale().W), int(howtoplay.text.GetScale().H)))
	howtoplay.text.ReplaceTexture(tex)

	tex = howtoplay.simra.NewTextTexture("(click to exit this page)",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(temporary.GetScale().W), int(temporary.GetScale().H)))
	temporary.ReplaceTexture(tex)

}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (howtoplay *howtoplay) Drive() {
	if howtoplay.nextScene != nil {
		howtoplay.simra.SetScene(howtoplay.nextScene)
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
