package main

import "fmt"

type Person struct {
	name   string
	age    int
	ismale bool
}

type Contact struct {
	email string
	phone string
}

type Address struct {
	houseNo int
	area    string
	state   string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
	Person_Company string
}

// so all these structs can make a complex data structure to hold a whole employee database

func main() {
	fmt.Println("Starting with structs")

	// method 1 to store values in struct
	var p1 Person
	p1.name = "Sahil"
	p1.age = 22
	p1.ismale = true

	fmt.Println("Sahil's Data :", p1)

	// method 2 to store values in struct
	p2 := Person{
		name:   "Shubh",
		age:    21,
		ismale: false,
	}

	fmt.Println("Shubh's Data :", p2)

	// method 3 to store values in struct (new keyword)
	var p3 = new(Person)
	p3.name = "Sarthak"
	p3.age = 23
	p3.ismale = true

	fmt.Println("Sarthak's Data :", *p3) // the new keyword makes a pointer which points to the location where new data is stored
	// hence printing a pointer gives the actual data over here

	// we can access single values in the struct using the dot operator
	fmt.Println("The Age of Sahil is", p1.age)

	// Creating a whole data of a single employee using complex structs
	var employee1 Employee

	employee1.Person_Details = Person{
		name:   "Kaushal",
		age:    25,
		ismale: true,
	}

	employee1.Person_Contact.email = "kash52@gmail.com"
	employee1.Person_Contact.phone = "9504603321"

	employee1.Person_Address = Address{
		houseNo: 1,
		area:    "Connaught Place",
		state:   "New Delhi",
	}

	employee1.Person_Company = "Uber"

	fmt.Println("Details of Employee 1", employee1)

}
