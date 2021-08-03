package main

import (
	"fmt"
)

type animal struct {
	name         string
	numberOfLegs int
	tail         bool
}

func main() {

	dog := animal{
		name:         "dog",
		numberOfLegs: 4,
		tail:         true,
	}

	fmt.Printf("Strruct: %v\n", dog)

	/*dog undergoes surgery*/

	dog.tail = false

	fmt.Printf("Strruct: %v\n", dog)

}
