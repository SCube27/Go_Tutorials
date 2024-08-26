package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Starting of the program")
	sum := add(5, 6)
	defer fmt.Println("Sum is", sum) // the statement with defer executes at the last just before the end of program
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")
}

// if there are multiple defer keywords then every defer statement is pushed in a call stack
// once the program is over the items from the stack are printed one by one
// output will be:
/*
Starting of the program
End of the program
Middle of the program
Sum is 11
*/
