// In an array of integers, a "peak" is an element which is greater than or
// equal to the adjacent integers and a "valley" is an element which is less
// than or equal to the adjacent integers. For example, in the array
// {5, 8, 6, 2, 3, 4, 6}, {8, 6} are peaks and {5, 2} are valleys.
// Given an array of integers, sort the array into an alternating sequence of
// peaks and valleys.

package main

import "fmt"

func rearrange(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		if (i % 2 == 0 && array[i-1] < array[i]) ||
			(i % 2 != 0 && array[i-1] > array[i]) {
				array[i-1], array[i] = array[i], array[i-1]
		}
	}
}

func main() {
	a := []int{2, 5, 8, 3, 10, 15, 12, 11, 9, 4, 5, 6, 7}
	fmt.Print(a)
	rearrange(a)
	fmt.Print(a)
}
