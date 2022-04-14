package main

import "fmt"

func main() {
	x := 4
	y := x << 1 // add one bit 0 in the right (bits "walks" to left)
	z := x >> 1 // remove an bit from the left (bits "moves" to right)

	fmt.Printf("x: %b\n", x) // out: 100
	fmt.Printf("y: %b\n", y) // out: 1000
	fmt.Printf("z: %b\n", z) // out: 10
}
