package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genFib(stopCh)
	fmt.Println("Hit ENTER to stop")
	go func() {
		fmt.Scanln()
		// stopCh <- struct{}{}
		close(stopCh)
	}()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	FIB_LOOP:
		for {
			select {
			case <-stopCh:
				break FIB_LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- x
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

/*
func timeout(d time.Duration) <-chan time.Time {
	stopCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		stopCh <- time.Now()
	}()
	return stopCh
}
*/
