package main

import "fmt"

func main() {
	// Outer function that returns a closure
	adder := adderFunction(10)

	// Using the closure to add values
	result1 := adder(5)
	result2 := adder(7)

	fmt.Println("Result 1:", result1) // Output: 15 (10 + 5)
	fmt.Println("Result 2:", result2) // Output: 17 (10 + 7)
}

// adderFunction returns a closure that adds the provided value to its argument
func adderFunction(base int) func(int) int {
	// Inner function (closure) that captures the 'base' variable
	return func(value int) int {
		return base + value
	}
}
