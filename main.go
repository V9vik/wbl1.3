package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	in := make(chan int)
	var worker int
	fmt.Scan(&worker)
	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			for {
				defer wg.Done()
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
		}()
	}
	wg.Wait()

}
