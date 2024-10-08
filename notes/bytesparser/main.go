package main

import (
	"fmt"
	"math"
)

func main() {
	number := newBytesToInt([]byte("2345"))
	fmt.Println(number)
}

func newBytesToInt(bs []byte) int {
	var order = len(bs) - 1
	var number = 0
	for i, _ := range bs {
		//toInt := dictionary.ByteToInt(b)
		number += int(math.Pow10(order)) * i
		order--
	}
	return number
}
