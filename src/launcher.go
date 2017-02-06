package main

import (
	"./constants"
	"./game"
	"./player"
	"fmt"
)

func main() {

	fmt.Println("\n############# GAME STARTED #############")

	playerOne := player.New(0, true, constants.PIT_COUNT)
	playerTwo := player.New(1, true, constants.PIT_COUNT)

	currentGame := game.New([]player.Player{playerOne, playerTwo})
	var position int

	fmt.Println(game.Render(currentGame))

	for !game.IsFinished(currentGame) {

		currentPlayer := game.GetCurrentPlayer(currentGame)

		if currentPlayer.HumanPlayer {
			position = game.GetPosition(currentGame)
		}

		currentGame = game.PlayTurn(currentGame, position)
		currentGame = game.SwitchPlayer(currentGame)

		fmt.Println(game.Render(currentGame))
		fmt.Println(game.RenderScore(currentGame))
	}
}
