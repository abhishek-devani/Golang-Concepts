package main

import "fmt"

const (
	first = iota
	second
	third
	forth = 'a'
	fifth
)

func main() {

	getType("Data")
	getType(25)

}

func getType(a interface{}) {

	switch t := a.(type) {
	case int64:
		fmt.Println("Type is an integer:", t)
	case float64:
		fmt.Println("Type is a float:", t)
	case string:
		fmt.Println("Type is a string:", t)
	case nil:
		fmt.Println("Type is nil:", t)
	case bool:
		fmt.Println("Type is a bool:", t)
	default:
		fmt.Println("Type is unknown:", t)
	}

}