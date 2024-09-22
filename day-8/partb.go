package main

import "fmt"

func applyLeftRightTurnsOnStartingState(leftTurnOperator, rightTurnOperator Matrix, leftRightTurns []int) (int, adjacency) {
	startingNodesAdjacency := createStartingAdjacencySlice(leftTurnOperator)
	return By{
		leftRightTransformFunc(leftTurnOperator, rightTurnOperator),
		keepCountingIfNodeEndsInZFunc(),
		adjacency(NodeState(startingNodesAdjacency))}.apply(leftRightTurns, newCounter())
}

func newCounter() counter {
	return counter{0, displayCountAndState}
}

func statefulDisplay() func(int, adjacency) {
	//i := 0
	return func(count int, a adjacency) {
		if a.containsFunc(stringEndsInZ) {
			fmt.Println(a.string())
		}
	}
}

func displayCountAndState(count int, a adjacency) {
	if a.containsFunc(stringEndsInZ) {
		fmt.Println(a.string())
	}
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

func createStartingAdjacencySlice(leftTurnOperator Matrix) []adjacency {
	startingNodes := make([]string, 0)
	for _, nodeId := range *leftTurnOperator.orderedKeys {
		if isStartingNode(nodeId) {
			startingNodes = append(startingNodes, nodeId)
		}
	}
	startingNodesAdjacency := make([]adjacency, 0)
	for _, sn := range startingNodes {
		startingNodesAdjacency = append(startingNodesAdjacency, adjacency(node(sn)))
	}
	return startingNodesAdjacency
}
