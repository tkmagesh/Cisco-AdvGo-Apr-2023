package main

import "fmt"

/*
func main() {
	ch := make(chan int)
	go func() {
		ch <- 200 // (2) non-blocking coz of (1)
	}()
	data := <-ch // (1) blocked, (3) unblocked with the data
	fmt.Println(data)
}
*/

func main() {
	ch := make(chan int)
	go func() {
		data := <-ch // (2) non-blocking
		fmt.Println(data)
	}()
	ch <- 200 // (1) blocked (3) unblocked
}
