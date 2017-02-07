package ai

import (
	"game"
	"board"
    "fmt"
)

func GetPosition(currentGame game.Game) (int) {

    depth := 6

    _, bestPosition := MinMax(currentGame, depth)
    return bestPosition
}

func MinMax(currentGame game.Game, depth int) (int, int) {

    player := game.GetCurrentPlayer(currentGame)

    if game.IsFinished(currentGame) || depth == 0 {
       // todo do something else to evaluation
        return currentGame.Score[player.Number] - currentGame.Score[1 - player.Number], 6 * player.Number
    }

    bestScore := 0
    bestPosition := 0

    // todo add go routine

    if player.Number == 0 {
        bestScore = -999
        for position := player.MinPosition; position < player.MaxPosition; position++ {
            if board.CanPlayerPlayPosition(player, currentGame.Board, position) {
               score :=  MinMax(game, depth - 1)
               if score > bestScore {
                   bestScore = score
                   bestPosition = position
               }
            }
        }
    }
    else {
        bestScore = 999
        for position := player.MinPosition; position < player.MaxPosition; position++ {
            if board.CanPlayerPlayPosition(player, currentGame.Board, position) {
                score :=  MinMax(game, depth - 1)
                if score < bestScore {
                    bestScore = score
                    bestPosition = position
                }
            }
        }
    }    

    return bestScore, bestPosition
}