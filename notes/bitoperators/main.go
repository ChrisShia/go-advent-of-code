package main

import (
	"fmt"
)

func main() {

	var _0001 = 1
	var _0011 = 3
	var _0101 = 5
	var _0110 = 6
	var _1100 = 12
	var _1111 = 15

	// If you want the output to have the same length, 8 bits for example
	// use the %08b notation which means:
	// print binary, use 4 digits for the output, pad with leading zeros
	fmt.Printf("%04b\n", _0001)
	fmt.Printf("%04b\n", _0011)
	fmt.Printf("%04b\n", _0101)
	fmt.Printf("%04b\n", _0110)
	fmt.Printf("%04b\n", _1100)

	fmt.Println()
	fmt.Println("Operations:")
	fmt.Printf("1. 0001 & 0011 : %04b\n", _0001&_0011)   //0001
	fmt.Printf("2. 0001 ^ 0011 : %04b\n", _0001^_0011)   //0010
	fmt.Printf("3. 0001 | 0011 : %04b\n", _0001|_0011)   //0011
	fmt.Printf("4. 0011 ^ 0101 : %04b\n", _0011^_0101)   //0110
	fmt.Printf("5. 0001 &^ 0011 : %04b\n", _0001&^_0011) //0000
	fmt.Printf("6. 0110 &^ 0011 : %04b\n", _0110&^_0011) //0100
	fmt.Printf("7. 1111 &^ 0001 : %04b\n", _1111&^_0001) //1110
	fmt.Printf("7. 1111 &^ 0001 : %04b\n", _1111&^_0001) //1110
}
