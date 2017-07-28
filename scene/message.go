package scene

import (
	"image"
	"image/color"
	"time"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

type message struct {
	sprite *simra.Sprite
	game   *game
}

func (m *message) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		// should be a command. ignore.
		return
	}

	switch c.commandtype {
	case commandShowMessage:
		message := c.data.(string)

		m.sprite = simra.NewSprite()
		simra.GetInstance().AddSprite2(m.sprite)

		tex := simra.NewTextTexture(
			message,
			40, // fontsize
			color.RGBA{255, 255, 255, 255},
			image.Rect(0, 0, config.ScreenWidth, 80),
		)
		m.sprite.ReplaceTexture2(tex)
		m.sprite.X, m.sprite.Y = config.ScreenWidth/2, 300
		m.sprite.W, m.sprite.H = config.ScreenWidth, 80

		go func() {
			select {
			case <-time.After(2 * time.Second):
				m.game.eventqueue <- newCommand(commandHideMessage, nil)
			}
		}()

	case commandHideMessage:
		simra.GetInstance().RemoveSprite(m.sprite)
	}
}
