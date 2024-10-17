package maths

import (
	"fmt"
	"testing"
	"time"
)

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

func Test_Mod(t *testing.T) {
	var tests = []struct {
		name     string
		i, j     int
		expected int
	}{
		{"", 10, 2, 0},
		{"", 10, 3, 1},
		{"", -4, 3, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Mod(tt.i, tt.j)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_PowOfTwo(t *testing.T) {
	var tests = []struct {
		name     string
		pow      int
		expected int
	}{
		{"2 to the pow of 0", 0, 1},
		{"2 to the pow of 0", 1, 2},
		{"2 to the pow of 0", 3, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := BitShiftPowOfTwo(tt.pow)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_Factorial(t *testing.T) {
	var tests = []struct {
		name     string
		n        int
		expected int
	}{
		{"factorial of 6", 6, 720},
		{"factorial of 20", 20, 2432902008176640000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := _factorial(tt.n)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_Time_Factorial(t *testing.T) {
	start := time.Now()
	f := _factorial(21)
	tm := time.Now()
	elapsed := tm.Sub(start)
	fmt.Println(f, " : ", elapsed)
}

func Test_Time_PowOFTwo(t *testing.T) {
	start := time.Now()
	f := PowOfTwo(20)
	tm := time.Now()
	elapsed := tm.Sub(start)
	fmt.Println(f, " : ", elapsed)
}

func Test_Time_BitShiftPowOfTwo(t *testing.T) {
	start := time.Now()
	f := BitShiftPowOfTwo(20)
	tm := time.Now()
	elapsed := tm.Sub(start)
	fmt.Println(f, " : ", elapsed)
}

func Test_sequentialMultiplicationFromTo(t *testing.T) {
	var tests = []struct {
		name     string
		from     int
		to       int
		expected int
	}{
		{"seq mul from 3 to 6", 3, 6, 360},
		{"", 1, 2, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := sequentialMultiplicationFromTo(tt.from, tt.to)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_NChooseK(t *testing.T) {
	var tests = []struct {
		name     string
		n        int
		k        int
		expected int
	}{
		{"6 choose 0", 6, 0, 1},
		{"6 choose 2", 6, 2, 15},
		{"20 choose 2", 20, 2, 190},
		{"20 choose 5", 20, 5, 15504},
		{"20 choose 10", 20, 10, 184756},
		{"30 choose 2", 30, 2, 435},
		{"30 choose 10", 30, 10, 30045015},
		{"30 choose 15", 30, 15, 155117520},
		{"30 choose 20", 30, 20, 30045015},
		{"30 choose 25", 30, 25, 155117520},
		{"40 choose 15", 40, 15, 40225345056},
		{"40 choose 20", 40, 20, 137846528820},
		{"50 choose 20", 50, 20, 47129212243960},
		{"50 choose 25", 50, 25, 126410606437752},
		{"60 choose 30", 60, 30, 118264581564861424},
		{"75 choose 35", 75, 30, 7116179530141809008},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := NChooseK(tt.n, tt.k)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_nChooseK(t *testing.T) {
	var tests = []struct {
		name     string
		n        int
		k        int
		expected int
	}{
		{"1 choose 0", 1, 0, 1},
		{"2 choose 1", 2, 1, 2},
		{"2 choose 1", 3, 2, 3},
		{"2 choose 1", 20, 19, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := nChooseK(tt.n, tt.k)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_divisionBySimplification(t *testing.T) {
	var tests = []struct {
		name          string
		numeratorFrom int
		numeratorTo   int
		denominator   int
		expected      int
	}{
		{"", 2, 2, 1, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := divisionBySimplification(tt.numeratorFrom, tt.numeratorTo, tt.denominator)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}

func Test_simplifyNumeratorDenominatorSequences(t *testing.T) {
	var tests = []struct {
		name  string
		numer []int
		denom []int
		//expectedNum []int
		//expectedDen []int
	}{
		{"", []int{6, 7, 8, 9}, []int{1, 2, 3, 4}},
		{"", []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num, den := simplifyNumeratorDenominatorSequences(tt.numer, tt.denom)
			fmt.Println(num)
			fmt.Println(den)
			//if num != tt.expectedNum {
			//	t.Errorf("got %v, want %v", ans, tt.expected)
			//}
		})
	}
}

func Test_multiply(t *testing.T) {
	var tests = []struct {
		name     string
		ints     []int
		expected int
	}{
		{"", []int{6, 7, 8, 13}, 4368},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := multiply(tt.ints...)
			if ans != tt.expected {
				t.Errorf("got %v, want %v", ans, tt.expected)
			}
		})
	}
}
