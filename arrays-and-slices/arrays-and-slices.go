package main

import (
	"fmt"
)

func main() {
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmad"
	students[2] = "Arnold"
	fmt.Printf("Student #1: %v\n", students[1])
}