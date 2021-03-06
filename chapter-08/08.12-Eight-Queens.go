// Write an algorithm to print all ways of arranging eight queens on an 8x8
// chess board so that none of them share the same row, column, or diagonal. In
// this case, "diagonal" means all diagonals, not just the two that bisect the
// board.

package main

import (
	"fmt"
	"time"
)

func CandidateBoards(board []int, start, end int, total, valid *int) {
	if start == end {
		*total++

		if validBoard(board) {
			fmt.Println(board)
			*valid++
		}
	} else {
		for i := start; i <= end; i++ {
			swapPositions(board, start, i)
			CandidateBoards(board, start+1, end, total, valid)
			swapPositions(board, start, i)
		}
	}
}

func validBoard(board []int) bool {
	length := len(board)
	for i := length-8; i < length-1; i++ {
		for j := i+1; j < length; j++ {
			if ! validPositions(i, j, board[i], board[j]) {
				return false
			}
		}
	}

	return true
}

func validPositions(row1, row2, col1, col2 int) bool {
	row := int8(row1 - row2)
	col := int8(col1 - col2)

	if (row == col) || (row == -1 * col) {
		return false
	}

	return true
}

func swapPositions(board []int, row1, row2 int) {
	board[row1], board[row2] = board[row2], board[row1]
}

func main() {
	initBoard := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	total, valid := 0, 0
	start := time.Now()

	low, high := len(initBoard)-8, len(initBoard)-1

	CandidateBoards(initBoard, low, high, &total, &valid)

	fmt.Println("Valid boards:", valid, "of", total)
	fmt.Println("Execution time:", time.Since(start))
}
