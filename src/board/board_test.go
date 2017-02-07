package board

import (
	"player"
	"reflect"
	"testing"
)

func TestNewBoardShouldReturnNewBoard(t *testing.T) {

	board, _ := New(12, 4)
	expectedBoard := Board{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}

	if !reflect.DeepEqual(board, expectedBoard) {
		t.Error("New Board doesn't return expected Board")
	}
}

func TestNewBoardShouldReturnSizeErrors(t *testing.T) {

	_, errors := New(13, 4)

	if errors == nil {
		t.Error("New Board should return an error with an invalid board size")
	}
}

func TestIsPickPossibleForPostionZeroShouldReturnPickPossible(t *testing.T) {
	board, _ := New(12, 4)
	board[0] = 2
	playerTwo := player.New(1, true, 12)
	if !IsPickPossible(board, playerTwo, 0) {
		t.Error("Pick possible should return true for player two")
	}
}

func TestIsPickPossibleForPostionZeroShouldReturnPickImpossible(t *testing.T) {
	board, _ := New(12, 4)
	board[0] = 2
	playerOne := player.New(0, true, 12)
	if IsPickPossible(board, playerOne, 0) {
		t.Error("Pick possible should return false for player one")
	}
}

func TestGetWinnerForMaxScoreShouldReturnPlayerOne(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)
	var score [2]int
	score[0] = 48
	score[1] = 0
	if GetWinner(playerOne, board, score) != 0 {
		t.Error("Get Winner should return player one")
	}
}

func TestGetWinnerForZeroScoreShouldReturnPlayerOne(t *testing.T) {
	board, _ := New(12, 4)
	playerTwo := player.New(1, true, 12)
	var score [2]int
	score[0] = 48
	score[1] = 0
	if GetWinner(playerTwo, board, score) != 0 {
		t.Error("Get Winner should return player one")
	}
}

func TestGetWinnerForContinueGameShouldReturnNoWinner(t *testing.T) {
	board, _ := New(12, 4)
	playerTwo := player.New(1, true, 12)
	var score [2]int
	score[0] = 20
	score[1] = 10
	if GetWinner(playerTwo, board, score) != -2 {
		t.Error("Get Winner should return continue game")
	}
}

func TestIsStarvingForEmptySideShouldReturnTrue(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)

	for i := 6; i <= 11; i++ {
		board[i] = 0
	}

	if !IsStarving(board, playerOne) {
		t.Error("Is starving should return starving")
	}
}

func TestIsStarvingForNewBoardShouldReturnFalse(t *testing.T) {
	board, _ := New(12, 4)
	playerTwo := player.New(1, true, 12)

	if IsStarving(board, playerTwo) {
		t.Error("Is starving should return no starving")
	}
}

func TestIsStarvingForEmptyBoardShouldReturnTrue(t *testing.T) {
	board, _ := New(12, 0)
	playerTwo := player.New(1, true, 12)

	if !IsStarving(board, playerTwo) {
		t.Error("Is starving should return starving")
	}
}

func TestCanFeedPlayerForNewBoardSouldReturnTrue(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)

	if !CanFeedPlayer(playerOne, board) {
		t.Error("Can feed player should return true")
	}
}

func TestCanFeedPlayerForEmptyPlayerTwoBoardSouldReturnTrue(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)

	for i := 6; i <= 11; i++ {
		board[i] = 0
	}

	if !CanFeedPlayer(playerOne, board) {
		t.Error("Can feed player should return true")
	}
}

func TestCanFeedPlayerForRandomPositionSouldReturnFalse(t *testing.T) {
	board, _ := New(12, 0)
	playerOne := player.New(0, true, 12)

	board[0] = 5
	board[1] = 0
	board[2] = 2
	board[3] = 0
	board[4] = 1
	board[5] = 0

	if CanFeedPlayer(playerOne, board) {
		t.Error("Can feed player should return false")
	}
}
