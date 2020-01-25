package scene

import (
	"fmt"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
)

type resource struct {
	simra   simra.Simraer
	balance int
	sprite  simra.Spriter
	game    *game
}

func (r *resource) initialize() {
	r.sprite = r.simra.NewSprite()
	r.simra.AddSprite(r.sprite)

	resstr := fmt.Sprintf("$ %d", r.balance)
	tex := r.simra.NewTextTexture(
		resstr,
		40, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 100, 80),
	)
	r.sprite.ReplaceTexture(tex)
	r.sprite.SetPosition(100, 100)
	r.sprite.SetScale(100, 80)
}

func (r *resource) updateResourceInfo() {
	resstr := fmt.Sprintf("$ %d", r.balance)
	tex := r.simra.NewTextTexture(
		resstr,
		40, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 100, 80),
	)
	r.sprite.ReplaceTexture(tex)
}

func (r *resource) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}
	switch c.commandtype {
	case commandSpawn:
		u, ok := c.data.(uniter)
		if !ok || !u.IsAlly() {
			// ignore
			break
		}
		r.balance -= u.GetCost()
		r.updateResourceInfo()

	case commandRecalled:
		u := c.data.(uniter)
		r.balance += u.GetCost()
		r.updateResourceInfo()

	case commandDead:
		u, ok := c.data.(uniter)
		if !ok || u.IsAlly() {
			// ignore
			break
		}
		r.balance += u.GetCost()
		r.updateResourceInfo()
	}
}
