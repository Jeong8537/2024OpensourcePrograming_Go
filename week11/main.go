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
func getInteger() int {
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	a = strings.TrimSpace(a)
	number, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal()
	}
	return number
}

func main() {
	fmt.Print("Start Number : ")
	n1 := getInteger()
	fmt.Print("End Number : ")
	n2 := getInteger()

	for i := n1; i < n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
