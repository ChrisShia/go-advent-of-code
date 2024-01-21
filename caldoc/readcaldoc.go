package main

import (
	"fmt"
	"go-advent-of-code/utils"
)

func main() {
	input := utils.ReadFile("/Users/christos/Practise-code/Go/go-advent/caldoc/input.txt")
	sum := utils.SumCalibrationSequence(input)
	fmt.Println("Sum is : ", sum)
}
