package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	fmt.Println("We are learning JSON")
	person := Person{Name: "John", Age: 30, IsAdult: true}
	fmt.Println("Person Object", person)

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	// Unmarshalling

	var person2 Person
	err = json.Unmarshal(jsonData, &person2)

	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Unmarshalled Person Object:", person2)
}
