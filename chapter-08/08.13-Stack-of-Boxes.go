// You have a stack of n boxes, with widths Wi, heights hi, and depths di. The
// boxes cannot be rotated and can only be stacked on top of one another if
// each box in the stack is strictly larger than the box above it in width,
// height, and depth. Implement a method to compute the height of the tallest
// possible stack. The height of a stack is the sum of the heights of each box.

package main

import (
	"fmt"
	"sort"
)

type Box struct {
	height, width, depth int
}

type Boxes []Box

func (this Boxes) Len() int {
	return len(this)
}

func (this Boxes) Less(i, j int) bool {
	return this[i].height > this[j].height
}

func (this Boxes) Swap(i, j int) {
	this[i].height, this[j].height = this[j].height, this[i].height
	this[i].width, this[j].width = this[j].width, this[i].width
	this[i].depth, this[j].depth = this[j].depth, this[i].depth
}

func (this Boxes) canStack(up, on int) bool {
	return this[up].height < this[on].height &&
		this[up].width < this[on].width &&
		this[up].depth < this[on].depth
}

func (this Boxes) Stack() int {
	sort.Sort(this)

	height := make([]int, this.Len())

	var getHeight func(int, int) int
	getHeight = func(currentBox, nextBox int) int {
		if nextBox >= this.Len() {
			return 0
		}

		// Include the next box.
		heightWithNextBox := 0
		if currentBox == -1 || this.canStack(nextBox, currentBox) {
			if height[nextBox] == 0 {
				height[nextBox] = getHeight(nextBox, nextBox+1)
				height[nextBox] += this[nextBox].height
			}
			heightWithNextBox = height[nextBox]
		}

		// Skip the next box.
		heightWithoutNextBox := getHeight(currentBox, nextBox+1)

		// Compare the two heights.
		if heightWithNextBox > heightWithoutNextBox {
			return heightWithNextBox
		} else {
			return heightWithoutNextBox
		}
	}

	return getHeight(-1, 0)
}

func main() {
	boxes := Boxes{{3, 3, 4}, {1, 2, 1}, {4, 5, 2}, {2, 2, 2}}
	fmt.Println(boxes)
	fmt.Println("Max box stack height =", boxes.Stack())
}
