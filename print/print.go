package main

import "fmt"

func main() {
	age := 22
	name := "Sahil"
	height := 5.81

	fmt.Println("Age:", age, "Name:", name, "Height:", height)

	// fmt.Printf("Age is %d\n", age)
	// fmt.Printf("Height is %.3f\n", height)
	// fmt.Printf("Name is %s\n", name)
	fmt.Printf("Type of name variable is %T\n", name)

	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height) // the format specifiers basically replace themselves with the variables (order of variables should be relative to format specifiers)
}
