package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Key, Size int
	Left, Right *Node
}

func (this Node) Init(key int) *Node {
	this.Key = key
	this.Size++
	return &this
}

func (this *Node) GetRandomNode() (*Node, int) {
	if this == nil {
		return nil, 0
	} else if this.Size == 1 {
		return this, 1
	}

	// Get a random number in the range of 1..size
	random := rand.New(rand.NewSource(time.Now().Unix())).Intn(this.Size)+1

	var randomNode func(*Node, int) *Node
	randomNode = func(root *Node, index int) *Node {
		left := 0
		if root.Left != nil {
			left = root.Left.Size
		}

		// Find the pointer position as per the random number.
		if index == left+1 {
			return root
		} else if index <= left {
			return randomNode(root.Left, index)
		} else {
			return randomNode(root.Right, index-left-1)
		}
	}

	return randomNode(this, random), random
}

func (this *Node) Insert(key int) {
	if this == nil {
		return
	}

	if key <= this.Key {
		if this.Left == nil {
			this.Left = Node{}.Init(key)
		} else {
			this.Left.Insert(key)
		}
	} else if this.Right == nil {
		this.Right = Node{}.Init(key)
	} else {
		this.Right.Insert(key)
	}

	this.Size++
}

func main() {
	T1 := Node{}.Init(50)
	for _, k := range []int{30, 20, 40, 70, 60, 80} {
		T1.Insert(k)
	}

	for i := 0; i < 10; i++ {
		node, random := T1.GetRandomNode()
		fmt.Printf("RandomNode(T1) = (%v, %v)\n", random, node.Key)
		time.Sleep(1 * time.Second)
	}
}
