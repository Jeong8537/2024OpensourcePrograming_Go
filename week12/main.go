package main

import (
	"fmt"
	"time"
)

func main() {
	// var dates [3]time.Time
	// dates[0] = time.Unix(0, 0)
	// dates[2] = time.Unix(1708012346, 0) // Unix Time, Zero Value, 2024-02-16
	// fmt.Println(dates[0], dates[1], dates[2])

	// var dates [3]time.Time = [3]time.Time(time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0))
	// fmt.Println(dates[0], dates[1], dates[2])

	//dates := [3]time.Time{time.Unix(0, 0), time.Unix(1, 0), time.Unix(1708012346, 0)}
	//fmt.Println(dates[0], dates[1], dates[2])

	// dates := [3]time.Time{
	// 	time.Unix(0, 0),
	// 	time.Unix(1, 0),
	// 	time.Unix(1708012346, 0),
	// }
	// fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1708012346, 0)}

	// fmt.Println(dates[0], dates[1], dates[2])
	// fmt.Println(dates)         // array
	// fmt.Printf("%#v\n", dates) //array literal
	// for i := 0; i <= 2; i++{

	// for i := 0; i < len(dates); i++ {
	// 	fmt.Printf("[%d] ", i)
	// 	fmt.Println(dates[i])
	// }
	for i, date := range dates { // like python for in, SAFE
		fmt.Println(i, date)
	}
	fmt.Printf("\n")
	for _, date := range dates { // like python for in, SAFE
		fmt.Println(date)
	}
}
