package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)    //increment the counter by 1
		go f1(i, wg) //scheduled to be executed through the scheduler
	}
	f2()
	wg.Wait() //BLOCK until the counter becomes 0
}

func f1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("f1[%d] started\n", id)
	time.Sleep(4 * time.Second)
	fmt.Printf("f1[%d] completed\n", id)

}

func f2() {
	fmt.Println("f2 invoked")
}
