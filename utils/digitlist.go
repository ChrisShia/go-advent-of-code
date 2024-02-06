package utils

import (
	"math"
)

type DigitList []int

func newDigitList() *DigitList {
	return &DigitList{}
}

func (dl *DigitList) number() int {
	return dl.intConcat()
}

func (dl *DigitList) appendDigits(digits []int) {
	if dl.isEmpty() && beginsWithZero(digits) {
		return
	} else {
		*dl = append(*dl, digits...)
	}
}

func (dl *DigitList) appendDigitN(digit, n int) {
	intSliceToAppend := createIntSliceAllElementsInitializedToDigit(digit, n)
	dl.appendDigits(intSliceToAppend)
}

func createIntSliceAllElementsInitializedToDigit(digit int, n int) []int {
	digitsToAppend := make([]int, n)
	for i := 0; i < n; i++ {
		digitsToAppend[i] = digit
	}
	return digitsToAppend
}

func (dl *DigitList) isEmpty() bool {
	return len(*dl) == 0
}

func beginsWithZero(digits []int) bool {
	return digits[0] == 0
}

func (dl *DigitList) intConcat() int {
	var number int
	orderOfNumber := len(*dl) - 1
	base10 := toBase10(orderOfNumber)
	for _, digit := range *dl {
		number = base10(digit)
	}
	return number
}

func toBase10(order int) func(digit int) int {
	sum := 0
	return func(digit int) int {
		sum += int(math.Pow10(order)) * digit
		order--
		return sum
	}
}
