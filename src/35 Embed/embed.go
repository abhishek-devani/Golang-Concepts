package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.json
var data string

func main() {

	fmt.Println(data)

}