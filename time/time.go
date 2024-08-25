package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting with time package")

	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	formattedTime := currentTime.Format("02-01-2006, 15:04:05") // formats the current time in following format
	ft1 := currentTime.Format("2006/01/02, 03:04 PM")           // another format
	fmt.Println(formattedTime)
	fmt.Println(ft1)

	layoutDate := "02/01/2006"
	todayDate := "25/08/2024"
	formatDate, _ := time.Parse(layoutDate, todayDate) // converts the date entered according to the layout date
	fmt.Println(formatDate)
}
