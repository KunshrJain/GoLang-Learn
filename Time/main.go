package main

import (
	"fmt"
	"time"
)

func main() {
	// "dd--mm--yyyy"
	// "yyyy/mm/dd"
	// "yyyy-mm-dd hh:mm:ss"

	curtime := time.Now()
	fmt.Println("Current Time is", curtime)

	formatted := curtime.Format("02-01-2006 15:04:05 , Monday")
	fmt.Println("Formatted Time is", formatted)
	formatted = curtime.Format("02-01-2006 15:04 PM , Monday")
	fmt.Println("Formatted Time is", formatted)

	curtime = curtime.Add(time.Hour * 24)
	formatted = curtime.Format("02-01-2006 15:04:05 , Tuesday")
	fmt.Println("Formatted Time is", formatted)

}
