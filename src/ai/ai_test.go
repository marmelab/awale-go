package ai

import (
	"board"
	"constants"
	"player"
	"testing"
	"time"
)

func TestRandomBoard(t *testing.T) {

	board := board.Board{0, 1, 8, 0, 7, 1, 3, 8, 8, 7, 2, 1}
	players := []player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, false, constants.PIT_COUNT),
	}
	AI_REFLECTION_TIME := time.Millisecond * 1500

	position := GetBestPosition(board, players, 1, AI_REFLECTION_TIME)

	if position == 0 {
		t.Error("New position doesn't return 0")
	}
}

func TestRandom1Board(t *testing.T) {

	board := board.Board{1, 1, 9, 0, 1, 8, 7, 7, 7, 1, 0, 1}
	players := []player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, false, constants.PIT_COUNT),
	}
	AI_REFLECTION_TIME := time.Millisecond * 1500

	position := GetBestPosition(board, players, 1, AI_REFLECTION_TIME)

	if position == 0 {
		t.Error("New position doesn't return 0")
	}
}
