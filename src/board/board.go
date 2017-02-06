package board

import (
	"../constants"
	"../player"
	"errors"
	"fmt"
)

type Board []int

func New(pitNumber int, pebbleNumber int) (Board, error) {
	board := make(Board, pitNumber, pitNumber)

	if pitNumber <= 0 || pitNumber%2 != 0 {
		return board, errors.New("Board size must be even.")
	}

	for i := range board {
		board[i] = pebbleNumber
	}
	return board, nil
}

func CanPlayerPlayPosition(player player.Player, board Board, position int) bool {
	isPlayerCanMove := (player.MinPosition <= position) && (position < player.MaxPosition)
	if !isPlayerCanMove {
		return false
	}

	movePossible := isPlayerCanMove && (board[position] != 0)
	sumPebble := SumArray(board[player.MinPick:player.MaxPick])

	if sumPebble == 0 {
		return movePossible
	}

	return movePossible
}

func DealPosition(board Board, position int) (int, Board) {
	seeds := board[position]
	board[position] = 0
	i := position

	for seeds > 0 {
		i += 1
		if i%constants.PIT_COUNT != position {
			board[i%constants.PIT_COUNT] += 1
			seeds -= 1
		}
	}

	return i % constants.PIT_COUNT, board
}

func Pick(player player.Player, board Board, position int, score [2]int) ([2]int, Board) {
	endPosition, newBoard := DealPosition(board, position)

	for IsPickPossible(newBoard, player, endPosition) {
		score[player.Number] += newBoard[endPosition]
		newBoard[endPosition] = 0
		endPosition -= 1
	}

	return score, newBoard
}

func IsPickPossible(board Board, player player.Player, position int) bool {
	return player.MinPick <= position && position < player.MaxPick &&
		2 <= board[position] && board[position] <= 3
}

func SumArray(array []int) int {
	total := 0
	for _, value := range array {
		total += value
	}
	return total
}
