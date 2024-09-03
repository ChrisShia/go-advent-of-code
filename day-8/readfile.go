package main

import (
	"bufio"
	"bytes"
)

const inputPath_ = "input/day-8.txt"

func populateLeftRightAdjacencyMatrices(scanner *bufio.Scanner, nodeSetter func(input [][]byte)) {
	var line []byte
	for scanner.Scan() {
		line = scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		fields := bytes.Fields(line)
		trimmedFields := trimUnnecessaryChars(fields)
		nodeSetter(trimmedFields)
	}
}

func trimUnnecessaryChars(lineFields [][]byte) [][]byte {
	res := make([][]byte, 0)
	for _, s := range lineFields {
		trimmedField := bytes.TrimFunc(s, isExcludedCharacter())
		if len(trimmedField) == 0 {
			continue
		}
		res = append(res, trimmedField)
	}
	return res
}

func binaryRepresentationOfLeftRight(line []byte) []int {
	var binarySequenceOfTurns = make([]int, 0)
	appendFunc := func(binValue int) { binarySequenceOfTurns = append(binarySequenceOfTurns, binValue) }
	for _, v := range line {
		if v == 'L' {
			appendFunc(0)
		}
		if v == 'R' {
			appendFunc(1)
		}
	}
	return binarySequenceOfTurns
}

func setNodesFromByteInput() func(left, right Matrix) func(input [][]byte) {
	return func(left, right Matrix) func(input [][]byte) {
		return setNodesFromAnyInput[[]byte](left, right, byteStringer())
	}
}

//func setNodesFromByteInputAndIdentifyStartingNodes() func(left, right *OrderedMap) func(input [][]byte) {
//	return func(left, right *OrderedMap) func(input [][]byte) {
//		return setNodesFromInputAndPopulateStartingNodes[[]byte](left, right, byteStringer())
//	}
//}

func byteStringer() func([]byte) string {
	return func(input []byte) string { return string(input) }
}

func setNodesFromAnyInput[T any](leftOrderedMap, rightOrderedMap Matrix, stringer func(T) string) func(input []T) {
	return func(input []T) {
		nodeId := stringer(input[0])
		leftAdjNodeId := stringer(input[1])
		rightAdjNodeId := stringer(input[2])
		leftOrderedMap.addSingleAdjacencyForNode(nodeId, leftAdjNodeId)
		rightOrderedMap.addSingleAdjacencyForNode(nodeId, rightAdjNodeId)
	}
}

func createLeftRightOperators(scanner *bufio.Scanner, adjacencySetterCreator func(left, right Matrix) func(input [][]byte)) (Matrix, Matrix) {
	orderedKeys := make([]string, 0)
	leftTurnOperator := Matrix{newOrderedMap(&orderedKeys)}
	rightTurnOperator := Matrix{newOrderedMap(&orderedKeys)}
	adjacencySetter := adjacencySetterCreator(leftTurnOperator, rightTurnOperator)
	var line []byte
	for scanner.Scan() {
		line = scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		fields := bytes.Fields(line)
		trimmedFields := trimUnnecessaryChars(fields)
		adjacencySetter(trimmedFields)
	}
	return leftTurnOperator, rightTurnOperator
}

func isExcludedCharacter() func(r rune) bool {
	return func(r rune) bool {
		return bytes.ContainsRune([]byte("=(,)"), r)
	}
}
