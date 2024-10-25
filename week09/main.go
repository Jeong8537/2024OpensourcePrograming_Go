package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // fixed seed, Unix time
	answer := rand.Intn(6) + 1   // dice 1 ~ 6
	fmt.Println(answer)
}
