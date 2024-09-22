package main

import "fmt"

func main() {
	ch := make(chan int64, 10)
	ch <- 1
	ch <- 2
	for i := int64(3); i <= int64(10); i++ {
		ch <- i
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
