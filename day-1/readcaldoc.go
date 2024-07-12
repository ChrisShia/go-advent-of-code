package main

import (
	"fmt"
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
)

func main() {
	input := utils.ReadFile("/Users/christos/Practise-code/Go/go-advent/input/day-1.txt")
	utils.DisplayDigitsInByteSequence(input)
	fmt.Println(utils.SumNumbersInByteSequence(input))
}

func experimentWithDigitDictionary() {
	allDigits := dictionary.Values()
	fmt.Println(allDigits)
}

func experimentWithByteSlices() {
	digitsToByteSliceDictionary := []byte("0123456789")
	var wordDigitsToBytesDictionary [][]byte
	wordDigitsToBytesDictionary = append(wordDigitsToBytesDictionary)
	fmt.Println(digitsToByteSliceDictionary[0])
	fmt.Println(wordDigitsToBytesDictionary[0])
}
