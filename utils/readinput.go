package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Read(inputPath string, lineTrimmer func(line []byte) []byte, fieldFunc func([]byte) [][]byte, fieldTrimmer func([]byte) []byte, lineFieldProcessorFunc func(input [][]byte, last bool)) {
	file := OpenFileLogFatal(inputPath)
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	scanWithLastLineAwareness(scanner, lineTrimmer, fieldFunc, fieldTrimmer, lineFieldProcessorFunc)
}

func scan(scanner *bufio.Scanner, lineTrimmer func(line []byte) []byte, fieldFunc func([]byte) [][]byte, fieldTrimmer func([]byte) []byte, lineFieldProcessorFunc func(input [][]byte, last bool)) {
	for scanner.Scan() {
		line := scanner.Bytes()
		trimmedLine := line
		if lineTrimmer != nil {
			trimmedLine = lineTrimmer(line)
		}
		readLine(trimmedLine, fieldFunc, fieldTrimmer, lineFieldProcessorFunc, false)
	}
}

func scanWithLastLineAwareness(scanner *bufio.Scanner, lineTrimmer func(line []byte) []byte, fieldFunc func([]byte) [][]byte, fieldTrimmer func([]byte) []byte, lineFieldProcessorFunc func(input [][]byte, last bool)) {
	for ok := scanner.Scan(); ok; {
		line := scanner.Bytes()
		trimmedLine := line
		if lineTrimmer != nil {
			trimmedLine = lineTrimmer(line)
		}
		var l = make([]byte, len(trimmedLine))
		copy(l, trimmedLine)
		ok = scanner.Scan()
		last := !ok
		readLine(l, fieldFunc, fieldTrimmer, lineFieldProcessorFunc, last)
	}
}

func readLine(line []byte, fieldFunc func([]byte) [][]byte, fieldTrimmer func([]byte) []byte, lineFieldProcessorFunc func(fields [][]byte, last bool), lastLine bool) {
	if len(line) == 0 {
		return
	}
	fields := extractFields(line, fieldFunc)
	trimmedFields := trimUnnecessaryChars(fields, fieldTrimmer)
	lineFieldProcessorFunc(trimmedFields, lastLine)
}

func extractFields(line []byte, fieldFunc func([]byte) [][]byte) [][]byte {
	var fields [][]byte
	if fieldFunc != nil {
		fields = fieldFunc(line)
	} else {
		fields = [][]byte{line}
	}
	return fields
}

func trimUnnecessaryChars(lineFields [][]byte, fieldTrimmer func([]byte) []byte) [][]byte {
	res := make([][]byte, 0)
	for _, f := range lineFields {
		trimmedField := f
		if fieldTrimmer != nil {
			trimmedField = fieldTrimmer(f)
		}
		if len(trimmedField) == 0 {
			continue
		}
		res = append(res, trimmedField)
	}
	return res
}

func IsExcludedCharacter(allExcludedChars []byte) func(r rune) bool {
	return func(r rune) bool {
		return bytes.ContainsRune(allExcludedChars, r)
	}
}

func FindLinesContainingByteSequences(file *os.File, bs ...[]byte) {
	//fanIn := make(chan searchResult)
	scanner := bufio.NewScanner(file)
	resultChannels := make([]chan searchResult, 0)
	for scanner.Scan() {
		lineResult := make(chan searchResult)
		resultChannels = append(resultChannels, lineResult)
		go find(scanner.Bytes(), &bs, lineResult)
	}
	//var wg sync.WaitGroup
	for _, result := range resultChannels {
		for r := range result {
			fmt.Println(r)
		}
		//go func() {
		//	wg.Add(1)
		//	for r := range result {
		//		fanIn <- r
		//	}
		//	wg.Done()
		//}()
	}
	//res := make([]searchResult, 0)
	//for r := range fanIn {
	//	res = append(res, r)
	//}
	//wg.Wait()
	//close(fanIn)
	//return res
}

type searchResult struct {
	s      []byte
	result map[int][]int
}

func find(input []byte, byteSequences *[][]byte, result chan<- searchResult) {
	if res, ok := Indices(input, byteSequences); ok {
		result <- res
	}
	close(result)
}

func Indices(input []byte, byteSequences *[][]byte) (searchResult, bool) {
	mp, ok := Find(input).FirstInstances(byteSequences)
	return searchResult{input, mp}, ok
}

type Find []byte

func (f Find) FirstInstances(identifiers *[][]byte) (map[int][]int, bool) {
	return IndexOfFirstInstance(f, identifiers)
}

func IndexOfFirstInstance(b []byte, identifiers *[][]byte) (map[int][]int, bool) {
	const nothingFound = false
	const found = true
	presentIdentifiers := make(map[int][]int)
	for k, sep := range *identifiers {
		index := bytes.Index(b, sep)
		if index != -1 {
			indices := presentIdentifiers[k]
			if indices == nil {
				indices = make([]int, 0)
			}
			indices = append(indices, index)
			presentIdentifiers[k] = indices
		}
	}
	if len(presentIdentifiers) == 0 {
		return nil, nothingFound
	}
	return presentIdentifiers, found
}
