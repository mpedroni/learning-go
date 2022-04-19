package main

import "fmt"

/*
	OBS: arrays are used in specific cases that involve optimized memory management.
	use slices instead
*/

func main() {
	// arrays have a fixed length
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", array) // out: [5]int

	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", slice) // out: []int

	// slices have a dynamic length
	slice = append(slice, 6)
	fmt.Println(slice) // out: [1, 2, 3, 4, 5, 6]
}
