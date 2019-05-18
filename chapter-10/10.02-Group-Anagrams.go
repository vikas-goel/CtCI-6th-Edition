// Write a method to sort an array of strings so that all the anagrams are next
// to each other.

package main

import (
	"fmt"
	"sort"
)

type AnagramsStrings []string

func (this AnagramsStrings) Len() int {
	return len(this)
}

// Treat each anagram as the same element.
func (this AnagramsStrings) Less(i, j int) bool {
	// If length of the two strings are not equal, then return the result
	// based on their lengths.
	if len(this[i]) < len(this[j]) {
		return true
	} else if len(this[i]) > len(this[j]) {
		return false
	}

	// Sort the strings and compare each character one by one.
	iBytes, jBytes := []byte(this[i]), []byte(this[j])
	iInts := make([]int, len(iBytes))
	jInts := make([]int, len(jBytes))

	for i := 0; i < len(iBytes); i++ {
		iInts[i] = int(iBytes[i])
		jInts[i] = int(jBytes[i])
	}

	sort.Ints(iInts)
	sort.Ints(jInts)

	for i := 0; i < len(iInts); i++ {
		if iInts[i] < jInts[i] {
			return true
		} else if iInts[i] > jInts[i] {
			return false
		}
	}

	return false
}

func (this AnagramsStrings) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	array := []string{"abc", "def", "abcd", "efgh", "abd", "bca", "fghe", "fed", "pqrs", "hefg", "cba", "hegf"}

	fmt.Println(array)
	sort.Sort(AnagramsStrings(array))
	fmt.Println(array)
}
