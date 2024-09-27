package main

import (
	"github.com/pterm/pterm"
	"time"
)

func applyLeftRightTurnsOnStartingState(leftTurnOperator, rightTurnOperator Matrix, leftRightTurns []int) {
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

func leftRightTransformFunc(leftTurnOperator Matrix, rightTurnOperator Matrix) func(leftOrRight int, node adjacency) adjacency {
	return func(leftOrRight int, node adjacency) adjacency {
		if leftOrRight == 0 {
			return leftTurnOperator.transform(node)
		}
		if leftOrRight == 1 {
			return rightTurnOperator.transform(node)
		}
		return node
	}
}

func initializeTeamOfWalkersAtStartingNodes(leftTurnOperator Matrix, multi *pterm.MultiPrinter) adjacency {
	startingNodes := make([]string, 0)
	for _, nodeId := range *leftTurnOperator.orderedKeys {
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
