package main

import (
	"fmt"
	"slices"
)

var m = map[int]int{
	987134265:          1000,
	987162345:          875,
	1324509876:         90,
	1320498756145:      123,
	1192348756:         927,
	21345013456:        284,
	2349873459867:      583,
	98723945674265:     1000,
	98723945672345:     875,
	132394567509876:    90,
	132049926456145:    123,
	119232394567756:    927,
	2123945675013456:   284,
	223945679873459867: 583,
}

func main() {
	sortedMap := newSortedMap(&m)
	fmt.Print(sortedMap.sortedKeys)
}

type SortedMap struct {
	m          *map[int]int
	sortedKeys []int
}

func newSortedMap(unsortedMap *map[int]int) *SortedMap {
	lenMap := len(*unsortedMap)
	var sortedKeys = make([]int, lenMap)
	for i := 0; i < lenMap; {
		sortedKeys[i] = (*unsortedMap)[i]
	}
	slices.Sort(sortedKeys)
	return &SortedMap{
		m:          unsortedMap,
		sortedKeys: sortedKeys,
	}
}
