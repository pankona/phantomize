package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

type instruction struct {
	simra  simra.Simraer
	sprite simra.Spriter
	game   *game
	texs   [2]*simra.Texture
}

func (in *instruction) initialize() {
	in.sprite = in.simra.NewSprite()
	in.simra.AddSprite(in.sprite)
	in.texs[0] = in.simra.NewTextTexture(
		"choose first unit from here ↓",
		40, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, config.ScreenWidth, 80),
	)

	in.texs[1] = in.simra.NewTextTexture(
		"↑ tap field to summon the unit",
		40, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, config.ScreenWidth, 80),
	)
}

func (in *instruction) OnEvent(i interface{}) {
	c := i.(*command)
	switch c.commandtype {
	case commandGameStarted:
		in.sprite.ReplaceTexture(in.texs[0])
		in.sprite.SetPosition(config.ScreenWidth/2+240, 260)
		in.sprite.SetScale(config.ScreenWidth, 80)

	case commandUpdateSelection:
		in.sprite.ReplaceTexture(in.texs[1])
		in.sprite.SetPosition(config.ScreenWidth/2+275, 260)
		in.sprite.SetScale(config.ScreenWidth, 80)

	case commandSpawn:
		in.simra.RemoveSprite(in.sprite)
	}
}
