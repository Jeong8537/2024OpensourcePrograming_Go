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

	// counts := 0
	var isPrime bool = true // int->bool | counts -> isPrime
	if n <= 1 {
		// counts = -1
		isPrime = false // readability
	} else {
		i := 2
		for i < n {
			if n%i == 0 {
				// counts = counts + 1
				isPrime = false // +, remove
			}
			i++
		}
	}

	// if counts == 0 {
	if isPrime { // ==, remove
		fmt.Printf("%d is a Prime Number.", n)
	} else {
		fmt.Printf("%d not a Prime Number.", n)
	}
}
