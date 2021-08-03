package main

import (
	"fmt"
)

func main() {
	/*Array definition modes*/
	var myarr [3]float32
	fmt.Printf("Array: %v:\n", myarr)

	var myarr2 = [3]float32{123, 23, 43}
	fmt.Printf("Array: %v:\n", myarr2)

	var myarr3 = [...]float32{123, 23, 43}
	fmt.Printf("Array: %v:\n", myarr3)

	fmt.Printf("myarr2 y myarr3 son iguales?? %v\n", myarr2 == myarr3)

	/*Array length*/

	fmt.Printf("Longitud de myarr: %d\n", len(myarr))
	fmt.Printf("Longitud de myarr3: %d\n", len(myarr3))

	/*Slices*/

	/*var slice1 = []int{3, 2, 4, 5, 3, 4, 3}*/

	slice1 := []int{3, 2, 4, 5, 3, 4, 3}

	fmt.Printf("Contenido del slice: %v\n", slice1)

	var nil_slice []int
	fmt.Printf("Valor del slice nil: %v\n", nil_slice)

	fmt.Printf("Valor de la comparación con nil: %v\n", nil_slice == nil)

	/*

		Slices Builtin functions

	*/

	/*len*/

	fmt.Printf("Uso de len en slices len() longitud de slice1: %v\n", len(slice1))
	fmt.Printf("Qué pasa cuando len se aplica sobre un slice = nil?:%v \n", len(nil_slice))

	/*append*/
	nil_slice = append(nil_slice, 4, 5, 4, 5, 4, 5)
	nil_slice = append(nil_slice, slice1...)
	fmt.Printf("nil_slice after appending: %v\n", nil_slice)

	/*cap*/
	var mySlice []int

	fmt.Printf("slice:%v length: %d capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 3, 4, 5)

	fmt.Printf("slice:%v length: %d capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 34)

	fmt.Printf("slice:%v length: %d capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	mySlice = append(mySlice, 33)

	fmt.Printf("slice:%v length: %d capacity: %d\n", mySlice, len(mySlice), cap(mySlice))

	/*make*/

	slice2 := make([]float64, 3, 10) /*initializes to zero three first positions in the slice*/

	fmt.Printf("slice2 length: %d\nslice2 capacity: %d\nslice2 value: %v\n", len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 3, 4, 5, 6)

	fmt.Printf("---After appending---\nslice2 length: %d\nslice2 capacity: %d\nslice2 value: %v\n", len(slice2), cap(slice2), slice2)

	/*make and slices*/

	x := make([]int, 4, 10)
	fmt.Printf("Después de make: %v - Capacidad: %d\n", x, cap(x))
	x = append(x, 1, 2)
	fmt.Printf("Después de hacer appending: %v\n - Capacidad: %d\n", x, cap(x))

	/*copy builtin function: is the lenth that matters*/

	y := make([]int, 5, 20)
	fmt.Printf("Y antes de copy: %v - Capacidad: %d\n", y, cap(y))

	numberOfItemsCopied := copy(y, x)
	fmt.Printf("y después de copy: %v - Número de posiciones copiadas: %d\n", y, numberOfItemsCopied)

}
