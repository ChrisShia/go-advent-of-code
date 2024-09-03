package main

import (
	"fmt"
	"log"
)

func count[T any](operation func(start T) T, keepCounting func(input T) bool, start T) int {
	result := start
	counter := 0
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
	}, startingNodeId}.apply(leftRightTurns)
}

type By struct {
	transform          func(leftOrRight int, nodeId string) string
	startingFromNodeId string
}

func (by By) apply(turns []int) string {
	resultantNodeId := by.startingFromNodeId
	for _, leftOrRight := range turns {
		if resultantNodeId == "" {
			log.Fatal("Not Defined")
		}
		resultantNodeId = by.transform(leftOrRight, resultantNodeId)
	}
	return resultantNodeId
}

type Matrix struct {
	*OrderedMap
}

func (m Matrix) transform(nodeId string) string {
	adjacencyOfNode := m.adjacencyMap[nodeId]
	return (*m.orderedKeys)[adjacencyOfNode-1]
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
