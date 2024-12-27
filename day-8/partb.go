package main

import (
	maths "github.com/ChrisShia/math-depot"
	"github.com/pterm/pterm"
	"io"
	"time"
)

func applyLeftRightTurnsOnStartingState(leftTurnOperator, rightTurnOperator maths.Matrix, leftRightTurns []int) {
	multiPrinter := pterm.DefaultMultiPrinter
	multiPrinter.UpdateDelay = time.Millisecond * 100
	t := initializeTeamOfWalkersAtStartingNodes(leftTurnOperator, &multiPrinter)
	multiPrinter.Start()
	By{
		turn:           leftRightTransformFunc(leftTurnOperator, rightTurnOperator),
		isEndPredicate: nodeIdEndsInZIsTheEndPredicate(),
		startingFrom:   t}.apply(leftRightTurns)
	defer multiPrinter.Stop()
	return
}

func nodeIdEndsInZIsTheEndPredicate() func(nodeId string) bool {
	return func(nodeId string) bool {
		if stringEndsInZ(nodeId) {
			return true
		}
		return false
	}
}

func leftRightTransformFunc(leftTurnOperator maths.Matrix, rightTurnOperator maths.Matrix) func(leftOrRight int, node adjacency) {
	return func(leftOrRight int, node adjacency) {
		if leftOrRight == 0 {
			node.progress(leftTurnOperator)
			return
		}
		if leftOrRight == 1 {
			node.progress(rightTurnOperator)
			return
		}
		return
	}
}

func initializeTeamOfWalkersAtStartingNodes(leftTurnOperator maths.Matrix, multi *pterm.MultiPrinter) adjacency {
	startingNodes := startingNodeIds(leftTurnOperator)
	startingWalkers := make([]adjacency, 0)
	teamWriter := (*multi).NewWriter()
	for _, sn := range startingNodes {
		writer := (*multi).NewWriter()
		w := newWalkerWithCachingVisualizer(sn, writer)
		startingWalkers = append(startingWalkers, adjacency(&w))
	}
	t := newTeam(startingWalkers, teamVisualizer(teamWriter))
	return adjacency(&t)
}

func startingNodeIds(leftTurnOperator maths.Matrix) []string {
	ids := make([]string, 0)
	for _, nodeId := range *leftTurnOperator.OrderedKeys {
		if isStartingNode(nodeId) {
			ids = append(ids, nodeId)
		}
	}
	return ids
}

func newWalkerWithCachingVisualizer(sn string, writer io.Writer) walker {
	return walker{
		position:       sn,
		firstPos:       sn,
		visualizer:     walkerVisualizerWithEmbeddedCache(writer),
		cachePredicate: func(pos string) bool { return stringEndsInZ(pos) },
	}
}
