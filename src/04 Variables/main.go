package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int
	var b = "10"
	c := 20

	fmt.Printf("a is a type of %v and b is a type of %v, and c is a type of %v\n",
		reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))

}