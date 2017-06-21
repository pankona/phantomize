package scene

import (
	"fmt"
	"image"

	"github.com/pankona/gomo-simra/simra"
)

type sampleUnit struct {
	*unitBase
	sprite simra.Sprite
}

func (u *sampleUnit) Initialize() {
	simra.GetInstance().AddSprite("player.png",
		image.Rect(0, 0, 384, 384),
		&u.sprite)
}

func (u *sampleUnit) SetPosition(p position) {
	u.position = p
}

func (u *sampleUnit) GetPosition() position {
	return u.position
}

func (u *sampleUnit) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case SPAWN:
		d, ok := c.data.(*sampleUnit)
		if !ok {
			panic("unexpected command received. fatal.")
		}
		if u.id != d.GetID() {
			// nop
			return
		}

		u.sprite.W = 64
		u.sprite.H = 64
		u.sprite.X = (float32)(d.position.x)
		u.sprite.Y = (float32)(d.position.y)

		// TODO: spawn myself
		fmt.Printf("@@@@@@ [SPAWN] i'm %s\n", u.GetID())
	default:
		// nop
	}
}
