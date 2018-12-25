package main

import "fmt"

func IsPalindromeCandidate(str string) bool {
	if len(str) == 0 {
		return false
	}

	// A map to keep track of each unique character frequency.
	count := make(map[rune]int)

	for _, ch := range str {
		count[ch]++
	}

	oddLength := len(str)%2 == 1

	foundOdd := false
	for _, freq := range count {
		// If the string length is odd and second character with
		// odd frequency found, then not palindrome.
		if freq % 2 == 1 {
			if !oddLength || foundOdd {
				return false
			}

			foundOdd = true
		}
	}

	return true
}

func main() {
	str := "tactcoa"
	fmt.Printf("isPalindromePermutation(%v) = %v\n", str, IsPalindromeCandidate(str))
}
