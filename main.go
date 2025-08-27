package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const N = 5 // секунд работы программы

func producer(ctx context.Context, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case out <- i:
			i++
		}
	}
}

func consumer(ctx context.Context, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			fmt.Println(v)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), N*time.Second)
	defer cancel()

	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(ctx, ch, &wg)
	go consumer(ctx, ch, &wg)

	wg.Wait()
}
