package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduled to be executed through the scheduler
	f2()
	time.Sleep(1 * time.Second) // blocking the execution of the main function and there by giving the opportunity to the scheduler to go and look for other goroutines scheduled and execute them
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
