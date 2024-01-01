package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		ch <- 10
		ch <- 100
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

}