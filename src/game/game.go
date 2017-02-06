package game

import (
	"../board"
	"../constants"
	"../player"
	"../render"
	"fmt"
)

const GAME_CONTINUE = -2

type Game struct {
	Board              board.Board
	Players            []player.Player
	Score              [2]int
	CurrentPlayerIndex int
	GameState          int
}

func New(players []player.Player) Game {
	gameBoard, _ := board.New(constants.PIT_COUNT, constants.PEBBLE_COUNT)
	var score [2]int
	return Game{
		gameBoard,
		players,
		score,
		0,
		GAME_CONTINUE,
	}
}

func Render(game Game) string {
	return render.Render(game.Board, game.CurrentPlayerIndex)
}

func RenderScore(game Game) string {
	return render.RenderScore(game.Score)
}

func IsFinished(game Game) bool {
	return game.GameState != GAME_CONTINUE
}

func CheckWinner(game Game) Game {
	if game.GameState != GAME_CONTINUE {
		return game
	}

	player := GetCurrentPlayer(game)
	game.GameState = board.GetWinner(player, game.Board, game.Score)
	return game
}

func GetCurrentPlayer(game Game) player.Player {
	return game.Players[game.CurrentPlayerIndex]
}

func PlayTurn(game Game, position int) Game {
	player := GetCurrentPlayer(game)

	isStarving := board.WillStravePlayer(player, game.Board, position, game.Score)
	if isStarving {
		_, newBoard := board.DealPosition(game.Board, position)
		game.Board = newBoard
		return game
	}

	score, newBoard := board.Pick(player, game.Board, position, game.Score)
	game.Board = newBoard
	game.Score = score
	return game
}

func GetPosition(game Game) int {
	position := 0

	fmt.Printf("Player (%d), wich position: ", game.CurrentPlayerIndex+1)
	fmt.Scanf("%d", &position)
	position = ConvertPositionPlayer(position, game.CurrentPlayerIndex)

	player := GetCurrentPlayer(game)

	for !board.CanPlayerPlayPosition(player, game.Board, position) {
		fmt.Printf("Wrong move, play again which position: ")
		fmt.Scanf("%d", &position)
		position = ConvertPositionPlayer(position, game.CurrentPlayerIndex)
	}
	return position
}

func ConvertPositionPlayer(position int, indexPlayer int) int {
	if indexPlayer == 1 {
		return constants.PIT_COUNT - position
	}
	return position - 1
}

func SwitchPlayer(game Game) Game {
	game.CurrentPlayerIndex = 1 - game.CurrentPlayerIndex
	return game
}
