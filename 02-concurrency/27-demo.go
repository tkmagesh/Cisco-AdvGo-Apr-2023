/* buffered channels */
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 100
	fmt.Println("Sent data to channel")
	fmt.Println("len(ch) = ", len(ch))
	ch <- 200
	fmt.Println("Sent data to channel")
	fmt.Println("len(ch) = ", len(ch))
	data := <-ch
	fmt.Println(data)
	fmt.Println("len(ch) = ", len(ch))
	data = <-ch
	fmt.Println(data)
	fmt.Println("len(ch) = ", len(ch))
}
