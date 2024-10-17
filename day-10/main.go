package main

import (
	"fmt"
	"go-advent-of-code/utils"
	"go-advent-of-code/utils/maths"
)

//References : ray casting, flood fill, breadth first search

const inputPath_ = "input/day-10.txt"

var oneDimensionalGrid_ []*neighbourPair
var gridWidth_ int

func main() {
	oneDimensionalGrid_ = make([]*neighbourPair, 0)
	utils.Read(inputPath_, nil, nil, nil, dayTenLineProcessor)
	setAnimalNeighbors()
	c := initializeCircuit(animal1dGridPosition_)
	var farthestPoint = traverseFarthestFromStart(c)
	//countForEnclosedPoints(c)
	fmt.Println(farthestPoint)
}

//func countForEnclosedPoints(c *circuit) {
//	startIndex := westMostIndexOnLongitude(c.northEnd)
//	endIndex := westMostIndexOnLongitude(c.southEnd)
//	for i := startIndex; i <= endIndex; i += gridWidth_ {
//		if c.contains(i) {
//			circuitNodes := c.lineStartToNodesMap[i]
//			if enclosedPoint, found := findInnerNodeAlongLine(circuitNodes); found {
//
//				break
//			}
//		}
//	}
//}

func (c *circuit) withinBoundaries(gridPoint int) bool {
	if c.contains(gridPoint) {
		return false
	}
	northBoundaryPoint := gridPoint
	southBoundaryPoint := gridPoint
	eastBoundaryPoint := gridPoint
	westBoundaryPoint := gridPoint
	for !c.contains(northBoundaryPoint) {
		northBoundaryPoint = move(northBoundaryPoint, north)
		if northBoundaryPoint == indexOutOfGridBounds {
			break
		}
	}
	for !c.contains(southBoundaryPoint) {
		southBoundaryPoint = move(southBoundaryPoint, south)
		if southBoundaryPoint == indexOutOfGridBounds {
			break
		}
	}
	for !c.contains(eastBoundaryPoint) {
		eastBoundaryPoint = move(eastBoundaryPoint, east)
		if eastBoundaryPoint == indexOutOfGridBounds {
			break
		}
	}
	for !c.contains(westBoundaryPoint) {
		westBoundaryPoint = move(westBoundaryPoint, west)
		if westBoundaryPoint == indexOutOfGridBounds {
			break
		}
	}
	var compass = []int{northBoundaryPoint, eastBoundaryPoint, southBoundaryPoint, westBoundaryPoint}
	walker{position: oneDimensionalGrid_[northBoundaryPoint].neighbour1}
}

func move(fromGridPoint int, toDirection direction) int {
	return pipe(fromGridPoint, toDirection)
}

func findInnerNodeAlongLine(circuitNodes []int) (int, bool) {
	found := true
	for i := 0; i < len(circuitNodes); i += 2 {
		eastAdjacentGridIndex := circuitNodes[i] + 1
		if eastAdjacentGridIndex == circuitNodes[i+1] {
			continue
		} else {
			return eastAdjacentGridIndex, found
		}
	}
	return -1, !found
}

func eastMostIndexOnLongitude(indexOnGrid int) int {
	return gridWidth_*(indexOnGrid/gridWidth_) + gridWidth_ - 1
}

func westMostIndexOnLongitude(indexOnGrid int) int {
	indexModWidth := maths.Mod(indexOnGrid+1, gridWidth_)
	return indexOnGrid - indexModWidth
}

func traverseFarthestFromStart(c *circuit) int {
	traversalCount := 1
	w1 := walker{position: oneDimensionalGrid_[c.startIndex].neighbour1, previousPosition: c.startIndex}
	w2 := walker{position: oneDimensionalGrid_[c.startIndex].neighbour2, previousPosition: c.startIndex}
	for w1.position != w2.position {
		w1.proceed()
		w2.proceed()
		//westMostValue, eastMostValue := westAndEastMostValue(w1, w2)
		southMostValue, northMostValue := southAndNorthMostValue(w1, w2)
		c.setCircuitPolarExtremes(northMostValue, southMostValue)
		c.mapToStartOfLine(w1.position)
		c.mapToStartOfLine(w2.position)
		traversalCount++
	}
	return traversalCount
}

