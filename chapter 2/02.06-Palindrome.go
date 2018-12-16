package main

import (
	"container/list"
	"fmt"
)

// Check whether a singly linked list is palindrome or not.
func Palindrome(l *list.List) bool {
	if l == nil || l.Len() == 0 {
		return false
	} else if l.Len() == 1 {
		return true
	}

	allChecked := false
	var compare func(start **list.Element, this *list.Element) bool

	compare = func(start **list.Element, this *list.Element) bool {
		if this == nil {
			return true
		}

		if compare(start, this.Next()) {
			// If this is the mid point for odd length linked list.
			if (*start) == this {
				allChecked = true
			}

			if allChecked {
				// If we reached here, then all comparison done
				// and the linked list is palindrome.
				return true
			}

			// If this is the mid point for even length linked list.
			if (*start).Next() == this {
				allChecked = true
			}

			if (*start).Value == this.Value {
				*start = (*start).Next()
				return true
			}

			return false
		}

		return false
	}

	start := l.Front()
	return compare(&start, l.Front().Next())
}

func printList(l *list.List, status bool) {
	if l == nil || l.Front() == nil {
		return
	}

	fmt.Printf("[%v", l.Front().Value)
	for elem := l.Front().Next(); elem != nil; elem = elem.Next() {
		fmt.Printf(", %v", elem.Value)
	}
	fmt.Println("] =", status)
}

func main() {
	l1, l2, l3, l4 := list.New(), list.New(), list.New(), list.New()
	for _, i := range []int{1, 2, 3, 4, 5, 4, 3, 2, 1} {
		l1.PushBack(i)
	}

	for _, i := range []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1} {
		l2.PushBack(i)
	}

	for _, i := range []string{"a", "b", "c", "d", "e", "d", "c", "b", "a"}{
		l3.PushBack(i)
	}

	for _, i := range []int{1, 2, 3, 4, 4, 2, 1} {
		l4.PushBack(i)
	}

	printList(l1, Palindrome(l1))
	printList(l2, Palindrome(l2))
	printList(l3, Palindrome(l3))
	printList(l4, Palindrome(l4))
}
