package ai

import (
	"board"
	"constants"
	"errors"
	"fmt"
	"game"
	"math"
	"player"
	"render"
	"time"
)

const MINSCORE int = math.MinInt32
const DEPTH int = 10
const INDICE_SCORE_GOOD float64 = 0.5
const INDICE_SCORE_BAD float64 = 0.2

type ScoredBoard struct {
	CurrentBoard   board.Board
	CurrentScoring int
}

type BestMove struct {
	Position int
	Error    error
}

func DisplayScoreBoard(position int, score ScoredBoard) string {
	return fmt.Sprintf("p: %d score: %d board: %s \n", position, score.CurrentScoring, render.Render(score.CurrentBoard, 1))
}

func GetPosition(currentGame game.Game) (int, error) {
	return GetBestPositionInTime(currentGame.Board, currentGame.Players, currentGame.CurrentPlayerIndex, currentGame.Score, time.Second)
}

func GetBestPositionInTime(currentBoard board.Board, players []player.Player, indexCurrentPlayer int, gameScore [2]int, duration time.Duration) (int, error) {
	results := make(chan BestMove, 1)
	timeout := make(chan bool, 1)
	position := -1
	var error error

	go func() {
		time.Sleep(duration)
		timeout <- true
	}()

	go GetBestPosition(currentBoard, players, indexCurrentPlayer, gameScore, results)

	finished := false

	for !finished {
		select {
		case finished = <-timeout:
		case result, channelStillOpen := <-results:
			if !channelStillOpen {
				finished = true
				break
			}
			position = result.Position
			if result.Error != nil {
				error = result.Error
				break
			}
		}
	}

	if error != nil {
		return position, error
	}

	if position == -1 {
		return position, errors.New("No result found in time")
	}

	return position, nil
}

func GetBestPosition(currentBoard board.Board, players []player.Player, indexCurrentPlayer int, gameScore [2]int, results chan BestMove) {
	currentPlayer := players[indexCurrentPlayer]
	defer close(results)
	var scoredBoards [][]ScoredBoard = make([][]ScoredBoard, constants.PIT_COUNT)
	firstPossibleBoards := 0

	legalPositionChanges := GetLegalPositionChangesForPlayer(currentPlayer, currentBoard)

	for _, position := range legalPositionChanges {
		scoredBoards[position] = make([]ScoredBoard, 0)

		nextGameScore, nextBoard := board.Pick(currentPlayer, currentBoard, position, gameScore)
		if board.GetWinner(currentPlayer, nextBoard, nextGameScore) == indexCurrentPlayer {
			results <- BestMove{
				Position: position,
				Error:    nil,
			}
			return
		}

		scoredBoard := ScoredBoard{
			CurrentBoard:   nextBoard,
			CurrentScoring: Score(nextBoard, players, indexCurrentPlayer, nextGameScore),
		}

		scoredBoards[position] = append(scoredBoards[position], scoredBoard)
		firstPossibleBoards++
	}

	if firstPossibleBoards == 0 {
		results <- BestMove{
			Position: -1,
			Error:    errors.New("No possible move"),
		}
		return
	}

	var bestPosition int
	depth := DEPTH
	for depth > 0 {
		bestScore := MINSCORE

		for _, position := range legalPositionChanges {
			score := AggregateScoring(scoredBoards[position])

			if score > bestScore {
				bestScore = score
				bestPosition = position
			}
		}

		results <- BestMove{
			Position: bestPosition,
			Error:    nil,
		}

		scoredBoards = GuessNextBoardsAggregated(scoredBoards, players, 1-indexCurrentPlayer, indexCurrentPlayer)
		depth--
	}
}

func AggregateScoring(scoredBoards []ScoredBoard) int {
	score := 0
	for _, scoredBoard := range scoredBoards {
		score += scoredBoard.CurrentScoring
	}
	return score
}

func GuessNextBoardsAggregated(scoredBoards [][]ScoredBoard, players []player.Player, indexCurrentPlayer int, scoringPlayer int) [][]ScoredBoard {
	var nextScoredBoardsByPosition [][]ScoredBoard = make([][]ScoredBoard, constants.PIT_COUNT)

	for position := 0; position < constants.PIT_COUNT; position++ {
		if len(scoredBoards) <= position {
			break
		}

		nextScoredBoardsByPosition[position] = make([]ScoredBoard, 0)
		nextScoredBoardsOnePosition := GuessNextBoards(scoredBoards[position], players, indexCurrentPlayer, scoringPlayer)
		nextScoredBoardsByPosition[position] = append(nextScoredBoardsByPosition[position], nextScoredBoardsOnePosition...)
	}
	return nextScoredBoardsByPosition
}

func GuessNextBoards(scoredBoards []ScoredBoard, players []player.Player, indexCurrentPlayer int, scoringPlayer int) []ScoredBoard {
	currentPlayer := players[indexCurrentPlayer]

	var nextScoredBoards []ScoredBoard
	var gameScore [2]int
	for _, scoredBoard := range scoredBoards {
		legalPositionChanges := GetLegalPositionChangesForPlayer(currentPlayer, scoredBoard.CurrentBoard)
		for _, position := range legalPositionChanges {
			nextGameScore, nextBoard := board.Pick(currentPlayer, scoredBoard.CurrentBoard, position, gameScore)

			nextScoredBoard := ScoredBoard{
				CurrentBoard:   nextBoard,
				CurrentScoring: Score(nextBoard, players, scoringPlayer, nextGameScore),
			}
			nextScoredBoards = append(nextScoredBoards, nextScoredBoard)
		}
	}

	return nextScoredBoards
}

func Score(currentBoard board.Board, players []player.Player, indexCurrentPlayer int, gameScore [2]int) int {

	// Bad for player
	playerCountPebble := board.GetCountPitWithOneTwoPebble(currentBoard, players[indexCurrentPlayer])
	opponentWithFullOfPebble := board.IsAPitWithMoreThanTwelvePebble(currentBoard, players[1-indexCurrentPlayer])

	// Good for player
	opponentCountPebble := board.GetCountPitWithOneTwoPebble(currentBoard, players[1-indexCurrentPlayer])
	playerWithFullOfPebble := board.IsAPitWithMoreThanTwelvePebble(currentBoard, players[indexCurrentPlayer])

	score := gameScore[indexCurrentPlayer] - gameScore[1-indexCurrentPlayer]

	scoreGoodForPlayer := +(INDICE_SCORE_GOOD * float64(opponentCountPebble)) + (INDICE_SCORE_BAD * float64(playerWithFullOfPebble))
	scoreBadForPlayer := -(INDICE_SCORE_GOOD * float64(playerCountPebble)) - (INDICE_SCORE_BAD * float64(opponentWithFullOfPebble))
	totalScore := (scoreGoodForPlayer + scoreBadForPlayer + float64(score)) * 1000

	return int(totalScore)
}

func GetLegalPositionChangesForPlayer(player player.Player, currentBoard board.Board) []int {
	var legalPositionChanges []int
	for position := player.MinPosition; position < player.MaxPosition; position++ {
		if board.CanPlayerPlayPosition(player, currentBoard, position) {
			legalPositionChanges = append(legalPositionChanges, position)
		}
	}
	return legalPositionChanges
}
