package main

import "fmt"

func main() {
	x := 10
	if x > 0 && x == 10 {
		fmt.Println("x equals to", x)
	} else {
		fmt.Println("x is smaller than", x)
	}

	ans := 21
	if ans > 21 {
		fmt.Println("ans is greater than", ans)
	} else if ans == 21 {
		fmt.Println("ans is equal to", ans)
	} else {
		fmt.Println("ans is smaller than", ans)
	}
}
