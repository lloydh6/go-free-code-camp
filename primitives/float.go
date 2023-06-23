package primitives

import (
	"fmt"
)

func primitives() {
	n := 3.14
	n = 13.7e72
	n = 2.1E14
	fmt.Printf("%v, %T\n", n, n)
}