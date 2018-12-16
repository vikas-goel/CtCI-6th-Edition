package main

import "fmt"

type List struct {
	Value interface{}
	Next *List
}

func (this *List) LoopDetection() *List {
	if this == nil || this.Next == nil {
		return nil
	}

	// The 'fast' pointer moves at the double speed of the 'slow' one.
	slow, fast := this, this
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// No loop.
	if slow != fast {
		return nil
	}

	// At this point both 'this' and 'slow' pointers are equidistant from
	// the beginning of the loop.
	for this != slow {
		this = this.Next
		slow = slow.Next
	}

	return this
}

func NewList(v interface{}) *List {
	l := new(List)
	l.Value = v
	return l
}

func main() {
	var loop *List
	l1 := NewList(0)
	temp := l1
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		temp.Next = NewList(i)
		temp = temp.Next
		if i == 5 {
			loop = temp
		}
	}

	// Create loop.
	fmt.Println("Creating loop at node with vlaue =", loop.Value)
	temp.Next = loop

	fmt.Println("Loop detected at node with value =",
		l1.LoopDetection().Value)
}
