package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Apple,orange,Banana"
	parts := strings.Split(data, ",  	")
	fmt.Println(parts)

	str := "One two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("Count of 'two':", count)

	str = "Hello, Go          "
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed Characters :", trimmed)

	str1 := "Kunsh"
	str2 := "Jain"
	str = strings.Join([]string{str1, str2}, " ")
	fmt.Println("Joined String :", str)
}
