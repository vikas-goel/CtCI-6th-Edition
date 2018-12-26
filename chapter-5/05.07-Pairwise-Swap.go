package main

import "fmt"

func SwapEvenOddBits(number uint) uint {
	// 1 left shift odd bits ORed with 1 right shift even bits.
	return ((number & 0x55555555) << 1) | ((number & 0xaaaaaaaa) >> 1)
}

func main() {
	for _, n := range []uint{1, 2, 3, 4, 5, 10, 100, 150} {
		fmt.Printf("{%b, %b}\n", n, SwapEvenOddBits(n))
	}
}
