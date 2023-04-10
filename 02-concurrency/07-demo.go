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

var counter = 0

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(counter)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	// counter++
	counter = counter + 1
}
