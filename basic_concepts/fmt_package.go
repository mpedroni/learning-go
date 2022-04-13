package main

import (
	"fmt"
)

func main() {
	x := "batman"
	y := "> superman"

	// Sprint(*) is like Print(*), but returns the evaluated string instead of display it in stdout
	// as well as Fprint(*) but these are focused in files stuff
	z := fmt.Sprint(x, " ", y)

	fmt.Println(z)
}
