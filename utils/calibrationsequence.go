package utils

import (
	"bytes"
	"fmt"
)

type CalibrationSequence []CalibrationLine

func NewCalibrationSequence(sequence []byte) CalibrationSequence {
	lines := bytes.Split(sequence, []byte("\n"))
	var calibrationSequence CalibrationSequence
	for _, line := range lines {
		calibrationLine := NewCalibrationLine(line)
		calibrationSequence = append(calibrationSequence, calibrationLine)
	}
	return calibrationSequence
}

func SumNumbersInByteSequence(sequence []byte) int {
	lines := bytes.Split(sequence, []byte("\n"))
	var sum int
	for _, line := range lines {
		calibrationLine := NewCalibrationLine(line)
		sum += calibrationLine.Number()
	}
	return sum
}

func DisplayDigitsInByteSequence(sequence []byte) {
	lines := bytes.Split(sequence, []byte("\n"))
	for _, line := range lines {
		calibrationLine := NewCalibrationLine(line)
		calibrationNumber := calibrationLine.Number()
		//s := calibrationLine.String()
		//fmt.Printf("%55s -> %4v\n", calibrationLine, calibrationNumber)
		fmt.Printf("%55s -> %4v\n", calibrationLine.String(), calibrationNumber)
	}
}
