package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content := "This is a test file.\nThis file is created using GoLang."
	_, err = io.WriteString(file, content)
	if err != nil {
		fmt.Println("Error while writing file", err)
		return
	}

	file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println("Error while opening file", err)
		return
	}
	defer file.Close()

	fmt.Println("File content:")

	data, _ := os.ReadFile("test.txt")
	fmt.Println(string(data))

	for {
		buff, err := file.Read(data)
		if err != nil {
			break
		}
		fmt.Print(string(data[:buff]))
	}

	//if you want the entire file in the memory we use ioutill

	// content, _ = ioutil.ReadFile("test.txt")
	// fmt.Println(string(content))

}
