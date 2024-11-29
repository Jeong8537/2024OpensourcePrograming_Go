package main

import "fmt"

func main() {
	var student struct {
		id   int
		name string
		gpa  float32
	}
	student.id = 202444057
	student.name = "Jeong Jang Hyeon"
	student.gpa = 3.0
	fmt.Println(student.gpa)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student2.id = 202444001
	student2.name = "Son Heungmin"
	student2.gpa = 4.5
	fmt.Println(student2.gpa)
}
