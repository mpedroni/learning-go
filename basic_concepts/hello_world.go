package main

import (
	"fmt"
)

func main() {
	bytes, err := fmt.Println("Hello world!")

	fmt.Println(bytes, err)
}
