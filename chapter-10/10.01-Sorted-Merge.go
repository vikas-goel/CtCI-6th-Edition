// You are given two sorted arrays, A and B, where A has a large enough buffer
// at the end to hold B. Write a method to merge B into A in sorted order.

package main

import "fmt"

func SortedMerge(A, B []int, lengthA int) {
	lengthB := len(B)

	// Starting from rear end, copy the elements of A and B in correct
	// order.
	for i := lengthA+lengthB-1; lengthA > 0 && lengthB > 0; i-- {
		if A[lengthA-1] > B[lengthB-1] {
			A[i] = A[lengthA-1]
			lengthA--
		} else {
			A[i] = B[lengthB-1]
			lengthB--
		}
	}

	// Copy the remaining elements of B.
	for i := 0; i < lengthB; i++ {
		A[i] = B[i]
	}
}

func main() {
	array1 := []int{8, 12, 14, 17, 18}
	array2 := []int{4, 7, 9, 10, 15}
	array3 := []int{5, 6, 11, 13, 16, 19}

	array := make([]int, len(array1)+len(array2)+len(array3))
	copy(array, array1)

	SortedMerge(array, array2, len(array1))
	fmt.Println(array)

	SortedMerge(array, array3, len(array1)+len(array2))
	fmt.Println(array)
}
