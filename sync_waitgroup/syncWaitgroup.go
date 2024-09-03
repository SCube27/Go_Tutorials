package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // once the function execution is done mark this goroutine as done
	fmt.Printf("Worker %d started\n", i)
	// some task is working over here
	fmt.Printf("Worker %d ended\n", i)
}

func main() {
	// fmt.Println("Starting with sync waitgroup!")

	var wg sync.WaitGroup

	// start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // increment the waitgroup conter by for for every goroutine
		go worker(i, &wg)
	}

	wg.Wait() // waiting for the goroutines till all are done with execution
	fmt.Println("All workers are completed with task!")
}
