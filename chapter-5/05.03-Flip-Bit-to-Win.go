package main

import "fmt"

func Flip(number int) (maxSequence int, flipIndex int, bitCounts []int) {
	maxSequence = 1

	if number == 0 {
		flipIndex = 1
		return
	}

	bitCounts = make([]int, 0)
	currentBit, currentCount := 0, 0

	// Prepare alternating bit (0s, 1s, 0s, 1s...) counts.
	for number != 0 {
		if number & 1 != currentBit {
			// Push the previous count to the array.
			bitCounts = append(bitCounts, currentCount)

			currentCount = 0	// Reset the current count.
			currentBit ^= 1		// Flip the current bit.
		}

		currentCount++
		number >>= 1
	}

	// Push the last count too.
	bitCounts = append(bitCounts, currentCount)

	for i, curIndex := 0, 0; i < len(bitCounts); i += 2 {
		zeroCount := bitCounts[i]
		leftOnes, rightOnes := 0, 0
		if i-1 > 0 {
			rightOnes = bitCounts[i-1]
		}

		if i+1 < len(bitCounts) {
			leftOnes = bitCounts[i+1]
		}

		curIndex += (zeroCount + rightOnes)

		if zeroCount == 1 {
			if leftOnes+rightOnes+zeroCount > maxSequence {
				// The flip index is the index of this zero.
				maxSequence = leftOnes+rightOnes+zeroCount
				flipIndex = curIndex
			}
		} else {
			if leftOnes+1 > maxSequence {
				// The flip index is the next to the ending of
				// left ones.
				maxSequence = leftOnes+1
				flipIndex = curIndex
			}

			if rightOnes+1 > maxSequence {
				// The flip index is the previous of starting
				// of right ones.
				maxSequence = rightOnes+1
				flipIndex = curIndex-zeroCount+1
			}
		}
	}

	return
}

func main() {
	for _, n := range []int{0, 14, 15, 1775, 225760, 225776, 225788} {
		seq, idx, count := Flip(n)
		fmt.Printf("%32b: Max sequence %2d @ index %2d. Map = %v\n",
			n, seq, idx, count)
	}
}
