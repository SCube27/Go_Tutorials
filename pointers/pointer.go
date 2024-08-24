package main

import "fmt"

func modifyValuebyReference(num *int) {
	*num = *num * 20
}

func main() {
	fmt.Println("Starting with pointers!")

	// var val int = 8
	// var ptr *int // declaring a pointer
	// ptr = &val // marking the pointer to the val variable

	val := 8
	ptr := &val

	fmt.Println("The pointer is pointing to the value", *ptr)
	fmt.Println("The address the pointer is pointing to is", ptr)

	value := 10
	modifyValuebyReference(&value)
	fmt.Println("The data in value variable is now changed as", value)
}
