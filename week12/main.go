package main

import (
	"fmt"
	"reflect"
)

func main() {
	// array slicing -> create slice
	// gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} // array := array literal
	// fmt.Println(gpas, reflect.TypeOf(gpas))
	// gpaSlice := gpas[1:4] // slice := slicing array
	// fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// array slicing make
	gpaSlice := []float64{4.1, 4.5, 3.9}
	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// make() function
	gpaSlice1 := make([]float64, 3)
	gpaSlice1[0] = 4.1
	gpaSlice1[1] = 4.5
	gpaSlice1[2] = 3.9

	fmt.Println(gpaSlice1, reflect.TypeOf(gpaSlice1))

}
