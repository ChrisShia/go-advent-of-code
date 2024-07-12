package main

import (
	"bufio"
	"bytes"
	"go-advent-of-code/dictionary"
	"log"
	"os"
)

const inputPath_ = "/Users/christos/Practise-code/Go/go-advent/input/day-7.txt"
const A byte = 'A'
const K byte = 'K'
const Q byte = 'Q'
const J byte = 'J'
const T byte = 'T'
const _9 byte = '9'
const _8 byte = '8'
const _7 byte = '7'
const _6 byte = '6'
const _5 byte = '5'
const _4 byte = '4'
const _3 byte = '3'
const _2 byte = '2'

type handType

func main() {
	file := openFile()
	defer closeFile(file)
	scanner := bufio.NewScanner(file)
	var line []byte
	var handAndBidSlice [][]byte
	var hands = make([]*hand, 0)
	for scanner.Scan() {
		line = scanner.Bytes()
		handAndBidSlice = bytes.Fields(line)
		bidInt := dictionary.BytesToInt(handAndBidSlice[1])
		hands = append(hands, newHand(handAndBidSlice[0], bidInt))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func openFile() *os.File {
	file, err := os.Open(inputPath_)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func closeFile(file *os.File) {
	func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
}

type hand struct {
	cards []byte
	bid   int
}

func newHand(cards []byte, bid int) *hand {
	return &hand{cards, bid}
}

func (h *hand) equal(o *hand) bool {
	if h == nil {
		return false
	}
	if o == nil {
		return true
	}
	if len(h.cards) != len(o.cards) {
		return false
	}
	for c := 0; c < len(h.cards); c++ {
		if h.cards[c] != o.cards[c] {
			return false
		}
	}
	return true
}

func (h *hand) greater(o *hand) bool {
	if h == nil {
		return false
	}
	if o == nil {
		return true
	}
	return true
}