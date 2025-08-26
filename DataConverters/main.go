package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Number is", num)
	fmt.Printf("Type of Number is %T\n", num)
	var data float64 = float64(num)
	fmt.Println("Data is", data)
	fmt.Printf("Type of Number is %T\n", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("String is", str)
	fmt.Printf("Type of String is %T\n", str)
	str1, _ := strconv.Atoi(str)
	fmt.Println("String is", str1)
	fmt.Printf("Type of String is %T\n", str1)

	num2 := "3.14"
	data1, _ := strconv.ParseFloat(num2, 64)
	fmt.Println("Data is", data1)
	fmt.Printf("Type of Data is %T\n", data1)

	//ParseBool
	//ParseFloat
	//ParseInt
	//ParseUint
}
