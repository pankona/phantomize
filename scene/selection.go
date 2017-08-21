package scene

import (
	"image"

	"github.com/pankona/gomo-simra/simra"
)

type selection struct {
	selecting *simra.Sprite
	cursor    *simra.Sprite
	cursorTex *simra.Texture
	class     string
	game      *game
}

func (s *selection) initialize(g *game) {
	s.game = g
	s.cursor = simra.NewSprite()
	s.cursorTex = simra.NewImageTexture("cursor.png", image.Rect(0, 0, 30, 30))
	s.cursor.W, s.cursor.H = 30, 30
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
		s.selecting, ok = c.data.(*simra.Sprite)
		if !ok {
			// ignore
			break
		}

		simra.LogDebug("selection updated: %v", s.selecting)

		s.cursor.X, s.cursor.Y = s.selecting.X, s.selecting.Y
		simra.GetInstance().AddSprite(s.cursor)
		s.cursor.ReplaceTexture(s.cursorTex)

	case commandUnsetSelection:
		s.selecting = nil
		simra.LogDebug("selection updated: nil")
		simra.GetInstance().RemoveSprite(s.cursor)

	default:
		// nop
	}
}
