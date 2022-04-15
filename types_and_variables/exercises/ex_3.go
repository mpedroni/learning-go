package main

import "fmt"

/*
	- Crie constantes tipadas e não-tipadas.
	- Demonstre seus valores.
*/

const x int = 10
const y = 10

func main() {
	fmt.Printf("x: %v\t%T\n", x, x)
	fmt.Printf("y: %v\t%T", y, y)
}
