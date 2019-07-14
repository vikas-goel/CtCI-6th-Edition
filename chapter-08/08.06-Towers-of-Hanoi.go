// In the classic problem of the Towers of Hanoi, you have 3 towers and N disks
// of different sizes which can slide onto any tower. The puzzle starts with
// disks sorted in ascending order of size from top to bottom (Le., each disk
// sits on top of an even larger one). You have the following constraints:
//   (1) Only one disk can be moved at a time.
//   (2) A disk is slid off the top of one tower onto another tower.
//   (3) A disk cannot be placed on top of a smaller disk.
// Write a program to move the disks from the first tower to the last using
// Stacks.

package main

import "fmt"

type Stack struct {
	capacity, size int
	nums []int
}

func (this *Stack) Push(num int) bool {
	if this.capacity == this.size {
		return false
	}

	this.nums[this.size] = num
	this.size++

	return true
}

func (this *Stack) Pop() (int, bool) {
	if this.size == 0 {
		return 0, false
	}

	this.size--
	element := this.nums[this.size]
	this.nums[this.size] = 0

	return element, true
}

type Tower struct {
	id int
	st Stack
}

func (this *Tower) MoveDisks(n int, dest, buff *Tower) {
	if n <= 0 {
		return
	}

	this.MoveDisks(n-1, buff, dest)

	this.MoveTopDisk(dest)
	PrintTowers(this, dest, buff)

	buff.MoveDisks(n-1, dest, this)
}

func (this *Tower) MoveTopDisk(dest *Tower) {
	topDisk, _ := this.st.Pop()
	dest.st.Push(topDisk)
}

func (this *Tower) Init(id, disks int) {
	this.id = id
	this.st.nums = make([]int, disks)
	this.st.capacity = disks
	this.st.size = 0
}

func (this *Tower) StackDisks() {
	for i := this.st.capacity; i > 0; i-- {
		this.st.Push(i)
	}
}

func PrintTowers(towers... *Tower) {
	for start, end := 0, len(towers); start != end; start++ {
		for i := 0; i < end; i++ {
			if towers[i].id == start {
				fmt.Print(towers[i].st, " ")
				break
			}
		}
	}
	fmt.Println()
}

func main() {
	numTowers, numDisks := 3, 5

	towers := make([]Tower, numTowers)
	for i := 0; i < numTowers; i++ {
		towers[i].Init(i, numDisks)
	}

	towers[0].StackDisks()
	PrintTowers(&towers[0], &towers[1], &towers[2])
	towers[0].MoveDisks(numDisks, &towers[1], &towers[2])
}
