package scene

import (
	"fmt"
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
)

type charainfo struct {
	icon   *simra.Sprite
	sprite [4]*simra.Sprite
	game   *game
}

func (ci *charainfo) initialize() {
	ci.icon = simra.NewSprite()
	ci.icon.W, ci.icon.H = 200, 200
	ci.icon.X, ci.icon.Y = 500, 120
	simra.GetInstance().AddSprite2(ci.icon)

	for i := 0; i < len(ci.sprite); i++ {
		ci.sprite[i] = simra.NewSprite()
		ci.sprite[i].X, ci.sprite[i].Y = 800, (float32)(165-(i*50))
		ci.sprite[i].H, ci.sprite[i].W = 100, 300
		simra.GetInstance().AddSprite2(ci.sprite[i])
	}
}

func (ci *charainfo) isCtrlButtonSelected(s *simra.Sprite) bool {
	ctrls := ci.game.ctrlButton
	for i, _ := range ctrls {
		if ctrls[i] == s {
			return true
		}
	}
	return false
}

func (ci *charainfo) showUnitInfo(s *simra.Sprite, unittype string) {

	asset := ci.game.assetNameByCtrlButton(s)
	tex := simra.NewImageTexture(asset, image.Rect(0, 0, 384, 384))
	ci.icon.ReplaceTexture2(tex)

	u := getUnitByUnitType(unittype)
	tex = simra.NewTextTexture(
		fmt.Sprintf("%s", unittype),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[0].ReplaceTexture2(tex)

	tex = simra.NewTextTexture(
		fmt.Sprintf("HP: %d", u.hp),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[1].ReplaceTexture2(tex)
	simra.LogDebug("@@@@@@@@@@@ showUnitInfo!! HP = %s", u.hp)

	tex = simra.NewTextTexture(
		fmt.Sprintf("COST: %d", u.cost),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[2].ReplaceTexture2(tex)

	tex = simra.NewTextTexture(
		fmt.Sprintf("SPEED: %0.1f", u.moveSpeed),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[3].ReplaceTexture2(tex)
}

func (ci *charainfo) showUnitStatus(u *unitBase) {
}

func (ci *charainfo) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandUpdateSelection:
		selecting := c.data.(*simra.Sprite)
		if ci.isCtrlButtonSelected(selecting) {
			ci.showUnitInfo(selecting, ci.game.unitIDBySprite(selecting))
		}
	}
}
