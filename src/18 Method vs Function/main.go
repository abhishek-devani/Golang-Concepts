package main

import "fmt"

type data int

type user struct {
	name string
	email string
}

func (v1 data) adder(v2 data) data {
	return v1 + v2
}

func main() {

	// Example 1
	a := data(10)
	b := data(20)
	res := a.adder(b)
	fmt.Println(res)

	// Example 2
	student := user {
		name: "abc",
		email: "xyz@xyz.com",
	}

	fmt.Println("user's name: ", student.name)
	fmt.Println("user's email(before): ", student.email)

	// Creating a pointer
	p := &student

	// Calling the show method
	p.correctEmail("abc@abc.com")
	fmt.Println("user's name: ", student.name)
	fmt.Println("user's email(after): ", student.email)

}

func (student *user) correctEmail(email string) {
	student.email = email
}