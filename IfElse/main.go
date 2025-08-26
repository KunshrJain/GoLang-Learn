package main

import "fmt"

func main() {
	x := 25
	if x > 5 {
		fmt.Println("x>5")
	} else {
		fmt.Println("x<5")
	}
	if x < 5 {
		fmt.Println("x<5")
	} else if x == 25 {
		fmt.Println("x=25")
	} else {
		fmt.Println("x>25")
	}
}
