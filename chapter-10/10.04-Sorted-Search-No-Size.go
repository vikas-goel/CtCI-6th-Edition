// You are given an array-like data structure Listy which lacks a size method.
// It does, however, have an elementAt (i) method that returns the element at
// index i in 0(1) time. If i is beyond the bounds of the data structure, it
// returns -1. (For this reason, the data structure only supports positive
// integers.) Given a Listy which contains sorted, positive integers, find the
// index at which an element x occurs. If x occurs multiple times, you may
// return any index.

package main

import "fmt"

type Listy map[int]int
func (this Listy) ElementAt(index int) int {
	if element, ok := this[index]; ok {
		return element
	}

	return -1
}

func binarySearch(key int, list Listy, start, end int) int {
	if end < start {
		return -1
	} else if start == end {
		if list.ElementAt(start) == key {
			return start
		}
		return -1
	}

	mid := start + (end - start)/2
	if list.ElementAt(mid) == key {
		return mid
	} else if key < list.ElementAt(mid) || list.ElementAt(mid) == -1 {
		return binarySearch(key, list, start, mid-1)
	} else  {
		return binarySearch(key, list, mid+1, end)
	}
}

func findIndexOf(key int, list Listy) int {
	end := 1
	for ; list.ElementAt(end-1) != -1; end *= 2 {}

	return binarySearch(key, list, 0, end-1)
}

func main() {
	list := make(map[int]int)
	for i, v := range []int{1, 1, 1, 3, 5, 5, 6, 7, 9, 10, 10, 10, 15} {
		list[i] = v
	}

	fmt.Println(list)

	fmt.Printf("search[ ")
	for _, v := range list {
		fmt.Printf("%d:%d ", findIndexOf(v, list), v)
	}
	fmt.Printf("]")
}
