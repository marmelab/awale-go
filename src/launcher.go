package main

import (
	"./game"
	"./player"
	"fmt"
)

func main() {

	fmt.Println("\n############# GAME STARTED #############")

	playerOne := player.New(0, true, game.PIT_COUNT)
	playerTwo := player.New(1, true, game.PIT_COUNT)

	party := game.New([]player.Player{playerOne, playerTwo})

	fmt.Println(game.Render(party))
}
