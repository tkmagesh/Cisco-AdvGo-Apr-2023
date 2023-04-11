package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Something went wrong :", err)
		}
	}()
	divisor := 0
	/*
		quotientCh, errCh := divide(100, divisor)
		select {
		case result := <-quotientCh:
			fmt.Printf("Dividing 100 by %d, quotient = %d\n", divisor, result)
		case err := <-errCh:
			fmt.Println(err)
		}
	*/
	timeout := time.After(5 * time.Second)
	quotientCh, _ := divide(100, divisor)
	select {
	case result := <-quotientCh:
		fmt.Printf("Dividing 100 by %d, quotient = %d\n", divisor, result)
	case <-timeout:
		fmt.Println("Timed out...")
		break
	}
}

func divide(x, y int) (<-chan int, <-chan error) {
	ch := make(chan int)
	errCh := make(chan error, 1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				errCh <- err.(error)
				fmt.Println("Error sent to error channel")
				debug.PrintStack()
			}
		}()
		fmt.Println("Performing divide operation.....")
		time.Sleep(3 * time.Second)
		ch <- x / y
	}()
	return ch, errCh

}
