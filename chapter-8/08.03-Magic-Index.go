package main

import "fmt"

func MagicIndexDistinct(array []int) int {
	length := len(array)
	if length == 0 {
		return -1
	}

	// Use binary search to find the index.
	for start, end := 0, length-1; start <= end; {
		mid := (start+end)/2
		if array[mid] == mid {
			// Spot on.
			return mid
		} else if array[mid] > mid {
			// The index would be on the left half as the mid
			// value is greater than the index.
			end = mid-1
		} else {
			// The index would be on the right half as the mid
			// value is smaller than the index.
			start = mid+1
		}
	}

	// Does not exist.
	return -1
}

func MagicIndexWithDups(array []int) int {
	length := len(array)
	if length == 0 {
		return -1
	}

	// Keep dividing the array in two halves and use depth first search.
	var findIndex func(int, int) int
	findIndex = func(start, end int) int {
		// Invalid range.
		if start > end {
			return -1
		}

		// Spot on.
		mid := (start+end)/2
		if array[mid] == mid {
			return mid
		} else if start == end {
			return -1
		}

		left := mid-1
		if left > array[mid] {
			// Optimize the left range.
			left = array[mid]
		}

		index := findIndex(start, left)
		if index >= 0 {
			// Found the index in the left half.
			return index
		}

		right := mid+1
		if right < array[mid] {
			// Optimize the right range.
			right = array[mid]
		}

		// Search the index in the optimized right half.
		return findIndex(right, end)
	}

	return findIndex(0, length-1)
}

func main() {
	distinct := []int{-40, -20, -1, 1, 2, 3, 5, 7, 9, 12, 13}
	fmt.Println("MagicIndex", distinct, "=", MagicIndexDistinct(distinct))

	withdups := []int{-10, -5, 2, 2, 2, 3, 4, 7, 9, 12, 13}
	fmt.Println("MagicIndex", withdups, "=", MagicIndexWithDups(withdups))
}
