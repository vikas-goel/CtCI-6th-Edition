package main

import (
	"container/list"
	"fmt"
)

func RemoveDups(l *list.List) {
	if l == nil || l.Len() < 2 {
		return
	}

	listElems := make(map[interface{}]struct{})
	for elem := l.Front(); elem != nil; {
		thisElem := elem
		elem = elem.Next()
		if _, ok := listElems[thisElem.Value]; ok {
			l.Remove(thisElem)
		} else {
			listElems[thisElem.Value] = struct{}{}
		}
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
	li := list.New()
	for _, i := range []int{0, 1, 99, 1, 10, 102, 103, 10, 98, 25, 102, 0} {
		li.PushBack(i)
	}

	printList(li)
	RemoveDups(li)
	printList(li)

	ls := list.New()
	for _, i := range []string{"a", "ab", "def", "ghi", "ab", "ijk", "a"} {
		ls.PushBack(i)
	}

	printList(ls)
	RemoveDups(ls)
	printList(ls)
}
