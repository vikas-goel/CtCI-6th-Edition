package main

import "fmt"

func CheckPermutation(str1, str2 string) bool {
	// Two different length strings cannot be permutation.
	if len(str1) != len(str2) {
		return false
	}

	chars := make(map[rune]int)

	// Create a map of count of characters in the first string.
	for _, ch := range str1 {
		chars[ch]++
	}

	// Keep track of the counts in the second string.
	// If the count of a character drops below 0, then there is a mismatch
	// in character or occurences of that.
	for _, ch := range str2 {
		chars[ch]--
		if chars[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	str1, str2 := "doge", "gode"
	fmt.Printf("(%v,%v) = %v\n", str1, str2, CheckPermutation(str1, str2))
}
