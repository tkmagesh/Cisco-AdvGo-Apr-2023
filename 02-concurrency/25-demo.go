package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genFib()
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib() <-chan int {
	ch := make(chan int)
	stopCh := time.After(5 * time.Second)
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
