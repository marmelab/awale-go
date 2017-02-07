package game

import (
	"constants"
	"player"
	"testing"
)

func TestNewGameReturn0InCurrentPlayerIndex(t *testing.T) {
	game := New([]player.Player{player.New(0, true, constants.PIT_COUNT), player.New(1, true, constants.PIT_COUNT)})

	if game.CurrentPlayerIndex != 0 {
		t.Error("New Game doesn't return expected CurrentPlayerIndex")
	}
}

func TestConvertPlayerPositionFromPlayerOneShouldReturn0(t *testing.T) {
	position := ConvertPlayerPosition(1, 0)
	if position != 0 {
		t.Error("Convert player position doesn't return 0")
	}
}

func TestConvertPlayerPositionFromPlayerOneShouldReturn5(t *testing.T) {
	position := ConvertPlayerPosition(6, 0)
	if position != 5 {
		t.Error("Convert player position doesn't return 5")
	}
}

func TestConvertPlayerPositionFromPlayerTwoShouldReturn1(t *testing.T) {
	position := ConvertPlayerPosition(11, 1)
	if position != 1 {
		t.Error("Convert player position doesn't return 1")
	}
}

func TestConvertPlayerPositionFromPlayerTwoShouldReturn6(t *testing.T) {
	position := ConvertPlayerPosition(6, 1)
	if position != 6 {
		t.Error("Convert player position doesn't return 6")
	}
}

func TestConvertPlayerPositionFromPlayerTwoShouldReturn9(t *testing.T) {
	position := ConvertPlayerPosition(3, 1)
	if position != 9 {
		t.Error("Convert player position doesn't return 9")
	}
}
