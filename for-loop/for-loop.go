package main

import "fmt"

func main() {
	// Type 1 normal loop
	for i := 0; i < 10; i++ { // normal for loop
		fmt.Println("Number:", i)
	}

	// Type 2 (Imitating while loop)
	counter := 0
	for { // starting a infinite loop
		fmt.Println("Infi Loop")
		counter++

		if counter == 10 { // adding a condition to break the loop
			break
		}
	}

	// Type 3 (using range keyword)
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers { // looping over an iterable object (slice) using the range keyword
		fmt.Println("Index:", index, "Value:", value)
	}

	data := "This is Sahil Shah"
	for index, char := range data {
		fmt.Printf("Index: %d Character: %c\n", index, char)
	}
}
