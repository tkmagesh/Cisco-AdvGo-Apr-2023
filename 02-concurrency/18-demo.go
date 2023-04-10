package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		fmt.Println("All the data received")
		break
	}
}

func fn(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 10
		fmt.Println("Sending :", i*10)
	}
	close(ch)
}
