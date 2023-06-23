package main

import (
	"fmt"
)

const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)
}