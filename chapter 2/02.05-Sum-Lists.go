package main

import (
	"container/list"
	"fmt"
)

// Sum of two numbers in reversed ordered singly linked lists.
func SumReverse(n1, n2 *list.List) *list.List {
	if n1 == nil || n1.Len() == 0 {
		return n2
	} else if n2 == nil || n2.Len() == 0 {
		return n1
	}

	result, carry := list.New(), 0
	v1, v2 := n1.Front(), n2.Front()

	// Add two respective digits of the numbers and track the carry-forward
	// until at least one of them is exhausted.
	for v1 != nil && v2 != nil {
		sum := v1.Value.(int) + v2.Value.(int) + carry
		carry = sum / 10
		sum = sum % 10
		result.PushBack(sum)

		v1 = v1.Next()
		v2 = v2.Next()
	}

	// Set the bigger number as v1 pointer.
	if (v1 == nil) {
		v1 = v2
	}

	// Process remainings of the bigger number while the carry bit is on.
	for ; v1 != nil && carry != 0; v1 = v1.Next() {
		sum := v1.Value.(int) + carry
		carry = sum / 10
		sum = sum % 10
		result.PushBack(sum)
	}

	// Process remaining number based on whether loop exited for
	// the carry (== 0) condition or the element (== nil).
	if carry != 0 {
		// Copy the last carry number.
		result.PushBack(carry)
	} else {
		// Copy the remainings of the number as-is.
		for ; v1 != nil; v1 = v1.Next() {
			result.PushBack(v1.Value.(int))
		}
	}

	return result
}

// Sum of two numbers in forward ordered singly linked lists.
//
// A doubly linked list logic would be same as SumReverse() except that the
// traversal will start from the Back (instead of Front) of the two lists and
// the pointers will move Prev (instead of Next).
func SumForward(n1, n2 *list.List) *list.List {
	if n1 == nil || n1.Len() == 0 {
		return n2
	} else if n2 == nil || n2.Len() == 0 {
		return n1
	}

	// Make n1 greater than or equal to n2 in the lists length.
	len1, len2 := n1.Len(), n2.Len()
	if len2 > len1 {
		n1, n2 = n2, n1
		len1, len2 = len2, len1
	}

	var sum func(e1, e2 *list.Element, len1, len2 int) (*list.Element, int)

	result := list.New()

	// Depth-sum rescursion.
	sum = func(e1, e2 *list.Element, len1, len2 int) (*list.Element, int) {
		// Base case.
		if e1 == nil && e2 == nil {
			return nil, 0
		}

		var re *list.Element
		var rv, c int

		// When the remaining elements in both lists is same, then
		// both lists will be counted. Until then, only the bigger.
		if len1 > len2 {
			re, c = sum(e1.Next(), e2, len1-1, len2)
			rv = e1.Value.(int) + c
		} else {
			re, c = sum(e1.Next(), e2.Next(), len1, len2)
			rv = e1.Value.(int) + e2.Value.(int) + c
		}

		c = rv / 10
		rv %= 10

		if re == nil {
			re = result.PushBack(rv)
		} else {
			re = result.InsertBefore(rv, re)
		}

		return re, c
	}

	re, c := sum(n1.Front(), n2.Front(), len1, len2)
	if c > 0 {
		result.InsertBefore(c, re)
	}

	return result
}

func printList(ll... *list.List) {
	for _, l := range ll {
		if l == nil || l.Front() == nil {
			continue
		}

		fmt.Printf("[%v", l.Front().Value)
		for elem := l.Front().Next(); elem != nil; elem = elem.Next() {
			fmt.Printf(", %v", elem.Value)
		}
		fmt.Print("] ")
	}
	fmt.Println()
}

func main() {
	l1, l2, l3 := list.New(), list.New(), list.New()
	for _, i := range []int{7, 5, 9, 9} {
		l1.PushBack(i)
	}

	for _, i := range []int{9, 9, 2} {
		l2.PushBack(i)
	}

	for _, i := range []int{9, 9, 9, 2} {
		l3.PushBack(i)
	}

	fmt.Println("Reverse sum")
	printList(l1, l2, SumReverse(l1, l2))
	printList(l2, l3, SumReverse(l2, l3))

	fmt.Println("Forward sum")
	printList(l1, l2, SumForward(l1, l2))
	printList(l2, l3, SumForward(l2, l3))
}
