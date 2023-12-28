// It is same as lambda function in java

package main

import "fmt"

func main() {
	
	func() {
		fmt.Println("anonymous function")
	}()

	val := func () int {
		return 10;
	}
	fmt.Println(val())

	func (v ...string)  {
		fmt.Println(v)
	}("a", "b")

	result := operate(3, 4, func(a, b int) int {
		return a * b
	})

	fmt.Println("Operator function: ", result)

}

func operate(a, b int, op func(int, int) int ) int {
	return op(a, b)
}