package dictionary

import (
	"testing"
)

func TestDigit_Byte(t *testing.T) {
	var tests = []struct {
		name  string
		input Digit
		want  byte
	}{
		{"0 is 48", ZERO, '0'},
		{"0 is 49", ONE, '1'},
		{"0 is 50", TWO, '2'},
		{"0 is 51", THREE, '3'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.Byte()
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestByteToNumerical(t *testing.T) {
	var tests = []struct {
		name  string
		input byte
		want  Digit
	}{
		{"a is nil", 'a', NonDigit},
		{"0 is ZERO", 0, ZERO},
		{"1 is ONE", 1, ONE},
		{"2 is TWO", 2, TWO},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := ByteToDigit(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	var tests = []struct {
		name  string
		input Digit
		want  string
	}{
		{"label for ZERO is zero", ZERO, "zero"},
		{"label for ONE is one", ONE, "one"},
		{"label for TWO is two", TWO, "two"},
		{"label for THREE is three", THREE, "three"},
		{"label for FOUR is four", FOUR, "four"},
		{"label for FIVE is five", FIVE, "five"},
		{"label for SIX is six", SIX, "six"},
		{"label for SEVEN is seven", SEVEN, "seven"},
		{"label for EIGHT is eight", EIGHT, "eight"},
		{"label for NINE is nine", NINE, "nine"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tt.input.String()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestWordToDigitMapper(t *testing.T) {
	var tests = []struct {
		name  string
		input []byte
		want  Digit
	}{
		{"one is 1", []byte("one"), ONE},
		{"thr is NonDigit", []byte("thr"), NonDigit},
		{"niner is NonDigit", []byte("niner"), NonDigit},
		{"one1 is NonDigit", []byte("one1"), NonDigit},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ans Digit
			mapper := WordToDigitMapper()
			for _, b := range tt.input {
				ans = mapper(b)
			}
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
