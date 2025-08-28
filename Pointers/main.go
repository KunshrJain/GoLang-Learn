package main

import "fmt"

func main() {
	var num int = 42
	var ptr *int = &num

	fmt.Printf("Value of num: %d\n", num)
	fmt.Printf("Address of num: %p\n", &num)
	fmt.Printf("Value of ptr: %p\n", ptr)
	fmt.Printf("Value pointed to by ptr: %d\n", *ptr)

	*ptr = 100
	fmt.Printf("New value of num: %d\n", num)
}
