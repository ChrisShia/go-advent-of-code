package utils

import (
	"bytes"
	"fmt"
)

type CalibrationSequence struct {
	Sequence []CalibrationLine
}

func SumCalibrationSequence(sequence []byte) float64 {
	lines := bytes.Split(sequence, []byte("\n"))
	var sum float64
	for _, line := range lines {
		calibrationLine := NewCalibrationLine(line)
		sum += calibrationLine.Number()
	}
	return sum
}

func DisplayCalibrationValues(sequence []byte) {
	lines := bytes.Split(sequence, []byte("\n"))
	for _, line := range lines {
		calibrationLine := NewCalibrationLine(line)
		fmt.Printf("%55s -> %4v\n", calibrationLine.value, calibrationLine.Number())
	}
}
