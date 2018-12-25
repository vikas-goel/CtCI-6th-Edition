package main

import (
	"fmt"
	"container/list"
	tree "./pkg4"
)

type Node = tree.Node

func Sequences(root *Node) []*list.List {
	result := make([]*list.List, 0)

	if root == nil {
		return append(result, list.New())
	}

	// Get all sequences for left and right branches.
	left := Sequences(root.Left)
	right := Sequences(root.Right)

	prefix := list.New()
	prefix.PushBack(root.Key)

	// Make new sequences with left & right sequences and current node.
	for _, lseq := range left {
		for _, rseq := range right {
			// For each left and right sequence combination,
			// prepare new sequences joining the current node.
			list := make([]*list.List, 0)
			weave(lseq, rseq, prefix, &list)
			result = append(result, list...)
		}
	}

	return result
}

func weave(list1, list2, prefix *list.List, result *[]*list.List) {
	if list1 == nil || list2 == nil || prefix == nil || result == nil {
		return
	}

	// If either of the lists is empty, then append the remaining list
	// to the prefix.
	if list1.Len() == 0 || list2.Len() == 0 {
		list := list.New()
		list.PushBackList(prefix)
		list.PushBackList(list1)
		list.PushBackList(list2)
		*result = append(*result, list)
		return
	}

	// Remove front item from each list one at a time, append that to the
	// prefix and generate sequence. Revert the changes after creating the
	// sequences.
	for _, list := range []*list.List{list1, list2} {
		front := list.Remove(list.Front()).(int)
		prefix.PushBack(front)

		weave(list1, list2, prefix, result)

		prefix.Remove(prefix.Back())
		list.PushFront(front)
	}
}

func main() {
	//root := tree.NewBST(2, 1, 3)
	root := tree.NewBST(50, 30, 20, 40, 70, 60, 80)
	seq := Sequences(root)

	for _, list := range seq {
		front := list.Front()
		fmt.Printf("{%d", front.Value.(int))
		for front = front.Next(); front != nil; front = front.Next() {
			fmt.Printf(", %d", front.Value.(int))
		}
		fmt.Println("}")
	}
}
