package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) //increment the counter by 1
	go f1()   //scheduled to be executed through the scheduler
	f2()
	wg.Wait() //BLOCK until the counter becomes 0

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
