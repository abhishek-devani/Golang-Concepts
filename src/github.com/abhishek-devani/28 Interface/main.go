package main

import "fmt"

// Defining an interface named Writer
type Writer interface {
	Write(data string) error
}

// Implementing the Writer interface for the File type
type File struct {
	Name string
}

// The File type satisfies the Writer interface by implementing the Write method
func (f File) Write(data string) error {
	fmt.Printf("Writing to file %s: %s\n", f.Name, data)
	// Logic for writing to a file goes here
	return nil
}

// Another implementation of the Writer interface for the Console type
type Console struct{}

// The Console type satisfies the Writer interface by implementing the Write method
func (c Console) Write(data string) error {
	fmt.Println("Writing to console:", data)
	// Logic for writing to the console goes here
	return nil
}

func main() {
	// Using the Writer interface as a parameter type
	writeDataToFile := func(w Writer, data string) {
		w.Write(data)
	}

	// Creating instances of types that implement the Writer interface
	file := File{Name: "example.txt"}
	console := Console{}

	// file.Write("Hello File!")
	// console.Write("Hello Console")

	// Calling the function with different implementations of the Writer interface
	writeDataToFile(file, "Hello, File!")
	writeDataToFile(console, "Hello, Console!")
}
