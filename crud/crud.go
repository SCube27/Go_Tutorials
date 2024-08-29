package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD...")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error while getting the data", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error to get the data", res.StatusCode)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	fmt.Println("Error while reading the data", err)
	// 	return
	// }

	// fmt.Println("Data:", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo) // the new decoder did all the above commented task in one line

	if err != nil {
		fmt.Println("Error while storing the data in struct", err)
		return
	}

	fmt.Println("Todo Task:", todo)
}
