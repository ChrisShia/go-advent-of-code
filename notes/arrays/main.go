package main

import "fmt"

func main() {

	//aints := []int{1, 2, 3, 4, 5}
	//astrings := []string{"sav", "chris"}
	//
	//bints := aints
	//bstrings := astrings
	//
	//bints[0] = 0
	//bstrings[0] = "chris"
	//
	//fmt.Println(aints)
	//fmt.Println(astrings)

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
