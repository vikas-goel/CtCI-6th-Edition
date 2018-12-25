package main

import (
	"fmt"
	tree "./pkg4"
)

type Node = tree.Node

func SumPaths(root *Node, targetSum int) int {
	if root == nil {
		return 0
	}

	// Track count of running sums.
	sumCount := make(map[int]int)
	sumCount[0] = 1

	var computePaths func(*Node, int) int
	computePaths = func(node *Node, runningSum int) (pathCount int) {
		if node == nil {
			return
		}

		// Add current node value to get running of the current node.
		runningSum += node.Key

		// Find number of paths exist so far.
		diff := runningSum - targetSum
		if num, ok := sumCount[diff]; ok {
			pathCount += num
		}

		// Increment the count of running sum.
		// Compute path count including children.
		// Deceremnt count of running sum before exiting.
		sumCount[runningSum]++
		pathCount += computePaths(node.Left, runningSum)
		pathCount += computePaths(node.Right, runningSum)
		sumCount[runningSum]--

		return
	}

	return computePaths(root, 0)
}

func binTree() (root *Node) {
	root = tree.NewNode(10)
	root.Left = tree.NewNode(5)
	root.Right = tree.NewNode(-3)
	root.Left.Left = tree.NewNode(3)
	root.Left.Right = tree.NewNode(1)
	root.Left.Left.Left = tree.NewNode(3)
	root.Left.Left.Right = tree.NewNode(-2)
	root.Left.Right.Right = tree.NewNode(2)
	root.Right.Right = tree.NewNode(11)
	return
}

func main() {
	root := binTree()
	sum := 6
	root.PrintLevelorder()
	fmt.Println()
	fmt.Printf("Sum Count(%d) = %d\n", sum, SumPaths(root, sum))
}
