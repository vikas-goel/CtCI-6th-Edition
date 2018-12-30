package main

import "fmt"

type Nums struct {
	multiplier, number uint
}

func Multiply(multiplier, number uint) uint {
	// Make the multiplier smaller of the two number to minimze the number
	// of operations.
	if multiplier > number {
		multiplier, number = number, multiplier
	}

	// Base cases.
	if multiplier == 0 {
		return 0
	} else if multiplier == 1 {
		return number
	}

	// The add factor converts the multiplier from odd to even so that
	// the number can be doubled and the mulitplier can be halved.
	var oddFactor uint
	if multiplier & 1 == 1 {
		oddFactor = number
	}

	// Double the number, reduce the multiplier to half, and add the odd
	// factor. That is, if m is
	//	even: m x n = m/2 + 2*n
	//	 odd: m x n = m/2 + 2*n + n
	return Multiply(multiplier >> 1, number + number) + oddFactor
}

func main() {
	for _, n := range []Nums{{8, 7}, {4, 6}, {15,9}, {0, 6}} {
		fmt.Println("Multiply", n, "=", Multiply(n.multiplier, n.number))
	}
}
