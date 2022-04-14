package main

import "fmt"

// strings with backtick keeps formatation
func Backtick() {
	s := `Hello, 
		batman`
	fmt.Println(s)
}

func StringAsByteSlice() {
	s := "ascii abcd utf-8 蝙蝠俠"

	/*
	 *
	 * for with `range` iterates by character
	 * s[i] iterates by byte
	 *
	 * some utf-8 characters needs more than 1 byte to be represented, so the second for prints the "wrong" characters
	 *
	 * -- read about utf-8 and ascii for better compreension --
	 */
	for _, v := range s {
		fmt.Printf("%b\t%T\t%#U\t%#x\n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%b\t%T\t%#U\t%#x\n", s[i], s[i], s[i], s[i])
	}
}

func main() {
	Backtick()
	StringAsByteSlice()
}
