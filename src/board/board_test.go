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

func TestIsPickPossibleShouldReturnTrueForPositionZero(t *testing.T) {
	board, _ := New(12, 4)
	board[0] = 2
	if !IsPickPossible(board, 0, 5, 0) {
		t.Error("Pick possible should return true for player two")
	}
}

func TestIsPickPossibleShouldReturnFalseForPositionZero(t *testing.T) {
	board, _ := New(12, 4)
	board[0] = 2
	if IsPickPossible(board, 6, 11, 0) {
		t.Error("Pick possible should return false for player one")
	}
}

func TestGetWinnerForMaxScoreShouldReturnPlayerOne(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)
	score := [2]int{48, 0}
	if GetWinner(playerOne, board, score) != 0 {
		t.Error("Get Winner should return player one")
	}
}

func TestGetWinnerForZeroScoreShouldReturnPlayerOneIndex(t *testing.T) {
	board, _ := New(12, 4)
	playerTwo := player.New(1, true, 12)
	score := [2]int{48, 0}
	if GetWinner(playerTwo, board, score) != 0 {
		t.Error("Get Winner should return player one")
	}
}

func TestGetWinnerForContinueGameShouldReturnNoWinner(t *testing.T) {
	board, _ := New(12, 4)
	playerTwo := player.New(1, true, 12)
	score := [2]int{20, 10}
	if GetWinner(playerTwo, board, score) != -2 {
		t.Error("Get Winner should return continue game")
	}
}

func TestIsStarvingForEmptySideShouldReturnTrue(t *testing.T) {
	board, _ := New(12, 4)

	for i := 6; i <= 11; i++ {
		board[i] = 0
	}

	if !IsStarving(board, 6, 11) {
		t.Error("Is starving should return starving")
	}
}

func TestIsStarvingForNewBoardShouldReturnFalse(t *testing.T) {
	board, _ := New(12, 4)

	if IsStarving(board, 0, 5) {
		t.Error("Is starving should return no starving")
	}
}

func TestIsStarvingForEmptyBoardShouldReturnTrue(t *testing.T) {
	board, _ := New(12, 0)

	if !IsStarving(board, 0, 5) {
		t.Error("Is starving should return starving")
	}
}

func TestCanFeedPlayerForNewBoardShouldReturnTrue(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)

	if !CanFeedPlayer(playerOne, board) {
		t.Error("Can feed player should return true")
	}
}

func TestPlayerCanFeedReversePlayerWithAnEmptyBoard(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)

	for i := 6; i <= 11; i++ {
		board[i] = 0
	}

	if !CanFeedPlayer(playerOne, board) {
		t.Error("Can feed player should return true")
	}
}

func TestCanFeedPlayerForEmptySidePositionShouldReturnFalse(t *testing.T) {
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

func TestDealPositionForNewBoardShouldReturnFour(t *testing.T) {
	board, _ := New(12, 4)
	endPosition, newBoard := DealPosition(board, 0)

	if !reflect.DeepEqual(newBoard, Board{0, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 4}) {
		t.Error("New Board doesn't return expected Board")
	}

	if endPosition != 4 {
		t.Error("End position should return 4")
	}
}

func TestDealPositionForEmptySideBoardShouldReturnNine(t *testing.T) {
	board := Board{4, 4, 4, 4, 4, 4, 0, 0, 0, 0, 0, 0}
	endPosition, newBoard := DealPosition(board, 5)

	if !reflect.DeepEqual(newBoard, Board{4, 4, 4, 4, 4, 0, 1, 1, 1, 1, 0, 0}) {
		t.Error("New Board doesn't return expected Board")
	}

	if endPosition != 9 {
		t.Error("End position should return 9")
	}
}

func TestDealPositionFor13PebbleBoardShouldReturnSeven(t *testing.T) {
	board := Board{4, 4, 4, 4, 14, 4, 0, 0, 0, 0, 0, 0}
	endPosition, newBoard := DealPosition(board, 4)

	if !reflect.DeepEqual(newBoard, Board{5, 5, 5, 5, 0, 6, 2, 2, 1, 1, 1, 1}) {
		t.Error("New Board doesn't return expected Board")
	}

	if endPosition != 7 {
		t.Error("End position should return 7")
	}
}

func TestPickForNewBoardShouldReturnEmptyScore(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)
	score, _ := Pick(playerOne, board, 6, [2]int{0, 0})

	if !reflect.DeepEqual(score, [2]int{0, 0}) {
		t.Error("New score doesn't return expected score")
	}
}

func TestPickForEmptySideShouldReturnEmptyScore(t *testing.T) {
	board := Board{4, 4, 4, 4, 14, 4, 0, 0, 0, 0, 0, 0}
	playerOne := player.New(0, true, 12)
	score, _ := Pick(playerOne, board, 6, [2]int{0, 0})

	if !reflect.DeepEqual(score, [2]int{0, 0}) {
		t.Error("New score doesn't return expected score")
	}
}

func TestPickForPositionWithEmptyPitShouldReturn7Score(t *testing.T) {
	board := Board{4, 4, 4, 4, 4, 4, 0, 1, 2, 1, 4, 5}
	playerOne := player.New(0, true, 12)
	score, _ := Pick(playerOne, board, 5, [2]int{0, 0})

	if !reflect.DeepEqual(score, [2]int{7, 0}) {
		t.Error("New score doesn't return expected score")
	}
}

func TestCanPlayerPlayPositionForNewBoardReturnTrue(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)
	if !CanPlayerPlayPosition(playerOne, board, 0) {
		t.Error("Can player play position doesn't return true")
	}
}

func TestCanPlayerPlayPositionForWrongPositionReturnFalse(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)
	if CanPlayerPlayPosition(playerOne, board, 99) {
		t.Error("Can player play position doesn't return false")
	}
}

func TestCountPitWithOneTwoPebbleForNewBoardShoulReturn0(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)
	if GetCountPitWithOneTwoPebble(board, playerOne) != 0 {
		t.Error("Count Pebble doesn't return 0")
	}
}

func TestCountPitWithOneTwoPebbleForAllBoardWithOnePebbleShoulReturn6(t *testing.T) {
	board := Board{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	playerOne := player.New(0, true, 12)
	if GetCountPitWithOneTwoPebble(board, playerOne) != 6 {
		t.Error("Count Pebble doesn't return 6")
	}
}

func TestIsAPitWithMoreThanTwelvePebbleForNewBoardShoulReturnFalse(t *testing.T) {
	board, _ := New(12, 4)
	playerOne := player.New(0, true, 12)
	if IsAPitWithMoreThanTwelvePebble(board, playerOne) != 0 {
		t.Error("Pit with more than twelve pebble doesn't return False")
	}
}

func TestIsAPitWithMoreThanTwelvePebbleForNewBoardShoulReturnTrue(t *testing.T) {
	board := Board{1, 1, 0, 13, 1, 1, 1, 1, 1, 1, 1, 1}
	playerOne := player.New(0, true, 12)
	if IsAPitWithMoreThanTwelvePebble(board, playerOne) != 1 {
		t.Error("Pit with more than twelve pebble doesn't return True")
	}
}

func TestSumArrayShouldRetun10(t *testing.T) {
	array := make([]int, 2, 2)
	array[0] = 5
	array[1] = 5
	if SumArray(array) != 10 {
		t.Error("Sum array doesn't return expected sum 10")
	}
}
