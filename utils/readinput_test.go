package utils

import (
	"testing"
)

func TestContainsAnyIdentifier(t *testing.T) {
	identifiers := make([][]byte, 0)
	identifiers = append(identifiers, []byte("Time"), []byte("User"), []byte("Distance"))
	var tests = []struct {
		name        string
		input       string
		identifiers [][]byte
		expectedLen int
		expectedMap map[int][]int
	}{
		{"Test_ContainsAnyIdentifier_only_Time_id_present", "Time: 123465", identifiers, 1, map[int][]int{0: {0}}},
		{"test all identifiers present once", "Time: 123465, Distance: 8093745, User: 092374", identifiers, 3, map[int][]int{0: {0}, 1: {33}, 2: {14}}},
	}
	for _, tt := range tests {
		mp, _ := IndexOfFirstInstance([]byte(tt.input), &tt.identifiers)
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedLen != len(mp) {
				t.Errorf("got %v, want %v", len(mp), tt.expectedLen)
			}
			for i, m := range mp {
				for j, e := range m {
					if e != tt.expectedMap[i][j] {
						t.Errorf("got %v, want %v for identifier %s", e, tt.expectedMap[i][j], tt.identifiers[i])
					}
				}
			}
		})
	}
}

func TestFindLinesOfFileContainingSequences(t *testing.T) {
	identifiers := make([][]byte, 0)
	identifiers = append(identifiers, []byte("Time"), []byte("User"), []byte("Distance"))
	tests := []struct {
		name          string
		inputFilePath string
		identifiers   [][]byte
	}{
		{"test file", "/Users/christos/Practise-code/Go/go-advent/input/test/scaninput_test.txt", identifiers},
		//{"test file", "input/test/scaninput_test.txt", identifiers},
	}
	for _, tt := range tests {
		file := OpenFileLogFatal(tt.inputFilePath)
		FindLinesContainingByteSequences(file, tt.identifiers...)
		//result := FindLinesContainingByteSequences(file, tt.identifiers...)
		CloseFile(file)
		t.Run(tt.name, func(t *testing.T) {
			//fmt.Println(result)
			//for i, m := range result {
			//for j, e := range m {
			//	if e != tt.expectedMap[i][j] {
			//		t.Errorf("got %v, want %v for identifier %s", e, tt.expectedMap[i][j], tt.identifiers[i])
			//	}
			//}
			//}
		})
	}
}
