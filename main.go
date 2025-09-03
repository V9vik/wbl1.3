package main

import "fmt"

func main() {
	mapForCheck := make(map[int]int)
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	resultSlice := make([]int, 0, len(A))

	for _, i := range A {
		mapForCheck[i]++
	}

	for _, i := range B {
		mapForCheck[i]++
	}
	for num, c := range mapForCheck {
		if c >= 2 {
			resultSlice = append(resultSlice, num)
		}
	}
	fmt.Print(resultSlice)
}
