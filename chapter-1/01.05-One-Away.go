package main

import "fmt"

func IsOneEditAway(str1, str2 string) bool {
	len1, len2 := len(str1), len(str2)
	if len1 == len2 {
		return isOneReplaceAway(str1, str2)
	} else if len1-1 == len2 {	// str2 is shorter.
		return isOneInsertAway(str2, str1)
	} else {			// str1 is shorter.
		return isOneInsertAway(str1, str2)
	}

	return false
}

func isOneReplaceAway(str1, str2 string) bool {
	foundDiff := false
	str1Byte, str2Byte := []byte(str1), []byte(str2)

	for i, length := 0, len(str1); i < length; i++ {
		if str1Byte[i] != str2Byte[i] {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
	}

	return true
}

func isOneInsertAway(str1, str2 string) bool {
	str1Byte, str2Byte := []byte(str1), []byte(str2)

	for idx1, idx2, length := 0, 0, len(str1); idx1 < length; idx1++ {
		if str1Byte[idx1] != str2Byte[idx2] {
			if idx1 != idx2 {
				// Second edit needed.
				return false
			}
			// First edit found.
			idx2++
		}
		idx2++
	}

	return true
}

func main() {
	str1, str2 := "apple", "aple"
	fmt.Printf("isOneEditAway(%v,%v) = %v\n", str1, str2, IsOneEditAway(str1, str2))
}
