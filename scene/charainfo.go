package scene

import (
	"fmt"
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
)

type charainfo struct {
	simra       simra.Simraer
	icon        simra.Spriter
	sprite      [4]simra.Spriter
	recall      [2]simra.Spriter
	recallBGTex *simra.Texture
	displaying  simra.Spriter
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
	ci.icon = ci.simra.NewSprite()
	ci.icon.SetPosition(500, 120)
	ci.icon.SetScale(200, 200)

	for i := 0; i < len(ci.sprite); i++ {
		ci.sprite[i] = ci.simra.NewSprite()
		ci.sprite[i].SetPosition(800, 165-(i*50))
		ci.sprite[i].SetScale(300, 100)
	}

	ci.recall[0] = ci.simra.NewSprite()
	ci.recall[0].SetPosition(500, 95)
	ci.recall[0].SetScale(150, 40)
	ci.recallBGTex = ci.simra.NewImageTexture("cursor.png", image.Rect(0, 0, 30, 30))

	ci.recall[1] = ci.simra.NewSprite()
	ci.recall[1].SetPosition(500, 70)
	ci.recall[1].SetScale(300, 100)
}

func (ci *charainfo) isCtrlButtonSelected(s simra.Spriter) bool {
	ctrls := ci.game.ctrlButton
	for i := range ctrls {
		if ctrls[i] == s {
			return true
		}
	}
	return false
}

func (ci *charainfo) showUnitInfo(s simra.Spriter, unittype string) {
	ci.simra.AddSprite(ci.icon)
	for i := 0; i < len(ci.sprite); i++ {
		ci.simra.AddSprite(ci.sprite[i])
	}

	asset := ci.game.assetNameByCtrlButton(s)
	tex := ci.simra.NewImageTexture(asset, image.Rect(0, 0, 384, 384))
	ci.icon.ReplaceTexture(tex)

	u := getUnitByUnitType(ci.simra, unittype)
	tex = ci.simra.NewTextTexture(
		fmt.Sprintf("%s", unittype),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[0].ReplaceTexture(tex)

	tex = ci.simra.NewTextTexture(
		fmt.Sprintf("HP: %d", u.hp),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[1].ReplaceTexture(tex)

	tex = ci.simra.NewTextTexture(
		fmt.Sprintf("COST: %d", u.cost),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[2].ReplaceTexture(tex)

	tex = ci.simra.NewTextTexture(
		fmt.Sprintf("SPEED: %0.1f", u.moveSpeed),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[3].ReplaceTexture(tex)
}

func (ci *charainfo) showUnitStatus(s simra.Spriter, u uniter) {
	ci.simra.AddSprite(ci.icon)
	for i := 0; i < len(ci.sprite); i++ {
		ci.simra.AddSprite(ci.sprite[i])
	}

	asset := ci.game.assetNameByUnitType(u.GetUnitType())
	tex := ci.simra.NewImageTexture(asset, image.Rect(0, 0, 384, 384))
	ci.icon.ReplaceTexture(tex)

	tex = ci.simra.NewTextTexture(
		fmt.Sprintf("%s", u.GetUnitType()),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[0].ReplaceTexture(tex)

	tex = ci.simra.NewTextTexture(
		fmt.Sprintf("HP: %d", u.GetHP()),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[1].ReplaceTexture(tex)

	tex = ci.simra.NewTextTexture(
		fmt.Sprintf("COST: %d", u.GetCost()),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[2].ReplaceTexture(tex)

	tex = ci.simra.NewTextTexture(
		fmt.Sprintf("SPEED: %0.1f", u.GetMoveSpeed()),
		30, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 300, 80),
	)
	ci.sprite[3].ReplaceTexture(tex)

	// recall label only for players
	if u.IsAlly() {
		ci.simra.AddSprite(ci.recall[0])
		ci.recall[0].ReplaceTexture(ci.recallBGTex)
		ci.recall[0].AddTouchListener(&recallTouchListener{
			unit: u,
			game: ci.game,
		})

		ci.simra.AddSprite(ci.recall[1])
		tex = ci.simra.NewTextTexture(
			"RECALL",
			30, // fontsize
			color.RGBA{255, 0, 0, 255},
			image.Rect(0, 0, 300, 80),
		)
		ci.recall[1].ReplaceTexture(tex)
	}
}

func (ci *charainfo) hideCharaInfo() {
	ci.simra.RemoveSprite(ci.icon)
	for i := range ci.sprite {
		ci.simra.RemoveSprite(ci.sprite[i])
	}
	ci.recall[0].RemoveAllTouchListener()
	ci.simra.RemoveSprite(ci.recall[0])
	ci.simra.RemoveSprite(ci.recall[1])
}

func (ci *charainfo) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandUpdateSelection:
		switch selecting := c.data.(type) {
		case simra.Spriter:
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
		ci.hideCharaInfo()

	case commandDamage:
		d := c.data.(*damage)
		if ci.displaying == d.unit.GetSprite() {
			// update chara info
			ci.hideCharaInfo()
			ci.showUnitStatus(d.unit.GetSprite(), d.unit)
		}

	case commandRecalled:
		fallthrough
	case commandDead:
		d := c.data.(uniter)
		if ci.displaying == d.GetSprite() {
			ci.hideCharaInfo()
		}
	}
}
