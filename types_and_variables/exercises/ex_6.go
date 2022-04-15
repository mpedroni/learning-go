package main

import "fmt"

/*
	- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
	- Demonstre estes valores.
*/

const (
	_ = iota + 2022
	w
	x
	y
	z
)

func main() {
	fmt.Println(w, x, y, z)
}
