package main

import (
	"bufio"
	"fmt"
)

import (
	"bytes"
	"go-advent-of-code/dictionary"
	"log"
	"math"
	"os"
	"slices"
)

//ans a: 265018614
//ans b: 63179500

const inputPath_ = "/Users/christos/Practise-code/Go/go-advent/input/day-5.txt"

var seedIdsAndRanges_ []int
var srcToDesMaps_ = make(map[int]*SortedMap)
var destinationToSourceFuncMap = make(map[string]interface{})

func main() {
	file := openFile()
	defer closeFile(file)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Bytes()
	scanFirstLineAndAppendToIdsSlice(line)
	scanFileAndPopulateMap(scanner)
	fmt.Println(mapSeedsToLocations())
}

func mapSeedsToLocations() int {
	closestLocation := -1
	for s := 0; s < len(seedIdsAndRanges_); s += 2 {
		upperSeedBoundary := seedIdsAndRanges_[s] + seedIdsAndRanges_[s+1]
		seed := seedIdsAndRanges_[s]
		for seed < upperSeedBoundary {
			mappedValue := seed
			for i := 1; i <= len(srcToDesMaps_); i++ {
				partialMap := srcToDesMaps_[i]
				mappedValue = getMappingFromPartialMap(partialMap, mappedValue)
			}
			if mappedValue < closestLocation || closestLocation < 0 {
				closestLocation = mappedValue
			}
			seed++
		}
	}
	return closestLocation
}

func getMappingFromPartialMap(partialMap *SortedMap, mappedValue int) int {
	var rightBoundary int
	for k := 0; k < len(partialMap.sortedKeys); k++ {
		rightBoundary = partialMap.sortedKeys[k]
		if mappedValue <= rightBoundary {
			if isUpperBoundary(k) {
				mappedValue = partialMap.m[rightBoundary] - (rightBoundary - mappedValue)
			} else if isLowerBoundary(k) {
				if mappedValue == rightBoundary {
					mappedValue = partialMap.m[rightBoundary] - (rightBoundary - mappedValue)
				}
			}
			return mappedValue
		}
	}
	return mappedValue
}

func isLowerBoundary(k int) bool {
	return math.Remainder(float64(k), 2) == 0
}

func isUpperBoundary(k int) bool {
	return math.Remainder(float64(k), 2) != 0
}

func closeFile(file *os.File) {
	func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}

func scanFirstLineAndAppendToIdsSlice(line []byte) {
	_, seedIdsBytes, _ := bytes.Cut(line, []byte("seeds: "))
	seedIdsSlice := bytes.Fields(seedIdsBytes)
	for _, id := range seedIdsSlice {
		seedIdsAndRanges_ = append(seedIdsAndRanges_, dictionary.BytesToInt(id))
	}
}

func scanFileAndPopulateMap(scanner *bufio.Scanner) {
	var line []byte
	var srcToDesSortedMap *SortedMap
	var mapCounter int
	for scanner.Scan() {
		line = scanner.Bytes()
		if isTheStartOfMapListing(line) {
			addToMaps(srcToDesSortedMap, mapCounter)
			mapCounter++
			srcToDesSortedMap = newSortedMap()
			continue
		} else {
			if len(line) == 0 {
				continue
			}
			addLineElementsToSortedMap(line, srcToDesSortedMap)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	addLastMapToCollectionOfMaps(srcToDesSortedMap, mapCounter)
}

func addLastMapToCollectionOfMaps(srcToDesSortedMap *SortedMap, mapCounter int) {
	addToMaps(srcToDesSortedMap, mapCounter)
}

func addLineElementsToSortedMap(line []byte, srcToDesSortedMap *SortedMap) {
	desSrcOffset := make([]int, 0)
	for _, value := range bytes.Fields(line) {
		desSrcOffset = append(desSrcOffset, dictionary.BytesToInt(value))
	}
	var des = desSrcOffset[0]
	var src = desSrcOffset[1]
	var offset = desSrcOffset[2]
	srcToDesSortedMap.add(src, des)
	srcToDesSortedMap.add(src+offset-1, des+offset-1)
}

func addToMaps(srcToDesSortedMap *SortedMap, mapCounter int) {
	if srcToDesSortedMap != nil && srcToDesSortedMap.len() > 0 {
		srcToDesMaps_[mapCounter] = srcToDesSortedMap
	}
}

func isTheStartOfMapListing(line []byte) bool {
	return bytes.Contains(line, []byte("map"))
}

func openFile() *os.File {
	file, err := os.Open(inputPath_)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

type SortedMap struct {
	m          map[int]int
	sortedKeys []int
}

func newSortedMap() *SortedMap {
	return &SortedMap{
		make(map[int]int),
		make([]int, 0),
	}
}

func (sm *SortedMap) add(src, des int) {
	_, present := sm.m[src]
	if present {
		return
	} else {
		sm.m[src] = des
		sm.sortedKeys = append(sm.sortedKeys, src)
		slices.Sort(sm.sortedKeys)
	}
}

func (sm *SortedMap) len() int {
	return len(sm.m)
}
