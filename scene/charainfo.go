package scene

import (
	"fmt"
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
)

type charainfo struct {
	icon       *simra.Sprite
	sprite     [4]*simra.Sprite
	game       *game
	displaying *simra.Sprite
}

func (ci *charainfo) initialize() {
	ci.icon = simra.NewSprite()
	ci.icon.W, ci.icon.H = 200, 200
	ci.icon.X, ci.icon.Y = 500, 120

	for i := 0; i < len(ci.sprite); i++ {
		ci.sprite[i] = simra.NewSprite()
		ci.sprite[i].X, ci.sprite[i].Y = 800, (float32)(165-(i*50))
		ci.sprite[i].H, ci.sprite[i].W = 100, 300
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
	simra.GetInstance().AddSprite2(ci.icon)
	for i := 0; i < len(ci.sprite); i++ {
		simra.GetInstance().AddSprite2(ci.sprite[i])
	}

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

func (ci *charainfo) showUnitStatus(s *simra.Sprite, u uniter) {
	simra.GetInstance().AddSprite2(ci.icon)
	for i := 0; i < len(ci.sprite); i++ {
		simra.GetInstance().AddSprite2(ci.sprite[i])
	}

	asset := ci.game.assetNameByUnitType(u.GetUnitType())
	tex := simra.NewImageTexture(asset, image.Rect(0, 0, 384, 384))
	ci.icon.ReplaceTexture2(tex)

	tex = simra.NewTextTexture(
		fmt.Sprintf("%s", u.GetUnitType()),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[0].ReplaceTexture2(tex)

	tex = simra.NewTextTexture(
		fmt.Sprintf("HP: %d", u.GetHP()),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[1].ReplaceTexture2(tex)

	tex = simra.NewTextTexture(
		fmt.Sprintf("COST: %d", u.GetCost()),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[2].ReplaceTexture2(tex)

	tex = simra.NewTextTexture(
		fmt.Sprintf("SPEED: %0.1f", u.GetMoveSpeed()),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[3].ReplaceTexture2(tex)
}

func (ci *charainfo) hideCharaInfo() {
	for i, _ := range ci.sprite {
		simra.GetInstance().RemoveSprite(ci.icon)
		simra.GetInstance().RemoveSprite(ci.sprite[i])
	}
}

func (ci *charainfo) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandUpdateSelection:
		switch selecting := c.data.(type) {
		case *simra.Sprite:
			if ci.isCtrlButtonSelected(selecting) {
				ci.showUnitInfo(selecting, ci.game.unitIDBySprite(selecting))
			}
			ci.displaying = selecting
		case *unitTouchListener:
			ci.showUnitStatus(selecting.sprite, selecting.uniter)
			ci.displaying = selecting.sprite
		}

	case commandUnsetSelection:
		ci.hideCharaInfo()

	case commandDamage:
		d := c.data.(*damage)
		if ci.displaying == d.unit.GetSprite() {
			// update chara info
			ci.showUnitStatus(d.unit.GetSprite(), d.unit)
		}

	case commandDead:
		d := c.data.(uniter)
		if ci.displaying == d.GetSprite() {
			ci.hideCharaInfo()
		}
	}
}
