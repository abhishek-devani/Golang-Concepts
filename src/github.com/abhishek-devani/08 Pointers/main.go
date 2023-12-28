package main

import (
	"fmt"
)

func main() {
	
	var a *int
	var b *int

	a = ptr(10)
	b = ptr(10)

	c := *a + *b

	fmt.Println(c)

	a = generic(10)
	b = generic(10)

	fmt.Println(*a + *b)

}

// Generic
func generic[T any](v T)*T {
	return &v
}

func ptr(v int) *int {
	return &v
}