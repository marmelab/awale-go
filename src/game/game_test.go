package game

import (
	"board"
	"constants"
	"player"
	"reflect"
	"testing"
)

func TestNewGameReturn0InCurrentPlayerIndex(t *testing.T) {
	game := New([]player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, true, constants.PIT_COUNT),
	})

	if game.CurrentPlayerIndex != 0 {
		t.Error("New Game doesn't return expected CurrentPlayerIndex")
	}
}

func TestConvertPlayerPositionFromPlayerOne(t *testing.T) {
	for i := 1; i <= 6; i++ {
		position := ConvertPlayerPosition(i, 0)
		if position != (i - 1) {
			t.Error("Convert player position doesn't return", position)
		}
	}
}

func TestConvertPlayerPositionFromPlayerTwo(t *testing.T) {
	for i := 1; i <= 6; i++ {
		position := ConvertPlayerPosition(i, 1)
		if position != (constants.PIT_COUNT - i) {
			t.Error("Convert player position doesn't return", position)
		}
	}
}

func TestPlayTurnForNewGameShouldReturnSameResult(t *testing.T) {
	game := New([]player.Player{player.New(0, true, constants.PIT_COUNT), player.New(1, true, constants.PIT_COUNT)})
	expectedGame := New([]player.Player{player.New(0, true, constants.PIT_COUNT), player.New(1, true, constants.PIT_COUNT)})
	expectedGame.Board = board.Board{4, 0, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4}

	currentGame := PlayTurn(game, 1)
	if !reflect.DeepEqual(currentGame, expectedGame) {
		t.Error("New Game doesn't return expected Game")
	}
}
