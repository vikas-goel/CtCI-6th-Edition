package main

import "fmt"

// Assumes N > M and (i,j) > bits(M).
func Insert(N, M int, i, j uint) int {
	// Create mask to clear i through j bits of N.
	// The mask is comprised of all 1s [j+1:] and [i-1:].
	nMask := (^0 << (j + 1)) | ((1 << i) - 1)

	// Shift M by i positions.
	mMove := M << i

	// Merge the masked N and shifted M.
	return (N & nMask) | mMove
}

func main() {
	var N, M int = 1040, 19
	var i, j uint = 2, 6
	fmt.Printf("Insert(M=%b in N=%b) @(i=%d, j=%d) = %b\n",
		M, N, i, j, Insert(N, M, i, j))
}
