package main

import (
	"bufio"
	"bytes"
	"os"
)

func populateLeftRightAdjacencyMatrices(file *os.File, nodeSetter func(input [][]byte)) {
	scanner := bufio.NewScanner(file)
	var line []byte
	scanner.Scan()
	line = scanner.Bytes()
	leftRightTurns_ = binaryRepresentationOfLeftRight(line)
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

func nodesFromInputSetter[T any](leftOrderedMap, rightOrderedMap *OrderedMap, stringer func(T) string) func(input []T) {
	return func(input []T) {
		//nodeId := fmt.Sprintf("%v", input[0])
		nodeId := stringer(input[0])
		//leftAdjNodeId := fmt.Sprintf("%v", input[1])
		leftAdjNodeId := stringer(input[1])
		//rightAdjNodeId := fmt.Sprintf("%v", input[2])
		rightAdjNodeId := stringer(input[2])
		leftOrderedMap.addSingleAdjacencyForNode(nodeId, leftAdjNodeId)
		rightOrderedMap.addSingleAdjacencyForNode(nodeId, rightAdjNodeId)
	}
}

func isExcludedCharacter() func(r rune) bool {
	return func(r rune) bool {
		return bytes.ContainsRune([]byte("=(,)"), r)
	}
}

func createLeftRightOperators(file *os.File) (Matrix, Matrix) {
	orderedKeys := make([]string, 0)
	leftOrderedMap := newOrderedMap(&orderedKeys)
	rightOrderedMap := newOrderedMap(&orderedKeys)
	nodeSetter := nodesFromInputSetter[[]byte](leftOrderedMap, rightOrderedMap, func(input []byte) string { return string(input) })
	populateLeftRightAdjacencyMatrices(file, nodeSetter)
	leftTurnOperator := Matrix{leftOrderedMap}
	rightTurnOperator := Matrix{rightOrderedMap}
	return leftTurnOperator, rightTurnOperator
}
