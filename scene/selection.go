package scene

import "github.com/pankona/gomo-simra/simra"

type selection struct {
	selecting *simra.Sprite
	class     string
	game      *game
}

func (s *selection) OnEvent(i interface{}) {
	c, ok := i.(*command)
	if !ok {
		panic("unexpected command received. fatal.")
	}

	switch c.commandtype {
	case commandSpawned:
		if s.game.gameState == gameStateInitial {
			// this is first ally's summoning.
			// transition gameState from initial to running.
			s.game.eventqueue <- newCommand(commandGoToRunningState, s.game)
		}

	case commandUpdateSelection:
		s.selecting = c.data.(*simra.Sprite)
		simra.LogDebug("selection updated: %v", s.selecting)

	case commandUnsetSelection:
		s.selecting = nil
		simra.LogDebug("selection updated: nil")

	default:
		// nop
	}
}
