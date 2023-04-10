package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	go f1()
	f2()
	wg.Wait()

}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(4 * time.Second)
	fmt.Println("f1 completed")
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
