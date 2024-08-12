package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("c <- x")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		fmt.Println("In go func 1")
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	go func() {
		fmt.Println("In go func 2")
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("before fibon")
	fibonacci(c, quit)
}
