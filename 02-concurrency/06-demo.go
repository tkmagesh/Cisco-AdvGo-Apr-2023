package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var wg sync.WaitGroup

func main() {

	wg := &sync.WaitGroup{}

	rand.Seed(10)
	var count int
	flag.IntVar(&count, "count", 0, "Number of goroutines to execute")
	flag.Parse()

	fmt.Println("main started")
	fmt.Printf("Spinning %d goroutines...\n", count)
	fmt.Println("Hit ENTER to start")
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait()
	fmt.Println("main completed")
	fmt.Println("Hit ENTER to shutdown")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Printf("fn[%d] - started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] - completed\n", id)
}
