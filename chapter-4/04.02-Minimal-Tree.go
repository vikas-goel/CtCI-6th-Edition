package main

import (
	"fmt"
	tree "./pkg4"
)

func MinimalBST(sortedArray []int, start, end int) *tree.Node {
	// Invalid range.
	if start > end {
		return nil
	}

	// Create the middle element as the root node for this recursion.
	mid := (start+end)/2
	root := new(tree.Node)
	root.Key = sortedArray[mid]

	// Base case if there is no other element left in this range.
	if start == end {
		return root
	}

	// Build left and right branches of this root.
	root.Left = MinimalBST(sortedArray, start, mid-1)
	root.Right = MinimalBST(sortedArray, mid+1, end)

	return root
}

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree := MinimalBST(array, 0, len(array)-1)
	fmt.Print("Array:      ")
	fmt.Println(array)
	fmt.Print("In-order:   [")
	tree.PrintInorder()
	fmt.Println("]")
}
