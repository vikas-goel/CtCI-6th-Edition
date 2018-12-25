package main

import (
	"container/list"
	"fmt"
)

func Partition(l *list.List, x int) {
	if l == nil || l.Len() < 2 {
		return
	}

	for i, elem := 0, l.Front(); i < l.Len(); i++ {
		thisElem := elem
		elem = elem.Next()

		thisValue := thisElem.Value.(int)
		if thisValue < x {
			continue
		}

		l.MoveToBack(thisElem)
	}
}

func printList(l *list.List) {
	if l == nil || l.Front() == nil {
		return
	}

	fmt.Printf("[%v", l.Front().Value)
	for elem := l.Front().Next(); elem != nil; elem = elem.Next() {
		fmt.Printf(", %v", elem.Value)
	}
	fmt.Println("]")
}

func main() {
	l := list.New()
	for _, i := range []int{3, 5, 8, 5, 10, 2, 1} {
		l.PushBack(i)
	}

	printList(l)
	Partition(l, 5)
	printList(l)
}
