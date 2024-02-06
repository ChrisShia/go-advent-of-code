package utils

import (
	"go-advent-of-code/dictionary"
	"reflect"
	"testing"
)

func TestCalibrationLineExtractDigitSliceExcludeWords(t *testing.T) {
	var tests = []struct {
		name  string
		input CalibrationLine
		want  []dictionary.Digit
	}{
		{"9nine90 is 990", NewCalibrationLine([]byte("9nine90")), []dictionary.Digit{9, 9, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.extractDigitSliceExcludeWords()
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestCalibrationLineExtractDigitSliceIncludeWords(t *testing.T) {
	var tests = []struct {
		name  string
		input CalibrationLine
		want  []int
	}{
		{"9nine90 is 9990", NewCalibrationLine([]byte("9nine90")), []int{9, 9, 9, 0}},
		{"9nine90six is 99906", NewCalibrationLine([]byte("9nine90six")), []int{9, 9, 9, 0, 6}},
		{"eightalskdgughaeight56hui is 8856", NewCalibrationLine([]byte("eightalskdgughaeight56hui")), []int{8, 8, 5, 6}},
		{"abcone2threexyz is 123", NewCalibrationLine([]byte("abcone2threexyz")), []int{1, 2, 3}},
		{"3twonep is 31", NewCalibrationLine([]byte("3twonep")), []int{3, 2, 1}},
		{"aofiveightiusfghsgeightwo is 82", NewCalibrationLine([]byte("aofiveightiusfghsgeightwo")), []int{5, 8, 8, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.extractDigitSliceIncludeWords()
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestCalibrationLine_Number(t *testing.T) {
	var tests = []struct {
		name  string
		input CalibrationLine
		want  int
	}{
		{"9liuasdfgbasdgfjkhb09zero0 is 90", NewCalibrationLine([]byte("9liuasdfgbasdgfjkhb09zero0")), 90},
		{"6iaushdf7one9nine00six is 66", NewCalibrationLine([]byte("6iaushdf7one9nine00six")), 66},
		{"aoiusfghsgeighttwo is 82", NewCalibrationLine([]byte("aoiusfghsgeighttwo")), 82},
		{"aoiusfsevenineghsgeightwo is 72", NewCalibrationLine([]byte("aoiusfsevenineghsgeightwo")), 72},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.Number()
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
