package main

import (
	"net/http"
	"net/url"
	"encoding/json"
	"game"
	"board"
	"constants"
	"player"
	"strconv"
)

type PositionStruct struct {
    Position string
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
	json.NewEncoder(w).Encode(currentGame)
}

func awale(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	board, err := getBoardFromQueryString(query)
	check(err)

	isIA, err := getIsIAFromQueryString(query)
	check(err)

	decoder := json.NewDecoder(r.Body)
	var positionStruct PositionStruct
	err = decoder.Decode(&positionStruct)
	check(err)
	defer r.Body.Close()

	position, _ := strconv.Atoi(positionStruct.Position)

	playerOne := player.New(0, true, constants.PIT_COUNT)
	playerTwo := player.New(1, false, constants.PIT_COUNT)
	currentGame := game.New([]player.Player{playerOne, playerTwo})
	currentGame.Board = board
	currentGame.CurrentPlayerIndex = isIA

	currentGame = game.PlayTurn(currentGame, game.ConvertPlayerPosition(position, currentGame.CurrentPlayerIndex))

	json.NewEncoder(w).Encode(currentGame)
}

func getBoardFromQueryString(query url.Values) (board.Board, error) {
	queryString := query.Get("board")
	grid := []byte(queryString)

	var board board.Board
	err := json.Unmarshal(grid, &board)
	return board, err
}

func getIsIAFromQueryString(query url.Values) (int, error) {
	queryString := query.Get("ia")
	return strconv.Atoi(queryString)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
