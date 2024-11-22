package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) // err
// func test(strs ...string, int i) // err
func test(i int, strs ...string) { // variable parameter -> only final parameter
	fmt.Println(i, strs, "|", reflect.TypeOf((strs)))
}
func main() {
	//fmt.Println(os.Args, len(os.Args), reflect.TypeOf(os.Args))
	slices := os.Args[1:]
	fmt.Println(slices, slices[2], slices[1], slices[0])

	test(1, "123")
	test(-99, "123", "ABC")
	test(55)
	test(0, "123", "abc", "123a")
}
