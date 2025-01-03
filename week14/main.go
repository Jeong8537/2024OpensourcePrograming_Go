package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func totalCost(visitors []visitor) int {
	total := 0
	for _, v := range visitors {
		total = total + v.cost
	}
	return total
}

func main() {
	var numVisitor int
	fmt.Printf("How many people? | [input] : ")
	fmt.Scanln(&numVisitor)

	vs := make([]visitor, numVisitor)

	for i := 0; i < numVisitor; i++ {
		age := 0
		fmt.Print("How old are you? | [input] : ")
		fmt.Scanln(&age)

		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}
		default:
			vs[i] = visitor{age: age, cost: 7000}
		}
	}
	fmt.Printf("total price is %d won.", totalCost(vs))
}
