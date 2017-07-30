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
	text      simra.Sprite
	nextScene simra.Driver
	bgm       simra.Audioer
	beep      asset.File
}

// Initialize initializes title scene
// This is called from simra.
// simra.GetInstance().SetDesiredScreenSize should be called to determine
// screen size of this scene.
func (title *Title) Initialize() {
	simra.LogDebug("[IN]")

	simra.GetInstance().SetDesiredScreenSize(config.ScreenWidth, config.ScreenHeight)

	// initialize sprites
	title.initialize()

	simra.LogDebug("[OUT]")
}

func (title *Title) initialize() {
	title.text.W = config.ScreenWidth
	title.text.H = 80
	title.text.X = config.ScreenWidth / 2
	title.text.Y = config.ScreenHeight / 2
	simra.GetInstance().AddTextSprite("title",
		60, // fontsize
		color.RGBA{255, 0, 0, 255},
		image.Rect(0, 0, int(title.text.W), int(title.text.H)),
		&title.text)

	simra.GetInstance().AddTouchListener(title)

	title.bgm = simra.NewAudio()
	resource, err := asset.Open("bgm1.mp3")
	if err != nil {
		panic(err.Error())
	}
	title.bgm.Play(resource, true, func() {})

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
		a.Play(title.beep, false, func() {})

		title.bgm.Stop()
		simra.GetInstance().SetScene(title.nextScene)
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
