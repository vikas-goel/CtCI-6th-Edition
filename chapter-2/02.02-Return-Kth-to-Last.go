package main

import (
	"container/list"
	"fmt"
)

// Return Kth element from the end of the list.
// The function assumes that the length of the list is unknown.
func KthElement(l *list.List, k int) *list.Element {
	if l == nil {
		return nil
	}

	i, elem := 0, l.Front()
	for i = 0; i <= k && elem != nil; i++ {
		elem = elem.Next()
	}

	// Less than or equal to k elements present in the list.
	if i <= k {
		return nil
	}

	kth := l.Front()
	for ; elem != nil; elem = elem.Next() {
		kth = kth.Next()
	}

	return kth
}

func main() {
	l := list.New()
	for i := 1; i <= 100; i++ {
		l.PushBack(i)
	}

	for _, i := range []int{0, 1, 99, 100} {
		elem := KthElement(l, i)
		if elem != nil {
			fmt.Printf("%v-th element = %v\n", i, elem.Value)
		} else {
			fmt.Printf("%v-th element = nil\n", i)
		}
	}
}
