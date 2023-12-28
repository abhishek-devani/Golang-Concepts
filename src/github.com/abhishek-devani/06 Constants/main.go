package main

import "fmt"

const (
	first = iota
	second
	third
	forth = 'a'
	fifth
)

func main() {
	fmt.Println(first, second, third, forth, fifth)

	// rune is an alias for int32. They represent the same underlying type.
	// The use of rune is more semantic when working with Unicode characters.
	// Every character has sky code
	if rune('a') == third {
		fmt.Println("true")
	}

}