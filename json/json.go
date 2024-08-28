package main

import (
	"encoding/json"
	"fmt"
)

// the fields present after datatypes are the name of the fields that will be shown when struct is converted to JSON
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{Name: "Shubh", Age: 22, IsAdult: true}
	fmt.Println("Person data is", person)

	// convert the struct object into JSON (Marshaling)
	jsonData, err := json.Marshal(person) // the returned jsonData over here is an array of bytes so convert into string

	if err != nil {
		fmt.Println("Error while encoding the struct", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData)) // converting the byte array into string

	// convert the JSON object into struct (Unmarshaling)
	var Shubh Person
	error := json.Unmarshal(jsonData, &Shubh) // expects the byte array and address of a struct object where data is to be filled

	if error != nil {
		fmt.Println("Error while decoding the JSON", error)
		return
	}

	fmt.Println("Shubh Data:", Shubh)
}
