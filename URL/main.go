package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL's")
	myurl := "https://example.com:8080"

	parsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Fetching URL:", myurl)
	fmt.Println("Fetching URL:", parsedURL)

	// things like scheme host path rawQuery can also be parsed
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
}
