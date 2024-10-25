package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	// fmt.Println(answer)

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Print("숫자 입력 : ")
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		//guess, err := strconv.ParseInt(input, 10, 32)
		guess, err := strconv.Atoi(input) // string to integer
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if guess == answer {
			fmt.Println("정답입니다!")
			break
		} else if answer > guess {
			fmt.Println("입력하신 값은 정답보다 작은 값입니다. LOW")
		} else {
			fmt.Println("입력하신 값은 정답보다 큰 값입니다. HIGH")
		}
	}
}
