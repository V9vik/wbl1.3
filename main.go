package main

import "fmt"

func WriteBit(n int64, i uint, val bool) int64 {
	if i >= 64 {
		return n
	}

	mask := int64(1 << i)

	n &^= mask

	if val == true {
		n |= mask
	}

	return n
}

func main() {
	// ожидаем 5 (0101₂): из 4 (0100₂) включаем 0-й бит
	fmt.Println(WriteBit(4, 0, true))

	// ожидаем 1 (0001₂): из 5 (0101₂) выключаем 2-й бит
	fmt.Println(WriteBit(5, 2, false))
}
