package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)
	fmt.Println("Start")

	var wg sync.WaitGroup
	
	wg.Add(1)
	go func () {
		defer wg.Done()
		multiply(ch)
	}()
	ch <- 10

	defer close(ch)

	wg.Wait()

	// go mul(ch)
	fmt.Println("End")

}

func multiply(ch chan int) {
	fmt.Println(10 * <-ch)
}
