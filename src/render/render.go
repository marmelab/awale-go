package render

import (
	"../board"
	"fmt"
)

func Render(board board.Board, playerNumber int) string {
	var render string = "\n"
	var size int = len(board)
	var halfSize int = size / 2

	render += "                   "
	for i := 1; i <= halfSize; i++ {
		render += fmt.Sprintf("%2d    ", i)
	}
	render += "\n\n"

	render += IndicatorCurrentPlayer(playerNumber, 1)
	render += " Player 2    "
	for i := size - 1; i >= halfSize; i-- {
		render += fmt.Sprintf("%2d    ", board[i])
	}

	render += "\n"

	render += IndicatorCurrentPlayer(playerNumber, 0)
	render += " Player 1    "
	for _, row := range board[0:halfSize] {
		render += fmt.Sprintf("%2d    ", row)
	}

	render += "\n"

	return render
}

func IndicatorCurrentPlayer(playerNumber int, number int) string {
	if playerNumber == number {
		return "  =>  "
	}
	return "      "
}

func RenderScore(score [2]int) string {
	return fmt.Sprintf("Score:\tPlayer (1): %d\tPlayer (2): %d\n", score[0], score[1])
}
