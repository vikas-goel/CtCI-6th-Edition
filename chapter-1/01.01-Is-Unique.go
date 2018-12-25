package main

import "fmt"

// Function to check whether each character in the string appears exactly once.
func IsUnique(str string) bool {
	if len(str) < 2 {
		return true
	}

	// Create a hash set of the characters.
	chars := make(map[rune]struct{})

	for _, ch := range str {
		if _, ok := chars[ch]; ok {
			// Duplication detected.
			return false
		}

		// Add the character in the hash set.
		chars[ch] = struct{}{}
	}

	return true
}

func main() {
	string1 := "abcdefghi"
	string2 := "abcdeeghi"

	fmt.Printf("isUnique(%v) = %v\n", string1, IsUnique(string1))
	fmt.Printf("isUnique(%v) = %v\n", string2, IsUnique(string2))
}
