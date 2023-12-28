package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {

	// arr := []string{"a", "b", "c"}

	// fmt.Println(arr)
	// fmt.Println("Length: " ,len(arr))
	// fmt.Println("Capacity: ", cap(arr))

	// arr = append(arr, "d")

	// fmt.Println(arr)
	// fmt.Println("Length: ", len(arr))
	// fmt.Println("Capacity: ", cap(arr))

	// arr = append(arr, "e")

	// fmt.Println(arr)
	// fmt.Println("Length: ", len(arr))
	// fmt.Println("Capacity: ", cap(arr))

	// mySlice := arr[1:3]
	// fmt.Println(mySlice)
	// fmt.Println("Length: ", len(mySlice))
	// fmt.Println("Capacity: ", cap(mySlice))

	// mySlice[0] = "10"
	// fmt.Println(arr)

	// Sorting slices
	s1 := []int{-23, 567, -34, 6, 0, 12, -5}
	fmt.Println("Before Sorting: ", s1)
	
	sort.Ints(s1)
	fmt.Println("After Sorting: ", s1)

	fmt.Println("Check ints are sorted or not: ", sort.IntsAreSorted(s1))

	fmt.Println(reflect.TypeOf(s1)) // Output: []int
	fmt.Println(reflect.TypeOf(s1).Kind()) // Output: slice

	// Byte Slice
	s2 := []byte{'a', 'b', '#', '0'}

	for _, v := range s2 {
		fmt.Println(string(v))
	}

}