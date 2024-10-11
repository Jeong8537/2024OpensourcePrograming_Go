package main

import (
	"fmt"
	"strings"
)

func main() {
	var hello string = "He##o, wor#d!"
	helloFixed := strings.NewReplacer("#", "l")
	fmt.Println(hello)
	fmt.Println(helloFixed.Replace(hello))
}
