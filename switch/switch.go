package main

import "fmt"

func main() {
	lang := 1

	switch lang {
	case 1:
		fmt.Println("C++")
	case 2:
		fmt.Println("Python")
	case 3:
		fmt.Println("JavaScript")
	case 4:
		fmt.Println("TypeScript")
	case 5:
		fmt.Println("Golang")
	default:
		fmt.Println("No language")
	}

	// Can add multiple values in the switch case
	month := "April"

	switch month {
	case "October", "November", "December", "January":
		fmt.Println("Winter is coming")

	case "March", "April", "May":
		fmt.Println("Summer is here")

	case "June", "July", "August", "September":
		fmt.Println("Monsoon coming soon")
	}

	// Can add expressions in the switch case
	temperature := 26

	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 18:
		fmt.Println("Cold")
	case temperature > 25 && temperature < 34:
		fmt.Println("Warm")
	case temperature > 34 && temperature < 45:
		fmt.Println("Hot")
	}
}
