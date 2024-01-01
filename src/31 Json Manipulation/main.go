package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var wrapData map[string]interface{}
	data :=
	`
	{
		"person": {
		  "name": "John Doe",
		  "age": 30,
		  "isStudent": false,
		  "address": {
			"city": "New York",
			"zipCode": "10001"
		  },
		  "languages": ["English", "Spanish"],
		  "grades": {
			"math": 95,
			"science": 88,
			"history": 92
		  }
		},
		"company": {
		  "name": "ABC Corp",
		  "employees": [
			{
			  "name": "Alice",
			  "position": "Software Engineer",
			  "salary": 90000
			},
			{
			  "name": "Bob",
			  "position": "Data Scientist",
			  "salary": 95000
			}
		  ],
		  "location": "San Francisco"
		}
	}	  
	`

	newEmployeeJson := 
	`
		{
			"name": "abc",
			"position": "SE",
			"salary": 1000
		}
	`

	err := json.Unmarshal([]byte(data), &wrapData)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(wrapData)

	person, _ := wrapData["person"].(map[string]interface{})

	// Printing languages
	fmt.Println("Printing languages")
	languages, _ := person["languages"].([]interface{})
	fmt.Println(len(languages))
	for _, val := range languages {
		fmt.Println(val)
	}

	// Printing grades
	fmt.Println("\nPrinting grades")
	grades, _ := person["grades"].(map[string]interface{})
	fmt.Println(len(grades))
	for key, val := range grades {
		fmt.Println(key, " : ", val)
	}

	// Append newEmployee to employee
	var newEmployee map[string]interface{}
	err = json.Unmarshal([]byte(newEmployeeJson), &newEmployee)
	if err != nil {
		fmt.Println(err)
	}

	company, _ := wrapData["company"].(map[string]interface{})
	employees, _ := company["employees"].([]interface{})

	company["employees"] = append(employees, newEmployee)

	// Print the data in formatted way
	updatedData, err := json.MarshalIndent(wrapData, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nUpdated Data:")
	fmt.Println(string(updatedData))

}