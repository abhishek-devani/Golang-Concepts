package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	m sync.Mutex
	v map[string]int
}

func (s *SafeCounter) Inc(key string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.v[key]++
}

func (s *SafeCounter) value(key string) int {
	s.m.Lock()
	defer s.m.Unlock()
	return s.v[key]
}

func main() {

	sc := SafeCounter{
		v : make(map[string]int),
	}

	for i := 1; i < 500; i++ {
		go sc.Inc("someKey")
	}

	time.Sleep(time.Second)
	fmt.Println(sc.value("someKey"))

}