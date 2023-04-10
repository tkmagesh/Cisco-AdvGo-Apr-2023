package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := make(chan int)
	go fn(ch)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("All the data received")
}

//producer
func fn(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 10
		fmt.Println("Sending :", i*10)
	}
	close(ch)
}
