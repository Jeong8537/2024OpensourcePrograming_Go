package main

import (
	"fmt" // == c lang, #include <stdio.h>
	"reflect"
)

func main() {
	i := 13
	//f := 12.9
	var f float64 = 12.9
	c1 := 'Z'
	c2 := '김'

	fmt.Printf("value i : %d, value f : %f\n", i, f)
	// fmt.Printf("%d * %f = %d\n", i, f, i*f) // mismatched
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) // conversion
	// fmt.Printf("%d * %f = %d\n", i, f, i*int(f)) // conversion
	fmt.Println(c1, c2)    // 10 진수
	fmt.Printf("%x\n", c2) // unicode '김' 16진수 출력
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i), reflect.TypeOf(c1), reflect.TypeOf(c2))
}
