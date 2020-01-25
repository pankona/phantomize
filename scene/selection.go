package scene

import (
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/gomo-simra/simra/image"
	"github.com/pankona/gomo-simra/simra/simlog"
)

type selection struct {
	simra     simra.Simraer
	selecting simra.Spriter
	cursor    simra.Spriter
	cursorTex *simra.Texture
	class     string
	game      *game
}

func (s *selection) initialize(g *game) {
	s.game = g
	s.cursor = s.simra.NewSprite()
	s.cursorTex = s.simra.NewImageTexture("cursor.png", image.Rect(0, 0, 30, 30))
	s.cursor.SetScale(30, 30)
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
		s.selecting, ok = c.data.(simra.Spriter)
		if !ok {
			// ignore
			break
		}

		simlog.Debugf("selection updated: %v", s.selecting)

		s.cursor.SetPosition(s.selecting.GetPosition().X, s.selecting.GetPosition().Y)
		s.simra.AddSprite(s.cursor)
		s.cursor.ReplaceTexture(s.cursorTex)

	case commandUnsetSelection:
		s.selecting = nil
		simlog.Debug("selection updated: nil")
		s.simra.RemoveSprite(s.cursor)

	default:
		// nop
	}
}
