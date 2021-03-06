package main

import (
	"net/http"
	"encoding/json"
	"game"
	"board"
	"constants"
	"player"
	"ai"
	"strconv"
	"time"
	"fmt"
)

type PositionStruct struct {
	Position string
	Board    board.Board
}

type BoardStruct struct {
	Board    board.Board
	Score    [2]int
}

type AwaleStruct struct {
	Player game.Game
	IA     game.Game
}

func main() {
	http.HandleFunc("/new", newGame)
	http.HandleFunc("/move", awale)
	http.HandleFunc("/moveIA", awaleIA)
	http.ListenAndServe(":2000", nil)
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

	playerOne := player.New(0, true, constants.PIT_COUNT)
	playerTwo := player.New(1, false, constants.PIT_COUNT)
	currentGame := game.New([]player.Player{playerOne, playerTwo})
	currentGame.Board = positionStruct.Board
	currentGame.CurrentPlayerIndex = 0

	gamePlayer := game.PlayTurn(currentGame, game.ConvertPlayerPosition(position, currentGame.CurrentPlayerIndex))
	gamePlayer = game.CheckWinner(gamePlayer)
	gamePlayer = game.SwitchPlayer(gamePlayer)

	position, _ = ai.GetPosition(gamePlayer)
	gameIA := game.PlayTurn(gamePlayer, position)
	gameIA = game.CheckWinner(gameIA)
	gameIA = game.SwitchPlayer(gameIA)

	var awale AwaleStruct
	awale.Player = gamePlayer
	awale.IA = gameIA

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(awale)
}

func awaleIA(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var boardStruct BoardStruct
	err := decoder.Decode(&boardStruct)
	check(err)
	defer r.Body.Close()

	playerOne := player.New(0, true, constants.PIT_COUNT)
	playerTwo := player.New(1, false, constants.PIT_COUNT)

	position, _ := ai.GetBestPositionInTime(boardStruct.Board, []player.Player{playerOne, playerTwo}, 1, boardStruct.Score, time.Second)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%d", position)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
