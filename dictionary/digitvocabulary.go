package dictionary

import "bytes"

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

func ByteToNumerical(byte byte) Digit {
	for _, d := range Values() {
		digitByte := d.Byte()
		if digitByte == byte {
			return d
		}
	}
	return NonDigit
}

func WordToDigitMapper() func(byte byte) Digit {
	var word []byte
	return func(byte byte) Digit {
		word = append(word, byte)
		if len(word) == 1 && IsNumericalCharacter(byte) {
			return ByteToNumerical(byte)
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
