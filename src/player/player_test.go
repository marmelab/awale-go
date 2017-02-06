package player

import (
	"testing"
)

func TestNewPlayerOneReturnZeroInMinPosition(t *testing.T) {
	playerOne := New(0, true, 12)
	if playerOne.MinPosition != 0 {
		t.Error("New Player One doesn't return expected Min position")
	}
}

func TestNewPlayerTwoReturnSixInMinPosition(t *testing.T) {
	playerTwo := New(1, true, 12)
	if playerTwo.MinPosition != 6 {
		t.Error("New Player Two doesn't return expected Min position")
	}
}
