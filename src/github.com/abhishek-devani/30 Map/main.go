package main

import (
	"fmt"
)

func main() {

	mp := map[int]string{1: "a", 2: "b"}

	fmt.Println(mp)

	if mp[1] == "a" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	for k, v := range mp {
		fmt.Println(k, v)
	}

	mp[5] = "c"

	fmt.Println("Before Deletion " ,mp)
	delete(mp, 1)

	val, ok := mp[1]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("After Deletion " ,mp)
	}

}