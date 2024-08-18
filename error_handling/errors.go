package main

import "fmt"

func divide(a, b float64) (float64, error) { // returning two datatypes
	if b == 0 {
		return 0, fmt.Errorf("denominator can't be 0") // returns 0 and also an error message using the Errorf() method
	}
	return a / b, nil // since no error in this case return nil
}

func multiply(a, b int) (int, string) { // since using string as second return type, we've to return string
	if b == 0 || a == 0 {
		return 0, "The ans will be"
	}
	return a * b, "Multiplication is"
}

func main() {
	fmt.Println("Error Handling in the functions!")

	// ans, _ := divide(21, 0) // in this case the _ could be ignored and won't be necessary to use

	ans, err := divide(21, 3) // since storing error in a variable err, it needs to be used
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Division is", ans)

	data, msg := multiply(14, 0)
	fmt.Println(msg, data)
}
