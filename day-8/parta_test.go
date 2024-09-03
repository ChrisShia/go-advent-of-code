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
		nodeToBeTransformed string
		operator            Matrix
		want                string
	}{
		{"transform node AAA by left op", "AAA", left, "BBB"},
		{"transform node BBB by left op", "BBB", left, "AAA"},
		{"transform node ZZZ by left op", "ZZZ", left, "ZZZ"},
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
		want                         string
	}{
		{"apply LR on AAA", "AAA", []int{0, 1}, "ZZZ"},
		{"apply LLR on AAA", "AAA", []int{0, 0, 1}, "BBB"},
		{"apply LLRL on AAA", "AAA", []int{0, 0, 1, 0}, "AAA"},
		{"apply LLRLL on AAA", "AAA", []int{0, 0, 1, 0, 0}, "BBB"},
		{"apply LLRLLR on AAA", "AAA", []int{0, 0, 1, 0, 0, 1}, "ZZZ"},
		{"apply LLRLLRRR on AAA", "AAA", []int{0, 0, 1, 0, 0, 1, 1, 1}, "ZZZ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := applyLeftRightTurnsOnStartingNode(left, right, tt.startNodeId, tt.leftOrRightBinRepresentation)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func Test_Count(t *testing.T) {
	left, right := createAdjacencyMatrices()
	var tests = []struct {
		name                         string
		startNodeId                  string
		leftOrRightBinRepresentation []int
		want                         string
	}{
		{"apply LR on AAA", "AAA", []int{0, 1}, "ZZZ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := applyLeftRightTurnsOnStartingNode(left, right, tt.startNodeId, tt.leftOrRightBinRepresentation)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func createAdjacencyMatrices() (Matrix, Matrix) {
	input := [][]string{{"AAA", "BBB", "BBB"}, {"BBB", "AAA", "ZZZ"}, {"ZZZ", "ZZZ", "ZZZ"}}
	orderedKeys := make([]string, 0)
	leftMatrix := Matrix{newOrderedMap(&orderedKeys)}
	rightMatrix := Matrix{newOrderedMap(&orderedKeys)}
	nodeSetter := setNodesFromAnyInput[string](leftMatrix, rightMatrix, func(input string) string { return input })
	for _, adjacency := range input {
		nodeSetter(adjacency)
	}
	return leftMatrix, rightMatrix
}
