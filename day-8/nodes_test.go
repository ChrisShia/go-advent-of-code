package main

import (
	"fmt"
	"testing"
)

func Test_Multiply(t *testing.T) {
	left, right := createAdjacencyMatrices()
	leftRight := left.multiply(right)
	rightLeft := right.multiply(left)
	fmt.Println(leftRight.adjacencyMap)
	fmt.Println(rightLeft.adjacencyMap)
}

func Test_TransformNode(t *testing.T) {
	left, _ := createAdjacencyMatrices()
	var tests = []struct {
		name                string
		nodeToBeTransformed adjacency
		operator            Matrix
		want                adjacency
	}{
		{"transform node AAA by left op", node("AAA"), left, node("BBB")},
		{"transform node BBB by left op", node("BBB"), left, node("AAA")},
		{"transform node ZZZ by left op", node("ZZZ"), left, node("ZZZ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.operator.transform(tt.nodeToBeTransformed)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func Test_ApplyLeftRightTurnsOnStartingNode(t *testing.T) {
	left, right := createAdjacencyMatrices()
	var tests = []struct {
		name                         string
		startNodeId                  string
		leftOrRightBinRepresentation []int
		want                         adjacency
	}{
		{"apply LR on AAA", "AAA", []int{0, 1}, node("ZZZ")},
		{"apply LLR on AAA", "AAA", []int{0, 0, 1}, node("BBB")},
		{"apply LLRL on AAA", "AAA", []int{0, 0, 1, 0}, node("AAA")},
		{"apply LLRLL on AAA", "AAA", []int{0, 0, 1, 0, 0}, node("BBB")},
		{"apply LLRLLR on AAA", "AAA", []int{0, 0, 1, 0, 0, 1}, node("ZZZ")},
		{"apply LLRLLRRR on AAA", "AAA", []int{0, 0, 1, 0, 0, 1, 1, 1}, node("ZZZ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ans := applyLeftRightTurnsOnStartingNode(left, right, tt.startNodeId, tt.leftOrRightBinRepresentation)
			if ans.string() != tt.want.string() {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func Test_Count(t *testing.T) {
	left, right := createAdjacencyMatrices()
	var tests = []struct {
		name                         string
		startNodeId                  adjacency
		leftOrRightBinRepresentation []int
		want                         adjacency
	}{
		{"apply LR on AAA", node("AAA"), []int{0, 1}, node("ZZZ")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ans := applyLeftRightTurnsOnStartingNode(left, right, tt.startNodeId.string(), tt.leftOrRightBinRepresentation)
			if ans.string() != tt.want.string() {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func Test_ContainsFunc_ContainsNodeEndingInZ(t *testing.T) {
	nodeState := createNodeState()
	var tests = []struct {
		name  string
		state adjacency
		want  bool
	}{
		{"apply contains node that ends with Z", adjacency(nodeState), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.state.containsFunc(stringEndsInZ)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func createNodeState() NodeState {
	nodes := []node{"AFG", "BBZ", "BBB", "HGF", "JIK"}
	adjSlice := make([]adjacency, 0)
	for _, n := range nodes {
		adjSlice = append(adjSlice, adjacency(n))
	}
	return adjSlice
}

func createAdjacencyMatrices() (Matrix, Matrix) {
	input := [][]string{{"AAA", "BBB", "BBB"}, {"BBB", "AAA", "ZZZ"}, {"ZZZ", "ZZZ", "ZZZ"}}
	orderedKeys := make([]string, 0)
	leftOrderedMap := newOrderedMap(&orderedKeys)
	rightOrderedMap := newOrderedMap(&orderedKeys)
	nodeSetter := nodesFromInputSetter[string](leftOrderedMap, rightOrderedMap, func(input string) string { return input })
	for _, adjacency := range input {
		nodeSetter(adjacency)
	}
	leftMatrix := Matrix{leftOrderedMap}
	rightMatrix := Matrix{rightOrderedMap}
	return leftMatrix, rightMatrix
}
