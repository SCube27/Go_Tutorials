package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")

	var langs [5]string // declaration of an array

	langs[0] = "C++"
	langs[1] = "Python"
	langs[2] = "JavaScript"
	langs[3] = "TypeScript"
	langs[4] = "Golang"

	fmt.Println("Programming languages are", langs)

	var numbers = [5]int{1, 2, 3, 4, 5} // initializing the array during its declaration
	fmt.Println("Numbers are", numbers)

	fmt.Println("Length of Langs array is", len(langs)) // length of array

	fmt.Println("Programming Language at the index 4 is:", langs[4]) // accessing values of array
}
