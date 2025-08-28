package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Learning Web Services")
	resp, _ := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	fmt.Println("", resp)
	defer resp.Body.Close()
	// Read the response body
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("Response Body:", string(body))
}
