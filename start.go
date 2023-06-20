package main

import "fmt"

func main() {
	s := "this is a string"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
}