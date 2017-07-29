package scene

import "github.com/pankona/gomo-simra/simra"

type charainfo struct {
	sprite [4]*simra.Sprite
	game   *game
}

func (ci *charainfo) initialize() {
	for i := 0; i < len(ci.sprite); i++ {
		ci.sprite[i] = simra.NewSprite()
		simra.GetInstance().AddSprite2(ci.sprite[i])
	}
}

func (ci *charainfo) isSelectionCtrlButton(s *simra.Sprite) bool {
	ctrls := ci.game.ctrlButton
	for i, _ := range ctrls {
		if ctrls[i] == s {
			return true
		}
	}
	return false
}

func (ci *charainfo) showUnitInfo(unitID string) {
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
		if ci.isSelectionCtrlButton(selecting) {
			ci.showUnitInfo(ci.game.unitIDBySprite(selecting))
		}
	}
}
