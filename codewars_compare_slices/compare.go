package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var a1 = []int{19, 19, 19, 144, 144, 161, 121, 11, 19}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19, 19 * 19}
	fmt.Printf("Resultado: %v\n", Comp(a1, a2))
}
func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	if len(array1) == 0 || len(array2) == 0 {
		return false
	}

	if len(array1) != len(array2) {
		return false
	}

	sort.Ints(array1)
	sort.Ints(array2)

	for i, v := range array2 {

		if int(math.Pow(float64(array1[i]), 2)) == v {
			continue
		} else {
			return false
		}

	}

	return true
}
