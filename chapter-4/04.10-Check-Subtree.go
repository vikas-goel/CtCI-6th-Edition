package main

import (
	"fmt"
	tree "./pkg4"
)

type Node = tree.Node

func Subtree(T1, T2 *Node) bool {
	// Empty subtree.
	if T2 == nil {
		return true
	} else if T1 == nil {
		return false
	}

	// If the Key of the current node is same as Key of the subtree root,
	// then start matching the entire subtree.
	if T1.Key == T2.Key && matchTrees(T1, T2) {
		return true
	}

	// One of the branches of the main tree is a complete match.
	return Subtree(T1.Left, T2) || Subtree(T1.Right, T2)
}

func matchTrees(T1, T2 *Node) bool {
	if T1 == nil && T2 == nil {
		// Both tree reached to their ends.
		return true
	} else if T1 == nil || T2 == nil {
		// Only one reached end.
		return false
	} else if T1.Key != T2.Key {
		// Mismatch in data.
		return false
	}

	// If both left and right branches match.
	return matchTrees(T1.Left, T2.Left) && matchTrees(T1.Right, T2.Right)
}

func main() {
	T1 := tree.NewBST(50, 30, 20, 40, 70, 60, 80)
	T2 := tree.NewBST(30, 40, 20)

	fmt.Print("T1: ")
	T1.PrintPreorder()
	fmt.Println()
	fmt.Print("T2: ")
	T2.PrintPreorder()
	fmt.Println()

	fmt.Printf("Subtree(T1, T2) = %v\n", Subtree(T1, T2))
}
