package main

import (
	"fmt"
)

var packlevel string

func main() {

	var boolNoninitialized bool
	var boolInitialized = true
	var uint8Noninitialized uint8
	var binary2Show byte

	binary2Show = 0b11111111
	binary2Show >>= 2
	packlevel = "Hello!!!"

	fmt.Printf("Binary String: %b\n", binary2Show)
	fmt.Printf("Package Level Var: %v", packlevel)
	binary2Show <<= 2

	fmt.Printf("Binary String: %b\n", binary2Show)

	fmt.Printf("%d\n", 0b111011)
	fmt.Printf("%b\n", 59)
	fmt.Println(boolNoninitialized)
	fmt.Println(boolInitialized)
	fmt.Println(uint8Noninitialized)
}
