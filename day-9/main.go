package main

import (
	"bytes"
	"fmt"
	maths "github.com/ChrisShia/math-depot"
	"go-advent-of-code/utils"
	"math"
	"strconv"
)

//part a: 1853145119
//part b: 923

var inputPath_ = "input/day-9.txt"
var sequences_ []*intSequence

func main() {
	sequences_ = make([]*intSequence, 0)
	readInput()
	partA()
	partB()
}

func readInput() {
	utils.Read(inputPath_, nil, func(line []byte) [][]byte { return bytes.Fields(line) }, nil, dayNineLineProcessor)
}

func partA() {
	sum := 0
	for _, sequence := range sequences_ {
		sum += sequence.predictNext()
	}
	fmt.Println("Sum of Predictions : " + strconv.Itoa(sum))
}

func partB() {
	sum := 0
	for _, sequence := range sequences_ {
		sum += sequence.extrapolateBack()
	}
	fmt.Println("Sum of Backward Extrapolations : " + strconv.Itoa(sum))
}

type intSequence struct {
	values                      []int
	firstElementsOfSubsequences []int
}

func (t *intSequence) append(v int) {
	t.values = append(t.values, v)
}

func createSequence(ints ...int) *intSequence {
	seq := &intSequence{values: make([]int, len(ints)), firstElementsOfSubsequences: make([]int, len(ints))}
	for i, v := range ints {
		seq.values[i] = v
	}
	seq.subsequencesFirstElements()
	return seq
}

func (s *intSequence) subsequencesFirstElements() []int {
	s.firstElementsOfSubsequences[0] = s.values[0]
	for i := 1; i < len(s.values); i++ {
		x := s.values[i]
		for j := 0; j < i; j++ {
			x -= s.firstElementsOfSubsequences[j] * maths.NChooseK(i, j)
		}
		s.firstElementsOfSubsequences[i] = x
	}
	return s.firstElementsOfSubsequences
}

func (s *intSequence) predictNext() int {
	predictionIndex := len(s.values)
	prediction := 0
	for i, v := range s.firstElementsOfSubsequences {
		prediction += v * maths.NChooseK(predictionIndex, i)
	}
	return prediction
}

func (s *intSequence) extrapolateBack() int {
	cdIndex, _ := s.commonDifference()
	sum := s.values[0]
	for i := 1; i <= cdIndex; i++ {
		sum += int(math.Pow(float64(-1), float64(i))) * s.firstElementsOfSubsequences[i]
	}
	return sum
}

func (s *intSequence) commonDifference() (int, int) {
	var commonDifference int
	var index int
	ss := s.firstElementsOfSubsequences
	numberOfSubseq := len(ss)
	for i := numberOfSubseq - 1; i >= 0; i-- {
		if ss[i] == 0 {
			continue
		} else {
			commonDifference = ss[i]
			index = i
			break
		}
	}
	return index, commonDifference
}
