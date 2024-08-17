package main

import "fmt"

func simpleFunction() {
	fmt.Println("Inside the simple function")
}

func addition(a int, b int) int { // the datatype after the parameters is the return type of the function
	return a + b
}

func product(a int, b int) (ans int) { // already specified the variable that will be returned after the function completes
	ans = a * b
	return // only writing return works
	// return a + b also works
}

func main() {
	fmt.Println("Starting with functions in GO!")
	simpleFunction()

	ans := addition(21, 26)
	fmt.Println("Addition is", ans)

	value := product(21, 26)
	fmt.Println("Product is", value)
}
