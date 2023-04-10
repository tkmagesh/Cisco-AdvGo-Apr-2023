package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go func(no int) {
			fmt.Println("No : ", no)
		}(i) // i is resolved NOT when the function is executed BUT when the function is scheduled for execution
	}
	time.Sleep(2 * time.Second) // main function execution is blocked for 2 seconds
	fmt.Println("main completed")
}
