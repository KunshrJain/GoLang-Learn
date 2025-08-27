package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Grade  int
	Rollno int
}

func main() {
	var student Student

	student.Name = "Alice"
	student.Grade = 90
	student.Rollno = 1

	fmt.Printf("Student Name: %s\n", student.Name)
	fmt.Printf("Student Grade: %d\n", student.Grade)
	fmt.Printf("Student Roll No: %d\n", student.Rollno)
}
