package main

import (
	"net/http"
	"encoding/json"
	"game"
	"board"
	"constants"
	"player"
	"strconv"
)

type PositionStruct struct {
	Position string
	Ia       string
	Board    board.Board
}

func main() {
	http.HandleFunc("/new", newGame)
	http.HandleFunc("/move", awale)
	http.ListenAndServe(":8080", nil)
}

func newGame(w http.ResponseWriter, r *http.Request) {
	playerOne := player.New(0, true, constants.PIT_COUNT)
	playerTwo := player.New(1, false, constants.PIT_COUNT)
	currentGame := game.New([]player.Player{playerOne, playerTwo})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentGame)
}

func awale(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var positionStruct PositionStruct
	err := decoder.Decode(&positionStruct)
	check(err)
	defer r.Body.Close()

	position, _ := strconv.Atoi(positionStruct.Position)
	isIA, _ := strconv.Atoi(positionStruct.Ia)

	playerOne := player.New(0, true, constants.PIT_COUNT)
	playerTwo := player.New(1, false, constants.PIT_COUNT)
	currentGame := game.New([]player.Player{playerOne, playerTwo})
	currentGame.Board = positionStruct.Board
	currentGame.CurrentPlayerIndex = isIA

	currentGame = game.PlayTurn(currentGame, game.ConvertPlayerPosition(position, currentGame.CurrentPlayerIndex))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentGame)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
