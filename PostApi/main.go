package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformGetReq() {
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
	fmt.Printf("Formatted Response : %+v\n", todoo)

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

func PerformPostReq() {
	todoo := todo{
		UserID:    23,
		Title:     "New Todo Created",
		Completed: true,
	}

	jsonData, err := json.Marshal(todoo)
	if err != nil {
		fmt.Println("Json opening Error")
	}

	jsonString := string(jsonData)

	//converting String data to reader
	jsonReader := strings.NewReader(jsonString)

	MyUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(MyUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error in Post Request")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response Body:", string(body))
}

func main() {
	PerformGetReq()
	PerformPostReq()
}
