package ai

import (
	"board"
	"errors"
	"game"
	"player"
	"time"
)

const AI_REFLECTION_TIME time.Duration = time.Millisecond * 1500
const SCORING_WORKER_COUNT int = 4
const DEPTH_RECURSIVITY int = 4

type Node struct {
	Board              board.Board
	PositionChange     int
	RootPositionChange int
	IsOpponent         bool
	Players            []player.Player
	IndexCurrentPlayer int
	Depth              int
}

type Scoring struct {
	ScoreNode   Node
	ScoringTime time.Duration
	Score       int
}

func GetPosition(currentGame game.Game) (int, error) {
	return GetBestPosition(currentGame.Board, currentGame.Players, currentGame.CurrentPlayerIndex, AI_REFLECTION_TIME)
}

func GetBestPosition(currentBoard board.Board, players []player.Player, indexCurrentPlayer int, duration time.Duration) (int, error) {

	nodes := make(chan Node, 100)
	scores := make(chan Scoring)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(duration)
		timeout <- true
	}()

	player := players[indexCurrentPlayer]
	legalPositionChanges := GetLegalPositionChangesForPlayer(player, currentBoard)

	if len(legalPositionChanges) == 0 {
		return -1, errors.New("There's no legal position for this player.")
	}

	if len(legalPositionChanges) == 1 {
		return legalPositionChanges[0], nil
	}

	// Start scoring workers
	for i := 0; i < SCORING_WORKER_COUNT; i++ {
		go ScoringWorker(nodes, scores)
	}

	// Start board graph visitors
	for _, position := range legalPositionChanges {
		go RecursiveNodeVisitor(Node{currentBoard, position, position, false, players, indexCurrentPlayer, DEPTH_RECURSIVITY}, nodes)
	}

	return CaptureBestPositionChange(scores, timeout), nil
}

func CaptureBestPositionChange(scores chan Scoring, stopProcess chan bool) int {

	bestPositionChange := 0
	finished := false
	maxScore := 0

	for !finished {
		select {
		case finished = <-stopProcess:
		case scoring := <-scores:
			if scoring.Score > maxScore && bestPositionChange != scoring.ScoreNode.RootPositionChange {
				maxScore = scoring.Score
				bestPositionChange = scoring.ScoreNode.RootPositionChange
			}
		}
	}

	return bestPositionChange
}

func ScoringWorker(nodes <-chan Node, scores chan<- Scoring) {
	for node := range nodes {
		start := time.Now()
		score := Score(node)
		scores <- Scoring{node, time.Since(start), score}
	}
}

func Score(node Node) int {

	// Bad for player
	playerCountPebble := board.GetCountPitWithOneTwoPebble(node.Board, node.Players[node.IndexCurrentPlayer])
	opponentWithFullOfPebble := board.IsPitWithMoreTwelvePebble(node.Board, node.Players[1-node.IndexCurrentPlayer])

	// Good for player
	opponentCountPebble := board.GetCountPitWithOneTwoPebble(node.Board, node.Players[1-node.IndexCurrentPlayer])
	playerWithFullOfPebble := board.IsPitWithMoreTwelvePebble(node.Board, node.Players[node.IndexCurrentPlayer])

	scoreGoodForPlayer := (0.5 * float64(opponentCountPebble)) + (0.2 * float64(playerWithFullOfPebble))
	scoreBadForPlayer := -(0.5 * float64(playerCountPebble)) - (0.2 * float64(opponentWithFullOfPebble))
	totalScore := scoreGoodForPlayer + scoreBadForPlayer

	if node.IsOpponent {
		return -int(totalScore)
	}
	return int(totalScore)
}

func RecursiveNodeVisitor(rootNode Node, out chan Node) {
	for _, node := range NodeVisitor(rootNode) {
		out <- node
		defer func() { go RecursiveNodeVisitor(node, out) }()
	}
}

func NodeVisitor(node Node) []Node {
	player := node.Players[node.IndexCurrentPlayer]
	nodeChild := []Node{}

	if node.Depth == 0 {
		return nodeChild
	}

	legalPositionChanges := GetLegalPositionChangesForPlayer(player, node.Board)
	for _, positionChange := range legalPositionChanges {
		_, nodeBoard := board.DealPosition(node.Board, positionChange)
		nodeChild = append(nodeChild, Node{nodeBoard, positionChange, node.RootPositionChange, !node.IsOpponent, node.Players, 1 - node.IndexCurrentPlayer, node.Depth - 1})
	}

	return nodeChild
}

func GetLegalPositionChangesForPlayer(player player.Player, currentBoard board.Board) []int {
	var legalPositionChanges []int
	for position := player.MinPosition; position < player.MaxPosition; position++ {
		if board.CanPlayerPlayPosition(player, currentBoard, position) {
			legalPositionChanges = append(legalPositionChanges, position)
		}
	}
	return legalPositionChanges
}
