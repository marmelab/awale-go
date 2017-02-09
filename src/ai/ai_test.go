package ai

import (
	"board"
	"constants"
	"player"
	"reflect"
	"testing"
)

func TestLegalPositionChangesForPlayer1ShouldReturn05(t *testing.T) {

	board, _ := board.New(constants.PIT_COUNT, constants.PEBBLE_COUNT)
	playerOne := player.New(0, true, constants.PIT_COUNT)

	position := GetLegalPositionChangesForPlayer(playerOne, board)
	expectedPosition := []int{0, 1, 2, 3, 4, 5}

	if !reflect.DeepEqual(position, expectedPosition) {
		t.Error("New position doesn't expected position")
	}
}

func TestLegalPositionChangesForPlayer2ShouldReturn611(t *testing.T) {

	board, _ := board.New(constants.PIT_COUNT, constants.PEBBLE_COUNT)
	playerOne := player.New(1, true, constants.PIT_COUNT)

	position := GetLegalPositionChangesForPlayer(playerOne, board)
	expectedPosition := []int{6, 7, 8, 9, 10, 11}

	if !reflect.DeepEqual(position, expectedPosition) {
		t.Error("New position doesn't expected position")
	}
}
