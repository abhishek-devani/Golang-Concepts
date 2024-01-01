package main

import "fmt"

func main() {
	a := a()
	b := b()

	// Variadic Function
	sum := variadic(a, b, 20)
	fmt.Println(sum)

}

func variadic(v ...int) int {
	sum := 0
	for _, val := range v {
		sum += val
	}
	return sum
}

func a() int {
	return 10
}

func b() int {
	return 20
}
