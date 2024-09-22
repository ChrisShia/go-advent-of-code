package main

import (
	"fmt"
	"time"
)

var counter int
var addCounterChan chan int
var readCounterChan chan int

func main() {
	addCounterChan = make(chan int, 100)
	readCounterChan = make(chan int, 100)

	counter = 0

	go func() {
		for {
			select {
			case val := <-addCounterChan:
				counter += val
				if counter > 5 {
					counter = 0
				}
				readCounterChan <- counter
				fmt.Printf("%d \n", counter)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		go AddCounter(addCounterChan)
	}

	time.Sleep(time.Minute)

	for i := 0; i < 10; i++ {
		fmt.Printf("Total Count #%d is ... %d \n", (i + 1), GetCount(readCounterChan))
	}

}

func AddCounter(ch chan int) {
	ch <- 1
}

func GetCount(ch chan int) int {
	r := <-ch
	return r
}
