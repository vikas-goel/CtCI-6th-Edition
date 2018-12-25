package main

import (
	"container/list"
	"fmt"
	tree "./pkg4"
)

func ListOfDepths(root *tree.Node) []*list.List {
	if root == nil {
		return nil
	}

	currentList := list.New()
	currentList.PushBack(root)

	listOfDepths := make([]*list.List, 0, 1)
	listOfDepths = append(listOfDepths, currentList)

	for i := 0; i < len(listOfDepths); i++ {
		currentList = list.New()

		for elem := listOfDepths[i].Front(); elem != nil; elem = elem.Next() {
			node := elem.Value.(*tree.Node)
			if node.Left != nil {
				currentList.PushBack(node.Left)
			}

			if node.Right != nil {
				currentList.PushBack(node.Right)
			}
		}

		if currentList.Len() > 0 {
			listOfDepths = append(listOfDepths, currentList)
		}
	}

	return listOfDepths
}

func main() {
	root := tree.NewBST(50, 30, 20, 40, 70, 60, 80)
	list := ListOfDepths(root)

	fmt.Printf("Level-order tree: ")
	root.PrintLevelorder()
	fmt.Println()

	for i := 0; i < len(list); i++ {
		fmt.Printf("Depth %d:", i)
		for n := list[i].Front(); n != nil; n = n.Next() {
			fmt.Printf(" %v", n.Value.(*tree.Node).Key)
		}
		fmt.Println()
	}
}
