package game

import (
	"../board"
	"../player"
	"../render"
)

const PIT_COUNT = 12
const PEBBLE_COUNT = 4

type Game struct {
	Board              board.Board
	Players            []player.Player
	CurrentPlayerIndex int
}

func New(players []player.Player) Game {
	gameBoard, _ := board.New(PIT_COUNT, PEBBLE_COUNT)
	return Game{
		gameBoard,
		players,
		0,
	}
}

func Render(game Game) string {
	return render.Render(game.Board, game.CurrentPlayerIndex)
}
