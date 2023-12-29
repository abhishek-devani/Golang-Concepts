package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Start goroutine")

	var wg sync.WaitGroup

	wg.Add(1)

	go printHello(&wg)

	wg.Wait()

	fmt.Println("End")

}

func printHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Print hello")
}