package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "@@Some@ @string+@+@"

	fmt.Println(strings.Trim(str1, "@"))
	fmt.Println(strings.TrimLeft(str1, "@"))
	fmt.Println(strings.TrimRight(str1, "@"))

	str2 := "		Some String		"
	fmt.Println("Before trim: ", str2)
	fmt.Println(strings.TrimSpace(str2))

	// Trim suffix
	str3 := "Hello World"
	fmt.Println(strings.TrimSuffix(str3, "World"))

	// Trim prefix
	str4 := "Hello World"
	fmt.Println(strings.TrimPrefix(str4, "Hello"))

}