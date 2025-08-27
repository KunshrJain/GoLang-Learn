package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5} //This is a Slicee not an array *(somehthing similar but not the same)
	fmt.Println("Numbers are", numbers)
	fmt.Println("Numbers are", len(numbers))
	fmt.Printf("Number has data type %T\n", numbers)

	numbers = append(numbers, 3, 10)
	fmt.Println("Numbers are", numbers)

	numbers = make([]int, 3, 5) // len 3 capacity 5
	fmt.Println("Slice:", numbers)
	fmt.Println("Slice:", len(numbers))
	fmt.Println("Slice:", cap(numbers))

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 56)

	fmt.Println("Slice:", numbers)
	fmt.Println("Slice:", len(numbers))
	fmt.Println("Slice:", cap(numbers))

}
