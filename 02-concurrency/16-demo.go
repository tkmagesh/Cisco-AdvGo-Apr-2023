package main

import "fmt"

func main() {
	ch := make(chan int)
	go fn(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fn(ch chan int) {
	ch <- 10
	fmt.Println("Sending :", 10)
	ch <- 20
	fmt.Println("Sending :", 20)
	ch <- 30
	fmt.Println("Sending :", 30)
	ch <- 40
	fmt.Println("Sending :", 40)
	ch <- 50
	fmt.Println("Sending :", 50)
}
