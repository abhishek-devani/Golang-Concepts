// Init function does not take argument or does not return anything
// It is used to initialize something before main function runs
// For Ex. we need DB connection before or application starts
// We can have multiple Init function as well. It will run in order

package main

import "fmt"

func init() {
	fmt.Println("Init first")
}

func init() {
	fmt.Println("Init Second")
}

func main() {
	fmt.Println("Main Func")
}