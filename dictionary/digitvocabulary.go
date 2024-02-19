package dictionary

import (
	"bytes"
	"math"
)

type Digit int

type digitIdentifier []Digit

const (
	ZERO Digit = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	NonDigit = -1
)

var digitByteMap_ = []byte("0123456789")

func (d Digit) String() string {
	if d == -1 {
		return "NonDigit"
	}
	return []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}[d]
}

func (d Digit) Bytes() []byte {
	if d == -1 {
		return []byte{}
	}
	return [][]byte{[]byte("zero"), []byte("one"), []byte("two"), []byte("three"), []byte("four"), []byte("five"), []byte("six"), []byte("seven"), []byte("eight"), []byte("nine")}[d]
}

func (d Digit) Integer() int {
	return int(d)
}

func (d Digit) Byte() byte {
	if d == -1 {
		return '/'
	}
	return digitByteMap_[d]
}

func Values() []Digit {
	return []Digit{ZERO, ONE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE}
}

func (d Digit) IsValidDigit() bool {
	for _, digit := range Values() {
		if d == digit {
			return true
		}
	}
	return false
}

var oneTwoSixChecker_ digitIdentifier = []Digit{ONE, TWO, SIX}
var zeroFourFiveNineChecker_ digitIdentifier = []Digit{ZERO, FOUR, FIVE, NINE}
var threeSevenEightChecker_ digitIdentifier = []Digit{THREE, SEVEN, EIGHT}

func (identifier *digitIdentifier) IdentifyDigit(word []byte) Digit {
	for _, digit := range *identifier {
		bytesOfDigit := digit.Bytes()
		if len(bytesOfDigit) == len(word) {
			if bytes.Contains(bytesOfDigit, word) {
				return digit
			}
		}
	}
	return NonDigit
}

func IsNumericalCharacter(byte byte) bool {
	return byte >= ZERO.Byte() && byte <= NINE.Byte()
}

func ByteToDigit(byte byte) Digit {
	for _, d := range Values() {
		digitByte := d.Byte()
		if digitByte == byte {
			return d
		}
	}
	return NonDigit
}

func ByteToInt(byte byte) int {
	for _, d := range Values() {
		digitByte := d.Byte()
		if digitByte == byte {
			return d.Integer()
		}
	}
	return -1
}

func BytesToInt(bs []byte) int {
	var number int
	orderOfNumber := len(bs) - 1
	parse := byteParser(orderOfNumber)
	for _, b := range bs {
		number = parse(b)
	}
	return number
}

func ConcatenateInts(ints []int) int {
	var number int
	orderOfNumber := len(ints) - 1
	parse := intParser(orderOfNumber)
	for _, b := range ints {
		number = parse(b)
	}
	return number
}

func byteParser(order int) func(b byte) int {
	sum := 0
	return func(b byte) int {
		toInt := ByteToInt(b)
		sum += int(math.Pow10(order)) * toInt
		order--
		return sum
	}
}

func intParser(order int) func(i int) int {
	sum := 0
	return func(i int) int {
		sum += int(math.Pow10(order)) * i
		order--
		return sum
	}
}

func WordToDigitMapper() func(byte byte) Digit {
	var word []byte
	return func(byte byte) Digit {
		word = append(word, byte)
		if len(word) == 1 && IsNumericalCharacter(byte) {
			return ByteToDigit(byte)
		} else if len(word) < 3 {
			return NonDigit
		} else if len(word) < 4 {
			return oneTwoSixChecker_.IdentifyDigit(word)
		} else if len(word) < 5 {
			return zeroFourFiveNineChecker_.IdentifyDigit(word)
		} else if len(word) < 6 {
			return threeSevenEightChecker_.IdentifyDigit(word)
		}
		return NonDigit
	}
}
