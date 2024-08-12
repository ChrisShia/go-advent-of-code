package main

import "fmt"

func main() {

	m := make(map[string][]int)
	m["AA"] = []int{1, 2}
	m["BB"] = []int{6, 5}
	if m["AB"] == nil {
		fmt.Printf("its nil")
	} else {
		fmt.Println(m["AB"])
	}
}
