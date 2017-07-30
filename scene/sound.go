package scene

import (
	"github.com/pankona/gomo-simra/simra"

	"golang.org/x/mobile/asset"
)

type sound struct {
	game *game
}

/*
atk_fire.mp3
atk_hammer.mp3
atk_sword.mp3
//bgm1.mp3
//bgm2.mp3
//bgm3.mp3
enemy_atk.mp3
//enemy_dead.mp3
//enemy_spawn.mp3
//player_dead.mp3
recall.mp3
//select.mp3
//start_game.mp3
//summoning.mp3
*/

func (s *sound) play(assetName string) {
	a := simra.NewAudio()
	resource, err := asset.Open(assetName)
	if err != nil {
		panic(err.Error())
	}
	a.Play(resource, false, func() {})
}

func (s *sound) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}
	switch c.commandtype {
	case commandSpawn:
		d := c.data.(uniter)
		if d.IsAlly() {
			s.play("summoning.mp3")
		} else {
			s.play("enemy_spawn.mp3")
		}

	case commandDead:
		d := c.data.(uniter)
		if d.IsAlly() {
			s.play("player_dead.mp3")
		} else {
			s.play("enemy_dead.mp3")
		}

	case commandRecall:
		s.play("summoning.mp3")

	case commandRecalled:
		s.play("recall.mp3")

	case commandUpdateSelection:
		s.play("select.mp3")
	}
}
