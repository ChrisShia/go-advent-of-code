package main

import (
	"fmt"
	"time"
)

func main() {
	myTicker()
}

func myTicker() {
	//ticker := time.Tick(time.Second)
	for i := 1; i <= 100000000; i++ {
		if i == 50000000 {
			fmt.Printf("\rOn %d/100000000 with %d, ", i)
		}
		//fmt.Printf("\rOn %d/100000000 with %d, ", i)
	}
	fmt.Println("\nAll is said and done.")
}

func fromWebTicker() {
	ticker := time.Tick(time.Second * 5)
	for i := 1; i <= 10; i++ {
		<-ticker
		fmt.Printf("\rOn %d/10", i) // escape sequence is different in this environment
	}
	fmt.Println("\nAll is said and done.")
}
