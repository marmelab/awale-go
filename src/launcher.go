package main

import (
	"ai"
	"constants"
	"fmt"
	"game"
	"os"
	"os/signal"
	"player"
	"syscall"
)

func ExitGame() {
	fmt.Println("\n\nNo Winner !\n")
}

func main() {

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		ExitGame()
		os.Exit(0)
	}()

	fmt.Println("\n############# GAME STARTED #############")

	playerOne := askForPlayer("\n Player 1 \n", 0)
	playerTwo := askForPlayer("\n Player 2 \n", 1)

	currentGame := game.New([]player.Player{playerOne, playerTwo})
	var position int

	fmt.Println(game.Render(currentGame))

	for !game.IsFinished(currentGame) {

		currentPlayer := game.GetCurrentPlayer(currentGame)

		if currentPlayer.HumanPlayer {
			position = game.GetPosition(currentGame)
		} else {
			position = ai.GetPosition(currentGame)
			fmt.Println("Player (2), position: ", game.ConvertIAPositionToBoardPosition(position, currentGame.CurrentPlayerIndex))
		}

		currentGame = game.PlayTurn(currentGame, position)
		currentGame = game.CheckWinner(currentGame)
		currentGame = game.SwitchPlayer(currentGame)

		fmt.Println(game.Render(currentGame))
		fmt.Println(game.RenderScore(currentGame))
	}

	fmt.Println(RenderGameState(currentGame))
}

func askForPlayer(header string, indexPlayer int) player.Player {

	var isHuman string

	fmt.Println(header)
	fmt.Print("Are you an human ? (y/n): ")
	fmt.Scanf("%s", &isHuman)

	if isHuman == "y" || isHuman == "" {
		return player.New(indexPlayer, true, constants.PIT_COUNT)
	}
	return player.New(indexPlayer, false, constants.PIT_COUNT)
}

func RenderGameState(game game.Game) string {
	if game.GameState == -1 {
		return "\nNo winner, try again."
	}
	return fmt.Sprintf("Winner player: %d.", (game.GameState + 1))
}
