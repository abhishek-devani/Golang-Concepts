package custom

import (
	"fmt"
	// "reflect"
)

func PrintWithAnyArgs[T any](a ...T) {
	fmt.Println(a)
}

func PrintStringValue(a string) {
	fmt.Println(a)
}

func PrintAnyValue(val ...interface{}) {
	for _, v := range val {
		printValue(v)
	}
}

func printValue(value interface{}) {
	// v := reflect.ValueOf(value)
	// fmt.Printf("%v ", v.Interface())
	fmt.Print(value, " ")
}