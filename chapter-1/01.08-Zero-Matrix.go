package main

import "fmt"

func ZeroMatrix(matrix [][]int) {
	if matrix == nil {
		return
	}

	rows, cols := len(matrix), len(matrix[0])

	// Nothing to be done if there is only one cell in the matrix.
	if rows == 1 && cols == 1 {
		return
	}

	zeroRows, zeroCols := make([]bool, rows), make([]bool, cols)

	// Track rows and columns that need to be zeroed out.
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				zeroRows[r] = true
				zeroCols[c] = true
			}
		}
	}

	// Zero-out the rows.
	for r := 0; r < rows; r++ {
		if !zeroRows[r] {
			continue
		}

		for c := 0; c < cols; c++ {
			matrix[r][c] = 0
		}
	}

	// Zero-out the columns.
	for c := 0; c < cols; c++ {
		if !zeroCols[c] {
			continue
		}

		for r := 0; r < rows; r++ {
			matrix[r][c] = 0
		}
	}
}

func main() {
	m1 := [][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}
	fmt.Print(m1, " -> ")
	ZeroMatrix(m1)
	fmt.Println(m1)

	m2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}}
	fmt.Print(m2, " -> ")
	ZeroMatrix(m2)
	fmt.Println(m2)
}
