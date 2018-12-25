package main

import (
	"fmt"
	tree "./pkg4"
)

// Given a node, return the successor key.
func Successor(root *tree.Node) (key int, ok bool) {
	// No successor available.
	if root == nil {
		return 0, false
	}

	// If the node has a right child, then the successor is left-most
	// node of the right child.
	if root.Right != nil {
		root = root.Right
		for ; root.Left != nil; root = root.Left {}
		return root.Key, true
	}

	// If the node is root node, then successor not available.
	if root.Parent == nil {
		return 0, false
	}

	// Keep moving up in the tree until we hit origin of the left branch.
	for root.Parent != nil && root.Parent.Right == root {
		root = root.Parent
	}

	if root.Parent == nil {
		// If left branch not found, then there is no successor.
		return 0, false
	} else {
		return root.Parent.Key, true
	}
}

func main() {
	keys := []int{50, 30, 20, 40, 70, 60, 80}
	fmt.Println(keys)
	root := tree.NewBST(keys...)
	for _, k := range keys {
		_, node, _ := root.SearchBST(k)
		successor, ok := Successor(node)

		fmt.Printf("Successor(%v):", k)

		if ok {
			fmt.Println(successor)
		} else {
			fmt.Println("NA")
		}
	}
}
