package main

import (
	"fmt"
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
	"math"
	"unicode"
)

var gearToPartsMap_ = make(map[int][]int)
var partBuffer_ = make([]int, 0)
var gearIdBuffer_ int
var foundStar_ bool
var input_ = utils.ReadFile("/Users/christos/Practise-code/Go/go-advent/input/engine.txt")
var inputSize = len(input_)

func main() {
	for index, b := range input_ {
		if isNotDigit(b) {
			if len(partBuffer_) > 0 {
				if foundStar_ == true {
					partNumber := dictionary.ConcatenateInts(partBuffer_)
					gearToPartsMap_[gearIdBuffer_] = append(gearToPartsMap_[gearIdBuffer_], partNumber)
					foundStar_ = false
				}
				partBuffer_ = make([]int, 0)
			}
			continue
		}
		partBuffer_ = append(partBuffer_, dictionary.ByteToInt(b))
		if foundStar_ {
			continue
		}
		if isUpperEdge(index) {
			if isLeftCorner(index) {
				searchForStarAroundLeftCorner(index)
				continue
			}
			if isRightCorner(index) {
				searchForStarAroundRightCorner(index)
				continue
			}
			searchForStarAroundUpperEdgePoint(index)
			continue
		}
		if isBottomEdge(index, inputSize) {
			if isLeftCorner(index) {
				searchForStarAroundBottomLeftCorner(index)
				continue
			}
			searchForStarAroundBottomEdge(index)
			continue
		}
		if isLeftEdge(index) {
			searchForStarAroundLeftEdge(index)
			continue
		}
		if isRightEdge(index) {
			searchForStarAroundRightEdge(index)
			continue
		}
		searchForStarAroundInnerPoint(index)
	}
	sum := 0
	for _, partNumbers := range gearToPartsMap_ {
		if len(partNumbers) == 2 {
			sum += partNumbers[0] * partNumbers[1]
		}
	}
	fmt.Printf("The sum of the gear products is : %d", sum)
}

func isLeftCorner(index int) bool {
	return isLeftEdge(index)
}

func isLeftEdge(index int) bool {
	return math.Remainder(float64(index), float64(141)) == 0
}

func isUpperEdge(index int) bool {
	return index < 140
}

func isNotDigit(b byte) bool {
	return !unicode.IsDigit(rune(b))
}

func isBottomEdge(index int, inputSize int) bool {
	return index+141 > inputSize
}

func isStarSymbol(b byte) bool {
	return b == '*'
}

func searchForStarAroundLeftCorner(index int) {
	if isStarSymbol(input_[index+1]) {
		gearIdBuffer_ = index + 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+141]) {
		gearIdBuffer_ = index + 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index+142]) {
		gearIdBuffer_ = index + 142
		foundStar_ = true
	}
}

func isRightCorner(index int) bool {
	return isRightEdge(index)
}

func isRightEdge(index int) bool {
	return math.Remainder(float64(index+2), float64(141)) == 0
}

func searchForStarAroundBottomEdge(index int) {
	if isStarSymbol(input_[index-142]) {
		gearIdBuffer_ = index - 142
		foundStar_ = true
	}
	if isStarSymbol(input_[index-141]) {
		gearIdBuffer_ = index - 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index-140]) {
		gearIdBuffer_ = index - 140
		foundStar_ = true
	}
	if isStarSymbol(input_[index-1]) {
		gearIdBuffer_ = index - 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+1]) {
		gearIdBuffer_ = index + 1
		foundStar_ = true
	}
}

func searchForStarAroundRightCorner(index int) {
	if isStarSymbol(input_[index-1]) {
		gearIdBuffer_ = index - 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+140]) {
		gearIdBuffer_ = index + 140
		foundStar_ = true
	}
	if isStarSymbol(input_[index+141]) {
		gearIdBuffer_ = index + 141
		foundStar_ = true
	}
}

func searchForStarAroundUpperEdgePoint(index int) {
	if isStarSymbol(input_[index-1]) {
		gearIdBuffer_ = index - 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+1]) {
		gearIdBuffer_ = index + 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+140]) {
		gearIdBuffer_ = index + 140
		foundStar_ = true
	}
	if isStarSymbol(input_[index+141]) {
		gearIdBuffer_ = index + 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index+142]) {
		gearIdBuffer_ = index + 142
		foundStar_ = true
	}
}

func searchForStarAroundBottomLeftCorner(index int) {
	if isStarSymbol(input_[index-141]) {
		gearIdBuffer_ = index - 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index-140]) {
		gearIdBuffer_ = index - 140
		foundStar_ = true
	}
	if isStarSymbol(input_[index+1]) {
		gearIdBuffer_ = index + 1
		foundStar_ = true
	}
}

func searchForStarAroundLeftEdge(index int) {
	if isStarSymbol(input_[index-141]) {
		gearIdBuffer_ = index - 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index-140]) {
		gearIdBuffer_ = index - 140
		foundStar_ = true
	}
	if isStarSymbol(input_[index+1]) {
		gearIdBuffer_ = index + 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+141]) {
		gearIdBuffer_ = index + 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index+142]) {
		gearIdBuffer_ = index + 142
		foundStar_ = true
	}
}

func searchForStarAroundRightEdge(index int) {
	if isStarSymbol(input_[index-142]) {
		gearIdBuffer_ = index - 142
		foundStar_ = true
	}
	if isStarSymbol(input_[index-141]) {
		gearIdBuffer_ = index - 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index-1]) {
		gearIdBuffer_ = index - 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+140]) {
		gearIdBuffer_ = index + 140
		foundStar_ = true
	}
	if isStarSymbol(input_[index+141]) {
		gearIdBuffer_ = index + 141
		foundStar_ = true
	}
}

func searchForStarAroundInnerPoint(index int) {
	if isStarSymbol(input_[index-142]) {
		gearIdBuffer_ = index - 142
		foundStar_ = true
	}
	if isStarSymbol(input_[index-141]) {
		gearIdBuffer_ = index - 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index-140]) {
		gearIdBuffer_ = index - 140
		foundStar_ = true
	}
	if isStarSymbol(input_[index-1]) {
		gearIdBuffer_ = index - 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+1]) {
		gearIdBuffer_ = index + 1
		foundStar_ = true
	}
	if isStarSymbol(input_[index+140]) {
		gearIdBuffer_ = index + 140
		foundStar_ = true
	}
	if isStarSymbol(input_[index+141]) {
		gearIdBuffer_ = index + 141
		foundStar_ = true
	}
	if isStarSymbol(input_[index+142]) {
		gearIdBuffer_ = index + 142
		foundStar_ = true
	}
}
