package main

import (
	"fmt"
)

var integer int
var float float32
var boolean bool
var str string

func main() {
	fmt.Println("==== The zero values (initial value of a variable) for each type ====")
	fmt.Printf("\n%v\t %T", integer, integer) // 0
	fmt.Printf("\n%v\t %T", float, float)     // 0.0
	fmt.Printf("\n%v\t %T", boolean, boolean) // false
	fmt.Printf("\n\"\"\t %T", str)            // ""
	fmt.Println("\nnil\t pointers, functions, interfaces, slices, channels, maps...")
}
