// Imagine a robot sitting on the upper left corner of grid with r rows and c
// columns. The robot can only move in two directions, right and down, but
// certain cells are "off limits" such that the robot cannot step on them.
// Design an algorithm to find a path for the robot from the top left to the
// bottom right.

package main

import "fmt"

type Cell struct {
	R, C int
}

func RobotPath(grid [][]bool) (path []Cell) {
	if grid == nil || len(grid) == 0 {
		return
	}

	unreachable := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		unreachable[i] = make([]bool, len(grid[0]))
	}

	path = make([]Cell, 0, len(grid)+len(grid[0])-1)

	origin := func(row, col int) bool {
		return row == 0 && col == 0
	}

	var getPath func(int, int) bool
	getPath = func(row, col int) bool {
		if row<0 || col<0 || !grid[row][col] || unreachable[row][col] {
			return false
		}

		if origin(row,col) || getPath(row-1,col) || getPath(row,col-1) {
			path = append(path, Cell{row, col})
			return true
		}

		unreachable[row][col] = true
		return false
	}

	getPath(len(grid)-1, len(grid[0])-1)
	return
}

func main() {
	matrix := [][]bool{
		{true, true, false, true},
		{true, true, false, true},
		{false, true, true, false},
		{true, false, true, true},
		{true, true, false, true},
		{false, true, true, true}}
	fmt.Printf("Grid:%v\n%v\n", matrix, RobotPath(matrix))
}
