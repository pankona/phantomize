package scene

import (
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/fps"
	"github.com/pankona/phantomize/scene/config"
)

type message struct {
	simra simra.Simraer
	game  *game
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

		sprite := m.simra.NewSprite()
		m.simra.AddSprite(sprite)

		tex := m.simra.NewTextTexture(message,
			40, color.RGBA{255, 255, 255, 255}, image.Rect(0, 0, config.ScreenWidth, 80))
		sprite.ReplaceTexture(tex)
		sprite.SetPosition(config.ScreenWidth/2, 300)
		sprite.SetScale(config.ScreenWidth, 80)

		go func() {
			select {
			case <-fps.After(2 * framePerSec):
				m.game.eventqueue <- newCommand(commandHideMessage, sprite)
			}
		}()

	case commandHideMessage:
		s := c.data.(simra.Spriter)
		m.simra.RemoveSprite(s)
	}
}
