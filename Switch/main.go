package main

import "fmt"

func main() {
	day := 3
	switch day {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Pta nhi yaar")
	}

	month := "Wednesday"
	switch month {
	case "Monday":
		fmt.Println("Monday")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Tuesday/Wednesday/Thursday")
	case "Friday", "Saturday":
		fmt.Println("Friday/Saturday")
	default:
		fmt.Println("Pta nhi yaar")
	}

	temp := 25
	switch {
	case temp < 20:
		fmt.Println("below 20")
	case temp == 25 && temp < 30:
		fmt.Println("25 or below 30")
	default:
		fmt.Println("Pta nhi yaar")
	}
}
