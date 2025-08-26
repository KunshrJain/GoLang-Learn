package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//infinite Loop

	/*	for {
			fmt.Println("Infinite loop")
			counter++
		}
	*/

	number := []int{11, 42, 83, 14, 23}
	for index, val := range number {
		fmt.Println("Index:", index, "Value:", val)
	}

	str := "Hello World"
	fmt.Printf("String: %c\n", str[0])
}
