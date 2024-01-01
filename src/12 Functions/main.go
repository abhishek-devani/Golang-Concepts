package main

import "fmt"

func main() {
	a := a()
	b := b()

	fmt.Println(a + b)

	passByValue(a, b)
	fmt.Println(passByRef(&a, &b))

}

func a() int {
	return 10
}

func b() int {
	return 20
}

func passByValue(a, b int) {
	fmt.Println(a+b)
}

func passByRef(a, b *int) int {
	return *a + *b
}