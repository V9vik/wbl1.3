package main

import (
	"fmt"
)

func main() {
	xs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	cha1 := make(chan int)
	cha2 := make(chan int)

	go func() {
		for _, x := range xs {
			cha1 <- x
		}
		defer close(cha1)
	}()

	go func() {
		for x := range cha1 {
			cha2 <- x * 2
		}
		defer close(cha2)
	}()

	for v := range cha2 {
		fmt.Println(v)
	}

}
