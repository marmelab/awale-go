package scoring

import (
	"board"
	"player"
)

func GetCountPitWithOneTwoPebble(board board.Board, player player.Player) int {
	countPeeeble := 0
	for position := player.MinPosition; position < player.MaxPosition; position++ {
		if (board[position] == 1) || (board[position] == 2) {
			countPeeeble += board[position]
		}
	}
	return countPeeeble
}

func IsPitWhithMoreTwelvePebbble(board board.Board, player player.Player) int {
	for position := player.MinPosition; position < player.MaxPosition; position++ {
		if board[position] >= 12 {
			return 1
		}
	}
	return 0
}
