package main

import (
	"bytes"
	"fmt"
	"github.com/ChrisShia/goread/read"
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
	timeSeq := read.BSeq("Time")
	times, extractTimes := newExtractor(timeSeq)
	distanceWord := read.BSeq("Distance")
	distances, extractDistance := newExtractor(distanceWord)
	resultPool := read.FindLinesContainingByteSequences(file, timeSeq, distanceWord)
	extractTimes.From(resultPool)
	extractDistance.From(resultPool)
	return *times, *distances
}

func newExtractor(identifier read.BSeq) (*[]int, read.Extractor) {
	ints := make([]int, 0)
	return &ints, func(pool *read.SearchResultPool) {
		index, _ := pool.SearchFor.Index(identifier)
		for _, sr := range pool.Results {
			if i, ok := sr.Result[index]; ok {
				ints = append(ints, extractValuesFromSearchResult(i, sr.B, identifier)...)
			}
		}
	}
}

func extractValuesFromSearchResult(inLineIndices []int, line []byte, identifier read.BSeq) []int {
	extractedValues := make([]int, 0)
	for _, indexInLine := range inLineIndices {
		dataField := line[indexInLine+len(identifier):]
		fields := bytes.Fields(dataField)
		for _, field := range fields {
			if value, onlyNumerical := toInt(field); onlyNumerical {
				extractedValues = append(extractedValues, value)
			}
		}
	}
	return extractedValues
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
