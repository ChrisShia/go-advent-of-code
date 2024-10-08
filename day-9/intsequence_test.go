package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		name     string
		seq      []int
		expected []int
	}{
		{"", []int{10, 13, 16, 21, 30, 45, 68}, []int{}},
		{"", []int{13, 25, 60, 142, 312, 634, 1201, 2141, 3623, 5863, 9130, 13752, 20122, 28704, 40039, 54751, 73553, 97253, 126760, 163090, 207372}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intSeq := createSequence(tt.seq...)
			ans := intSeq.subsequencesFirstElements()
			fmt.Println(ans)
		})
	}
}

func Test_predictNext(t *testing.T) {
	var tests = []struct {
		name     string
		seq      []int
		expected int
	}{
		{"", []int{10, 13, 16, 21, 30, 45}, 68},
		{"", []int{13, 25, 60, 142, 312, 634, 1201, 2141, 3623, 5863, 9130, 13752, 20122, 28704, 40039, 54751, 73553, 97253, 126760, 163090, 207372}, 260854},
		{"", []int{9, 5, 1, -3, -7, -11, -15, -19, -23, -27, -31, -35, -39, -43, -47, -51, -55, -59, -63, -67, -71}, -75},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intSeq := createSequence(tt.seq...)
			ans := intSeq.predictNext()
			if ans != tt.expected {
				t.Errorf("PredictNext() = %d, want %d", ans, tt.expected)
			}
		})
	}
}

func Test_predictPrevious(t *testing.T) {
	var tests = []struct {
		name     string
		seq      []int
		expected int
	}{
		{"", []int{10, 13, 16, 21, 30, 45}, 5},
		{"", []int{0, 3, 6, 9, 12, 15}, -3},
		{"", []int{1, 3, 6, 10, 15, 21}, 0},
		{"", []int{9, 5, 1, -3, -7, -11, -15, -19, -23, -27, -31, -35, -39, -43, -47, -51, -55, -59, -63, -67, -71}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intSeq := createSequence(tt.seq...)
			ans := intSeq.extrapolateBack()
			if ans != tt.expected {
				t.Errorf("PredictPrevious() = %d, want %d", ans, tt.expected)
			}
		})
	}
}
