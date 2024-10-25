package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	// fmt.Sprintf("%d\n", rand.Intn(6)+1)
	r := fmt.Sprintf("%10d\n", rand.Intn(6)+1)
	// a, _ := strconv.Atoi(r)
	fmt.Println(reflect.TypeOf(r))
	fmt.Printf("%T\n", 2.1)

	i := 1
	for i <= 100 { // while
		fmt.Printf("%3d점\n", i)
		i = i + 1
	}
}
