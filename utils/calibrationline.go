package utils

import (
	"fmt"
	"go-advent-of-code/dictionary"
)

type CalibrationLine struct {
	value     []byte
	digitList *DigitList
}

func NewCalibrationLine(byteLine []byte) CalibrationLine {
	return CalibrationLine{byteLine, newDigitList()}
}

func (cl *CalibrationLine) String() string {
	return fmt.Sprintf("%s", string(cl.value))
}

func (cl *CalibrationLine) Number() int {
	cl.identifyFirstAndLastDigitsInCalLine()
	return cl.digitList.number()
}

func (cl *CalibrationLine) identifyFirstAndLastDigitsInCalLine() {
	digitsFound := cl.extractDigitSliceIncludeWords()
	numberOfDigitsFound := len(digitsFound)
	if numberOfDigitsFound == 0 {
		return
	}
	if numberOfDigitsFound == 1 {
		cl.appendToDigitListN(digitsFound)
		return
	}
	if numberOfDigitsFound > 2 {
		cl.appendOnlyFirstAndLastDigits(digitsFound, numberOfDigitsFound)
		return
	}
	cl.appendToDigitList(digitsFound)
}

func (cl *CalibrationLine) appendToDigitListN(digitsFound []int) {
	cl.digitList.appendDigitN(digitsFound[0], 2)
}

func (cl *CalibrationLine) appendOnlyFirstAndLastDigits(digitsFound []int, numberOfDigitsFound int) {
	digitsToAppend := []int{digitsFound[0], digitsFound[numberOfDigitsFound-1]}
	cl.appendToDigitList(digitsToAppend)
}

func (cl *CalibrationLine) appendToDigitList(digitsToAppend []int) {
	cl.digitList.appendDigits(digitsToAppend)
}

func (cl *CalibrationLine) extractDigitSliceExcludeWords() []dictionary.Digit {
	var digitsFound []dictionary.Digit
	for _, char := range cl.value {
		if isNumericalCharacter(char) {
			integerRepresentation := dictionary.ByteToNumerical(char)
			digitsFound = append(digitsFound, integerRepresentation)
		}
	}
	return digitsFound
}

func isNumericalCharacter(char byte) bool {
	return dictionary.IsNumericalCharacter(char)
}

func (cl *CalibrationLine) extractDigitSliceIncludeWords() []int {
	var digitsFound []int
	for index := 0; index < len(cl.value); index++ {
		var byteBuffer []byte
		remainingSlice := len(cl.value) - index
		if remainingSlice >= 5 {
			byteBuffer = cl.value[index : index+5]
		} else {
			byteBuffer = cl.value[index:]
		}
		mapper := dictionary.WordToDigitMapper()
		for _, b := range byteBuffer {
			digit := mapper(b)
			if digit.IsValidDigit() {
				digitsFound = append(digitsFound, digit.Integer())
				break
			}
		}
	}
	return digitsFound
}
