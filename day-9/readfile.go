package main

import (
	"bufio"
	"bytes"
	"go-advent-of-code/dictionary"
	"go-advent-of-code/utils"
)

func read(inputPath string, lineFieldProcessorFunc func(input [][]byte)) {
	file := utils.OpenFileLogFatal(inputPath)
	defer utils.CloseFile(file)
	scanner := bufio.NewScanner(file)
	var line []byte
	for scanner.Scan() {
		line = scanner.Bytes()
		readLine(line, lineFieldProcessorFunc)
	}
}

func readLine(line []byte, lineFieldProcessorFunc func(input [][]byte)) {
	if len(line) == 0 {
		return
	}
	fields := bytes.Fields(line)
	trimmedFields := trimUnnecessaryChars(fields)
	lineFieldProcessorFunc(trimmedFields)
}

func lineFieldProcessor() func(input [][]byte) {
	return dayNineLineProcessor
}

func dayNineLineProcessor(input [][]byte) {
	values := make([]int, 0)
	for _, bs := range input {
		toInt := dictionary.BytesToInt(bs)
		values = append(values, toInt)
	}
	sequences_ = append(sequences_, createSequence(values...))
}

func trimUnnecessaryChars(lineFields [][]byte) [][]byte {
	res := make([][]byte, 0)
	for _, s := range lineFields {
		trimmedField := bytes.TrimFunc(s, isExcludedCharacter(nil))
		if len(trimmedField) == 0 {
			continue
		}
		res = append(res, trimmedField)
	}
	return res
}

func isExcludedCharacter(allExcludedChars []byte) func(r rune) bool {
	return func(r rune) bool {
		return bytes.ContainsRune(allExcludedChars, r)
	}
}
