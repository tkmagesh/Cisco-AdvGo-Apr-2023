package main

/*
	To Detect the data race
		go run -race 07-demo.go
		go build -race 07-demo.go // DO NOT create a production build with race detector

*/

import (
	"fmt"
	"sync"
)

type Counter struct {
	mutex sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.mutex.Lock()
	{
		c.count++
	}
	c.mutex.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			counter.increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}
