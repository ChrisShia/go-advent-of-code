package main

import (
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
	"math"
	"unicode"
)

var partNumberSlice_ = make([]int, 0)
var partBuffer_ = make([]int, 0)
var input_ = utils.ReadFile("engine/engine.txt")
var inputSize = len(input_)
var foundSymbol_ bool

func main() {
	var partNumberSum_ int
	for index, b := range input_ {
		if isNotDigit(b) {
			if len(partBuffer_) > 0 {
				if foundSymbol_ == true {
					intPartBuffer := dictionary.ConcatenateInts(partBuffer_)
					partNumberSlice_ = append(partNumberSlice_, intPartBuffer)
					partNumberSum_ += intPartBuffer
					foundSymbol_ = false
				}
				partBuffer_ = make([]int, 0)
			}
			continue
		}
		partBuffer_ = append(partBuffer_, dictionary.ByteToInt(b))
		if foundSymbol_ {
			continue
		}
		if isUpperEdge(index) {
			if isLeftCorner(index) {
				searchForSymbolAroundLeftCorner(index)
				continue
			}
			if isRightCorner(index) {
				searchForSymbolAroundRightCorner(index)
				continue
			}
			searchForSymbolAroundUpperEdgePoint(index)
			continue
		}
		if isBottomEdge(index, inputSize) {
			if isLeftCorner(index) {
				searchForSymbolAroundBottomLeftCorner(index)
				continue
			}
			searchForSymbolAroundBottomEdge(index)
			continue
		}
		if isLeftEdge(index) {
			searchForSymbolAroundLeftEdge(index)
			continue
		}
		if isRightEdge(index) {
			searchForSymbolAroundRightEdge(index)
			continue
		}
		searchForSymbolAroundInnerPoint(index)
	}
}

func searchForSymbolAroundRightEdge(index int) {
	if isSymbol(input_[index-142]) || isSymbol(input_[index-141]) || isSymbol(input_[index-1]) || isSymbol(input_[index+140]) || isSymbol(input_[index+141]) {
		foundSymbol_ = true
	}
}

func searchForSymbolAroundRightCorner(index int) {
	if isSymbol(input_[index-1]) || isSymbol(input_[index+140]) || isSymbol(input_[index+141]) {
		foundSymbol_ = true
	}
}

func isRightCorner(index int) bool {
	return isRightEdge(index)
}

func isRightEdge(index int) bool {
	return math.Remainder(float64(index+2), float64(141)) == 0
}

func searchForSymbolAroundInnerPoint(index int) {
	if isSymbol(input_[index-142]) || isSymbol(input_[index-141]) || isSymbol(input_[index-140]) || isSymbol(input_[index-1]) || isSymbol(input_[index+1]) || isSymbol(input_[index+140]) || isSymbol(input_[index+141]) || isSymbol(input_[index+142]) {
		foundSymbol_ = true
	}
}

func searchForSymbolAroundBottomLeftCorner(index int) {
	if isSymbol(input_[index-141]) || isSymbol(input_[index-140]) || isSymbol(input_[index+1]) {
		foundSymbol_ = true
	}
}

func searchForSymbolAroundBottomEdge(index int) {
	if isSymbol(input_[index-142]) || isSymbol(input_[index-141]) || isSymbol(input_[index-140]) || isSymbol(input_[index-1]) || isSymbol(input_[index+1]) {
		foundSymbol_ = true
	}
}

func isNotDigit(b byte) bool {
	return !unicode.IsDigit(rune(b))
}

func searchForSymbolAroundLeftEdge(index int) {
	if isSymbol(input_[index-141]) || isSymbol(input_[index-140]) || isSymbol(input_[index+1]) || isSymbol(input_[index+141]) || isSymbol(input_[index+142]) {
		foundSymbol_ = true
	}
}

func isBottomEdge(index int, inputSize int) bool {
	return index+141 > inputSize
}

func searchForSymbolAroundLeftCorner(index int) {
	if isSymbol(input_[index+1]) || isSymbol(input_[index+141]) || isSymbol(input_[index+142]) {
		foundSymbol_ = true
	}
}

func isSymbol(b byte) bool {
	return !isDotOrDigit(b)
}

func isDotOrDigit(b byte) bool {
	return unicode.IsDigit(rune(b)) || b == '.'
}

func isLeftCorner(index int) bool {
	return isLeftEdge(index)
}

func isUpperEdge(index int) bool {
	return index < 140
}

func isLeftEdge(index int) bool {
	return math.Remainder(float64(index), float64(141)) == 0
}

func searchForSymbolAroundUpperEdgePoint(index int) {
	if isSymbol(input_[index-1]) || isSymbol(input_[index+1]) || isSymbol(input_[index+140]) || isSymbol(input_[index+141]) || isSymbol(input_[index+142]) {
		foundSymbol_ = true
	}
}
