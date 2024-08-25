package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Starting with strings")

	str := "abc,def,ghi"
	parts := strings.Split(str, ",") // splitting the string str on the basis of comma (,)
	fmt.Println(parts)

	nums := "one two three four five three six seven three"
	count := strings.Count(nums, "three") // counts the number of occurrences of a substring from the given string
	fmt.Println("The count of three is", count)

	data := "    Hello Go!    "
	trimmed := strings.TrimSpace(data) // trims the unwanted space from the given string
	fmt.Println("Trimmed string :", trimmed)

	str1 := "Sahil"
	str2 := "Shah"
	result := strings.Join([]string{str1, str2}, " ") // joins two strings and adds the separator string in between
	res := str1 + " " + str2                          // this also works though
	fmt.Println(result)
	fmt.Println(res)
}
