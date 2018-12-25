package main

import (
	"fmt"
	tree "./pkg4"
)

func ValidBST(root, left, right *tree.Node) bool {
	if root == nil {
		return true
	}

	if left != nil && root.Key <= left.Key {
		return false
	}

	if right != nil && root.Key > right.Key {
		return false
	}

	return ValidBST(root.Left, left, root) && ValidBST(root.Right, root, right)

}

func binTree() (root *tree.Node) {
	root = tree.NewNode(3)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(5)
	root.Left.Left = tree.NewNode(1)
	root.Left.Right = tree.NewNode(4)
	return
}

func main() {
	root := tree.NewBST(50, 30, 20, 40, 70, 60, 80)
	fmt.Printf("%11v: { ", "In-order")
	root.PrintInorder()
	fmt.Println("}")
	fmt.Printf("%11v? %v\n", "Valid BST", ValidBST(root, nil, nil))

	root = tree.NewBST(4, 2, 5, 1, 3)
	fmt.Printf("%11v: { ", "In-order")
	root.PrintInorder()
	fmt.Println("}")
	fmt.Printf("%11v? %v\n", "Valid BST", ValidBST(root, nil, nil))

	root = binTree()
	fmt.Printf("%11v: { ", "In-order")
	root.PrintInorder()
	fmt.Println("}")
	fmt.Printf("%11v? %v\n", "Valid BST", ValidBST(root, nil, nil))
}
