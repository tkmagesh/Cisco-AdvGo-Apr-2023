package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	rootCtx := context.Background()
	valCtx := context.WithValue(rootCtx, "app-key", "app-value")
	cancelCtx, cancel := context.WithCancel(valCtx)

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
	fmt.Println("[f1] app-key = ", ctx.Value("app-key"))

	//overriding the values in the context
	f1ValCtx := context.WithValue(ctx, "app-key", "f1-app-value")
	timeoutCtx, cancel := context.WithTimeout(f1ValCtx, 2*time.Second)

	defer cancel()

	wg.Add(1)
	go f11(timeoutCtx, wg)

	wg.Add(1)
	go f12(timeoutCtx, wg)
LOOP:
	for {
		select {
		case done := <-ctx.Done():
			fmt.Println("[f1] Cancel signal received, done :", done)
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("[f1] producing data")
		}
	}
}

func f11(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[f11] app-key = ", ctx.Value("app-key"))
LOOP:
	for {
		select {
		case done := <-ctx.Done():
			fmt.Println("[f11] Cancel signal received, done :", done)
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("[f11] producing data")
		}
	}
}

func f12(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[f12] app-key = ", ctx.Value("app-key"))
LOOP:
	for {
		select {
		case done := <-ctx.Done():
			fmt.Println("[f12] Cancel signal received, done :", done)
			break LOOP
		default:
			time.Sleep(300 * time.Millisecond)
			fmt.Println("[f12] producing data")
		}
	}
}

func f2(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[f2] app-key = ", ctx.Value("app-key"))
LOOP:
	for {
		select {
		case done := <-ctx.Done():
			fmt.Println("[f2] Cancel signal received, done :", done)
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("[f2] producing data")
		}
	}
}
