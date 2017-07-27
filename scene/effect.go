package scene

import (
	"fmt"
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

	// smoke animation
	numOfAnimation := 3
	w := 512 / numOfAnimation
	h := 528 / 4
	resource := "smoke.png"
	animationSet := simra.NewAnimationSet()
	for i := 0; i < numOfAnimation; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, 0, ((int)(w)*(i+1))-1, int(h))))
	}
	// TODO: don't relay on time. use fps based animation control
	animationSet.SetInterval(100 * time.Millisecond)
	e.animations[resource] = animationSet

	// attack animation (1)
	numOfAnimation = 5
	w = 600 / numOfAnimation
	h = 120
	resource = "atkeffect1.png"
	animationSet = simra.NewAnimationSet()
	for i := 0; i < numOfAnimation; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, 0, ((int)(w)*(i+1))-1, int(h))))
	}
	// TODO: don't relay on time. use fps based animation control
	animationSet.SetInterval(100 * time.Millisecond)
	e.animations[resource] = animationSet

	// attack animation (2)
	numOfAnimation = 7
	w = 840 / numOfAnimation
	h = 120
	resource = "atkeffect2.png"
	animationSet = simra.NewAnimationSet()
	for i := 0; i < numOfAnimation; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, 0, ((int)(w)*(i+1))-1, int(h))))
	}
	// TODO: don't relay on time. use fps based animation control
	animationSet.SetInterval(100 * time.Millisecond)
	e.animations[resource] = animationSet

	// attack animation (3)
	numOfAnimation = 8
	w = 960 / numOfAnimation
	h = 120
	resource = "atkeffect3.png"
	animationSet = simra.NewAnimationSet()
	for i := 0; i < numOfAnimation; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, 0, ((int)(w)*(i+1))-1, int(h))))
	}
	// TODO: don't relay on time. use fps based animation control
	animationSet.SetInterval(100 * time.Millisecond)
	e.animations[resource] = animationSet

	// attack animation (4)
	w = 600 / 5
	h = 120
	resource = "atkeffect4.png"
	animationSet = simra.NewAnimationSet()
	for i := 0; i < 5; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, 0, ((int)(w)*(i+1))-1, int(h))))
	}
	for i := 0; i < 3; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, h, ((int)(w)*(i+1))-1, int(h))))
	}
	// TODO: don't relay on time. use fps based animation control
	animationSet.SetInterval(100 * time.Millisecond)
	e.animations[resource] = animationSet

	// attack animation (5)
	w = 600 / 6
	h = 120
	resource = "atkeffect5.png"
	animationSet = simra.NewAnimationSet()
	for i := 0; i < 6; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, 0, ((int)(w)*(i+1))-1, int(h))))
	}
	for i := 0; i < 6; i++ {
		animationSet.AddTexture(simra.NewImageTexture(resource,
			image.Rect((int)(w)*i, h, ((int)(w)*(i+1))-1, int(h))))
	}
	// TODO: don't relay on time. use fps based animation control
	animationSet.SetInterval(100 * time.Millisecond)
	e.animations[resource] = animationSet

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
			break
		}
		sprite := simra.NewSprite()
		sprite.W = 512 / 3
		sprite.H = 528 / 4
		x, y := p.GetPosition()
		sprite.X, sprite.Y = x-10, y+20

		animationSet := e.animations["smoke.png"]
		sprite.AddAnimationSet("smoke.png", animationSet)
		simra.GetInstance().AddSprite2(sprite)
		sprite.StartAnimation("smoke.png", true, func() {})
		e.effects[p.GetID()] = sprite

	case commandSpawned:
		p, ok := c.data.(uniter)
		if !ok {
			// ignore
			break
		}
		sprite := e.effects[p.GetID()]
		sprite.StopAnimation()
		delete(e.effects, p.GetID())
		simra.GetInstance().RemoveSprite(sprite)

	case commandAttack:
		p, ok := c.data.(uniter)
		if !ok {
			// ignore
			break
		}
		target := p.GetTarget()
		tx, ty := target.GetPosition()

		sprite := simra.NewSprite()
		sprite.W = 64
		sprite.H = 64
		sprite.X, sprite.Y = tx, ty
		var atkeffect string
		switch p.GetUnitType() {
		case "player1":
			atkeffect = "atkeffect1.png"
		case "player2":
			atkeffect = "atkeffect2.png"
		case "player3":
			atkeffect = "atkeffect3.png"
		case "enemy1":
			atkeffect = "atkeffect4.png"
		case "enemy2":
			atkeffect = "atkeffect5.png"
		default:
			simra.LogError("[%s]'s atkeffect is not loaded!", p.GetUnitType())
			panic("atkeffect is not loaded!")
		}

		animationSet := e.animations[atkeffect]
		sprite.AddAnimationSet(atkeffect, animationSet)
		simra.GetInstance().AddSprite2(sprite)
		sprite.StartAnimation(atkeffect, true, func() {})
		e.effects[p.GetID()] = sprite

	case commandAttackEnd:
		p, ok := c.data.(uniter)
		if !ok {
			// ignore
			break
		}

		fmt.Printf("[effect][%s] ends attacking\n", p.GetID())

		sprite := e.effects[p.GetID()]
		sprite.StopAnimation()
		delete(e.effects, p.GetID())
		simra.GetInstance().RemoveSprite(sprite)
	}
}
