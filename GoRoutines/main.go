package main

import (
	"fmt"
	"sync"
	"time"
)

func SayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond)
}

func SayBye() {
	fmt.Println("Bye")
	time.Sleep(1000 * time.Millisecond)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d Starting work \n", id)
	fmt.Printf("Worker %d Ending work \n", id)
}

func main() {
	fmt.Println("Learning GO Routines :")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	//wg.Wait()
	fmt.Println("Workers Task completed")
}
