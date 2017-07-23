package scene

import (
	"image"
	"time"

	"github.com/pankona/gomo-simra/simra"
)

type effect struct {
	game       *game
	animations map[string]*simra.AnimationSet
	effects    map[string]*simra.Sprite
}

func (e *effect) initialize() {
	e.animations = make(map[string]*simra.AnimationSet)
	e.effects = make(map[string]*simra.Sprite)

	numOfAnimation := 3
	w := 512 / 3
	h := 528 / 4

	resource := "smoke.png"
	animationSet := simra.NewAnimationSet()
	for i := 0; i < numOfAnimation; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, 0, ((int)(w)*(i+1))-1, int(h))))
	}
	// TODO: don't relay on time. use fps based animation control
	animationSet.SetInterval(100 * time.Millisecond)
	e.animations["smoke"] = animationSet
}

func (e *effect) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandSpawn:
		p, ok := c.data.(uniter)
		if !ok {
			// ignore
		}
		sprite := simra.NewSprite()
		sprite.W = 512 / 3
		sprite.H = 528 / 4
		x, y := p.GetPosition()
		sprite.X, sprite.Y = x-10, y+20

		animationSet := e.animations["smoke"]
		sprite.AddAnimationSet("summoning", animationSet)
		simra.GetInstance().AddSprite2(sprite)
		sprite.StartAnimation("summoning", true, func() {})
		e.effects[p.GetID()] = sprite

	case commandSpawned:
		p, ok := c.data.(uniter)
		if !ok {
			// ignore
		}
		sprite := e.effects[p.GetID()]
		sprite.StopAnimation()
		delete(e.effects, p.GetID())
		simra.GetInstance().RemoveSprite(sprite)
	}
}
