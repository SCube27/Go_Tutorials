package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getRequest() {
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

func postRequest() {
	todo := Todo{
		UserId:    21,
		Title:     "Complete all the tasks",
		Completed: false,
	}

	// convert the todo in JSON format for sending the post request
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error occured while marshaling", err)
		return
	}

	// convert the jsonData(byte slice) to string
	jsonString := string(jsonData)

	// converting the jsonString to reader as the post method requires a reader object
	jsonReader := strings.NewReader(jsonString)

	url := "https://jsonplaceholder.typicode.com/todos"

	// calling the post method
	res, err := http.Post(url, "application/json", jsonReader)

	if err != nil {
		fmt.Println("Error occured while reading response", err)
		return
	}

	// reading the http response object and converting it in byte slice
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error occured while reading http response", err)
		return
	}

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response:", string(data)) // converting the byte slice in string
}

func updateRequest() {
	todo := Todo{
		UserId:    26,
		Title:     "Be Consistent!",
		Completed: true,
	}

	url := "https://jsonplaceholder.typicode.com/todos/1"

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error occured while marshaling", err)
		return
	}

	jsonString := string(jsonData)
	jsonReader := strings.NewReader(jsonString)

	req, err := http.NewRequest("PUT", url, jsonReader) // creating a new request to send

	if err != nil {
		fmt.Println("Error occured in PUT Request", err)
		return
	}

	req.Header.Set("Content-type", "application/json") // setting the headers

	// send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error occured while sending request", err)
		return
	}
	defer res.Body.Close()

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error occured while reading response", err)
		return
	}

	fmt.Println("Rsponse Status:", res.Status)
	fmt.Println("Response Data:", string(resData))
}

func deleteRequest() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		fmt.Println("Error while creating request", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error while getting response", err)
		return
	}

	resData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error while reading response", err)
		return
	}

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response Data:", string(resData))
}

func main() {
	fmt.Println("Learning CRUD...")
	getRequest()
	postRequest()
	updateRequest()
	deleteRequest()
}
