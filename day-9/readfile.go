package main

import (
	"go-advent-of-code/dictionary"
)

func dayNineLineProcessor(input [][]byte, last bool) {
	values := make([]int, 0)
	for _, bs := range input {
		toInt := dictionary.BytesToInt(bs)
		values = append(values, toInt)
	}
	sequences_ = append(sequences_, createSequence(values...))
}
