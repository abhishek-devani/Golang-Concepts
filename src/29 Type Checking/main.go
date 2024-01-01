package main

import "fmt"

func main() {

	var val interface{} = "abc"

	if str, ok := val.(string); ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Not a string")
	}

}
