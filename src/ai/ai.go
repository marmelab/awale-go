package ai

import (
	"board"
	"game"
	"player"
	"time"
)

const AI_REFLECTION_TIME time.Duration = time.Millisecond * 1500
const SCORING_WORKER_COUNT int = 6

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

func GetPosition(currentGame game.Game) int {
	return GetBestPosition(currentGame.Board, currentGame.Players, currentGame.CurrentPlayerIndex, AI_REFLECTION_TIME)
}

func GetBestPosition(currentBoard board.Board, players []player.Player, indexCurrentPlayer int, duration time.Duration) int {

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
		return -1 // todo return error
	}

	if len(legalPositionChanges) == 1 {
		return legalPositionChanges[0]
	}

	// Start scoring workers
	for i := 0; i < SCORING_WORKER_COUNT; i++ {
		go ScoringWorker(nodes, scores)
	}

	// Start board graph visitors
	for _, position := range legalPositionChanges {
		go RecursiveNodeVisitor(Node{currentBoard, position, position, false, players, indexCurrentPlayer, 1}, nodes)
	}

	return CaptureBestPositionChange(scores, timeout)
}

func CaptureBestPositionChange(scores chan Scoring, stopProcess chan bool) int {

	bestPositionChange := 0
	finished := false
	maxScore := 0

	for !finished {
		select {
		case finished = <-stopProcess:
		case scoring := <-scores:
			if scoring.Score > maxScore {
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

	//availablePositionChanges := board.GetLegalPositionChangesForPlayer(node.Player, node.Board)

	zoningScore := 1
	supremacyScore := 1

	totalScore := zoningScore + supremacyScore

	if node.IsOpponent {
		return -totalScore
	}

	return totalScore
}

func RecursiveNodeVisitor(rootNode Node, out chan Node) {
	for _, node := range NodeVisitor(rootNode) {
		out <- node
		defer func() { go RecursiveNodeVisitor(node, out) }()
	}
}

func NodeVisitor(node Node) []Node {
	player := node.Players[node.IndexCurrentPlayer]
	out := []Node{}

	legalPositionChanges := GetLegalPositionChangesForPlayer(player, node.Board)
	for _, positionChange := range legalPositionChanges {
		_, nodeBoard := board.DealPosition(node.Board, positionChange)
		out = append(out, Node{nodeBoard, positionChange, node.RootPositionChange, !node.IsOpponent, node.Players, 1 - node.IndexCurrentPlayer, node.Depth + 1})
	}
	return out
}

func GetLegalPositionChangesForPlayer(player player.Player, currentBoard board.Board) []int {
	legalPositionChanges := make([]int, (player.MaxPosition - player.MinPosition))
	for position := player.MinPosition; position < player.MaxPosition; position++ {
		if board.CanPlayerPlayPosition(player, currentBoard, position) {
			legalPositionChanges = append(legalPositionChanges, position)
		}
	}
	return legalPositionChanges
}
