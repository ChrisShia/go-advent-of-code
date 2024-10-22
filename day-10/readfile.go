package main

import maths "github.com/ChrisShia/math-depot"

const indexOutOfGridBounds = -1
const one = 1

var atLastLine_ bool
var startingPosition_ = -1
var pipeShapesCount_ = 0
var linesReadCount_ = 0

func dayTenLineProcessor(pipes [][]byte, atLastLine bool) {
	pipeShapes := pipes[0]
	setGridWidthOnce(pipeShapes)
	setLastLineFlag(atLastLine)
	for _, p := range pipeShapes {
		currentIndex := len(oneDimensionalGrid_)
		neighbour1, neighbour2 := pipesTo(p, currentIndex)
		oneDimensionalGrid_ = append(oneDimensionalGrid_, newNeighbourPair(neighbour1, neighbour2, p))
		pipeShapesCount_++
	}
	linesReadCount_++
}

func setLastLineFlag(atLastLine bool) {
	atLastLine_ = atLastLine
}

func pipesTo(pipeShape byte, referenceIndex int) (int, int) {
	switch pipeShape {
	case WestSouthPipe:
		return pipe(referenceIndex, west), pipe(referenceIndex, south)
	case WestEastPipe:
		return pipe(referenceIndex, west), pipe(referenceIndex, east)
	case WestNorthPipe:
		return pipe(referenceIndex, west), pipe(referenceIndex, north)
	case NorthEastPipe:
		return pipe(referenceIndex, north), pipe(referenceIndex, east)
	case NorthSouthPipe:
		return pipe(referenceIndex, north), pipe(referenceIndex, south)
	case SouthEastPipe:
		return pipe(referenceIndex, east), pipe(referenceIndex, south)
	case S:
		startingPosition_ = referenceIndex
		return referenceIndex, referenceIndex
	case Dot:
		return referenceIndex, referenceIndex
	default:
		return referenceIndex, referenceIndex
	}
}

func atFirstLineOfGrid() bool {
	return len(oneDimensionalGrid_) < gridWidth_
}

func atLeftGridEdge(position int) bool {
	return maths.Mod(position, gridWidth_) == 0
}

func atRightGridEdge(position int) bool {
	return maths.Mod(position+one, gridWidth_) == 0
}

func pipe(referenceIndex int, to direction) int {
	return to.direct(referenceIndex)
}

type direction func(int) int

func (d direction) direct(fromIndex int) int {
	return d(fromIndex)
}

var north direction = func(referenceIndex int) int {
	if atFirstLineOfGrid() {
		return indexOutOfGridBounds
	}
	return positionOfNorthIndex(referenceIndex)
}

func positionOfNorthIndex(referenceIndex int) int {
	return referenceIndex - gridWidth_
}

var south direction = func(referenceIndex int) int {
	if atLastLine_ {
		return indexOutOfGridBounds
	}
	return positionOfSouthNeighbour(referenceIndex)
}

func positionOfSouthNeighbour(referenceIndex int) int {
	return referenceIndex + gridWidth_
}

var west direction = func(referenceIndex int) int {
	if atLeftGridEdge(referenceIndex) {
		return indexOutOfGridBounds
	}
	return positionOfWestNeighbour(referenceIndex)
}

func positionOfWestNeighbour(referenceIndex int) int {
	return referenceIndex - one
}

var east direction = func(referenceIndex int) int {
	if atRightGridEdge(referenceIndex) {
		return indexOutOfGridBounds
	}
	return positionOfEastNeighbour(referenceIndex)
}

func positionOfEastNeighbour(referenceIndex int) int {
	return referenceIndex + one
}

func setGridWidthOnce(fields []byte) {
	if gridWidth_ == 0 {
		gridWidth_ = len(fields)
	}
	return
}

const (
	WestSouthPipe  = '7'
	WestEastPipe   = '-'
	WestNorthPipe  = 'J'
	NorthSouthPipe = '|'
	NorthEastPipe  = 'L'
	SouthEastPipe  = 'F'
	Dot            = '.'
	S              = 'S'
)

type neighbourPair struct {
	neighbour1, neighbour2 int
	pipeShape              byte
}

func newNeighbourPair(neighbour1, neighbour2 int, pipeShape byte) *neighbourPair {
	return &neighbourPair{neighbour1, neighbour2, pipeShape}
}
