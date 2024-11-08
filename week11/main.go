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

			}
			// fmt.Printf("%d ", i)
			i = i + 2
		}
	}
	return true
}

// input
// 10
// 19

// output
// 11 13 17 19
func main() {
	fmt.Print("Start Number : ")
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	a = strings.TrimSpace(a)
	n1, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal()
	}
	fmt.Print("End Number : ")
	// in := bufio.NewReader(os.Stdin)
	b, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	b = strings.TrimSpace(b)
	n2, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal()
	}

	for i := n1; i < n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
