package main

import "fmt"

func Conversion(A, B int) (count int) {
	// Find the bits difference using XOR and then count 1s.
	for diff := A^B; diff != 0; diff >>= 1 {
		if diff & 1 == 1 {
			count++
		}
	}

	return count
}

func main() {
	A, B := 29, 15
	fmt.Printf("Flips(%b, %b) = %d\n", A, B, Conversion(A, B))
}
