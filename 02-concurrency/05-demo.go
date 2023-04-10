package main

import (
	"fmt"
	"time"
)

func main() {
	fn()
	time.Sleep(1 * time.Second)
}

func fn() {
	go f1() //scheduled to be executed through the scheduler
	f2()

}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
