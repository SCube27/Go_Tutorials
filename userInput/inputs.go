package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Please Enter Your Name!")

	// var name string
	// fmt.Scan(&name) // using & to store the console input on the address of name variable
	// fmt.Println("Hey Hello", name)

	reader := bufio.NewReader(os.Stdin) // this reads from the standard input (Stdin) from the os package
	name, _ := reader.ReadString('\n')  // reads the string till gets a new line
	fmt.Println("Hey Hello", name)
}
