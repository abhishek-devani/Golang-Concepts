package main

import "fmt"

// Address struct represents a physical address
type Address struct {
    Street  string
    City    string
    Country string
}

// Person struct represents a person with an embedded Address
type Person struct {
    Name    string
    Age     int
    Contact struct {
        Email    string
        Phone    string
        Residence Address // Embedded Address struct
    }
}

func main() {
    // Create a Person instance
    johnDoe := Person{
        Name: "John Doe",
        Age:  30,
        Contact: struct {
            Email    string
            Phone    string
            Residence Address
        }{
            Email: "john.doe@example.com",
            Phone: "+1234567890",
            Residence: Address {
                Street:  "123 Main St",
                City:    "Anytown",
                Country: "USA",
            },
        },
    }

    // Access fields of the nested struct
    fmt.Println("Name:", johnDoe.Name)
    fmt.Println("Age:", johnDoe.Age)
    fmt.Println("Email:", johnDoe.Contact.Email)
    fmt.Println("Phone:", johnDoe.Contact.Phone)
    fmt.Println("Street:", johnDoe.Contact.Residence.Street)
    fmt.Println("City:", johnDoe.Contact.Residence.City)
    fmt.Println("Country:", johnDoe.Contact.Residence.Country)
}
