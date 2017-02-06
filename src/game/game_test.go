package game

import (
	"../constants"
	"../player"
	"testing"
)

func TestNewGameReturn0InCurrentPlayerIndex(t *testing.T) {
	game := New([]player.Player{player.New(0, true, constants.PIT_COUNT), player.New(1, true, constants.PIT_COUNT)})

	if game.CurrentPlayerIndex != 0 {
		t.Error("New Game doesn't return expected CurrentPlayerIndex")
	}
}
