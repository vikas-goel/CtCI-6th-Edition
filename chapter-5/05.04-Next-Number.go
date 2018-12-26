package main

import "fmt"

func GetNext(number int) (larger int, smaller int) {
	return getNextLarger(number), getNextSmaller(number)
}

func getNextLarger(number int) int {
	var c0, c1 uint
	larger := number

	// Find number of trailing 0s.
	for c0 = 0; larger & 1 == 0 && larger != 0; larger >>= 1 {
		c0++
	}

	// Find number of 1s on the left of trailing 0s.
	for c1 = 0; larger & 1 == 1; larger >>= 1 {
		c1++
	}

	// Cannot have a larger number if there is no 1s or leading 0s.
	if c1 == 0 || c0 + c1 == 31 {
		return -1
	}

	// Flip the rightmost non-trailing zero.
	larger = number | (1 << (c0 + c1))

	// Clear all the (c1) bits to the right of the flipped bit.
	larger &= ^((1 << (c0 + c1)) - 1)

	// Add trailing (c1-1) 1s.
	larger |= ((1 << (c1-1)) - 1)

	return larger
}

func getNextSmaller(number int) int {
	var c0, c1 uint
	smaller := number

	// Find number of trailing 1s.
	for c1 = 0; smaller & 1 == 1; smaller >>= 1 {
		c1++
	}

	// Find number of 0s on the left of trailing 1s.
	for c0 = 0; smaller & 1 == 0 && smaller != 0; smaller >>= 1 {
		c0++
	}

	// Cannot have a smaller number if there is no 0s or leading 1s.
	if c0 == 0 || c0 + c1 == 31 {
		return -1
	}

	// Clear all bits starting from right most non-trailing 1.
	smaller = number & (^0 << (c0 + c1 + 1))

	// Sequence of c1+1 1s.
	mask := (1 << (c1 + 1)) - 1

	smaller |= (mask << (c0 - 1))

	return smaller
}

func main() {
	num := 13948
	l, s := GetNext(num)
	fmt.Printf("N(%d,%b): L(%d,%b), S(%d,%b)\n", num, num, l, l, s, s)
}
