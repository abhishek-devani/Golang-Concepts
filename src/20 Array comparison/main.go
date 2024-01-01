package main

import "fmt"

func main() {

	q1 := [3]int{9, 7, 6}

	// size of an array determines at compile time
	q2 := [...]int{9, 7, 6}
	q3 := [3]int{9, 7, 3}

	fmt.Println(q1 == q2)
	fmt.Println(q2 == q3)
	fmt.Println(q1 == q3)

}