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

	fmt.Println(dates[0], dates[1], dates[2])
	fmt.Println(dates)         // array
	fmt.Printf("%#v\n", dates) //array literal
}
