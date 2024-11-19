package main

import (
	"fmt"
	"github.com/ChrisShia/goread/read"
)

//References : green's calculus theorem relating surface area and the line integral on its boundary.

const inputPath_ = "input/day-10.txt"

var oneDimensionalGrid_ []*neighbourPair
var gridWidth_ int

func main() {
	oneDimensionalGrid_ = make([]*neighbourPair, 0)
	read.Read(inputPath_, nil, nil, nil, dayTenLineProcessor)
	c := initializeCircuit()
	var farthestPoint = traverseCircuitFromStartingPosition(c)
	fmt.Println(farthestPoint)
}

func traverseCircuitFromStartingPosition(c *circuit) int {
	traversalCount := 1
	w1 := walker{position: oneDimensionalGrid_[c.startIndex].neighbour1, previousPosition: c.startIndex}
	w2 := walker{position: oneDimensionalGrid_[c.startIndex].neighbour2, previousPosition: c.startIndex}
	integral := 0
	for w1.position != w2.position {
		//integral += w1.position
		w1.proceed()
		w2.proceed()
		traversalCount++
	}
	return traversalCount
}

type circuit struct {
	startIndex          int
	lineStartToNodesMap map[int][]int
}

func initializeCircuit() *circuit {
	setStartingPositionNeighbors()
	return &circuit{startIndex: startingPosition_}
}

type walker struct {
	position         int
	previousPosition int
}

func (w *walker) proceed() int {
	neighbour1 := oneDimensionalGrid_[w.position].neighbour1
	neighbour2 := oneDimensionalGrid_[w.position].neighbour2
	if neighbour1 != w.previousPosition {
		w.Visit(neighbour1)
		return neighbour1
	}
	w.Visit(neighbour2)
	return neighbour2
}

func (w *walker) Visit(neighbour int) {
	w.previousPosition = w.position
	w.position = neighbour
}

func setStartingPositionNeighbors() {
	westIndex := -1
	if !atLeftGridEdge(startingPosition_) {
		westIndex = startingPosition_ - 1
	}
	northIndex := -1
	if !animalAtFirstLine() {
		northIndex = positionOfNorthIndex(startingPosition_)
	}
	eastIndex := -1
	if !atRightGridEdge(startingPosition_) {
		eastIndex = startingPosition_ + 1
	}
	southIndex := -1
	if !animalAtLastLine() {
		southIndex = positionOfSouthNeighbour(startingPosition_)
	}
	assignNeighbourOfAnimalIfAppropriate(westIndex)
	assignNeighbourOfAnimalIfAppropriate(northIndex)
	assignNeighbourOfAnimalIfAppropriate(eastIndex)
	assignNeighbourOfAnimalIfAppropriate(southIndex)
}

func assignNeighbourOfAnimalIfAppropriate(index int) {
	if index >= 0 && oneDimensionalGrid_[index] != nil {
		if animalInNeighbourhoodOf(index) {
			setAnimalNeighbourIfAbsent(index)
		}
	}
}

func setAnimalNeighbourIfAbsent(neighbourIndex int) bool {
	if oneDimensionalGrid_[startingPosition_].neighbour1 == startingPosition_ {
		oneDimensionalGrid_[startingPosition_].neighbour1 = neighbourIndex
		return true
	} else if oneDimensionalGrid_[startingPosition_].neighbour2 == startingPosition_ {
		oneDimensionalGrid_[startingPosition_].neighbour2 = neighbourIndex
		return true
	}
	return false
}

func animalInNeighbourhoodOf(index int) bool {
	return oneDimensionalGrid_[index].neighbour1 == startingPosition_ || oneDimensionalGrid_[index].neighbour2 == startingPosition_
}

func animalAtLastLine() bool {
	return startingPosition_ > len(oneDimensionalGrid_)-gridWidth_
}

func animalAtFirstLine() bool {
	return startingPosition_ < gridWidth_
}
