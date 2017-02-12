package board

import (
	"constants"
	"errors"
	"player"
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

	if IsStarving(board, player.MinPick, player.MaxPick) {
		isStarving := WillStarvePlayer(player, board, position)
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

	var newScore [2]int
	copy(newScore[:], score[:])

	for IsPickPossible(newBoard, player.MinPick, player.MaxPick, endPosition) {
		newScore[player.Number] += newBoard[endPosition]
		newBoard[endPosition] = 0
		endPosition -= 1
	}

	return newScore, newBoard
}

func IsPickPossible(board Board, minPick int, maxPick int, position int) bool {
	return minPick <= position && position < maxPick &&
		2 <= board[position] && board[position] <= 3
}

func WillStarvePlayer(player player.Player, board Board, position int) bool {
	//  Fake pick to simulate next turn
	_, newBoard := Pick(player, board, position, [2]int{0, 0})
	return IsStarving(newBoard, player.MinPick, player.MaxPick)
}

func IsStarving(board Board, minPick int, maxPick int) bool {
	return (SumArray(board[minPick:maxPick]) == 0)
}

func CanFeedPlayer(player player.Player, board Board) bool {
	for i := player.MinPosition; i < player.MaxPosition; i++ {
		starving := WillStarvePlayer(player, board, i)
		if !starving {
			return true
		}
	}
	return false
}

func GetWinner(player player.Player, board Board, score [2]int) int {
	minScore := ((constants.PIT_COUNT * constants.PEBBLE_COUNT) / 2)
	starving := IsStarving(board, player.MinPick, player.MaxPick)
	if starving || score[player.Number] > minScore {
		return player.Number
	} else if score[1-player.Number] > minScore {
		return 1 - player.Number
	}

	return constants.GAME_CONTINUE
}

func GetCountPitWithOneTwoPebble(board Board, player player.Player) int {
	countPeeeble := 0
	for position := player.MinPosition; position < player.MaxPosition; position++ {
		if (board[position] == 1) || (board[position] == 2) {
			countPeeeble += board[position]
		}
	}
	return countPeeeble
}

func IsAPitWithMoreThanTwelvePebble(board Board, player player.Player) int {
	for position := player.MinPosition; position < player.MaxPosition; position++ {
		if board[position] >= 12 {
			return 1
		}
	}
	return 0
}

func SumArray(array []int) int {
	total := 0
	for _, value := range array {
		total += value
	}
	return total
}

func InitBoardTest(b [12]int) Board {
	return Board{b[6], b[7], b[8], b[9], b[10], b[11], b[5], b[4], b[3], b[2], b[1], b[0]}
}
