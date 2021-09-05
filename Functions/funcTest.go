package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func mul(i int, j int) int {
	return i * j
}

func div(i int, j int) int {
	return i / j
}

func main() {
	opMap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressionsTests := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"two", "+", "3"},
		[]string{"5"},
	}

	for _, expression := range expressionsTests {
		if len(expression) != 3 {
			fmt.Println("Invalid expression:", expression)
			continue
		}
		operand1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		operation, ok := opMap[expression[1]]
		if !ok {
			fmt.Println("Operator not Supported:", expression[1])
			continue
		}
		operand2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println("Result:", operation(operand1, operand2))

	}

}
