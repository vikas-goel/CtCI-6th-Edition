// Given a sorted array of n integers that has been rotated an unknown number
// of times, write code to find an element in the array. You may assume that
// the array was originally sorted in increasing order.

package main

import "fmt"

func findPivot(array []int, start, end int) int {
	if end < start {
		return -1
	}

	if start == end {
		return start
	}

	mid := start + (end - start)/2
	if array[mid] > array[mid+1] {
		return mid
	} else if array[mid] < array[mid-1] {
		return mid-1
	}

	if array[start] > array[mid] {
		return findPivot(array, start, mid-1)
	} else {
		return findPivot(array, mid+1, end)
	}
}

func binarySearch(key int, sortedArray []int, start, end int) int {
	if end < start {
		return -1
	} else if start == end {
		if sortedArray[start] == key {
			return start
		}
		return -1
	}

	mid := start + (end - start)/2
	if sortedArray[mid] == key {
		return mid
	} else if key < sortedArray[mid] {
		return binarySearch(key, sortedArray, start, mid-1)
	} else  {
		return binarySearch(key, sortedArray, mid+1, end)
	}
}

func findIndexOf(key, pivot int, array []int) (index int) {
	if pivot == -1 {
		index = binarySearch(key, array, 0, len(array)-1)
	} else if key == array[pivot] {
		index = pivot
	} else if key < array[0] {
		index = binarySearch(key, array, pivot+1, len(array)-1)
	} else {
		index = binarySearch(key, array, 0, pivot)
	}

	return
}

func main() {
	array := []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
	fmt.Println(array)

	pivot := findPivot(array, 0, len(array)-1)

	for _, key := range array {
		fmt.Printf("Key=%d, Index=%d\n", key, findIndexOf(key, pivot, array))
	}
}
