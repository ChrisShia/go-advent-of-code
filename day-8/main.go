package main

import (
	"bufio"
	"fmt"
	"go-advent-of-code/utils"
)

//		part a : 20093
//		part b :

func main() {

	file := utils.OpenFileLogFatal(inputPath_)
	defer utils.CloseFile(file)
	scanner := bufio.NewScanner(file)
	var line []byte
	scanner.Scan()
	line = scanner.Bytes()
	leftRightTurns := binaryRepresentationOfLeftRight(line)
	solvePartA(scanner, leftRightTurns)
	//solvePartB(scanner, leftRightTurns)
}

func solvePartA(scanner *bufio.Scanner, leftRightTurns []int) {
	leftTurnOperator, rightTurnOperator := createLeftRightOperators(scanner, setNodesFromByteInput())
	counter := count[string](
		func(start string) string {
			return applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator, start, leftRightTurns)
		},
		func(input string) bool {
			keepCounting := true
			stop := false
			if input == "ZZZ" {
				return stop
			}
			return keepCounting
		},
		"AAA")
	fmt.Println(counter * len(leftRightTurns))
}

//func solvePartB(scanner *bufio.Scanner, leftRightTurns []int) {
//	leftTurnOperator, rightTurnOperator := createLeftRightOperators(scanner, setNodesFromByteInput())
//	startingNodes := make([]string, 0)
//	for _, nodeId := range *leftTurnOperator.orderedKeys {
//		if isStartingNode(nodeId) {
//			startingNodes = append(startingNodes, nodeId)
//		}
//	}
//	counter := count[[]string](
//		func(fromNodes []string) []string {
//			toNodes := make([]string, 0)
//			for _, n := range fromNodes {
//				toNodes = append(toNodes,
//					applyLeftRightTurnsOnStartingNode(leftTurnOperator, rightTurnOperator, n, leftRightTurns))
//			}
//			return toNodes
//		},
//		func(nodes []string) bool {
//			keepCounting := true
//			stop := false
//			for _, n := range nodes {
//				if !stringEndsInZ(n) {
//					return keepCounting
//				}
//			}
//			return stop
//		},
//		startingNodes)
//	fmt.Println(counter * len(leftRightTurns))
//}
