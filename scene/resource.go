package scene

import (
	"fmt"
	"image"
	"image/color"

	"github.com/pankona/gomo-simra/simra"
)

type resource struct {
	balance int
	sprite  *simra.Sprite
	game    *game
}

func (r *resource) initialize() {
	r.sprite = simra.NewSprite()
	simra.GetInstance().AddSprite2(r.sprite)

	resstr := fmt.Sprintf("$ %d", r.balance)
	tex := simra.NewTextTexture(
		resstr,
		40, // fontsize
		color.RGBA{255, 255, 255, 255},
		image.Rect(0, 0, 100, 80),
	)
	r.sprite.ReplaceTexture2(tex)
	r.sprite.X, r.sprite.Y = 100, 100
	r.sprite.W, r.sprite.H = 100, 80
}

func (r *resource) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}
	switch c.commandtype {
	case commandSpawned:
		u, ok := c.data.(uniter)
		if !ok || !u.IsAlly() {
			// ignore
			break
		}
		r.balance -= u.GetCost()

		// update texture
		resstr := fmt.Sprintf("$ %d", r.balance)
		tex := simra.NewTextTexture(
			resstr,
			40, // fontsize
			color.RGBA{255, 255, 255, 255},
			image.Rect(0, 0, 100, 80),
		)
		r.sprite.ReplaceTexture2(tex)

	case commandDead:
		u, ok := c.data.(uniter)
		if !ok || u.IsAlly() {
			// ignore
			break
		}
		r.balance += u.GetCost()

		// update texture
		resstr := fmt.Sprintf("$ %d", r.balance)
		tex := simra.NewTextTexture(
			resstr,
			40, // fontsize
			color.RGBA{255, 255, 255, 255},
			image.Rect(0, 0, 100, 80),
		)
		r.sprite.ReplaceTexture2(tex)
	}

}
