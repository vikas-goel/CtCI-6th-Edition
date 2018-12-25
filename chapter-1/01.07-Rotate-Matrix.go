package main

import "fmt"

func RotateMatrix(matrix [][]int) bool {
	if matrix == nil {
		return false
	}

	size, cols := len(matrix), len(matrix[0])

	// Ensure it is a square matrix.
	if size != cols {
		return false
	}

	// Nothing to be done if there is only one cell in the matrix.
	if size == 1 {
		return true
	}

	// Each outer iteration will handle start and endtom size of the
	// respective one. So, need to run only half of the sizeension.
	for start, end := 0, size-1; start < size/2; start++ {
		// Each inner iteration will handle left and right columns of
		// the respective one.
		for i := 0; i < end-start; i++ {
			// Clockwise 90-degree rotation.
			first := matrix[start][start+i]
			matrix[start][start+i] = matrix[end-i][start]
			matrix[end-i][start] = matrix[end][end-i]
			matrix[end][end-i] = matrix[start+i][end]
			matrix[start+i][end] = first
		}
		end--
	}

	return true
}

func main() {
	m1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Print(m1, " -> ")
	RotateMatrix(m1)
	fmt.Println(m1)

	m2 := [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}, {13,14,15,16}}
	fmt.Print(m2, " -> ")
	RotateMatrix(m2)
	fmt.Println(m2)
}
