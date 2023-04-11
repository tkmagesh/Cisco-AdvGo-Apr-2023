package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	defer cancel()
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go f1(cancelCtx, wg)

	wg.Add(1)
	go f2(cancelCtx, wg)

	wg.Wait()
	fmt.Println("Exiting main")
}

func f1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[f1] Cancel signal received")
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("[f1] producing data")
		}
	}
}

func f2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[f2] Cancel signal received")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("[f2] producing data")
		}
	}
}
