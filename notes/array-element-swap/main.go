package main

import "fmt"

func main() {

	ints := []int{1, 2, 3, 4}
	ints[0], ints[1], ints[2], ints[3] = ints[1], ints[0], ints[0], ints[0]
	fmt.Println(ints)

}
