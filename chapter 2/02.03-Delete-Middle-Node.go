package main

import "fmt"

type List struct {
	Value interface{}
	Next *List
}

func (this *List) DeleteMiddle() {
	if this == nil || this.Next == nil {
		return
	}

	// Copy value of next node in the current and then delink the next node.
	next := this.Next
	this.Value = next.Value
	this.Next = next.Next
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

func main() {
	var del *List
	l1 := NewList(0)
	temp := l1
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		temp.Next = NewList(i)
		temp = temp.Next
		if i == 6 {
			del = temp
		}
	}

	l1.Print()
	del.DeleteMiddle()
	fmt.Print("-> ")
	l1.Print()
	fmt.Println()
}
