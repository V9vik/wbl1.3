package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func byCondition(done chan bool) {
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Способ #1: остановка по условию")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	done <- true
}

func byChannel(stop chan struct{}) {
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Способ #2: остановка через канал")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	close(stop)
}

func byContext(ctx context.Context, cancel context.CancelFunc) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Способ #3: остановка через контекст")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(3 * time.Second)
	cancel()
}

func byGoexit(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; ; i++ {
			if i == 3 {
				fmt.Println("Способ #4: остановка через runtime.Goexit()")
				runtime.Goexit()
			}
			fmt.Println("Работаю...")
			time.Sleep(time.Second)
		}
	}()
}

func main() {
	var wg sync.WaitGroup

	cond := make(chan bool)
	byCondition(cond)

	ch := make(chan struct{})
	byChannel(ch)

	ctx, cancel := context.WithCancel(context.Background())
	byContext(ctx, cancel)

	byGoexit(&wg)

	time.Sleep(5 * time.Second)
	wg.Wait()
	fmt.Println("Все способы остановки продемонстрированы")
}
