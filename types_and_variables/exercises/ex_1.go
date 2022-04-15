package main

import "fmt"

// Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

func main() {
	x := 42

	// the `x` add a 0b and 0x preffix for binaries and hexadecimals values
	fmt.Printf("%d\t%#b\t%#x\n", x, x, x)
}
