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

	var isPrime bool = true
	if n <= 1 {

		isPrime = false
	} else {
		i := 2
		for i < n {
			if n%i == 0 {
				isPrime = false
				break // first divisor, loop break
			}
			fmt.Printf("%d ", i)
			i++
		}
	}

	if isPrime {
		fmt.Printf("%d is a Prime Number.", n)
	} else {
		fmt.Printf("%d not a Prime Number.", n)
	}
}
