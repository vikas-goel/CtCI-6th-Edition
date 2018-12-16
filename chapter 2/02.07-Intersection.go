package main

import "fmt"

type List struct {
	Value interface{}
	Next *List
}

func (this *List) Len() int {
	if this == nil {
		return 0
	}

	i := 0
	for ; this.Next != nil; this, i = this.Next, i+1 {}
	return i
}

func (this *List) Print() {
	if this == nil {
		return
	}

	fmt.Printf("[%v", this.Value)
	for this = this.Next; this != nil; this = this.Next {
		fmt.Printf(", %v", this.Value)
	}
	fmt.Print("] ")
}

func NewList(v interface{}) *List {
	l := new(List)
	l.Value = v
	return l
}

// Find intersection of two lists.
func Intersection(n1, n2 *List) *List {
	if n1 == nil || n1.Len() == 0 || n2 == nil || n2.Len() == 0 {
		return nil
	}

	n1Len, n2Len := n1.Len(), n2.Len()
	if n2Len > n1Len {
		n1, n2 = n2, n1
		n1Len, n2Len = n2Len, n1Len
	}

	// Move forward the head pointer of the bigger list by the difference
	// in length of the two lists.
	for i := n1Len - n2Len; i != 0; i-- {
		n1 = n1.Next
	}

	// Walk the two lists parallel and compare each step.
	for ; n2 != nil; n1, n2 = n1.Next, n2.Next {
		if n1 == n2 {
			return n1
		}
	}

	return nil
}

func main() {
	var intersect, temp *List
	l1, l2 := NewList(0), NewList(0)

	temp = l1
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		temp.Next = NewList(i)
		temp = temp.Next
		if i == 6 {
			intersect = temp
		}
	}

	temp = l2
	for _, i := range []int{3, 4, 5} {
		temp.Next = NewList(i)
		temp = temp.Next
	}
	temp.Next = intersect

	l1.Print()
	l2.Print()
	fmt.Println(" = ", Intersection(l1, l2).Value.(int))
}
