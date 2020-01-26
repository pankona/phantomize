package scene

type sound struct {
	game *game
}

func (s *sound) play(assetName string) {
	//a := simra.NewAudio()
	//resource, err := asset.Open(assetName)
	//if err != nil {
	//	panic(err.Error())
	//}
	//err = a.Play(resource, false, func(err error) {
	//	if err != nil {
	//		simlog.Error(err.Error())
	//	}
	//})
	//if err != nil {
	//	panic(err.Error())
	//}
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

	case commandAttacking:
		d := c.data.(uniter)
		switch d.GetUnitType() {
		case "player1":
			s.play("atk_sword.mp3")
		case "player2":
			s.play("atk_hammer.mp3")
		case "player3":
			s.play("atk_fire.mp3")
		case "enemy1":
			s.play("enemy_atk.mp3")
		case "enemy2":
			s.play("enemy_atk.mp3")
		}
	}
}
