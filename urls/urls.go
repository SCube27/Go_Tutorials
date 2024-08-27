package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Starting with URL handling")

	link := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

	parsedURL, err := url.Parse(link) // parsing / converting the raw string into a URL type

	if err != nil {
		fmt.Println("Error while parsing the URL", err)
		return
	}

	// breaking down the URL structure
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)

	// modifying the url components
	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=sash21"

	// constructing a new url string from the url object
	newUrl := parsedURL.String()
	fmt.Println("New URL:", newUrl)
}
