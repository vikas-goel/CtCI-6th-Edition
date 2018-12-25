package main

import "fmt"
import "strings"

var isSubstring = strings.Contains

func IsRotation(str, of string) bool {
	if len(str) != len(of) {
		return false
	}

	// The 'of' string must be subset of the string when joined the
	// rotated string to itself.
	return isSubstring(str+str, of)
}

func main() {
	orig := "waterbottle"
	str1 := "erbottlewat"
	str2 := "erbottlewar"
	fmt.Printf("isRotated(%v, %v) = %v\n", str1, orig, IsRotation(str1, orig))
	fmt.Printf("isRotated(%v, %v) = %v\n", str2, orig, IsRotation(str2, orig))
}
