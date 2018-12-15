package main

import "fmt"

func URLify(str []byte, trueLength int)  {
	if str == nil || trueLength == 0 {
		return
	}

	spaceCount := 0

	// Count number of spaces in the actual string.
	for i := 0; i < trueLength; i++ {
		if str[i] == ' ' {
			spaceCount++
		}
	}

	// Ending position in the modified string.
	lastIndex := trueLength + 2*spaceCount
	if lastIndex < len(str) {
		str[lastIndex] = 0
	}

	// Stop iteration when the pointers of actual & modified string meet.
	for i := trueLength-1; i != lastIndex-1; i-- {
		if str[i] == ' ' {
			// Replace space character with %20.
			str[lastIndex-1] = '0'
			str[lastIndex-2] = '2'
			str[lastIndex-3] = '%'
			lastIndex -= 3
		} else {
			// Copy non-space character as-is.
			str[lastIndex-1] = str[i]
			lastIndex -= 1
		}
	}
}

func main() {
	str1, len1 := "Mr John Smith      ", 13
	byteStr := []byte(str1)
	URLify(byteStr, len1)
	fmt.Printf("URLify(%v) = %v\n", str1, string(byteStr))
}
