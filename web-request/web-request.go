package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Starting web request with native packages")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // making a GET request to fetch data

	if err != nil {
		fmt.Println("Error occured while fetching the data", err)
		return
	}
	defer res.Body.Close() // once processes are done close the connection with API

	// reading the response body
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading the response body", err)
		return
	}

	fmt.Println("Response:", string(data)) // since the data is in form of bytes we've to convert it into string
}
