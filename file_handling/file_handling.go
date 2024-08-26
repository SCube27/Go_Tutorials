package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		// Step 1: Create a file
		file, err := os.Create("sample.txt")

		if err != nil { // handling error
			fmt.Println("Error while creating file", err)
			return
		}
		defer file.Close() // when the function execution ends close the file

		// Step 2: Write content in the created file
		content := "Some content to enter the file!"
		fmt.Println("File created!")
		_, error := io.WriteString(file, content+"\n") // writing the content string in the file

		if error != nil {
			fmt.Println("Error while writing file", error)
			return
		}
	*/

	/*
		// Step 3: Reading the file (Using Buffer)
		file, err := os.Open("sample.txt")

		if err != nil {
			fmt.Println("Error occured while opening the file", err)
			return
		}
		defer file.Close()

		// creating a buffer to read the file content
		buffer := make([]byte, 1024)

		for {
			n, err := file.Read(buffer) // reading from the file and storing in the buffer

			if err == io.EOF { // if reached end of file while reading the file
				break
			}
			if err != nil {
				fmt.Println("Error while reading the file", err)
				return
			}

			// process the read content
			fmt.Println(string(buffer[:n]))
		}
	*/

	// Step 3: Reading the file (Using ioutil package functions)
	content, err := os.ReadFile("sample.txt")

	if err != nil {
		fmt.Println("Error reading the file!", err)
		return
	}

	fmt.Println(string(content))
}

/*
1. The method 2 has the issue that it loads all the data of file into the memory all at once, so incase if the file is too large then the memory can't handle the file data and may raise a memory overflow error.
2. So reading the data in chunks and storing in buffer and then printing the chunk is a better option to go with.
*/
