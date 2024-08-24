package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 10
	fmt.Printf("The type of variable is %T\n", num)

	var data float64 = float64(num) // Type conversion from int to float
	fmt.Printf("The type of variable is %T\n", data)

	num = 123
	str := strconv.Itoa(num) // converting int to string
	fmt.Printf("The type of str is %T\n", str)

	num_string := "1234"
	num_int, _ := strconv.Atoi(num_string) // converting string to int
	fmt.Printf("The type of num_int is %T\n", num_int)

	number_str := "3.14"
	number_float, _ := strconv.ParseFloat(number_str, 64) // converting string to float
	fmt.Printf("The type of number_float is %T\n", number_float)
}
