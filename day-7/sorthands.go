package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
	"log"
	"os"
	"sort"
)

//part a: 250946742
//part b: 251824095

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

func (c *Card) StrengthJWeakest() int {
	switch *c {
	case 'A':
		return 13
	case 'K':
		return 12
	case 'Q':
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
	case 'J':
		return 1
	default:
		return -1
	}
}

const HighCard int = 1
const OnePair int = 2
const TwoPair int = 3
const ThreeOfAKind int = 4
const FullHouse int = 5
const FourOfAKind int = 6
const FiveOfAKind int = 7

func main() {
	file := utils.OpenFileLogFatal(inputPath_)
	defer utils.CloseFile(file)
	hands := readHandsFromFile(file)
	//typeRule := TypeRule(func(h *hand) map[Card]int { return h.cardFrequencyMap })
	typeRule := TypeRule(func(h *hand) map[Card]int { return applyTrumpCard(h.cardFrequencyMap, Card('J')) })
	strengthRule := StrengthRule(func(c *Card) int { return c.StrengthJWeakest() })
	gameRules := &GameRules{typeRule, strengthRule}
	byTotalStrength := By(func(h1, h2 *hand) bool { return gameRules.lessThan(h1, h2) })
	byTotalStrength.SortSliceStable(hands)
	sum := 0
	for _, h := range hands {
		for _, c := range h.cards {
			fmt.Printf("%v ", gameRules.strengthRule.Strength(&c))
		}
		fmt.Printf("%v ", h.bid)
		fmt.Println()
	}
	for key, h := range hands {
		rank := key + 1
		sum += rank * h.bid
	}
	fmt.Printf("Result is : %v\n", sum)
}

func readHandsFromFile(file *os.File) []*hand {
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
		cardFrequencyMap := make(map[Card]int)
		for _, c := range cardsByteSlice {
			card := Card(c)
			cards = append(cards, card)
			cardFrequencyMap[card]++
		}
		hands = append(hands, newHand(cards, bid, cardFrequencyMap))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return hands
}

func applyTrumpCard(cardMap map[Card]int, trump Card) map[Card]int {
	trumps := cardMap[trump]
	if trumps == 0 {
		return cardMap
	}
	trumplessMap := make(map[Card]int)
	mostFreqCard, _ := mostFrequentButExcludeTrump(cardMap, trump)
	if mostFreqCard == 0 {
		return cardMap
	}
	virtualFrequency := cardMap[mostFreqCard] + trumps
	for c, f := range cardMap {
		if c == trump {
			continue
		}
		if c == mostFreqCard {
			trumplessMap[c] = virtualFrequency
			continue
		}
		trumplessMap[c] = f
	}
	return trumplessMap
}

func mostFrequentButExcludeTrump(cardMap map[Card]int, trump Card) (Card, int) {
	var freq = 0
	var mostFreqCard Card
	for c, f := range cardMap {
		if c == trump {
			continue
		}
		if f > freq {
			freq = f
			mostFreqCard = c
		}
	}
	return mostFreqCard, freq
}

type hand struct {
	cards            []Card
	bid              int
	cardFrequencyMap map[Card]int
}

func highestFreq(cardMap map[Card]int) int {
	_, freq := highestFrequencyCard(cardMap)
	return freq
}

func highestFrequencyCard(cardMap map[Card]int) (Card, int) {
	var highFreq = 0
	var highestFreqCard Card
	for c, f := range cardMap {
		if f > highFreq {
			highFreq = f
			highestFreqCard = c
		}
	}
	return highestFreqCard, highFreq
}

func handType(cardMap map[Card]int) int {
	switch len(cardMap) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		switch highestFreq(cardMap) {
		case 2:
			return TwoPair
		case 3:
			return ThreeOfAKind
		}
	case 2:
		switch highestFreq(cardMap) {
		case 3:
			return FullHouse
		case 4:
			return FourOfAKind
		}
	case 1:
		return FiveOfAKind
	}
	return 0
}

func newHand(cards []Card, bid int, handType map[Card]int) *hand {
	return &hand{cards, bid, handType}
}

func (gr *GameRules) compare(h1, h2 *hand) int {
	if h1 == nil {
		return -1
	}
	if h2 == nil {
		return 1
	}
	typeOfHand1 := gr.typeRule.Type(h1)
	typeOfHand2 := gr.typeRule.Type(h2)
	if typeOfHand1 > typeOfHand2 {
		return 1
	} else if typeOfHand1 < typeOfHand2 {
		return -1
	}
	for c := 0; c < len(h1.cards); c++ {
		strength1 := gr.strengthRule.Strength(&h1.cards[c])
		strength2 := gr.strengthRule.Strength(&h2.cards[c])
		if strength1 < strength2 {
			return -1
		}
		if strength1 > strength2 {
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

func (gr *GameRules) lessThan(h, o *hand) bool {
	switch gr.compare(h, o) {
	case -1:
		return true

	default:
		return false
	}
}

type GameRules struct {
	typeRule     TypeRule
	strengthRule StrengthRule
}

type TypeRule func(h *hand) map[Card]int

func (tr *TypeRule) Type(h *hand) int {
	ruleBasedCardMap := (*tr)(h)
	return handType(ruleBasedCardMap)
}

type StrengthRule func(c *Card) int

func (sr *StrengthRule) Strength(c *Card) int {
	return (*sr)(c)
}

type By func(h1, h2 *hand) bool

func (by *By) SortSliceStable(hands []*hand) {
	sort.SliceStable(hands, by.compareFunc(hands))
}

func (by *By) compareFunc(hands []*hand) func(i int, j int) bool {
	return func(i, j int) bool {
		return (*by)(hands[i], hands[j])
	}
}
