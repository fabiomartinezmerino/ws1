package main

import (
	"fmt"
)

func main() {

	var stringPiece string = "Hello :)"
	var byteSlice []byte = []byte(stringPiece)
	var runeSlice []rune = []rune(stringPiece)

	fmt.Printf("String to Bytes: %v\n", byteSlice)
	fmt.Printf("String to runes: %v", runeSlice)

}
