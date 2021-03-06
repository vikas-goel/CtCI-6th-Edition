// Given an infinite number of quarters (25 cents), dimes (10 cents), nickels
// (5 cents), and pennies (1 cent), write code to calculate the number of ways
// of representing n cents.

package main

import "fmt"

func Count(denoms []int, cents int) int {
	numDenoms := len(denoms)
	ways := make([][]int, cents+1)

	for i := 0; i <= cents; i++ {
		ways[i] = make([]int, numDenoms)
	}

	// Zero cent can be done in 1 way only irrespective of denominations.
	for i := 0; i < numDenoms; i++ {
		ways[0][i] = 1
	}

	for i := 1; i <= cents; i++ {
		for j := 0; j < numDenoms; j++ {
			// Include denom[j] to count the ways.
			if i - denoms[j] >= 0 {
				ways[i][j] += ways[i-denoms[j]][j]
			}

			// Exclude denom[j] to count the ways.
			if j >= 1 {
				ways[i][j] += ways[i][j-1]
			}
		}
	}

	return ways[cents][numDenoms-1]
}

func main() {
	cents := 100
	denoms := []int{1, 5, 10, 25}

	fmt.Println("Coins", denoms)
	fmt.Println("Change count for", cents, "=", Count(denoms, 100))
}
