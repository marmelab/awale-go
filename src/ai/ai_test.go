package ai

import (
	"board"
	"constants"
	"os"
	"player"
	"reflect"
	"testing"
	"time"
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

func TestGuessNextBoards(t *testing.T) {

	players := []player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, true, constants.PIT_COUNT),
	}

	gameBoard := board.InitBoardTest([12]int{
		5, 4, 7, 2, 0, 1,
		1, 5, 3, 4, 0, 6,
	})

	scoredBoards := make([]ScoredBoard, 0)
	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentBoard: gameBoard,
	})

	nextScoredBoards := GuessNextBoards(scoredBoards, players, 1, 1)

	firstBoard := board.InitBoardTest([12]int{
		5, 4, 7, 2, 1, 0,
		1, 5, 3, 4, 0, 6,
	})
	if !reflect.DeepEqual(nextScoredBoards[0].CurrentBoard, firstBoard) {
		t.Error("Expected boards to contain next board playing with first column")
	}

	secondBoard := board.InitBoardTest([12]int{
		5, 5, 8, 0, 0, 1,
		1, 5, 3, 4, 0, 6,
	})
	if !reflect.DeepEqual(nextScoredBoards[1].CurrentBoard, secondBoard) {
		t.Error("Expected boards to contain next board playing with second column")
	}

	thirdBoard := board.InitBoardTest([12]int{
		6, 5, 0, 2, 0, 1,
		2, 6, 4, 5, 1, 6,
	})
	if !reflect.DeepEqual(nextScoredBoards[2].CurrentBoard, thirdBoard) {
		t.Error("Expected boards to contain next board playing with third column")
	}

	fourthBoard := board.InitBoardTest([12]int{
		6, 0, 7, 2, 0, 1,
		2, 6, 4, 4, 0, 6,
	})
	if !reflect.DeepEqual(nextScoredBoards[3].CurrentBoard, fourthBoard) {
		t.Error("Expected boards to contain next board playing with fourth column")
	}

	fifthBoard := board.InitBoardTest([12]int{
		0, 4, 7, 2, 0, 1,
		2, 6, 4, 5, 1, 6,
	})
	if !reflect.DeepEqual(nextScoredBoards[4].CurrentBoard, fifthBoard) {
		t.Error("Expected boards to contain next board playing with fifth column")
	}
}

func TestScoreSecondPlayerShouldBeHigher(t *testing.T) {

	gameScore := [2]int{1, 0}
	players := []player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, true, constants.PIT_COUNT),
	}

	gameBoard := board.InitBoardTest([12]int{
		5, 4, 7, 2, 0, 1,
		1, 5, 4, 4, 3, 4,
	})

	firstPlayerScore := Score(gameBoard, players, 0, gameScore)
	secondPlayerScore := Score(gameBoard, players, 1, gameScore)
	if firstPlayerScore < secondPlayerScore {
		t.Error("Expected first player score to be higher than second player score")
	}
}

func TestGetBestPositionInTimeIsReturned(t *testing.T) {
	gameScore := [2]int{1, 0}
	players := []player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, true, constants.PIT_COUNT),
	}

	gameBoard := board.InitBoardTest([12]int{
		0, 0, 0, 0, 0, 1,
		1, 5, 4, 4, 3, 4,
	})

	position, err := GetBestPositionInTime(gameBoard, players, 0, gameScore, 2*time.Second)

	if err != nil {
		t.Error("Expected to have at least one result in time")
	}
	if position != 1 {
		t.Error("Expected 1, got ", position)
	}
}

func TestGetBestPositionInTimeIsLastPosition(t *testing.T) {
	gameScore := [2]int{22, 24}
	players := []player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, true, constants.PIT_COUNT),
	}

	gameBoard := board.InitBoardTest([12]int{
		0, 0, 5, 0, 0, 1,
		1, 5, 4, 4, 3, 12,
	})

	position, err := GetBestPositionInTime(gameBoard, players, 0, gameScore, 2*time.Second)
	if err != nil {
		t.Error("Expected to have at least one result in time")
	}
	if position != 5 {
		t.Error("Expected 5, got ", position)
	}
}

func TestWhatHappensWhenBoardFull(t *testing.T) {
	gameScore := [2]int{22, 24}
	players := []player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, true, constants.PIT_COUNT),
	}

	gameBoard := board.InitBoardTest([12]int{
		0, 0, 0, 0, 0, 0,
		3, 5, 4, 4, 3, 12,
	})

	_, err := GetBestPositionInTime(gameBoard, players, 1, gameScore, time.Second)

	if err == nil {
		t.Error("Expected not to have an error")
	}
}

func TestGuessNextBoardsAggregated(t *testing.T) {
	t.Skip("skipping test")

	players := []player.Player{
		player.New(0, true, constants.PIT_COUNT),
		player.New(1, true, constants.PIT_COUNT),
	}

	firstBoard := board.InitBoardTest([12]int{
		5, 4, 7, 2, 0, 1,
		1, 5, 3, 4, 0, 6,
	})
	firstScoreBoard := ScoredBoard{
		CurrentBoard:   firstBoard,
		CurrentScoring: 10,
	}

	secondBoard := board.InitBoardTest([12]int{
		6, 5, 8, 3, 1, 2,
		1, 5, 3, 4, 0, 0,
	})
	secondScoredBoards := ScoredBoard{
		CurrentBoard:   secondBoard,
		CurrentScoring: 11,
	}

	var scoredBoards [][]ScoredBoard = make([][]ScoredBoard, 1)
	scoredBoards[0] = make([]ScoredBoard, 0)
	scoredBoards[0] = append(scoredBoards[0], firstScoreBoard)
	scoredBoards[0] = append(scoredBoards[0], secondScoredBoards)

	newScoreBoards := GuessNextBoardsAggregated(scoredBoards, players, 0, 0)

	f, e := os.Create("./GuessNextBoardsAggregated.txt")
	if e != nil {
		panic(e)
	}

	for i, scoredBoards := range newScoreBoards {
		for _, scoredBoard := range scoredBoards {
			f.WriteString(DisplayScoreBoard(i, scoredBoard))
		}
	}
	f.Sync()
}
