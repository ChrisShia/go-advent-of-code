package utils

import (
	"bufio"
	"bytes"
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
