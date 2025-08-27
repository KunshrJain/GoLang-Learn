package main

import "fmt"

func add(a, b int) (int, string) {
	return a + b, "done"
}

func mul(a, b int) (int, string) {
	if b == 0 {
		return 0, "Error hai bhai isme"
	}
	return a / b, "done"
}

func main() {
	ans, string := add(3, 4)
	fmt.Println("Result:", ans, string)
	ans1, string1 := mul(12, 0)
	fmt.Println("Result:", ans1, string1)
	ans1, string1 = mul(12, 3)
	fmt.Println("Result:", ans1, string1)
}
