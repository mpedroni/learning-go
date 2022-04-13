package main

import (
	"fmt"
)

func main() {
	// gopher (marmota) operator -> create variables (type inference)
	// gopher only works inside code blocks
	x := 10
	y := "shazam"

	// prints the variable value (%v) and type (%T)
	fmt.Printf("x: %v, %T", x, x)
	fmt.Printf("\nx: %v, %T", y, y)

	a, b := 42, 43
	fmt.Printf("\n\na: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
}
