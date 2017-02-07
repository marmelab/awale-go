package main

import (
	"constants"
	"fmt"
	"game"
	"player"
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
		currentGame = game.CheckWinner(currentGame)
		currentGame = game.SwitchPlayer(currentGame)

		fmt.Println(game.Render(currentGame))
		fmt.Println(game.RenderScore(currentGame))
	}

	fmt.Println(RenderGameState(currentGame))
}

func RenderGameState(game game.Game) string {
	if game.GameState == -1 {
		return "\nNo winner, try again."
	}
	return fmt.Sprintf("Winner player: %d.", (game.GameState + 1))
}
