package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)
	for i := 1; i <= 5; i++ {

		go func(chan bool) {
			worker(i, ch)
		}(ch)
		<-ch
	}
}

func worker(i int, check chan bool) {
	
	fmt.Printf("Worker %d starting\n", i)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
	
	check <- true

}