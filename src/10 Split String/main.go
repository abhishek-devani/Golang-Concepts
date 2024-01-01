package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "Hi, Welcome, Mr. Robot"
	fmt.Println("Before split", str1)

	strArr := strings.Split(str1, ",")
	fmt.Println(strArr)

	fmt.Println(strings.ToUpper(str1)) // Modify Original String
	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToTitle(str1)) // Copy of String

}