package scene

import (
	"image"
	"image/color"

	"golang.org/x/mobile/asset"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

// Title represents a scene object for Title
type Title struct {
	simra     simra.Simraer
	text      simra.Spriter
	nextScene simra.Driver
	bgm       simra.Audioer
	beep      asset.File
}

// Initialize initializes title scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (title *Title) Initialize(sim simra.Simraer) {
	title.simra = sim
	title.simra.SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)
	title.initialize()
}

func (title *Title) initialize() {
	title.text.SetScale(config.ScreenWidth, 80)
	title.text.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	title.simra.AddSprite(title.text)
	tex := title.simra.NewTextTexture("phantomize",
		60, color.RGBA{255, 0, 0, 255}, image.Rect(0, 0, int(title.text.GetScale().W), int(title.text.GetScale().H)))
	title.text.ReplaceTexture(tex)

	bg := title.simra.NewSprite()
	bg.SetScale(config.ScreenWidth, config.ScreenHeight)
	bg.SetPosition(config.ScreenWidth/2, config.ScreenHeight/2)
	title.simra.AddSprite(bg)
	tex = title.simra.NewImageTexture("title.png", image.Rect(0, 0, 1280, 720))
	bg.ReplaceTexture(tex)

	text2 := title.simra.NewSprite()
	text2.SetScale(config.ScreenWidth, 80)
	text2.SetPosition(config.ScreenWidth/2, config.ScreenHeight/6*1)
	title.simra.AddSprite(text2)
	tex = title.simra.NewTextTexture("tap to start!",
		60, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, int(title.text.GetScale().W), int(title.text.GetScale().H)))
	text2.ReplaceTexture(tex)

	title.simra.AddTouchListener(title)

	title.bgm = simra.NewAudio()
	_, err := asset.Open("bgm1.mp3")
	if err != nil {
		panic(err.Error())
	}
	//title.bgm.Play(resource, true, func(err error) {})

	title.beep, err = asset.Open("start_game.mp3")
	if err != nil {
		panic(err.Error())
	}
}

// Drive is called from simra.
// This is used to update sprites position.
// This will be called 60 times per sec.
func (title *Title) Drive() {
	if title.nextScene != nil {
		a := simra.NewAudio()
		a.Play(title.beep, false, func(err error) {})

		title.bgm.Stop()
		title.simra.SetScene(title.nextScene)
	}
}

// OnTouchBegin is called when Title scene is Touched.
func (title *Title) OnTouchBegin(x, y float32) {
}

// OnTouchMove is called when Title scene is Touched and moved.
func (title *Title) OnTouchMove(x, y float32) {
}

// OnTouchEnd is called when Title scene is Touched and it is released.
func (title *Title) OnTouchEnd(x, y float32) {
	// scene end. go to next scene
	//title.nextScene = &menu{}
	title.nextScene = &game{}
}
