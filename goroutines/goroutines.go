package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello World!")
	time.Sleep(2000 * time.Millisecond) // simulating some work
	fmt.Println("sayHello function ended!")
}

func sayHi() {
	fmt.Println("Hi Sahil :)")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("starting with goroutines!")
	go sayHello() // this goes on for its execution separately until then the further lines are executed
	go sayHi()

	// wait for a moment until all the goroutines are executed
	time.Sleep(3000 * time.Millisecond)
}
