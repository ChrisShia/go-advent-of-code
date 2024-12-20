package main

import (
	maths "github.com/ChrisShia/math-depot"
	"github.com/pterm/pterm"
	"time"
)

func applyLeftRightTurnsOnStartingState(leftTurnOperator, rightTurnOperator maths.Matrix, leftRightTurns []int) {
	multiPrinter := pterm.DefaultMultiPrinter
	multiPrinter.UpdateDelay = time.Millisecond * 200
	t := initializeTeamOfWalkersAtStartingNodes(leftTurnOperator, &multiPrinter)
	multiPrinter.Start()
	By{
		turn:         leftRightTransformFunc(leftTurnOperator, rightTurnOperator),
		keepCounting: keepCountingIfNodeEndsInZFunc(),
		startingFrom: t}.apply(leftRightTurns)
	defer multiPrinter.Stop()
	return
}

func keepCountingIfNodeEndsInZFunc() func(nodeId string) bool {
	return func(nodeId string) bool {
		if stringEndsInZ(nodeId) {
			return false
		}
		return true
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
	startingNodes := make([]string, 0)
	for _, nodeId := range *leftTurnOperator.OrderedKeys {
		if isStartingNode(nodeId) {
			startingNodes = append(startingNodes, nodeId)
		}
	}
	startingWalker := make([]adjacency, 0)
	teamWriter := (*multi).NewWriter()
	for _, sn := range startingNodes {
		writer := (*multi).NewWriter()
		w := walker{
			pos:            sn,
			firstPos:       sn,
			visualizer:     walkerVisualizerWithEmbeddedCache(writer),
			cachePredicate: func(pos string) bool { return stringEndsInZ(pos) },
		}
		startingWalker = append(startingWalker, adjacency(&w))
	}
	t := newTeam(startingWalker, teamVisualizer(teamWriter))
	return adjacency(&t)
}
