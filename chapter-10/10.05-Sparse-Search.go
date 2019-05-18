// Given a sorted array of strings that is interspersed with empty strings,
// write a method to find the location of a given string.

package main

import (
	"fmt"
	"strings"
)

func binarySearch(str string, sortedArray []string, start, end int) int {
	if end < start {
		return -1
	}

	mid := start + (end - start)/2

	// If mid element is emtpty string, then find the nearest non-empty
	// string as the mid element.
	if strings.Compare(sortedArray[mid], "") == 0 {
		i := 1
		for {
			if mid-i < start && mid+1 > end {
				return -1
			}

			if mid-i > start &&
				strings.Compare(sortedArray[mid-i], "") != 0 {
				mid -= i
				break
			} else if mid+i < end &&
				strings.Compare(sortedArray[mid+i], "") != 0 {
				mid += i
				break
			}
		}
	}

	if diff := strings.Compare(str, sortedArray[mid]); diff == 0 {
		return mid
	} else if diff < 0 {
		return binarySearch(str, sortedArray, start, mid-1)
	} else  {
		return binarySearch(str, sortedArray, mid+1, end)
	}
}

func main() {
	array := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	fmt.Println(array)

	element := "ball"
	fmt.Printf("Element=%v @ Index=%d\n", element,
		binarySearch(element, array, 0, len(array)-1))
}
