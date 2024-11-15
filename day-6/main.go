package main

import (
	"bytes"
	"fmt"
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
	"math"
)

// Answer a: 131376
// Answer b: 34123437

const inputPath_ = "input/day-6.txt"

func main() {
	times, distances := extractTimesAndDistances()
	product := 1
	for i := 0; i < len(times); i++ {
		product *= winningMoves(times[i], distances[i])
	}
	fmt.Println(product)
}

func winningMoves(t int, d int) int {
	quadSol1 := (float64(t) + determinant(t, d)) / 2
	quadSol2 := (float64(t) - determinant(t, d)) / 2
	if quadSol1 > float64(int(quadSol1)) {
		quadSol1 = math.Round(quadSol1 - 0.5)
	} else {
		quadSol1--
	}
	if quadSol2 > float64(int(quadSol2)) {
		quadSol2 = math.Round(quadSol2 - 0.5)
	}
	rootDiff := quadSol1 - quadSol2
	return int(rootDiff)
}

func determinant(t int, d int) float64 {
	return math.Sqrt(math.Pow(float64(t), 2) - 4*float64(d))
}

func extractTimesAndDistances() ([]int, []int) {
	file := utils.OpenFileLogFatal(inputPath_)
	defer utils.CloseFile(file)
	timeSeq := utils.BSeq("Time")
	times, extractTimes := newExtractor(timeSeq)
	distanceWord := utils.BSeq("Distance")
	distances, extractDistance := newExtractor(distanceWord)
	resultPool := utils.FindLinesContainingByteSequences(file, timeSeq, distanceWord)
	extractTimes.From(resultPool)
	extractDistance.From(resultPool)
	return *times, *distances
}

func newExtractor(identifier utils.BSeq) (*[]int, utils.Extractor) {
	ints := make([]int, 0)
	return &ints, func(pool *utils.SearchResultPool) {
		index, _ := pool.SearchFor.Index(identifier)
		for _, sr := range pool.Results {
			if i, ok := sr.Result[index]; ok {
				for _, indexInLine := range i {
					dataField := sr.B[indexInLine+len(identifier):]
					fields := bytes.Fields(dataField)
					for _, field := range fields {
						if value, onlyNumerical := toInt(field); onlyNumerical {
							ints = append(ints, value)
						}
					}
				}
			}
		}
	}
}

func toInt(field []byte) (int, bool) {
	return dictionary.BytesToInt(field), containsOnlyNumericalChars(field)
}

func containsOnlyNumericalChars(field []byte) bool {
	for _, f := range field {
		if !dictionary.IsNumericalCharacter(f) {
			return false
		}
	}
	return true
}
