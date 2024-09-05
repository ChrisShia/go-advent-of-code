package main

import (
	"fmt"
	"go-advent-of-code/utils"
	"log"
	"os"
)

const inputPath_ = "input/day-8.txt"

var leftRightTurns_ []int

func main() {
	file := utils.OpenFileLogFatal(inputPath_)
	defer utils.CloseFile(file)
	leftTurnOperator, rightTurnOperator := createLeftRightOperators(file)
	//resultNodeId := applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator, "AAA", leftRightTurns_)
	counter := count[string](
		func(start string) string {
			return applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator, start, leftRightTurns_)
		},
		func(input string) bool {
			if input == "ZZZ" {
				return false
			}
			return true
		}, "AAA")
	fmt.Println(counter * len(leftRightTurns_))
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

func count[T any](operation func(start T) T, keepCounting func(input T) bool, start T) int {
	result := operation(start)
	counter := 1
	for keepCounting(result) {
		result = operation(result)
		counter++
	}
	fmt.Println(result)
	return counter
}

func applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator Matrix, startingNodeId string, leftRightTurns []int) string {
	return By{func(leftOrRight int, nodeId string) string {
		if leftOrRight == 0 {
			return leftTurnOperator.transform(nodeId)
		}
		if leftOrRight == 1 {
			return rightTurnOperator.transform(nodeId)
		}
		return nodeId
	}, startingNodeId}.walk(leftRightTurns)
}

type By struct {
	transform          func(leftOrRight int, nodeId string) string
	startingFromNodeId string
}

func (by By) walk(turns []int) string {
	resultantNodeId := by.startingFromNodeId
	for _, leftOrRight := range turns {
		if resultantNodeId == "" {
			log.Fatal("Not Defined")
		}
		resultantNodeId = by.transform(leftOrRight, resultantNodeId)
	}
	return resultantNodeId
}

//func (by *By) reduce(operation []byte) transform {
//	var result = make([]transform, 0)
//	for index := 0; index < len(operation)-1; index++ {
//		by(operation[index], operation[index+1])
//	}
//}

//type transform struct {
//	m Matrix
//}

type Matrix struct {
	om *OrderedMap
}

func (m Matrix) transform(nodeId string) string {
	adjacencyOfNode := m.om.adjacencyMap[nodeId]
	return (*m.om.orderedKeys)[adjacencyOfNode-1]
}

func (m Matrix) multiply(o Matrix) Matrix {
	orderedKeys := *m.om.orderedKeys
	resAdjacencyMap := make(map[string]int)
	productMap := OrderedMap{resAdjacencyMap, &orderedKeys}
	for i := range orderedKeys {
		rowSlice := m.row(i + 1)
		for _, key := range orderedKeys {
			columnNonZeroElement := o.om.adjacencyMap[key]
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
	orderedKeys := *m.om.orderedKeys
	var r = make([]int, 0)
	for _, nodeId := range orderedKeys {
		nodeAdjacency := m.om.adjacencyMap[nodeId]
		if rowIndex == nodeAdjacency {
			r = append(r, 1)
		} else {
			r = append(r, 0)
		}
	}
	return r
}

func (m Matrix) column(colIndex int) []int {
	orderedKeys := *m.om.orderedKeys
	var col = make([]int, 0)
	nodeId := orderedKeys[colIndex]
	adj := m.om.adjacencyMap[nodeId]
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

//type Column func(dim int) []any
//
//func (from *Column) get(index int) any {
//	return (*from)()[index]
//}
//
//func orthonormalIntColumn(nonZeroElement int) Column {
//	return func() []any {
//		ints := make([]int, 0)
//		if
//	}
//}
