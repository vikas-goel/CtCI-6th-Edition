// Write a method to compute all permutations of a string whose characters are
// not necessarily unique. The list of permutations should not have duplicates.

package main

import (
	"fmt"
	"log"
	"os"
)

func Permutations(set []byte) {
	setLength := len(set)
	if setLength < 2 {
		return
	}

	charMap := make(map[byte]int)
	for i := 0; i < setLength; i++ {
		charMap[set[i]]++
	}

	chars := make([]byte, 0)
	for key, _ := range charMap {
		chars = append(chars, key)
	}

	var prepare func([]byte, int)
	prepare = func(prefix []byte, prefixLength int) {
		if prefixLength == setLength {
			// One permuatation is complete.
			fmt.Printf("%c\n", prefix)
			return
		}

		for i := 0; i < len(chars); i++ {
			if charMap[chars[i]] == 0 {
				continue
			}

			// Fix a character and then recurse for remaining.
			charMap[chars[i]]--
			prefix[prefixLength] = chars[i]
			prepare(prefix, prefixLength+1)
			charMap[chars[i]]++
		}
	}

	prepare(make([]byte, setLength), 0)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "<string>")
	}

	orig := []byte(os.Args[1])
	Permutations(orig)
}
