package main

import (
	"fmt" // == c lang, #include <stdio.h>
	"reflect"
)

func main() {
	i := 13
	//f := 12.9
	var f float64 = 12.9

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	// fmt.Printf("%d * %f = %d\n", i, f, i*f) // mismatched
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) // conversion
	// fmt.Printf("%d * %f = %d\n", i, f, i*int(f)) // conversion
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
}
