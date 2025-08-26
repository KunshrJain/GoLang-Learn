package main

import "fmt"

func add(a, b int) int {
	result := a + b
	return result
}

func main() {
	fmt.Println("Starting -")
	sum := add(5, 10)
	defer fmt.Println("Sum:", sum)
	defer fmt.Println("Middle Part")
	fmt.Println("Ending -")
}
