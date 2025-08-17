package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int)
	var worker int
	fmt.Scan(&worker)
	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			for v := range in {
				defer wg.Done()
				fmt.Println(v)
			}
		}()
	}
	wg.Wait()

}
