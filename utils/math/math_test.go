package math

import "testing"

func Test_HCF(t *testing.T) {
	var tests = []struct {
		name     string
		i, j     int
		expected int
	}{
		{"", 5, 10, 5},
		{"", 0, 10, 0},
		{"", 0, -2, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := HCF(tt.i, tt.j)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_GCD(t *testing.T) {
	var tests = []struct {
		name     string
		i, j     int
		expected int
	}{
		{"", 5, 10, 5},
		{"", 0, 10, 10},
		{"", 10, 0, 10},
		{"", 0, -2, -2},
		//TODO: check this, write answer should be -2
		{"", 4, -2, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := GCD(tt.i, tt.j)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_LCM(t *testing.T) {
	var tests = []struct {
		name     string
		i, j     int
		expected int
	}{
		{"", 5, 10, 10},
		{"", 5, 14, 70},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := LCM(tt.i, tt.j)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_Lcm(t *testing.T) {
	var tests = []struct {
		name     string
		i        []int
		expected int
	}{
		{"", []int{5, 10}, 10},
		{"", []int{5, 14}, 70},
		{"", []int{5, 14, 25}, 350},
		{"", []int{1, 3, 4, 5, 7}, 420},
		{"", []int{20659, 20093, 14999, 17263, 22357, 16697}, 22103062509257},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Lcm(tt.i...)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}
