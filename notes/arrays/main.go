package main

import (
	"fmt"
)

func main() {
	arr := [10]int{}
	coef := arr[:2]
	coef[0], coef[1] = 1, 1
	evenBuf := make([]int, len(coef)+1, 10)
	oddBuf := make([]int, len(coef)+2, 10)
	copy(evenBuf[1:], coef)
	for i := 0; i < len(evenBuf); i++ {
		evenBuf[i] *= 2
	}
	copy(oddBuf[1:], evenBuf)
	for i := 0; i < len(oddBuf); i++ {
		oddBuf[i] *= 2
	}
	for i := 0; i < len(evenBuf); i++ {
		oddBuf[i] += evenBuf[i]
	}
	fmt.Println(evenBuf)
	fmt.Println(oddBuf)
	copy(coef[:len(oddBuf)], oddBuf[:])
	coef = coef[:len(oddBuf)]
	clear(evenBuf)
	clear(oddBuf)
	fmt.Println("------------------")
	evenBuf = evenBuf[:len(coef)+1]
	oddBuf = oddBuf[:len(coef)+2]
	copy(evenBuf[1:], coef)
	for i := 0; i < len(evenBuf); i++ {
		evenBuf[i] *= 3
	}
	copy(oddBuf[1:], evenBuf)
	for i := 0; i < len(oddBuf); i++ {
		oddBuf[i] *= 3
	}
	for i := 0; i < len(evenBuf); i++ {
		oddBuf[i] += evenBuf[i]
	}
	fmt.Println(evenBuf)
	fmt.Println(oddBuf)
	copy(coef[:len(oddBuf)], oddBuf)
	coef = coef[:len(oddBuf)]

}

func testPointerSliceWrapper() {
	wrapper1 := pointerSliceWrapper{&[]string{"one", "two", "three"}, "ttt"}
	(*wrapper1.s)[0] = "two"
	wrapper1.str = "yyyy"
	fmt.Println(*wrapper1.s)
	fmt.Println(wrapper1.str)
}

type pointerSliceWrapper struct {
	s   *[]string
	str string
}

func (p *pointerSliceWrapper) String() string {
	p.str = "yyy"
	return p.str
}

type sliceWrapper struct {
	s []string
}
