package main

/*
	To Detect the data race
		go run -race 07-demo.go
		go build -race 07-demo.go // DO NOT create a production build with race detector

*/

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count int64
}

func (c *Counter) increment() {
	atomic.AddInt64(&c.count, 1)
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
