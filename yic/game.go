package yic

import (
	"github.com/GodsBoss/ld46/pkg/engine"
)

func NewGame(storage Storage) *engine.Game {
	lvls := createLevels()
	game := &engine.Game{
		States: map[string]engine.State{
			titleStateID: &title{},
			playingStateID: &playing{
				levels: lvls,
			},
			levelSelectStateID: &levelSelect{
				levels: lvls,
			},
			hiscoreStateID: &hiscore{
				storage: storage,
			},
			gameOverStateID: &gameOver{},
		},
	}
	game.Transition(titleStateID)
	return game
}

type Storage interface {
	Get(key string) (string, bool)

	Set(key string, value string) error

	Clear()
}
