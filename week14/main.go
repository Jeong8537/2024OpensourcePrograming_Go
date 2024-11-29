package main

import "fmt"

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("\nName?: ")
		fmt.Scanln(&name)
		if name == "q" {
			break
		} else {
			fmt.Print("Input your age: ")
			fmt.Scanln(&age)

			ages[name] = age
		}
		for k, v := range ages {
			fmt.Printf("%s is %d years old", k, v)
		}
	}
}
