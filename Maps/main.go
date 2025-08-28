package main

import (
	"fmt"
)

func main() {
	StudGrades := make(map[string]int)
	StudGrades["Alice"] = 90
	StudGrades["Bob"] = 85
	StudGrades["Charlie"] = 92

	for name, grade := range StudGrades {
		fmt.Printf("Student: %s, Grade: %d\n", name, grade)
	}

	StudGrades["bob"] = 100

	for name, grade := range StudGrades {
		fmt.Printf("Student: %s, Grade: %d\n", name, grade)
	}

	delete(StudGrades, "Alice")

	for name, grade := range StudGrades {
		fmt.Printf("Student: %s, Grade: %d\n", name, grade)
	}

	grades, exists := StudGrades["Bob"]
	if exists {
		fmt.Printf("Bob's grade is: %d\n", grades)
	} else {
		fmt.Println("Bob's grade not found.")
	}
	grades, exists = StudGrades["Alice"]
	if exists {
		fmt.Printf("Alice's grade is: %d\n", grades)
	} else {
		fmt.Println("Alices's grade not found.")
	}
}
