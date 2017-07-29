package scene

import (
	"fmt"
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
)

type charainfo struct {
	icon        *simra.Sprite
	sprite      [4]*simra.Sprite
	recall      [2]*simra.Sprite
	recallBGTex *simra.Texture
	displaying  *simra.Sprite
	game        *game
}

type recallTouchListener struct {
	unit uniter
	game *game
}

func (r *recallTouchListener) OnTouchBegin(x, y float32) {
	// nop
}

func (r *recallTouchListener) OnTouchMove(x, y float32) {
	// nop
}

func (r *recallTouchListener) OnTouchEnd(x, y float32) {
	action := r.unit.GetAction()
	if action != nil && action.actiontype == actionRecall {
		// now recalling. ignore
		return
	}

	r.game.eventqueue <- newCommand(commandRecall, r.unit)
}

func (ci *charainfo) initialize() {
	ci.icon = simra.NewSprite()
	ci.icon.X, ci.icon.Y = 500, 120
	ci.icon.W, ci.icon.H = 200, 200

	for i := 0; i < len(ci.sprite); i++ {
		ci.sprite[i] = simra.NewSprite()
		ci.sprite[i].X, ci.sprite[i].Y = 800, (float32)(165-(i*50))
		ci.sprite[i].H, ci.sprite[i].W = 100, 300
	}

	ci.recall[0] = simra.NewSprite()
	ci.recall[0].X, ci.recall[0].Y = 500, 95
	ci.recall[0].W, ci.recall[0].H = 150, 40
	ci.recallBGTex = simra.NewImageTexture("cursor.png", image.Rect(0, 0, 30, 30))

	ci.recall[1] = simra.NewSprite()
	ci.recall[1].X, ci.recall[1].Y = 500, 70
	ci.recall[1].H, ci.recall[1].W = 100, 300
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

	// recall label only for players
	if u.IsAlly() {
		simra.GetInstance().AddSprite2(ci.recall[0])
		ci.recall[0].ReplaceTexture2(ci.recallBGTex)
		ci.recall[0].AddTouchListener(&recallTouchListener{
			unit: u,
			game: ci.game,
		})

		simra.GetInstance().AddSprite2(ci.recall[1])
		tex = simra.NewTextTexture(
			"RECALL",
			30, // fontsize
			color.RGBA{255, 0, 0, 255},
			image.Rect(0, 0, 300, 80),
		)
		ci.recall[1].ReplaceTexture2(tex)
	}
}

func (ci *charainfo) hideCharaInfo() {
	simra.LogDebug("@@@@@@ hideCharaInfo!!")
	simra.GetInstance().RemoveSprite(ci.icon)
	for i, _ := range ci.sprite {
		simra.GetInstance().RemoveSprite(ci.sprite[i])
	}
	simra.GetInstance().RemoveSprite(ci.recall[0])
	simra.GetInstance().RemoveSprite(ci.recall[1])
}

func (ci *charainfo) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandUpdateSelection:
		fmt.Println("@@@@@@ selection is updated")
		switch selecting := c.data.(type) {
		case *simra.Sprite:
			if ci.isCtrlButtonSelected(selecting) {
				ci.hideCharaInfo()
				ci.showUnitInfo(selecting, ci.game.unitIDBySprite(selecting))
			}
			ci.displaying = selecting
		case *unitTouchListener:
			ci.hideCharaInfo()
			ci.showUnitStatus(selecting.sprite, selecting.uniter)
			ci.displaying = selecting.sprite
		}

	case commandUnsetSelection:
		fmt.Println("@@@@@@ selection is unset")
		ci.hideCharaInfo()

	case commandDamage:
		d := c.data.(*damage)
		if ci.displaying == d.unit.GetSprite() {
			// update chara info
			ci.hideCharaInfo()
			ci.showUnitStatus(d.unit.GetSprite(), d.unit)
		}

	case commandDead:
		d := c.data.(uniter)
		if ci.displaying == d.GetSprite() {
			ci.hideCharaInfo()
		}
	}
}
