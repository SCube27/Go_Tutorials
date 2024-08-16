package main

import (
	"fmt"
	myutil "tutorials/utils"
)

func main() {
	fmt.Println("Hello World in GO!")
	myutil.PrintMessage("Hello World from myutil")

	var name string = "Sahil"
	var version = 2.0
	fmt.Println(name)
	fmt.Println(version) // datatype decided during compile time

	var money int = 700000000
	fmt.Println(money)

	var dimensions float64 = 26.21
	fmt.Println(dimensions)

	var decision bool = true
	decision = false // can be reassigned
	fmt.Println(decision)

	const pi float64 = 3.14 // can't be reassigned
	fmt.Println("Value of pi:", pi)

	username := "Prince" // can be used when taking a value from a called function
	fmt.Println(username)

	var Public = "Data is public"   // variable name starts from capital hence can be exported
	var private = "Data is private" // can't be exported
	fmt.Println(Public)
	fmt.Println(private)
}
