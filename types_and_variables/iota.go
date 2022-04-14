package main

import "fmt"

// `iota` generates successive untyped integers

// OBS: `untyped` because const are untyped until it be used;
// it will be assume the first "convenient" number type (float or int) depending of the context
const (
	a = iota
	b = iota
	c = iota
	_ = iota // ignoring value 3
	e = iota
)

const (
	// the first expression will be replied to subsequent constants
	w = (iota + 1) * 10
	x
	y
	z
)

func main() {
	fmt.Println(a, b, c, e) // out: 0 2 3 4

	fmt.Println(w, x, y, z) // out: 10 20 30 40
}
