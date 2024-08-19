package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	numbers = append(numbers, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15) // size can be changed even after declaration

	fmt.Println(numbers)
	fmt.Println(len(numbers))

	// cars := []string{"koenigsegg jesko", "bugatti chiron", "porsche 911 gt3 rs", "mclaren 765lt"}
	// fmt.Println(cars)

	langs := make([]string, 5, 7) // creating a slice using the make() function
	fmt.Println(langs)
	fmt.Println(cap(langs))

	langs = append(langs, "Golang")
	langs = append(langs, "Python")
	langs = append(langs, "JavaScript") // at this point the value exceeds the slice capacity, so it doubles its capacity

	fmt.Println(langs)
	fmt.Println(cap(langs)) // the capacity doubles down
}
