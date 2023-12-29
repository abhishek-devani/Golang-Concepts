package main

import "fmt"

func main() {

	x := 0
	y := 20
	printAllOperations(x, y)

}

func printAllOperations(x, y int) {
	defer func () {
		if r := recover(); r != nil {
			fmt.Println("recovering from panic")
			printOtherOperations(x, y)
		}
	}()
	sum, divide, multiply := x+y, y/x, x*y
	fmt.Printf("sum=%v, divide=%v, multiply=%v \n", sum, divide, multiply)

}

func printOtherOperations(x, y int) {
	sum, mul := x + y, x*y
	fmt.Printf("sum=%v, multiply=%v \n", sum, mul)
}