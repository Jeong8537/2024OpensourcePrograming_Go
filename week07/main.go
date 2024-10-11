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
	//var realScore int32
	fmt.Print("정수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print(strings.TrimSpace(score))
	score = strings.TrimSpace(score)                // \n, tab, space remove
	realScore, _ := strconv.ParseInt(score, 16, 32) // int32 type <= str type change

	if realScore >= 60 {
		fmt.Println("A")
		fmt.Printf("%d\n", realScore)
	} else {
		fmt.Println("BCDF")
		fmt.Printf("%d\n", realScore)
	}
}
