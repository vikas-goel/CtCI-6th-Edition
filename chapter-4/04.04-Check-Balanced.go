package main

import (
	"fmt"
	tree "./pkg4"
)

func IsBalanced(root *tree.Node) (ok bool, height int) {
	if root == nil {
		return true, 0
	}

	lok, lheight := IsBalanced(root.Left)
	rok, rheight := IsBalanced(root.Right)

	if !lok {
		return lok, lheight
	} else if !rok {
		return rok, rheight
	}

	if lheight > rheight {
		lheight, rheight = rheight, lheight
	}

	height = 1+rheight
	ok = rheight <= lheight+1

	return
}

func binTree1() (root *tree.Node) {
	root = tree.NewNode(10)
	root.Left = tree.NewNode(12)
	root.Right = tree.NewNode(15)
	root.Left.Left = tree.NewNode(25)
	root.Left.Right = tree.NewNode(30)
	root.Right.Left = tree.NewNode(36)
	return
}

func binTree2() (root *tree.Node) {
	root = tree.NewNode(1)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Left.Left = tree.NewNode(4)
	root.Left.Right = tree.NewNode(5)
	root.Left.Right.Left = tree.NewNode(6)
	root.Left.Right.Right = tree.NewNode(7)
	root.Left.Left.Right = tree.NewNode(8)
	root.Left.Left.Right.Left = tree.NewNode(9)
	return
}

func binTree3() (root *tree.Node) {
	root = tree.NewNode(1)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Left.Left = tree.NewNode(4)
	root.Left.Right = tree.NewNode(5)
	root.Left.Left.Left = tree.NewNode(7)
	return
}

func binTree4() (root *tree.Node) {
	root = tree.NewNode(1)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Left.Left = tree.NewNode(4)
	root.Left.Right = tree.NewNode(5)
	root.Right.Right = tree.NewNode(8)
	root.Right.Right.Left = tree.NewNode(6)
	root.Right.Right.Right = tree.NewNode(7)
	return
}

func binTree5() (root *tree.Node) {
	root = tree.NewNode(1)
	root.Left = tree.NewNode(2)
	root.Right = tree.NewNode(3)
	root.Left.Left = tree.NewNode(4)
	root.Left.Right = tree.NewNode(5)
	root.Right.Left = tree.NewNode(6)
	root.Right.Right = tree.NewNode(7)
	root.Right.Left.Right = tree.NewNode(8)
	root.Right.Right.Right = tree.NewNode(9)
	return
}

func main() {
	roots := []*tree.Node{binTree1(), binTree2(), binTree3(), binTree4(), binTree5(), tree.NewBST(50, 30, 20, 40, 70, 60, 80)}

	for _, root := range roots {
		ok, _ := IsBalanced(root)
		root.PrintInorder()
		fmt.Printf(": %v\n", ok)
	}
}
