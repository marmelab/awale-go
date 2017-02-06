package board

import (
	"errors"
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
