package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		i := 3

		for i*i <= n {
			if n%i == 0 {
				return false
				// isPrime = false
				// break
			}
			fmt.Printf("%d ", i)
			i = i + 2
		}
	}
	return true
}

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

	if isPrime(n) {
		fmt.Printf("%d is a Prime Number.", n)
	} else {
		fmt.Printf("%d not a Prime Number.", n)
	}
}
