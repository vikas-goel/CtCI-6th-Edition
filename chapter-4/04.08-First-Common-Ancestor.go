package main

import (
	"fmt"
	tree "./pkg4"
)

type Node = tree.Node

// Returns first ancestor of given two keys in the tree.
// If either of the keys are not present in the tree, then ancestor not present.
// Also, if one of the keys is the root, then ancestor not present.
func FirstCommonAncestor(tree *Node, key1, key2 int) *Node {
	if tree == nil {
		return nil
	}

	// fca = first common ancestor.
	// node1, node2 = node representing key1, key2 respectively.
	var fca, node1, node2 *Node

	var findNodes func(*Node) (bool, bool)
	findNodes = func(root *Node) (found1, found2 bool) {
		// Stop recursion if any of the condition is true
		//	root is nil
		//	fca already found
		//	both keys already found
		if root == nil || fca != nil || (node1 != nil && node2 != nil) {
			return
		}

		// Check whether the current node is one of the keys.
		if node1 == nil && root.Key == key1 {
			node1 = root
			found1 = true
		} else if node2 == nil && root.Key == key2 {
			node2 = root
			found2 = true
		}

		// Find the keys in both branches of the current node.
		lfound1, lfound2 := findNodes(root.Left)
		rfound1, rfound2 := findNodes(root.Right)

		// If both keys are found immediate to this node's branches,
		// then this is the first common ancestor.
		if fca == nil && (lfound1 || rfound1) && (lfound2 || rfound2) {
			fca = root
		}

		// Return status of keys due to self node or its branches.
		return (lfound1||rfound1||found1), (lfound2||rfound2||found2)
	}

	findNodes(tree)
	return fca
}

func binTree() (root *tree.Node) {
	root = tree.NewNode(1)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Left.Left = tree.NewNode(4)
	root.Left.Right = tree.NewNode(5)
	root.Right.Left = tree.NewNode(6)
	root.Right.Right = tree.NewNode(7)
	root.Right.Left.Left = tree.NewNode(10)
	root.Right.Left.Right = tree.NewNode(8)
	root.Right.Right.Left = tree.NewNode(11)
	root.Right.Right.Right = tree.NewNode(9)
	return
}

func main() {
	key1, key2 := 11, 5
	tree := binTree()
	ancestor := FirstCommonAncestor(tree, key1, key2)
	if ancestor != nil {
		fmt.Printf("Ancestor of (%d, %d) = %d\n",
			key1, key2, ancestor.Key)
	} else {
		fmt.Printf("Ancestor of (%d, %d) not found.\n",
			key1, key2)
	}
}
