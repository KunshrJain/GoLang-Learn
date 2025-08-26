package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("What you name BBG :")
	//var name string
	//fmt.Scan(&name)
	//fmt.Println("Hello my BBG :", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello my BBG :", name)

}
