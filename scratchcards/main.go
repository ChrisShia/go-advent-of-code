package main

import (
	"bytes"
	"fmt"
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
	"math"
	"sort"
)

// Correct answer for 1st part: 24160

var input_ = utils.ReadFile("scratchcards/scratchcards.txt")
var sumOfPointsFromAllCards_ int
var noOfCardEntitiesSlice_ []int

func main() {
	var found = true
	var card []byte
	var remainingCards = input_
	//aggregator := cardPointAggregator()
	copiesPopulator := cardCopiesPopulator()
	for found {
		card, remainingCards, found = bytes.Cut(remainingCards, []byte{'\n'})
		copiesPopulator(card)
	}
	sum := 0
	for _, n := range noOfCardEntitiesSlice_ {
		sum += n
	}
	fmt.Println(sum)
}

func cardPointAggregator() func(card []byte) {
	return cardProcessor(splitOnVerticalLine(), collectPointsPowersOfTwo)
}

func cardCopiesPopulator() func(card []byte) {
	noOfCardEntitiesSlice_ = make([]int, 0)
	return cardProcessor(splitOnVerticalLine(), collectCopies)
}

func cardProcessor(cardSplitter func(card []byte) (winning, actual []int),
	collector func(winning, actual []int, id int)) func(card []byte) {
	cardId := 1
	return func(card []byte) {
		_, winningAndActual, _ := bytes.Cut(card, []byte(": "))
		winning, actual := cardSplitter(winningAndActual)
		collector(winning, actual, cardId)
		cardId++
	}
}

func splitOnVerticalLine() func(card []byte) ([]int, []int) {
	return func(card []byte) ([]int, []int) {
		beforeLine, afterLine, _ := bytes.Cut(card, []byte(" | "))
		winningNumbers := bytesToSortedInts(beforeLine)
		actualNumbers := bytesToSortedInts(afterLine)
		return winningNumbers, actualNumbers
	}
}

func bytesToSortedInts(nums []byte) []int {
	a := bytes.Fields(nums)
	numbers := make([]int, 0)
	for _, bs := range a {
		numbers = append(numbers, dictionary.BytesToInt(bs))
	}
	sort.Ints(numbers)
	return numbers
}

func collectPointsPowersOfTwo(sortedWinningNumbers, sortedActualNumbers []int, cardId int) {
	commonNumbers := commonInts(&sortedWinningNumbers, &sortedActualNumbers)
	sumOfPointsFromAllCards_ += int(math.Pow(float64(2), float64(len(*commonNumbers)-1)))
}

func collectCopies(sortedWinningNumbers, sortedActualNumbers []int, cardId int) {
	commonNumbers := commonInts(&sortedWinningNumbers, &sortedActualNumbers)
	noOfMatches := len(*commonNumbers)
	var noOfEntitiesForCurrentId int
	if cardId == 1 || cardId > len(noOfCardEntitiesSlice_) {
		noOfCardEntitiesSlice_ = append(noOfCardEntitiesSlice_, 1)
		noOfEntitiesForCurrentId = 1
		for i := 1; i <= noOfMatches; i++ {
			noOfCardEntitiesSlice_ = append(noOfCardEntitiesSlice_, 1)
		}
	} else {
		noOfCardEntitiesSlice_[cardId-1]++
		noOfEntitiesForCurrentId = noOfCardEntitiesSlice_[cardId-1]
		for i := 1; i <= noOfMatches; i++ {
			if cardId+i <= len(noOfCardEntitiesSlice_) {
				noOfCardEntitiesSlice_[cardId+i-1] += noOfEntitiesForCurrentId
				continue
			}
			noOfCardEntitiesSlice_ = append(noOfCardEntitiesSlice_, noOfEntitiesForCurrentId)
		}
	}

}

func commonInts(sortedA, sortedB *[]int) *[]int {
	commonNumbers := make([]int, 0)
	for i, j := 0, 0; i < len(*sortedA) && j < len(*sortedB); {
		if (*sortedA)[i] == (*sortedB)[j] {
			commonNumbers = append(commonNumbers, (*sortedB)[j])
		}
		if (*sortedA)[i] > (*sortedB)[j] {
			j++
		} else {
			i++
		}
	}
	return &commonNumbers
}
