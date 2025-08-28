package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning Crud API")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Either decode directly:
	var todoo todo
	err = json.NewDecoder(resp.Body).Decode(&todoo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Printf("Formatted Response (direct decode): %+v\n", todoo)

	// OR if you want to print raw + formatted:
	// Reset request again
	resp, err = http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Raw Response Body:", string(body))
}
