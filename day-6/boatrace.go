package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go-advent-of-code/dictionary"
	"log"
	"math"
	"os"
)

// Answer a: 131376
// Answer b: 34123437

const inputPath_ = "input/day-6.txt"

func main() {
	timeByteSlice, distByteSlice := extractTimesAndDistances()
	product := 1
	for i := 0; i < len(timeByteSlice); i++ {
		time := dictionary.BytesToInt(timeByteSlice[i])
		distance := dictionary.BytesToInt(distByteSlice[i])
		product *= winningMoves(time, distance)
	}
	fmt.Println(product)
}

func winningMoves(t int, d int) int {
	quadSol1 := (float64(t) + determinant(t, d)) / 2
	quadSol2 := (float64(t) - determinant(t, d)) / 2
	if quadSol1 > float64(int(quadSol1)) {
		quadSol1 = math.Round(quadSol1 - 0.5)
	} else {
		quadSol1--
	}
	if quadSol2 > float64(int(quadSol2)) {
		quadSol2 = math.Round(quadSol2 - 0.5)
	}
	rootDiff := quadSol1 - quadSol2
	return int(rootDiff)
}

func determinant(t int, d int) float64 {
	return math.Sqrt(math.Pow(float64(t), 2) - 4*float64(d))
}

func extractTimesAndDistances() ([][]byte, [][]byte) {
	file := openFile()
	defer closeFile(file)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	_, times, _ := bytes.Cut(scanner.Bytes(), []byte("Time: "))
	scanner.Scan()
	_, distances, _ := bytes.Cut(scanner.Bytes(), []byte("Distance: "))
	timeByteSlice := bytes.Fields(times)
	distByteSlice := bytes.Fields(distances)
	return timeByteSlice, distByteSlice
}

func openFile() *os.File {
	file, err := os.Open(inputPath_)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func closeFile(file *os.File) {
	func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
}
