package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

const (
	ctrlPanelHeight = 220
)

func (g *game) initCtrlPanel() {
	g.ctrlPanel.W = config.ScreenWidth
	g.ctrlPanel.H = ctrlPanelHeight
	g.ctrlPanel.X = config.ScreenWidth / 2
	g.ctrlPanel.Y = g.ctrlPanel.H / 2
	simra.GetInstance().AddSprite("panel.png",
		image.Rect(0, 0, 1280, 240),
		&g.ctrlPanel)

	g.ctrlButton = make([]simra.Sprite, 9)
	for i := range g.ctrlButton {
		g.ctrlButton[i].W = 64
		g.ctrlButton[i].H = 64
		g.ctrlButton[i].X = (float32)(1010 + (64+45)*(i%3))
		g.ctrlButton[i].Y = (float32)(44 + (64+5)*(i/3))
		simra.GetInstance().AddSprite("player.png",
			image.Rect(0, 0, 384, 384),
			&g.ctrlButton[i])

		g.ctrlButton[i].AddTouchListener(&ctrlButtonTouchListener{id: i, game: g})
	}
}

type ctrlButtonTouchListener struct {
	id   int
	game *game
}

func (c *ctrlButtonTouchListener) OnTouchBegin(x, y float32) {
	// nop
}

func (c *ctrlButtonTouchListener) OnTouchMove(x, y float32) {
	// nop
}

func (c *ctrlButtonTouchListener) OnTouchEnd(x, y float32) {
	c.game.eventqueue <- newCommand(commandUpdateSelection, &c.game.ctrlButton[c.id])
}
