package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene/config"
)

const (
	ctrlPanelHeight = 220
)

func (g *game) assetNameByCtrlButton(s simra.Spriter) string {
	switch s {
	case g.ctrlButton[0]:
		return "ctrlbutton1.png"
	case g.ctrlButton[1]:
		return "ctrlbutton2.png"
	case g.ctrlButton[2]:
		return "ctrlbutton3.png"
	}
	return ""
}

func (g *game) initCtrlPanel() {
	g.ctrlPanel.SetScale(config.ScreenWidth, ctrlPanelHeight)
	g.ctrlPanel.SetPosition(config.ScreenWidth/2, g.ctrlPanel.GetScale().H/2)
	g.simra.AddSprite(g.ctrlPanel)

	var tex *simra.Texture
	tex = g.simra.NewImageTexture("panel.png", image.Rect(0, 0, 1280, 240))
	g.ctrlPanel.ReplaceTexture(tex)

	g.ctrlButton = make([]simra.Spriter, 3)
	for i := range g.ctrlButton {
		g.ctrlButton[i] = g.simra.NewSprite()
		g.ctrlButton[i].SetScale(64, 64)
		g.ctrlButton[i].SetPosition(1000+(64+50)*(i%3), 44+(64+5)*2-(64+5)*(i/3))
		g.simra.AddSprite(g.ctrlButton[i])
		tex = g.simra.NewImageTexture(g.assetNameByCtrlButton(g.ctrlButton[i]),
			image.Rect(0, 0, 384, 384))
		g.ctrlButton[i].ReplaceTexture(tex)

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
	c.game.eventqueue <- newCommand(commandUpdateSelection, c.game.ctrlButton[c.id])
}
