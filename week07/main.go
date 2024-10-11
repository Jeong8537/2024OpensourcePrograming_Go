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
	fmt.Print("정수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print(strings.TrimSpace(score))
	score = strings.TrimSpace(score)              // \n, tab, space remove
	realScore, _ := strconv.ParseFloat(score, 64) // float64 type => str type change

	if realScore >= 90 {
		fmt.Println("A")
	} else {
		fmt.Println("BCDF")
	}
}
