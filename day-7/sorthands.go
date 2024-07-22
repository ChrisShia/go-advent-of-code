package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
	"log"
	"sort"
)

//part a: 250946742
//part b:

//Notes:
//1. occurrence map should be lazily instantiated upon request, instead of given as constructor parameter

const inputPath_ = "input/day-7.txt"

type Card byte

func (c *Card) Strength() int {
	switch *c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	default:
		return -1
	}
}

//func (c Card) StrengthJWeakest() int {
//	switch c {
//	case 'A':
//		return 13
//	case 'K':
//		return 12
//	case 'Q':
//		return 11
//	case 'T':
//		return 10
//	case '9':
//		return 9
//	case '8':
//		return 8
//	case '7':
//		return 7
//	case '6':
//		return 6
//	case '5':
//		return 5
//	case '4':
//		return 4
//	case '3':
//		return 3
//	case '2':
//		return 2
//	case 'J':
//		return 1
//	default:
//		return -1
//	}
//}

const HIGH_CARD int = 1
const ONE_PAIR int = 2
const TWO_PAIR int = 3
const THREE_OF_A_KIND int = 4
const FULL_HOUSE int = 5
const FOUR_OF_A_KIND int = 6
const FIVE_OF_A_KIND int = 7

func main() {
	file := utils.OpenFileLogFatal(inputPath_)
	defer utils.CloseFile(file)
	scanner := bufio.NewScanner(file)
	var line []byte
	var handAndBidSlice [][]byte
	var hands = make([]*hand, 0)
	for scanner.Scan() {
		line = scanner.Bytes()
		handAndBidSlice = bytes.Fields(line)
		bid := dictionary.BytesToInt(handAndBidSlice[1])
		cardsByteSlice := handAndBidSlice[0]
		cards := make([]Card, 0)
		typMap := make(map[int]int)
		for _, c := range cardsByteSlice {
			card := Card(c)
			cards = append(cards, card)
			typMap[card.Strength()]++
		}
		hands = append(hands, newHand(cards, bid, typMap))
	}
	byTotalStrength := By(func(h1, h2 *hand) bool { return h1.lessThan(h2) })
	byTotalStrength.SortSliceStable(hands)
	sum := 0
	for _, h := range hands {
		for _, c := range h.cards {
			fmt.Printf("%v ", c.Strength())
		}
		fmt.Printf("%v ", h.bid)
		fmt.Println()
	}
	for key, h := range hands {
		rank := key + 1
		sum += rank * h.bid
	}
	fmt.Printf("Result is : %v\n", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type hand struct {
	cards         []Card
	bid           int
	occurrenceMap map[int]int
}

func (h *hand) highestFreq() int {
	var res = 1
	for _, f := range h.occurrenceMap {
		if f > res {
			res = f
		}
	}
	return res
}

func (h *hand) Type() int {
	return h.handType()
}

func (h *hand) handType() int {
	switch len(h.occurrenceMap) {
	case 5:
		return HIGH_CARD
	case 4:
		return ONE_PAIR
	case 3:
		switch h.highestFreq() {
		case 2:
			return TWO_PAIR
		case 3:
			return THREE_OF_A_KIND
		}
	case 2:
		switch h.highestFreq() {
		case 3:
			return FULL_HOUSE
		case 4:
			return FOUR_OF_A_KIND
		}
	case 1:
		return FIVE_OF_A_KIND
	}
	return 0
}

func newHand(cards []Card, bid int, handType map[int]int) *hand {
	return &hand{cards, bid, handType}
}

func (h *hand) compare(o *hand) int {
	if o == nil {
		return 1
	}
	if h.Type() > o.Type() {
		return 1
	} else if h.Type() < o.Type() {
		return -1
	}
	for c := 0; c < len(h.cards); c++ {
		if h.cards[c].Strength() < o.cards[c].Strength() {
			return -1
		}
		if h.cards[c].Strength() > o.cards[c].Strength() {
			return 1
		}
	}
	return 0
}

func (h *hand) equals(o *hand) bool {
	for _, ch := range h.cards {
		for _, co := range o.cards {
			if ch != co {
				return false
			}
		}
	}
	return true
}

func (h *hand) lessThan(o *hand) bool {
	switch h.compare(o) {
	case -1:
		return true
	default:
		return false
	}
}

type By func(h1, h2 *hand) bool

func (by By) SortSliceStable(hands []*hand) {
	sort.SliceStable(hands, by.cmpFunc(hands))
}

func (by By) cmpFunc(hands []*hand) func(i int, j int) bool {
	return func(i, j int) bool {
		return by(hands[i], hands[j])
	}
}
