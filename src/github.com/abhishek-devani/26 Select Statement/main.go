package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan string)

	go functionOne(ch1)
	go functionTwo(ch2)

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}

}

func functionOne(ch1 chan int) {
	ch1 <- 10
}

func functionTwo(ch2 chan string) {
	time.Sleep(time.Second)
	ch2 <- "Function 2"
}