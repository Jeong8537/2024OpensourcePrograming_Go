package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Printf("%f\n", math.Sqrt(19.0)) // root
	fmt.Print("Input Number : ")
	in := bufio.NewReader(os.Stdin)
	number, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	number = strings.TrimSpace(number)
	n, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal()
	}

	var isPrime bool = true
	if n <= 1 {
		isPrime = false
	} else if n == 2 {
		isPrime = true
	} else if n%2 == 0 { // Among even number, only 2 is Prime Number
		isPrime = false
	} else {
		i := 3
		// for i < n {
		// for i < int(math.Sqrt(float64(n))) { // Error
		// for i <= int(math.Sqrt(float64(n))) { // Same as current code

		//for float64(i) < math.Sqrt(float64(n)) { // Same as current code
		for i*i <= n {
			if n%i == 0 {
				isPrime = false
				break // first divisor, loop break
			}
			fmt.Printf("%d ", i)
			// i++
			i = i + 2
		}
	}

	if isPrime {
		fmt.Printf("%d is a Prime Number.", n)
	} else {
		fmt.Printf("%d not a Prime Number.", n)
	}
}
