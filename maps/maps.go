package main

import "fmt"

func main() {
	fmt.Println("Starting with maps")

	// mapping a person name (string) with their company (string)
	personCompany := make(map[string]string)

	// here the person name is the key and the company name is the value
	personCompany["Sahil"] = "Cloudanix"
	personCompany["Shubh"] = "GIC"
	personCompany["Sarthak"] = "Pubmatic"

	fmt.Println("The company of Sahil is", personCompany["Sahil"])

	personCompany["Sarthak"] = "PhonePe" // The values in maps can be changed
	fmt.Println("The company of Sarthak is", personCompany["Sarthak"])

	// The values from the map can be deleted too
	delete(personCompany, "Sarthak")
	fmt.Println("The company of Sarthak is", personCompany["Sarthak"]) // this returns an empty string

	// A key mapping from the map returns two things, 1. its value and 2. if it exists in the map or not
	company, exists := personCompany["Shubh"]
	fmt.Println("Shubh exists in the map?", exists) // returns either true or false
	fmt.Println("The company of Shubh is", company)

	// since range returns the index and value, in maps the index is key and value is value for the key
	for index, value := range personCompany {
		fmt.Printf("The person name is %s and his company is %s\n", index, value)
	}
}
