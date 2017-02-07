package board

import (
	"constants"
	"player"
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

func CanPlayerPlayPosition(player player.Player, board Board, position int) bool {
	isPlayerCanMove := (player.MinPosition <= position) && (position < player.MaxPosition)
	if !isPlayerCanMove {
		return false
	}

	movePossible := isPlayerCanMove && (board[position] != 0)

	if IsStarving(board, player) {
		var score [2]int
		isStarving := WillStarvePlayer(player, board, position, score)
		canFeed := CanFeedPlayer(player, board)
		return movePossible && (!isStarving || !canFeed)
	}

	return movePossible
}

func DealPosition(board Board, position int) (int, Board) {
	newBoard := make([]int, len(board))
	copy(newBoard, board)

	seeds := newBoard[position]
	newBoard[position] = 0
	i := position

	for seeds > 0 {
		i += 1
		if i%constants.PIT_COUNT != position {
			newBoard[i%constants.PIT_COUNT] += 1
			seeds -= 1
		}
	}

	return i % constants.PIT_COUNT, newBoard
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

func WillStarvePlayer(player player.Player, board Board, position int, score [2]int) bool {
	//  Fake pick to simulate next turn
	_, newBoard := Pick(player, board, position, score)
	return IsStarving(newBoard, player)
}

func IsStarving(board Board, player player.Player) {
	return (SumArray(board[player.MinPick:player.MaxPick]) == 0)
}

func CanFeedPlayer(player player.Player, board Board) bool {
	cannot_feed := false
	var score [2]int
	for i := player.MinPosition; i <= player.MaxPosition; i++ {
		starving := WillStarvePlayer(player, board, i, score)
		cannot_feed = cannot_feed && starving
	}
	return !cannot_feed
}

func GetWinner(player player.Player, board Board, score [2]int) int {
	minScore := ((constants.PIT_COUNT * constants.PEBBLE_COUNT) / 2)

	if IsStarving(board, player) || score[player.Number] > minScore {
		return player.Number
	} else if score[1-player.Number] > minScore {
		return 1 - player.Number
	}

	return -2
}

func SumArray(array []int) int {
	total := 0
	for _, value := range array {
		total += value
	}
	return total
}
