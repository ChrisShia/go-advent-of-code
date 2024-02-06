package utils

import (
	"reflect"
	"testing"
)

func TestDigitListAppend(t *testing.T) {
	var tests = []struct {
		name     string
		appendTo DigitList
		append   DigitList
		want     DigitList
	}{
		{"append 456 to 123 gives 123456", DigitList{1, 2, 3}, DigitList{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"append 0456 to 123 gives 1230456", DigitList{1, 2, 3}, DigitList{0, 4, 5, 6}, []int{1, 2, 3, 0, 4, 5, 6}},
		{"append 0456 to {} gives {}", DigitList{}, DigitList{0, 4, 5, 6}, []int{}},
		{"append 6752 to {} gives 6752", DigitList{}, DigitList{6, 7, 5, 2}, []int{6, 7, 5, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.appendTo.appendDigits(tt.append)
			ans := tt.appendTo
			if !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestDigitListIntConcat(t *testing.T) {
	var tests = []struct {
		name  string
		input DigitList
		want  int
	}{
		{"concat {1,2,3} gives 123", DigitList{1, 2, 3}, 123},
		{"concat {6,7,2,5,6} gives 67256", DigitList{6, 7, 2, 5, 6}, 67256},
		{"concat {0,7,2,5,6} gives 67256", DigitList{0, 7, 2, 5, 6}, 7256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.intConcat()
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
