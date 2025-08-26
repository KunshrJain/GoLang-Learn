package main

import "fmt"

func main() {
	fmt.Println("We are learning Array in GoLang")

	var name [5]string
	name[0] = "Kunsh"
	name[1] = "Manan"
	name[2] = "Arnav"
	name[3] = "Akshat"

	fmt.Println("Names are:", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are:", numbers)
	fmt.Println("Length of this is:", len(numbers))
}
