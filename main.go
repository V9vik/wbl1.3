package main

import (
	"fmt"
	"math"
	"sort"
)

func bucket(v float64) int {
	return int(math.Trunc(v/10.0)) * 10
}

func main() {
	vals := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, v := range vals {
		k := bucket(v)
		groups[k] = append(groups[k], v)
	}

	keys := make([]int, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, groups[k])
	}
}
