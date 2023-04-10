package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := fn()
	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("All the data received")
}

//producer
func fn() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
			fmt.Println("Sending :", i*10)
		}
		close(ch)
	}()
	return ch
}