func southAndNorthMostValue(w1 walker, w2 walker) (int, int) {
	southMostValue := w1.position
	northMostValue := w1.position
	if w2.position > southMostValue {
		southMostValue = w2.position
	} else if w2.position < northMostValue {
		northMostValue = w2.position
	}
	return southMostValue, northMostValue
}

func westAndEastMostValue(w1 walker, w2 walker) (int, int) {
	westMostValue := w1.position
	eastMostValue := w1.position
	if maths.Mod(w2.position, gridWidth_) < maths.Mod(westMostValue, gridWidth_) {
		westMostValue = w2.position
	} else if maths.Mod(w2.position, gridWidth_) > maths.Mod(eastMostValue, gridWidth_) {
		eastMostValue = w2.position
	}
	return westMostValue, eastMostValue
}

func (c *circuit) setCircuitPolarExtremes(northMostValue int, southMostValue int) {
	//if maths.Mod(westMostValue, gridWidth_) < maths.Mod(c.westEnd, gridWidth_) {
	//	c.westEnd = westMostValue
	//}
	//if maths.Mod(eastMostValue, gridWidth_) > maths.Mod(c.eastEnd, gridWidth_) {
	//	c.eastEnd = eastMostValue
	//}
	if northMostValue/gridWidth_ < c.northEnd/gridWidth_ {
		c.northEnd = northMostValue
	}
	if southMostValue/gridWidth_ > c.southEnd/gridWidth_ {
		c.southEnd = southMostValue
	}
}

func (c *circuit) contains(index int) bool {
	indexStartOfLine := (index / gridWidth_) * gridWidth_
	for _, i := range c.lineStartToNodesMap[indexStartOfLine] {
		if i == index {
			return true
		}
	}
	return false
}

func (c *circuit) mapToStartOfLine(position int) {
	startOfLineIndex := (position / gridWidth_) * gridWidth_
	if c.lineStartToNodesMap[startOfLineIndex] == nil {
		c.lineStartToNodesMap[startOfLineIndex] = make([]int, 0)
	}
	c.lineStartToNodesMap[startOfLineIndex] = append(c.lineStartToNodesMap[startOfLineIndex], position)
}

type circuit struct {
	startIndex, northEnd, southEnd int
	lineStartToNodesMap            map[int][]int
}

func initializeCircuit(start int) *circuit {
	return &circuit{startIndex: start, northEnd: len(oneDimensionalGrid_), lineStartToNodesMap: make(map[int][]int)}
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
	} else {
		w.Visit(neighbour2)
		return neighbour2
	}
	return -1
}

func (w *walker) Visit(neighbour int) {
	w.previousPosition = w.position
	w.position = neighbour
}

func setAnimalNeighbors() {
	westIndex := -1
	if !atLeftGridEdge(animal1dGridPosition_) {
		westIndex = animal1dGridPosition_ - 1
	}
	northIndex := -1
	if !animalAtFirstLine() {
		northIndex = positionOfNorthIndex(animal1dGridPosition_)
	}
	eastIndex := -1
	if !atRightGridEdge(animal1dGridPosition_) {
		eastIndex = animal1dGridPosition_ + 1
	}
	southIndex := -1
	if !animalAtLastLine() {
		southIndex = positionOfSouthNeighbour(animal1dGridPosition_)
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
	if oneDimensionalGrid_[animal1dGridPosition_].neighbour1 == animal1dGridPosition_ {
		oneDimensionalGrid_[animal1dGridPosition_].neighbour1 = neighbourIndex
		return true
	} else if oneDimensionalGrid_[animal1dGridPosition_].neighbour2 == animal1dGridPosition_ {
		oneDimensionalGrid_[animal1dGridPosition_].neighbour2 = neighbourIndex
		return true
	}
	return false
}

func animalInNeighbourhoodOf(index int) bool {
	return oneDimensionalGrid_[index].neighbour1 == animal1dGridPosition_ || oneDimensionalGrid_[index].neighbour2 == animal1dGridPosition_
}

func animalAtLastLine() bool {
	return animal1dGridPosition_ > len(oneDimensionalGrid_)-gridWidth_
}

func animalAtFirstLine() bool {
	return animal1dGridPosition_ < gridWidth_
}
