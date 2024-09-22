package main

import (
	"fmt"
	"go-advent-of-code/utils"
	"strings"
)

const inputPath_ = "input/day-8.txt"

var leftRightTurns_ []int

func main() {
	file := utils.OpenFileLogFatal(inputPath_)
	defer utils.CloseFile(file)
	leftTurnOperator, rightTurnOperator := createLeftRightOperators(file)
	fmt.Println(applyLeftRightTurnsOnStartingState(leftTurnOperator, rightTurnOperator, leftRightTurns_))
}

func applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator Matrix, startingNodeId string, leftRightTurns []int) (int, adjacency) {
	return By{leftRightTransformFunc(leftTurnOperator, rightTurnOperator),
		func(nodeId string) bool {
			if nodeId == "ZZZ" {
				return false
			}
			return true
		}, adjacency(node(startingNodeId))}.apply(leftRightTurns, newCounter())
}

func isStartingNode(s string) bool {
	bytes := []byte(s)
	if len(bytes) == 0 {
		return false
	}
	return byteSliceEndsInA(bytes)
}

func byteSliceEndsInA(bs []byte) bool {
	return 'A' == bs[len(bs)-1]
}

func stringEndsInZ(s string) bool {
	return 'Z' == s[len(s)-1]
}

type By struct {
	transform    func(leftOrRight int, nodeId adjacency) adjacency
	keepCounting func(inputId string) bool
	startingFrom adjacency
}

func (by By) apply(path path, c counter) (int, adjacency) {
	resultantNode := by.startingFrom
	for !resultantNode.isEnd(by.keepCounting) {
		leftOrRight := path.step(c.count)
		resultantNode = by.transform(leftOrRight, resultantNode)
		c.increment(resultantNode)
	}
	return c.count, resultantNode
}

type counter struct {
	count       int
	displayFunc func(count int, input adjacency)
}

func (c *counter) increment(a adjacency) {
	c.count++
	c.displayFunc(c.count, a)
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
	string() string
	isEnd(func(id string) bool) bool
	containsFunc(func(id string) bool) bool
}

type node string
type NodeState []adjacency

func (n node) containsFunc(f func(id string) bool) bool {
	return f(n.string())
}

func (n node) isEnd(keepCountingFunc func(id string) bool) bool {
	if keepCountingFunc(n.string()) {
		return false
	}
	return true
}

func (n node) string() string {
	return string(n)
}

func (ns NodeState) containsFunc(f func(id string) bool) bool {
	for _, n := range ns {
		if n.containsFunc(f) {
			return true
		}
	}
	return false
}

func (ns NodeState) isEnd(keepCountingFunc func(id string) bool) bool {
	keepCounting := false
	isEnd := true
	for _, n := range ns {
		if keepCountingFunc(n.string()) {
			return keepCounting
		}
	}
	return isEnd
}

func (ns NodeState) string() string {
	stringBuilder := strings.Builder{}
	for _, n := range ns {
		stringBuilder.WriteString(n.string())
		stringBuilder.WriteByte(' ')
	}
	return stringBuilder.String()
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
