// In Go, defer is a keyword used to schedule a function call to be run after the function that contains the defer statement completes
// It's often used to ensure that resources are released or cleanup operations are performed regardless of whether the surrounding 
//	function exits normally or with an error

package main

import "fmt"

func main() {

	fmt.Println("Start")

	// Follows LIFO (stack)
	defer fmt.Println("End")

	defer func (a, b int) {
		fmt.Println(a + b)
	}(5, 6)

	defer func (a, b int) {
		fmt.Println(a * b)
	}(3, 4)
	

}