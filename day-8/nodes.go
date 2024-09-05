package main

import (
	"fmt"
	"go-advent-of-code/utils"
	"os"
)

const inputPath_ = "input/day-8.txt"

var leftRightTurns_ []int

func main() {
	file := utils.OpenFileLogFatal(inputPath_)
	defer utils.CloseFile(file)
	leftTurnOperator, rightTurnOperator := createLeftRightOperators(file)
	fmt.Println(applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator, "AAA", leftRightTurns_))
}

func createLeftRightOperators(file *os.File) (Matrix, Matrix) {
	orderedKeys := make([]string, 0)
	leftOrderedMap := newOrderedMap(&orderedKeys)
	rightOrderedMap := newOrderedMap(&orderedKeys)
	nodeSetter := setNodesFromInput[[]byte](leftOrderedMap, rightOrderedMap, func(input []byte) string { return string(input) })
	populateLeftRightAdjacencyMatrices(file, nodeSetter)
	leftTurnOperator := Matrix{leftOrderedMap}
	rightTurnOperator := Matrix{rightOrderedMap}
	return leftTurnOperator, rightTurnOperator
}

func applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator Matrix, startingNodeId string, leftRightTurns []int) (int, adjacency) {
	return By{func(leftOrRight int, node adjacency) adjacency {
		if leftOrRight == 0 {
			return leftTurnOperator.transform(node)
		}
		if leftOrRight == 1 {
			return rightTurnOperator.transform(node)
		}
		return node
	},
		func(node adjacency) bool {
			if node.id() == "ZZZ" {
				return false
			}
			return true
		}, adjacency(node(startingNodeId))}.apply(leftRightTurns)
}

func applyLeftRightTurnsOnStartingState(leftTurnOperator, rightTurnOperator Matrix, startingNodeId string, leftRightTurns []int) (int, adjacency) {
	return By{func(leftOrRight int, node adjacency) adjacency {
		if leftOrRight == 0 {
			return leftTurnOperator.transform(node)
		}
		if leftOrRight == 1 {
			return rightTurnOperator.transform(node)
		}
		return node
	},
		func(node adjacency) bool {
			if node.id() == "ZZZ" {
				return false
			}
			return true
		}, adjacency(node(startingNodeId))}.apply(leftRightTurns)
}

type By struct {
	transform          func(leftOrRight int, nodeId adjacency) adjacency
	keepCounting       func(input adjacency) bool
	startingFromNodeId adjacency
}

func (by By) apply(path path) (int, adjacency) {
	resultantNodeId := by.startingFromNodeId
	c := 0
	for by.keepCounting(resultantNodeId) {
		leftOrRight := path.step(c)
		resultantNodeId = by.transform(leftOrRight, resultantNodeId)
		c++
	}
	return c, resultantNodeId
}

type path []int

func (p path) step(i int) int {
	if i < len(p) {
		return p[i]
	}
	return p[mod(i, len(p))]
}

func mod(i, j int) int {
	return i % j
}

type adjacency interface {
	adjacent(m Matrix) adjacency
	id() string
}

type node string
type NodeState []adjacency

func (n node) id() string {
	return n.string()
}

func (n node) string() string {
	return string(n)
}

func (ns NodeState) id() string {
	return ""
}

func (ns NodeState) adjacent(m Matrix) adjacency {
	adjacentState := make([]adjacency, 0)
	for _, n := range ns {
		adjacentState = append(adjacentState, n.adjacent(m))
	}
	return adjacency(NodeState(adjacentState))
}

func (n node) adjacent(m Matrix) adjacency {
	adjacencyOfNode := m.adjacencyMap[n.string()]
	return node((*m.orderedKeys)[adjacencyOfNode-1])
}

type Matrix struct {
	*OrderedMap
}

func (m Matrix) transform(node adjacency) adjacency {
	return node.adjacent(m)
}

func (m Matrix) multiply(o Matrix) Matrix {
	orderedKeys := *m.orderedKeys
	resAdjacencyMap := make(map[string]int)
	productMap := OrderedMap{resAdjacencyMap, &orderedKeys}
	for i := range orderedKeys {
		rowSlice := m.row(i + 1)
		for _, key := range orderedKeys {
			columnNonZeroElement := o.adjacencyMap[key]
			ijElement := rowSlice[columnNonZeroElement-1]
			if ijElement == 0 {
				continue
			}
			resAdjacencyMap[key] = i + 1
		}
	}
	matrix := Matrix{&productMap}
	return matrix
}

func (m Matrix) row(rowIndex int) []int {
	orderedKeys := *m.orderedKeys
	var r = make([]int, 0)
	for _, nodeId := range orderedKeys {
		nodeAdjacency := m.adjacencyMap[nodeId]
		if rowIndex == nodeAdjacency {
			r = append(r, 1)
		} else {
			r = append(r, 0)
		}
	}
	return r
}

func (m Matrix) column(colIndex int) []int {
	orderedKeys := *m.orderedKeys
	var col = make([]int, 0)
	nodeId := orderedKeys[colIndex]
	adj := m.adjacencyMap[nodeId]
	for k := range orderedKeys {
		if adj == k+1 {
			col = append(col, 1)
		} else {
			col = append(col, 0)
		}
	}
	return col
}

type OrderedMap struct {
	adjacencyMap map[string]int
	orderedKeys  *[]string
}

func newOrderedMap(orderedKeys *[]string) *OrderedMap {
	return &OrderedMap{make(map[string]int), orderedKeys}
}

func (om *OrderedMap) getOrder(key string) int {
	for k, v := range *om.orderedKeys {
		if v == key {
			return k + 1
		}
	}
	return -1
}

func (om *OrderedMap) addSingleAdjacencyForNode(nodeId string, adjNodeId string) {
	if len(nodeId) == 0 || len(adjNodeId) == 0 {
		return
	}
	nodeOrder := om.getOrder(nodeId)
	adjNodeOrder := om.getOrder(adjNodeId)
	if nodeOrder == -1 {
		*om.orderedKeys = append(*om.orderedKeys, nodeId)
	}
	if adjNodeId != nodeId && adjNodeOrder == -1 {
		*om.orderedKeys = append(*om.orderedKeys, adjNodeId)
	}
	om.adjacencyMap[nodeId] = om.getOrder(adjNodeId)
}
