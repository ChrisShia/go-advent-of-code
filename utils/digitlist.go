package utils

import (
	"math"
)

type DigitList struct {
	digits []int
}

func newDigitList() *DigitList {
	return &DigitList{}
}

func (dl *DigitList) number() float64 {
	return dl.intConcat()
}

func (dl *DigitList) appendDigits(digits []int) {
	if dl.isEmpty() && beginsWithZero(digits) {
		return
	} else {
		dl.digits = append(dl.digits, digits...)
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
	return len(dl.digits) == 0
}

func beginsWithZero(digits []int) bool {
	return digits[0] == 0
}

func (dl *DigitList) intConcat() float64 {
	var number float64
	orderOfNumber := len(dl.digits) - 1
	base10 := toBase10(orderOfNumber)
	for _, digit := range dl.digits {
		number = base10(digit)
	}
	return number
}

func toBase10(order int) func(digit int) float64 {
	sum := float64(0)
	return func(digit int) float64 {
		sum += math.Pow10(order) * float64(digit)
		order--
		return sum
	}
}
