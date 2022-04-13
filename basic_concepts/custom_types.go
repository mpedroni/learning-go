package main

import (
	"fmt"
)

type superhero int

var batman superhero

func main() {
	x := 10

	fmt.Printf("%T\n", batman) // out: superhero
	fmt.Printf("%T\n", x)      // out: int

	// don't work because superhero != int; `batman` needs to be CONVERTED (not "casted" or any else) to int
	// x = batman
	x = int(batman)
}
